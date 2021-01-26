.PHONY: build clean install

build:
	go build -o go2uml

clean:
	rm -f go2uml

install: build
	@cp -f go2uml $$GOPATH/bin/