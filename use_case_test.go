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

package asposeslidescloud

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"testing"
)

/*
   Test for create empty presentation
*/
func TestCreateEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.CreatePresentation(fileName, nil, "", "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create presentation from request
*/
func TestCreateFromRequest(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data, e := ioutil.ReadFile("TestData/" + fileName)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.CreatePresentation(fileName, data, "password", "", folderName, "")
}

/*
   Test for create presentation from storage
*/
func TestCreateFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	newFileName := "test2.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+newFileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	sourcePath := folderName + "/" + fileName
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.CreatePresentationFromSource(newFileName, sourcePath, "password", "", "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create presentation from template
*/
func TestCreateFromTemplate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	templateFileName := "TemplateCV.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	templatePath := folderName + "/" + templateFileName
	_, e = c.SlidesApi.CopyFile("TempTests/"+templateFileName, templatePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data := "<staff><person><name>John Doe</name><address><line1>10 Downing Street</line1><line2>London</line2></address><phone>+457 123456</phone><bio>Hi, I'm John and this is my CV</bio><skills><skill><title>C#</title><level>Excellent</level></skill><skill><title>C++</title><level>Good</level></skill><skill><title>Java</title><level>Average</level></skill></skills></person></staff>"
	_, _, e = c.SlidesApi.CreatePresentationFromTemplate(fileName, templatePath, data, "", "", nil, "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create presentation from HTML
*/
func TestCreateFromHtml(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ImportFromHtml(fileName, "<html><body>New Content</body></html>", "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for append presentation from HTML
*/
func TestAppendFromHtml(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	GetResponse1, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideCount := len(GetResponse1.GetSlideList())
	_, _, e = c.SlidesApi.ImportFromHtml(fileName, "<html><body>New Content</body></html>", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	GetResponse2, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	newSlideCount := len(GetResponse2.GetSlideList())
	if newSlideCount != slideCount+1 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", slideCount+1, newSlideCount)
		return
	}
}

/*
   Test for create presentation from PDF
*/
func TestCreateFromPdf(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source, e := ioutil.ReadFile("TestData/test.pdf")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ImportFromPdf(fileName, source, "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for append presentation from PDF
*/
func TestAppendFromPdf(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	GetResponse1, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideCount := len(GetResponse1.GetSlideList())
	source, e := ioutil.ReadFile("TestData/test.pdf")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ImportFromPdf(fileName, source, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	GetResponse2, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	newSlideCount := len(GetResponse2.GetSlideList())
	if newSlideCount != slideCount+4 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", slideCount+4, newSlideCount)
		return
	}
}

/*
   Test for convert presentation from request to request
*/
func TestConvertPostFromRequest(t *testing.T) {
	password := "password"
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert presentation from storage to request
*/
func TestConvertPostFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pdf"
	c := GetTestApiClient()
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
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	outPath := "TestData/test.pdf"
	c := GetTestApiClient()
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
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert presentation with options from request
*/
func TestConvertWithOptionsFromRequest(t *testing.T) {
	password := "password"
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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

	options := NewPdfExportOptions()
	options.DrawSlidesFrame = true
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
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
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

	options := NewImageExportOptions()
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
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = GetTestApiClient().SlidesApi.DownloadSlideOnline(source, 1, "pdf", nil, nil, "password", "", "", nil)
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
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert slide from storage to request
*/
func TestConvertSlidePostFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
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
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	outPath := "TestData/test.pdf"
	c := GetTestApiClient()
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
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert slide with options from request
*/
func TestConvertSlideWithOptionsFromRequest(t *testing.T) {
	password := "password"
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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

	options := NewPdfExportOptions()
	options.DrawSlidesFrame = true
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
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
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

	options := NewPdfExportOptions()
	options.DrawSlidesFrame = true
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
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = GetTestApiClient().SlidesApi.DownloadShapeOnline(source, 1, 3, "png", nil, nil, "", "password", "", "", nil)
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
	c := GetTestApiClient()
	source, e := ioutil.ReadFile("TestData/test.pptx")
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
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert shape from storage to request
*/
func TestConvertShapePostFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadShape(fileName, 1, 1, "png", nil, nil, nil, "", "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert subshape from storage to request
*/
func TestConvertSubShapePostFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadSubshape(fileName, 1, "4/shapes", 1, "png", nil, nil, nil, "", "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for convert shape from storage to storage
*/
func TestConvertShapePutFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	outPath := "TestData/test.png"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SaveShape(fileName, 1, 1, "png", outPath, nil, nil, nil, "", "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for convert subshape from storage to storage
*/
func TestConvertSubshapePutFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	outPath := "TestData/test.png"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SaveSubshape(fileName, 1, "4/shapes", 1, "png", outPath, nil, nil, nil, "", "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for conversion with font fallback rules.
*/

func TestConvertWithFontFallBackRules(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fontRule0 := NewFontFallbackRule()
	fontRule0.SetRangeStartIndex(0x0B80)
	fontRule0.SetRangeEndIndex(0x0BFF)
	fontRule0.SetFallbackFontList([]string{"Vijaya"})

	fontRule1 := NewFontFallbackRule()
	fontRule1.SetRangeStartIndex(0x0B80)
	fontRule1.SetRangeEndIndex(0x0BFF)
	fontRule1.SetFallbackFontList([]string{"Segoe UI Emoji", "Segoe UI Symbol", "Arial"})

	fontRules := []IFontFallbackRule{fontRule0, fontRule1}

	exportOptions := NewImageExportOptions()
	exportOptions.SetFontFallbackRules(fontRules)

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "png", exportOptions, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for Get notes slide from storage
*/
func TestNotesSlideGetFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetNotesSlide(fileName, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetText()) == 0 {
		t.Errorf("Note text is empty.")
		return
	}
}

/*
   Test for exists notes slide from storage
*/
func TestNotesSlideExistsFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.NotesSlideExists(fileName, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetExists() {
		t.Errorf("Note does not exist.")
		return
	}
}

/*
   Test for download notes slide from storage
*/
func TestNotesSlideDownloadFromStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.DownloadNotesSlide(fileName, 1, "png", nil, nil, "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for Get notes slide from request
*/
func TestNotesSlideGetFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := GetTestApiClient().SlidesApi.GetNotesSlideOnline(source, 1, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetText()) == 0 {
		t.Errorf("Note text is empty.")
		return
	}
}

/*
   Test for exists notes slide from request
*/
func TestNotesSlideExistsFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := GetTestApiClient().SlidesApi.NotesSlideExistsOnline(source, 1, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetExists() {
		t.Errorf("Note does not exist.")
		return
	}
}

/*
   Test for download notes slide from request
*/
func TestNotesSlideDownloadFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = GetTestApiClient().SlidesApi.DownloadNotesSlideOnline(source, 1, "png", nil, nil, "password", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for notes slide shapes
*/
func TestNotesSlideShapes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeCount int32 = 3
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.GetShapesLinks()))
		return
	}

	dto := NewShape()
	dto.X = 100
	dto.Y = 100
	dto.Width = 500
	dto.Height = 200
	dto.ShapeType = "Rectangle"
	dto.Text = "New shape"
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "notesSlide", dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}
}

/*
   Test for notes slide paragraphs
*/
func TestNotesSlideParagraphs(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphCount int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}

	portion := NewPortion()
	portion.Text = "New Paragraph"
	dto := NewParagraph()
	dto.Alignment = "Right"
	dto.PortionList = []IPortion{portion}
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	dto = NewParagraph()
	dto.Alignment = "Center"
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}
}

/*
   Test for notes slide portions
*/
func TestNotesSlidePortions(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionCount int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}

	dto := NewPortion()
	dto.Text = "New portion"
	dto.FontBold = "True"
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetText() != dto.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	dto2 := NewPortion()
	dto2.Text = "Updated portion"
	dto2.FontHeight = 22
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount+1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetFontHeight() != dto2.GetFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.GetFontHeight(), portion.GetFontHeight())
		return
	}
	if portion.GetText() != dto2.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}
}

/*
   Test create notes slide
*/
func TestCreateNotesSlideUseCase(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	notesSlideText := "Notes slides string"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewNotesSlide()
	dto.SetText(notesSlideText)
	notesSlide, _, e := c.SlidesApi.CreateNotesSlide(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if notesSlide.GetText() != notesSlideText {
		t.Errorf("Notes slides text error. Expected \"%v\" but was \"%v\"", notesSlide.GetText(), notesSlideText)
		return
	}
}

/*
   Test update notes slide
*/
func TestUpdateNotesSlideUseCase(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	notesSlideText := "Notes slides string"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewNotesSlide()
	dto.SetText(notesSlideText)
	notesSlide, _, e := c.SlidesApi.UpdateNotesSlide(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if notesSlide.GetText() != notesSlideText {
		t.Errorf("Notes slides text error. Expected \"%v\" but was \"%v\"", notesSlide.GetText(), notesSlideText)
		return
	}
}

/*
   Test for deleting notes slide
*/
func TestDeleteNotesSlideUseCase(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	slide, _, e := c.SlidesApi.DeleteNotesSlide(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if slide.GetNotesSlide() != nil {
		t.Errorf("NoteSlide expected to be nil.")
		return
	}
}

/*
   Test for master slides
*/
func TestMasterSlides(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	sourceFile := "TemplateCV.pptx"
	sourcePath := folderName + "/" + sourceFile
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+sourceFile, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	masterSlides, _, e := c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 1 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 1, len(masterSlides.GetSlideList()))
		return
	}
	masterSlide, _, e := c.SlidesApi.GetMasterSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.GetName() != "Office Theme" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Office Theme", masterSlide.GetName())
		return
	}

	masterSlide, _, e = c.SlidesApi.CopyMasterSlide(fileName, sourcePath, 1, "", "", nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.GetName() != "Digital portfolio" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Digital portfolio", masterSlide.GetName())
		return
	}
	masterSlides, _, e = c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 2 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 2, len(masterSlides.GetSlideList()))
		return
	}
}

/*
   Test for master slide shapes
*/
func TestMasterSlideShapes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeCount int32 = 6
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.GetShapesLinks()))
		return
	}

	dto := NewShape()
	dto.X = 100
	dto.Y = 100
	dto.Width = 500
	dto.Height = 200
	dto.ShapeType = "Rectangle"
	dto.Text = "New shape"
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "masterSlide", dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}
}

/*
   Test for master slide paragraphs
*/
func TestMasterSlideParagraphs(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphCount int32 = 5
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}

	portion := NewPortion()
	portion.Text = "New Paragraph"
	dto := NewParagraph()
	dto.Alignment = "Right"
	dto.PortionList = []IPortion{portion}
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	dto = NewParagraph()
	dto.Alignment = "Center"
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}
}

/*
   Test for master slide portions
*/
func TestMasterSlidePortions(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 3
	var portionCount int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}

	dto := NewPortion()
	dto.Text = "New portion"
	dto.FontBold = "True"
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetText() != dto.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	dto2 := NewPortion()
	dto2.Text = "Updated portion"
	dto2.FontHeight = 22
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount+1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetFontHeight() != dto2.GetFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.GetFontHeight(), portion.GetFontHeight())
		return
	}
	if portion.GetText() != dto2.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}
}

/*
   Test for master slide animation
*/
func TestMasterSlideAnimation(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	effect1 := NewEffect()
	effect1.Type_ = "Blink"
	effect1.ShapeIndex = 2
	effect2 := NewEffect()
	effect2.Type_ = "Appear"
	effect2.ShapeIndex = 3
	dto := NewSlideAnimation()
	dto.MainSequence = []IEffect{effect1, effect2}
	animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	var shapeIndex int32 = 3
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "masterSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence())-1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence())-1, len(animation.GetMainSequence()))
		return
	}
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}

	animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	dto2 := NewInteractiveSequence()
	dto2.TriggerShapeIndex = 1

	effect := NewEffect()
	effect.Type_ = "Fly"
	effect.Subtype = "Bottom"
	effect.PresetClassType = "Entrance"
	effect.ShapeIndex = 3
	effect.TriggerType = "OnClick"
	dto2.Effects = []IEffect{effect}

	animation, _, e = c.SlidesApi.CreateAnimationInteractiveSequence(fileName, 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
	}
}

/*
   Test for deleting master slide
*/
func TestMasterSlideDeleteUnused(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var ignorePreserveField bool = true

	masterSlides, _, e := c.SlidesApi.DeleteUnusedMasterSlides(fileName, &ignorePreserveField, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 1 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 1, len(masterSlides.GetSlideList()))
		return
	}
}

/*
   Test for deleting master slide from request
*/
func TestMasterSlideDeleteUnusedOnline(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var ignorePreserveField bool = true

	result, _, e := c.SlidesApi.DeleteUnusedMasterSlidesOnline(source, &ignorePreserveField, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if resultStat.Size() == 0 {
		t.Errorf("Wrong response size.")
		return
	}
}

/*
   Test for layout slides
*/
func TestLayoutSlides(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	sourceFile := "TemplateCV.pptx"
	sourcePath := folderName + "/" + sourceFile
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+sourceFile, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	layoutSlides, _, e := c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.GetSlideList()) != 11 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 11, len(layoutSlides.GetSlideList()))
		return
	}
	layoutSlide, _, e := c.SlidesApi.GetLayoutSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.GetName() != "Title Slide" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title Slide", layoutSlide.GetName())
		return
	}

	layoutSlide, _, e = c.SlidesApi.CopyLayoutSlide(fileName, sourcePath, 2, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.GetName() != "Title and Content" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title and Content", layoutSlide.GetName())
		return
	}
	layoutSlides, _, e = c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.GetSlideList()) != 12 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 12, len(layoutSlides.GetSlideList()))
		return
	}
}

/*
   Test for layout slide shapes
*/
func TestLayoutSlideShapes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeCount int32 = 6
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.GetShapesLinks()))
		return
	}

	dto := NewShape()
	dto.X = 100
	dto.Y = 100
	dto.Width = 500
	dto.Height = 200
	dto.ShapeType = "Rectangle"
	dto.Text = "New shape"
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "layoutSlide", dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}
}

/*
   Test for layout slide paragraphs
*/
func TestLayoutSlideParagraphs(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphCount int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}

	portion := NewPortion()
	portion.Text = "New Paragraph"
	dto := NewParagraph()
	dto.Alignment = "Right"
	dto.PortionList = []IPortion{portion}
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	dto = NewParagraph()
	dto.Alignment = "Center"
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount+1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}
}

