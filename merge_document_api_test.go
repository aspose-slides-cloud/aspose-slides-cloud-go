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

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method
*/
func TestPostPresentationMerge(t *testing.T) {
    request := createPostPresentationMergeRequest()
    e := initializeTest("PostPresentationMerge", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.MergeDocumentApi.PostPresentationMerge(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostPresentationMergeRequest() PostPresentationMergeRequest {
    var request PostPresentationMergeRequest
    request.name = createTestParamValue("PostPresentationMerge", "name", "string").(string)
    request.request = createTestParamValue("PostPresentationMerge", "request", "PresentationsMergeRequest").(IPresentationsMergeRequest)
    request.password = createTestParamValue("PostPresentationMerge", "password", "string").(string)
    request.storage = createTestParamValue("PostPresentationMerge", "storage", "string").(string)
    request.folder = createTestParamValue("PostPresentationMerge", "folder", "string").(string)
    return request
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method with invalid name
*/
func TestPostPresentationMergeInvalidname(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "name", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method with invalid request
*/
func TestPostPresentationMergeInvalidrequest(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.request = invalidizeTestParamValue(request.request, "request", "PresentationsMergeRequest").(IPresentationsMergeRequest)

    e := initializeTest("PostPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "request", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method with invalid password
*/
func TestPostPresentationMergeInvalidpassword(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "password", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method with invalid storage
*/
func TestPostPresentationMergeInvalidstorage(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "storage", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for MergeDocumentApi.PostPresentationMerge method with invalid folder
*/
func TestPostPresentationMergeInvalidfolder(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "folder", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method
*/
func TestPutPresentationMerge(t *testing.T) {
    request := createPutPresentationMergeRequest()
    e := initializeTest("PutPresentationMerge", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.MergeDocumentApi.PutPresentationMerge(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutPresentationMergeRequest() PutPresentationMergeRequest {
    var request PutPresentationMergeRequest
    request.name = createTestParamValue("PutPresentationMerge", "name", "string").(string)
    request.request = createTestParamValue("PutPresentationMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    request.password = createTestParamValue("PutPresentationMerge", "password", "string").(string)
    request.storage = createTestParamValue("PutPresentationMerge", "storage", "string").(string)
    request.folder = createTestParamValue("PutPresentationMerge", "folder", "string").(string)
    return request
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method with invalid name
*/
func TestPutPresentationMergeInvalidname(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "name", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method with invalid request
*/
func TestPutPresentationMergeInvalidrequest(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.request = invalidizeTestParamValue(request.request, "request", "OrderedMergeRequest").(IOrderedMergeRequest)

    e := initializeTest("PutPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "request", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method with invalid password
*/
func TestPutPresentationMergeInvalidpassword(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "password", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method with invalid storage
*/
func TestPutPresentationMergeInvalidstorage(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "storage", r.Code, e)
}

/* MergeDocumentApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for MergeDocumentApi.PutPresentationMerge method with invalid folder
*/
func TestPutPresentationMergeInvalidfolder(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().MergeDocumentApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "folder", r.Code, e)
}
