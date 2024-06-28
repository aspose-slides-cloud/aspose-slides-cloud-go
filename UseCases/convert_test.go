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
	"io/ioutil"
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v24"
)

/*
   Test for convert presentation from request to request
*/
func TestConvertPostFromRequest(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.Convert(source, "pdf", password, "", "", nil, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	slides := []int32{2, 4}
	resultSlides, _, e := c.SlidesApi.Convert(source, "pdf", password, "", "", slides, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultSlidesStat, e := os.Stat(resultSlides.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultStat.Size() <= resultSlidesStat.Size() {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", resultSlidesStat.Size(), resultStat.Size())
		return
	}
}

/*
   Test for convert presentation from request to storage
*/
func TestConvertPutFromRequest(t *testing.T) {
	outPath := "TestData/test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.ConvertAndSave(source, "pdf", outPath, "password", "", "", nil, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert presentation from storage to request
*/
func TestConvertPostFromStorage(t *testing.T) {
	fileName := "test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "html5", nil, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert presentation from storage to storage
*/
func TestConvertPutFromStorage(t *testing.T) {
	outPath := "TestData/test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SavePresentation(fileName, "pdf", outPath, nil, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert presentation with options from request
*/
func TestConvertWithOptionsFromRequest(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.Convert(source, "pdf", password, "", "", nil, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := slidescloud.NewPdfExportOptions()
	drawSlidesFrame := true
	options.DrawSlidesFrame = &drawSlidesFrame
	resultOptions, _, e := c.SlidesApi.Convert(source, "pdf", password, "", "", nil, options)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultOptionsStat, e := os.Stat(resultOptions.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultOptionsStat.Size() == resultStat.Size() {
		t.Errorf("Wrong response size. Expected not %v but was %v.", resultOptionsStat.Size(), resultStat.Size())
		return
	}
}

/*
   Test for convert presentation with options from storage
*/
func TestConvertWithOptionsFromStorage(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DownloadPresentation(fileName, "png", nil, password, folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := slidescloud.NewImageExportOptions()
	options.Width = 480
	options.Height = 360
	resultOptions, _, e := c.SlidesApi.DownloadPresentation(fileName, "png", options, password, folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultOptionsStat, e := os.Stat(resultOptions.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultOptionsStat.Size() >= resultStat.Size() {
		t.Errorf("Wrong response size. Expected less than %v but was %v.", resultStat.Size(), resultOptionsStat.Size())
		return
	}
}

/*
   Test for convert slide from request to request
*/
func TestConvertSlidePostFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = slidescloud.GetTestSlidesApiClient().SlidesApi.DownloadSlideOnline(source, 1, "pdf", nil, nil, "password", "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert slide from request to storage
*/
func TestConvertSlidePutFromRequest(t *testing.T) {
	outPath := "TestData/test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.SaveSlideOnline(source, 1, "pdf", outPath, nil, nil, "password", "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert slide from storage to request
*/
func TestConvertSlidePostFromStorage(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadSlide(fileName, 1, "pdf", nil, nil, nil, "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert slide from storage to storage
*/
func TestConvertSlidePutFromStorage(t *testing.T) {
	outPath := "TestData/test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SaveSlide(fileName, 1, "pdf", outPath, nil, nil, nil, "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert slide with options from request
*/
func TestConvertSlideWithOptionsFromRequest(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DownloadSlideOnline(source, 1, "pdf", nil, nil, password, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := slidescloud.NewPdfExportOptions()
	drawSlidesFrame := true
	options.DrawSlidesFrame = &drawSlidesFrame
	resultOptions, _, e := c.SlidesApi.DownloadSlideOnline(source, 1, "pdf", nil, nil, password, "", "", options)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultOptionsStat, e := os.Stat(resultOptions.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultOptionsStat.Size() == resultStat.Size() {
		t.Errorf("Wrong response size. Expected not %v but was %v.", resultOptionsStat.Size(), resultStat.Size())
		return
	}
}

/*
   Test for convert slide with options from storage
*/
func TestConvertSlideWithOptionsFromStorage(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DownloadSlide(fileName, 1, "pdf", nil, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := slidescloud.NewPdfExportOptions()
	drawSlidesFrame := true
	options.DrawSlidesFrame = &drawSlidesFrame
	resultOptions, _, e := c.SlidesApi.DownloadSlide(fileName, 1, "pdf", options, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultOptionsStat, e := os.Stat(resultOptions.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultOptionsStat.Size() == resultStat.Size() {
		t.Errorf("Wrong response size. Expected not %v but was %v.", resultOptionsStat.Size(), resultStat.Size())
		return
	}
}

/*
   Test for convert shape from request to request
*/
func TestConvertShapePostFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = slidescloud.GetTestSlidesApiClient().SlidesApi.DownloadShapeOnline(source, 1, 3, "png", nil, nil, "", "password", "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert shape from request to storage
*/
func TestConvertShapePutFromRequest(t *testing.T) {
	outPath := "TestData/test.png"
	c := slidescloud.GetTestSlidesApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.SaveShapeOnline(source, 1, 1, "png", outPath, nil, nil, "", "password", "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert shape from storage to request
*/
func TestConvertShapePostFromStorage(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadShape(fileName, 1, 1, "png", nil, nil, nil, "", "password", folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert subshape from storage to request
*/
func TestConvertSubShapePostFromStorage(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var subShape = "1"
	var shapeIndex int32 = 4
	_, _, e = c.SlidesApi.DownloadShape(fileName, 1, shapeIndex, "png", nil, nil, nil, "", "password", folderName, "", "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert shape from storage to storage
*/
func TestConvertShapePutFromStorage(t *testing.T) {
	outPath := "TestData/test.png"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SaveShape(fileName, 1, 1, "png", outPath, nil, nil, nil, "", "password", folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert subshape from storage to storage
*/
func TestConvertSubshapePutFromStorage(t *testing.T) {
	outPath := "TestData/test.png"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var subShape = "1"
	var shapeIndex int32 = 4

	_, e = c.SlidesApi.SaveShape(fileName, 1, shapeIndex, "png", outPath, nil, nil, nil, "", "password", folderName, "", "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for conversion with font fallback rules.
*/
func TestConvertWithFontFallBackRules(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fontRule0 := slidescloud.NewFontFallbackRule()
	fontRule0.SetRangeStartIndex(0x0B80)
	fontRule0.SetRangeEndIndex(0x0BFF)
	fontRule0.SetFallbackFontList([]string{"Vijaya"})

	fontRule1 := slidescloud.NewFontFallbackRule()
	fontRule1.SetRangeStartIndex(0x0B80)
	fontRule1.SetRangeEndIndex(0x0BFF)
	fontRule1.SetFallbackFontList([]string{"Segoe UI Emoji", "Segoe UI Symbol", "Arial"})

	fontRules := []slidescloud.IFontFallbackRule{fontRule0, fontRule1}

	exportOptions := slidescloud.NewImageExportOptions()
	exportOptions.SetFontFallbackRules(fontRules)

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "png", exportOptions, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for conversion with slide layout options.
*/
func TestConvertWithSlideLayoutOptions(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	slideLayoutOptions := slidescloud.NewHandoutLayoutingOptions()
	slideLayoutOptions.SetHandout("Handouts2")
	printSlideNumbers := true
	slideLayoutOptions.SetPrintSlideNumbers(&printSlideNumbers)

	exportOptions := slidescloud.NewPdfExportOptions()
	exportOptions.SetSlidesLayoutOptions(slideLayoutOptions)

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "pdf", exportOptions, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
