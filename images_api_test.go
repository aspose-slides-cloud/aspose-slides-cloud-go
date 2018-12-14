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

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method
*/
func TestGetSlidesImageWithFormat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    e := initializeTest("GetSlidesImageWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ImagesApi.GetSlidesImageWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetSlidesImageWithFormatRequest() GetSlidesImageWithFormatRequest {
    var request GetSlidesImageWithFormatRequest
    request.name = createTestParamValue("GetSlidesImageWithFormat", "name", "string").(string)
    request.index = createTestParamValue("GetSlidesImageWithFormat", "index", "int32").(int32)
    request.format = createTestParamValue("GetSlidesImageWithFormat", "format", "string").(string)
    request.password = createTestParamValue("GetSlidesImageWithFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesImageWithFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesImageWithFormat", "storage", "string").(string)
    return request
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid name
*/
func TestGetSlidesImageWithFormatInvalidname(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesImageWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "name", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid index
*/
func TestGetSlidesImageWithFormatInvalidindex(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.index = invalidizeTestParamValue(request.index, "index", "int32").(int32)
    e := initializeTest("GetSlidesImageWithFormat", "index", request.index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "index", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid format
*/
func TestGetSlidesImageWithFormatInvalidformat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("GetSlidesImageWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "format", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid password
*/
func TestGetSlidesImageWithFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesImageWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "password", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid folder
*/
func TestGetSlidesImageWithFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesImageWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "folder", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Gets image in specified format.
   Test for ImagesApi.GetSlidesImageWithFormat method with invalid storage
*/
func TestGetSlidesImageWithFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesImageWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().ImagesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "storage", int32(r.StatusCode), e)
}

/* ImagesApiServiceTests Read presentation images info.
   Test for ImagesApi.GetSlidesImages method
*/
func TestGetSlidesImages(t *testing.T) {
    request := createGetSlidesImagesRequest()
    e := initializeTest("GetSlidesImages", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ImagesApi.GetSlidesImages(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesImagesRequest() GetSlidesImagesRequest {
    var request GetSlidesImagesRequest
    request.name = createTestParamValue("GetSlidesImages", "name", "string").(string)
    request.password = createTestParamValue("GetSlidesImages", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesImages", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesImages", "storage", "string").(string)
    return request
}

/* ImagesApiServiceTests Read presentation images info.
   Test for ImagesApi.GetSlidesImages method with invalid name
*/
func TestGetSlidesImagesInvalidname(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "name", r.Code, e)
}

/* ImagesApiServiceTests Read presentation images info.
   Test for ImagesApi.GetSlidesImages method with invalid password
*/
func TestGetSlidesImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "password", r.Code, e)
}

/* ImagesApiServiceTests Read presentation images info.
   Test for ImagesApi.GetSlidesImages method with invalid folder
*/
func TestGetSlidesImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "folder", r.Code, e)
}

/* ImagesApiServiceTests Read presentation images info.
   Test for ImagesApi.GetSlidesImages method with invalid storage
*/
func TestGetSlidesImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "storage", r.Code, e)
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method
*/
func TestGetSlidesSlideImages(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    e := initializeTest("GetSlidesSlideImages", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ImagesApi.GetSlidesSlideImages(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlideImagesRequest() GetSlidesSlideImagesRequest {
    var request GetSlidesSlideImagesRequest
    request.name = createTestParamValue("GetSlidesSlideImages", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideImages", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlideImages", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideImages", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideImages", "storage", "string").(string)
    return request
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method with invalid name
*/
func TestGetSlidesSlideImagesInvalidname(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlideImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "name", r.Code, e)
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method with invalid slideIndex
*/
func TestGetSlidesSlideImagesInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesSlideImages", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "slideIndex", r.Code, e)
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method with invalid password
*/
func TestGetSlidesSlideImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlideImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "password", r.Code, e)
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method with invalid folder
*/
func TestGetSlidesSlideImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlideImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "folder", r.Code, e)
}

/* ImagesApiServiceTests Read slide images info.
   Test for ImagesApi.GetSlidesSlideImages method with invalid storage
*/
func TestGetSlidesSlideImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlideImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ImagesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "storage", r.Code, e)
}
