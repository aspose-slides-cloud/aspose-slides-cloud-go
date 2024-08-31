/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package usecasetests

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"testing"
)

/*
   Test for Get image
*/
func TestImagesGet(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	presentationResult, _, e := c.SlidesApi.GetPresentationImages(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	slideResult, _, e := c.SlidesApi.GetSlideImages(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(slideResult.GetList()) >= len(presentationResult.GetList()) {
		t.Errorf("Wrong image count. Expected less than %v but was %v.", len(presentationResult.GetList()), len(slideResult.GetList()))
		return
	}
}

/*
   Test for download all images from storage
*/
func TestImagesDownloadStorage(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	downloadResult, _, e := c.SlidesApi.DownloadImagesDefaultFormat(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadStat, e := os.Stat(downloadResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadZip, e := zip.OpenReader(downloadResult.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer downloadZip.Close()
	downloadCount := 0
	for _, _ = range downloadZip.File {
		downloadCount++
	}

	pngResult, _, e := c.SlidesApi.DownloadImages(fileName, "png", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngStat, e := os.Stat(pngResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngZip, e := zip.OpenReader(pngResult.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer pngZip.Close()
	pngCount := 0
	for _, _ = range pngZip.File {
		pngCount++
	}
	if downloadStat.Size() == pngStat.Size() {
		t.Errorf("Wrong file size. Expected not %v but was %v.", downloadStat.Size(), pngStat.Size())
		return
	}
	if downloadCount != pngCount {
		t.Errorf("Wrong image count. Expected %v but was %v.", downloadCount, pngCount)
		return
	}
}

/*
   Test for download all images from request
*/
func TestImagesDownloadRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	downloadResult, _, e := c.SlidesApi.DownloadImagesDefaultFormatOnline(source, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadStat, e := os.Stat(downloadResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadZip, e := zip.OpenReader(downloadResult.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer downloadZip.Close()
	downloadCount := 0
	for _, _ = range downloadZip.File {
		downloadCount++
	}

	pngResult, _, e := c.SlidesApi.DownloadImagesOnline(source, "png", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngStat, e := os.Stat(pngResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngZip, e := zip.OpenReader(pngResult.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer pngZip.Close()
	pngCount := 0
	for _, _ = range pngZip.File {
		pngCount++
	}
	if downloadStat.Size() == pngStat.Size() {
		t.Errorf("Wrong file size. Expected not %v but was %v.", downloadStat.Size(), pngStat.Size())
		return
	}
	if downloadCount != pngCount {
		t.Errorf("Wrong image count. Expected %v but was %v.", downloadCount, pngCount)
		return
	}
}

/*
   Test for download image from storage
*/
func TestImageDownloadStorage(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	downloadResult, _, e := c.SlidesApi.DownloadImageDefaultFormat(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadStat, e := os.Stat(downloadResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	pngResult, _, e := c.SlidesApi.DownloadImage(fileName, slideIndex, "png", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngStat, e := os.Stat(pngResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if downloadStat.Size() == pngStat.Size() {
		t.Errorf("Wrong file size. Expected not %v but was %v.", downloadStat.Size(), pngStat.Size())
		return
	}
}

/*
   Test for download image from request
*/
func TestImageDownloadRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	downloadResult, _, e := c.SlidesApi.DownloadImageDefaultFormatOnline(source, slideIndex, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	downloadStat, e := os.Stat(downloadResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	pngResult, _, e := c.SlidesApi.DownloadImageOnline(source, slideIndex, "png", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	pngStat, e := os.Stat(pngResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if downloadStat.Size() == pngStat.Size() {
		t.Errorf("Wrong file size. Expected not %v but was %v.", downloadStat.Size(), pngStat.Size())
		return
	}
}

/*
   Test for replace image
*/
func TestReplaceImage(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	image, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.ReplaceImage(fileName, 1, image, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for replace image from request
*/
func TestReplaceImageRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	image, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.ReplaceImageOnline(source, 1, image, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for delete picture cropped areas
*/
func TestDeletePictureCroppedAreas(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	 _, e = c.SlidesApi.DeletePictureCroppedAreas(fileName, 2, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for delete picture cropped areas for wrong shape type
*/
func TestDeletePictureCroppedAreasWrongShapeType(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, e := c.SlidesApi.DeletePictureCroppedAreas(fileName, 2, 3, password, folderName, "")
	if e == nil {
		t.Errorf("Should throw an exception if shape is not PictureFrame")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}