/*
   Test for layout slide portions
*/
func TestLayoutSlidePortions(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionCount int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}

	dto := NewPortion()
	dto.Text = "New portion"
	dto.FontBold = "True"
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetText() != dto.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	dto2 := NewPortion()
	dto2.Text = "Updated portion"
	dto2.FontHeight = 22
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount+1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetFontHeight() != dto2.GetFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.GetFontHeight(), portion.GetFontHeight())
		return
	}
	if portion.GetText() != dto2.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount+1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}
}

/*
   Test for layout slide animation
*/
func TestLayoutSlideAnimation(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	effect1 := NewEffect()
	effect1.Type_ = "Blink"
	effect1.ShapeIndex = 2
	effect2 := NewEffect()
	effect2.Type_ = "Appear"
	effect2.ShapeIndex = 3
	dto := NewSlideAnimation()
	dto.MainSequence = []IEffect{effect1, effect2}
	animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	var shapeIndex int32 = 3
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "layoutSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence())-1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence())-1, len(animation.GetMainSequence()))
		return
	}
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}

	animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
}

/*
   Test for deleting layout slide
*/
func TestSlideDeleteUnused(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	layoutSlidesBefore, _, e := c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(layoutSlidesBefore.GetSlideList()) != 11 {
		t.Errorf("Expected %v, but was %v", 11, layoutSlidesBefore.GetSlideList())
		return
	}

	layoutSlidesAfter, _, e := c.SlidesApi.DeleteUnusedLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(layoutSlidesAfter.GetSlideList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, layoutSlidesAfter.GetSlideList())
		return
	}
}

/*
   Test for deleting layout slide from request
*/
func TestLayoutSlideDeleteUnusedOnline(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteUnusedLayoutSlidesOnline(source, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if resultStat.Size() == 0 {
		t.Errorf("Wrong response size.")
		return
	}
}

/*
   Test for get shapes
*/
func TestShapesGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get shapes by type
*/
func TestShapesGetByType(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "Chart")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get sub-shapes
*/
func TestSubShapesGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	path := "4/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shapes, _, e := c.SlidesApi.GetSubshapes(fileName, slideIndex, path, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get shape
*/
func TestShapeGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if shape.GetType() != "Chart" {
		t.Errorf("Expected %v, but was %v", "Chart", shape.GetType())
		return
	}
}

/*
   Test for get sub-shape
*/
func TestSubShapeGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1
	path := "4/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if shape.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Chart", shape.GetType())
		return
	}
}

/*
   Test for add shape
*/
func TestShapeAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	dto.ShapeType = "Callout1"
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty shape
*/
func TestShapeEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("Shape with undefined type should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty graphical object
*/
func TestGraphicalObjectEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewGraphicalObject()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("GraphicalObject should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add picture frame
*/
func TestPictureFrameAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewPictureFrame()
	fill := NewPictureFill()
	fill.Base64Data = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXY5g+ffp/AAZTAsWGL27gAAAAAElFTkSuQmCC"
	dto.PictureFillFormat = fill
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IPictureFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty picture frame
*/
func TestPictureFrameEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewPictureFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("PictureFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add audio frame
*/
func TestAudioFrameAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewAudioFrame()
	dto.Base64Data = "bXAzc2FtcGxl"
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IAudioFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty audio frame
*/
func TestAudioFrameEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewAudioFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("AudioFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add video frame
*/
func TestVideoFrameAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewVideoFrame()
	dto.Base64Data = "bXAzc2FtcGxl"
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IVideoFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty video frame
*/
func TestVideoFrameEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewVideoFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("VideoFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty OLE object frame
*/
func TestOleObjectFrameEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewOleObjectFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("OleObjectFrame should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add smart art
*/
func TestSmartArtAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSmartArt()
	dto.X = 0
	dto.Y = 0
	dto.Width = 300
	dto.Height = 200
	dto.Layout = "BasicProcess"
	dto.QuickStyle = "SimpleFill"
	dto.ColorStyle = "ColoredFillAccent1"
	node1 := NewSmartArtNode()
	node1.Text = "First"
	node1.OrgChartLayout = "Initial"
	subNode1 := NewSmartArtNode()
	subNode1.Text = "SubFirst"
	subNode1.OrgChartLayout = "Initial"
	node1.Nodes = []ISmartArtNode{subNode1}
	node2 := NewSmartArtNode()
	node2.Text = "Second"
	node2.OrgChartLayout = "Initial"
	dto.Nodes = []ISmartArtNode{node1, node2}
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(ISmartArt)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add smart art text formatting
*/
func TestSmartArtTextFormatting(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion := NewPortion()
	portion.SetText("New text")
	portion.SetFontHeight(24)
	portion.SetFontBold("True")
	portion.SetSpacing(3)
	fillFormat := NewSolidFill()
	fillFormat.SetColor("#FFFFFF00")
	portion.FillFormat = fillFormat

	targetNodePath := "1/nodes/1/nodes"
	var slideIndex int32 = 7
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	response, _, e := c.SlidesApi.UpdateSubshapePortion(fileName, slideIndex, targetNodePath, shapeIndex, paragraphIndex, portionIndex, portion, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != portion.GetText() {
		t.Errorf("Expected %v, but was %v", portion.GetText(), response.GetText())
		return
	}

	if response.GetFontHeight() != portion.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", portion.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetFontBold() != portion.GetFontBold() {
		t.Errorf("Expected %v, but was %v", portion.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetSpacing() != portion.GetSpacing() {
		t.Errorf("Expected %v, but was %v", portion.GetSpacing(), response.GetSpacing())
		return
	}

	if response.GetFillFormat().(ISolidFill).GetColor() != fillFormat.GetColor() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetColor(), response.GetFillFormat().(ISolidFill).GetColor())
		return
	}
}

/*
   Test for add empty smart art
*/
func TestSmartArtEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSmartArt()
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(ISmartArt)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty chart
   See Chart tests for non-empty chart examples
*/
func TestChartEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewChart()
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IChart)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}

	dto = NewChart()
	dto.ChartType = "ClusteredColumn"
	dto.X = 100
	dto.Y = 100
	dto.Width = 400
	dto.Height = 400

	title := NewChartTitle()
	title.HasTitle = true
	title.Text = "Column Chart"
	dto.Title = title

	category1 := NewChartCategory()
	category1.Value = "Category1"
	category2 := NewChartCategory()
	category2.Value = "Category2"
	category3 := NewChartCategory()
	category3.Value = "Category3"
	dto.Categories = []IChartCategory{category1, category2, category3}

	series1 := NewOneValueSeries()
	dataPoint11 := NewOneValueChartDataPoint()
	dataPoint11.Value = 20

	dataPoint12 := NewOneValueChartDataPoint()
	dataPoint12.Value = 50

	dataPoint13 := NewOneValueChartDataPoint()
	dataPoint13.Value = 30
	series1.DataPoints = []IOneValueChartDataPoint{dataPoint11, dataPoint12, dataPoint13}

	series2 := NewOneValueSeries()
	dataPoint21 := NewOneValueChartDataPoint()
	dataPoint21.Value = 30

	dataPoint22 := NewOneValueChartDataPoint()
	dataPoint22.Value = 10

	dataPoint23 := NewOneValueChartDataPoint()
	dataPoint23.Value = 60
	series2.DataPoints = []IOneValueChartDataPoint{dataPoint21, dataPoint22, dataPoint23}
	dto.Series = []ISeries{series1, series2}

	result, _, e = c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
	}
}

/*
   Test for add table
*/
func TestTableAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewTable()
	dto.X = 30
	dto.Y = 20
	dto.Style = "MediumStyle2Accent1"
	row1 := NewTableRow()
	cell11 := NewTableCell()
	cell11.Text = "0.1"
	cell12 := NewTableCell()
	cell12.Text = "0.2"
	cell13 := NewTableCell()
	cell13.Text = "0.3"
	cell14 := NewTableCell()
	cell14.Text = "0.4"
	row1.Cells = []ITableCell{cell11, cell12, cell13, cell14}
	row2 := NewTableRow()
	cell21 := NewTableCell()
	cell21.Text = "1"
	cell22 := NewTableCell()
	cell22.Text = "2-3"
	cell22.ColSpan = 2
	cell22.RowSpan = 2
	cell24 := NewTableCell()
	cell24.Text = "4"
	row2.Cells = []ITableCell{cell21, cell22, cell24}
	row3 := NewTableRow()
	cell31 := NewTableCell()
	cell31.Text = "first"
	cell32 := NewTableCell()
	cell32.Text = "last"
	row3.Cells = []ITableCell{cell31, cell32}
	row4 := NewTableRow()
	cell41 := NewTableCell()
	cell41.Text = "3.1"
	cell42 := NewTableCell()
	cell42.Text = "3.2"
	cell43 := NewTableCell()
	cell43.Text = "3.3"
	cell44 := NewTableCell()
	cell44.Text = "3.4"
	row4.Cells = []ITableCell{cell41, cell42, cell43, cell44}
	row5 := NewTableRow()
	cell51 := NewTableCell()
	cell51.Text = "4.1"
	cell52 := NewTableCell()
	cell52.Text = "4.2"
	cell53 := NewTableCell()
	cell53.Text = "4.3"
	cell54 := NewTableCell()
	cell54.Text = "4.4"
	row5.Cells = []ITableCell{cell51, cell52, cell53, cell54}
	dto.Rows = []ITableRow{row1, row2, row3, row4, row5}
	column1 := NewTableColumn()
	column1.Width = 100
	column2 := NewTableColumn()
	column2.Width = 100
	column3 := NewTableColumn()
	column3.Width = 100
	column4 := NewTableColumn()
	column4.Width = 100
	dto.Columns = []ITableColumn{column1, column2, column3, column4}
	dto.FirstRow = true
	dto.HorizontalBanding = true
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(ITable)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty table
*/
func TestTableEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewTable()
	_, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("Table with undefinined cell data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty group shape
*/
func TestGroupShapeEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewGroupShape()
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IGroupShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add connector
*/
func TestConnectorAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewConnector()
	dto.ShapeType = "BentConnector3"
	start := NewResourceUri()
	start.Href = "https://api.aspose.cloud/v3.0/slides/myPresentation.pptx/slides/1/shapes/1"
	dto.StartShapeConnectedTo = start
	end := NewResourceUri()
	end.Href = "https://api.aspose.cloud/v3.0/slides/myPresentation.pptx/slides/1/shapes/2"
	dto.EndShapeConnectedTo = end
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IConnector)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty connector
*/
func TestConnectorEmpty(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewConnector()
	result, _, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IConnector)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for create shubshape
*/
func TestCreateSubShape(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	dto.SetShapeType("Rectangle")
	dto.SetX(200)
	dto.SetY(200)
	dto.SetWidth(50)
	dto.SetHeight(50)

	shapePath := "4/shapes"
	shapeBase, _, e := c.SlidesApi.CreateSubshape(fileName, slideIndex, shapePath, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeBase.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Shape", shapeBase.GetType())
		return
	}
}

/*
   Test for update shape
*/
func TestShapeUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	dto.SetWidth(200)
	dto.SetHeight(200)
	fillFormat := NewSolidFill()
	fillFormat.SetColor("#FFF5FF8A")
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetWidth() != dto.GetWidth() {
		t.Errorf("Expected %v, but was %v", dto.GetWidth(), response.GetWidth())
		return
	}

	if response.GetHeight() != dto.GetHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetHeight(), response.GetHeight())
		return
	}

	if response.GetFillFormat().(ISolidFill).GetColor() != fillFormat.GetColor() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetColor(), response.GetFillFormat().(ISolidFill).GetColor())
		return
	}
}

/*
   Test for update sub-shape
*/
func TestSubShapeUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1
	path := "4/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	dto.SetWidth(200)
	dto.SetHeight(200)
	fillFormat := NewGradientFill()
	fillFormat.SetShape("Linear")
	fillFormat.SetDirection("FromCorner1")
	gradientStop1 := NewGradientFillStop()
	gradientStop1.Color = "#FFF5FF8A"
	gradientStop1.SetPosition(0)
	gradientStop2 := NewGradientFillStop()
	gradientStop2.Color = "#FFF5FF8A"
	gradientStop2.SetPosition(1)
	fillFormat.SetStops([]IGradientFillStop{gradientStop1, gradientStop2})
	dto.SetFillFormat(fillFormat)

	response, _, e := c.SlidesApi.UpdateSubshape(fileName, slideIndex, path, shapeIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetWidth() != dto.GetWidth() {
		t.Errorf("Expected %v, but was %v", dto.GetWidth(), response.GetWidth())
		return
	}

	if response.GetHeight() != dto.GetHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetHeight(), response.GetHeight())
		return
	}

	if response.GetFillFormat().GetType() != fillFormat.GetType() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for delete shapes
*/
func TestShapesDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete shapes by indexes
*/
func TestShapesDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	shapesIndexes := []int32{2}

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, shapesIndexes, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shapes
*/
func TestSubShapesDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	path := "4/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapes(fileName, slideIndex, path, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shapes by indexes
*/
func TestSubShapesDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	path := "4/shapes"
	shapesIndexes := []int32{2}

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapes(fileName, slideIndex, path, shapesIndexes, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete shape
*/
func TestShapeDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 4

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteShape(fileName, slideIndex, shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shape
*/
func TestSubShapeDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1
	path := "4/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshape(fileName, slideIndex, path, shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for align shapes
*/
func TestShapesAlign(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shape1Index int32 = 1
	var shape2Index int32 = 2
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape11, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape12, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape11.GetX() == shape12.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape11.GetX(), shape12.GetX())
		return
	}
	if shape11.GetY() == shape12.GetY() {
		t.Errorf("Wrong Y value. Expected not %v but was %v.", shape11.GetY(), shape12.GetY())
		return
	}

	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignTop", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape21, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape22, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape21.GetX() == shape22.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape21.GetX(), shape22.GetX())
		return
	}
	if math.Abs(shape21.GetY()-shape22.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape21.GetY(), shape22.GetY())
		return
	}

	var alignToSlide bool = true
	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignLeft", &alignToSlide, []int32{1, 2}, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape31, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape32, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if math.Abs(shape31.GetX()-shape32.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", shape31.GetX(), shape32.GetX())
		return
	}
	if math.Abs(shape31.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", 0, shape31.GetX())
		return
	}
	if math.Abs(shape31.GetY()-shape32.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape31.GetY(), shape32.GetY())
		return
	}
}

