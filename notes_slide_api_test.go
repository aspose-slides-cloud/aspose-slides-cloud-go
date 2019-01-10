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

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method
*/
func TestDeleteNotesSlide(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    e := initializeTest("DeleteNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideApi.DeleteNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteNotesSlideRequest() DeleteNotesSlideRequest {
    var request DeleteNotesSlideRequest
    request.name = createTestParamValue("DeleteNotesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteNotesSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteNotesSlide", "password", "string").(string)
    request.folder = createTestParamValue("DeleteNotesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteNotesSlide", "storage", "string").(string)
    return request
}

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method with invalid name
*/
func TestDeleteNotesSlideInvalidname(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "name", r.Code, e)
}

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method with invalid slideIndex
*/
func TestDeleteNotesSlideInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "slideIndex", r.Code, e)
}

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method with invalid password
*/
func TestDeleteNotesSlideInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "password", r.Code, e)
}

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method with invalid folder
*/
func TestDeleteNotesSlideInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "folder", r.Code, e)
}

/* NotesSlideApiServiceTests Remove Notes Slide.
   Test for NotesSlideApi.DeleteNotesSlide method with invalid storage
*/
func TestDeleteNotesSlideInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "storage", r.Code, e)
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method
*/
func TestGetNotesSlide(t *testing.T) {
    request := createGetNotesSlideRequest()
    e := initializeTest("GetNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideApi.GetNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetNotesSlideRequest() GetNotesSlideRequest {
    var request GetNotesSlideRequest
    request.name = createTestParamValue("GetNotesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetNotesSlide", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlide", "storage", "string").(string)
    return request
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method with invalid name
*/
func TestGetNotesSlideInvalidname(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "name", r.Code, e)
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method with invalid slideIndex
*/
func TestGetNotesSlideInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "slideIndex", r.Code, e)
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method with invalid password
*/
func TestGetNotesSlideInvalidpassword(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "password", r.Code, e)
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method with invalid folder
*/
func TestGetNotesSlideInvalidfolder(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "folder", r.Code, e)
}

/* NotesSlideApiServiceTests Read Notes slide info.
   Test for NotesSlideApi.GetNotesSlide method with invalid storage
*/
func TestGetNotesSlideInvalidstorage(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "storage", r.Code, e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method
*/
func TestGetNotesSlideWithFormat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    e := initializeTest("GetNotesSlideWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideApi.GetNotesSlideWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetNotesSlideWithFormatRequest() GetNotesSlideWithFormatRequest {
    var request GetNotesSlideWithFormatRequest
    request.name = createTestParamValue("GetNotesSlideWithFormat", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetNotesSlideWithFormat", "slideIndex", "int32").(int32)
    request.format = createTestParamValue("GetNotesSlideWithFormat", "format", "string").(string)
    testwidth := createTestParamValue("GetNotesSlideWithFormat", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.width = new(int32)
        *request.width = v
    }
    testheight := createTestParamValue("GetNotesSlideWithFormat", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.height = new(int32)
        *request.height = v
    }
    request.password = createTestParamValue("GetNotesSlideWithFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetNotesSlideWithFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetNotesSlideWithFormat", "storage", "string").(string)
    return request
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid name
*/
func TestGetNotesSlideWithFormatInvalidname(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "name", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid slideIndex
*/
func TestGetNotesSlideWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "slideIndex", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid format
*/
func TestGetNotesSlideWithFormatInvalidformat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "format", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid width
*/
func TestGetNotesSlideWithFormatInvalidwidth(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.width = new(int32)
    *request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)

    e := initializeTest("GetNotesSlideWithFormat", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "width", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid height
*/
func TestGetNotesSlideWithFormatInvalidheight(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.height = new(int32)
    *request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)

    e := initializeTest("GetNotesSlideWithFormat", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "height", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid password
*/
func TestGetNotesSlideWithFormatInvalidpassword(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "password", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid folder
*/
func TestGetNotesSlideWithFormatInvalidfolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "folder", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Convert Notes Slide to the specified picture format.
   Test for NotesSlideApi.GetNotesSlideWithFormat method with invalid storage
*/
func TestGetNotesSlideWithFormatInvalidstorage(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().NotesSlideApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "storage", int32(r.StatusCode), e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method
*/
func TestPostAddNotesSlide(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    e := initializeTest("PostAddNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideApi.PostAddNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostAddNotesSlideRequest() PostAddNotesSlideRequest {
    var request PostAddNotesSlideRequest
    request.name = createTestParamValue("PostAddNotesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostAddNotesSlide", "slideIndex", "int32").(int32)
    request.dto = createTestParamValue("PostAddNotesSlide", "dto", "NotesSlide").(INotesSlide)
    request.password = createTestParamValue("PostAddNotesSlide", "password", "string").(string)
    request.folder = createTestParamValue("PostAddNotesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("PostAddNotesSlide", "storage", "string").(string)
    return request
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid name
*/
func TestPostAddNotesSlideInvalidname(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostAddNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "name", r.Code, e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid slideIndex
*/
func TestPostAddNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostAddNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "slideIndex", r.Code, e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid dto
*/
func TestPostAddNotesSlideInvaliddto(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "NotesSlide").(INotesSlide)

    e := initializeTest("PostAddNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "dto", r.Code, e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid password
*/
func TestPostAddNotesSlideInvalidpassword(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostAddNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "password", r.Code, e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid folder
*/
func TestPostAddNotesSlideInvalidfolder(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostAddNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "folder", r.Code, e)
}

/* NotesSlideApiServiceTests Add new Notes Slide.
   Test for NotesSlideApi.PostAddNotesSlide method with invalid storage
*/
func TestPostAddNotesSlideInvalidstorage(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostAddNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "storage", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method
*/
func TestPutUpdateNotesSlide(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    e := initializeTest("PutUpdateNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.NotesSlideApi.PutUpdateNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutUpdateNotesSlideRequest() PutUpdateNotesSlideRequest {
    var request PutUpdateNotesSlideRequest
    request.name = createTestParamValue("PutUpdateNotesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutUpdateNotesSlide", "slideIndex", "int32").(int32)
    request.dto = createTestParamValue("PutUpdateNotesSlide", "dto", "NotesSlide").(INotesSlide)
    request.password = createTestParamValue("PutUpdateNotesSlide", "password", "string").(string)
    request.folder = createTestParamValue("PutUpdateNotesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("PutUpdateNotesSlide", "storage", "string").(string)
    return request
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid name
*/
func TestPutUpdateNotesSlideInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "name", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid slideIndex
*/
func TestPutUpdateNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "slideIndex", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid dto
*/
func TestPutUpdateNotesSlideInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "NotesSlide").(INotesSlide)

    e := initializeTest("PutUpdateNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "dto", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid password
*/
func TestPutUpdateNotesSlideInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "password", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid folder
*/
func TestPutUpdateNotesSlideInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "folder", r.Code, e)
}

/* NotesSlideApiServiceTests Update Notes Slide properties.
   Test for NotesSlideApi.PutUpdateNotesSlide method with invalid storage
*/
func TestPutUpdateNotesSlideInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().NotesSlideApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "storage", r.Code, e)
}
