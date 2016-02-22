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