/*
   Test for align shapes in group
*/
func TestShapesAlignGroup(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	var path string = "4/shapes"
	var shape1Index int32 = 1
	var shape2Index int32 = 2
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape11, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape12, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape11.GetX() == shape12.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape11.GetX(), shape12.GetX())
		return
	}
	if shape11.GetY() == shape12.GetY() {
		t.Errorf("Wrong Y value. Expected not %v but was %v.", shape11.GetY(), shape12.GetY())
		return
	}

	_, _, e = c.SlidesApi.AlignSubshapes(fileName, slideIndex, path, "AlignTop", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape21, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape22, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape21.GetX() == shape22.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape21.GetX(), shape22.GetX())
		return
	}
	if math.Abs(shape21.GetY()-shape22.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape21.GetY(), shape22.GetY())
		return
	}

	var alignToSlide bool = true
	_, _, e = c.SlidesApi.AlignSubshapes(fileName, slideIndex, path, "AlignLeft", &alignToSlide, []int32{1, 2}, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape31, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape1Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape32, _, e := c.SlidesApi.GetSubshape(fileName, slideIndex, path, shape2Index, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if math.Abs(shape31.GetX()-shape32.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", shape31.GetX(), shape32.GetX())
		return
	}
	if math.Abs(shape31.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", 0, shape31.GetX())
		return
	}
	if math.Abs(shape31.GetY()-shape32.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape31.GetY(), shape32.GetY())
		return
	}
}

/*
   Test for Get shape geometry paths
*/
func TestShapeGeometryGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	paths, _, e := c.SlidesApi.GetShapeGeometryPath(fileName, 4, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paths.GetPaths() == nil {
		t.Errorf("Error: paths should not be nil.")
		return
	}
	if len(paths.GetPaths()) != 1 {
		t.Errorf("Wrong paths count. Expected %v but was %v.", 1, len(paths.GetPaths()))
		return
	}
}

/*
   Test for set shape geometry paths
*/
func TestShapeGeometrySet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewGeometryPaths()
	path := NewGeometryPath()
	startSegment := NewMoveToPathSegment()
	startSegment.X = 0
	startSegment.Y = 0
	line1 := NewLineToPathSegment()
	line1.X = 0
	line1.Y = 200
	line2 := NewLineToPathSegment()
	line2.X = 200
	line2.Y = 300
	line3 := NewLineToPathSegment()
	line3.X = 400
	line3.Y = 200
	line4 := NewLineToPathSegment()
	line4.X = 400
	line4.Y = 0
	endSegment := NewClosePathSegment()
	path.PathData = []IPathSegment{startSegment, line1, line2, line3, line4, endSegment}
	dto.Paths = []IGeometryPath{path}
	shape, _, e := c.SlidesApi.SetShapeGeometryPath(fileName, 4, 1, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape == nil {
		t.Errorf("Error: shape should not be nil.")
		return
	}
}

/*
   Test for Zoom Frame add
*/
func TestZoomFrameAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewZoomFrame()
	dto.SetX(20)
	dto.SetY(20)
	dto.SetWidth(200)
	dto.SetHeight(100)
	dto.SetTargetSlideIndex(2)

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "ZoomFrame" {
		t.Errorf("Expected %v, but was %v", "ZoomFrame", response.GetType())
		return
	}

	if response.(IZoomFrame).GetTargetSlideIndex() != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.(IZoomFrame).GetTargetSlideIndex())
		return
	}
}

/*
   Test for Section Zoom Frame add
*/
func TestSectionZoomFrameAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSectionZoomFrame()
	dto.SetX(20)
	dto.SetY(20)
	dto.SetWidth(200)
	dto.SetHeight(100)
	dto.SetTargetSectionIndex(2)

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "SectionZoomFrame" {
		t.Errorf("Expected %v, but was %v", "SectionZoomFrame", response.GetType())
		return
	}

	if response.(ISectionZoomFrame).GetTargetSectionIndex() != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.(ISectionZoomFrame).GetTargetSectionIndex())
		return
	}
}

/*
   Test for OleObjectFrame add by link
*/
func TestOleObjectFrameAddByLink(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewOleObjectFrame()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(200)
	dto.SetHeight(200)
	dto.SetLinkPath("OleObject.xlsx")
	dto.SetObjectProgId("Excel.Sheet.8")

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "OleObjectFrame" {
		t.Errorf("Expected %v, but was %v", "OleObjectFrame", response.GetType())
		return
	}

	if response.(IOleObjectFrame).GetLinkPath() != dto.GetLinkPath() {
		t.Errorf("Expected %v, but was %v", dto.GetLinkPath(), response.(IOleObjectFrame).GetLinkPath())
		return
	}
}

/*
   Test for OleObjectFrame add embedded
*/
func TestOleObjectFrameAddEmbedded(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	source, e := ioutil.ReadFile("TestData/oleObject.xlsx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewOleObjectFrame()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(200)
	dto.SetHeight(200)
	dto.SetEmbeddedFileBase64Data(base64.StdEncoding.EncodeToString(source))
	dto.SetEmbeddedFileExtension("xlsx")

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "OleObjectFrame" {
		t.Errorf("Expected %v, but was %v", "OleObjectFrame", response.GetType())
		return
	}

	if response.(IOleObjectFrame).GetEmbeddedFileBase64Data() == "" {
		t.Errorf("EmbeddedFileBase64Data is not defined.")
		return
	}
}

/*
   Test for GroupShape add
*/
func TestGroupShapeAdd(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 5
	path := "1/shapes"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewGroupShape()
	c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")

	shape1 := NewShape()
	shape1.SetShapeType("Rectangle")
	shape1.SetX(50)
	shape1.SetY(400)
	shape1.SetWidth(50)
	shape1.SetHeight(50)

	shape2 := NewShape()
	shape2.SetShapeType("Ellipse")
	shape2.SetX(150)
	shape2.SetY(400)
	shape2.SetWidth(50)
	shape2.SetHeight(50)

	shape3 := NewShape()
	shape3.SetShapeType("Triangle")
	shape3.SetX(250)
	shape3.SetY(400)
	shape3.SetWidth(50)
	shape3.SetHeight(50)

	c.SlidesApi.CreateSubshape(fileName, slideIndex, path, shape1, nil, nil, password, folderName, "")
	c.SlidesApi.CreateSubshape(fileName, slideIndex, path, shape2, nil, nil, password, folderName, "")
	c.SlidesApi.CreateSubshape(fileName, slideIndex, path, shape3, nil, nil, password, folderName, "")

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(shapes.GetShapesLinks()))
		return
	}

	subShapes, _, e := c.SlidesApi.GetSubshapes(fileName, slideIndex, path, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(subShapes.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(subShapes.GetShapesLinks()))
		return
	}
}

/*
   Test for Get chart
*/
func TestChartGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, 3, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(chart.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(chart.(IChart).GetSeries()))
		return
	}
	if len(chart.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(chart.(IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart
*/
func TestChartCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300
	series1 := NewOneValueSeries()
	series1.Name = "Series1"
	point11 := NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []IOneValueChartDataPoint{point11, point12, point13}
	series2 := NewOneValueSeries()
	series2.Name = "Series2"
	point21 := NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []ISeries{series1, series2}
	category1 := NewChartCategory()
	category1.Value = "Category1"
	category2 := NewChartCategory()
	category2.Value = "Category2"
	category3 := NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for update chart
*/
func TestChartUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300
	series1 := NewOneValueSeries()
	series1.Name = "Series1"
	point11 := NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []IOneValueChartDataPoint{point11, point12, point13}
	series2 := NewOneValueSeries()
	series2.Name = "Series2"
	point21 := NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []ISeries{series1, series2}
	category1 := NewChartCategory()
	category1.Value = "Category1"
	category2 := NewChartCategory()
	category2.Value = "Category2"
	category3 := NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.UpdateShape(fileName, 3, 1, chart, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart series
*/
func TestChartSeriesCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	series := NewOneValueSeries()
	series.Name = "Series1"
	point1 := NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := NewOneValueChartDataPoint()
	point3.Value = 14
	point4 := NewOneValueChartDataPoint()
	point4.Value = 70
	series.DataPoints = []IOneValueChartDataPoint{point1, point2, point3, point4}
	result, _, e := c.SlidesApi.CreateChartSeries(fileName, 3, 1, series, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 4 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 4, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for update chart series
*/
func TestChartSeriesUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	series := NewOneValueSeries()
	series.Name = "Series1"
	point1 := NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := NewOneValueChartDataPoint()
	point3.Value = 14
	point4 := NewOneValueChartDataPoint()
	point4.Value = 70
	series.DataPoints = []IOneValueChartDataPoint{point1, point2, point3, point4}
	result, _, e := c.SlidesApi.UpdateChartSeries(fileName, 3, 1, 2, series, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for delete chart series
*/
func TestChartSeriesDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartSeries(fileName, 3, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart category
*/
func TestChartCategoryCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	category := NewChartCategory()
	category.Value = "NewCategory"
	point1 := NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := NewOneValueChartDataPoint()
	point3.Value = 14
	category.DataPoints = []IOneValueChartDataPoint{point1, point2, point3}
	result, _, e := c.SlidesApi.CreateChartCategory(fileName, 3, 1, category, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 5 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 5, len(result.(IChart).GetCategories()))
		return
	}
	if len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()) != 5 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 5, len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()[4].GetValue() != category.DataPoints[0].GetValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].GetValue(), result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()[4].GetValue())
		return
	}
}

/*
   Test for update chart category
*/
func TestChartCategoryUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	category := NewChartCategory()
	category.Value = "NewCategory"
	point1 := NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := NewOneValueChartDataPoint()
	point3.Value = 14
	category.DataPoints = []IOneValueChartDataPoint{point1, point2, point3}
	result, _, e := c.SlidesApi.UpdateChartCategory(fileName, 3, 1, 2, category, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
	if len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()[1].GetValue() != category.DataPoints[0].GetValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].GetValue(), result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()[1].GetValue())
		return
	}
}

/*
   Test for delete chart category
*/
func TestChartCategoryDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartCategory(fileName, 3, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).GetCategories()))
		return
	}
	if len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()) != 3 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()[0].(IOneValueSeries).GetDataPoints()))
		return
	}
}

/*
   Test for create chart data point
*/
func TestChartDataPointCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	point := NewOneValueChartDataPoint()
	point.Value = 40
	_, response, e := c.SlidesApi.CreateChartDataPoint(fileName, 3, 1, 2, point, "password", folderName, "")
	if e == nil {
		t.Errorf("Must have failed because adding data points only works with Scatter & Bubble charts")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for update chart data point
*/
func TestChartDataPointUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	point := NewOneValueChartDataPoint()
	point.Value = 40
	result, _, e := c.SlidesApi.UpdateChartDataPoint(fileName, 3, 1, 2, 2, point, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
	if len(result.(IChart).GetSeries()[1].(IOneValueSeries).GetDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(IChart).GetSeries()[1].(IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(IChart).GetSeries()[1].(IOneValueSeries).GetDataPoints()[1].GetValue() != point.Value {
		t.Errorf("Wrong data point value. Expected %v but was %v.", point.Value, result.(IChart).GetSeries()[1].(IOneValueSeries).GetDataPoints()[1].GetValue())
		return
	}
}

/*
   Test for delete chart data point
*/
func TestChartDataPointDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartDataPoint(fileName, 3, 1, 2, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
	if result.(IChart).GetSeries()[1].(IOneValueSeries).GetDataPoints()[1] != nil {
		t.Errorf("Data point should be nil.")
		return
	}
}

/*
   Test for sunburst chart
*/
func TestChartSunburst(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := NewChart()
	chart.ChartType = "Sunburst"
	chart.Width = 400
	chart.Height = 300
	series1 := NewOneValueSeries()
	series1.Name = "Series1"
	point1 := NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := NewOneValueChartDataPoint()
	point3.Value = 70
	point4 := NewOneValueChartDataPoint()
	point4.Value = 60
	series1.DataPoints = []IOneValueChartDataPoint{point1, point2, point3, point4}
	chart.Series = []ISeries{series1}
	category1 := NewChartCategory()
	category1.Value = "Leaf1"
	category1.Level = 3
	category1.ParentCategories = []string{"Branch1", "Stem1"}
	category2 := NewChartCategory()
	category2.Value = "Leaf2"
	category2.Level = 3
	category2.ParentCategories = []string{"Branch1", "Stem1"}
	category3 := NewChartCategory()
	category3.Value = "Branch2"
	category3.Level = 2
	category3.ParentCategories = []string{"Stem1"}
	category4 := NewChartCategory()
	category4.Value = "Stem2"
	category4.Level = 1
	chart.Categories = []IChartCategory{category1, category2, category3, category4}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).GetSeries()) != 1 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 1, len(result.(IChart).GetSeries()))
		return
	}
	if len(result.(IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).GetCategories()))
		return
	}
}

