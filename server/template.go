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
	"html/template"
	"net/http"
)

const (
	templateDirPath = "views/"
)

// map that will holds all template directory
var templateList map[string]*template.Template

// init rednerTemplate with configurations
func initTemplate() map[string]*template.Template {
	list := map[string]*template.Template{
		"index": template.Must(template.ParseFiles("views/header.tmpl", "views/footer.tmpl", "views/index/body.tmpl", "views/index/index.tmpl")),
		"404":   template.Must(template.ParseFiles("views/header.tmpl", "views/404.tmpl")),
	}
	// return configurations
	return list
}

// basic method to render html pages
func renderHTML(w http.ResponseWriter, page string, status int, data interface{}) {
	// if the page exists
	if t, ok := templateList[page]; ok {
		// set headers
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		// write response
		err := t.ExecuteTemplate(w, page, data)
		// if error response
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
