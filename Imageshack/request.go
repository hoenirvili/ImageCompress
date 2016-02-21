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

package Imageshack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Get request
func (i ImageShack) Get(url, mime string) (resp *http.Response, respErr error) {
	request, err := http.NewRequest("GET", url, new(bytes.Buffer))
	request.Header.Add("Content-Type", mime)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// ImageJSON request imageshack
func (i ImageShack) ImageJSON(url string) (apiJSON *ImageShackJSON) {
	body := new(bytes.Buffer)
	req, err := http.NewRequest("GET", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	imgJSON := &ImageShackJSON{}

	err = json.NewDecoder(resp.Body).Decode(imgJSON)
	if err != nil {
		log.Fatal(err)
	}

	return imgJSON
}

// ImageByte reads images from http response and serialize it
// into byte
func (i ImageShack) ImageByte(url string) (bodyByte []byte) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	readed, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	return readed
}