/*
   Test for Multi level category axis
*/
func TestMultiLevelCategoryAxis(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewChart()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(500)
	dto.SetHeight(400)
	dto.SetChartType("ClusteredColumn")

	series1 := NewOneValueSeries()
	series1.SetType("ClusteredColumn")
	dataPoint1 := NewOneValueChartDataPoint()
	dataPoint1.SetValue(1)
	dataPoint2 := NewOneValueChartDataPoint()
	dataPoint2.SetValue(2)
	dataPoint3 := NewOneValueChartDataPoint()
	dataPoint3.SetValue(3)
	dataPoint4 := NewOneValueChartDataPoint()
	dataPoint4.SetValue(4)
	dataPoint5 := NewOneValueChartDataPoint()
	dataPoint5.SetValue(5)
	dataPoint6 := NewOneValueChartDataPoint()
	dataPoint6.SetValue(6)
	dataPoint7 := NewOneValueChartDataPoint()
	dataPoint7.SetValue(7)
	dataPoint8 := NewOneValueChartDataPoint()
	dataPoint8.SetValue(8)
	series1.DataPoints = []IOneValueChartDataPoint{dataPoint1, dataPoint2, dataPoint3, dataPoint4, dataPoint5, dataPoint6, dataPoint7, dataPoint8}
	dto.Series = []ISeries{series1}

	category1 := NewChartCategory()
	category1.SetValue("Category1")
	category1.SetParentCategories([]string{"Sub-category 1", "Root 1"})
	category2 := NewChartCategory()
	category2.SetValue("Category2")
	category3 := NewChartCategory()
	category3.SetValue("Category3")
	category3.SetParentCategories([]string{"Sub-category 2"})
	category4 := NewChartCategory()
	category4.SetValue("Category4")
	category5 := NewChartCategory()
	category5.SetValue("Category5")
	category5.SetParentCategories([]string{"Sub-category 3", "Root 2"})
	category6 := NewChartCategory()
	category6.SetValue("Category6")
	category7 := NewChartCategory()
	category7.SetValue("Category7")
	category7.SetParentCategories([]string{"Sub-category 4"})
	category8 := NewChartCategory()
	category8.SetValue("Category8")
	dto.Categories = []IChartCategory{category1, category2, category3, category4, category5, category6, category7, category8}

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.(IChart).GetChartType() != dto.GetChartType() {
		t.Errorf("Expected %v, but was %v", dto.GetChartType(), response.(IChart).GetChartType())
		return
	}

	if len(response.(IChart).GetSeries()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response)
		return
	}

	if len(response.(IChart).GetCategories()) != 8 {
		t.Errorf("Expected %v, but was %v", 8, response)
		return
	}
}

/*
   Test for Hide chart legend
*/
func TestTemplate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart.(IChart).GetLegend().SetHasLegend(false)

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, chart, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.(IChart).GetLegend().GetHasLegend() != false {
		t.Errorf("Expected %v, but was %v", false, true)
		return
	}
}

/*
   Test for chart grid lines format
*/
func TestChartGridLinesFormat(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	horizontalAxis := NewAxis()
	horizontalAxis.MajorGridLinesFormat = NewChartLinesFormat()
	horzMajLineFormat := NewLineFormat()
	horzMajLineFormat.FillFormat = NewNoFill()
	horizontalAxis.MajorGridLinesFormat.SetLineFormat(horzMajLineFormat)

	horizontalAxis.MinorGridLinesFormat = NewChartLinesFormat()
	horzMinorLineFormat := NewLineFormat()
	horzMinorLineFormat.FillFormat = NewSolidFill()
	horzMinorLineFormat.FillFormat.(ISolidFill).SetColor("Black")
	horizontalAxis.MinorGridLinesFormat.SetLineFormat(horzMinorLineFormat)

	verticalAxis := NewAxis()
	verticalAxis.MajorGridLinesFormat = NewChartLinesFormat()
	vertMajorLineFormat := NewLineFormat()
	gradientFill := NewGradientFill()
	gradientFill.Direction = "FromCorner1"
	stop1 := NewGradientFillStop()
	stop1.Color = "White"
	stop1.Position = 0
	stop2 := NewGradientFillStop()
	stop2.Color = "Black"
	stop2.Position = 1
	gradientFill.SetStops([]IGradientFillStop{stop1, stop2})
	vertMajorLineFormat.FillFormat = gradientFill
	verticalAxis.MajorGridLinesFormat.SetLineFormat(vertMajorLineFormat)

	verticalAxis.MinorGridLinesFormat = NewChartLinesFormat()
	vertMinorLineFormat := NewLineFormat()
	vertMinorLineFormat.FillFormat = NewNoFill()
	verticalAxis.MinorGridLinesFormat.SetLineFormat(vertMinorLineFormat)

	axes := NewAxes()
	axes.SetHorizontalAxis(horizontalAxis)
	axes.SetVerticalAxis(verticalAxis)

	chart.(IChart).SetAxes(axes)

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, chart, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	resonseFill := response.(IChart).GetAxes().GetHorizontalAxis().GetMajorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "NoFill" {
		t.Errorf("Expected %v, but was %v", "NoFill", resonseFill)
		return
	}

	resonseFill = response.(IChart).GetAxes().GetHorizontalAxis().GetMinorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", resonseFill)
		return
	}

	resonseFill = response.(IChart).GetAxes().GetVerticalAxis().GetMajorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "Gradient" {
		t.Errorf("Expected %v, but was %v", "NoFill", resonseFill)
		return
	}

	resonseFill = response.(IChart).GetAxes().GetVerticalAxis().GetMinorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "NoFill" {
		t.Errorf("Expected %v, but was %v", "Solid", resonseFill)
		return
	}
}

/*
   Test for chart series groups
*/
func TestChartSeriesGroups(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1
	var seriesGroupIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(chart.(IChart).GetSeriesGroups()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(chart.(IChart).GetSeriesGroups()))
		return
	}

	seriesGroup := chart.(IChart).GetSeriesGroups()[0]
	seriesGroup.SetOverlap(10)

	chart, _, e = c.SlidesApi.SetChartSeriesGroup(fileName, slideIndex, shapeIndex, seriesGroupIndex, seriesGroup, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	seriesGroup = chart.(IChart).GetSeriesGroups()[0]

	if seriesGroup.GetOverlap() != 10 {
		t.Errorf("Expected %v, but was %v", 10, seriesGroup.GetOverlap())
		return
	}
}

/*
   Test for chart legend
*/
func TestChartLegend(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	legendDto := NewLegend()
	legendDto.Overlay = true
	solidFill := NewSolidFill()
	solidFill.Color = "#77CEF9"
	legendDto.FillFormat = solidFill

	response, _, e := c.SlidesApi.SetChartLegend(fileName, slideIndex, shapeIndex, legendDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetOverlay() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetOverlay())
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat())
		return
	}
}

/*
   Test for setchart axis
*/
func TestChartAxisSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	axisDto := NewAxis()
	axisDto.HasTitle = true
	axisDto.IsAutomaticMaxValue = false
	axisDto.MaxValue = 10

	response, _, e := c.SlidesApi.SetChartAxis(fileName, slideIndex, shapeIndex, "VerticalAxis", axisDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetHasTitle() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetHasTitle())
		return
	}

	if response.GetIsAutomaticMaxValue() != false {
		t.Errorf("Expected %v, but was %v", false, response.GetIsAutomaticMaxValue())
		return
	}

	if response.GetMaxValue() != 10 {
		t.Errorf("Expected %v, but was %v", 10, response.GetMaxValue())
		return
	}
}

/*
   Test for chart walls
*/
func TestChartWallSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 8
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	wallDto := NewChartWall()
	fillFormat := NewSolidFill()
	fillFormat.SetColor("#77CEF9")
	wallDto.SetFillFormat(fillFormat)

	response, _, e := c.SlidesApi.SetChartWall(fileName, slideIndex, shapeIndex, "BackWall", wallDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for shape line format
*/
func TestShapeFormatLine(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	lineFormat := NewLineFormat()
	lineFormat.Style = "ThickThin"
	lineFormat.Width = 7
	lineFormat.DashStyle = "Dash"
	dto.LineFormat = lineFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetLineFormat().GetWidth() != dto.GetLineFormat().GetWidth() {
		t.Errorf("Wrong line width. Expected %v but was %v.", dto.GetLineFormat().GetWidth(), result.GetLineFormat().GetWidth())
		return
	}
}

/*
   Test for shape fill format
*/
func TestShapeFormatFill(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	fillFormat := NewSolidFill()
	fillFormat.Color = "#FFFFFF00"
	dto.FillFormat = fillFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	_, ok = result.GetFillFormat().(ISolidFill)
	if !ok {
		t.Errorf("Wrong fill type.")
		return
	}
	if result.GetFillFormat().(ISolidFill).GetColor() != dto.GetFillFormat().(ISolidFill).GetColor() {
		t.Errorf("Wrong fill color. Expected %v but was %v.", dto.GetFillFormat().(ISolidFill).GetColor(), result.GetFillFormat().(ISolidFill).GetColor())
		return
	}
}

/*
   Test for shape effect format
*/
func TestShapeFormatEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	effectFormat := NewEffectFormat()
	innerShadow := NewInnerShadowEffect()
	innerShadow.Direction = 35
	innerShadow.BlurRadius = 30
	innerShadow.Distance = 40
	innerShadow.ShadowColor = "#FFFFFF00"
	effectFormat.InnerShadow = innerShadow
	dto.EffectFormat = effectFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetEffectFormat().GetInnerShadow().GetDirection() != dto.GetEffectFormat().GetInnerShadow().GetDirection() {
		t.Errorf("Wrong inner shadow direction. Expected %v but was %v.", dto.GetEffectFormat().GetInnerShadow().GetDirection(), result.GetEffectFormat().GetInnerShadow().GetDirection())
		return
	}
}

/*
   Test for shape 3D format
*/
func TestShapeFormat3D(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	threeDFormat := NewThreeDFormat()
	threeDFormat.Depth = 4

	bevelTop := NewShapeBevel()
	bevelTop.BevelType = "Circle"
	bevelTop.Height = 6
	bevelTop.Width = 6
	threeDFormat.BevelTop = bevelTop

	camera := NewCamera()
	camera.CameraType = "OrthographicFront"
	threeDFormat.Camera = camera

	lightRig := NewLightRig()
	lightRig.LightType = "ThreePt"
	lightRig.Direction = "Top"
	threeDFormat.LightRig = lightRig
	dto.ThreeDFormat = threeDFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetThreeDFormat().GetDepth() != dto.GetThreeDFormat().GetDepth() {
		t.Errorf("Wrong 3D depth. Expected %v but was %v.", dto.GetThreeDFormat().GetDepth(), result.GetThreeDFormat().GetDepth())
		return
	}
}

