# Unit test
test:
	go test -v ./apis/imageshack/
	go test -v ./apis/imgur/
# Build the project
build:
	go build
# Clean from runing api tests.
clean:
	rm -rf *.png
	rm -rf *.jpg
# Run the server
run:
	go build
	./ImageCompress

# Optimal build
optbin:
	# The -s ldflag will omit the symbol table and debugging information when building your executable.
	go build -o gobin -ldflags=-s
