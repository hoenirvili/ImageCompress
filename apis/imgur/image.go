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

//Image type struct JSON
type Image struct {
	Data struct {
		ID          string // The ID for the image
		Title       string // The title of the image.
		Description string // Description of the image.
		DateTime    int    // Time inserted into the gallery, epoch time
		Type        string // Image MIME type.
		Width       int    // The width of the image in pixels
		Height      int    // The height of the image in pixels
		Size        int    // The size of the image in bytes
		Views       int    // The number of image views
		Link        string // Link to to the image
	}
	Success bool   // success request
	Status  uint16 //
}
