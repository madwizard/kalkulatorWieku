.PHONY: all build clean

all: build

build:
	fyne-cross windows -arch=amd64 --app-id "github.com/madwizard/kalkulatorWieku"

clean:
	rm -rf fyne-cross/

run:
	go run ./kalkulator.go