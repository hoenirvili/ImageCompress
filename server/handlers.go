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
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// index internal handler just server index page
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("%+v", r)
	index, err := template.ParseFiles("views/header.tmpl", "views/body.tmpl", "views/footer.tmpl", "views/base.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = index.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}
