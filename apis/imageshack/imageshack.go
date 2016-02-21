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

package imageshack

// "success":true,
//    "process_time":32,
//    "result":{
//       "id":"idggZf2Aj",
//       "server":661,
//       "bucket":8761,
//       "filename":"ggZf2A.jpg",
//       "original_filename":"e1ab262c288441e252ea507676e6c8a8.jpg",
//       "direct_link":"imagizer.imageshack.com\/img661\/8761\/ggZf2A.jpg",
// ImageShackJSON struct that hold JSON response serialized
type ImageShackJSON struct {
	Result struct {
		Direct_link string
	}
}

// ImageShack nacked struct
type ImageShack struct {
}

// NewImageShack alloc and returns a new *ImageShack{} struct
func NewImageShack() (i *ImageShack) {
	return &ImageShack{}
}
