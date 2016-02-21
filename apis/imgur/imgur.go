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

import "io"

// TestBaseURL just testing purpose
const TestBaseURL = "https://api.imgur.com/3/gallery/nuSZyEE"

// Addon for production purpose attachemnt
// for HTTP header request.
// Special flags for imgur api
const (
	userLimit       = "X-RateLimit-UserLimit"
	userRemaining   = "X-RateLimit-UserRemaining"
	userReset       = "X-RateLimit-UserReset"
	clientLimit     = "X-RateLimit-ClientLimit"
	clientRemaining = "X-RateLimit-ClientRemaining"
)

// Basic Imgur api configuration
const (
	baseURL  = "https://api.imgur.com/3"
	stateURL = "https://api.imgur.com/3/credits"
)

// Imgur type struct for config
// api app info
type Imgur struct {
	clientID     string
	clientSecret string
	body         io.Reader
}

// SetClientID set's the id of the app to
// identify who is requesting for info
func (i *Imgur) SetClientID(id string) {
	i.clientID = id
}

// SetClientSecret set's the secretID of the app to
// identify who is requesting for info
func (i *Imgur) SetClientSecret(secret string) {
	i.clientSecret = secret
}

// SetBody set's the body type of the request
// Note that body is io.Reader and the type should
// implement io.Reader interface
func (i *Imgur) SetBody(body io.Reader) {
	i.body = body
}

// ClientID get's the client id
func (i Imgur) ClientID() string {
	return i.clientID
}

//ClientSecret get's client secretID
func (i Imgur) ClientSecret() string {
	return i.clientSecret
}

// Body get's the body
func (i Imgur) Body() (body io.Reader) {
	return i.body
}

//NewImgur alloc a new pointer to Imgur{} struct
func NewImgur() (i *Imgur) {
	return &Imgur{}
}
