test:
	go test -v ./apis/imageshack/
	go test -v ./apis/imgur/
build:
	go build
clean:
	rm -rf *.png
	rm -rf *.jpg
run:
	./ImageCompress
