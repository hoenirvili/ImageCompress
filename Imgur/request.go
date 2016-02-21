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

package Imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Basic request struct for api
// Unspecified elem of struct is nil
// But for consistency and clearer
var basicRequest = &http.Request{
	Method:           "",
	URL:              nil,
	Proto:            "HTTP 2.0",
	ProtoMajor:       2,
	ProtoMinor:       0,
	Header:           make(http.Header),
	Body:             nil,
	ContentLength:    0,
	TransferEncoding: nil,
	Close:            false,
	Host:             "api.imgur.com",
	Form:             nil,
	PostForm:         nil,
	MultipartForm:    nil,
	Trailer:          nil,
	RemoteAddr:       "",
	TLS:              nil,
	Cancel:           nil,
}

// NewRequest func return a *http.Request
// that can be used with Client.Do() func
func (i Imgur) NewRequest(methodRequest, urlRequest, mimeRequest string) (*http.Request, *ErrorStat) {
	u, err := url.Parse(urlRequest)

	if err != nil {
		return nil, &ErrorStat{msg: fmt.Sprintf("Can't parse this url %s", urlRequest)}
	}

	rc, ok := i.body.(io.ReadCloser)

	if !ok && i.body != nil {
		rc = ioutil.NopCloser(i.body)
	}

	// Attach all basic config for request
	basicRequest.Method = methodRequest
	basicRequest.URL = u
	basicRequest.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", i.clientID))
	basicRequest.Header.Add("Content-Type", mimeRequest)
	basicRequest.Body = rc

	// Note that in go you don't need break stmt
	if i.body != nil {
		switch v := i.body.(type) {
		case *bytes.Buffer:
			basicRequest.ContentLength = int64(v.Len())
		case *bytes.Reader:
			basicRequest.ContentLength = int64(v.Len())
		case *strings.Reader:
			basicRequest.ContentLength = int64(v.Len())
		}
	}

	// return it
	return basicRequest, nil
}

// Get request to imgurl api with setting the mime content and url
// returns a pointer to http.Reponse and in case of error  a pointer to ErrorStat
func (i Imgur) Get(url, mime string) (resp *http.Response, errStat *ErrorStat) {
	request, err := i.NewRequest("GET", url, mime)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, errReq := client.Do(request)

	err = errorStatus(response.StatusCode)
	if err != nil {
		return nil, err
	}
	// other error
	if errReq != nil {
		return nil, &ErrorStat{msg: fmt.Sprintf("Error : %s\n", errReq.Error())}
	}

	return response, nil
}

// ImageJSON returns serialez get response
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageJSON(url string) (retI *Image) {
	mime := "application/json"
	resp, err := i.Get(url, mime)
	if err != nil {
		err.Print()
		return nil
	}

	img := &Image{}

	errJSON := json.NewDecoder(resp.Body).Decode(img)
	if errJSON != nil {
		log.Fatal(errJSON)
	}

	return img
}

// ImageByte http request and return image serialized body
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageByte(url string) (byteBody []byte) {
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

//SaveImage saves images
func (i Imgur) SaveImage(url, path string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	readed, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	errWrite := ioutil.WriteFile(path, readed, 0644)

	if errWrite != nil {
		log.Fatal(errWrite)
	}

	defer resp.Body.Close()
}
