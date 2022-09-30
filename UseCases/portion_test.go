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
	"strings"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v22"
)

/*
   Test for get portions
*/
func TestPortionsGet(t *testing.T) {
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := slidescloud.NewPortion()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := slidescloud.NewPortion()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := slidescloud.NewPortion()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	portionText := "portion 1"
	fontName := "Arial"

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = "#FFF5FF8A"
	dto := slidescloud.NewPortion()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
   Test for get portion effective
*/
func TestPortionEffectiveGet(t *testing.T) {
	var slideIndex int32 = 6
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	c := slidescloud.GetTestApiClient()
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
