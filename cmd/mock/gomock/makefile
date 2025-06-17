.PHONY: mock
mock: 
	go install go.uber.org/mock/mockgen@latest
	mockgen -source=main.go -destination=mock/mock_main.go -exclude_interfaces=Bird -package=mock
