build:
	go build -o output/nsregistry cmd/nsregistry/main.go

clean:
	rm -rf output
run:
	go run cmd/nsregistry/main.go