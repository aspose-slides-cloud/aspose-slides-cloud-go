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
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + fileName, "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + fileName, "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + newFileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        sourcePath := folderName + "/" + fileName
        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName, sourcePath, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        templatePath := folderName + "/" + templateFileName
        _, e = c.SlidesApi.CopyFile("TempTests/" + templateFileName, templatePath, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + fileName, "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResponse1, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        slideCount := len(getResponse1.getSlideList())
        _, _, e = c.SlidesApi.ImportFromHtml(fileName, "<html><body>New Content</body></html>", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResponse2, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        newSlideCount := len(getResponse2.getSlideList())
	if newSlideCount != slideCount + 1 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", slideCount + 1, newSlideCount)
		return
	}
}

/* 
   Test for create presentation from PDF
*/
func TestCreateFromPdf(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.DeleteFile(folderName + "/" + fileName, "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResponse1, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        slideCount := len(getResponse1.getSlideList())
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
        getResponse2, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        newSlideCount := len(getResponse2.getSlideList())
	if newSlideCount != slideCount + 4 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", slideCount + 4, newSlideCount)
		return
	}
}

/* 
   Test for convert presentation from request to request
*/
func TestConvertPostFromRequest(t *testing.T) {
        password := "password"
        c := getTestApiClient()
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

        resultSlides, _, e := c.SlidesApi.Convert(source, "pdf", password, "", "", []int32 { 2, 4 }, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        resultSlidesStat, e := os.Stat(resultSlides.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultSlidesStat.Size() >= resultStat.Size() {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", resultSlidesStat.Size(), resultStat.Size())
		return
	}
}

/* 
   Test for convert presentation from request to storage
*/
func TestConvertPutFromRequest(t *testing.T) {
        outPath := "TestData/test.pdf"
        c := getTestApiClient()
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
	if !resultExists.getExists() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !resultExists.getExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/* 
   Test for convert presentation with options from request
*/
func TestConvertWithOptionsFromRequest(t *testing.T) {
        password := "password"
        c := getTestApiClient()
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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

        _, _, e = getTestApiClient().SlidesApi.DownloadSlideOnline(source, 1, "pdf", nil, nil, "password", "", "", nil)
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
        c := getTestApiClient()
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
	if !resultExists.getExists() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !resultExists.getExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/* 
   Test for convert slide with options from request
*/
func TestConvertSlideWithOptionsFromRequest(t *testing.T) {
        password := "password"
        c := getTestApiClient()
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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

        _, _, e = getTestApiClient().SlidesApi.DownloadShapeOnline(source, 1, 3, "png", nil, nil, "", "password", "", "", nil)
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
        c := getTestApiClient()
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
	if !resultExists.getExists() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
   Test for convert shape from storage to storage
*/
func TestConvertShapePutFromStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        outPath := "TestData/test.png"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !resultExists.getExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/* 
   Test for get notes slide from storage
*/
func TestNotesSlideGetFromStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.GetNotesSlide(fileName, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getText()) == 0 {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.NotesSlideExists(fileName, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.getExists() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
   Test for get notes slide from request
*/
func TestNotesSlideGetFromRequest(t *testing.T) {
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := getTestApiClient().SlidesApi.GetNotesSlideOnline(source, 1, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getText()) == 0 {
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

        result, _, e := getTestApiClient().SlidesApi.NotesSlideExistsOnline(source, 1, "password")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.getExists() {
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

        _, _, e = getTestApiClient().SlidesApi.DownloadNotesSlideOnline(source, 1, "png", nil, nil, "password", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.getShapesLinks()))
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
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        dto.Text = "Updated shape"
        shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
		return
	}

        portion := NewPortion()
        portion.Text = "New Paragraph"
        dto := NewParagraph()
        dto.Alignment = "Right"
        dto.PortionList = []IPortion { portion }
        paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        dto = NewParagraph()
        dto.Alignment = "Center"
        paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getText() != dto.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        dto2 := NewPortion()
        dto2.Text = "Updated portion"
        dto2.FontHeight = 22
        portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount + 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getFontHeight() != dto2.getFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.getFontHeight(), portion.getFontHeight())
		return
	}
	if portion.getText() != dto2.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        _, e = c.SlidesApi.CopyFile("TempTests/" + sourceFile, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        masterSlides, _, e := c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.getSlideList()) != 1 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 1, len(masterSlides.getSlideList()))
		return
	}
        masterSlide, _, e := c.SlidesApi.GetMasterSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.getName() != "Office Theme" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Office Theme", masterSlide.getName())
		return
	}

        masterSlide, _, e = c.SlidesApi.CopyMasterSlide(fileName, sourcePath, 1, "", "", nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.getName() != "Digital portfolio" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Digital portfolio", masterSlide.getName())
		return
	}
        masterSlides, _, e = c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.getSlideList()) != 2 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 2, len(masterSlides.getSlideList()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.getShapesLinks()))
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
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        dto.Text = "Updated shape"
        shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
		return
	}

        portion := NewPortion()
        portion.Text = "New Paragraph"
        dto := NewParagraph()
        dto.Alignment = "Right"
        dto.PortionList = []IPortion { portion }
        paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        dto = NewParagraph()
        dto.Alignment = "Center"
        paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getText() != dto.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        dto2 := NewPortion()
        dto2.Text = "Updated portion"
        dto2.FontHeight = 22
        portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount + 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getFontHeight() != dto2.getFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.getFontHeight(), portion.getFontHeight())
		return
	}
	if portion.getText() != dto2.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}

        effect1 := NewEffect()
        effect1.Type_ = "Blink"
        effect1.ShapeIndex = 2
        effect2 := NewEffect()
        effect2.Type_ = "Appear"
        effect2.ShapeIndex = 3
        dto := NewSlideAnimation()
        dto.MainSequence = []IEffect { effect1, effect2 }
        animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != len(dto.getMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.getMainSequence()), len(animation.getMainSequence()))
		return
	}
        var shapeIndex int32 = 3
        animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "masterSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != len(dto.getMainSequence()) - 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.getMainSequence()) - 1, len(animation.getMainSequence()))
		return
	}
        animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}

        animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.getMainSequence()))
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
	dto2.Effects = []IEffect { effect }

	animation, _, e = c.SlidesApi.CreateAnimationInteractiveSequence(fileName, 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        _, e = c.SlidesApi.CopyFile("TempTests/" + sourceFile, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        layoutSlides, _, e := c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.getSlideList()) != 11 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 11, len(layoutSlides.getSlideList()))
		return
	}
        layoutSlide, _, e := c.SlidesApi.GetLayoutSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.getName() != "Title Slide" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title Slide", layoutSlide.getName())
		return
	}

        layoutSlide, _, e = c.SlidesApi.CopyLayoutSlide(fileName, sourcePath, 2, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.getName() != "Title and Content" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title and Content", layoutSlide.getName())
		return
	}
        layoutSlides, _, e = c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.getSlideList()) != 12 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 12, len(layoutSlides.getSlideList()))
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
        var shapeCount int32 = 5
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.getShapesLinks()))
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
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        dto.Text = "Updated shape"
        shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(IShape).getText() != dto.getText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.getText(), shape.(IShape).getText())
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) + 1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.getShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.getShapesLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
		return
	}

        portion := NewPortion()
        portion.Text = "New Paragraph"
        dto := NewParagraph()
        dto.Alignment = "Right"
        dto.PortionList = []IPortion { portion }
        paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        dto = NewParagraph()
        dto.Alignment = "Center"
        paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount + 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.getAlignment() != dto.getAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.getAlignment(), paragraph.getAlignment())
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.getParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.getParagraphLinks()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getText() != dto.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        dto2 := NewPortion()
        dto2.Text = "Updated portion"
        dto2.FontHeight = 22
        portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount + 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.getFontBold() != dto.getFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.getFontBold(), portion.getFontBold())
		return
	}
	if portion.getFontHeight() != dto2.getFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.getFontHeight(), portion.getFontHeight())
		return
	}
	if portion.getText() != dto2.getText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.getText(), portion.getText())
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()) + 1)
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount + 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.getItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.getItems()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}

        effect1 := NewEffect()
        effect1.Type_ = "Blink"
        effect1.ShapeIndex = 2
        effect2 := NewEffect()
        effect2.Type_ = "Appear"
        effect2.ShapeIndex = 3
        dto := NewSlideAnimation()
        dto.MainSequence = []IEffect { effect1, effect2 }
        animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != len(dto.getMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.getMainSequence()), len(animation.getMainSequence()))
		return
	}
        var shapeIndex int32 = 3
        animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}

        _, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "layoutSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != len(dto.getMainSequence()) - 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.getMainSequence()) - 1, len(animation.getMainSequence()))
		return
	}
        animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}

        animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}
}

