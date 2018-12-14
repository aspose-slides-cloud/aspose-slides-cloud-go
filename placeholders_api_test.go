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

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method
*/
func TestGetSlidesPlaceholder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    e := initializeTest("GetSlidesPlaceholder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PlaceholdersApi.GetSlidesPlaceholder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesPlaceholderRequest() GetSlidesPlaceholderRequest {
    var request GetSlidesPlaceholderRequest
    request.name = createTestParamValue("GetSlidesPlaceholder", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesPlaceholder", "slideIndex", "int32").(int32)
    request.placeholderIndex = createTestParamValue("GetSlidesPlaceholder", "placeholderIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesPlaceholder", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesPlaceholder", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesPlaceholder", "storage", "string").(string)
    return request
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid name
*/
func TestGetSlidesPlaceholderInvalidname(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesPlaceholder", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "name", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid slideIndex
*/
func TestGetSlidesPlaceholderInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesPlaceholder", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "slideIndex", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid placeholderIndex
*/
func TestGetSlidesPlaceholderInvalidplaceholderIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.placeholderIndex = invalidizeTestParamValue(request.placeholderIndex, "placeholderIndex", "int32").(int32)
    e := initializeTest("GetSlidesPlaceholder", "placeholderIndex", request.placeholderIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "placeholderIndex", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid password
*/
func TestGetSlidesPlaceholderInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesPlaceholder", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "password", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid folder
*/
func TestGetSlidesPlaceholderInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesPlaceholder", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "folder", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholder info.
   Test for PlaceholdersApi.GetSlidesPlaceholder method with invalid storage
*/
func TestGetSlidesPlaceholderInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesPlaceholder", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "storage", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method
*/
func TestGetSlidesPlaceholders(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    e := initializeTest("GetSlidesPlaceholders", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PlaceholdersApi.GetSlidesPlaceholders(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesPlaceholdersRequest() GetSlidesPlaceholdersRequest {
    var request GetSlidesPlaceholdersRequest
    request.name = createTestParamValue("GetSlidesPlaceholders", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesPlaceholders", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesPlaceholders", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesPlaceholders", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesPlaceholders", "storage", "string").(string)
    return request
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method with invalid name
*/
func TestGetSlidesPlaceholdersInvalidname(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesPlaceholders", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "name", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method with invalid slideIndex
*/
func TestGetSlidesPlaceholdersInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesPlaceholders", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "slideIndex", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method with invalid password
*/
func TestGetSlidesPlaceholdersInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesPlaceholders", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "password", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method with invalid folder
*/
func TestGetSlidesPlaceholdersInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesPlaceholders", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "folder", r.Code, e)
}

/* PlaceholdersApiServiceTests Read slide placeholders info.
   Test for PlaceholdersApi.GetSlidesPlaceholders method with invalid storage
*/
func TestGetSlidesPlaceholdersInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesPlaceholders", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PlaceholdersApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "storage", r.Code, e)
}