/*
   Test for Get hyperlink for shape
*/
func TestHyperlinkGetShape(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape, _, e := c.SlidesApi.GetShape(fileName, 2, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.GetHyperlinkClick() == nil {
		t.Errorf("Error: click hyperlink should not be nil.")
		return
	}
	if shape.GetHyperlinkClick().GetActionType() != "Hyperlink" {
		t.Errorf("Wrong action type. Expected %v but was %v.", "Hyperlink", shape.GetHyperlinkClick().GetActionType())
		return
	}
	if shape.GetHyperlinkMouseOver() != nil {
		t.Errorf("Error: mouse over hyperlink should be nil.")
		return
	}
}

/*
   Test for Get hyperlink for portion
*/
func TestHyperlinkGetPortion(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion, _, e := c.SlidesApi.GetPortion(fileName, 2, 1, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetHyperlinkClick() != nil {
		t.Errorf("Error: click hyperlink should be nil.")
		return
	}
	if portion.GetHyperlinkMouseOver() == nil {
		t.Errorf("Error: mouse over hyperlink should not be nil.")
		return
	}
	if portion.GetHyperlinkMouseOver().GetActionType() != "JumpLastSlide" {
		t.Errorf("Wrong action type. Expected not %v but was %v.", "JumpLastSlide", portion.GetHyperlinkMouseOver().GetActionType())
		return
	}
}

/*
   Test for create hyperlink for shape
*/
func TestHyperlinkCreateShape(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape := NewShape()
	hyperlink := NewHyperlink()
	hyperlink.ActionType = "Hyperlink"
	hyperlink.ExternalUrl = "https://docs.aspose.cloud/slides"
	shape.HyperlinkClick = hyperlink
	updatedShape, _, e := c.SlidesApi.UpdateShape(fileName, 1, 1, shape, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedShape.GetHyperlinkClick() == nil {
		t.Errorf("Error: click hyperlink should not be nil.")
		return
	}
	if updatedShape.GetHyperlinkClick().GetExternalUrl() != shape.GetHyperlinkClick().GetExternalUrl() {
		t.Errorf("Wrong URL. Expected %v but was %v.", shape.GetHyperlinkClick().GetExternalUrl(), updatedShape.GetHyperlinkClick().GetExternalUrl())
		return
	}
}

/*
   Test for create hyperlink for portion
*/
func TestHyperlinkCreatePortion(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewPortion()
	dto.Text = "Link text"
	hyperlink := NewHyperlink()
	hyperlink.ActionType = "JumpLastSlide"
	dto.HyperlinkMouseOver = hyperlink
	updatedPortion, _, e := c.SlidesApi.CreatePortion(fileName, 1, 1, 1, dto, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedPortion.GetHyperlinkMouseOver() == nil {
		t.Errorf("Error: mouse over hyperlink should not be nil.")
		return
	}
	if updatedPortion.GetHyperlinkMouseOver().GetActionType() != "JumpLastSlide" {
		t.Errorf("Wrong action type. Expected %v but was %v.", dto.GetHyperlinkMouseOver().GetActionType(), updatedPortion.GetHyperlinkMouseOver().GetActionType())
		return
	}
}

/*
   Test for delete hyperlink
*/
func TestHyperlinkDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	shape := NewPictureFrame()
	hyperlink := NewHyperlink()
	hyperlink.IsDisabled = true
	shape.HyperlinkClick = hyperlink
	updatedShape, _, e := c.SlidesApi.UpdateShape(fileName, 1, 1, shape, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedShape.GetHyperlinkClick() != nil {
		t.Errorf("Error: click hyperlink should be nil.")
		return
	}
}

/*
   Test for Get math
*/
func TestMathGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion, _, e := c.SlidesApi.GetPortion(fileName, 2, 3, 1, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 3 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 3, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[2].(IFractionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for Get empty math
*/
func TestMathGetNull(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion, _, e := c.SlidesApi.GetPortion(fileName, 2, 1, 1, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() != nil {
		t.Errorf("Error: math paragraph should be nil.")
		return
	}
}

/*
   Test for create math
*/
func TestMathCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewPortion()
	mathParagraph := NewMathParagraph()
	blockElement := NewBlockElement()
	functionElement := NewFunctionElement()
	limitElement := NewLimitElement()
	text1 := NewTextElement()
	text1.Value = "lim"
	limitElement.Base = text1
	text2 := NewTextElement()
	text2.Value = "x->0"
	limitElement.Limit = text2
	functionElement.Name = limitElement
	fractionElement := NewFractionElement()
	sinusElement := NewFunctionElement()
	text3 := NewTextElement()
	text3.Value = "sin"
	sinusElement.Name = text3
	text4 := NewTextElement()
	text4.Value = "x"
	sinusElement.Base = text4
	fractionElement.Numerator = sinusElement
	text5 := NewTextElement()
	text5.Value = "x"
	fractionElement.Denominator = text5
	functionElement.Base = fractionElement
	blockElement.MathElementList = []IMathElement{functionElement}
	mathParagraph.MathBlockList = []IBlockElement{blockElement}
	dto.MathParagraph = mathParagraph
	portion, _, e := c.SlidesApi.CreatePortion(fileName, 1, 1, 1, dto, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 1 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[0].(IFunctionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for update math
*/
func TestMathUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewPortion()
	mathParagraph := NewMathParagraph()
	blockElement := NewBlockElement()
	functionElement := NewFunctionElement()
	limitElement := NewLimitElement()
	text1 := NewTextElement()
	text1.Value = "lim"
	limitElement.Base = text1
	text2 := NewTextElement()
	text2.Value = "x->0"
	limitElement.Limit = text2
	functionElement.Name = limitElement
	fractionElement := NewFractionElement()
	sinusElement := NewFunctionElement()
	text3 := NewTextElement()
	text3.Value = "sin"
	sinusElement.Name = text3
	text4 := NewTextElement()
	text4.Value = "x"
	sinusElement.Base = text4
	fractionElement.Numerator = sinusElement
	text5 := NewTextElement()
	text5.Value = "x"
	fractionElement.Denominator = text5
	functionElement.Base = fractionElement
	blockElement.MathElementList = []IMathElement{functionElement}
	mathParagraph.MathBlockList = []IBlockElement{blockElement}
	dto.MathParagraph = mathParagraph
	portion, _, e := c.SlidesApi.UpdatePortion(fileName, 2, 3, 1, 1, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 1 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[0].(IFunctionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for download math
*/
func TestMathDownload(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	mathMl, _, e := c.SlidesApi.DownloadPortionAsMathMl(fileName, 2, 3, 1, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	mathMlStat, e := os.Stat(mathMl.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if mathMlStat.Size() <= 0 {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", 0, mathMlStat.Size())
		return
	}
}

/*
   Test for download empty math
*/
func TestMathDownloadNull(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, response, e := c.SlidesApi.DownloadPortionAsMathMl(fileName, 2, 1, 1, 1, "password", folderName, "")
	if e == nil {
		t.Errorf("An ordinary paragraph should not have been converted to MathML.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for save math
*/
func TestMathSave(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	outPath := folderName + "/mathml.xml"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SavePortionAsMathMl(fileName, 2, 3, 1, 1, outPath, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for header/footer on all slides
*/
func TestHeaderFooterSlides(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewHeaderFooter()
	dto.IsFooterVisible = true
	dto.FooterText = "footer"
	dto.IsDateTimeVisible = false
	_, _, e = c.SlidesApi.SetPresentationHeaderFooter(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e := c.SlidesApi.GetSlideHeaderFooter(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on one slide
*/
func TestHeaderFooterSlide(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewHeaderFooter()
	dto.IsFooterVisible = true
	dto.FooterText = "footer"
	dto.IsDateTimeVisible = false
	result, _, e := c.SlidesApi.SetSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on notes slide
*/
func TestHeaderFooterNotesSlide(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewNotesSlideHeaderFooter()
	dto.IsHeaderVisible = true
	dto.FooterText = "footer"
	dto.IsDateTimeVisible = false
	result, _, e := c.SlidesApi.SetNotesSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetNotesSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for Get sections
*/
func TestSectionsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetSections(fileName, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
}

/*
   Test for replace sections
*/
func TestSectionsReplace(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSections()
	section1 := NewSection()
	section1.Name = "Section1"
	section1.FirstSlideIndex = 1
	section2 := NewSection()
	section2.Name = "Section2"
	section2.FirstSlideIndex = 3
	dto.SectionList = []ISection{section1, section2}
	result, _, e := c.SlidesApi.SetSections(fileName, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != len(dto.GetSectionList()) {
		t.Errorf("Wrong section count. Expected %v but was %v.", len(dto.GetSectionList()), len(result.GetSectionList()))
		return
	}
	if len(result.GetSectionList()[0].GetSlideList()) != int(section2.FirstSlideIndex-section1.FirstSlideIndex) {
		t.Errorf("Wrong slide count. Expected %v but was %v.", section2.FirstSlideIndex-section1.FirstSlideIndex, len(result.GetSectionList()[0].GetSlideList()))
		return
	}
}

/*
   Test for create section
*/
func TestSectionsPost(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.CreateSection(fileName, "NewSection", 5, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 4 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 4, len(result.GetSectionList()))
		return
	}
}

/*
   Test for update section
*/
func TestSectionsPut(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var sectionIndex int32 = 2
	sectionName := "UpdatedSection"
	result, _, e := c.SlidesApi.UpdateSection(fileName, sectionIndex, sectionName, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
	if result.GetSectionList()[sectionIndex-1].GetName() != sectionName {
		t.Errorf("Wrong section name. Expected %v but was %v.", sectionName, result.GetSectionList()[sectionIndex-1].GetName())
		return
	}
}

/*
   Test for move section
*/
func TestSectionsMove(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.MoveSection(fileName, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
}

/*
   Test for clear sections
*/
func TestSectionsClear(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSections(fileName, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 0 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 0, len(result.GetSectionList()))
		return
	}
}

/*
   Test for delete sections
*/
func TestSectionsDeleteMany(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSections(fileName, []int32{2, 3}, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 1 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 1, len(result.GetSectionList()))
		return
	}
}

/*
   Test for delete section
*/
func TestSectionsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSection(fileName, 2, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 2 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 2, len(result.GetSectionList()))
		return
	}
}

/*
   Test for builtin property
*/
func TestPropertyBuiltin(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	propertyName := "Author"
	updatedPropertyValue := "New Value"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}

	property := NewDocumentProperty()
	property.Value = updatedPropertyValue
	result, _, e = c.SlidesApi.SetDocumentProperty(fileName, propertyName, property, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}

	_, _, e = c.SlidesApi.DeleteDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e = c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() == updatedPropertyValue {
		t.Errorf("Wrong property value. Expected not %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}
}

/*
   Test for custom property
*/
func TestPropertyCustom(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	propertyName := "CustomProperty2"
	updatedPropertyValue := "New Value"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	property := NewDocumentProperty()
	property.Value = updatedPropertyValue
	result, _, e := c.SlidesApi.SetDocumentProperty(fileName, propertyName, property, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected false.")
		return
	}

	_, _, e = c.SlidesApi.DeleteDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, response, e := c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e == nil {
		t.Errorf("The property must have been deleted")
		return
	}
	if response.StatusCode != 404 {
		t.Errorf("Wrong status code. Expected %v but was %v.", 404, response.StatusCode)
		return
	}
}

/*
   Test for properties bulk update
*/
func TestPropertyBulkUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	propertyName := "Author"
	customPropertyName := "CustomProperty2"
	updatedPropertyValue := "New Value"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	count := len(result.GetList())

	property1 := NewDocumentProperty()
	property1.Name = propertyName
	property1.Value = updatedPropertyValue
	property2 := NewDocumentProperty()
	property2.Name = customPropertyName
	property2.Value = updatedPropertyValue
	properties := NewDocumentProperties()
	properties.List = []IDocumentProperty{property1, property2}
	result, _, e = c.SlidesApi.SetDocumentProperties(fileName, properties, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetList()) != count+1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count+1, len(result.GetList()))
		return
	}

	result, _, e = c.SlidesApi.DeleteDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetList()) != count-1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count-1, len(result.GetList()))
		return
	}
}

/*
   Test for slide properties
*/
func TestPropertySlideProperties(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	GetResult, _, e := c.SlidesApi.GetSlideProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideProperties()
	dto.FirstSlideNumber = GetResult.GetFirstSlideNumber() + 2
	putResult, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.GetOrientation() != GetResult.GetOrientation() {
		t.Errorf("Wrong orientation. Expected %v but was %v.", GetResult.GetOrientation(), putResult.GetOrientation())
		return
	}
	if putResult.GetFirstSlideNumber() == GetResult.GetFirstSlideNumber() {
		t.Errorf("Wrong FirstSlideNumber. Expected not %v but was %v.", GetResult.GetFirstSlideNumber(), putResult.GetFirstSlideNumber())
		return
	}
}

/*
   Test for preset slide size
*/
func TestPropertySlideSizePreset(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideProperties()
	dto.SizeType = "B4IsoPaper"
	result, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetSizeType() != "B4IsoPaper" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "B4IsoPaper", result.GetSizeType())
		return
	}
	if result.GetWidth() != 852 {
		t.Errorf("Wrong width. Expected %v but was %v.", 852, result.GetWidth())
		return
	}
	if result.GetHeight() != 639 {
		t.Errorf("Wrong height. Expected %v but was %v.", 639, result.GetHeight())
		return
	}
}

/*
   Test for custom slide size
*/
func TestPropertySlideSizeCustom(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var width int32 = 800
	var height int32 = 500
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideProperties()
	dto.Width = width
	dto.Height = height
	result, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetSizeType() != "Custom" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "Custom", result.GetSizeType())
		return
	}
	if result.GetWidth() != width {
		t.Errorf("Wrong width. Expected %v but was %v.", width, result.GetWidth())
		return
	}
	if result.GetHeight() != height {
		t.Errorf("Wrong height. Expected %v but was %v.", height, result.GetHeight())
		return
	}
}

/*
   Test for protection properties
*/
func TestPropertyProtection(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	GetResult, _, e := c.SlidesApi.GetProtectionProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewProtectionProperties()
	dto.EncryptDocumentProperties = GetResult.GetEncryptDocumentProperties()
	dto.ReadOnlyRecommended = !GetResult.GetReadOnlyRecommended()
	putResult, _, e := c.SlidesApi.SetProtection(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.GetEncryptDocumentProperties() != GetResult.GetEncryptDocumentProperties() {
		t.Errorf("Wrong EncryptDocumentProperties. Expected %v but was %v.", GetResult.GetEncryptDocumentProperties(), putResult.GetEncryptDocumentProperties())
		return
	}
	if putResult.GetReadOnlyRecommended() == GetResult.GetReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended. Expected not %v but was %v.", GetResult.GetReadOnlyRecommended(), putResult.GetReadOnlyRecommended())
		return
	}
}

/*
   Test for delete protection
*/
func TestPropertyProtectionDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteProtection(fileName, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetIsEncrypted() {
		t.Errorf("Wrong IsEncrypted value. Expected false.")
		return
	}
	if result.GetReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended value. Expected false.")
		return
	}
	if result.GetReadPassword() != "" {
		t.Errorf("Wrong ReadPassword value. Expected empty string.")
		return
	}
}

/*
   Test for online protection properties
*/
func TestPropertyProtectionOnline(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	dto := NewProtectionProperties()
	dto.ReadPassword = "newPassword"
	result, _, e := GetTestApiClient().SlidesApi.SetProtectionOnline(source, dto, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultStat.Size() == int64(len(source)) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), resultStat.Size())
		return
	}
}

/*
   Test for online delete protection
*/
func TestPropertyProtectionUnprotectOnline(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e := GetTestApiClient().SlidesApi.DeleteProtectionOnline(source, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultStat.Size() == int64(len(source)) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), resultStat.Size())
		return
	}
}

/*
   Test for view properties
*/
func TestViewPropertiesGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetViewProperties(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetShowComments() != "True" {
		t.Errorf("Expected %v, but was %v", "True", response.GetShowComments())
		return
	}
}

/*
   Test for
*/
func TestViewPropertiesSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewViewProperties()
	dto.SetShowComments("False")
	slideViewProperties := NewCommonSlideViewProperties()
	slideViewProperties.SetScale(50)
	dto.SetSlideViewProperties(slideViewProperties)

	response, _, e := c.SlidesApi.SetViewProperties(fileName, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetShowComments() != "False" {
		t.Errorf("Expected %v, but was %v", "False", response.GetShowComments())
		return
	}

	if response.GetSlideViewProperties().GetScale() != 50 {
		t.Errorf("Expected %v, but was %v", 50, response.GetSlideViewProperties().GetScale())
		return
	}
}

/*
   Test for protection check
*/
func TestProtectionCheck(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	protectionProperties, _, e := c.SlidesApi.GetProtectionProperties(fileName, "", folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if protectionProperties.GetIsEncrypted() != true {
		t.Errorf("Expected %v, but was %v", true, protectionProperties.GetIsEncrypted())
		return
	}

	if protectionProperties.GetReadPassword() != "" {
		t.Errorf("Expected %v, but was %v", "", protectionProperties.GetReadPassword())
		return
	}

	protectionProperties, _, e = c.SlidesApi.GetProtectionProperties(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if protectionProperties.GetIsEncrypted() != true {
		t.Errorf("Expected %v, but was %v", true, protectionProperties.GetIsEncrypted())
		return
	}

	if protectionProperties.GetReadPassword() == "" {
		t.Errorf("Expected %v, but was %v", "not null", protectionProperties.GetReadPassword())
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestGoodAuth(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestBadAuth(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.AppSid = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, r, e := testApiClient.SlidesApi.GetApiInfo()
	if e == nil {
		t.Errorf("Must have failed.")
		return
	}
	if r == nil {
		t.Errorf("Null response not expected.")
		return
	}
	statusCode := r.StatusCode
	if statusCode != 401 {
		t.Errorf("Unexpected error code: %v.", statusCode)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestGoodAuthToken(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.AppSid = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestBadAuthToken(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.OAuthToken = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for Get image
*/
func TestImageGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
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
func TestImageDownloadAllStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
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
func TestImageDownloadAllRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	password := "password"
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

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
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
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
	var slideIndex int32 = 1
	password := "password"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

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
   Test for merge from storage
*/
func TestMergeStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	fileName2 := "test-unprotected.pptx"
	fileNamePdf := "test.pdf"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileNamePdf, folderName+"/"+fileNamePdf, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := NewPresentationsMergeRequest()
	request.PresentationPaths = []string{folderName + "/" + fileName2, folderName + "/" + fileNamePdf}
	_, _, e = c.SlidesApi.Merge(fileName, request, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for ordered merge from storage
*/
func TestMergeOrderedStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	fileName2 := "test-unprotected.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := NewOrderedMergeRequest()
	presentation := NewPresentationToMerge()
	presentation.Path = folderName + "/" + fileName2
	presentation.Slides = []int32{2, 1}
	request.Presentations = []IPresentationToMerge{presentation}
	_, _, e = c.SlidesApi.OrderedMerge(fileName, request, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge from request
*/
func TestMergeRequest(t *testing.T) {
	source1, e := ioutil.ReadFile("TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = GetTestApiClient().SlidesApi.MergeOnline([][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge and save from request
*/
func TestMergeAndSaveRequest(t *testing.T) {
	source1, e := ioutil.ReadFile("TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	outPath := "TestData/out.pptx"
	c := GetTestApiClient()

	_, e = c.SlidesApi.MergeAndSaveOnline(outPath, [][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for ordered merge from request
*/
func TestMergeOrderedRequest(t *testing.T) {
	source1, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := NewOrderedMergeRequest()
	presentation1 := NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = "password"
	presentation2 := NewPresentationToMerge()
	presentation2.Path = "file2"
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []IPresentationToMerge{presentation1, presentation2}
	_, _, e = GetTestApiClient().SlidesApi.MergeOnline([][]byte{source1, source2}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for combined merge from request
*/
func TestMergeCombinedRequest(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName2 := "test-unprotected.pptx"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := NewOrderedMergeRequest()
	presentation1 := NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = "password"
	presentation2 := NewPresentationToMerge()
	presentation2.Source = "Storage"
	presentation2.Path = folderName + "/" + fileName2
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []IPresentationToMerge{presentation1, presentation2}
	_, _, e = GetTestApiClient().SlidesApi.MergeOnline([][]byte{source}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for
*/
func TestMergeOrderedUrl(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := NewOrderedMergeRequest()
	presentation1 := NewPresentationToMerge()
	presentation1.Source = "Storage"
	presentation1.Path = folderName + "/" + fileName
	presentation1.Slides = []int32{1, 2}
	presentation1.Password = password
	presentation2 := NewPresentationToMerge()
	presentation2.Source = "Url"
	presentation2.Path = "https://drive.google.com/uc?export=download&id=1ycMzd7e--Ro9H8eH2GL5fPP7-2HjX4My"
	presentation2.Slides = []int32{1}
	request.Presentations = []IPresentationToMerge{presentation1, presentation2}

	_, _, e = c.SlidesApi.MergeOnline(nil, request, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for split from storage
*/
func TestSplitStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result1, _, e := c.SlidesApi.Split(fileName, nil, "", nil, nil, nil, nil, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.Split(fileName, nil, "", nil, nil, &from, &to, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result2.GetSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.GetSlides()) <= len(result2.GetSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.GetSlides()), len(result1.GetSlides()))
		return
	}
	url := result1.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for split from request
*/
func TestSplitRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	password := "password"
	c := GetTestApiClient()

	result1, _, e := c.SlidesApi.SplitOnline(source, "png", nil, nil, nil, nil, password, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	zip1, e := zip.OpenReader(result1.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer zip1.Close()
	count1 := 0
	for _, _ = range zip1.File {
		count1++
	}

	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.SplitOnline(source, "png", nil, nil, &from, &to, password, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	zip2, e := zip.OpenReader(result2.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer zip2.Close()
	count2 := 0
	for _, _ = range zip2.File {
		count2++
	}
	if count2 != 2 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", 2, count2)
		return
	}
	if count1 <= count2 {
		t.Errorf("Wrong slide count. Expected greater than %v but was %v.", count2, count1)
		return
	}
}

/*
   Test for split from request to storage
*/
func TestSplitRequestToStorage(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	password := "password"
	c := GetTestApiClient()

	result1, _, e := c.SlidesApi.SplitAndSaveOnline(source, "png", "", nil, nil, nil, nil, password, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.SplitAndSaveOnline(source, "png", "", nil, nil, &from, &to, password, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result2.GetSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.GetSlides()) <= len(result2.GetSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.GetSlides()), len(result1.GetSlides()))
		return
	}
	url := result1.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for split with options
*/
func TestSplitWithOptions(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := NewPdfExportOptions()
	options.JpegQuality = 50
	result, _, e := c.SlidesApi.Split(fileName, options, "", nil, nil, nil, nil, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	url := result.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for Get text
*/
func TestTextGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var withEmpty bool = true
	result, _, e := c.SlidesApi.GetPresentationTextItems(fileName, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultWithEmpty, _, e := c.SlidesApi.GetPresentationTextItems(fileName, &withEmpty, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResult, _, e := c.SlidesApi.GetSlideTextItems(fileName, slideIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResultWithEmpty, _, e := c.SlidesApi.GetSlideTextItems(fileName, slideIndex, &withEmpty, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetItems()) >= len(resultWithEmpty.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(result.GetItems()), len(resultWithEmpty.GetItems()))
		return
	}
	if len(slideResult.GetItems()) >= len(result.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.GetItems()), len(result.GetItems()))
		return
	}
	if len(slideResult.GetItems()) >= len(slideResultWithEmpty.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.GetItems()), len(slideResultWithEmpty.GetItems()))
		return
	}
}

/*
   Test for replace text on storage
*/
func TestTextReplaceStorage(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1
	oldValue := "text"
	newValue := "new_text"
	c := GetTestApiClient()
	var ignoreCase bool = true

	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultWithEmpty, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResult, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResultWithEmpty, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if result.GetMatches() >= resultWithEmpty.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", result.GetMatches(), resultWithEmpty.GetMatches())
		return
	}
	if slideResult.GetMatches() >= result.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.GetMatches(), result.GetMatches())
		return
	}
	if slideResult.GetMatches() >= slideResultWithEmpty.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.GetMatches(), slideResultWithEmpty.GetMatches())
		return
	}
}

/*
   Test for replace text on request
*/
func TestTextReplaceRequest(t *testing.T) {
	password := "password"
	var slideIndex int32 = 1
	oldValue := "text"
	newValue := "new_text"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

	var withEmpty bool = true
	_, _, e = c.SlidesApi.ReplacePresentationTextOnline(source, oldValue, newValue, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplacePresentationTextOnline(source, oldValue, newValue, &withEmpty, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplaceSlideTextOnline(source, slideIndex, oldValue, newValue, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplaceSlideTextOnline(source, slideIndex, oldValue, newValue, &withEmpty, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for highlight shape text
*/
func TestShapeTextHighlight(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"
	ignoreCase := true
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.HighlightShapeText(fileName, slideIndex, shapeIndex, textToHighlight, highlightColor, nil, &ignoreCase, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}

/*
   Test for highlight shape text with regex
*/
func TestShapeTextHighlightRegex(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	highlightRegex := "h.ghl[abci]ght"
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"
	ignoreCase := false
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.HighlightShapeRegex(fileName, slideIndex, shapeIndex, highlightRegex, highlightColor, nil, &ignoreCase, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}

/*
   Test for text watermark on storage
*/
// func TestWatermarkTextStorage(t *testing.T) {
// 	folderName := "TempSlidesSDK"
// 	fileName := "test.pptx"
// 	password := "password"
// 	var slideIndex int32 = 1
// 	watermarkText := "watermarkText"
// 	c := GetTestApiClient()
// 	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	shapeCount := len(GetResult.GetShapesLinks()) + 1

// 	_, e = c.SlidesApi.CreateWatermark(fileName, nil, nil, watermarkText, "", "", password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// 	GetShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if GetShapeResult.GetName() != "watermark" {
// 		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", GetShapeResult.GetName())
// 		return
// 	}
// 	if GetShapeResult.(IShape).GetText() != watermarkText {
// 		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, GetShapeResult.(IShape).GetText())
// 		return
// 	}

// 	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount-1 != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// }

/*
   Test for text DTO watermark on storage
*/
// func TestWatermarkTextDTOStorage(t *testing.T) {
// 	folderName := "TempSlidesSDK"
// 	fileName := "test.pptx"
// 	password := "password"
// 	var slideIndex int32 = 1
// 	watermarkText := "watermarkText"
// 	c := GetTestApiClient()
// 	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	shapeCount := len(GetResult.GetShapesLinks()) + 1

// 	watermark := NewShape()
// 	watermark.Text = watermarkText
// 	_, e = c.SlidesApi.CreateWatermark(fileName, watermark, nil, watermarkText, "", "", password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// 	GetShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if GetShapeResult.GetName() != "watermark" {
// 		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", GetShapeResult.GetName())
// 		return
// 	}
// 	if GetShapeResult.(IShape).GetText() != watermarkText {
// 		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, GetShapeResult.(IShape).GetText())
// 		return
// 	}

// 	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount-1 != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// }

/*
   Test for image watermark on storage
*/
// func TestWatermarkImageStorage(t *testing.T) {
// 	folderName := "TempSlidesSDK"
// 	fileName := "test.pptx"
// 	password := "password"
// 	var slideIndex int32 = 1
// 	c := GetTestApiClient()
// 	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	shapeCount := len(GetResult.GetShapesLinks()) + 1

// 	source, e := ioutil.ReadFile("TestData/watermark.png")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	_, e = c.SlidesApi.CreateImageWatermark(fileName, source, nil, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// 	GetShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if GetShapeResult.GetName() != "watermark" {
// 		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", GetShapeResult.GetName())
// 		return
// 	}

// 	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount-1 != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// }

/*
   Test for image DTO watermark on storage
*/
// func TestWatermarkImageDTOStorage(t *testing.T) {
//         folderName := "TempSlidesSDK"
//         fileName := "test.pptx"
//         password := "password"
//         watermarkName := "myWatermark"
//         var slideIndex int32 = 1
//         c := GetTestApiClient()
//         _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         GetResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         shapeCount := len(GetResult.GetShapesLinks()) + 1

// 	source, e := ioutil.ReadFile("TestData/watermark.png")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         watermark := NewPictureFrame()
//         fillFormat := NewPictureFill()
//         fillFormat.Base64Data = base64.StdEncoding.EncodeToString(source)
//         watermark.FillFormat = fillFormat
//         watermark.Name = watermarkName
//         _, e = c.SlidesApi.CreateImageWatermark(fileName, nil, watermark, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(GetResult.GetShapesLinks()))
// 		return
// 	}
//         GetShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if GetShapeResult.GetName() != watermarkName {
// 		t.Errorf("Wrong shape name. Expected %v but was %v.", watermarkName, GetShapeResult.GetName())
// 		return
// 	}

//         _, e = c.SlidesApi.DeleteWatermark(fileName, watermarkName, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         GetResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if shapeCount - 1 != len(GetResult.GetShapesLinks()) {
// 		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(GetResult.GetShapesLinks()))
// 		return
// 	}
// }

/*
   Test for text watermark from request
*/
// func TestWatermarkTextRequest(t *testing.T) {
//         password := "password"
// 	source, e := ioutil.ReadFile("TestData/test.pptx")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         c := GetTestApiClient()

//         postResult, _, e := c.SlidesApi.CreateWatermarkOnline(source, nil, nil, "watermark", "", "", password)
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         postStat, e := os.Stat(postResult.Name())
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if len(source) == int(postStat.Size()) {
// 		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
// 		return
// 	}

//         deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
//         deleteStat, e := os.Stat(deleteResult.Name())
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// 	if deleteStat.Size() >= postStat.Size() {
// 		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
// 		return
// 	}
// }

/*
   Test for text DTO watermark from request
*/
func TestWatermarkTextDTORequest(t *testing.T) {
	password := "password"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

	watermark := NewShape()
	watermark.Text = "watermarkText"
	postResult, _, e := c.SlidesApi.CreateWatermarkOnline(source, watermark, nil, "", "", "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for image watermark from request
*/
func TestWatermarkImageRequest(t *testing.T) {
	password := "password"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	watermark, e := ioutil.ReadFile("TestData/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

	postResult, _, e := c.SlidesApi.CreateImageWatermarkOnline(source, watermark, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for image DTO watermark from request
*/
func TestWatermarkImageDTORequest(t *testing.T) {
	password := "password"
	watermarkName := "myWatermark"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	watermark, e := ioutil.ReadFile("TestData/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c := GetTestApiClient()

	dto := NewPictureFrame()
	fillFormat := NewPictureFill()
	fillFormat.Base64Data = base64.StdEncoding.EncodeToString(watermark)
	dto.FillFormat = fillFormat
	dto.Name = watermarkName
	postResult, _, e := c.SlidesApi.CreateImageWatermarkOnline(source, nil, dto, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, watermarkName, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for Get animation
*/
func TestAnimationGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	var shapeIndex int32 = 3
	var paragraphIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.GetAnimation(fileName, slideIndex, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}

	animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}

	animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &shapeIndex, &paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
}

/*
   Test for set animation
*/
func TestAnimationSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideAnimation()
	effect1 := NewEffect()
	effect1.Type_ = "Blink"
	effect1.ShapeIndex = 2
	effect1.ParagraphIndex = 1

	effect2 := NewEffect()
	effect2.Type_ = "Box"
	effect2.Subtype = "In"
	effect2.PresetClassType = "Entrance"
	effect2.ShapeIndex = 4
	dto.MainSequence = []IEffect{effect1, effect2}
	dto.InteractiveSequences = []IInteractiveSequence{}
	animation, _, e := c.SlidesApi.SetAnimation(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation effect
*/
func TestAnimationCreateEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = 3
	animation, _, e := c.SlidesApi.CreateAnimationEffect(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 2 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 2, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation intercative sequence
*/
func TestAnimationCreateInteractiveSequence(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewInteractiveSequence()
	dto.TriggerShapeIndex = 2
	effect := NewEffect()
	effect.Type_ = "Blast"
	effect.ShapeIndex = 3
	dto.Effects = []IEffect{effect}
	animation, _, e := c.SlidesApi.CreateAnimationInteractiveSequence(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 2 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 2, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation interactive sequence effect
*/
func TestAnimationCreateInteractiveSequenceEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = 3
	animation, _, e := c.SlidesApi.CreateAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for update animation effect
*/
func TestAnimationUpdateEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = 3
	animation, _, e := c.SlidesApi.UpdateAnimationEffect(fileName, slideIndex, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for update animation interactive sequence effect
*/
func TestAnimationUpdateInteractiveSequenceEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = 3
	animation, _, e := c.SlidesApi.UpdateAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation
*/
func TestAnimationDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimation(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation main sequence
*/
func TestAnimationDeleteMainSequence(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimationMainSequence(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation main sequence effect
*/
func TestAnimationDeleteMainSequenceEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimationEffect(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequences
*/
func TestAnimationDeleteInteractiveSequences(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequences(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequence
*/
func TestAnimationDeleteInteractiveSequence(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequence(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequence effect
*/
func TestAnimationDeleteInteractiveSequenceEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	var slideIndex int32 = 1
	password := "password"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for pipeline
*/
func TestPipelineTemplate(t *testing.T) {
	file1 := NewRequestInputFile()
	file1.Index = 0

	file2 := NewRequestInputFile()
	file2.Index = 1

	input := NewInput()
	input.TemplateData = file1
	input.Template = file2

	output := NewOutputFile()

	task := NewSave()
	task.Format = "pptx"
	task.Output = output

	pipeline := NewPipeline()
	pipeline.Input = input
	pipeline.Tasks = []ITask{task}

	data1, e := ioutil.ReadFile("TestData/TemplatingCVDataWithBase64.xml")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data2, e := ioutil.ReadFile("TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	files := [][]byte{data1, data2}

	_, _, e = GetTestApiClient().SlidesApi.Pipeline(pipeline, files)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for SlidesApi.PostSlideSaveAs with timeout method
*/
func TestTimeout(t *testing.T) {
	return
	e := initializeTest("PostSlideSaveAs", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.Timeout = 1
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.DownloadSlide("test.pptx", 1, "svg", nil, nil, nil, "password", "TempSlidesSDK", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
 * Test for Shape
 */
func TestShape(t *testing.T) {
	e := initializeTest("GetSlideShape", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e := GetTestApiClient().SlidesApi.GetShape("test.pptx", 1, 1, "password", "TempSlidesSDK", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IShape).GetText() != "1" {
		t.Errorf("Error: Text expected to equal 1.")
		return
	}
}

/*
   Test for Chart creation
*/
func TestChart(t *testing.T) {
	chart := NewChart()
	if chart.GetType() != "Chart" {
		t.Errorf("Unexpected chart type: %v.", chart.GetType())
		return
	}
}

/*
 * Test for nullable properties
 */
func TestNullableProperties(t *testing.T) {
	var folderName = "TempSlidesSDK"
	var fileName = "test.pptx"
	var password = "password"
	var min1 = 44.3
	var min2 = 12.0
	var max1 = 104.3
	var max2 = 87.0

	e := initializeTest("NoFunction", "No method", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := GetTestApiClient()
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var dto1 Chart
	dto1.SetType("Chart")
	dto1.SetChartType("Line")
	dto1.SetWidth(400)
	dto1.SetHeight(300)
	var title ChartTitle
	title.SetHasTitle(true)
	title.SetText("MyTitle")
	dto1.SetTitle(&title)
	var series OneValueSeries
	series.SetType("ClusteredColumn")
	series.SetDataPointType("OneValue")
	series.SetName("Series1")
	var point1 OneValueChartDataPoint
	point1.SetValue(40)
	var point2 OneValueChartDataPoint
	point2.SetValue(50)
	points := make([]IOneValueChartDataPoint, 2)
	points[0] = &point1
	points[1] = &point2
	series.SetDataPoints(points)
	serieses := make([]ISeries, 1)
	serieses[0] = &series
	dto1.SetSeries(serieses)
	var axes Axes
	var axis1 Axis
	axis1.SetIsAutomaticMinValue(false)
	axis1.SetMinValue(min1)
	axis1.SetIsAutomaticMaxValue(false)
	axis1.SetMaxValue(max1)
	axes.SetHorizontalAxis(&axis1)
	dto1.SetAxes(&axes)
	dto1.SetX(12)
	dto1.SetY(14)
	_, _, e = c.SlidesApi.CreateShape(fileName, 1, &dto1, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e := c.SlidesApi.GetShape(fileName, 1, 5, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min1 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

	var dto2 Chart
	dto2.SetType("Chart")
	dto2.SetChartType("Line")
	var axis2 Axis
	axis2.SetMinValue(min2)
	axes.SetHorizontalAxis(&axis2)
	dto2.SetAxes(&axes)
	_, _, e = c.SlidesApi.UpdateShape(fileName, 1, 5, &dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, 1, 5, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

	var axis3 Axis
	axis3.SetMaxValue(max2)
	axes.SetHorizontalAxis(&axis3)
	_, _, e = c.SlidesApi.UpdateShape(fileName, 1, 5, &dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, 1, 5, password, folderName, "")
	if e != nil {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max2 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}
}

/*
   Test for create comment
*/
func TestCommentCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideComment()
	dto.Text = commentText
	dto.Author = author

	childCommentDto := NewSlideComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	dto.ChildComments = []ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}

	if response.GetList()[0].GetText() != commentText {
		t.Errorf("Expected %v, but was %v", commentText, response.GetList()[0].GetText())
		return
	}

	if response.GetList()[0].GetAuthor() != author {
		t.Errorf("Expected %v, but was %v", author, response.GetList()[0].GetAuthor())
		return
	}

	childComment := response.GetList()[0].GetChildComments()[0]
	if childComment.GetText() != childCommentText {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetText())
		return
	}

	if childComment.GetAuthor() != author {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetAuthor())
		return
	}
}

/*
   Test for create comment from request
*/
func TestCommentCreateOnline(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideComment()
	dto.Text = commentText
	dto.Author = author

	childCommentDto := NewSlideComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	dto.ChildComments = []ISlideCommentBase{childCommentDto}

	source, e := ioutil.ReadFile("TestData/test.pptx")
	_, _, e = c.SlidesApi.CreateCommentOnline(source, slideIndex, dto, nil, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for get slide comments
*/
func TestSlideCommentsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlideComments(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetList()))
		return
	}

	if len(response.GetList()[0].GetChildComments()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()[0].GetChildComments()))
		return
	}
}

/*
   Test for delete comments
*/
func TestCommentsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.DeleteComments(fileName, "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlideComments(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetList()))
		return
	}
}

/*
   Test for delete comments from request
*/
func TestCommentsDeleteOnline(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile("TestData/test.pptx")

	_, _, e = c.SlidesApi.DeleteCommentsOnline(document, "", password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for delete slide comments
*/
func TestSlideCommentsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSlideComments(fileName, slideIndex, "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetList()))
		return
	}
}

/*
   Test for delete slide comments from request
*/
func TestSlideCommentsDeleteOnline(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile("TestData/test.pptx")
	_, _, e = c.SlidesApi.DeleteSlideCommentsOnline(document, slideIndex, "", password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create modern comment
*/
func TestModernCommentCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var textSelectionStartIndex int32 = 1
	var textSelectionLength int32 = 5
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	childCommentDto := NewSlideModernComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	childCommentDto.Status = "Resolved"

	dto := NewSlideModernComment()
	dto.Text = commentText
	dto.Author = author
	dto.Status = "Active"
	dto.TextSelectionStart = textSelectionStartIndex
	dto.TextSelectionLength = textSelectionLength
	dto.ChildComments = []ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}

	childComment := response.GetList()[0].GetChildComments()[0]
	if childComment.GetText() != childCommentText {
		t.Errorf("Expected %v, but was %v", childCommentText, childComment.GetText())
		return
	}
}

/*
   Test for create shape modern comment
*/
func TestShapeModernCommentCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 3
	var shapeIndex int32 = 1
	var textSelectionStartIndex int32 = 1
	var textSelectionLength int32 = 5
	commentText := "Comment text"
	author := "Test author"
	childCommentText := "Child comment text"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	childCommentDto := NewSlideModernComment()
	childCommentDto.Text = childCommentText
	childCommentDto.Author = author
	childCommentDto.Status = "Resolved"

	dto := NewSlideModernComment()
	dto.Text = commentText
	dto.Author = author
	dto.Status = "Active"
	dto.TextSelectionStart = textSelectionStartIndex
	dto.TextSelectionLength = textSelectionLength
	dto.ChildComments = []ISlideCommentBase{childCommentDto}

	response, _, e := c.SlidesApi.CreateComment(fileName, slideIndex, dto, &shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetList()))
		return
	}
}

/*
   Test for get paragraph
*/
func TestParagraphGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.GetPortionList())
		return
	}
}

/*
   Test for get paragraphs
*/
func TestParagraphsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetParagraphs(fileName, slideIndex, shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.GetParagraphLinks())
		return
	}
}

/*
   Test for get sub-shape paragraph
*/
func TestSubshapeParagraphGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapeParagraph(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.GetPortionList())
		return
	}
}

/*
   Test for get sub-shape paragraphs
*/
func TestSubshapeParagraphsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapeParagraphs(fileName, slideIndex, "3/shapes", shapeIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.GetParagraphLinks())
		return
	}
}

/*
   Test for create paragraph
*/
func TestParagraphCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewParagraph()
	dto.MarginLeft = 2
	dto.MarginRight = 2
	dto.Alignment = "Center"

	response, _, e := c.SlidesApi.CreateParagraph(fileName, slideIndex, shapeIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetMarginLeft() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginLeft())
		return
	}

	if response.GetMarginRight() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginRight())
		return
	}

	if response.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Expected %v, but was %v", dto.GetAlignment(), response.GetAlignment())
		return
	}
}

/*
   Test for create paragraph with portions
*/
func TestParagraphCreateWithPortions(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion1 := NewPortion()
	portion1.Text = "Portion1"
	portion2 := NewPortion()
	portion2.Text = "Portion2"
	portion2.FontBold = "True"
	dto := NewParagraph()
	dto.PortionList = []IPortion{portion1, portion2}

	response, _, e := c.SlidesApi.CreateParagraph(fileName, slideIndex, shapeIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetPortionList()))
		return
	}
}

/*
   Test for create sub-shape paragraph
*/
func TestSubshapeParagraphCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewParagraph()
	dto.MarginLeft = 2
	dto.MarginRight = 2
	dto.Alignment = "Center"

	response, _, e := c.SlidesApi.CreateSubshapeParagraph(fileName, slideIndex, "3/shapes", shapeIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetMarginLeft() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginLeft())
		return
	}

	if response.GetMarginRight() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginRight())
		return
	}

	if response.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Expected %v, but was %v", dto.GetAlignment(), response.GetAlignment())
		return
	}
}

/*
   Test for update paragraph
*/
func TestParagraphUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewParagraph()
	dto.MarginLeft = 2
	dto.MarginRight = 2
	dto.Alignment = "Center"

	response, _, e := c.SlidesApi.UpdateParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetMarginLeft() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginLeft())
		return
	}

	if response.GetMarginRight() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginRight())
		return
	}

	if response.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Expected %v, but was %v", dto.GetAlignment(), response.GetAlignment())
		return
	}
}

/*
   Test for update sub-shape paragraph
*/
func TestSubshapeParagraphUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewParagraph()
	dto.MarginLeft = 2
	dto.MarginRight = 2
	dto.Alignment = "Center"

	response, _, e := c.SlidesApi.UpdateSubshapeParagraph(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetMarginLeft() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginLeft())
		return
	}

	if response.GetMarginRight() != dto.GetMarginLeft() {
		t.Errorf("Expected %v, but was %v", dto.GetMarginLeft(), response.GetMarginRight())
		return
	}

	if response.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Expected %v, but was %v", dto.GetAlignment(), response.GetAlignment())
		return
	}
}

/*
   Test for delete paragraphs
*/
func TestParagraphsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteParagraphs(fileName, slideIndex, shapeIndex, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, response.GetParagraphLinks())
		return
	}
}

/*
   Test for delete paragraphs by indexes
*/
func TestParagraphsDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	indexes := []int32{2}
	response, _, e := c.SlidesApi.DeleteParagraphs(fileName, slideIndex, shapeIndex, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response.GetParagraphLinks())
		return
	}
}

/*
   Test for delete sub-shape paragraphs
*/
func TestSubshapeParagraphsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapeParagraphs(fileName, slideIndex, "3/shapes", shapeIndex, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, response.GetParagraphLinks())
		return
	}
}

/*
   Test for delete sub-shape paragraphs by indexes
*/
func TestSubshapeParagraphsDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	indexes := []int32{1}
	response, _, e := c.SlidesApi.DeleteSubshapeParagraphs(fileName, slideIndex, "3/shapes", shapeIndex, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response.GetParagraphLinks())
		return
	}
}

/*
   Test for delete paragraph
*/
func TestParagraphDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response.GetParagraphLinks())
		return
	}
}

/*
   Test for sub-shape delete paragraph
*/
func TestSubshapeParagraphDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapeParagraph(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetParagraphLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response.GetParagraphLinks())
		return
	}
}

/*
   Test for get paragraph rect
*/
func TestParagraphRectGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetParagraphRectangle(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetX() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetX())
		return
	}

	if response.GetY() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetY())
		return
	}

	if response.GetWidth() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetWidth())
		return
	}

	if response.GetHeight() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetHeight())
		return
	}
}

/*
   Test for paragraph default portionformat
*/
func TestParagraphDefaultPortionFormat(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion1 := NewPortion()
	portion1.Text = "Portion1"
	portion2 := NewPortion()
	portion2.Text = "Portion2"

	portionFormat := NewPortionFormat()
	portionFormat.FontHeight = 20
	portionFormat.LatinFont = "Arial"

	dto := NewParagraph()
	dto.DefaultPortionFormat = portionFormat
	dto.PortionList = []IPortion{portion1, portion2}

	response, _, e := c.SlidesApi.CreateParagraph(fileName, slideIndex, shapeIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetPortionList()))
		return
	}

	if response.GetDefaultPortionFormat().GetLatinFont() != dto.GetDefaultPortionFormat().GetLatinFont() {
		t.Errorf("Expected %v, but was %v", response.GetDefaultPortionFormat().GetLatinFont(), response.GetDefaultPortionFormat().GetLatinFont())
		return
	}

	if response.GetDefaultPortionFormat().GetFontHeight() != dto.GetDefaultPortionFormat().GetFontHeight() {
		t.Errorf("Expected %v, but was %v", response.GetDefaultPortionFormat().GetFontHeight(), response.GetDefaultPortionFormat().GetFontHeight())
		return
	}
}

