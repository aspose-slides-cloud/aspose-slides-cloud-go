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

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method
*/
func TestGetLayoutSlide(t *testing.T) {
    request := createGetLayoutSlideRequest()
    e := initializeTest("GetLayoutSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.LayoutSlidesApi.GetLayoutSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetLayoutSlideRequest() GetLayoutSlideRequest {
    var request GetLayoutSlideRequest
    request.name = createTestParamValue("GetLayoutSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetLayoutSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetLayoutSlide", "password", "string").(string)
    request.folder = createTestParamValue("GetLayoutSlide", "folder", "string").(string)
    request.storage = createTestParamValue("GetLayoutSlide", "storage", "string").(string)
    return request
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method with invalid name
*/
func TestGetLayoutSlideInvalidname(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "name", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method with invalid slideIndex
*/
func TestGetLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "slideIndex", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method with invalid password
*/
func TestGetLayoutSlideInvalidpassword(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "password", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method with invalid folder
*/
func TestGetLayoutSlideInvalidfolder(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "folder", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlide info.
   Test for LayoutSlidesApi.GetLayoutSlide method with invalid storage
*/
func TestGetLayoutSlideInvalidstorage(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "storage", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlides info.
   Test for LayoutSlidesApi.GetLayoutSlidesList method
*/
func TestGetLayoutSlidesList(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    e := initializeTest("GetLayoutSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.LayoutSlidesApi.GetLayoutSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetLayoutSlidesListRequest() GetLayoutSlidesListRequest {
    var request GetLayoutSlidesListRequest
    request.name = createTestParamValue("GetLayoutSlidesList", "name", "string").(string)
    request.password = createTestParamValue("GetLayoutSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("GetLayoutSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("GetLayoutSlidesList", "storage", "string").(string)
    return request
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlides info.
   Test for LayoutSlidesApi.GetLayoutSlidesList method with invalid name
*/
func TestGetLayoutSlidesListInvalidname(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetLayoutSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "name", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlides info.
   Test for LayoutSlidesApi.GetLayoutSlidesList method with invalid password
*/
func TestGetLayoutSlidesListInvalidpassword(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetLayoutSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "password", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlides info.
   Test for LayoutSlidesApi.GetLayoutSlidesList method with invalid folder
*/
func TestGetLayoutSlidesListInvalidfolder(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetLayoutSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "folder", r.Code, e)
}

/* LayoutSlidesApiServiceTests Read presentation layoutSlides info.
   Test for LayoutSlidesApi.GetLayoutSlidesList method with invalid storage
*/
func TestGetLayoutSlidesListInvalidstorage(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetLayoutSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "storage", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method
*/
func TestPostCopyLayoutSlideFromSourcePresentation(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostCopyLayoutSlideFromSourcePresentationRequest() PostCopyLayoutSlideFromSourcePresentationRequest {
    var request PostCopyLayoutSlideFromSourcePresentationRequest
    request.name = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "name", "string").(string)
    request.cloneFrom = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", "string").(string)
    request.cloneFromPosition = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", "int32").(int32)
    request.cloneFromPassword = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", "string").(string)
    request.cloneFromStorage = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", "string").(string)
    request.password = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "password", "string").(string)
    request.folder = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "folder", "string").(string)
    request.storage = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "storage", "string").(string)
    return request
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidname(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "name", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFrom = invalidizeTestParamValue(request.cloneFrom, "cloneFrom", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromPosition = invalidizeTestParamValue(request.cloneFromPosition, "cloneFromPosition", "int32").(int32)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromPassword = invalidizeTestParamValue(request.cloneFromPassword, "cloneFromPassword", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromStorage = invalidizeTestParamValue(request.cloneFromStorage, "cloneFromStorage", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "password", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "folder", r.Code, e)
}

/* LayoutSlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "storage", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method
*/
func TestPutLayoutSlide(t *testing.T) {
    request := createPutLayoutSlideRequest()
    e := initializeTest("PutLayoutSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.LayoutSlidesApi.PutLayoutSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutLayoutSlideRequest() PutLayoutSlideRequest {
    var request PutLayoutSlideRequest
    request.name = createTestParamValue("PutLayoutSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutLayoutSlide", "slideIndex", "int32").(int32)
    request.slideDto = createTestParamValue("PutLayoutSlide", "slideDto", "LayoutSlide").(ILayoutSlide)
    request.password = createTestParamValue("PutLayoutSlide", "password", "string").(string)
    request.folder = createTestParamValue("PutLayoutSlide", "folder", "string").(string)
    request.storage = createTestParamValue("PutLayoutSlide", "storage", "string").(string)
    return request
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid name
*/
func TestPutLayoutSlideInvalidname(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "name", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid slideIndex
*/
func TestPutLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "slideIndex", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid slideDto
*/
func TestPutLayoutSlideInvalidslideDto(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.slideDto = invalidizeTestParamValue(request.slideDto, "slideDto", "LayoutSlide").(ILayoutSlide)
    e := initializeTest("PutLayoutSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "slideDto", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid password
*/
func TestPutLayoutSlideInvalidpassword(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "password", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid folder
*/
func TestPutLayoutSlideInvalidfolder(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "folder", r.Code, e)
}

/* LayoutSlidesApiServiceTests Update a layoutSlide.
   Test for LayoutSlidesApi.PutLayoutSlide method with invalid storage
*/
func TestPutLayoutSlideInvalidstorage(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().LayoutSlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "storage", r.Code, e)
}