/* 
   Test for add shape
*/
func TestShapeAdd(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        node1.Nodes = []ISmartArtNode { subNode1 }
        node2 := NewSmartArtNode()
        node2.Text = "Second"
        node2.OrgChartLayout = "Initial"
        dto.Nodes = []ISmartArtNode { node1, node2 }
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
   Test for add empty smart art
*/
func TestSmartArtEmpty(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	dto.Categories = []IChartCategory { category1, category2, category3 }

	series1 := NewOneValueSeries()
	dataPoint11 := NewOneValueChartDataPoint()
	dataPoint11.Value = 20

	dataPoint12 := NewOneValueChartDataPoint()
	dataPoint12.Value = 50

	dataPoint13 := NewOneValueChartDataPoint()
	dataPoint13.Value = 30
	series1.DataPoints = []IOneValueChartDataPoint { dataPoint11, dataPoint12, dataPoint13 }

	series2 := NewOneValueSeries()
	dataPoint21 := NewOneValueChartDataPoint()
	dataPoint21.Value = 30

	dataPoint22 := NewOneValueChartDataPoint()
	dataPoint22.Value = 10

	dataPoint23 := NewOneValueChartDataPoint()
	dataPoint23.Value = 60
	series2.DataPoints = []IOneValueChartDataPoint { dataPoint21, dataPoint22, dataPoint23 }
	dto.Series = []ISeries { series1, series2 }

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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        row1.Cells = []ITableCell { cell11, cell12, cell13, cell14 }
        row2 := NewTableRow()
        cell21 := NewTableCell()
        cell21.Text = "1"
        cell22 := NewTableCell()
        cell22.Text = "2-3"
        cell22.ColSpan = 2
        cell22.RowSpan = 2
        cell24 := NewTableCell()
        cell24.Text = "4"
        row2.Cells = []ITableCell { cell21, cell22, cell24 }
        row3 := NewTableRow()
        cell31 := NewTableCell()
        cell31.Text = "first"
        cell32 := NewTableCell()
        cell32.Text = "last"
        row3.Cells = []ITableCell { cell31, cell32 }
        row4 := NewTableRow()
        cell41 := NewTableCell()
        cell41.Text = "3.1"
        cell42 := NewTableCell()
        cell42.Text = "3.2"
        cell43 := NewTableCell()
        cell43.Text = "3.3"
        cell44 := NewTableCell()
        cell44.Text = "3.4"
        row4.Cells = []ITableCell { cell41, cell42, cell43, cell44 }
        row5 := NewTableRow()
        cell51 := NewTableCell()
        cell51.Text = "4.1"
        cell52 := NewTableCell()
        cell52.Text = "4.2"
        cell53 := NewTableCell()
        cell53.Text = "4.3"
        cell54 := NewTableCell()
        cell54.Text = "4.4"
        row5.Cells = []ITableCell { cell51, cell52, cell53, cell54 }
        dto.Rows = []ITableRow { row1, row2, row3, row4, row5 }
        column1 := NewTableColumn()
        column1.Width = 100
        column2 := NewTableColumn()
        column2.Width = 100
        column3 := NewTableColumn()
        column3.Width = 100
        column4 := NewTableColumn()
        column4.Width = 100
        dto.Columns = []ITableColumn { column1, column2, column3, column4 }
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        dto := NewGroupShape()
        _, response, e := c.SlidesApi.CreateShape(fileName, 1, dto, nil, nil, "password", folderName, "")
	if e == nil {
		t.Errorf("GroupShape should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/* 
   Test for add connector
*/
func TestConnectorAdd(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
   Test for align shapes
*/
func TestShapesAlign(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        var slideIndex int32 = 3
        var shape1Index int32 = 1
        var shape2Index int32 = 2
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if shape11.getX() == shape12.getX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape11.getX(), shape12.getX())
		return
	}
	if shape11.getY() == shape12.getY() {
		t.Errorf("Wrong Y value. Expected not %v but was %v.", shape11.getY(), shape12.getY())
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
	if shape21.getX() == shape22.getX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape21.getX(), shape22.getX())
		return
	}
	if math.Abs(shape21.getY() - shape22.getY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape21.getY(), shape22.getY())
		return
	}

	var alignToSlide bool = true
        _, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignLeft", &alignToSlide, []int32 { 1, 2 }, password, folderName, "")
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
	if math.Abs(shape31.getX() - shape32.getX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", shape31.getX(), shape32.getX())
		return
	}
	if math.Abs(shape31.getX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", 0, shape31.getX())
		return
	}
	if math.Abs(shape31.getY() - shape32.getY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape31.getY(), shape32.getY())
		return
	}
}

