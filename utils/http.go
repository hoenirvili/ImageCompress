// Copyright [2016] [hoenir]
//
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

package utils

import (
	"io"
	"log"
	"net/http"
)

// ResponseStringReader reads resp.Body *http.Response
// into a string containter and closes the resp.Body.Close()
func ResponseStringReader(resp *http.Response) (stringContainer string, err error) {
	var (
		buffer    []byte
		n         int
		container []byte
	)
	// declare buffer to be 1kb
	buffer = make([]byte, 1024)
	for {
		n, err = resp.Body.Read(buffer)
		// test if err
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return "", err
		}

		// apend buffer to container
		container = append(buffer)
	}

	return string(container), nil
}

// ResponseByteReader reads resp.Body *http.Response
// into a byte container anc closes the resp.Body.Close()[
func ResponseByteReader(resp *http.Response) (byteContainer []byte, err error) {
	var (
		buffer    []byte
		n         int
		container []byte
	)
	// declare buffer to be 1kb
	buffer = make([]byte, 1024)
	for {
		n, err = resp.Body.Read(buffer)
		// test if err
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return nil, err
		}

		// apend buffer to container
		container = append(buffer)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return container, nil
}
