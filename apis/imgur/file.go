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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/hoenirvili/ImageCompress/internal"
	"github.com/hoenirvili/ImageCompress/utils"
)

// DownloadImage make's http get request and save's the file
// with mime extension
func (i Imgur) DownloadImage(url, mime, name string) error {
	var (
		out *os.File
	)

	switch mime {
	case "image/png":
		var (
			err error // for scope reasons
			res string
		)
		res, err = utils.Concat(name, ".png")
		if err != nil {
			return err
		}
		out, err = os.Create(res)
		if err != nil {
			return err
		}
	case "image/jpeg":
		var (
			err error // for scope reasons
			res string
		)
		res, err = utils.Concat(name, ".jpg")
		if err != nil {
			return err
		}
		out, err = os.Create(res)
		if err != nil {
			return nil
		}
	}

	resp, err := i.Get(url)
	if err != nil {
		return err
	}

	_, errCopy := io.Copy(out, resp.Body)
	if errCopy != nil {
		return err
	}

	defer func() {
		if err := out.Close(); err != nil {
			log.Fatal(err)
		}

	}()

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

//SaveImageByte saves images form byte serialez to hdd
func (i Imgur) SaveImageByte(img []byte, path string) error {
	//TODO better opt way
	err := ioutil.WriteFile(path, img, 0644)
	if err != nil {
		return internal.ErrorStat{Message: fmt.Sprintf("%s", "Can't save image from Imgurl")}
	}

	return nil
}
