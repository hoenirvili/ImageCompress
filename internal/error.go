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

package internal

import (
	"fmt"
	"os"
)

// ErrorStat general struct error type.
// this struct can be used the code that is not
// a specific http/api/protocol error.
// A more generic error type
type ErrorStat struct {
	Message string // the actual error message
}

// struct ErrorStat implements Stringer inteface
// Now we can use this with the fmt package.
func (e ErrorStat) String() string {
	return fmt.Sprintf("[ > ]\tError : %s\n", e.Message)
}

// Print func it's just a util function that we can use
// independently as it is.
func (e ErrorStat) Print() {
	fmt.Fprintf(os.Stderr, "[ > ]\t %s\n", e.Message)
}

// struct ErrorStat implements Error interface
// Now we can use this with return error types.
func (e ErrorStat) Error() string {
	return fmt.Sprintf("[ > ]\tError : %s\n", e.Message)
}