/* 
   Test for get chart
*/
func TestChartGet(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        chart, _, e := c.SlidesApi.GetShape(fileName, 3, 1, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(chart.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(chart.(IChart).getSeries()))
		return
	}
	if len(chart.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(chart.(IChart).getCategories()))
		return
	}
}

/* 
   Test for create chart
*/
func TestChartCreate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        series1.DataPoints = []IOneValueChartDataPoint { point11, point12, point13 }
        series2 := NewOneValueSeries()
        series2.Name = "Series2"
        point21 := NewOneValueChartDataPoint()
        point21.Value = 55
        point22 := NewOneValueChartDataPoint()
        point22.Value = 35
        point23 := NewOneValueChartDataPoint()
        point23.Value = 90
        series2.DataPoints = []IOneValueChartDataPoint { point21, point22, point23 }
        chart.Series = []ISeries { series1, series2 }
        category1 := NewChartCategory()
        category1.Value = "Category1"
        category2 := NewChartCategory()
        category2.Value = "Category2"
        category3 := NewChartCategory()
        category3.Value = "Category3"
        chart.Categories = []IChartCategory { category1, category2, category3 }
        result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).getCategories()))
		return
	}
}

/* 
   Test for update chart
*/
func TestChartUpdate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        series1.DataPoints = []IOneValueChartDataPoint { point11, point12, point13 }
        series2 := NewOneValueSeries()
        series2.Name = "Series2"
        point21 := NewOneValueChartDataPoint()
        point21.Value = 55
        point22 := NewOneValueChartDataPoint()
        point22.Value = 35
        point23 := NewOneValueChartDataPoint()
        point23.Value = 90
        series2.DataPoints = []IOneValueChartDataPoint { point21, point22, point23 }
        chart.Series = []ISeries { series1, series2 }
        category1 := NewChartCategory()
        category1.Value = "Category1"
        category2 := NewChartCategory()
        category2.Value = "Category2"
        category3 := NewChartCategory()
        category3.Value = "Category3"
        chart.Categories = []IChartCategory { category1, category2, category3 }
        result, _, e := c.SlidesApi.UpdateShape(fileName, 3, 1, chart, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).getCategories()))
		return
	}
}

