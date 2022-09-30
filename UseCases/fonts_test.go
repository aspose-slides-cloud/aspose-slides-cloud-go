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
   Test for get fonts
*/
func TestFontsGet(t *testing.T) {
	c := slidescloud.GetTestApiClient()
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

	c := slidescloud.GetTestApiClient()
	document, e := ioutil.ReadFile(localTestFile)

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
	fontName := "Calibri"
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "", "")

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
	fontName := "Calibri"

	c := slidescloud.GetTestApiClient()
	document, e := ioutil.ReadFile(localTestFile)

	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for set embedded font from request
*/
func TestSetEmbeddedFontFromRequest(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	document, e := ioutil.ReadFile(localFontFile)

	c = slidescloud.GetTestApiClient()
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false
	response, _, e := c.SlidesApi.SetEmbeddedFontFromRequest(document, fileName, &onlyUsed, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetList()[2].GetIsEmbedded() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
	fontName := "Calibri"
	if response.GetList()[2].GetFontName() != fontName {
		t.Errorf("Expected %v, but was %v", fontName, response.GetList()[2].GetFontName())
		return
	}
}

/*
   Test for set embedded font from request
*/
func TestSetEmbeddedFontFromRequestOnline(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	document, e := ioutil.ReadFile(localTestFile)

	c = slidescloud.GetTestApiClient()
	fontFile, e := ioutil.ReadFile(localFontFile)

	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontFromRequestOnline(document, fontFile, &onlyUsed, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for set embedded font
*/
func TestEmbeddedFontDelete(t *testing.T) {
	fontName := "Calibri"
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "", "")

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

/*
   Test for replace font
*/
func TestReplaceFont(t *testing.T) {

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	sourceFontName := "Calibri"
	targetFontName := "Times New Roman"
	var embed bool = true
	response, _, e := c.SlidesApi.ReplaceFont(fileName, sourceFontName, targetFontName, &embed, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetList()[2].GetIsEmbedded() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
	if response.GetList()[2].GetFontName() != targetFontName {
		t.Errorf("Expected %v, but was %v", targetFontName, response.GetList()[2].GetFontName())
		return
	}
}

/*
   Test for replace font online
*/
func TestReplaceFontOnline(t *testing.T) {

	c := slidescloud.GetTestApiClient()
	document, e := ioutil.ReadFile(localTestFile)

	sourceFontName := "Calibri"
	targetFontName := "Times New Roman"
	var embed bool = true
	_, _, e = c.SlidesApi.ReplaceFontOnline(document, sourceFontName, targetFontName, &embed, password, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for font substitution
*/
func TestFontSubstitution(t *testing.T) {

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	targetFontName := "Times New Roman"

	fontRule1 := slidescloud.NewFontSubstRule()
	fontRule1.SetSourceFont("Arial")
	fontRule1.SetTargetFont(targetFontName)
	fontRule1.SetNotFoundOnly(false)
	fontRule2 := slidescloud.NewFontSubstRule()
	fontRule2.SetSourceFont("Calibri")
	fontRule2.SetTargetFont(targetFontName)

	exportOptions := slidescloud.NewImageExportOptions()
	exportOptions.SetFontSubstRules([]slidescloud.IFontSubstRule{fontRule1, fontRule2})

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "PNG", exportOptions, password, folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