/*
   Test for get paragraph effective
*/
func TestParagraphEffectiveGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetParagraphEffective(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetDefaultTabSize() != 72 {
		t.Errorf("Expected %v, but was %v", false, true)
		return
	}
}

/*
   Test for get sub-shapeparagraph effective
*/
func TestSubshapeParagraphEffectiveGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapeParagraphEffective(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetDefaultTabSize() != 72 {
		t.Errorf("Expected %v, but was %v", false, true)
		return
	}
}

/*
   Test for get portions
*/
func TestPortionsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetPortions(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetItems()))
		return
	}
}

/*
   Test for get sub-shape portions
*/
func TestSubshapePortionsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapePortions(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetItems()))
		return
	}
}

/*
   Test for get portion
*/
func TestPortionGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetPortion(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if !strings.Contains(response.GetText(), portionText) {
		t.Errorf("Expected %v, but was %v", "contains \""+portionText+"\"", "\""+response.GetText()+"\"")
		return
	}
}

/*
   Test for get sub-sshape portion
*/
func TestSubshapePortionGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapePortion(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if !strings.Contains(response.GetText(), portionText) {
		t.Errorf("Expected %v, but was %v", "contains \""+portionText+"\"", "\""+response.GetText()+"\"")
		return
	}
}

/*
   Test for create portion
*/
func TestPortionCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := NewPortion()
	dto.Text = portionText
	dto.SetFontBold("True")
	dto.FontHeight = 20
	dto.LatinFont = fontName
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.CreatePortion(fileName, slideIndex, shapeIndex, paragraphIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), response.GetText())
		return
	}

	if response.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Expected %v, but was %v", dto.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetFontHeight() != dto.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetLatinFont() != dto.GetLatinFont() {
		t.Errorf("Expected %v, but was %v", dto.GetLatinFont(), response.GetLatinFont())
		return
	}

	if response.GetFillFormat().GetType() != dto.GetFillFormat().GetType() {
		t.Errorf("Expected %v, but was %v", dto.GetFillFormat().GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for create sub-shape portion
*/
func TestSubshapePortionCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := NewPortion()
	dto.Text = portionText
	dto.SetFontBold("True")
	dto.FontHeight = 20
	dto.LatinFont = fontName
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.CreateSubshapePortion(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, dto, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), response.GetText())
		return
	}

	if response.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Expected %v, but was %v", dto.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetFontHeight() != dto.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetLatinFont() != dto.GetLatinFont() {
		t.Errorf("Expected %v, but was %v", dto.GetLatinFont(), response.GetLatinFont())
		return
	}

	if response.GetFillFormat().GetType() != dto.GetFillFormat().GetType() {
		t.Errorf("Expected %v, but was %v", dto.GetFillFormat().GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for update portion
*/
func TestPortionUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := NewPortion()
	dto.Text = portionText
	dto.SetFontBold("True")
	dto.FontHeight = 20
	dto.LatinFont = fontName
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.UpdatePortion(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), response.GetText())
		return
	}

	if response.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Expected %v, but was %v", dto.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetFontHeight() != dto.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetLatinFont() != dto.GetLatinFont() {
		t.Errorf("Expected %v, but was %v", dto.GetLatinFont(), response.GetLatinFont())
		return
	}

	if response.GetFillFormat().GetType() != dto.GetFillFormat().GetType() {
		t.Errorf("Expected %v, but was %v", dto.GetFillFormat().GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for update sub-shape portion
*/
func TestSubshapePortionUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := NewPortion()
	dto.Text = portionText
	dto.SetFontBold("True")
	dto.FontHeight = 20
	dto.LatinFont = fontName
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.UpdateSubshapePortion(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, portionIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), response.GetText())
		return
	}

	if response.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Expected %v, but was %v", dto.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetFontHeight() != dto.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetLatinFont() != dto.GetLatinFont() {
		t.Errorf("Expected %v, but was %v", dto.GetLatinFont(), response.GetLatinFont())
		return
	}

	if response.GetFillFormat().GetType() != dto.GetFillFormat().GetType() {
		t.Errorf("Expected %v, but was %v", dto.GetFillFormat().GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for delete portions
*/
func TestPortionsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeletePortions(fileName, slideIndex, shapeIndex, paragraphIndex, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetItems()))
		return
	}
}