/* 
   Test for create chart series
*/
func TestChartSeriesCreate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        series.DataPoints = []IOneValueChartDataPoint { point1, point2, point3, point4 }
        result, _, e := c.SlidesApi.CreateChartSeries(fileName, 3, 1, series, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 4 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 4, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
}

/* 
   Test for update chart series
*/
func TestChartSeriesUpdate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        series.DataPoints = []IOneValueChartDataPoint { point1, point2, point3, point4 }
        result, _, e := c.SlidesApi.UpdateChartSeries(fileName, 3, 1, 2, series, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
}

/* 
   Test for delete chart series
*/
func TestChartSeriesDelete(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteChartSeries(fileName, 3, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
}

/* 
   Test for create chart category
*/
func TestChartCategoryCreate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        category.DataPoints = []IOneValueChartDataPoint { point1, point2, point3 }
        result, _, e := c.SlidesApi.CreateChartCategory(fileName, 3, 1, category, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 5 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 5, len(result.(IChart).getCategories()))
		return
	}
	if len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()) != 5 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 5, len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()))
		return
	}
	if result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()[4].getValue() != category.DataPoints[0].getValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].getValue(), result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()[4].getValue())
		return
	}
}

/* 
   Test for update chart category
*/
func TestChartCategoryUpdate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        category.DataPoints = []IOneValueChartDataPoint { point1, point2, point3 }
        result, _, e := c.SlidesApi.UpdateChartCategory(fileName, 3, 1, 2, category, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
	if len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()))
		return
	}
	if result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()[1].getValue() != category.DataPoints[0].getValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].getValue(), result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()[1].getValue())
		return
	}
}

/* 
   Test for delete chart category
*/
func TestChartCategoryDelete(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteChartCategory(fileName, 3, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(IChart).getCategories()))
		return
	}
	if len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()) != 3 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()[0].(IOneValueSeries).getDataPoints()))
		return
	}
}

/* 
   Test for create chart data point
*/
func TestChartDataPointCreate(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
	if len(result.(IChart).getSeries()[1].(IOneValueSeries).getDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(IChart).getSeries()[1].(IOneValueSeries).getDataPoints()))
		return
	}
	if result.(IChart).getSeries()[1].(IOneValueSeries).getDataPoints()[1].getValue() != point.Value {
		t.Errorf("Wrong data point value. Expected %v but was %v.", point.Value, result.(IChart).getSeries()[1].(IOneValueSeries).getDataPoints()[1].getValue())
		return
	}
}

