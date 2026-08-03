ifeq ($(OS),Windows_NT)
    TARGET := sca.exe
else
    TARGET := sca
endif

sca:
	go build -o build/$(TARGET) .
