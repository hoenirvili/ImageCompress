all:
	go build
	./ImageCompress

clean:
	rm -rf *.png
	rm -rf *.jpg

