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
	"fmt"
	"io/ioutil"
	"net/http"
)

// DumpResponse for debugging informatin
func DumpResponse(resp *http.Response) {
	for key, val := range resp.Header {
		fmt.Print(key)
		fmt.Print(" : ")
		fmt.Println(val)
	}
	fmt.Printf("Status : %s\n", resp.Status)
	readed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("===================== BODY ====================")
	fmt.Println()
	fmt.Println(string(readed))

	// for any case
	defer resp.Body.Close()
}
