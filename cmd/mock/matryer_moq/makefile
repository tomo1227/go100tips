.PHONY: mock
mock: 
	go install github.com/matryer/moq@latest
	moq -out mock/mock_main.go -pkg mock . FourDimensionalPocket
