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

package imageshack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hoenirvili/ImageCompress/internal"
)

// Get request
func (i ImageShack) Get(url string) (*http.Response, error) {

	// make new pointer to *http.Resquest
	// with GET method, url and a new byes.Buffer body
	req, err := http.NewRequest("GET", url, new(bytes.Buffer))
	if err != nil {
		return nil, &internal.ErrorStat{Message: "Can't create new GET request to Image Shack"}
	}

	// alloc new *http.Client
	client := &http.Client{}

	// make client request with the http.Request above declared
	resp, err := client.Do(req)
	if err != nil {
		return nil, &internal.ErrorStat{Message: "GET request to ImageShack api failed"}
	}

	if resp.StatusCode > http.StatusPartialContent {
		return nil, &internal.ErrorStat{Message: fmt.Sprintf("%s %s", resp.Status, "POST request to ImageShack api failed")}
	}

	// return it
	return resp, nil
}

// ImageJSON request imageshack
func (i ImageShack) ImageJSON(url string) (*Image, error) {

	resp, err := i.Get(url)
	if err != nil {
		return nil, err
	}

	// aloc imgJSON struct for decoding the response body
	imgJSON := &Image{}

	// decode into JSON
	err = json.NewDecoder(resp.Body).Decode(imgJSON)
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

	// return the JSON
	return imgJSON, nil
}

// ImageByte reads images from http response
// and serialize it into byte
func (i ImageShack) ImageByte(url string) ([]byte, error) {
	resp, err := i.Get(url)
	if err != nil {
		return nil, err
	}

	readed, err := ioutil.ReadAll(resp.Body)
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
	return readed, nil
}
