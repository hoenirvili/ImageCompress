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

package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hoenirvili/ImageCompress/internal"
)

// Get request to imgurl api with setting the mime content and url
// returns a pointer to http.Reponse and in case of error  a pointer to ErrorStat
func (i Imgur) Get(url string) (*http.Response, *internal.ErrorStat) {
	req, err := http.NewRequest("GET", url, new(bytes.Buffer))
	if err != nil {
		return nil, &internal.ErrorStat{Message: "Can't create new request to Imgurl"}
	}
	// add the corresponding header
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", i.clientID))
	//aloc new *http.Client
	client := &http.Client{}

	//make client reqeust with the http.Request above declared.
	resp, err := client.Do(req)
	if err != nil {
		return nil, &internal.ErrorStat{Message: "Get request to Imgurl api failed"}
	}

	// check http status to be ok(200)
	errStatHTTP := errorHTTPStatus(resp.StatusCode)
	if errStatHTTP != nil {
		return nil, errStatHTTP
	}

	// return it
	return resp, nil
}

// ImageJSON returns serialez get response
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageJSON(url string) (*Image, error) {
	resp, errGet := i.Get(url)
	if errGet != nil {
		return nil, errGet
	}

	imgJSON := &Image{}
	errJSON := json.NewDecoder(resp.Body).Decode(imgJSON)
	if errJSON != nil {
		return nil, errJSON
	}
	return imgJSON, nil
}

// ImageByte http request and return image serialized body
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageByte(url string) ([]byte, error) {
	resp, errGet := i.Get(url)
	if errGet != nil {
		return nil, errGet
	}

	readed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return readed, nil
}
