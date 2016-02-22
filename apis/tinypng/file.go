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

package tinypng

import (
	"io/ioutil"

	"github.com/hoenirvili/ImageCompress/utils"
)

// SaveImage saves the new image compressed
func (t Tiny) SaveImage(img []byte, name, mime string) error {

	if mime == "image/jpeg" {
		//Write compressed file
		nameJPG, err := utils.Concat(name, ".jpg")
		if err != nil {
			nameJPG = "defaultCompressed.jpg"
		}
		errWrite := ioutil.WriteFile(nameJPG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	} else if mime == "image/png" {
		namePNG, err := utils.Concat(name, ".png")
		if err != nil {
			namePNG = "defaultCompressed.png"
		}
		errWrite := ioutil.WriteFile(namePNG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	}
	return nil
}