/* 
   Test for delete chart data point
*/
func TestChartDataPointDelete(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteChartDataPoint(fileName, 3, 1, 2, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
		return
	}
	if result.(IChart).getSeries()[1].(IOneValueSeries).getDataPoints()[1] != nil {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        series1.DataPoints = []IOneValueChartDataPoint { point1, point2, point3, point4 }
        chart.Series = []ISeries { series1 }
        category1 := NewChartCategory()
        category1.Value = "Leaf1"
        category1.Level = 3
        category1.ParentCategories = []string { "Branch1", "Stem1" }
        category2 := NewChartCategory()
        category2.Value = "Leaf2"
        category2.Level = 3
        category2.ParentCategories = []string { "Branch1", "Stem1" }
        category3 := NewChartCategory()
        category3.Value = "Branch2"
        category3.Level = 2
        category3.ParentCategories = []string { "Stem1" }
        category4 := NewChartCategory()
        category4.Value = "Stem2"
        category4.Level = 1
        chart.Categories = []IChartCategory { category1, category2, category3, category4 }
        result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(IChart).getSeries()) != 1 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 1, len(result.(IChart).getSeries()))
		return
	}
	if len(result.(IChart).getCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(IChart).getCategories()))
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

	c := getTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getLineFormat().getWidth() != dto.getLineFormat().getWidth() {
		t.Errorf("Wrong line width. Expected %v but was %v.", dto.getLineFormat().getWidth(), result.getLineFormat().getWidth())
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

	c := getTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	/* SLIDESCLOUD-1355
        _, ok = result.getFillFormat().(ISolidFill)
        if !ok {
		t.Errorf("Wrong fill type.")
		return
	}
	if result.getFillFormat().(ISolidFill).getColor() != dto.getFillFormat().(ISolidFill).getColor() {
		t.Errorf("Wrong fill color. Expected %v but was %v.", dto.getFillFormat().(ISolidFill).getColor(), result.getFillFormat().(ISolidFill).getColor())
		return
	}
	*/
}

/* 
   Test for shape effect format
*/
func TestShapeFormatEffect(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
        var slideIndex int32 = 1
        var shapeIndex int32 = 1

	c := getTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getEffectFormat().getInnerShadow().getDirection() != dto.getEffectFormat().getInnerShadow().getDirection() {
		t.Errorf("Wrong inner shadow direction. Expected %v but was %v.", dto.getEffectFormat().getInnerShadow().getDirection(), result.getEffectFormat().getInnerShadow().getDirection())
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

	c := getTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getThreeDFormat().getDepth() != dto.getThreeDFormat().getDepth() {
		t.Errorf("Wrong 3D depth. Expected %v but was %v.", dto.getThreeDFormat().getDepth(), result.getThreeDFormat().getDepth())
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !result.getIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.getIsDateTimeVisible() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !result.getIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.getIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
        result, _, e = c.SlidesApi.GetSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.getIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if result.getIsDateTimeVisible() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if !result.getIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if result.getIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
        result, _, e = c.SlidesApi.GetNotesSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !result.getIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if result.getIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/* 
   Test for get sections
*/
func TestSectionsGet(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.GetSections(fileName, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.getSectionList()))
		return
	}
}

/* 
   Test for replace sections
*/
func TestSectionsReplace(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        dto.SectionList = []ISection { section1, section2 }
        result, _, e := c.SlidesApi.SetSections(fileName, dto, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != len(dto.getSectionList()) {
		t.Errorf("Wrong section count. Expected %v but was %v.", len(dto.getSectionList()), len(result.getSectionList()))
		return
	}
	if len(result.getSectionList()[0].getSlideList()) != int(section2.FirstSlideIndex - section1.FirstSlideIndex) {
		t.Errorf("Wrong slide count. Expected %v but was %v.", section2.FirstSlideIndex - section1.FirstSlideIndex, len(result.getSectionList()[0].getSlideList()))
		return
	}
}

/* 
   Test for create section
*/
func TestSectionsPost(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.CreateSection(fileName, "NewSection", 5, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 4 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 4, len(result.getSectionList()))
		return
	}
}

/* 
   Test for update section
*/
func TestSectionsPut(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(result.getSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.getSectionList()))
		return
	}
	if result.getSectionList()[sectionIndex - 1].getName() != sectionName {
		t.Errorf("Wrong section name. Expected %v but was %v.", sectionName, result.getSectionList()[sectionIndex - 1].getName())
		return
	}
}

/* 
   Test for move section
*/
func TestSectionsMove(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.MoveSection(fileName, 1, 2, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.getSectionList()))
		return
	}
}

/* 
   Test for clear sections
*/
func TestSectionsClear(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteSections(fileName, nil, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 0 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 0, len(result.getSectionList()))
		return
	}
}

/* 
   Test for delete sections
*/
func TestSectionsDeleteMany(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteSections(fileName, []int32 { 2, 3 }, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 1 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 1, len(result.getSectionList()))
		return
	}
}

