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
	"encoding/json"
	"fmt"
	"os"

	"github.com/hoenirvili/ImageCompress/internal"
)

// {
//    "data": {
//        "error": "This method requires authentication",
//        "request": "\/3\/account.json",
//        "method": "GET",
//    },
//    "success": false,
//    "status": 403
// }
// ErrorImgurJSON error api construct error
// we can use this struct in order to parse api imgur response error
// note that we don't export this func because it's just belongs inside
// imgur package because it's internal to api imgur error response json
type errorImgurJSON struct {
	Data struct {
		Error   string
		Request string
		Method  string
	}
	Success bool
	Status  uint16
}

// Implement error interface
func (err errorImgurJSON) Error() string {
	return fmt.Sprintf(" Error Code : %d\n Error method : %s\n Error message : %s\n Request: %s\n", err.Status, err.Data.Method, err.Data.Error, err.Data.Request)
}

// Print just print the Code errors after JSON parse
func (err errorImgurJSON) Print() {
	fmt.Fprintf(os.Stderr, " Error Code : %d\n Error method : %s\n Error message : %s\n Request: %s\n", err.Status, err.Data.Method, err.Data.Error, err.Data.Request)
}

// ErrorResponseJSON check header error and handle it.
func errorResponseJSON(statusCode int, response []byte) *errorImgurJSON {
	switch statusCode {
	case 400, 401, 403, 404, 429, 500:
		jsonErr := &errorImgurJSON{}
		err := json.Unmarshal(response, jsonErr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't Unmarshall error json")
			os.Exit(1)
		}
		return jsonErr
	}
	return nil
}

// erroHttpStatus just checks if it's an error and parses
// corresponding massage suggested in imgur api handle error.
func errorHTTPStatus(statusCode int) *internal.ErrorStat {
	switch statusCode {
	case 400:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "Parameter is missing or a parameter has a value that is out of bounds or otherwise incorrect.image uploads fail due to images that are corrupt or do not meet the format requirements.")}
	case 401:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "The request requires user authentication.")}
	case 403:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "Forbidden. You don't have access to this action.")}
	case 404:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "Resource does not exist.")}
	case 429:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "Rate limiting on the application or on the user's IP address.")}
	case 500:
		return &internal.ErrorStat{Message: fmt.Sprintf("Error : %d %s\n", statusCode, "Unexpected internal error. Something is broken with the Imgur service.")}
	}
	// if it's not these types of erros just return nil.
	return nil
}
