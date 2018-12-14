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
	"testing"
)

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method
*/
func TestDeleteNotesSlideParagraph(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    e := initializeTest("DeleteNotesSlideParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlideParagraphRequest() DeleteNotesSlideParagraphRequest {
    var request DeleteNotesSlideParagraphRequest
    request.name = createTestParamValue("DeleteNotesSlideParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlideParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlideParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteNotesSlideParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeleteNotesSlideParagraph", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteNotesSlideParagraph", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlideParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlideParagraph", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid name
*/
func TestDeleteNotesSlideParagraphInvalidname(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid path
*/
func TestDeleteNotesSlideParagraphInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid paragraphIndex
*/
func TestDeleteNotesSlideParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid password
*/
func TestDeleteNotesSlideParagraphInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid folder
*/
func TestDeleteNotesSlideParagraphInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraph method with invalid storage
*/
func TestDeleteNotesSlideParagraphInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method
*/
func TestDeleteNotesSlideParagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    e := initializeTest("DeleteNotesSlideParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlideParagraphsRequest() DeleteNotesSlideParagraphsRequest {
    var request DeleteNotesSlideParagraphsRequest
    request.name = createTestParamValue("DeleteNotesSlideParagraphs", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlideParagraphs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlideParagraphs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteNotesSlideParagraphs", "shapeIndex", "int32").(int32)
    request.paragraphs = createTestParamValue("DeleteNotesSlideParagraphs", "paragraphs", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteNotesSlideParagraphs", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlideParagraphs", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlideParagraphs", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid name
*/
func TestDeleteNotesSlideParagraphsInvalidname(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid path
*/
func TestDeleteNotesSlideParagraphsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid paragraphs
*/
func TestDeleteNotesSlideParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.paragraphs = invalidizeTestParamValue(request.paragraphs, "paragraphs", "[]int32").([]int32)
    e := initializeTest("DeleteNotesSlideParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "paragraphs", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid password
*/
func TestDeleteNotesSlideParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid folder
*/
func TestDeleteNotesSlideParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideParagraphs method with invalid storage
*/
func TestDeleteNotesSlideParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlideParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method
*/
func TestDeleteNotesSlidePortion(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    e := initializeTest("DeleteNotesSlidePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlidePortionRequest() DeleteNotesSlidePortionRequest {
    var request DeleteNotesSlidePortionRequest
    request.name = createTestParamValue("DeleteNotesSlidePortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlidePortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlidePortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteNotesSlidePortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeleteNotesSlidePortion", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("DeleteNotesSlidePortion", "portionIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteNotesSlidePortion", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlidePortion", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlidePortion", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid name
*/
func TestDeleteNotesSlidePortionInvalidname(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlidePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid path
*/
func TestDeleteNotesSlidePortionInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlidePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid portionIndex
*/
func TestDeleteNotesSlidePortionInvalidportionIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "portionIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid password
*/
func TestDeleteNotesSlidePortionInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlidePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid folder
*/
func TestDeleteNotesSlidePortionInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlidePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortion method with invalid storage
*/
func TestDeleteNotesSlidePortionInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlidePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method
*/
func TestDeleteNotesSlidePortions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    e := initializeTest("DeleteNotesSlidePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlidePortionsRequest() DeleteNotesSlidePortionsRequest {
    var request DeleteNotesSlidePortionsRequest
    request.name = createTestParamValue("DeleteNotesSlidePortions", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlidePortions", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlidePortions", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteNotesSlidePortions", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeleteNotesSlidePortions", "paragraphIndex", "int32").(int32)
    request.portions = createTestParamValue("DeleteNotesSlidePortions", "portions", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteNotesSlidePortions", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlidePortions", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlidePortions", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid name
*/
func TestDeleteNotesSlidePortionsInvalidname(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlidePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid path
*/
func TestDeleteNotesSlidePortionsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlidePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlidePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid portions
*/
func TestDeleteNotesSlidePortionsInvalidportions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.portions = invalidizeTestParamValue(request.portions, "portions", "[]int32").([]int32)
    e := initializeTest("DeleteNotesSlidePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "portions", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid password
*/
func TestDeleteNotesSlidePortionsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlidePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid folder
*/
func TestDeleteNotesSlidePortionsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlidePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlidePortions method with invalid storage
*/
func TestDeleteNotesSlidePortionsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlidePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method
*/
func TestDeleteNotesSlideShape(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    e := initializeTest("DeleteNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlideShapeRequest() DeleteNotesSlideShapeRequest {
    var request DeleteNotesSlideShapeRequest
    request.name = createTestParamValue("DeleteNotesSlideShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlideShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlideShape", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteNotesSlideShape", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteNotesSlideShape", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlideShape", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlideShape", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid name
*/
func TestDeleteNotesSlideShapeInvalidname(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid slideIndex
*/
func TestDeleteNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid path
*/
func TestDeleteNotesSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid shapeIndex
*/
func TestDeleteNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid password
*/
func TestDeleteNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid folder
*/
func TestDeleteNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShape method with invalid storage
*/
func TestDeleteNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method
*/
func TestDeleteNotesSlideShapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    e := initializeTest("DeleteNotesSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlideShapesRequest() DeleteNotesSlideShapesRequest {
    var request DeleteNotesSlideShapesRequest
    request.name = createTestParamValue("DeleteNotesSlideShapes", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlideShapes", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteNotesSlideShapes", "path", "string").(string)
    request.shapes = createTestParamValue("DeleteNotesSlideShapes", "shapes", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteNotesSlideShapes", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlideShapes", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlideShapes", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid name
*/
func TestDeleteNotesSlideShapesInvalidname(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid slideIndex
*/
func TestDeleteNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid path
*/
func TestDeleteNotesSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid shapes
*/
func TestDeleteNotesSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.shapes = invalidizeTestParamValue(request.shapes, "shapes", "[]int32").([]int32)
    e := initializeTest("DeleteNotesSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "shapes", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid password
*/
func TestDeleteNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid folder
*/
func TestDeleteNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for NotesSlideShapesApi.DeleteNotesSlideShapes method with invalid storage
*/
func TestDeleteNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method
*/
func TestGetNotesSlideShape(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    e := initializeTest("GetNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapeRequest() GetNotesSlideShapeRequest {
    var request GetNotesSlideShapeRequest
    request.name = createTestParamValue("GetNotesSlideShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShape", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetNotesSlideShape", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlideShape", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShape", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShape", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid name
*/
func TestGetNotesSlideShapeInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid slideIndex
*/
func TestGetNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid path
*/
func TestGetNotesSlideShapeInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid shapeIndex
*/
func TestGetNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid password
*/
func TestGetNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid folder
*/
func TestGetNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShape method with invalid storage
*/
func TestGetNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method
*/
func TestGetNotesSlideShapeParagraph(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    e := initializeTest("GetNotesSlideShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapeParagraphRequest() GetNotesSlideShapeParagraphRequest {
    var request GetNotesSlideShapeParagraphRequest
    request.name = createTestParamValue("GetNotesSlideShapeParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShapeParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetNotesSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetNotesSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlideShapeParagraph", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapeParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid name
*/
func TestGetNotesSlideShapeParagraphInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid path
*/
func TestGetNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid password
*/
func TestGetNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid folder
*/
func TestGetNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraph method with invalid storage
*/
func TestGetNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method
*/
func TestGetNotesSlideShapeParagraphs(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    e := initializeTest("GetNotesSlideShapeParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapeParagraphsRequest() GetNotesSlideShapeParagraphsRequest {
    var request GetNotesSlideShapeParagraphsRequest
    request.name = createTestParamValue("GetNotesSlideShapeParagraphs", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapeParagraphs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShapeParagraphs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetNotesSlideShapeParagraphs", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlideShapeParagraphs", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapeParagraphs", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapeParagraphs", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid name
*/
func TestGetNotesSlideShapeParagraphsInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid path
*/
func TestGetNotesSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid password
*/
func TestGetNotesSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid folder
*/
func TestGetNotesSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapeParagraphs method with invalid storage
*/
func TestGetNotesSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method
*/
func TestGetNotesSlideShapePortion(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    e := initializeTest("GetNotesSlideShapePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapePortionRequest() GetNotesSlideShapePortionRequest {
    var request GetNotesSlideShapePortionRequest
    request.name = createTestParamValue("GetNotesSlideShapePortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapePortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShapePortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetNotesSlideShapePortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetNotesSlideShapePortion", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("GetNotesSlideShapePortion", "portionIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlideShapePortion", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapePortion", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapePortion", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid name
*/
func TestGetNotesSlideShapePortionInvalidname(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid path
*/
func TestGetNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid portionIndex
*/
func TestGetNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "portionIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid password
*/
func TestGetNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid folder
*/
func TestGetNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortion method with invalid storage
*/
func TestGetNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method
*/
func TestGetNotesSlideShapePortions(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    e := initializeTest("GetNotesSlideShapePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapePortionsRequest() GetNotesSlideShapePortionsRequest {
    var request GetNotesSlideShapePortionsRequest
    request.name = createTestParamValue("GetNotesSlideShapePortions", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapePortions", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShapePortions", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetNotesSlideShapePortions", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetNotesSlideShapePortions", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlideShapePortions", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapePortions", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapePortions", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid name
*/
func TestGetNotesSlideShapePortionsInvalidname(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid path
*/
func TestGetNotesSlideShapePortionsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShapePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid password
*/
func TestGetNotesSlideShapePortionsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid folder
*/
func TestGetNotesSlideShapePortionsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapePortions method with invalid storage
*/
func TestGetNotesSlideShapePortionsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method
*/
func TestGetNotesSlideShapeWithFormat(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    e := initializeTest("GetNotesSlideShapeWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetNotesSlideShapeWithFormatRequest() GetNotesSlideShapeWithFormatRequest {
    var request GetNotesSlideShapeWithFormatRequest
    request.name = createTestParamValue("GetNotesSlideShapeWithFormat", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapeWithFormat", "slideIndex", "int32").(int32)
    request.shapeIndex = createTestParamValue("GetNotesSlideShapeWithFormat", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("GetNotesSlideShapeWithFormat", "format", "string").(string)
    request.password = createTestParamValue("GetNotesSlideShapeWithFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapeWithFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapeWithFormat", "storage", "string").(string)
    request.scaleX = createTestParamValue("GetNotesSlideShapeWithFormat", "scaleX", "float64").(float64)
    request.scaleY = createTestParamValue("GetNotesSlideShapeWithFormat", "scaleY", "float64").(float64)
    request.bounds = createTestParamValue("GetNotesSlideShapeWithFormat", "bounds", "string").(string)
    request.outPath = createTestParamValue("GetNotesSlideShapeWithFormat", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("GetNotesSlideShapeWithFormat", "fontsFolder", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid name
*/
func TestGetNotesSlideShapeWithFormatInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "name", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid slideIndex
*/
func TestGetNotesSlideShapeWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "slideIndex", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid shapeIndex
*/
func TestGetNotesSlideShapeWithFormatInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapeWithFormat", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "shapeIndex", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid format
*/
func TestGetNotesSlideShapeWithFormatInvalidformat(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "format", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid password
*/
func TestGetNotesSlideShapeWithFormatInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "password", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid folder
*/
func TestGetNotesSlideShapeWithFormatInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "folder", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid storage
*/
func TestGetNotesSlideShapeWithFormatInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "storage", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid scaleX
*/
func TestGetNotesSlideShapeWithFormatInvalidscaleX(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)
    e := initializeTest("GetNotesSlideShapeWithFormat", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "scaleX", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid scaleY
*/
func TestGetNotesSlideShapeWithFormatInvalidscaleY(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)
    e := initializeTest("GetNotesSlideShapeWithFormat", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "scaleY", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid bounds
*/
func TestGetNotesSlideShapeWithFormatInvalidbounds(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "bounds", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid outPath
*/
func TestGetNotesSlideShapeWithFormatInvalidoutPath(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "outPath", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.GetNotesSlideShapeWithFormat method with invalid fontsFolder
*/
func TestGetNotesSlideShapeWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetNotesSlideShapeWithFormatRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("GetNotesSlideShapeWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapeWithFormat(request)
    assertError(t, "GetNotesSlideShapeWithFormat", "fontsFolder", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method
*/
func TestGetNotesSlideShapes(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    e := initializeTest("GetNotesSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.GetNotesSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideShapesRequest() GetNotesSlideShapesRequest {
    var request GetNotesSlideShapesRequest
    request.name = createTestParamValue("GetNotesSlideShapes", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideShapes", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetNotesSlideShapes", "path", "string").(string)
    request.password = createTestParamValue("GetNotesSlideShapes", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideShapes", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideShapes", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid name
*/
func TestGetNotesSlideShapesInvalidname(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid slideIndex
*/
func TestGetNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid path
*/
func TestGetNotesSlideShapesInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid password
*/
func TestGetNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid folder
*/
func TestGetNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Read slide shapes or shape info.
   Test for NotesSlideShapesApi.GetNotesSlideShapes method with invalid storage
*/
func TestGetNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method
*/
func TestPostNotesSlideAddNewParagraph(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    e := initializeTest("PostNotesSlideAddNewParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostNotesSlideAddNewParagraphRequest() PostNotesSlideAddNewParagraphRequest {
    var request PostNotesSlideAddNewParagraphRequest
    request.name = createTestParamValue("PostNotesSlideAddNewParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostNotesSlideAddNewParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostNotesSlideAddNewParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostNotesSlideAddNewParagraph", "shapeIndex", "int32").(int32)
    request.dto = createTestParamValue("PostNotesSlideAddNewParagraph", "dto", "Paragraph").(IParagraph)
    request.password = createTestParamValue("PostNotesSlideAddNewParagraph", "password", "string").(string)
    request.folder = createTestParamValue("PostNotesSlideAddNewParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("PostNotesSlideAddNewParagraph", "storage", "string").(string)
    request.position = createTestParamValue("PostNotesSlideAddNewParagraph", "position", "int32").(int32)
    return request
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid name
*/
func TestPostNotesSlideAddNewParagraphInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostNotesSlideAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid slideIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid path
*/
func TestPostNotesSlideAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostNotesSlideAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid dto
*/
func TestPostNotesSlideAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)
    e := initializeTest("PostNotesSlideAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid password
*/
func TestPostNotesSlideAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostNotesSlideAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid folder
*/
func TestPostNotesSlideAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostNotesSlideAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid storage
*/
func TestPostNotesSlideAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostNotesSlideAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new paragraph.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewParagraph method with invalid position
*/
func TestPostNotesSlideAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "position", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method
*/
func TestPostNotesSlideAddNewPortion(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    e := initializeTest("PostNotesSlideAddNewPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostNotesSlideAddNewPortionRequest() PostNotesSlideAddNewPortionRequest {
    var request PostNotesSlideAddNewPortionRequest
    request.name = createTestParamValue("PostNotesSlideAddNewPortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostNotesSlideAddNewPortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostNotesSlideAddNewPortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostNotesSlideAddNewPortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PostNotesSlideAddNewPortion", "paragraphIndex", "int32").(int32)
    request.dto = createTestParamValue("PostNotesSlideAddNewPortion", "dto", "Portion").(IPortion)
    request.password = createTestParamValue("PostNotesSlideAddNewPortion", "password", "string").(string)
    request.folder = createTestParamValue("PostNotesSlideAddNewPortion", "folder", "string").(string)
    request.storage = createTestParamValue("PostNotesSlideAddNewPortion", "storage", "string").(string)
    request.position = createTestParamValue("PostNotesSlideAddNewPortion", "position", "int32").(int32)
    return request
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid name
*/
func TestPostNotesSlideAddNewPortionInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostNotesSlideAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid slideIndex
*/
func TestPostNotesSlideAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid path
*/
func TestPostNotesSlideAddNewPortionInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostNotesSlideAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid paragraphIndex
*/
func TestPostNotesSlideAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid dto
*/
func TestPostNotesSlideAddNewPortionInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)
    e := initializeTest("PostNotesSlideAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid password
*/
func TestPostNotesSlideAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostNotesSlideAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid folder
*/
func TestPostNotesSlideAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostNotesSlideAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid storage
*/
func TestPostNotesSlideAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostNotesSlideAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new portion.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewPortion method with invalid position
*/
func TestPostNotesSlideAddNewPortionInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "position", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method
*/
func TestPostNotesSlideAddNewShape(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    e := initializeTest("PostNotesSlideAddNewShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostNotesSlideAddNewShapeRequest() PostNotesSlideAddNewShapeRequest {
    var request PostNotesSlideAddNewShapeRequest
    request.name = createTestParamValue("PostNotesSlideAddNewShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostNotesSlideAddNewShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostNotesSlideAddNewShape", "path", "string").(string)
    request.dto = createTestParamValue("PostNotesSlideAddNewShape", "dto", "ShapeBase").(IShapeBase)
    request.password = createTestParamValue("PostNotesSlideAddNewShape", "password", "string").(string)
    request.folder = createTestParamValue("PostNotesSlideAddNewShape", "folder", "string").(string)
    request.storage = createTestParamValue("PostNotesSlideAddNewShape", "storage", "string").(string)
    request.shapeToClone = createTestParamValue("PostNotesSlideAddNewShape", "shapeToClone", "int32").(int32)
    request.position = createTestParamValue("PostNotesSlideAddNewShape", "position", "int32").(int32)
    return request
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid name
*/
func TestPostNotesSlideAddNewShapeInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostNotesSlideAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid slideIndex
*/
func TestPostNotesSlideAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid path
*/
func TestPostNotesSlideAddNewShapeInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostNotesSlideAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid dto
*/
func TestPostNotesSlideAddNewShapeInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)
    e := initializeTest("PostNotesSlideAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid password
*/
func TestPostNotesSlideAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostNotesSlideAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid folder
*/
func TestPostNotesSlideAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostNotesSlideAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid storage
*/
func TestPostNotesSlideAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostNotesSlideAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid shapeToClone
*/
func TestPostNotesSlideAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.shapeToClone = invalidizeTestParamValue(request.shapeToClone, "shapeToClone", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "shapeToClone", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Creates new shape.
   Test for NotesSlideShapesApi.PostNotesSlideAddNewShape method with invalid position
*/
func TestPostNotesSlideAddNewShapeInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostNotesSlideAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "position", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method
*/
func TestPostNotesSlideShapeSaveAs(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    e := initializeTest("PostNotesSlideShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostNotesSlideShapeSaveAsRequest() PostNotesSlideShapeSaveAsRequest {
    var request PostNotesSlideShapeSaveAsRequest
    request.name = createTestParamValue("PostNotesSlideShapeSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostNotesSlideShapeSaveAs", "slideIndex", "int32").(int32)
    request.shapeIndex = createTestParamValue("PostNotesSlideShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PostNotesSlideShapeSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostNotesSlideShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PostNotesSlideShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostNotesSlideShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostNotesSlideShapeSaveAs", "storage", "string").(string)
    request.scaleX = createTestParamValue("PostNotesSlideShapeSaveAs", "scaleX", "float64").(float64)
    request.scaleY = createTestParamValue("PostNotesSlideShapeSaveAs", "scaleY", "float64").(float64)
    request.bounds = createTestParamValue("PostNotesSlideShapeSaveAs", "bounds", "string").(string)
    request.outPath = createTestParamValue("PostNotesSlideShapeSaveAs", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("PostNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid name
*/
func TestPostNotesSlideShapeSaveAsInvalidname(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid format
*/
func TestPostNotesSlideShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid options
*/
func TestPostNotesSlideShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)
    e := initializeTest("PostNotesSlideShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid password
*/
func TestPostNotesSlideShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid folder
*/
func TestPostNotesSlideShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid storage
*/
func TestPostNotesSlideShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)
    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)
    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPostNotesSlideShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid outPath
*/
func TestPostNotesSlideShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "outPath", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Render shape to specified picture format.
   Test for NotesSlideShapesApi.PostNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPostNotesSlideShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("PostNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideShapesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method
*/
func TestPutUpdateNotesSlideShape(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    e := initializeTest("PutUpdateNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutUpdateNotesSlideShapeRequest() PutUpdateNotesSlideShapeRequest {
    var request PutUpdateNotesSlideShapeRequest
    request.name = createTestParamValue("PutUpdateNotesSlideShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutUpdateNotesSlideShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutUpdateNotesSlideShape", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutUpdateNotesSlideShape", "shapeIndex", "int32").(int32)
    request.dto = createTestParamValue("PutUpdateNotesSlideShape", "dto", "ShapeBase").(IShapeBase)
    request.password = createTestParamValue("PutUpdateNotesSlideShape", "password", "string").(string)
    request.folder = createTestParamValue("PutUpdateNotesSlideShape", "folder", "string").(string)
    request.storage = createTestParamValue("PutUpdateNotesSlideShape", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid name
*/
func TestPutUpdateNotesSlideShapeInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid path
*/
func TestPutUpdateNotesSlideShapeInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid dto
*/
func TestPutUpdateNotesSlideShapeInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)
    e := initializeTest("PutUpdateNotesSlideShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid password
*/
func TestPutUpdateNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid folder
*/
func TestPutUpdateNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShape method with invalid storage
*/
func TestPutUpdateNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method
*/
func TestPutUpdateNotesSlideShapeParagraph(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutUpdateNotesSlideShapeParagraphRequest() PutUpdateNotesSlideShapeParagraphRequest {
    var request PutUpdateNotesSlideShapeParagraphRequest
    request.name = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.dto = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "dto", "Paragraph").(IParagraph)
    request.password = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "password", "string").(string)
    request.folder = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid name
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid path
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid dto
*/
func TestPutUpdateNotesSlideShapeParagraphInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid password
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid folder
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph method with invalid storage
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "storage", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method
*/
func TestPutUpdateNotesSlideShapePortion(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    e := initializeTest("PutUpdateNotesSlideShapePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutUpdateNotesSlideShapePortionRequest() PutUpdateNotesSlideShapePortionRequest {
    var request PutUpdateNotesSlideShapePortionRequest
    request.name = createTestParamValue("PutUpdateNotesSlideShapePortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutUpdateNotesSlideShapePortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "portionIndex", "int32").(int32)
    request.dto = createTestParamValue("PutUpdateNotesSlideShapePortion", "dto", "Portion").(IPortion)
    request.password = createTestParamValue("PutUpdateNotesSlideShapePortion", "password", "string").(string)
    request.folder = createTestParamValue("PutUpdateNotesSlideShapePortion", "folder", "string").(string)
    request.storage = createTestParamValue("PutUpdateNotesSlideShapePortion", "storage", "string").(string)
    return request
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid name
*/
func TestPutUpdateNotesSlideShapePortionInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "name", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "slideIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid path
*/
func TestPutUpdateNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "path", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "shapeIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "paragraphIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid portionIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "portionIndex", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid dto
*/
func TestPutUpdateNotesSlideShapePortionInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "dto", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid password
*/
func TestPutUpdateNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "password", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid folder
*/
func TestPutUpdateNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "folder", r.Code, e)
}

/* NotesSlideShapesApiServiceTests Updates shape properties.
   Test for NotesSlideShapesApi.PutUpdateNotesSlideShapePortion method with invalid storage
*/
func TestPutUpdateNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutUpdateNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideShapesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "storage", r.Code, e)
}
