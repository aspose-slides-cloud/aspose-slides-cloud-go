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

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method
*/
func TestGetSlidesPresentationTextItems(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    e := initializeTest("GetSlidesPresentationTextItems", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.TextApi.GetSlidesPresentationTextItems(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesPresentationTextItemsRequest() GetSlidesPresentationTextItemsRequest {
    var request GetSlidesPresentationTextItemsRequest
    request.name = createTestParamValue("GetSlidesPresentationTextItems", "name", "string").(string)
    request.withEmpty = createTestParamValue("GetSlidesPresentationTextItems", "withEmpty", "bool").(bool)
    request.password = createTestParamValue("GetSlidesPresentationTextItems", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesPresentationTextItems", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesPresentationTextItems", "storage", "string").(string)
    return request
}

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method with invalid name
*/
func TestGetSlidesPresentationTextItemsInvalidname(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesPresentationTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "name", r.Code, e)
}

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method with invalid withEmpty
*/
func TestGetSlidesPresentationTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.withEmpty = invalidizeTestParamValue(request.withEmpty, "withEmpty", "bool").(bool)
    e := initializeTest("GetSlidesPresentationTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "withEmpty", r.Code, e)
}

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method with invalid password
*/
func TestGetSlidesPresentationTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesPresentationTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "password", r.Code, e)
}

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method with invalid folder
*/
func TestGetSlidesPresentationTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesPresentationTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "folder", r.Code, e)
}

/* TextApiServiceTests Extract presentation text items.
   Test for TextApi.GetSlidesPresentationTextItems method with invalid storage
*/
func TestGetSlidesPresentationTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesPresentationTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "storage", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method
*/
func TestGetSlidesSlideTextItems(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    e := initializeTest("GetSlidesSlideTextItems", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.TextApi.GetSlidesSlideTextItems(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlideTextItemsRequest() GetSlidesSlideTextItemsRequest {
    var request GetSlidesSlideTextItemsRequest
    request.name = createTestParamValue("GetSlidesSlideTextItems", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideTextItems", "slideIndex", "int32").(int32)
    request.withEmpty = createTestParamValue("GetSlidesSlideTextItems", "withEmpty", "bool").(bool)
    request.password = createTestParamValue("GetSlidesSlideTextItems", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideTextItems", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideTextItems", "storage", "string").(string)
    return request
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid name
*/
func TestGetSlidesSlideTextItemsInvalidname(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlideTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "name", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid slideIndex
*/
func TestGetSlidesSlideTextItemsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesSlideTextItems", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "slideIndex", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid withEmpty
*/
func TestGetSlidesSlideTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.withEmpty = invalidizeTestParamValue(request.withEmpty, "withEmpty", "bool").(bool)
    e := initializeTest("GetSlidesSlideTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "withEmpty", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid password
*/
func TestGetSlidesSlideTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlideTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "password", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid folder
*/
func TestGetSlidesSlideTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlideTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "folder", r.Code, e)
}

/* TextApiServiceTests Extract slide text items.
   Test for TextApi.GetSlidesSlideTextItems method with invalid storage
*/
func TestGetSlidesSlideTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlideTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "storage", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method
*/
func TestPostSlidesPresentationReplaceText(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    e := initializeTest("PostSlidesPresentationReplaceText", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.TextApi.PostSlidesPresentationReplaceText(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostSlidesPresentationReplaceTextRequest() PostSlidesPresentationReplaceTextRequest {
    var request PostSlidesPresentationReplaceTextRequest
    request.name = createTestParamValue("PostSlidesPresentationReplaceText", "name", "string").(string)
    request.oldValue = createTestParamValue("PostSlidesPresentationReplaceText", "oldValue", "string").(string)
    request.newValue = createTestParamValue("PostSlidesPresentationReplaceText", "newValue", "string").(string)
    request.ignoreCase = createTestParamValue("PostSlidesPresentationReplaceText", "ignoreCase", "bool").(bool)
    request.password = createTestParamValue("PostSlidesPresentationReplaceText", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesPresentationReplaceText", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesPresentationReplaceText", "storage", "string").(string)
    return request
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid name
*/
func TestPostSlidesPresentationReplaceTextInvalidname(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "name", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid oldValue
*/
func TestPostSlidesPresentationReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.oldValue = invalidizeTestParamValue(request.oldValue, "oldValue", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "oldValue", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid newValue
*/
func TestPostSlidesPresentationReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.newValue = invalidizeTestParamValue(request.newValue, "newValue", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "newValue", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid ignoreCase
*/
func TestPostSlidesPresentationReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.ignoreCase = invalidizeTestParamValue(request.ignoreCase, "ignoreCase", "bool").(bool)
    e := initializeTest("PostSlidesPresentationReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "ignoreCase", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid password
*/
func TestPostSlidesPresentationReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "password", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid folder
*/
func TestPostSlidesPresentationReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "folder", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesPresentationReplaceText method with invalid storage
*/
func TestPostSlidesPresentationReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostSlidesPresentationReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "storage", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method
*/
func TestPostSlidesSlideReplaceText(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    e := initializeTest("PostSlidesSlideReplaceText", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.TextApi.PostSlidesSlideReplaceText(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostSlidesSlideReplaceTextRequest() PostSlidesSlideReplaceTextRequest {
    var request PostSlidesSlideReplaceTextRequest
    request.name = createTestParamValue("PostSlidesSlideReplaceText", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlidesSlideReplaceText", "slideIndex", "int32").(int32)
    request.oldValue = createTestParamValue("PostSlidesSlideReplaceText", "oldValue", "string").(string)
    request.newValue = createTestParamValue("PostSlidesSlideReplaceText", "newValue", "string").(string)
    request.ignoreCase = createTestParamValue("PostSlidesSlideReplaceText", "ignoreCase", "bool").(bool)
    request.password = createTestParamValue("PostSlidesSlideReplaceText", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesSlideReplaceText", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesSlideReplaceText", "storage", "string").(string)
    return request
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid name
*/
func TestPostSlidesSlideReplaceTextInvalidname(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "name", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid slideIndex
*/
func TestPostSlidesSlideReplaceTextInvalidslideIndex(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostSlidesSlideReplaceText", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "slideIndex", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid oldValue
*/
func TestPostSlidesSlideReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.oldValue = invalidizeTestParamValue(request.oldValue, "oldValue", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "oldValue", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid newValue
*/
func TestPostSlidesSlideReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.newValue = invalidizeTestParamValue(request.newValue, "newValue", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "newValue", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid ignoreCase
*/
func TestPostSlidesSlideReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.ignoreCase = invalidizeTestParamValue(request.ignoreCase, "ignoreCase", "bool").(bool)
    e := initializeTest("PostSlidesSlideReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "ignoreCase", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid password
*/
func TestPostSlidesSlideReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "password", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid folder
*/
func TestPostSlidesSlideReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "folder", r.Code, e)
}

/* TextApiServiceTests Replace text with a new value.
   Test for TextApi.PostSlidesSlideReplaceText method with invalid storage
*/
func TestPostSlidesSlideReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostSlidesSlideReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().TextApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "storage", r.Code, e)
}
