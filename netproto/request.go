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

package netproto

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// HeaderToBase64 Base64 encoding defined by rfc RFC2045-MIME
// plus concat the "Basic: "word.
func HeaderToBase64(username, key string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s", username, key)))
	return fmt.Sprintf("Basic %s", sEnc)
}

// Basic request
var basicRequest = &http.Request{
	Method:           "",
	URL:              nil,
	Proto:            "HTTP 1.0",
	ProtoMajor:       1,
	ProtoMinor:       1,
	Header:           make(http.Header),
	Body:             nil,
	ContentLength:    0,
	TransferEncoding: nil,
	Close:            false,
	Host:             "",
	Form:             nil,
	PostForm:         nil,
	MultipartForm:    nil,
	Trailer:          nil,
	RemoteAddr:       "",
	TLS:              nil,
	Cancel:           nil,
}

// // NewRequest func return a *http.Request
// // that can be used with Client.Do() func
// func (i Imgur) NewRequest(methodRequest, urlRequest, mimeRequest string) (*http.Request, *errors.ErrorStat) {
// 	u, err := url.Parse(urlRequest)
//
// 	if err != nil {
// 		return nil, &errors.ErrorStat{Message: fmt.Sprintf("Can't parse this url %s", urlRequest)}
// 	}
//
// 	rc, ok := i.body.(io.ReadCloser)
//
// 	if !ok && i.body != nil {
// 		rc = ioutil.NopCloser(i.body)
// 	}
//
// 	// Attach all basic config for request
// 	basicRequest.Method = methodRequest
// 	basicRequest.URL = u
//
// 	basicRequest.Header.Add("Content-Type", mimeRequest)
// 	basicRequest.Body = rc
//
// 	// Note that in go you don't need break stmt
// 	if i.body != nil {
// 		switch v := i.body.(type) {
// 		case *bytes.Buffer:
// 			basicRequest.ContentLength = int64(v.Len())
// 		case *bytes.Reader:
// 			basicRequest.ContentLength = int64(v.Len())
// 		case *strings.Reader:
// 			basicRequest.ContentLength = int64(v.Len())
// 		}
// 	}
//
// 	// return it
// 	return basicRequest, nil
// }
