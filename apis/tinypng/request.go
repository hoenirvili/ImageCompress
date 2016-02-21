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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Base64 encoding defined by rfc RFC2045-MIME
// plus concat the "Basic: "word.
func sanitizeHeader(username, key string) (header string) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s", username, key)))
	return fmt.Sprintf("Basic %s", sEnc)
}

// Post request to tiny
func (t *Tiny) Post(mime string) {
	//serialized body for post request
	postBytesReader := bytes.NewReader(t.body)

	// new *http.Request{}
	request, err := http.NewRequest("POST", urlSender, postBytesReader)
	if err != nil {
		log.Fatal(err)
	}
	//sanize header for auth header
	authValue := sanitizeHeader(username, key)
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

// func (t Tiny) Get(url string) *http.Response
