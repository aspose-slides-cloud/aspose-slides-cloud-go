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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v22"
)

/*
   Test for create empty presentation
*/
func TestCreateEmpty(t *testing.T) {
	c := slidescloud.GetTestApiClient()
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
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data, e := ioutil.ReadFile("../TestData/" + fileName)
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
	newFileName := "test2.pptx"
	c := slidescloud.GetTestApiClient()
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
	templateFileName := "TemplateCV.pptx"
	c := slidescloud.GetTestApiClient()
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
	c := slidescloud.GetTestApiClient()
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
	c := slidescloud.GetTestApiClient()
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
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.DeleteFile(folderName+"/"+fileName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source, e := ioutil.ReadFile("../TestData/test.pdf")
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
	c := slidescloud.GetTestApiClient()
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
	source, e := ioutil.ReadFile("../TestData/test.pdf")
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
