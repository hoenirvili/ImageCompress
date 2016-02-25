# ImageCompress

[![Build Status](https://travis-ci.org/hoenirvili/ImageCompress.svg?branch=master)](https://travis-ci.org/hoenirvili/ImageCompress)

###### A simple image compression service written in Go.


![gopher image](doc/gopher.png)


## About

This service is based on 3 apis.

**Imgur** api and **ImageShack** api where you can select what image to compress by passing the link to the service the image will be sent after to **TinyPNG** for compression.

#### Format
For now the service supports **JPG** and **PNG** formats.


Sizes in bytes.

```
└───> stat CFzq6zN.jpg
	File: 'CFzq6zN.jpg'
	Size: 152015    	Blocks: 312        IO Block: 4096   regular file

└───> stat CFzq6zNCompressed.jpg
	File: 'CFzq6zNCompressed.jpg'
	Size: 72360     	Blocks: 144        IO Block: 4096   regular file
```
