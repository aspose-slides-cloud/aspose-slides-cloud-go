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

/* DocumentApiServiceTests Get API info.
   Test for DocumentApi.GetSlidesApiInfo method
*/
func TestGetSlidesApiInfo(t *testing.T) {
    e := initializeTest("GetSlidesApiInfo", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.GetSlidesApiInfo()
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

/* DocumentApiServiceTests Read presentation info.
   Test for DocumentApi.GetSlidesDocument method
*/
func TestGetSlidesDocument(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    e := initializeTest("GetSlidesDocument", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.GetSlidesDocument(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesDocumentRequest() GetSlidesDocumentRequest {
    var request GetSlidesDocumentRequest
    request.name = createTestParamValue("GetSlidesDocument", "name", "string").(string)
    request.password = createTestParamValue("GetSlidesDocument", "password", "string").(string)
    request.storage = createTestParamValue("GetSlidesDocument", "storage", "string").(string)
    request.folder = createTestParamValue("GetSlidesDocument", "folder", "string").(string)
    return request
}

/* DocumentApiServiceTests Read presentation info.
   Test for DocumentApi.GetSlidesDocument method with invalid name
*/
func TestGetSlidesDocumentInvalidname(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "name", r.Code, e)
}

/* DocumentApiServiceTests Read presentation info.
   Test for DocumentApi.GetSlidesDocument method with invalid password
*/
func TestGetSlidesDocumentInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "password", r.Code, e)
}

/* DocumentApiServiceTests Read presentation info.
   Test for DocumentApi.GetSlidesDocument method with invalid storage
*/
func TestGetSlidesDocumentInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "storage", r.Code, e)
}