/* 
   Test for delete section
*/
func TestSectionsDelete(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.DeleteSection(fileName, 2, nil, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getSectionList()) != 2 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 2, len(result.getSectionList()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.getName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.getName())
		return
	}
	if !result.getBuiltIn() {
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
	if result.getName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.getName())
		return
	}
	if result.getValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.getValue())
		return
	}
	if !result.getBuiltIn() {
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
	if result.getName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.getName())
		return
	}
	if result.getValue() == updatedPropertyValue {
		t.Errorf("Wrong property value. Expected not %v but was %v.", updatedPropertyValue, result.getValue())
		return
	}
	if !result.getBuiltIn() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.getName())
		return
	}
	if result.getValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.getValue())
		return
	}
	if result.getBuiltIn() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        result, _, e := c.SlidesApi.GetDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	count := len(result.getList())

        property1 := NewDocumentProperty()
        property1.Name = propertyName
        property1.Value = updatedPropertyValue
        property2 := NewDocumentProperty()
        property2.Name = customPropertyName
        property2.Value = updatedPropertyValue
        properties := NewDocumentProperties()
        properties.List = []IDocumentProperty { property1, property2 }
        result, _, e = c.SlidesApi.SetDocumentProperties(fileName, properties, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getList()) != count + 1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count + 1, len(result.getList()))
		return
	}

        result, _, e = c.SlidesApi.DeleteDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.getList()) != count - 1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count - 1, len(result.getList()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        getResult, _, e := c.SlidesApi.GetSlideProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        dto := NewSlideProperties()
        dto.FirstSlideNumber = getResult.getFirstSlideNumber() + 2
        putResult, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.getOrientation() != getResult.getOrientation() {
		t.Errorf("Wrong orientation. Expected %v but was %v.", getResult.getOrientation(), putResult.getOrientation())
		return
	}
	if putResult.getFirstSlideNumber() == getResult.getFirstSlideNumber() {
		t.Errorf("Wrong FirstSlideNumber. Expected not %v but was %v.", getResult.getFirstSlideNumber(), putResult.getFirstSlideNumber())
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getSizeType() != "B4IsoPaper" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "B4IsoPaper", result.getSizeType())
		return
	}
	if result.getWidth() != 852 {
		t.Errorf("Wrong width. Expected %v but was %v.", 852, result.getWidth())
		return
	}
	if result.getHeight() != 639 {
		t.Errorf("Wrong height. Expected %v but was %v.", 639, result.getHeight())
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if result.getSizeType() != "Custom" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "Custom", result.getSizeType())
		return
	}
	if result.getWidth() != width {
		t.Errorf("Wrong width. Expected %v but was %v.", width, result.getWidth())
		return
	}
	if result.getHeight() != height {
		t.Errorf("Wrong height. Expected %v but was %v.", height, result.getHeight())
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        getResult, _, e := c.SlidesApi.GetProtectionProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        dto := NewProtectionProperties()
        dto.EncryptDocumentProperties = getResult.getEncryptDocumentProperties()
        dto.ReadOnlyRecommended = !getResult.getReadOnlyRecommended()
        putResult, _, e := c.SlidesApi.SetProtection(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.getEncryptDocumentProperties() != getResult.getEncryptDocumentProperties() {
		t.Errorf("Wrong EncryptDocumentProperties. Expected %v but was %v.", getResult.getEncryptDocumentProperties(), putResult.getEncryptDocumentProperties())
		return
	}
	if putResult.getReadOnlyRecommended() == getResult.getReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended. Expected not %v but was %v.", getResult.getReadOnlyRecommended(), putResult.getReadOnlyRecommended())
		return
	}
}

