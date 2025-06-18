package main

import (
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/dig"
)

type FourDimensionalPocket interface {
	GetTakecopter() (string, error)
	GetDokodemodoa() (string, error)
}

type Doraemon struct{}

func (d Doraemon) GetTakecopter() (string, error) {
	return "タケコプター🚁", nil
}

func (d Doraemon) GetDokodemodoa() (string, error) {
	return "どこでもドア🚪", nil
}

func NewDoraemon() FourDimensionalPocket {
	return Doraemon{}
}

func NewLogger() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return slog.New(handler)
}

type Bird interface {
	fly() string
}

type Nobita struct {
	pocket FourDimensionalPocket
	logger *slog.Logger
}

func NewNobita(pocket FourDimensionalPocket, logger *slog.Logger) Bird {
	return Nobita{
		pocket: pocket,
		logger: logger,
	}
}

func (n Nobita) fly() string {
	takecopter, err := n.pocket.GetTakecopter()
	if err != nil {
		n.logger.Error("タケコプターの取得に失敗", "error", err)
		return "飛べません🐧"
	}
	n.logger.Info("タケコプターを取得しました")
	return fmt.Sprintf("%sを使って飛んでいく!", takecopter)
}

func main() {
	container := dig.New()

	container.Provide(NewLogger)
	container.Provide(NewDoraemon)
	container.Provide(NewNobita)

	err := container.Invoke(func(b Bird) {
		fmt.Println(b.fly())
	})

	if err != nil {
		slog.New(slog.NewTextHandler(os.Stderr, nil)).Error("DIに失敗", "error", err)
	}
}
