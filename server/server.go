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

package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	port = ":8080"
)

// Init the server components middleware
func Init() {
	// load html into cache emplateList
	// predefine package variable that renderHTML operates on
	templateList = initTemplate()
	// make a new router
	router := httprouter.New()
	// custon 404 not found page
	router.NotFound = &undefineHandler{}

	// define routs
	router.GET("/", index)
	// static resources
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	// create http server and listen on port.
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