/* DocumentApiServiceTests Read presentation info.
   Test for DocumentApi.GetSlidesDocument method with invalid folder
*/
func TestGetSlidesDocumentInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "folder", r.Code, e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method
*/
func TestGetSlidesDocumentWithFormat(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    e := initializeTest("GetSlidesDocumentWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.GetSlidesDocumentWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetSlidesDocumentWithFormatRequest() GetSlidesDocumentWithFormatRequest {
    var request GetSlidesDocumentWithFormatRequest
    request.name = createTestParamValue("GetSlidesDocumentWithFormat", "name", "string").(string)
    request.format = createTestParamValue("GetSlidesDocumentWithFormat", "format", "string").(string)
    testjpegQuality := createTestParamValue("GetSlidesDocumentWithFormat", "jpegQuality", "int32")
    switch v := testjpegQuality.(type) { 
    case int32:
        request.jpegQuality = new(int32)
        *request.jpegQuality = v
    }
    request.password = createTestParamValue("GetSlidesDocumentWithFormat", "password", "string").(string)
    request.storage = createTestParamValue("GetSlidesDocumentWithFormat", "storage", "string").(string)
    request.folder = createTestParamValue("GetSlidesDocumentWithFormat", "folder", "string").(string)
    request.outPath = createTestParamValue("GetSlidesDocumentWithFormat", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("GetSlidesDocumentWithFormat", "fontsFolder", "string").(string)
    return request
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid name
*/
func TestGetSlidesDocumentWithFormatInvalidname(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "name", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid format
*/
func TestGetSlidesDocumentWithFormatInvalidformat(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "format", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid jpegQuality
*/
func TestGetSlidesDocumentWithFormatInvalidjpegQuality(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.jpegQuality = new(int32)
    *request.jpegQuality = invalidizeTestParamValue(request.jpegQuality, "jpegQuality", "int32").(int32)

    e := initializeTest("GetSlidesDocumentWithFormat", "jpegQuality", request.jpegQuality)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "jpegQuality", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid password
*/
func TestGetSlidesDocumentWithFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "password", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid storage
*/
func TestGetSlidesDocumentWithFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "storage", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid folder
*/
func TestGetSlidesDocumentWithFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "folder", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid outPath
*/
func TestGetSlidesDocumentWithFormatInvalidoutPath(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "outPath", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Export presentation to some format.
   Test for DocumentApi.GetSlidesDocumentWithFormat method with invalid fontsFolder
*/
func TestGetSlidesDocumentWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetSlidesDocumentWithFormatRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("GetSlidesDocumentWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.GetSlidesDocumentWithFormat(request)
    assertError(t, "GetSlidesDocumentWithFormat", "fontsFolder", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method
*/
func TestPostSlidesDocument(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    e := initializeTest("PostSlidesDocument", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PostSlidesDocument(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostSlidesDocumentRequest() PostSlidesDocumentRequest {
    var request PostSlidesDocumentRequest
    request.name = createTestParamValue("PostSlidesDocument", "name", "string").(string)
    request.data = createTestParamValue("PostSlidesDocument", "data", "string").(string)
    request.templatePath = createTestParamValue("PostSlidesDocument", "templatePath", "string").(string)
    request.templateStorage = createTestParamValue("PostSlidesDocument", "templateStorage", "string").(string)
    testisImageDataEmbedded := createTestParamValue("PostSlidesDocument", "isImageDataEmbedded", "bool")
    switch v := testisImageDataEmbedded.(type) { 
    case bool:
        request.isImageDataEmbedded = new(bool)
        *request.isImageDataEmbedded = v
    }
    request.password = createTestParamValue("PostSlidesDocument", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesDocument", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesDocument", "folder", "string").(string)
    return request
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid name
*/
func TestPostSlidesDocumentInvalidname(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "name", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid data
*/
func TestPostSlidesDocumentInvaliddata(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.data = invalidizeTestParamValue(request.data, "data", "string").(string)

    e := initializeTest("PostSlidesDocument", "data", request.data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "data", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid templatePath
*/
func TestPostSlidesDocumentInvalidtemplatePath(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.templatePath = invalidizeTestParamValue(request.templatePath, "templatePath", "string").(string)

    e := initializeTest("PostSlidesDocument", "templatePath", request.templatePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "templatePath", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid templateStorage
*/
func TestPostSlidesDocumentInvalidtemplateStorage(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.templateStorage = invalidizeTestParamValue(request.templateStorage, "templateStorage", "string").(string)

    e := initializeTest("PostSlidesDocument", "templateStorage", request.templateStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "templateStorage", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid isImageDataEmbedded
*/
func TestPostSlidesDocumentInvalidisImageDataEmbedded(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.isImageDataEmbedded = new(bool)
    *request.isImageDataEmbedded = invalidizeTestParamValue(request.isImageDataEmbedded, "isImageDataEmbedded", "bool").(bool)

    e := initializeTest("PostSlidesDocument", "isImageDataEmbedded", request.isImageDataEmbedded)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "isImageDataEmbedded", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid password
*/
func TestPostSlidesDocumentInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "password", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid storage
*/
func TestPostSlidesDocumentInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "storage", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PostSlidesDocument method with invalid folder
*/
func TestPostSlidesDocumentInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "folder", r.Code, e)
}

/* DocumentApiServiceTests Performs slides pipeline. Http-request contains pipeline DTO.
   Test for DocumentApi.PostSlidesPipeline method
*/
func TestPostSlidesPipeline(t *testing.T) {
    request := createPostSlidesPipelineRequest()
    e := initializeTest("PostSlidesPipeline", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PostSlidesPipeline(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostSlidesPipelineRequest() PostSlidesPipelineRequest {
    var request PostSlidesPipelineRequest
    request.pipeline = createTestParamValue("PostSlidesPipeline", "pipeline", "Pipeline").(IPipeline)
    return request
}

/* DocumentApiServiceTests Performs slides pipeline. Http-request contains pipeline DTO.
   Test for DocumentApi.PostSlidesPipeline method with invalid pipeline
*/
func TestPostSlidesPipelineInvalidpipeline(t *testing.T) {
    request := createPostSlidesPipelineRequest()
    request.pipeline = invalidizeTestParamValue(request.pipeline, "pipeline", "Pipeline").(IPipeline)

    e := initializeTest("PostSlidesPipeline", "pipeline", request.pipeline)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesPipeline(request)
    assertError(t, "PostSlidesPipeline", "pipeline", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method
*/
func TestPostSlidesSaveAs(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    e := initializeTest("PostSlidesSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PostSlidesSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostSlidesSaveAsRequest() PostSlidesSaveAsRequest {
    var request PostSlidesSaveAsRequest
    request.name = createTestParamValue("PostSlidesSaveAs", "name", "string").(string)
    request.format = createTestParamValue("PostSlidesSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostSlidesSaveAs", "options", "ExportOptions").(IExportOptions)
    request.password = createTestParamValue("PostSlidesSaveAs", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesSaveAs", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesSaveAs", "folder", "string").(string)
    request.outPath = createTestParamValue("PostSlidesSaveAs", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("PostSlidesSaveAs", "fontsFolder", "string").(string)
    return request
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid name
*/
func TestPostSlidesSaveAsInvalidname(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "name", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid format
*/
func TestPostSlidesSaveAsInvalidformat(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "format", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid options
*/
func TestPostSlidesSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PostSlidesSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "options", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid password
*/
func TestPostSlidesSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "password", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid storage
*/
func TestPostSlidesSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "storage", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid folder
*/
func TestPostSlidesSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "folder", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid outPath
*/
func TestPostSlidesSaveAsInvalidoutPath(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "outPath", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Saves presentation with options
   Test for DocumentApi.PostSlidesSaveAs method with invalid fontsFolder
*/
func TestPostSlidesSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "fontsFolder", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method
*/
func TestPostSlidesSplit(t *testing.T) {
    request := createPostSlidesSplitRequest()
    e := initializeTest("PostSlidesSplit", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PostSlidesSplit(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostSlidesSplitRequest() PostSlidesSplitRequest {
    var request PostSlidesSplitRequest
    request.name = createTestParamValue("PostSlidesSplit", "name", "string").(string)
    request.options = createTestParamValue("PostSlidesSplit", "options", "ExportOptions").(IExportOptions)
    request.format = createTestParamValue("PostSlidesSplit", "format", "string").(string)
    testwidth := createTestParamValue("PostSlidesSplit", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.width = new(int32)
        *request.width = v
    }
    testheight := createTestParamValue("PostSlidesSplit", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.height = new(int32)
        *request.height = v
    }
    testto := createTestParamValue("PostSlidesSplit", "to", "int32")
    switch v := testto.(type) { 
    case int32:
        request.to = new(int32)
        *request.to = v
    }
    testfrom := createTestParamValue("PostSlidesSplit", "from", "int32")
    switch v := testfrom.(type) { 
    case int32:
        request.from = new(int32)
        *request.from = v
    }
    request.destFolder = createTestParamValue("PostSlidesSplit", "destFolder", "string").(string)
    request.password = createTestParamValue("PostSlidesSplit", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesSplit", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesSplit", "folder", "string").(string)
    return request
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid name
*/
func TestPostSlidesSplitInvalidname(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSplit", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "name", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid options
*/
func TestPostSlidesSplitInvalidoptions(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PostSlidesSplit", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "options", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid format
*/
func TestPostSlidesSplitInvalidformat(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)

    e := initializeTest("PostSlidesSplit", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "format", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid width
*/
func TestPostSlidesSplitInvalidwidth(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.width = new(int32)
    *request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)

    e := initializeTest("PostSlidesSplit", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "width", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid height
*/
func TestPostSlidesSplitInvalidheight(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.height = new(int32)
    *request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)

    e := initializeTest("PostSlidesSplit", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "height", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid to
*/
func TestPostSlidesSplitInvalidto(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.to = new(int32)
    *request.to = invalidizeTestParamValue(request.to, "to", "int32").(int32)

    e := initializeTest("PostSlidesSplit", "to", request.to)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "to", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid from
*/
func TestPostSlidesSplitInvalidfrom(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.from = new(int32)
    *request.from = invalidizeTestParamValue(request.from, "from", "int32").(int32)

    e := initializeTest("PostSlidesSplit", "from", request.from)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "from", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid destFolder
*/
func TestPostSlidesSplitInvaliddestFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.destFolder = invalidizeTestParamValue(request.destFolder, "destFolder", "string").(string)

    e := initializeTest("PostSlidesSplit", "destFolder", request.destFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "destFolder", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid password
*/
func TestPostSlidesSplitInvalidpassword(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSplit", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "password", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid storage
*/
func TestPostSlidesSplitInvalidstorage(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSplit", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "storage", r.Code, e)
}

/* DocumentApiServiceTests Splitting presentations. Create one image per slide.
   Test for DocumentApi.PostSlidesSplit method with invalid folder
*/
func TestPostSlidesSplitInvalidfolder(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSplit", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "folder", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method
*/
func TestPutNewPresentation(t *testing.T) {
    request := createPutNewPresentationRequest()
    e := initializeTest("PutNewPresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PutNewPresentation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutNewPresentationRequest() PutNewPresentationRequest {
    var request PutNewPresentationRequest
    request.name = createTestParamValue("PutNewPresentation", "name", "string").(string)
    request.stream = createTestParamValue("PutNewPresentation", "stream", "[]byte").([]byte)
    request.templatePath = createTestParamValue("PutNewPresentation", "templatePath", "string").(string)
    request.templatePassword = createTestParamValue("PutNewPresentation", "templatePassword", "string").(string)
    request.templateStorage = createTestParamValue("PutNewPresentation", "templateStorage", "string").(string)
    request.password = createTestParamValue("PutNewPresentation", "password", "string").(string)
    request.storage = createTestParamValue("PutNewPresentation", "storage", "string").(string)
    request.folder = createTestParamValue("PutNewPresentation", "folder", "string").(string)
    return request
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid name
*/
func TestPutNewPresentationInvalidname(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutNewPresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "name", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid stream
*/
func TestPutNewPresentationInvalidstream(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.stream = invalidizeTestParamValue(request.stream, "stream", "[]byte").([]byte)

    e := initializeTest("PutNewPresentation", "stream", request.stream)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "stream", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid templatePath
*/
func TestPutNewPresentationInvalidtemplatePath(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.templatePath = invalidizeTestParamValue(request.templatePath, "templatePath", "string").(string)

    e := initializeTest("PutNewPresentation", "templatePath", request.templatePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "templatePath", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid templatePassword
*/
func TestPutNewPresentationInvalidtemplatePassword(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.templatePassword = invalidizeTestParamValue(request.templatePassword, "templatePassword", "string").(string)

    e := initializeTest("PutNewPresentation", "templatePassword", request.templatePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "templatePassword", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid templateStorage
*/
func TestPutNewPresentationInvalidtemplateStorage(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.templateStorage = invalidizeTestParamValue(request.templateStorage, "templateStorage", "string").(string)

    e := initializeTest("PutNewPresentation", "templateStorage", request.templateStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "templateStorage", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid password
*/
func TestPutNewPresentationInvalidpassword(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutNewPresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "password", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid storage
*/
func TestPutNewPresentationInvalidstorage(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutNewPresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "storage", r.Code, e)
}

/* DocumentApiServiceTests Create presentation 
   Test for DocumentApi.PutNewPresentation method with invalid folder
*/
func TestPutNewPresentationInvalidfolder(t *testing.T) {
    request := createPutNewPresentationRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutNewPresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutNewPresentation(request)
    assertError(t, "PutNewPresentation", "folder", r.Code, e)
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method
*/
func TestPutSlidesConvert(t *testing.T) {
    request := createPutSlidesConvertRequest()
    e := initializeTest("PutSlidesConvert", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PutSlidesConvert(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPutSlidesConvertRequest() PutSlidesConvertRequest {
    var request PutSlidesConvertRequest
    request.format = createTestParamValue("PutSlidesConvert", "format", "string").(string)
    request.document = createTestParamValue("PutSlidesConvert", "document", "[]byte").([]byte)
    request.password = createTestParamValue("PutSlidesConvert", "password", "string").(string)
    request.outPath = createTestParamValue("PutSlidesConvert", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("PutSlidesConvert", "fontsFolder", "string").(string)
    return request
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method with invalid format
*/
func TestPutSlidesConvertInvalidformat(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)

    e := initializeTest("PutSlidesConvert", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "format", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method with invalid document
*/
func TestPutSlidesConvertInvaliddocument(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.document = invalidizeTestParamValue(request.document, "document", "[]byte").([]byte)

    e := initializeTest("PutSlidesConvert", "document", request.document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "document", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method with invalid password
*/
func TestPutSlidesConvertInvalidpassword(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesConvert", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "password", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method with invalid outPath
*/
func TestPutSlidesConvertInvalidoutPath(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutSlidesConvert", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "outPath", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Convert presentation from request content to format specified.
   Test for DocumentApi.PutSlidesConvert method with invalid fontsFolder
*/
func TestPutSlidesConvertInvalidfontsFolder(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutSlidesConvert", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().DocumentApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "fontsFolder", int32(r.StatusCode), e)
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method
*/
func TestPutSlidesDocumentFromHtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    e := initializeTest("PutSlidesDocumentFromHtml", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PutSlidesDocumentFromHtml(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSlidesDocumentFromHtmlRequest() PutSlidesDocumentFromHtmlRequest {
    var request PutSlidesDocumentFromHtmlRequest
    request.name = createTestParamValue("PutSlidesDocumentFromHtml", "name", "string").(string)
    request.html = createTestParamValue("PutSlidesDocumentFromHtml", "html", "string").(string)
    request.password = createTestParamValue("PutSlidesDocumentFromHtml", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesDocumentFromHtml", "storage", "string").(string)
    request.folder = createTestParamValue("PutSlidesDocumentFromHtml", "folder", "string").(string)
    return request
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method with invalid name
*/
func TestPutSlidesDocumentFromHtmlInvalidname(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "name", r.Code, e)
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method with invalid html
*/
func TestPutSlidesDocumentFromHtmlInvalidhtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.html = invalidizeTestParamValue(request.html, "html", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "html", request.html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "html", r.Code, e)
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method with invalid password
*/
func TestPutSlidesDocumentFromHtmlInvalidpassword(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "password", r.Code, e)
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method with invalid storage
*/
func TestPutSlidesDocumentFromHtmlInvalidstorage(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "storage", r.Code, e)
}

/* DocumentApiServiceTests Create presentation document from html
   Test for DocumentApi.PutSlidesDocumentFromHtml method with invalid folder
*/
func TestPutSlidesDocumentFromHtmlInvalidfolder(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "folder", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method
*/
func TestPutSlidesSlideSize(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    e := initializeTest("PutSlidesSlideSize", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.DocumentApi.PutSlidesSlideSize(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSlidesSlideSizeRequest() PutSlidesSlideSizeRequest {
    var request PutSlidesSlideSizeRequest
    request.name = createTestParamValue("PutSlidesSlideSize", "name", "string").(string)
    request.password = createTestParamValue("PutSlidesSlideSize", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlideSize", "storage", "string").(string)
    request.folder = createTestParamValue("PutSlidesSlideSize", "folder", "string").(string)
    testwidth := createTestParamValue("PutSlidesSlideSize", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.width = new(int32)
        *request.width = v
    }
    testheight := createTestParamValue("PutSlidesSlideSize", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.height = new(int32)
        *request.height = v
    }
    request.sizeType = createTestParamValue("PutSlidesSlideSize", "sizeType", "string").(string)
    request.scaleType = createTestParamValue("PutSlidesSlideSize", "scaleType", "string").(string)
    return request
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid name
*/
func TestPutSlidesSlideSizeInvalidname(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "name", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid password
*/
func TestPutSlidesSlideSizeInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "password", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid storage
*/
func TestPutSlidesSlideSizeInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "storage", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid folder
*/
func TestPutSlidesSlideSizeInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "folder", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid width
*/
func TestPutSlidesSlideSizeInvalidwidth(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.width = new(int32)
    *request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)

    e := initializeTest("PutSlidesSlideSize", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "width", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid height
*/
func TestPutSlidesSlideSizeInvalidheight(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.height = new(int32)
    *request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)

    e := initializeTest("PutSlidesSlideSize", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "height", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid sizeType
*/
func TestPutSlidesSlideSizeInvalidsizeType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.sizeType = invalidizeTestParamValue(request.sizeType, "sizeType", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "sizeType", request.sizeType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "sizeType", r.Code, e)
}

/* DocumentApiServiceTests Set slide size for the presentation.
   Test for DocumentApi.PutSlidesSlideSize method with invalid scaleType
*/
func TestPutSlidesSlideSizeInvalidscaleType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.scaleType = invalidizeTestParamValue(request.scaleType, "scaleType", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "scaleType", request.scaleType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().DocumentApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "scaleType", r.Code, e)
}
