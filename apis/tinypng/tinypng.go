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

// curl --user api:$YOUR_API_KEY \
//      --data-binary @unoptimized.png -i https://api.tinify.com/shrink

package tinypng

// contains unexpected
// const vars inside tiny package
const (
	urlSender = "https://api.tinify.com/shrink"
	username  = "api:"
	key       = "47yoyAZ-RSj34OwwRCJ5EImb8nUDLHx8"
)

// TinyJSON struct for parsing JSON tiny api response
type TinyJSON struct {
	Input struct {
		Size int
		Type string
	}
	Output struct {
		Size   int
		Type   string
		Width  int
		Height int
		Ratio  float64
		URL    string
	}
}

// Tiny struct
// contains just the body that will
// be used to post request to the api compression system
type Tiny struct {
	body []byte // fill the body with the post request content
}

// NewTiny return a new pointer to *Tiny
func NewTiny() (t *Tiny) {
	return &Tiny{}
}

//SetBody for post request
func (t *Tiny) SetBody(b []byte) {
	t.body = b
}

//Body get value for body
func (t Tiny) Body() []byte {
	return t.body
}
