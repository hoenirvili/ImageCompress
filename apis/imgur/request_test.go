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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"
)

const (
	clientID     = "40aea5f08c0f717"
	clientSecret = "a72c35d27b38d27114b4503e5b9acc835861ed8c"
	urlTest      = "https://api.imgur.com/3/gallery/image/i0xn0Dx"
)

func dumpJSON(v *Image) {
	fmt.Println("Data : ")
	fmt.Printf("\t ID : ")
	fmt.Println(v.Data.ID)

	fmt.Printf("\t Title : ")
	fmt.Println(v.Data.Title)

	fmt.Printf("\t Description : ")
	fmt.Println(v.Data.Description)

	fmt.Printf("\t DataTime : ")
	fmt.Println(v.Data.DateTime)

	fmt.Printf("\t Type : ")
	fmt.Println(v.Data.Type)

	fmt.Printf("\t Width : ")
	fmt.Println(v.Data.Width)

	fmt.Printf("\t Height : ")
	fmt.Println(v.Data.Height)

	fmt.Printf("\t Size : ")
	fmt.Println(v.Data.Size)

	fmt.Printf("\t Views : ")
	fmt.Println(v.Data.Views)

	fmt.Printf("\t Link : ")
	fmt.Println(v.Data.Link)

	fmt.Printf(" Success: ")
	fmt.Println(v.Success)

	fmt.Printf(" Link : ")
	fmt.Println(v.Status)

	fmt.Println()
}
func dumpResponse(resp *http.Response) {
	fmt.Println()
	fmt.Println()
	fmt.Printf("[ %s ] Protocol : %s \n", resp.Status, resp.Proto)
	for val, key := range resp.Header {
		fmt.Printf("%s : %s\n", key, val)
	}
	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("============= Body BYTE ======================")
	fmt.Println()
	fmt.Println(reader)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println(string(reader))
}
func TestJSON(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println("[ > ]\t Respnse JSON")
	imgur := NewImgur()
	imgur.SetClientID(clientID)
	imgur.SetClientSecret(clientSecret)
	imgur.SetBody(new(bytes.Buffer))
	v, err := imgur.ImageJSON(urlTest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dumpJSON(v)
}

func TestGetRequest(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println(" [ > ]\t Get response")
	imgur := NewImgur()
	resp, err := imgur.Get(urlTest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dumpResponse(resp)
}

func TestDumpRequest(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println("[ > ]\t Http Request")
	fmt.Println()

	req, err := http.NewRequest("GET", urlTest, new(bytes.Buffer))
	if err != nil {
		log.Fatal(err)
	}
	// add the corresponding header
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))
	rsp, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(rsp))
}
