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

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method
*/
func TestGetMasterSlide(t *testing.T) {
    request := createGetMasterSlideRequest()
    e := initializeTest("GetMasterSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.MasterSlidesApi.GetMasterSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetMasterSlideRequest() GetMasterSlideRequest {
    var request GetMasterSlideRequest
    request.name = createTestParamValue("GetMasterSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetMasterSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetMasterSlide", "password", "string").(string)
    request.folder = createTestParamValue("GetMasterSlide", "folder", "string").(string)
    request.storage = createTestParamValue("GetMasterSlide", "storage", "string").(string)
    return request
}

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method with invalid name
*/
func TestGetMasterSlideInvalidname(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetMasterSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "name", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method with invalid slideIndex
*/
func TestGetMasterSlideInvalidslideIndex(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetMasterSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "slideIndex", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method with invalid password
*/
func TestGetMasterSlideInvalidpassword(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetMasterSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "password", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method with invalid folder
*/
func TestGetMasterSlideInvalidfolder(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetMasterSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "folder", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlide info.
   Test for MasterSlidesApi.GetMasterSlide method with invalid storage
*/
func TestGetMasterSlideInvalidstorage(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetMasterSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "storage", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlides info.
   Test for MasterSlidesApi.GetMasterSlidesList method
*/
func TestGetMasterSlidesList(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    e := initializeTest("GetMasterSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.MasterSlidesApi.GetMasterSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetMasterSlidesListRequest() GetMasterSlidesListRequest {
    var request GetMasterSlidesListRequest
    request.name = createTestParamValue("GetMasterSlidesList", "name", "string").(string)
    request.password = createTestParamValue("GetMasterSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("GetMasterSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("GetMasterSlidesList", "storage", "string").(string)
    return request
}

/* MasterSlidesApiServiceTests Read presentation masterSlides info.
   Test for MasterSlidesApi.GetMasterSlidesList method with invalid name
*/
func TestGetMasterSlidesListInvalidname(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetMasterSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "name", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlides info.
   Test for MasterSlidesApi.GetMasterSlidesList method with invalid password
*/
func TestGetMasterSlidesListInvalidpassword(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetMasterSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "password", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlides info.
   Test for MasterSlidesApi.GetMasterSlidesList method with invalid folder
*/
func TestGetMasterSlidesListInvalidfolder(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetMasterSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "folder", r.Code, e)
}

/* MasterSlidesApiServiceTests Read presentation masterSlides info.
   Test for MasterSlidesApi.GetMasterSlidesList method with invalid storage
*/
func TestGetMasterSlidesListInvalidstorage(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetMasterSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "storage", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method
*/
func TestPostCopyMasterSlideFromSourcePresentation(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostCopyMasterSlideFromSourcePresentationRequest() PostCopyMasterSlideFromSourcePresentationRequest {
    var request PostCopyMasterSlideFromSourcePresentationRequest
    request.name = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "name", "string").(string)
    request.cloneFrom = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", "string").(string)
    request.cloneFromPosition = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", "int32").(int32)
    request.cloneFromPassword = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", "string").(string)
    request.cloneFromStorage = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", "string").(string)
    request.applyToAll = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "applyToAll", "bool").(bool)
    request.password = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "password", "string").(string)
    request.folder = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "folder", "string").(string)
    request.storage = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "storage", "string").(string)
    return request
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidname(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "name", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFrom = invalidizeTestParamValue(request.cloneFrom, "cloneFrom", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromPosition = invalidizeTestParamValue(request.cloneFromPosition, "cloneFromPosition", "int32").(int32)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromPassword = invalidizeTestParamValue(request.cloneFromPassword, "cloneFromPassword", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromStorage = invalidizeTestParamValue(request.cloneFromStorage, "cloneFromStorage", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid applyToAll
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidapplyToAll(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.applyToAll = invalidizeTestParamValue(request.applyToAll, "applyToAll", "bool").(bool)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.applyToAll)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "password", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "folder", r.Code, e)
}

/* MasterSlidesApiServiceTests Copy masterSlide from source presentation.
   Test for MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MasterSlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "storage", r.Code, e)
}
