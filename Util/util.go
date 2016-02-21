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

package Util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Concat concat two string into one
func Concat(first, second string) (result string, err error) {
	//init
	nCopied := 0

	// get lens
	lenFirst := len(first)
	lenSecond := len(second)

	// sum of both lens
	n := lenFirst + lenSecond

	// alloc holder
	holder := make([]byte, n)

	// copy the first holder from pos nCopied:=0
	nCopied = copy(holder[nCopied:], first)

	// if everything is ok copy the second one.
	if nCopied != n-lenSecond {
		return "", errors.New("Can't copy first string")
	}

	// coy the second holder from pos nCopied:= previousNcopied
	nCopied = copy(holder[nCopied:], second)

	if nCopied != n-lenFirst {
		return "", errors.New("Can't copy the second string")
	}

	// TODO find a better way to convert it to string
	return string(holder), nil
}

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