/*
   Test for delete portions indexes
*/
func TestPortionsDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	indexes := []int32{1}
	response, _, e := c.SlidesApi.DeletePortions(fileName, slideIndex, shapeIndex, paragraphIndex, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetItems()))
		return
	}
}

/*
   Test for delete sub-shape portions
*/
func TestSubshapePortionsDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapePortions(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetItems()))
		return
	}
}

/*
   Test for delete sub-shape portions indexes
*/
func TestSubshapePortionsDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	indexes := []int32{1}
	response, _, e := c.SlidesApi.DeleteSubshapePortions(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetItems()))
		return
	}
}

/*
   Test for delete portion
*/
func TestPortionDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeletePortion(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetItems()))
		return
	}
}

/*
   Test for delete sub-shape portion
*/
func TestSubshapePortionDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSubshapePortion(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetItems()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetItems()))
		return
	}
}

/*
   Test for portion rect
*/
func TestPortionRect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetPortionRectangle(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetX() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetX())
		return
	}

	if response.GetY() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetY())
		return
	}

	if response.GetWidth() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetWidth())
		return
	}

	if response.GetHeight() <= 0 {
		t.Errorf("Expected %v, but was %v", "greater then 0", response.GetHeight())
		return
	}
}

/*
   Test for import shapes from svg
*/
func TestShapesFromSvgImport(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 5
	var x int32 = 50
	var y int32 = 50
	var width int32 = 300
	var height int32 = 300

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile("TestData/shapes.svg")
	shapes := []int32{1, 3, 5}
	response, _, e := c.SlidesApi.ImportShapesFromSvg(fileName, slideIndex, document, &x, &y, &width, &height, shapes, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for get portion effective
*/
func TestPortionEffectiveGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetPortionEffective(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFontHeight() != 18 {
		t.Errorf("Expected %v, but was %v", 18, response.GetFontHeight())
		return
	}
}

/*
   Test for get sub-shape portion effective
*/
func TestSubshapePortionEffectiveGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSubshapePortionEffective(fileName, slideIndex, "3/shapes", shapeIndex, paragraphIndex, portionIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFontHeight() != 18 {
		t.Errorf("Expected %v, but was %v", 18, response.GetFontHeight())
		return
	}
}

/*
   Test for get slides
*/
func TestSlidesGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != 8 {
		t.Errorf("Expected %v, but was %v", 8, len(response.GetSlideList()))
		return
	}
}

/*
   Test for get slide
*/
func TestSlideGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.GetSlide(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create slide
*/
func TestSlideCreate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var position int32 = 1
	var slidesCount int = 8

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.CreateSlide(fileName, "layoutSlides/3", &position, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount+1 {
		t.Errorf("Expected %v, but was %v", slidesCount+1, len(response.GetSlideList()))
		return
	}

	response, _, e = c.SlidesApi.CreateSlide(fileName, "layoutSlides/3", nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount+2 {
		t.Errorf("Expected %v, but was %v", slidesCount+2, len(response.GetSlideList()))
		return
	}
}

/*
   Test for copy slide
*/
func TestSlideCopy(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slidesCount int = 8

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.CopySlide(fileName, 3, nil, "", "", "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount+1 {
		t.Errorf("Expected %v, but was %v", slidesCount+1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for copy slide from source
*/
func TestSlideCopyFromSource(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	sourceFileName := "TemplateCV.pptx"
	password := "password"
	var slidesCount int = 8
	var slideIndex int32 = 1
	var position int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile("TempTests/"+sourceFileName, folderName+"/"+sourceFileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.CopySlide(fileName, slideIndex, &position, folderName+"/"+sourceFileName, "", "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount+1 {
		t.Errorf("Expected %v, but was %v", slidesCount+1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for move slide
*/
func TestSlideMove(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slidesCount int = 8
	var slideIndex int32 = 1
	var position int32 = 2

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.MoveSlide(fileName, slideIndex, position, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount {
		t.Errorf("Expected %v, but was %v", slidesCount, len(response.GetSlideList()))
		return
	}
}

/*
   Test for reorder slides
*/
func TestSlidesReorder(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slidesCount int = 8

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	oldPositions := []int32{1, 2, 3, 4, 5, 6}
	newPositions := []int32{6, 5, 4, 3, 2, 1}

	response, _, e := c.SlidesApi.ReorderSlides(fileName, oldPositions, newPositions, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount {
		t.Errorf("Expected %v, but was %v", slidesCount, len(response.GetSlideList()))
		return
	}
}

/*
   Test for update slide
*/
func TestSlideUpdate(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	layoutSlide := NewResourceUri()
	layoutSlide.Href = "layoutSlides/3"
	dto := NewSlide()
	dto.LayoutSlide = layoutSlide

	response, _, e := c.SlidesApi.UpdateSlide(fileName, slideIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if !strings.Contains(response.GetLayoutSlide().GetHref(), "layoutSlides/3") {
		t.Errorf("Expected %v, but was %v", "contains \"layoutSlides/3\"", "\""+response.GetLayoutSlide().GetHref()+"\"")
		return
	}
}

/*
   Test for delete slide
*/
func TestSlidesDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSlides(fileName, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for delete slides by indexes
*/
func TestSlideDeleteIndexes(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slidesCount int = 8

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	indexes := []int32{1, 3, 5}
	response, _, e := c.SlidesApi.DeleteSlides(fileName, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount-len(indexes) {
		t.Errorf("Expected %v, but was %v", slidesCount-len(indexes), len(response.GetSlideList()))
		return
	}
}

/*
   Test for delete slide
*/
func TestSlideDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slidesCount int = 8
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteSlide(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount-1 {
		t.Errorf("Expected %v, but was %v", slidesCount-1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for get background
*/
func TestBackgroundGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 5

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetBackground(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for set background
*/
func TestBackgroundSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	color := "#FFF5FF8A"
	var slideIndex int32 = 5

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewSlideBackground()
	fillFormat := NewSolidFill()
	fillFormat.Color = color
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.SetBackground(fileName, slideIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}

	if response.GetFillFormat().(ISolidFill).GetColor() != color {
		t.Errorf("Expected %v, but was %v", color, response.GetFillFormat().(ISolidFill).GetColor())
		return
	}
}

/*
   Test for set background color
*/
func TestBackgroundColorSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	color := "#FFF5FF8A"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.SetBackgroundColor(fileName, slideIndex, color, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}

	if response.GetFillFormat().(ISolidFill).GetColor() != color {
		t.Errorf("Expected %v, but was %v", color, response.GetFillFormat().(ISolidFill).GetColor())
		return
	}
}

/*
   Test for delete background
*/
func TestBackgroundDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 5

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.DeleteBackground(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "NoFill" {
		t.Errorf("Expected %v, but was %v", "NoFill", response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for text format 3D
*/
func TestTextFormat3D(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	var slideIndex int32 = 1

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := NewShape()
	dto.ShapeType = "Rectangle"
	dto.X = 100
	dto.Y = 100
	dto.Height = 100
	dto.Width = 200
	dto.Text = "Sample text"

	textFormat := NewTextFrameFormat()
	textFormat.Transform = "ArchUpPour"

	threeDFormat := NewThreeDFormat()
	bevelBottom := NewShapeBevel()
	bevelBottom.BevelType = "Circle"
	bevelBottom.Height = 3.5
	bevelBottom.Width = 3.5
	threeDFormat.BevelBottom = bevelBottom
	bevelTop := NewShapeBevel()
	bevelTop.BevelType = "Circle"
	bevelTop.Height = 4
	bevelTop.Width = 4
	threeDFormat.BevelTop = bevelTop
	threeDFormat.ExtrusionColor = "#FF008000"
	threeDFormat.ExtrusionHeight = 6
	threeDFormat.ContourColor = "#FF25353D"
	threeDFormat.ContourWidth = 1.5
	threeDFormat.Depth = 3
	threeDFormat.Material = "Plastic"
	lightRig := NewLightRig()
	lightRig.LightType = "Balanced"
	lightRig.Direction = "Top"
	lightRig.XRotation = 0
	lightRig.YRotation = 0
	lightRig.ZRotation = 40
	threeDFormat.LightRig = lightRig
	camera := NewCamera()
	camera.CameraType = "PerspectiveContrastingRightFacing"
	threeDFormat.Camera = camera
	textFormat.ThreeDFormat = threeDFormat
	dto.TextFrameFormat = textFormat

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Shape", response.GetType())
		return
	}
}

/*
   Test for get fonts
*/
func TestFontsGet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"

	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetFonts(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetList()))
		return
	}
}

/*
   Test for get fonts
*/
func TestFontsGetOnline(t *testing.T) {
	password := "password"

	c := GetTestApiClient()
	document, e := ioutil.ReadFile("TestData/test.pptx")

	response, _, e := c.SlidesApi.GetFontsOnline(document, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetList()))
		return
	}
}

/*
   Test for set embedded font
*/
func TestEmbeddedFontSet(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	fontName := "Calibri"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetList()[2].GetIsEmbedded() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
}

/*
   Test for set embedded font online
*/
func TestEmbeddedFontSetOnline(t *testing.T) {
	password := "password"
	fontName := "Calibri"

	c := GetTestApiClient()
	document, e := ioutil.ReadFile("TestData/test.pptx")

	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for set embedded font
*/
func TestEmbeddedFontDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	password := "password"
	fontName := "Calibri"
	c := GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetList()[2].GetIsEmbedded() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}

	response, _, e = c.SlidesApi.DeleteEmbeddedFont(fileName, fontName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetList()[2].GetIsEmbedded() != false {
		t.Errorf("Expected %v, but was %v", false, response.GetList()[2].GetIsEmbedded())
		return
	}
}

/*
   Test for delete embedded font online
*/
// func TestEmbeddedFontDeleteOnline(t *testing.T) {
// 	password := "password"
// 	fontName := "Calibri"

// 	c := GetTestApiClient()
// 	document, e := ioutil.ReadFile("TestData/test.pptx")
// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}

// 	var onlyUsed bool = false
// 	_, _, e = c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password)

// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}

// 	_, _, e = c.SlidesApi.DeleteEmbeddedFontOnline(response, fontName, password)

// 	if e != nil {
// 		t.Errorf("Error: %v.", e)
// 		return
// 	}
// }