/* 
   Test for delete protection
*/
func TestPropertyProtectionDelete(t *testing.T) {
	folderName := "TempSlidesSDK"
	fileName := "test.pptx"
	c := getTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteProtection(fileName, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.getIsEncrypted() {
		t.Errorf("Wrong IsEncrypted value. Expected false.")
		return
	}
	if result.getReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended value. Expected false.")
		return
	}
	if result.getReadPassword() != "" {
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
        result, _, e := getTestApiClient().SlidesApi.SetProtectionOnline(source, dto, "password")
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
        result, _, e := getTestApiClient().SlidesApi.DeleteProtectionOnline(source, "password")
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
   Test for get image
*/
func TestImageGet(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(slideResult.getList()) >= len(presentationResult.getList()) {
		t.Errorf("Wrong image count. Expected less than %v but was %v.", len(presentationResult.getList()), len(slideResult.getList()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()

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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        c := getTestApiClient()

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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName2, folderName + "/" + fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        request := NewPresentationsMergeRequest()
        request.PresentationPaths = []string { folderName + "/" + fileName2 }
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName2, folderName + "/" + fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        request := NewOrderedMergeRequest()
        presentation := NewPresentationToMerge()
        presentation.Path = folderName + "/" + fileName2
        presentation.Slides = []int32 { 2, 1 }
        request.Presentations = []IPresentationToMerge { presentation }
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

        _, _, e = getTestApiClient().SlidesApi.MergeOnline([][]byte { source1, source2 }, nil, "")
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
        c := getTestApiClient()

        _, e = c.SlidesApi.MergeAndSaveOnline(outPath, [][]byte { source1, source2 }, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.getExists() {
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
        presentation2.Slides = []int32 { 1, 2 }
        request.Presentations = []IPresentationToMerge { presentation1, presentation2 }
        _, _, e = getTestApiClient().SlidesApi.MergeOnline([][]byte { source1, source2 }, request, "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName2, folderName + "/" + fileName2, "", "", "")
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
        presentation2.Slides = []int32 { 1, 2 }
        request.Presentations = []IPresentationToMerge { presentation1, presentation2 }
        _, _, e = getTestApiClient().SlidesApi.MergeOnline([][]byte { source }, request, "")
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(result2.getSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.getSlides()) <= len(result2.getSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.getSlides()), len(result1.getSlides()))
		return
	}
        url := result1.getSlides()[0].getHref()
        path := url[strings.Index(url, "/storage/file/") + len("/storage/file/"):]
        resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.getExists() {
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
        c := getTestApiClient()

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
        c := getTestApiClient()

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
	if len(result2.getSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.getSlides()) <= len(result2.getSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.getSlides()), len(result1.getSlides()))
		return
	}
        url := result1.getSlides()[0].getHref()
        path := url[strings.Index(url, "/storage/file/") + len("/storage/file/"):]
        resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.getExists() {
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
        url := result.getSlides()[0].getHref()
        path := url[strings.Index(url, "/storage/file/") + len("/storage/file/"):]
        resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.getExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for get text
*/
func TestTextGet(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        var slideIndex int32 = 1
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(result.getItems()) >= len(resultWithEmpty.getItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(result.getItems()), len(resultWithEmpty.getItems()))
		return
	}
	if len(slideResult.getItems()) >= len(result.getItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.getItems()), len(result.getItems()))
		return
	}
	if len(slideResult.getItems()) >= len(slideResultWithEmpty.getItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.getItems()), len(slideResultWithEmpty.getItems()))
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
        c := getTestApiClient()
        var ignoreCase bool = true

        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        result, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        resultWithEmpty, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        slideResult, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        _, e = c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        slideResultWithEmpty, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if result.getMatches() >= resultWithEmpty.getMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", result.getMatches(), resultWithEmpty.getMatches())
		return
	}
	if slideResult.getMatches() >= result.getMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.getMatches(), result.getMatches())
		return
	}
	if slideResult.getMatches() >= slideResultWithEmpty.getMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.getMatches(), slideResultWithEmpty.getMatches())
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
        c := getTestApiClient()

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
   Test for text watermark on storage
*/
func TestWatermarkTextStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        var slideIndex int32 = 1
        watermarkText := "watermarkText"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapeCount := len(getResult.getShapesLinks()) + 1

        _, e = c.SlidesApi.CreateWatermark(fileName, nil, nil, watermarkText, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.getShapesLinks()))
		return
	}
        getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.getName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.getName())
		return
	}
	if getShapeResult.(IShape).getText() != watermarkText {
		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, getShapeResult.(IShape).getText())
		return
	}

        _, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount - 1 != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(getResult.getShapesLinks()))
		return
	}
}

/*
   Test for text DTO watermark on storage
*/
func TestWatermarkTextDTOStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        var slideIndex int32 = 1
        watermarkText := "watermarkText"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapeCount := len(getResult.getShapesLinks()) + 1

        watermark := NewShape()
        watermark.Text = watermarkText
        _, e = c.SlidesApi.CreateWatermark(fileName, watermark, nil, watermarkText, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.getShapesLinks()))
		return
	}
        getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.getName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.getName())
		return
	}
	if getShapeResult.(IShape).getText() != watermarkText {
		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, getShapeResult.(IShape).getText())
		return
	}

        _, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount - 1 != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(getResult.getShapesLinks()))
		return
	}
}

/*
   Test for image watermark on storage
*/
func TestWatermarkImageStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        var slideIndex int32 = 1
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapeCount := len(getResult.getShapesLinks()) + 1

	source, e := ioutil.ReadFile("TestData/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        _, e = c.SlidesApi.CreateImageWatermark(fileName, source, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.getShapesLinks()))
		return
	}
        getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.getName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.getName())
		return
	}

        _, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount - 1 != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(getResult.getShapesLinks()))
		return
	}
}

/*
   Test for image DTO watermark on storage
*/
func TestWatermarkImageDTOStorage(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        password := "password"
        watermarkName := "myWatermark"
        var slideIndex int32 = 1
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        shapeCount := len(getResult.getShapesLinks()) + 1

	source, e := ioutil.ReadFile("TestData/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        watermark := NewPictureFrame()
        fillFormat := NewPictureFill()
        fillFormat.Base64Data = base64.StdEncoding.EncodeToString(source)
        watermark.FillFormat = fillFormat
        watermark.Name = watermarkName
        _, e = c.SlidesApi.CreateImageWatermark(fileName, nil, watermark, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.getShapesLinks()))
		return
	}
        getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.getName() != watermarkName {
		t.Errorf("Wrong shape name. Expected %v but was %v.", watermarkName, getShapeResult.getName())
		return
	}

        _, e = c.SlidesApi.DeleteWatermark(fileName, watermarkName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount - 1 != len(getResult.getShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(getResult.getShapesLinks()))
		return
	}
}

