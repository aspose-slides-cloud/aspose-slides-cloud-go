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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v22"
)

/*
   Test for get paragraph
*/
func TestParagraphGet(t *testing.T) {
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewParagraph()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion1 := slidescloud.NewPortion()
	portion1.Text = "Portion1"
	portion2 := slidescloud.NewPortion()
	portion2.Text = "Portion2"
	portion2.FontBold = "True"
	dto := slidescloud.NewParagraph()
	dto.PortionList = []slidescloud.IPortion{portion1, portion2}

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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewParagraph()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewParagraph()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewParagraph()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	portion1 := slidescloud.NewPortion()
	portion1.Text = "Portion1"
	portion2 := slidescloud.NewPortion()
	portion2.Text = "Portion2"

	portionFormat := slidescloud.NewPortionFormat()
	portionFormat.FontHeight = 20
	portionFormat.LatinFont = "Arial"

	dto := slidescloud.NewParagraph()
	dto.DefaultPortionFormat = portionFormat
	dto.PortionList = []slidescloud.IPortion{portion1, portion2}

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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
