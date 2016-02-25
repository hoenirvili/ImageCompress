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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"testing"
)

const (
	testURL = "https://api.imageshack.com/v2/images/pbzPCsEij"
)

func dumpJSON(v *Image) {
	fmt.Println("Result")
	fmt.Printf("\t %s\n", v.Result.Direct_link)
}

func dumpResponse(resp *http.Response) {
	fmt.Println()
	fmt.Println()
	fmt.Printf(" [ %s ] : %s \n", resp.Status, resp.Proto)
	for val, key := range resp.Header {
		fmt.Printf("%s : %s\n", key, val)
	}
	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("==================== Body Byte =================")
	fmt.Println(reader)
	fmt.Println()
	fmt.Println("================================================")
	fmt.Println()
	fmt.Println(string(reader))
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
func TestJSON(t *testing.T) {
	shack := NewImageShack()
	v, err := shack.ImageJSON(testURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println()
	dumpJSON(v)
	fmt.Println()
	fmt.Println()
}

func TestGet(t *testing.T) {
	shack := NewImageShack()
	fmt.Println(" [ > ]\t Get response")
	resp, err := shack.Get(testURL)
	if err != nil {
		log.Fatal(err)
	}
	dumpResponse(resp)
}

func TestDumpRequest(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println(" [ > ]\t Http Request")
	fmt.Println()

	req, err := http.NewRequest(http.MethodGet, testURL, new(bytes.Buffer))
	if err != nil {
		log.Fatal(err)
	}
	rsp, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(rsp))
}
