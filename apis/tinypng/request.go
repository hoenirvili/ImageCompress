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

package tinypng

// Post send http post and parse the response
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hoenirvili/ImageCompress/internal"
	"github.com/hoenirvili/ImageCompress/netproto"
	"github.com/hoenirvili/ImageCompress/utils"
)

// Post request to tiny
func (t Tiny) Post(mime string) (*http.Response, error) {
	//serialized body for post request
	postBytesReader := bytes.NewReader(t.body)

	// new *http.Request{}
	request, err := http.NewRequest("POST", urlSender, postBytesReader)
	if err != nil {
		return nil, &internal.ErrorStat{Message: "Can't create new POST request to Tiny"}
	}

	//sanize header for auth header
	authValue := netproto.HeaderToBase64(username, key)
	request.Header.Add("Content-Type", mime)
	request.Header.Add("Authorization", authValue)

	// new *http.Client struct
	cli := &http.Client{}

	// make POST request to tiny api
	resp, err := cli.Do(request)
	if err != nil {
		return nil, &internal.ErrorStat{Message: "POST request to Tiny api failed"}
	}

	if resp.StatusCode > http.StatusPartialContent {
		return nil, &internal.ErrorStat{Message: fmt.Sprintf("%s %s", resp.Status, "POST request to Tiny api failed")}
	}
	// for debugging reasons
	fmt.Printf("%+v\n", resp)

	return resp, nil
}

// Get image from tiny api
func (t Tiny) Get(url string) ([]byte, error) {
	respImage, errGET := http.Get(url)
	if errGET != nil {
		return nil, internal.ErrorStat{Message: fmt.Sprintf("Error on geting image from Tiny api")}
	}
	//todo
	readedIMG, errREAD := ioutil.ReadAll(respImage.Body)
	if errREAD != nil {
		return nil, internal.ErrorStat{Message: fmt.Sprintf("Error on reading response body image from Tiny api")}
	}

	func() {
		err := respImage.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return readedIMG, nil
}

// PostGetJSON funcition that sends image and recives json
func (t Tiny) PostGetJSON(mime string) (*TinyJSON, error) {
	resp, err := t.Post(mime)
	if err != nil {
		return nil, err
	}

	// allc TinyJSON struct for parsing JSON response
	v := &TinyJSON{}

	// parse JSON from response
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	// close body response
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return v, nil
}

// //ImageJSON get json response from sending image to tiny api
// func (t Tiny) ImageJSON(url string) (*TinyJSON, error) {
// 	resp, err := t.Post("application/json")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// allc TinyJSON struct for parsing JSON response
// 	v := &TinyJSON{}
//
// 	// parse JSON from response
// 	err = json.NewDecoder(resp.Body).Decode(v)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// close body response
// 	defer func() {
// 		err = resp.Body.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
//
// 	return v, nil
// }

// SaveImage saves the new image compressed
func (t Tiny) SaveImage(img []byte, path, mime string) error {

	if mime == "image/jpeg" {
		//Write compressed file
		nameJPG, err := utils.Concat(path, ".jpg")
		if err != nil {
			nameJPG = "defaultCompressed.jpg"
		}
		errWrite := ioutil.WriteFile(nameJPG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	} else if mime == "image/png" {
		namePNG, err := utils.Concat(path, ".png")
		if err != nil {
			namePNG = "defaultCompressed.png"
		}
		errWrite := ioutil.WriteFile(namePNG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	}
	return nil
}

// Post request to tiny
func (t *Tiny) PostTry(mime string) {
	//serialized body for post request
	postBytesReader := bytes.NewReader(t.body)

	// new *http.Request{}
	request, err := http.NewRequest("POST", urlSender, postBytesReader)
	if err != nil {
		log.Fatal(err)
	}
	//sanize header for auth header
	authValue := netproto.HeaderToBase64(username, key)
	request.Header.Add("Content-Type", mime)
	request.Header.Add("Authorization", authValue)
	// alloc clinet struct
	// default
	cli := &http.Client{}

	// make request
	resp, err := cli.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)

	// alog TinyJSON struct for parsing JSON response
	v := &TinyJSON{}

	// parse JSON from response
	errJSON := json.NewDecoder(resp.Body).Decode(v)
	if errJSON != nil {
		log.Fatal(errJSON)
	}
	//===========================================================
	//get request the link image from JSON response
	respImage, errGET := http.Get(v.Output.URL)
	if errGET != nil {
		log.Fatal(errGET)
	}
	// read all bytes body from getRequest
	readedIMG, errREAD := ioutil.ReadAll(respImage.Body)
	if errREAD != nil {
		log.Fatal(errREAD)
	}
	fmt.Println()
	fmt.Printf("%+v\n", respImage)
	if mime == "image/jpeg" {
		//Write compressed file
		errWrite := ioutil.WriteFile("imageCompressed.jpg", readedIMG, 0644)
		if errWrite != nil {
			log.Fatal(errWrite)
		}
	} else if mime == "image/png" {
		errWrite := ioutil.WriteFile("imageCompressed.png", readedIMG, 0644)
		if errWrite != nil {
			log.Fatal(errWrite)
		}
	}

	// close all Bodys from all requests
	defer resp.Body.Close()
	defer respImage.Body.Close()
}
