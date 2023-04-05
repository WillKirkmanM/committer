path:
	go build && mv committer ~/.local/bin/

cleanPath:
	rm ~/.local/bin/committer

clean:
	rm committer

production:
	go build -ldflags="-s -w"

build:
	go build
