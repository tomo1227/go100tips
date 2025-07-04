package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
)

type ElasticsearchHandler struct {
	client *elasticsearch.Client
	level  slog.Level
}

func NewElasticsearchHandler(level slog.Level) *ElasticsearchHandler {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		panic(err)
	}
	return &ElasticsearchHandler{client: es, level: level}
}

func (h *ElasticsearchHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *ElasticsearchHandler) Handle(_ context.Context, record slog.Record) error {
	entry := map[string]interface{}{
		"time":    record.Time.Format(time.RFC3339),
		"level":   record.Level.String(),
		"message": record.Message,
	}
	record.Attrs(func(attr slog.Attr) bool {
		entry[attr.Key] = attr.Value.Any()
		return true
	})

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	res, err := h.client.Index(
		"go-logs",
		bytes.NewReader(data),
		h.client.Index.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch error: %s", res.String())
	}

	return err
}

func (h *ElasticsearchHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *ElasticsearchHandler) WithGroup(_ string) slog.Handler {
	return h
}
