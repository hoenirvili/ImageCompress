# Unit test
ROOT_PATH=/home/hoenir/Work/Go/src/github.com/hoenirvili/ImageCompress

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

sass:
	sass --watch $(ROOT_PATH)/static/sass/bundle.sass:$(ROOT_PATH)/static/css/bundle.css
