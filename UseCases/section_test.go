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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v24"
)

/*
   Test for download empty math
*/
func TestMathDownloadNull(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, response, e := c.SlidesApi.DownloadMathPortion(fileName, 2, 1, 1, 1, "mathml", "password", folderName, "")
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
	outPath := folderName + "/mathml.xml"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.SaveMathPortion(fileName, 2, 3, 1, 1, "mathml", outPath, "password", folderName, "")
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
   Test for header/footer on all slides
*/
func TestHeaderFooterSlides(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewHeaderFooter()
	isFooterVisible := true
	dto.IsFooterVisible = &isFooterVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
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
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on one slide
*/
func TestHeaderFooterSlide(t *testing.T) {
	var slideIndex int32 = 1
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewHeaderFooter()
	isFooterVisible := true
	dto.IsFooterVisible = &isFooterVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
	result, _, e := c.SlidesApi.SetSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on notes slide
*/
func TestHeaderFooterNotesSlide(t *testing.T) {
	var slideIndex int32 = 1
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewNotesSlideHeaderFooter()
	isHeaderVisible := true
	dto.IsHeaderVisible = &isHeaderVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
	result, _, e := c.SlidesApi.SetNotesSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetNotesSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for Get sections
*/
func TestSectionsGet(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetSections(fileName, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSections()
	section1 := slidescloud.NewSection()
	section1.Name = "Section1"
	section1.FirstSlideIndex = 1
	section2 := slidescloud.NewSection()
	section2.Name = "Section2"
	section2.FirstSlideIndex = 3
	dto.SectionList = []slidescloud.ISection{section1, section2}
	result, _, e := c.SlidesApi.SetSections(fileName, dto, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.CreateSection(fileName, "NewSection", 5, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var sectionIndex int32 = 2
	sectionName := "UpdatedSection"
	result, _, e := c.SlidesApi.UpdateSection(fileName, sectionIndex, sectionName, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.MoveSection(fileName, 1, 2, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSections(fileName, nil, nil, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSections(fileName, []int32{2, 3}, nil, password, folderName, "")
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
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteSection(fileName, 2, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 2 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 2, len(result.GetSectionList()))
		return
	}
}
