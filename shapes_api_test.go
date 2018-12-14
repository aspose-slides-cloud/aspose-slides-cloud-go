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

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method
*/
func TestDeleteParagraph(t *testing.T) {
    request := createDeleteParagraphRequest()
    e := initializeTest("DeleteParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeleteParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteParagraphRequest() DeleteParagraphRequest {
    var request DeleteParagraphRequest
    request.name = createTestParamValue("DeleteParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeleteParagraph", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteParagraph", "password", "string").(string)
    request.folder = createTestParamValue("DeleteParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteParagraph", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid name
*/
func TestDeleteParagraphInvalidname(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid slideIndex
*/
func TestDeleteParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid path
*/
func TestDeleteParagraphInvalidpath(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid shapeIndex
*/
func TestDeleteParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid paragraphIndex
*/
func TestDeleteParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeleteParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid password
*/
func TestDeleteParagraphInvalidpassword(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid folder
*/
func TestDeleteParagraphInvalidfolder(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraph method with invalid storage
*/
func TestDeleteParagraphInvalidstorage(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "storage", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method
*/
func TestDeleteParagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()
    e := initializeTest("DeleteParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeleteParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteParagraphsRequest() DeleteParagraphsRequest {
    var request DeleteParagraphsRequest
    request.name = createTestParamValue("DeleteParagraphs", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteParagraphs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteParagraphs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteParagraphs", "shapeIndex", "int32").(int32)
    request.paragraphs = createTestParamValue("DeleteParagraphs", "paragraphs", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteParagraphs", "password", "string").(string)
    request.folder = createTestParamValue("DeleteParagraphs", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteParagraphs", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid name
*/
func TestDeleteParagraphsInvalidname(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid slideIndex
*/
func TestDeleteParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid path
*/
func TestDeleteParagraphsInvalidpath(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid shapeIndex
*/
func TestDeleteParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid paragraphs
*/
func TestDeleteParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.paragraphs = invalidizeTestParamValue(request.paragraphs, "paragraphs", "[]int32").([]int32)
    e := initializeTest("DeleteParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "paragraphs", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid password
*/
func TestDeleteParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid folder
*/
func TestDeleteParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteParagraphs method with invalid storage
*/
func TestDeleteParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "storage", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method
*/
func TestDeletePortion(t *testing.T) {
    request := createDeletePortionRequest()
    e := initializeTest("DeletePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeletePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeletePortionRequest() DeletePortionRequest {
    var request DeletePortionRequest
    request.name = createTestParamValue("DeletePortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeletePortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeletePortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeletePortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeletePortion", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("DeletePortion", "portionIndex", "int32").(int32)
    request.password = createTestParamValue("DeletePortion", "password", "string").(string)
    request.folder = createTestParamValue("DeletePortion", "folder", "string").(string)
    request.storage = createTestParamValue("DeletePortion", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid name
*/
func TestDeletePortionInvalidname(t *testing.T) {
    request := createDeletePortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeletePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid slideIndex
*/
func TestDeletePortionInvalidslideIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeletePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid path
*/
func TestDeletePortionInvalidpath(t *testing.T) {
    request := createDeletePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeletePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid shapeIndex
*/
func TestDeletePortionInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeletePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid paragraphIndex
*/
func TestDeletePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeletePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid portionIndex
*/
func TestDeletePortionInvalidportionIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("DeletePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "portionIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid password
*/
func TestDeletePortionInvalidpassword(t *testing.T) {
    request := createDeletePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeletePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid folder
*/
func TestDeletePortionInvalidfolder(t *testing.T) {
    request := createDeletePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeletePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortion method with invalid storage
*/
func TestDeletePortionInvalidstorage(t *testing.T) {
    request := createDeletePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeletePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "storage", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method
*/
func TestDeletePortions(t *testing.T) {
    request := createDeletePortionsRequest()
    e := initializeTest("DeletePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeletePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeletePortionsRequest() DeletePortionsRequest {
    var request DeletePortionsRequest
    request.name = createTestParamValue("DeletePortions", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeletePortions", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeletePortions", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeletePortions", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("DeletePortions", "paragraphIndex", "int32").(int32)
    request.portions = createTestParamValue("DeletePortions", "portions", "[]int32").([]int32)
    request.password = createTestParamValue("DeletePortions", "password", "string").(string)
    request.folder = createTestParamValue("DeletePortions", "folder", "string").(string)
    request.storage = createTestParamValue("DeletePortions", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid name
*/
func TestDeletePortionsInvalidname(t *testing.T) {
    request := createDeletePortionsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeletePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid slideIndex
*/
func TestDeletePortionsInvalidslideIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeletePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid path
*/
func TestDeletePortionsInvalidpath(t *testing.T) {
    request := createDeletePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeletePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid shapeIndex
*/
func TestDeletePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeletePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid paragraphIndex
*/
func TestDeletePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("DeletePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid portions
*/
func TestDeletePortionsInvalidportions(t *testing.T) {
    request := createDeletePortionsRequest()
    request.portions = invalidizeTestParamValue(request.portions, "portions", "[]int32").([]int32)
    e := initializeTest("DeletePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "portions", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid password
*/
func TestDeletePortionsInvalidpassword(t *testing.T) {
    request := createDeletePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeletePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid folder
*/
func TestDeletePortionsInvalidfolder(t *testing.T) {
    request := createDeletePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeletePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeletePortions method with invalid storage
*/
func TestDeletePortionsInvalidstorage(t *testing.T) {
    request := createDeletePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeletePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "storage", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method
*/
func TestDeleteSlideShape(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    e := initializeTest("DeleteSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeleteSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteSlideShapeRequest() DeleteSlideShapeRequest {
    var request DeleteSlideShapeRequest
    request.name = createTestParamValue("DeleteSlideShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteSlideShape", "path", "string").(string)
    request.shapeIndex = createTestParamValue("DeleteSlideShape", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideShape", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideShape", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideShape", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid name
*/
func TestDeleteSlideShapeInvalidname(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid slideIndex
*/
func TestDeleteSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid path
*/
func TestDeleteSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid shapeIndex
*/
func TestDeleteSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("DeleteSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid password
*/
func TestDeleteSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid folder
*/
func TestDeleteSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShape method with invalid storage
*/
func TestDeleteSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "storage", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method
*/
func TestDeleteSlideShapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    e := initializeTest("DeleteSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.DeleteSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteSlideShapesRequest() DeleteSlideShapesRequest {
    var request DeleteSlideShapesRequest
    request.name = createTestParamValue("DeleteSlideShapes", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideShapes", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("DeleteSlideShapes", "path", "string").(string)
    request.shapes = createTestParamValue("DeleteSlideShapes", "shapes", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteSlideShapes", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideShapes", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideShapes", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid name
*/
func TestDeleteSlideShapesInvalidname(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "name", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid slideIndex
*/
func TestDeleteSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid path
*/
func TestDeleteSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("DeleteSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "path", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid shapes
*/
func TestDeleteSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.shapes = invalidizeTestParamValue(request.shapes, "shapes", "[]int32").([]int32)
    e := initializeTest("DeleteSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "shapes", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid password
*/
func TestDeleteSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "password", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid folder
*/
func TestDeleteSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "folder", r.Code, e)
}

/* ShapesApiServiceTests Removes a shape, specified shapes or all shapes.
   Test for ShapesApi.DeleteSlideShapes method with invalid storage
*/
func TestDeleteSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "storage", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method
*/
func TestGetParagraphPortion(t *testing.T) {
    request := createGetParagraphPortionRequest()
    e := initializeTest("GetParagraphPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetParagraphPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetParagraphPortionRequest() GetParagraphPortionRequest {
    var request GetParagraphPortionRequest
    request.name = createTestParamValue("GetParagraphPortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetParagraphPortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetParagraphPortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetParagraphPortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetParagraphPortion", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("GetParagraphPortion", "portionIndex", "int32").(int32)
    request.password = createTestParamValue("GetParagraphPortion", "password", "string").(string)
    request.folder = createTestParamValue("GetParagraphPortion", "folder", "string").(string)
    request.storage = createTestParamValue("GetParagraphPortion", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid name
*/
func TestGetParagraphPortionInvalidname(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetParagraphPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid slideIndex
*/
func TestGetParagraphPortionInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid path
*/
func TestGetParagraphPortionInvalidpath(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetParagraphPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid shapeIndex
*/
func TestGetParagraphPortionInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid paragraphIndex
*/
func TestGetParagraphPortionInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid portionIndex
*/
func TestGetParagraphPortionInvalidportionIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "portionIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid password
*/
func TestGetParagraphPortionInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetParagraphPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid folder
*/
func TestGetParagraphPortionInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetParagraphPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortion method with invalid storage
*/
func TestGetParagraphPortionInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetParagraphPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "storage", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method
*/
func TestGetParagraphPortions(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    e := initializeTest("GetParagraphPortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetParagraphPortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetParagraphPortionsRequest() GetParagraphPortionsRequest {
    var request GetParagraphPortionsRequest
    request.name = createTestParamValue("GetParagraphPortions", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetParagraphPortions", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetParagraphPortions", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetParagraphPortions", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetParagraphPortions", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("GetParagraphPortions", "password", "string").(string)
    request.folder = createTestParamValue("GetParagraphPortions", "folder", "string").(string)
    request.storage = createTestParamValue("GetParagraphPortions", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid name
*/
func TestGetParagraphPortionsInvalidname(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetParagraphPortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid slideIndex
*/
func TestGetParagraphPortionsInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid path
*/
func TestGetParagraphPortionsInvalidpath(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetParagraphPortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid shapeIndex
*/
func TestGetParagraphPortionsInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid paragraphIndex
*/
func TestGetParagraphPortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetParagraphPortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid password
*/
func TestGetParagraphPortionsInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetParagraphPortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid folder
*/
func TestGetParagraphPortionsInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetParagraphPortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetParagraphPortions method with invalid storage
*/
func TestGetParagraphPortionsInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetParagraphPortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "storage", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method
*/
func TestGetShapeParagraph(t *testing.T) {
    request := createGetShapeParagraphRequest()
    e := initializeTest("GetShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetShapeParagraphRequest() GetShapeParagraphRequest {
    var request GetShapeParagraphRequest
    request.name = createTestParamValue("GetShapeParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetShapeParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetShapeParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetShapeParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetShapeParagraph", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("GetShapeParagraph", "password", "string").(string)
    request.folder = createTestParamValue("GetShapeParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("GetShapeParagraph", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid name
*/
func TestGetShapeParagraphInvalidname(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid slideIndex
*/
func TestGetShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid path
*/
func TestGetShapeParagraphInvalidpath(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid shapeIndex
*/
func TestGetShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid paragraphIndex
*/
func TestGetShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("GetShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid password
*/
func TestGetShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid folder
*/
func TestGetShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetShapeParagraph method with invalid storage
*/
func TestGetShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetShapeParagraph(request)
    assertError(t, "GetShapeParagraph", "storage", r.Code, e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method
*/
func TestGetShapeWithFormat(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    e := initializeTest("GetShapeWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetShapeWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetShapeWithFormatRequest() GetShapeWithFormatRequest {
    var request GetShapeWithFormatRequest
    request.name = createTestParamValue("GetShapeWithFormat", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetShapeWithFormat", "slideIndex", "int32").(int32)
    request.shapeIndex = createTestParamValue("GetShapeWithFormat", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("GetShapeWithFormat", "format", "string").(string)
    request.password = createTestParamValue("GetShapeWithFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetShapeWithFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetShapeWithFormat", "storage", "string").(string)
    request.scaleX = createTestParamValue("GetShapeWithFormat", "scaleX", "float64").(float64)
    request.scaleY = createTestParamValue("GetShapeWithFormat", "scaleY", "float64").(float64)
    request.bounds = createTestParamValue("GetShapeWithFormat", "bounds", "string").(string)
    request.outPath = createTestParamValue("GetShapeWithFormat", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("GetShapeWithFormat", "fontsFolder", "string").(string)
    return request
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid name
*/
func TestGetShapeWithFormatInvalidname(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetShapeWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "name", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid slideIndex
*/
func TestGetShapeWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetShapeWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "slideIndex", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid shapeIndex
*/
func TestGetShapeWithFormatInvalidshapeIndex(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetShapeWithFormat", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "shapeIndex", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid format
*/
func TestGetShapeWithFormatInvalidformat(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("GetShapeWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "format", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid password
*/
func TestGetShapeWithFormatInvalidpassword(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetShapeWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "password", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid folder
*/
func TestGetShapeWithFormatInvalidfolder(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetShapeWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "folder", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid storage
*/
func TestGetShapeWithFormatInvalidstorage(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetShapeWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "storage", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid scaleX
*/
func TestGetShapeWithFormatInvalidscaleX(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)
    e := initializeTest("GetShapeWithFormat", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "scaleX", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid scaleY
*/
func TestGetShapeWithFormatInvalidscaleY(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)
    e := initializeTest("GetShapeWithFormat", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "scaleY", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid bounds
*/
func TestGetShapeWithFormatInvalidbounds(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "string").(string)
    e := initializeTest("GetShapeWithFormat", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "bounds", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid outPath
*/
func TestGetShapeWithFormatInvalidoutPath(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("GetShapeWithFormat", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "outPath", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.GetShapeWithFormat method with invalid fontsFolder
*/
func TestGetShapeWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetShapeWithFormatRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("GetShapeWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.GetShapeWithFormat(request)
    assertError(t, "GetShapeWithFormat", "fontsFolder", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method
*/
func TestGetSlideShape(t *testing.T) {
    request := createGetSlideShapeRequest()
    e := initializeTest("GetSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlideShapeRequest() GetSlideShapeRequest {
    var request GetSlideShapeRequest
    request.name = createTestParamValue("GetSlideShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetSlideShape", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetSlideShape", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlideShape", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideShape", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideShape", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid name
*/
func TestGetSlideShapeInvalidname(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid slideIndex
*/
func TestGetSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid path
*/
func TestGetSlideShapeInvalidpath(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid shapeIndex
*/
func TestGetSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid password
*/
func TestGetSlideShapeInvalidpassword(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid folder
*/
func TestGetSlideShapeInvalidfolder(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShape method with invalid storage
*/
func TestGetSlideShapeInvalidstorage(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "storage", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method
*/
func TestGetSlideShapeParagraphs(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    e := initializeTest("GetSlideShapeParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetSlideShapeParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlideShapeParagraphsRequest() GetSlideShapeParagraphsRequest {
    var request GetSlideShapeParagraphsRequest
    request.name = createTestParamValue("GetSlideShapeParagraphs", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideShapeParagraphs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetSlideShapeParagraphs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetSlideShapeParagraphs", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlideShapeParagraphs", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideShapeParagraphs", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideShapeParagraphs", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid name
*/
func TestGetSlideShapeParagraphsInvalidname(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid path
*/
func TestGetSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("GetSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid password
*/
func TestGetSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid folder
*/
func TestGetSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapeParagraphs method with invalid storage
*/
func TestGetSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "storage", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method
*/
func TestGetSlideShapes(t *testing.T) {
    request := createGetSlideShapesRequest()
    e := initializeTest("GetSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.GetSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlideShapesRequest() GetSlideShapesRequest {
    var request GetSlideShapesRequest
    request.name = createTestParamValue("GetSlideShapes", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideShapes", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetSlideShapes", "path", "string").(string)
    request.password = createTestParamValue("GetSlideShapes", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideShapes", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideShapes", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid name
*/
func TestGetSlideShapesInvalidname(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "name", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid slideIndex
*/
func TestGetSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid path
*/
func TestGetSlideShapesInvalidpath(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("GetSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "path", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid password
*/
func TestGetSlideShapesInvalidpassword(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "password", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid folder
*/
func TestGetSlideShapesInvalidfolder(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "folder", r.Code, e)
}

/* ShapesApiServiceTests Read slide shapes or shape info.
   Test for ShapesApi.GetSlideShapes method with invalid storage
*/
func TestGetSlideShapesInvalidstorage(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "storage", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method
*/
func TestPostAddNewParagraph(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    e := initializeTest("PostAddNewParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PostAddNewParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostAddNewParagraphRequest() PostAddNewParagraphRequest {
    var request PostAddNewParagraphRequest
    request.name = createTestParamValue("PostAddNewParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostAddNewParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostAddNewParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostAddNewParagraph", "shapeIndex", "int32").(int32)
    request.dto = createTestParamValue("PostAddNewParagraph", "dto", "Paragraph").(IParagraph)
    request.password = createTestParamValue("PostAddNewParagraph", "password", "string").(string)
    request.folder = createTestParamValue("PostAddNewParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("PostAddNewParagraph", "storage", "string").(string)
    request.position = createTestParamValue("PostAddNewParagraph", "position", "int32").(int32)
    return request
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid name
*/
func TestPostAddNewParagraphInvalidname(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "name", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid slideIndex
*/
func TestPostAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid path
*/
func TestPostAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "path", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid shapeIndex
*/
func TestPostAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid dto
*/
func TestPostAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)
    e := initializeTest("PostAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "dto", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid password
*/
func TestPostAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "password", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid folder
*/
func TestPostAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "folder", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid storage
*/
func TestPostAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "storage", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewParagraph method with invalid position
*/
func TestPostAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "position", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method
*/
func TestPostAddNewPortion(t *testing.T) {
    request := createPostAddNewPortionRequest()
    e := initializeTest("PostAddNewPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PostAddNewPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostAddNewPortionRequest() PostAddNewPortionRequest {
    var request PostAddNewPortionRequest
    request.name = createTestParamValue("PostAddNewPortion", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostAddNewPortion", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostAddNewPortion", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostAddNewPortion", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PostAddNewPortion", "paragraphIndex", "int32").(int32)
    request.dto = createTestParamValue("PostAddNewPortion", "dto", "Portion").(IPortion)
    request.password = createTestParamValue("PostAddNewPortion", "password", "string").(string)
    request.folder = createTestParamValue("PostAddNewPortion", "folder", "string").(string)
    request.storage = createTestParamValue("PostAddNewPortion", "storage", "string").(string)
    request.position = createTestParamValue("PostAddNewPortion", "position", "int32").(int32)
    return request
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid name
*/
func TestPostAddNewPortionInvalidname(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "name", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid slideIndex
*/
func TestPostAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid path
*/
func TestPostAddNewPortionInvalidpath(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "path", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid shapeIndex
*/
func TestPostAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid paragraphIndex
*/
func TestPostAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PostAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid dto
*/
func TestPostAddNewPortionInvaliddto(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)
    e := initializeTest("PostAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "dto", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid password
*/
func TestPostAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "password", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid folder
*/
func TestPostAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "folder", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid storage
*/
func TestPostAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "storage", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewPortion method with invalid position
*/
func TestPostAddNewPortionInvalidposition(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "position", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method
*/
func TestPostAddNewShape(t *testing.T) {
    request := createPostAddNewShapeRequest()
    e := initializeTest("PostAddNewShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PostAddNewShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostAddNewShapeRequest() PostAddNewShapeRequest {
    var request PostAddNewShapeRequest
    request.name = createTestParamValue("PostAddNewShape", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostAddNewShape", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PostAddNewShape", "path", "string").(string)
    request.dto = createTestParamValue("PostAddNewShape", "dto", "ShapeBase").(IShapeBase)
    request.password = createTestParamValue("PostAddNewShape", "password", "string").(string)
    request.folder = createTestParamValue("PostAddNewShape", "folder", "string").(string)
    request.storage = createTestParamValue("PostAddNewShape", "storage", "string").(string)
    request.shapeToClone = createTestParamValue("PostAddNewShape", "shapeToClone", "int32").(int32)
    request.position = createTestParamValue("PostAddNewShape", "position", "int32").(int32)
    return request
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid name
*/
func TestPostAddNewShapeInvalidname(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "name", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid slideIndex
*/
func TestPostAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid path
*/
func TestPostAddNewShapeInvalidpath(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PostAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "path", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid dto
*/
func TestPostAddNewShapeInvaliddto(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)
    e := initializeTest("PostAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "dto", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid password
*/
func TestPostAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "password", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid folder
*/
func TestPostAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "folder", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid storage
*/
func TestPostAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "storage", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid shapeToClone
*/
func TestPostAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.shapeToClone = invalidizeTestParamValue(request.shapeToClone, "shapeToClone", "int32").(int32)
    e := initializeTest("PostAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "shapeToClone", r.Code, e)
}

/* ShapesApiServiceTests Creates new shape.
   Test for ShapesApi.PostAddNewShape method with invalid position
*/
func TestPostAddNewShapeInvalidposition(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "position", r.Code, e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method
*/
func TestPostShapeSaveAs(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    e := initializeTest("PostShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PostShapeSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostShapeSaveAsRequest() PostShapeSaveAsRequest {
    var request PostShapeSaveAsRequest
    request.name = createTestParamValue("PostShapeSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostShapeSaveAs", "slideIndex", "int32").(int32)
    request.shapeIndex = createTestParamValue("PostShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PostShapeSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PostShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostShapeSaveAs", "storage", "string").(string)
    request.scaleX = createTestParamValue("PostShapeSaveAs", "scaleX", "float64").(float64)
    request.scaleY = createTestParamValue("PostShapeSaveAs", "scaleY", "float64").(float64)
    request.bounds = createTestParamValue("PostShapeSaveAs", "bounds", "string").(string)
    request.outPath = createTestParamValue("PostShapeSaveAs", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("PostShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid name
*/
func TestPostShapeSaveAsInvalidname(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid slideIndex
*/
func TestPostShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid shapeIndex
*/
func TestPostShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PostShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid format
*/
func TestPostShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("PostShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid options
*/
func TestPostShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)
    e := initializeTest("PostShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid password
*/
func TestPostShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid folder
*/
func TestPostShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid storage
*/
func TestPostShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid scaleX
*/
func TestPostShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)
    e := initializeTest("PostShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid scaleY
*/
func TestPostShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)
    e := initializeTest("PostShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid bounds
*/
func TestPostShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "string").(string)
    e := initializeTest("PostShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid outPath
*/
func TestPostShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("PostShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "outPath", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Render shape to specified picture format.
   Test for ShapesApi.PostShapeSaveAs method with invalid fontsFolder
*/
func TestPostShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("PostShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ShapesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method
*/
func TestPutSetParagraphPortionProperties(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    e := initializeTest("PutSetParagraphPortionProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PutSetParagraphPortionProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSetParagraphPortionPropertiesRequest() PutSetParagraphPortionPropertiesRequest {
    var request PutSetParagraphPortionPropertiesRequest
    request.name = createTestParamValue("PutSetParagraphPortionProperties", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSetParagraphPortionProperties", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutSetParagraphPortionProperties", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutSetParagraphPortionProperties", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PutSetParagraphPortionProperties", "paragraphIndex", "int32").(int32)
    request.portionIndex = createTestParamValue("PutSetParagraphPortionProperties", "portionIndex", "int32").(int32)
    request.dto = createTestParamValue("PutSetParagraphPortionProperties", "dto", "Portion").(IPortion)
    request.password = createTestParamValue("PutSetParagraphPortionProperties", "password", "string").(string)
    request.folder = createTestParamValue("PutSetParagraphPortionProperties", "folder", "string").(string)
    request.storage = createTestParamValue("PutSetParagraphPortionProperties", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid name
*/
func TestPutSetParagraphPortionPropertiesInvalidname(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutSetParagraphPortionProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "name", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid slideIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphPortionProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid path
*/
func TestPutSetParagraphPortionPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutSetParagraphPortionProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "path", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphPortionProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphPortionProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid portionIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidportionIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphPortionProperties", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "portionIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid dto
*/
func TestPutSetParagraphPortionPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)
    e := initializeTest("PutSetParagraphPortionProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "dto", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid password
*/
func TestPutSetParagraphPortionPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutSetParagraphPortionProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "password", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid folder
*/
func TestPutSetParagraphPortionPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutSetParagraphPortionProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "folder", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphPortionProperties method with invalid storage
*/
func TestPutSetParagraphPortionPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutSetParagraphPortionProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "storage", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method
*/
func TestPutSetParagraphProperties(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    e := initializeTest("PutSetParagraphProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PutSetParagraphProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSetParagraphPropertiesRequest() PutSetParagraphPropertiesRequest {
    var request PutSetParagraphPropertiesRequest
    request.name = createTestParamValue("PutSetParagraphProperties", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSetParagraphProperties", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutSetParagraphProperties", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutSetParagraphProperties", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("PutSetParagraphProperties", "paragraphIndex", "int32").(int32)
    request.dto = createTestParamValue("PutSetParagraphProperties", "dto", "Paragraph").(IParagraph)
    request.password = createTestParamValue("PutSetParagraphProperties", "password", "string").(string)
    request.folder = createTestParamValue("PutSetParagraphProperties", "folder", "string").(string)
    request.storage = createTestParamValue("PutSetParagraphProperties", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid name
*/
func TestPutSetParagraphPropertiesInvalidname(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutSetParagraphProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "name", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid slideIndex
*/
func TestPutSetParagraphPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid path
*/
func TestPutSetParagraphPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutSetParagraphProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "path", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)
    e := initializeTest("PutSetParagraphProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "paragraphIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid dto
*/
func TestPutSetParagraphPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)
    e := initializeTest("PutSetParagraphProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "dto", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid password
*/
func TestPutSetParagraphPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutSetParagraphProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "password", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid folder
*/
func TestPutSetParagraphPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutSetParagraphProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "folder", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSetParagraphProperties method with invalid storage
*/
func TestPutSetParagraphPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutSetParagraphProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "storage", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method
*/
func TestPutSlideShapeInfo(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    e := initializeTest("PutSlideShapeInfo", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ShapesApi.PutSlideShapeInfo(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSlideShapeInfoRequest() PutSlideShapeInfoRequest {
    var request PutSlideShapeInfoRequest
    request.name = createTestParamValue("PutSlideShapeInfo", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlideShapeInfo", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutSlideShapeInfo", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutSlideShapeInfo", "shapeIndex", "int32").(int32)
    request.dto = createTestParamValue("PutSlideShapeInfo", "dto", "ShapeBase").(IShapeBase)
    request.password = createTestParamValue("PutSlideShapeInfo", "password", "string").(string)
    request.folder = createTestParamValue("PutSlideShapeInfo", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlideShapeInfo", "storage", "string").(string)
    return request
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid name
*/
func TestPutSlideShapeInfoInvalidname(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutSlideShapeInfo", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "name", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid slideIndex
*/
func TestPutSlideShapeInfoInvalidslideIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutSlideShapeInfo", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "slideIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid path
*/
func TestPutSlideShapeInfoInvalidpath(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)
    e := initializeTest("PutSlideShapeInfo", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "path", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid shapeIndex
*/
func TestPutSlideShapeInfoInvalidshapeIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)
    e := initializeTest("PutSlideShapeInfo", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "shapeIndex", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid dto
*/
func TestPutSlideShapeInfoInvaliddto(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)
    e := initializeTest("PutSlideShapeInfo", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "dto", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid password
*/
func TestPutSlideShapeInfoInvalidpassword(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutSlideShapeInfo", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "password", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid folder
*/
func TestPutSlideShapeInfoInvalidfolder(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutSlideShapeInfo", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "folder", r.Code, e)
}

/* ShapesApiServiceTests Updates shape properties.
   Test for ShapesApi.PutSlideShapeInfo method with invalid storage
*/
func TestPutSlideShapeInfoInvalidstorage(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutSlideShapeInfo", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ShapesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "storage", r.Code, e)
}