/*
   Test for text watermark from request
*/
func TestWatermarkTextRequest(t *testing.T) {
        password := "password"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        c := getTestApiClient()

        postResult, _, e := c.SlidesApi.CreateWatermarkOnline(source, nil, nil, "watermark", "", "", password)
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
   Test for text DTO watermark from request
*/
func TestWatermarkTextDTORequest(t *testing.T) {
        password := "password"
	source, e := ioutil.ReadFile("TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        c := getTestApiClient()

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
        c := getTestApiClient()

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
        c := getTestApiClient()

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
   Test for get animation
*/
func TestAnimationGet(t *testing.T) {
        folderName := "TempSlidesSDK"
        fileName := "test.pptx"
        var slideIndex int32 = 1
        var shapeIndex int32 = 3
        var paragraphIndex int32 = 1
        password := "password"
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.GetAnimation(fileName, slideIndex, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
		return
	}

        animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.getInteractiveSequences()))
		return
	}

        animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &shapeIndex, &paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.getMainSequence()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	dto.MainSequence = []IEffect { effect1, effect2 }
	dto.InteractiveSequences = []IInteractiveSequence {}
        animation, _, e := c.SlidesApi.SetAnimation(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != len(dto.getMainSequence()) {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", len(dto.getMainSequence()), len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(animation.getMainSequence()) != 2 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 2, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        dto := NewInteractiveSequence()
        dto.TriggerShapeIndex = 2
	effect := NewEffect()
	effect.Type_ = "Blast"
	effect.ShapeIndex = 3
        dto.Effects = []IEffect { effect }
        animation, _, e := c.SlidesApi.CreateAnimationInteractiveSequence(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 2 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 2, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
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
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimation(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimationMainSequence(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimationEffect(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequences(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequence(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.getInteractiveSequences()))
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
        c := getTestApiClient()
        _, e := c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.getMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.getMainSequence()))
		return
	}
	if len(animation.getInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.getInteractiveSequences()))
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
        pipeline.Tasks = []ITask { task }

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
        files := [][]byte { data1, data2 }

	_, _, e = getTestApiClient().SlidesApi.Pipeline(pipeline, files)
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

	r, _, e := getTestApiClient().SlidesApi.GetShape("test.pptx", 1, 1, "password", "TempSlidesSDK", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IShape).getText() != "1" {
		t.Errorf("Error: Text expected to equal 1.")
		return
	}
}

/* 
   Test for Chart creation
*/
func TestChart(t *testing.T) {
	chart := NewChart()
	if chart.getType() != "Chart" {
		t.Errorf("Unexpected chart type: %v.", chart.getType())
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
        var min1  = 44.3
        var min2 = 12.0
        var max1 = 104.3
        var max2 = 87.0

	e := initializeTest("NoFunction", "No method", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := getTestApiClient()
	_, e = c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        var dto1 Chart
        dto1.setType("Chart")
        dto1.setChartType("Line")
        dto1.setWidth(400)
        dto1.setHeight(300)
        var title ChartTitle
        title.setHasTitle(true)
        title.setText("MyTitle")
        dto1.setTitle(&title)
        var series OneValueSeries
        series.setType("ClusteredColumn")
        series.setDataPointType("OneValue")
        series.setName("Series1")
        var point1 OneValueChartDataPoint
        point1.setValue(40)
        var point2 OneValueChartDataPoint
        point2.setValue(50)
        points := make([]IOneValueChartDataPoint, 2)
        points[0] = &point1
        points[1] = &point2
        series.setDataPoints(points)
        serieses := make([]ISeries, 1)
        serieses[0] = &series
        dto1.setSeries(serieses)
        var axes Axes
        var axis1 Axis
        axis1.setIsAutomaticMinValue(false)
        axis1.setMinValue(min1)
        axis1.setIsAutomaticMaxValue(false)
        axis1.setMaxValue(max1)
        axes.setHorizontalAxis(&axis1)
        dto1.setAxes(&axes)
        dto1.setX(12)
        dto1.setY(14)
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
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min1 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

        var dto2 Chart
        dto2.setType("Chart")
        dto2.setChartType("Line")
        var axis2 Axis
        axis2.setMinValue(min2)
        axes.setHorizontalAxis(&axis2)
        dto2.setAxes(&axes)
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
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

        var axis3 Axis
        axis3.setMaxValue(max2)
        axes.setHorizontalAxis(&axis3)
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
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max2 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}
}
