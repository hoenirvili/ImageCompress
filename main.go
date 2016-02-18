// Copyright [2016] [hoenir]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hoenirvili/ImageCompress/Imageshack"
	"github.com/hoenirvili/ImageCompress/Imgur"
	"github.com/hoenirvili/ImageCompress/Tinypng"
)

const (
	clientID     = "40aea5f08c0f717"
	clientSecret = "a72c35d27b38d27114b4503e5b9acc835861ed8c"
)

func concat(first, second string) (string, error) {
	//init
	nCopied := 0

	// get lens
	lenFirst := len(first)
	lenSecond := len(second)

	// sum of both lens
	n := lenFirst + lenSecond

	// alloc holder
	holder := make([]byte, n)

	// copy the first holder from pos nCopied:=0
	nCopied = copy(holder[nCopied:], first)

	// if everything is ok copy the second one.
	if nCopied != n-lenSecond {
		return "", errors.New("Can't copy first string")
	}

	// coy the second holder from pos nCopied:= previousNcopied
	nCopied = copy(holder[nCopied:], second)

	if nCopied != n-lenFirst {
		return "", errors.New("Can't copy the second string")
	}

	// TODO find a better way to convert it to string
	return string(holder), nil
}

// dumpResponse for debugging informatin
func dumpResponse(resp *http.Response) {
	for key, val := range resp.Header {
		fmt.Print(key)
		fmt.Print(" : ")
		fmt.Println(val)
	}
	fmt.Printf("Status : %s\n", resp.Status)
	readed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("===================== BODY ====================")
	fmt.Println()
	fmt.Println(string(readed))

	// for any case
	defer resp.Body.Close()
}

// imgurToTiny just send image to tinyPNG
func imgurToTiny() {
	// alloc
	imgur := Imgur.NewImgur()
	tiny := Tinypng.NewTiny()
	// set
	imgur.SetClientID(clientID)
	imgur.SetClientSecret(clientSecret)
	imgur.SetBody(new(bytes.Buffer))
	// get request JSON response
	v := imgur.ImageJSON("https://api.imgur.com/3/gallery/image/i0xn0Dx")
	if v.Data.Type == "image/png" || v.Data.Type == "image/jpeg" {
		byteImage := imgur.ImageByte(v.Data.Link)
		tiny.SetBody(byteImage)
		tiny.Post(v.Data.Type)
	} else {
		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
		os.Exit(1)
	}
}
func shackToTiny() {
	// alloc
	shack := Imageshack.NewImageShack()
	tiny := Tinypng.NewTiny()
	v := shack.ImageJSON("https://api.imageshack.com/v2/images/pbzPCsEij")
	url, err := concat("https://", v.Result.Direct_link)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(url, ".jpg") {
		byteImage := shack.ImageByte(url)
		tiny.SetBody(byteImage)
		tiny.Post("image/jpeg")
	} else if strings.Contains(url, ".png") {
		byteImage := shack.ImageByte(url)
		tiny.SetBody(byteImage)
		tiny.Post("image/png")
	} else {
		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
		os.Exit(1)
	}
}
func main() {

	imgurToTiny()
	shackToTiny()

}
