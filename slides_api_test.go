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

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method
*/
func TestCopyFile(t *testing.T) {
    request := createCopyFileRequest()
    e := initializeTest("CopyFile", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.CopyFile(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createCopyFileRequest() CopyFileRequest {
    var request CopyFileRequest
    request.srcPath = createTestParamValue("CopyFile", "srcPath", "string").(string)
    request.destPath = createTestParamValue("CopyFile", "destPath", "string").(string)
    request.srcStorageName = createTestParamValue("CopyFile", "srcStorageName", "string").(string)
    request.destStorageName = createTestParamValue("CopyFile", "destStorageName", "string").(string)
    request.versionId = createTestParamValue("CopyFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid srcPath
*/
func TestCopyFileInvalidsrcPath(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.srcPath, "CopyFile", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcPath = nullValue
    } else {
        request.srcPath = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "srcPath", request.srcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destPath
*/
func TestCopyFileInvaliddestPath(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.destPath, "CopyFile", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destPath = nullValue
    } else {
        request.destPath = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "destPath", request.destPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid srcStorageName
*/
func TestCopyFileInvalidsrcStorageName(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.srcStorageName, "CopyFile", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcStorageName = nullValue
    } else {
        request.srcStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "srcStorageName", request.srcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destStorageName
*/
func TestCopyFileInvaliddestStorageName(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.destStorageName, "CopyFile", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destStorageName = nullValue
    } else {
        request.destStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "destStorageName", request.destStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid versionId
*/
func TestCopyFileInvalidversionId(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.versionId, "CopyFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.versionId = nullValue
    } else {
        request.versionId = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "versionId", request.versionId, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method
*/
func TestCopyFolder(t *testing.T) {
    request := createCopyFolderRequest()
    e := initializeTest("CopyFolder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.CopyFolder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createCopyFolderRequest() CopyFolderRequest {
    var request CopyFolderRequest
    request.srcPath = createTestParamValue("CopyFolder", "srcPath", "string").(string)
    request.destPath = createTestParamValue("CopyFolder", "destPath", "string").(string)
    request.srcStorageName = createTestParamValue("CopyFolder", "srcStorageName", "string").(string)
    request.destStorageName = createTestParamValue("CopyFolder", "destStorageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid srcPath
*/
func TestCopyFolderInvalidsrcPath(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.srcPath, "CopyFolder", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcPath = nullValue
    } else {
        request.srcPath = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "srcPath", request.srcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destPath
*/
func TestCopyFolderInvaliddestPath(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.destPath, "CopyFolder", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destPath = nullValue
    } else {
        request.destPath = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "destPath", request.destPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid srcStorageName
*/
func TestCopyFolderInvalidsrcStorageName(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.srcStorageName, "CopyFolder", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcStorageName = nullValue
    } else {
        request.srcStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "srcStorageName", request.srcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destStorageName
*/
func TestCopyFolderInvaliddestStorageName(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.destStorageName, "CopyFolder", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destStorageName = nullValue
    } else {
        request.destStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "destStorageName", request.destStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method
*/
func TestCreateFolder(t *testing.T) {
    request := createCreateFolderRequest()
    e := initializeTest("CreateFolder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.CreateFolder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createCreateFolderRequest() CreateFolderRequest {
    var request CreateFolderRequest
    request.path = createTestParamValue("CreateFolder", "path", "string").(string)
    request.storageName = createTestParamValue("CreateFolder", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method with invalid path
*/
func TestCreateFolderInvalidpath(t *testing.T) {
    request := createCreateFolderRequest()

    invalidValue := invalidizeTestParamValue(request.path, "CreateFolder", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("CreateFolder", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CreateFolder", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method with invalid storageName
*/
func TestCreateFolderInvalidstorageName(t *testing.T) {
    request := createCreateFolderRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "CreateFolder", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("CreateFolder", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CreateFolder", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method
*/
func TestDeleteFile(t *testing.T) {
    request := createDeleteFileRequest()
    e := initializeTest("DeleteFile", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.DeleteFile(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteFileRequest() DeleteFileRequest {
    var request DeleteFileRequest
    request.path = createTestParamValue("DeleteFile", "path", "string").(string)
    request.storageName = createTestParamValue("DeleteFile", "storageName", "string").(string)
    request.versionId = createTestParamValue("DeleteFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid path
*/
func TestDeleteFileInvalidpath(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid storageName
*/
func TestDeleteFileInvalidstorageName(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "DeleteFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid versionId
*/
func TestDeleteFileInvalidversionId(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.versionId, "DeleteFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.versionId = nullValue
    } else {
        request.versionId = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "versionId", request.versionId, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method
*/
func TestDeleteFolder(t *testing.T) {
    request := createDeleteFolderRequest()
    e := initializeTest("DeleteFolder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.DeleteFolder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteFolderRequest() DeleteFolderRequest {
    var request DeleteFolderRequest
    request.path = createTestParamValue("DeleteFolder", "path", "string").(string)
    request.storageName = createTestParamValue("DeleteFolder", "storageName", "string").(string)
    testrecursive := createTestParamValue("DeleteFolder", "recursive", "bool")
    switch v := testrecursive.(type) { 
    case bool:
        request.recursive = new(bool)
        *request.recursive = v
    }
    return request
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid path
*/
func TestDeleteFolderInvalidpath(t *testing.T) {
    request := createDeleteFolderRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteFolder", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteFolder", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid storageName
*/
func TestDeleteFolderInvalidstorageName(t *testing.T) {
    request := createDeleteFolderRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "DeleteFolder", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("DeleteFolder", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid recursive
*/
func TestDeleteFolderInvalidrecursive(t *testing.T) {
    request := createDeleteFolderRequest()
    request.recursive = new(bool)

    invalidValue := invalidizeTestParamValue(request.recursive, "DeleteFolder", "recursive", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.recursive = nullValue
    } else {
        *request.recursive = invalidValue.(bool)
    }

    e := initializeTest("DeleteFolder", "recursive", request.recursive)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "recursive", request.recursive, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method
*/
func TestDeleteNotesSlide(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    e := initializeTest("DeleteNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid name
*/
func TestDeleteNotesSlideInvalidname(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid slideIndex
*/
func TestDeleteNotesSlideInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid password
*/
func TestDeleteNotesSlideInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid folder
*/
func TestDeleteNotesSlideInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid storage
*/
func TestDeleteNotesSlideInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method
*/
func TestDeleteNotesSlideParagraph(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    e := initializeTest("DeleteNotesSlideParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlideParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid name
*/
func TestDeleteNotesSlideParagraphInvalidname(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlideParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlideParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid path
*/
func TestDeleteNotesSlideParagraphInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlideParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteNotesSlideParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid paragraphIndex
*/
func TestDeleteNotesSlideParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeleteNotesSlideParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid password
*/
func TestDeleteNotesSlideParagraphInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlideParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid folder
*/
func TestDeleteNotesSlideParagraphInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlideParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid storage
*/
func TestDeleteNotesSlideParagraphInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlideParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method
*/
func TestDeleteNotesSlideParagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    e := initializeTest("DeleteNotesSlideParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlideParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid name
*/
func TestDeleteNotesSlideParagraphsInvalidname(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlideParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlideParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid path
*/
func TestDeleteNotesSlideParagraphsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlideParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteNotesSlideParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid paragraphs
*/
func TestDeleteNotesSlideParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphs, "DeleteNotesSlideParagraphs", "paragraphs", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.paragraphs = nullValue
    } else {
        request.paragraphs = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "paragraphs", request.paragraphs, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid password
*/
func TestDeleteNotesSlideParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlideParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid folder
*/
func TestDeleteNotesSlideParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlideParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid storage
*/
func TestDeleteNotesSlideParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlideParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method
*/
func TestDeleteNotesSlidePortion(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    e := initializeTest("DeleteNotesSlidePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlidePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid name
*/
func TestDeleteNotesSlidePortionInvalidname(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlidePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlidePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid path
*/
func TestDeleteNotesSlidePortionInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlidePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteNotesSlidePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeleteNotesSlidePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid portionIndex
*/
func TestDeleteNotesSlidePortionInvalidportionIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "DeleteNotesSlidePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid password
*/
func TestDeleteNotesSlidePortionInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlidePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid folder
*/
func TestDeleteNotesSlidePortionInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlidePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid storage
*/
func TestDeleteNotesSlidePortionInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlidePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method
*/
func TestDeleteNotesSlidePortions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    e := initializeTest("DeleteNotesSlidePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlidePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid name
*/
func TestDeleteNotesSlidePortionsInvalidname(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlidePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlidePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid path
*/
func TestDeleteNotesSlidePortionsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlidePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteNotesSlidePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeleteNotesSlidePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid portions
*/
func TestDeleteNotesSlidePortionsInvalidportions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.portions, "DeleteNotesSlidePortions", "portions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.portions = nullValue
    } else {
        request.portions = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "portions", request.portions, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid password
*/
func TestDeleteNotesSlidePortionsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlidePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid folder
*/
func TestDeleteNotesSlidePortionsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlidePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid storage
*/
func TestDeleteNotesSlidePortionsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlidePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method
*/
func TestDeleteNotesSlideShape(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    e := initializeTest("DeleteNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid name
*/
func TestDeleteNotesSlideShapeInvalidname(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid slideIndex
*/
func TestDeleteNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid path
*/
func TestDeleteNotesSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid shapeIndex
*/
func TestDeleteNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid password
*/
func TestDeleteNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid folder
*/
func TestDeleteNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid storage
*/
func TestDeleteNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method
*/
func TestDeleteNotesSlideShapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    e := initializeTest("DeleteNotesSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteNotesSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid name
*/
func TestDeleteNotesSlideShapesInvalidname(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteNotesSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid slideIndex
*/
func TestDeleteNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteNotesSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid path
*/
func TestDeleteNotesSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteNotesSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid shapes
*/
func TestDeleteNotesSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.shapes, "DeleteNotesSlideShapes", "shapes", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.shapes = nullValue
    } else {
        request.shapes = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "shapes", request.shapes, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid password
*/
func TestDeleteNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteNotesSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid folder
*/
func TestDeleteNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteNotesSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid storage
*/
func TestDeleteNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteNotesSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method
*/
func TestDeleteParagraph(t *testing.T) {
    request := createDeleteParagraphRequest()
    e := initializeTest("DeleteParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid name
*/
func TestDeleteParagraphInvalidname(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid slideIndex
*/
func TestDeleteParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid path
*/
func TestDeleteParagraphInvalidpath(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid shapeIndex
*/
func TestDeleteParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid paragraphIndex
*/
func TestDeleteParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeleteParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid password
*/
func TestDeleteParagraphInvalidpassword(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid folder
*/
func TestDeleteParagraphInvalidfolder(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid storage
*/
func TestDeleteParagraphInvalidstorage(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method
*/
func TestDeleteParagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()
    e := initializeTest("DeleteParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid name
*/
func TestDeleteParagraphsInvalidname(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid slideIndex
*/
func TestDeleteParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid path
*/
func TestDeleteParagraphsInvalidpath(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid shapeIndex
*/
func TestDeleteParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid paragraphs
*/
func TestDeleteParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphs, "DeleteParagraphs", "paragraphs", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.paragraphs = nullValue
    } else {
        request.paragraphs = invalidValue.([]int32)
    }

    e := initializeTest("DeleteParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "paragraphs", request.paragraphs, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid password
*/
func TestDeleteParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid folder
*/
func TestDeleteParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid storage
*/
func TestDeleteParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method
*/
func TestDeletePortion(t *testing.T) {
    request := createDeletePortionRequest()
    e := initializeTest("DeletePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeletePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid name
*/
func TestDeletePortionInvalidname(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeletePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid slideIndex
*/
func TestDeletePortionInvalidslideIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeletePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid path
*/
func TestDeletePortionInvalidpath(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeletePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid shapeIndex
*/
func TestDeletePortionInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeletePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid paragraphIndex
*/
func TestDeletePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeletePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid portionIndex
*/
func TestDeletePortionInvalidportionIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "DeletePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid password
*/
func TestDeletePortionInvalidpassword(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeletePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid folder
*/
func TestDeletePortionInvalidfolder(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeletePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid storage
*/
func TestDeletePortionInvalidstorage(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeletePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method
*/
func TestDeletePortions(t *testing.T) {
    request := createDeletePortionsRequest()
    e := initializeTest("DeletePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeletePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid name
*/
func TestDeletePortionsInvalidname(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeletePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid slideIndex
*/
func TestDeletePortionsInvalidslideIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeletePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid path
*/
func TestDeletePortionsInvalidpath(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeletePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid shapeIndex
*/
func TestDeletePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeletePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid paragraphIndex
*/
func TestDeletePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "DeletePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid portions
*/
func TestDeletePortionsInvalidportions(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.portions, "DeletePortions", "portions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.portions = nullValue
    } else {
        request.portions = invalidValue.([]int32)
    }

    e := initializeTest("DeletePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "portions", request.portions, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid password
*/
func TestDeletePortionsInvalidpassword(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeletePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid folder
*/
func TestDeletePortionsInvalidfolder(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeletePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid storage
*/
func TestDeletePortionsInvalidstorage(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeletePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method
*/
func TestDeleteSlideAnimation(t *testing.T) {
    request := createDeleteSlideAnimationRequest()
    e := initializeTest("DeleteSlideAnimation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationRequest() DeleteSlideAnimationRequest {
    var request DeleteSlideAnimationRequest
    request.name = createTestParamValue("DeleteSlideAnimation", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimation", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimation", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimation", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid name
*/
func TestDeleteSlideAnimationInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid slideIndex
*/
func TestDeleteSlideAnimationInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimation", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid password
*/
func TestDeleteSlideAnimationInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid folder
*/
func TestDeleteSlideAnimationInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid storage
*/
func TestDeleteSlideAnimationInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method
*/
func TestDeleteSlideAnimationEffect(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()
    e := initializeTest("DeleteSlideAnimationEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimationEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationEffectRequest() DeleteSlideAnimationEffectRequest {
    var request DeleteSlideAnimationEffectRequest
    request.name = createTestParamValue("DeleteSlideAnimationEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.effectIndex = createTestParamValue("DeleteSlideAnimationEffect", "effectIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimationEffect", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimationEffect", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid name
*/
func TestDeleteSlideAnimationEffectInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid slideIndex
*/
func TestDeleteSlideAnimationEffectInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid effectIndex
*/
func TestDeleteSlideAnimationEffectInvalideffectIndex(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effectIndex, "DeleteSlideAnimationEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.effectIndex = nullValue
    } else {
        request.effectIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "effectIndex", request.effectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "effectIndex", request.effectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid password
*/
func TestDeleteSlideAnimationEffectInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid folder
*/
func TestDeleteSlideAnimationEffectInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid storage
*/
func TestDeleteSlideAnimationEffectInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method
*/
func TestDeleteSlideAnimationInteractiveSequence(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()
    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationInteractiveSequenceRequest() DeleteSlideAnimationInteractiveSequenceRequest {
    var request DeleteSlideAnimationInteractiveSequenceRequest
    request.name = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "slideIndex", "int32").(int32)
    request.sequenceIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "sequenceIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimationInteractiveSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimationInteractiveSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid sequenceIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidsequenceIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.sequenceIndex, "DeleteSlideAnimationInteractiveSequence", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.sequenceIndex = nullValue
    } else {
        request.sequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "sequenceIndex", request.sequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "sequenceIndex", request.sequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimationInteractiveSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimationInteractiveSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimationInteractiveSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method
*/
func TestDeleteSlideAnimationInteractiveSequenceEffect(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()
    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationInteractiveSequenceEffectRequest() DeleteSlideAnimationInteractiveSequenceEffectRequest {
    var request DeleteSlideAnimationInteractiveSequenceEffectRequest
    request.name = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.sequenceIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.effectIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidsequenceIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.sequenceIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.sequenceIndex = nullValue
    } else {
        request.sequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid effectIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalideffectIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effectIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.effectIndex = nullValue
    } else {
        request.effectIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", request.effectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", request.effectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method
*/
func TestDeleteSlideAnimationInteractiveSequences(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()
    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationInteractiveSequencesRequest() DeleteSlideAnimationInteractiveSequencesRequest {
    var request DeleteSlideAnimationInteractiveSequencesRequest
    request.name = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimationInteractiveSequences", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimationInteractiveSequences", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimationInteractiveSequences", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimationInteractiveSequences", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimationInteractiveSequences", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method
*/
func TestDeleteSlideAnimationMainSequence(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()
    e := initializeTest("DeleteSlideAnimationMainSequence", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideAnimationMainSequence(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideAnimationMainSequenceRequest() DeleteSlideAnimationMainSequenceRequest {
    var request DeleteSlideAnimationMainSequenceRequest
    request.name = createTestParamValue("DeleteSlideAnimationMainSequence", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideAnimationMainSequence", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideAnimationMainSequence", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideAnimationMainSequence", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideAnimationMainSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid name
*/
func TestDeleteSlideAnimationMainSequenceInvalidname(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideAnimationMainSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid slideIndex
*/
func TestDeleteSlideAnimationMainSequenceInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideAnimationMainSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid password
*/
func TestDeleteSlideAnimationMainSequenceInvalidpassword(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideAnimationMainSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid folder
*/
func TestDeleteSlideAnimationMainSequenceInvalidfolder(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideAnimationMainSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid storage
*/
func TestDeleteSlideAnimationMainSequenceInvalidstorage(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideAnimationMainSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method
*/
func TestDeleteSlideByIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    e := initializeTest("DeleteSlideByIndex", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideByIndex(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlideByIndexRequest() DeleteSlideByIndexRequest {
    var request DeleteSlideByIndexRequest
    request.name = createTestParamValue("DeleteSlideByIndex", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideByIndex", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideByIndex", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideByIndex", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideByIndex", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid name
*/
func TestDeleteSlideByIndexInvalidname(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideByIndex", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid slideIndex
*/
func TestDeleteSlideByIndexInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideByIndex", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideByIndex", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid password
*/
func TestDeleteSlideByIndexInvalidpassword(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideByIndex", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid folder
*/
func TestDeleteSlideByIndexInvalidfolder(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideByIndex", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid storage
*/
func TestDeleteSlideByIndexInvalidstorage(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideByIndex", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method
*/
func TestDeleteSlideShape(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    e := initializeTest("DeleteSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid name
*/
func TestDeleteSlideShapeInvalidname(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid slideIndex
*/
func TestDeleteSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid path
*/
func TestDeleteSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid shapeIndex
*/
func TestDeleteSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "DeleteSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid password
*/
func TestDeleteSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid folder
*/
func TestDeleteSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid storage
*/
func TestDeleteSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method
*/
func TestDeleteSlideShapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    e := initializeTest("DeleteSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid name
*/
func TestDeleteSlideShapesInvalidname(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid slideIndex
*/
func TestDeleteSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid path
*/
func TestDeleteSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DeleteSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid shapes
*/
func TestDeleteSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.shapes, "DeleteSlideShapes", "shapes", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.shapes = nullValue
    } else {
        request.shapes = invalidValue.([]int32)
    }

    e := initializeTest("DeleteSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "shapes", request.shapes, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid password
*/
func TestDeleteSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid folder
*/
func TestDeleteSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid storage
*/
func TestDeleteSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method
*/
func TestDeleteSlidesCleanSlidesList(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    e := initializeTest("DeleteSlidesCleanSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlidesCleanSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlidesCleanSlidesListRequest() DeleteSlidesCleanSlidesListRequest {
    var request DeleteSlidesCleanSlidesListRequest
    request.name = createTestParamValue("DeleteSlidesCleanSlidesList", "name", "string").(string)
    request.slides = createTestParamValue("DeleteSlidesCleanSlidesList", "slides", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteSlidesCleanSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesCleanSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesCleanSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid name
*/
func TestDeleteSlidesCleanSlidesListInvalidname(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlidesCleanSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid slides
*/
func TestDeleteSlidesCleanSlidesListInvalidslides(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.slides, "DeleteSlidesCleanSlidesList", "slides", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.slides = nullValue
    } else {
        request.slides = invalidValue.([]int32)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "slides", request.slides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "slides", request.slides, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid password
*/
func TestDeleteSlidesCleanSlidesListInvalidpassword(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlidesCleanSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid folder
*/
func TestDeleteSlidesCleanSlidesListInvalidfolder(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlidesCleanSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid storage
*/
func TestDeleteSlidesCleanSlidesListInvalidstorage(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlidesCleanSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method
*/
func TestDeleteSlidesDocumentProperties(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    e := initializeTest("DeleteSlidesDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlidesDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlidesDocumentPropertiesRequest() DeleteSlidesDocumentPropertiesRequest {
    var request DeleteSlidesDocumentPropertiesRequest
    request.name = createTestParamValue("DeleteSlidesDocumentProperties", "name", "string").(string)
    request.password = createTestParamValue("DeleteSlidesDocumentProperties", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesDocumentProperties", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid name
*/
func TestDeleteSlidesDocumentPropertiesInvalidname(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlidesDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid password
*/
func TestDeleteSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlidesDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid folder
*/
func TestDeleteSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlidesDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid storage
*/
func TestDeleteSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlidesDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method
*/
func TestDeleteSlidesDocumentProperty(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    e := initializeTest("DeleteSlidesDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlidesDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlidesDocumentPropertyRequest() DeleteSlidesDocumentPropertyRequest {
    var request DeleteSlidesDocumentPropertyRequest
    request.name = createTestParamValue("DeleteSlidesDocumentProperty", "name", "string").(string)
    request.propertyName = createTestParamValue("DeleteSlidesDocumentProperty", "propertyName", "string").(string)
    request.password = createTestParamValue("DeleteSlidesDocumentProperty", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesDocumentProperty", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid name
*/
func TestDeleteSlidesDocumentPropertyInvalidname(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlidesDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid propertyName
*/
func TestDeleteSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.propertyName, "DeleteSlidesDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.propertyName = nullValue
    } else {
        request.propertyName = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "propertyName", request.propertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid password
*/
func TestDeleteSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlidesDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid folder
*/
func TestDeleteSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlidesDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid storage
*/
func TestDeleteSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlidesDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method
*/
func TestDeleteSlidesSlideBackground(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    e := initializeTest("DeleteSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.DeleteSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createDeleteSlidesSlideBackgroundRequest() DeleteSlidesSlideBackgroundRequest {
    var request DeleteSlidesSlideBackgroundRequest
    request.name = createTestParamValue("DeleteSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlidesSlideBackground", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesSlideBackground", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid name
*/
func TestDeleteSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.name, "DeleteSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid slideIndex
*/
func TestDeleteSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "DeleteSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid password
*/
func TestDeleteSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.password, "DeleteSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid folder
*/
func TestDeleteSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "DeleteSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid storage
*/
func TestDeleteSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "DeleteSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method
*/
func TestDownloadFile(t *testing.T) {
    request := createDownloadFileRequest()
    e := initializeTest("DownloadFile", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.DownloadFile(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createDownloadFileRequest() DownloadFileRequest {
    var request DownloadFileRequest
    request.path = createTestParamValue("DownloadFile", "path", "string").(string)
    request.storageName = createTestParamValue("DownloadFile", "storageName", "string").(string)
    request.versionId = createTestParamValue("DownloadFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid path
*/
func TestDownloadFileInvalidpath(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.path, "DownloadFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid storageName
*/
func TestDownloadFileInvalidstorageName(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "DownloadFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid versionId
*/
func TestDownloadFileInvalidversionId(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.versionId, "DownloadFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.versionId = nullValue
    } else {
        request.versionId = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "versionId", request.versionId, int32(statusCode), e)
}

/* SlidesApiServiceTests Get disc usage
   Test for SlidesApi.GetDiscUsage method
*/
func TestGetDiscUsage(t *testing.T) {
    request := createGetDiscUsageRequest()
    e := initializeTest("GetDiscUsage", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetDiscUsage(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetDiscUsageRequest() GetDiscUsageRequest {
    var request GetDiscUsageRequest
    request.storageName = createTestParamValue("GetDiscUsage", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get disc usage
   Test for SlidesApi.GetDiscUsage method with invalid storageName
*/
func TestGetDiscUsageInvalidstorageName(t *testing.T) {
    request := createGetDiscUsageRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "GetDiscUsage", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("GetDiscUsage", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetDiscUsage(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetDiscUsage", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method
*/
func TestGetFileVersions(t *testing.T) {
    request := createGetFileVersionsRequest()
    e := initializeTest("GetFileVersions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetFileVersions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetFileVersionsRequest() GetFileVersionsRequest {
    var request GetFileVersionsRequest
    request.path = createTestParamValue("GetFileVersions", "path", "string").(string)
    request.storageName = createTestParamValue("GetFileVersions", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method with invalid path
*/
func TestGetFileVersionsInvalidpath(t *testing.T) {
    request := createGetFileVersionsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetFileVersions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetFileVersions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFileVersions", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method with invalid storageName
*/
func TestGetFileVersionsInvalidstorageName(t *testing.T) {
    request := createGetFileVersionsRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "GetFileVersions", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("GetFileVersions", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFileVersions", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method
*/
func TestGetFilesList(t *testing.T) {
    request := createGetFilesListRequest()
    e := initializeTest("GetFilesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetFilesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetFilesListRequest() GetFilesListRequest {
    var request GetFilesListRequest
    request.path = createTestParamValue("GetFilesList", "path", "string").(string)
    request.storageName = createTestParamValue("GetFilesList", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method with invalid path
*/
func TestGetFilesListInvalidpath(t *testing.T) {
    request := createGetFilesListRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetFilesList", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetFilesList", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFilesList", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method with invalid storageName
*/
func TestGetFilesListInvalidstorageName(t *testing.T) {
    request := createGetFilesListRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "GetFilesList", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("GetFilesList", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFilesList", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method
*/
func TestGetLayoutSlide(t *testing.T) {
    request := createGetLayoutSlideRequest()
    e := initializeTest("GetLayoutSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetLayoutSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid name
*/
func TestGetLayoutSlideInvalidname(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetLayoutSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid slideIndex
*/
func TestGetLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetLayoutSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid password
*/
func TestGetLayoutSlideInvalidpassword(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetLayoutSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid folder
*/
func TestGetLayoutSlideInvalidfolder(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetLayoutSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid storage
*/
func TestGetLayoutSlideInvalidstorage(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetLayoutSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method
*/
func TestGetLayoutSlidesList(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    e := initializeTest("GetLayoutSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetLayoutSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid name
*/
func TestGetLayoutSlidesListInvalidname(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetLayoutSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid password
*/
func TestGetLayoutSlidesListInvalidpassword(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetLayoutSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid folder
*/
func TestGetLayoutSlidesListInvalidfolder(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetLayoutSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid storage
*/
func TestGetLayoutSlidesListInvalidstorage(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetLayoutSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method
*/
func TestGetMasterSlide(t *testing.T) {
    request := createGetMasterSlideRequest()
    e := initializeTest("GetMasterSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetMasterSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid name
*/
func TestGetMasterSlideInvalidname(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetMasterSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid slideIndex
*/
func TestGetMasterSlideInvalidslideIndex(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetMasterSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetMasterSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid password
*/
func TestGetMasterSlideInvalidpassword(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetMasterSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid folder
*/
func TestGetMasterSlideInvalidfolder(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetMasterSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid storage
*/
func TestGetMasterSlideInvalidstorage(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetMasterSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method
*/
func TestGetMasterSlidesList(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    e := initializeTest("GetMasterSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetMasterSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid name
*/
func TestGetMasterSlidesListInvalidname(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetMasterSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid password
*/
func TestGetMasterSlidesListInvalidpassword(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetMasterSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid folder
*/
func TestGetMasterSlidesListInvalidfolder(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetMasterSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid storage
*/
func TestGetMasterSlidesListInvalidstorage(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetMasterSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method
*/
func TestGetNotesSlide(t *testing.T) {
    request := createGetNotesSlideRequest()
    e := initializeTest("GetNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid name
*/
func TestGetNotesSlideInvalidname(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid slideIndex
*/
func TestGetNotesSlideInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid password
*/
func TestGetNotesSlideInvalidpassword(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid folder
*/
func TestGetNotesSlideInvalidfolder(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid storage
*/
func TestGetNotesSlideInvalidstorage(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method
*/
func TestGetNotesSlideShape(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    e := initializeTest("GetNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid name
*/
func TestGetNotesSlideShapeInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid slideIndex
*/
func TestGetNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid path
*/
func TestGetNotesSlideShapeInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid shapeIndex
*/
func TestGetNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid password
*/
func TestGetNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid folder
*/
func TestGetNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid storage
*/
func TestGetNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method
*/
func TestGetNotesSlideShapeParagraph(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    e := initializeTest("GetNotesSlideShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid name
*/
func TestGetNotesSlideShapeParagraphInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid path
*/
func TestGetNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetNotesSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetNotesSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid password
*/
func TestGetNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid folder
*/
func TestGetNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid storage
*/
func TestGetNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method
*/
func TestGetNotesSlideShapeParagraphs(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    e := initializeTest("GetNotesSlideShapeParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShapeParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid name
*/
func TestGetNotesSlideShapeParagraphsInvalidname(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShapeParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShapeParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid path
*/
func TestGetNotesSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShapeParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetNotesSlideShapeParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid password
*/
func TestGetNotesSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShapeParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid folder
*/
func TestGetNotesSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShapeParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid storage
*/
func TestGetNotesSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShapeParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method
*/
func TestGetNotesSlideShapePortion(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    e := initializeTest("GetNotesSlideShapePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShapePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid name
*/
func TestGetNotesSlideShapePortionInvalidname(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShapePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShapePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid path
*/
func TestGetNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShapePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetNotesSlideShapePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetNotesSlideShapePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid portionIndex
*/
func TestGetNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "GetNotesSlideShapePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid password
*/
func TestGetNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShapePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid folder
*/
func TestGetNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShapePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid storage
*/
func TestGetNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShapePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method
*/
func TestGetNotesSlideShapePortions(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    e := initializeTest("GetNotesSlideShapePortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShapePortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid name
*/
func TestGetNotesSlideShapePortionsInvalidname(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShapePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShapePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid path
*/
func TestGetNotesSlideShapePortionsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShapePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetNotesSlideShapePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetNotesSlideShapePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid password
*/
func TestGetNotesSlideShapePortionsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShapePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid folder
*/
func TestGetNotesSlideShapePortionsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShapePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid storage
*/
func TestGetNotesSlideShapePortionsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShapePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method
*/
func TestGetNotesSlideShapes(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    e := initializeTest("GetNotesSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetNotesSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid name
*/
func TestGetNotesSlideShapesInvalidname(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid slideIndex
*/
func TestGetNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid path
*/
func TestGetNotesSlideShapesInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetNotesSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid password
*/
func TestGetNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid folder
*/
func TestGetNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid storage
*/
func TestGetNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method
*/
func TestGetNotesSlideWithFormat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    e := initializeTest("GetNotesSlideWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetNotesSlideWithFormat(request)
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
    request.fontsFolder = createTestParamValue("GetNotesSlideWithFormat", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid name
*/
func TestGetNotesSlideWithFormatInvalidname(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetNotesSlideWithFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid slideIndex
*/
func TestGetNotesSlideWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetNotesSlideWithFormat", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid format
*/
func TestGetNotesSlideWithFormatInvalidformat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.format, "GetNotesSlideWithFormat", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid width
*/
func TestGetNotesSlideWithFormatInvalidwidth(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.width = new(int32)

    invalidValue := invalidizeTestParamValue(request.width, "GetNotesSlideWithFormat", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.width = nullValue
    } else {
        *request.width = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "width", request.width, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid height
*/
func TestGetNotesSlideWithFormatInvalidheight(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.height = new(int32)

    invalidValue := invalidizeTestParamValue(request.height, "GetNotesSlideWithFormat", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.height = nullValue
    } else {
        *request.height = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "height", request.height, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid password
*/
func TestGetNotesSlideWithFormatInvalidpassword(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetNotesSlideWithFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid folder
*/
func TestGetNotesSlideWithFormatInvalidfolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetNotesSlideWithFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid storage
*/
func TestGetNotesSlideWithFormatInvalidstorage(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetNotesSlideWithFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid fontsFolder
*/
func TestGetNotesSlideWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "GetNotesSlideWithFormat", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method
*/
func TestGetParagraphPortion(t *testing.T) {
    request := createGetParagraphPortionRequest()
    e := initializeTest("GetParagraphPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetParagraphPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid name
*/
func TestGetParagraphPortionInvalidname(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetParagraphPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid slideIndex
*/
func TestGetParagraphPortionInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetParagraphPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid path
*/
func TestGetParagraphPortionInvalidpath(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetParagraphPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid shapeIndex
*/
func TestGetParagraphPortionInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetParagraphPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid paragraphIndex
*/
func TestGetParagraphPortionInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetParagraphPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid portionIndex
*/
func TestGetParagraphPortionInvalidportionIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "GetParagraphPortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid password
*/
func TestGetParagraphPortionInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetParagraphPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid folder
*/
func TestGetParagraphPortionInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetParagraphPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid storage
*/
func TestGetParagraphPortionInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetParagraphPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method
*/
func TestGetParagraphPortions(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    e := initializeTest("GetParagraphPortions", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetParagraphPortions(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid name
*/
func TestGetParagraphPortionsInvalidname(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetParagraphPortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid slideIndex
*/
func TestGetParagraphPortionsInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetParagraphPortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid path
*/
func TestGetParagraphPortionsInvalidpath(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetParagraphPortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid shapeIndex
*/
func TestGetParagraphPortionsInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetParagraphPortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid paragraphIndex
*/
func TestGetParagraphPortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetParagraphPortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid password
*/
func TestGetParagraphPortionsInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetParagraphPortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid folder
*/
func TestGetParagraphPortionsInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetParagraphPortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid storage
*/
func TestGetParagraphPortionsInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetParagraphPortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method
*/
func TestGetSlideAnimation(t *testing.T) {
    request := createGetSlideAnimationRequest()
    e := initializeTest("GetSlideAnimation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlideAnimation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlideAnimationRequest() GetSlideAnimationRequest {
    var request GetSlideAnimationRequest
    request.name = createTestParamValue("GetSlideAnimation", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideAnimation", "slideIndex", "int32").(int32)
    request.shapeIndex = createTestParamValue("GetSlideAnimation", "shapeIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlideAnimation", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideAnimation", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid name
*/
func TestGetSlideAnimationInvalidname(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid slideIndex
*/
func TestGetSlideAnimationInvalidslideIndex(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideAnimation", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid shapeIndex
*/
func TestGetSlideAnimationInvalidshapeIndex(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetSlideAnimation", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideAnimation", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid password
*/
func TestGetSlideAnimationInvalidpassword(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid folder
*/
func TestGetSlideAnimationInvalidfolder(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid storage
*/
func TestGetSlideAnimationInvalidstorage(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method
*/
func TestGetSlideShape(t *testing.T) {
    request := createGetSlideShapeRequest()
    e := initializeTest("GetSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid name
*/
func TestGetSlideShapeInvalidname(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid slideIndex
*/
func TestGetSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid path
*/
func TestGetSlideShapeInvalidpath(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid shapeIndex
*/
func TestGetSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid password
*/
func TestGetSlideShapeInvalidpassword(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid folder
*/
func TestGetSlideShapeInvalidfolder(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid storage
*/
func TestGetSlideShapeInvalidstorage(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method
*/
func TestGetSlideShapeParagraph(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    e := initializeTest("GetSlideShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlideShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlideShapeParagraphRequest() GetSlideShapeParagraphRequest {
    var request GetSlideShapeParagraphRequest
    request.name = createTestParamValue("GetSlideShapeParagraph", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("GetSlideShapeParagraph", "path", "string").(string)
    request.shapeIndex = createTestParamValue("GetSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.paragraphIndex = createTestParamValue("GetSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlideShapeParagraph", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideShapeParagraph", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid name
*/
func TestGetSlideShapeParagraphInvalidname(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid slideIndex
*/
func TestGetSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid path
*/
func TestGetSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "GetSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid password
*/
func TestGetSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid folder
*/
func TestGetSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid storage
*/
func TestGetSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method
*/
func TestGetSlideShapeParagraphs(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    e := initializeTest("GetSlideShapeParagraphs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlideShapeParagraphs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid name
*/
func TestGetSlideShapeParagraphsInvalidname(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlideShapeParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlideShapeParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid path
*/
func TestGetSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetSlideShapeParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "GetSlideShapeParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid password
*/
func TestGetSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlideShapeParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid folder
*/
func TestGetSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlideShapeParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid storage
*/
func TestGetSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlideShapeParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method
*/
func TestGetSlideShapes(t *testing.T) {
    request := createGetSlideShapesRequest()
    e := initializeTest("GetSlideShapes", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlideShapes(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid name
*/
func TestGetSlideShapesInvalidname(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid slideIndex
*/
func TestGetSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid path
*/
func TestGetSlideShapesInvalidpath(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "GetSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid password
*/
func TestGetSlideShapesInvalidpassword(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid folder
*/
func TestGetSlideShapesInvalidfolder(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid storage
*/
func TestGetSlideShapesInvalidstorage(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Get API info.
   Test for SlidesApi.GetSlidesApiInfo method
*/
func TestGetSlidesApiInfo(t *testing.T) {
    e := initializeTest("GetSlidesApiInfo", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesApiInfo()
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method
*/
func TestGetSlidesDocument(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    e := initializeTest("GetSlidesDocument", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesDocument(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid name
*/
func TestGetSlidesDocumentInvalidname(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesDocument", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid password
*/
func TestGetSlidesDocumentInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesDocument", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid storage
*/
func TestGetSlidesDocumentInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesDocument", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid folder
*/
func TestGetSlidesDocumentInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesDocument", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method
*/
func TestGetSlidesDocumentProperties(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    e := initializeTest("GetSlidesDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesDocumentPropertiesRequest() GetSlidesDocumentPropertiesRequest {
    var request GetSlidesDocumentPropertiesRequest
    request.name = createTestParamValue("GetSlidesDocumentProperties", "name", "string").(string)
    request.password = createTestParamValue("GetSlidesDocumentProperties", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesDocumentProperties", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid name
*/
func TestGetSlidesDocumentPropertiesInvalidname(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid password
*/
func TestGetSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid folder
*/
func TestGetSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid storage
*/
func TestGetSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method
*/
func TestGetSlidesDocumentProperty(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    e := initializeTest("GetSlidesDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesDocumentPropertyRequest() GetSlidesDocumentPropertyRequest {
    var request GetSlidesDocumentPropertyRequest
    request.name = createTestParamValue("GetSlidesDocumentProperty", "name", "string").(string)
    request.propertyName = createTestParamValue("GetSlidesDocumentProperty", "propertyName", "string").(string)
    request.password = createTestParamValue("GetSlidesDocumentProperty", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesDocumentProperty", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid name
*/
func TestGetSlidesDocumentPropertyInvalidname(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid propertyName
*/
func TestGetSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.propertyName, "GetSlidesDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.propertyName = nullValue
    } else {
        request.propertyName = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "propertyName", request.propertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid password
*/
func TestGetSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid folder
*/
func TestGetSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid storage
*/
func TestGetSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method
*/
func TestGetSlidesImageWithDefaultFormat(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()
    e := initializeTest("GetSlidesImageWithDefaultFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesImageWithDefaultFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetSlidesImageWithDefaultFormatRequest() GetSlidesImageWithDefaultFormatRequest {
    var request GetSlidesImageWithDefaultFormatRequest
    request.name = createTestParamValue("GetSlidesImageWithDefaultFormat", "name", "string").(string)
    request.index = createTestParamValue("GetSlidesImageWithDefaultFormat", "index", "int32").(int32)
    request.password = createTestParamValue("GetSlidesImageWithDefaultFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesImageWithDefaultFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesImageWithDefaultFormat", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid name
*/
func TestGetSlidesImageWithDefaultFormatInvalidname(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesImageWithDefaultFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid index
*/
func TestGetSlidesImageWithDefaultFormatInvalidindex(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.index, "GetSlidesImageWithDefaultFormat", "index", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.index = nullValue
    } else {
        request.index = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "index", request.index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "index", request.index, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid password
*/
func TestGetSlidesImageWithDefaultFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesImageWithDefaultFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid folder
*/
func TestGetSlidesImageWithDefaultFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesImageWithDefaultFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid storage
*/
func TestGetSlidesImageWithDefaultFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesImageWithDefaultFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method
*/
func TestGetSlidesImageWithFormat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    e := initializeTest("GetSlidesImageWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesImageWithFormat(request)
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

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid name
*/
func TestGetSlidesImageWithFormatInvalidname(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesImageWithFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid index
*/
func TestGetSlidesImageWithFormatInvalidindex(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.index, "GetSlidesImageWithFormat", "index", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.index = nullValue
    } else {
        request.index = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesImageWithFormat", "index", request.index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "index", request.index, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid format
*/
func TestGetSlidesImageWithFormatInvalidformat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.format, "GetSlidesImageWithFormat", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid password
*/
func TestGetSlidesImageWithFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesImageWithFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid folder
*/
func TestGetSlidesImageWithFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesImageWithFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid storage
*/
func TestGetSlidesImageWithFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesImageWithFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method
*/
func TestGetSlidesImages(t *testing.T) {
    request := createGetSlidesImagesRequest()
    e := initializeTest("GetSlidesImages", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesImages(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid name
*/
func TestGetSlidesImagesInvalidname(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesImages", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid password
*/
func TestGetSlidesImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesImages", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid folder
*/
func TestGetSlidesImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesImages", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid storage
*/
func TestGetSlidesImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesImages", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method
*/
func TestGetSlidesPlaceholder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    e := initializeTest("GetSlidesPlaceholder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesPlaceholder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid name
*/
func TestGetSlidesPlaceholderInvalidname(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesPlaceholder", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid slideIndex
*/
func TestGetSlidesPlaceholderInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesPlaceholder", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholder", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid placeholderIndex
*/
func TestGetSlidesPlaceholderInvalidplaceholderIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.placeholderIndex, "GetSlidesPlaceholder", "placeholderIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.placeholderIndex = nullValue
    } else {
        request.placeholderIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholder", "placeholderIndex", request.placeholderIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "placeholderIndex", request.placeholderIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid password
*/
func TestGetSlidesPlaceholderInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesPlaceholder", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid folder
*/
func TestGetSlidesPlaceholderInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesPlaceholder", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid storage
*/
func TestGetSlidesPlaceholderInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesPlaceholder", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method
*/
func TestGetSlidesPlaceholders(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    e := initializeTest("GetSlidesPlaceholders", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesPlaceholders(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid name
*/
func TestGetSlidesPlaceholdersInvalidname(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesPlaceholders", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid slideIndex
*/
func TestGetSlidesPlaceholdersInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesPlaceholders", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholders", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid password
*/
func TestGetSlidesPlaceholdersInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesPlaceholders", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid folder
*/
func TestGetSlidesPlaceholdersInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesPlaceholders", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid storage
*/
func TestGetSlidesPlaceholdersInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesPlaceholders", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method
*/
func TestGetSlidesPresentationTextItems(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    e := initializeTest("GetSlidesPresentationTextItems", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesPresentationTextItems(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesPresentationTextItemsRequest() GetSlidesPresentationTextItemsRequest {
    var request GetSlidesPresentationTextItemsRequest
    request.name = createTestParamValue("GetSlidesPresentationTextItems", "name", "string").(string)
    testwithEmpty := createTestParamValue("GetSlidesPresentationTextItems", "withEmpty", "bool")
    switch v := testwithEmpty.(type) { 
    case bool:
        request.withEmpty = new(bool)
        *request.withEmpty = v
    }
    request.password = createTestParamValue("GetSlidesPresentationTextItems", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesPresentationTextItems", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesPresentationTextItems", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid name
*/
func TestGetSlidesPresentationTextItemsInvalidname(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesPresentationTextItems", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid withEmpty
*/
func TestGetSlidesPresentationTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.withEmpty = new(bool)

    invalidValue := invalidizeTestParamValue(request.withEmpty, "GetSlidesPresentationTextItems", "withEmpty", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.withEmpty = nullValue
    } else {
        *request.withEmpty = invalidValue.(bool)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "withEmpty", request.withEmpty, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid password
*/
func TestGetSlidesPresentationTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesPresentationTextItems", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid folder
*/
func TestGetSlidesPresentationTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesPresentationTextItems", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid storage
*/
func TestGetSlidesPresentationTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesPresentationTextItems", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method
*/
func TestGetSlidesSlide(t *testing.T) {
    request := createGetSlidesSlideRequest()
    e := initializeTest("GetSlidesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesSlideRequest() GetSlidesSlideRequest {
    var request GetSlidesSlideRequest
    request.name = createTestParamValue("GetSlidesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlide", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid name
*/
func TestGetSlidesSlideInvalidname(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid slideIndex
*/
func TestGetSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid password
*/
func TestGetSlidesSlideInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid folder
*/
func TestGetSlidesSlideInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid storage
*/
func TestGetSlidesSlideInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method
*/
func TestGetSlidesSlideBackground(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    e := initializeTest("GetSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesSlideBackgroundRequest() GetSlidesSlideBackgroundRequest {
    var request GetSlidesSlideBackgroundRequest
    request.name = createTestParamValue("GetSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlideBackground", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideBackground", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid name
*/
func TestGetSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid slideIndex
*/
func TestGetSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid password
*/
func TestGetSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid folder
*/
func TestGetSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid storage
*/
func TestGetSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method
*/
func TestGetSlidesSlideComments(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    e := initializeTest("GetSlidesSlideComments", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlideComments(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesSlideCommentsRequest() GetSlidesSlideCommentsRequest {
    var request GetSlidesSlideCommentsRequest
    request.name = createTestParamValue("GetSlidesSlideComments", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideComments", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlideComments", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideComments", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideComments", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid name
*/
func TestGetSlidesSlideCommentsInvalidname(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlideComments", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid slideIndex
*/
func TestGetSlidesSlideCommentsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesSlideComments", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideComments", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid password
*/
func TestGetSlidesSlideCommentsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlideComments", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid folder
*/
func TestGetSlidesSlideCommentsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlideComments", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid storage
*/
func TestGetSlidesSlideCommentsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlideComments", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method
*/
func TestGetSlidesSlideImages(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    e := initializeTest("GetSlidesSlideImages", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlideImages(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid name
*/
func TestGetSlidesSlideImagesInvalidname(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlideImages", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid slideIndex
*/
func TestGetSlidesSlideImagesInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesSlideImages", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideImages", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid password
*/
func TestGetSlidesSlideImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlideImages", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid folder
*/
func TestGetSlidesSlideImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlideImages", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid storage
*/
func TestGetSlidesSlideImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlideImages", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method
*/
func TestGetSlidesSlideTextItems(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    e := initializeTest("GetSlidesSlideTextItems", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlideTextItems(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesSlideTextItemsRequest() GetSlidesSlideTextItemsRequest {
    var request GetSlidesSlideTextItemsRequest
    request.name = createTestParamValue("GetSlidesSlideTextItems", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideTextItems", "slideIndex", "int32").(int32)
    testwithEmpty := createTestParamValue("GetSlidesSlideTextItems", "withEmpty", "bool")
    switch v := testwithEmpty.(type) { 
    case bool:
        request.withEmpty = new(bool)
        *request.withEmpty = v
    }
    request.password = createTestParamValue("GetSlidesSlideTextItems", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideTextItems", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideTextItems", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid name
*/
func TestGetSlidesSlideTextItemsInvalidname(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlideTextItems", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid slideIndex
*/
func TestGetSlidesSlideTextItemsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesSlideTextItems", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideTextItems", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid withEmpty
*/
func TestGetSlidesSlideTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.withEmpty = new(bool)

    invalidValue := invalidizeTestParamValue(request.withEmpty, "GetSlidesSlideTextItems", "withEmpty", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.withEmpty = nullValue
    } else {
        *request.withEmpty = invalidValue.(bool)
    }

    e := initializeTest("GetSlidesSlideTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "withEmpty", request.withEmpty, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid password
*/
func TestGetSlidesSlideTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlideTextItems", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid folder
*/
func TestGetSlidesSlideTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlideTextItems", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid storage
*/
func TestGetSlidesSlideTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlideTextItems", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method
*/
func TestGetSlidesSlidesList(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    e := initializeTest("GetSlidesSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesSlidesListRequest() GetSlidesSlidesListRequest {
    var request GetSlidesSlidesListRequest
    request.name = createTestParamValue("GetSlidesSlidesList", "name", "string").(string)
    request.password = createTestParamValue("GetSlidesSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid name
*/
func TestGetSlidesSlidesListInvalidname(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid password
*/
func TestGetSlidesSlidesListInvalidpassword(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid folder
*/
func TestGetSlidesSlidesListInvalidfolder(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid storage
*/
func TestGetSlidesSlidesListInvalidstorage(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method
*/
func TestGetSlidesTheme(t *testing.T) {
    request := createGetSlidesThemeRequest()
    e := initializeTest("GetSlidesTheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesTheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesThemeRequest() GetSlidesThemeRequest {
    var request GetSlidesThemeRequest
    request.name = createTestParamValue("GetSlidesTheme", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesTheme", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesTheme", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesTheme", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesTheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid name
*/
func TestGetSlidesThemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesTheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid slideIndex
*/
func TestGetSlidesThemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesTheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesTheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid password
*/
func TestGetSlidesThemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesTheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid folder
*/
func TestGetSlidesThemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesTheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid storage
*/
func TestGetSlidesThemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesTheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method
*/
func TestGetSlidesThemeColorScheme(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    e := initializeTest("GetSlidesThemeColorScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesThemeColorScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesThemeColorSchemeRequest() GetSlidesThemeColorSchemeRequest {
    var request GetSlidesThemeColorSchemeRequest
    request.name = createTestParamValue("GetSlidesThemeColorScheme", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesThemeColorScheme", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesThemeColorScheme", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesThemeColorScheme", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesThemeColorScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid name
*/
func TestGetSlidesThemeColorSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesThemeColorScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid slideIndex
*/
func TestGetSlidesThemeColorSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesThemeColorScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid password
*/
func TestGetSlidesThemeColorSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesThemeColorScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid folder
*/
func TestGetSlidesThemeColorSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesThemeColorScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid storage
*/
func TestGetSlidesThemeColorSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesThemeColorScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method
*/
func TestGetSlidesThemeFontScheme(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    e := initializeTest("GetSlidesThemeFontScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesThemeFontScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesThemeFontSchemeRequest() GetSlidesThemeFontSchemeRequest {
    var request GetSlidesThemeFontSchemeRequest
    request.name = createTestParamValue("GetSlidesThemeFontScheme", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesThemeFontScheme", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesThemeFontScheme", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesThemeFontScheme", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesThemeFontScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid name
*/
func TestGetSlidesThemeFontSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesThemeFontScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFontSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesThemeFontScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid password
*/
func TestGetSlidesThemeFontSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesThemeFontScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid folder
*/
func TestGetSlidesThemeFontSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesThemeFontScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid storage
*/
func TestGetSlidesThemeFontSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesThemeFontScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method
*/
func TestGetSlidesThemeFormatScheme(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    e := initializeTest("GetSlidesThemeFormatScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.GetSlidesThemeFormatScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createGetSlidesThemeFormatSchemeRequest() GetSlidesThemeFormatSchemeRequest {
    var request GetSlidesThemeFormatSchemeRequest
    request.name = createTestParamValue("GetSlidesThemeFormatScheme", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesThemeFormatScheme", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesThemeFormatScheme", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesThemeFormatScheme", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesThemeFormatScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid name
*/
func TestGetSlidesThemeFormatSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "GetSlidesThemeFormatScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFormatSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "GetSlidesThemeFormatScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid password
*/
func TestGetSlidesThemeFormatSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "GetSlidesThemeFormatScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid folder
*/
func TestGetSlidesThemeFormatSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "GetSlidesThemeFormatScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid storage
*/
func TestGetSlidesThemeFormatSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "GetSlidesThemeFormatScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method
*/
func TestMoveFile(t *testing.T) {
    request := createMoveFileRequest()
    e := initializeTest("MoveFile", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.MoveFile(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createMoveFileRequest() MoveFileRequest {
    var request MoveFileRequest
    request.srcPath = createTestParamValue("MoveFile", "srcPath", "string").(string)
    request.destPath = createTestParamValue("MoveFile", "destPath", "string").(string)
    request.srcStorageName = createTestParamValue("MoveFile", "srcStorageName", "string").(string)
    request.destStorageName = createTestParamValue("MoveFile", "destStorageName", "string").(string)
    request.versionId = createTestParamValue("MoveFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid srcPath
*/
func TestMoveFileInvalidsrcPath(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.srcPath, "MoveFile", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcPath = nullValue
    } else {
        request.srcPath = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "srcPath", request.srcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destPath
*/
func TestMoveFileInvaliddestPath(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.destPath, "MoveFile", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destPath = nullValue
    } else {
        request.destPath = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "destPath", request.destPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid srcStorageName
*/
func TestMoveFileInvalidsrcStorageName(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.srcStorageName, "MoveFile", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcStorageName = nullValue
    } else {
        request.srcStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "srcStorageName", request.srcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destStorageName
*/
func TestMoveFileInvaliddestStorageName(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.destStorageName, "MoveFile", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destStorageName = nullValue
    } else {
        request.destStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "destStorageName", request.destStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid versionId
*/
func TestMoveFileInvalidversionId(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.versionId, "MoveFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.versionId = nullValue
    } else {
        request.versionId = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "versionId", request.versionId, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method
*/
func TestMoveFolder(t *testing.T) {
    request := createMoveFolderRequest()
    e := initializeTest("MoveFolder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.MoveFolder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createMoveFolderRequest() MoveFolderRequest {
    var request MoveFolderRequest
    request.srcPath = createTestParamValue("MoveFolder", "srcPath", "string").(string)
    request.destPath = createTestParamValue("MoveFolder", "destPath", "string").(string)
    request.srcStorageName = createTestParamValue("MoveFolder", "srcStorageName", "string").(string)
    request.destStorageName = createTestParamValue("MoveFolder", "destStorageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid srcPath
*/
func TestMoveFolderInvalidsrcPath(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.srcPath, "MoveFolder", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcPath = nullValue
    } else {
        request.srcPath = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "srcPath", request.srcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destPath
*/
func TestMoveFolderInvaliddestPath(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.destPath, "MoveFolder", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destPath = nullValue
    } else {
        request.destPath = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "destPath", request.destPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid srcStorageName
*/
func TestMoveFolderInvalidsrcStorageName(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.srcStorageName, "MoveFolder", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.srcStorageName = nullValue
    } else {
        request.srcStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "srcStorageName", request.srcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destStorageName
*/
func TestMoveFolderInvaliddestStorageName(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.destStorageName, "MoveFolder", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destStorageName = nullValue
    } else {
        request.destStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "destStorageName", request.destStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method
*/
func TestObjectExists(t *testing.T) {
    request := createObjectExistsRequest()
    e := initializeTest("ObjectExists", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.ObjectExists(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createObjectExistsRequest() ObjectExistsRequest {
    var request ObjectExistsRequest
    request.path = createTestParamValue("ObjectExists", "path", "string").(string)
    request.storageName = createTestParamValue("ObjectExists", "storageName", "string").(string)
    request.versionId = createTestParamValue("ObjectExists", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid path
*/
func TestObjectExistsInvalidpath(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "ObjectExists", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid storageName
*/
func TestObjectExistsInvalidstorageName(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "ObjectExists", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid versionId
*/
func TestObjectExistsInvalidversionId(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.versionId, "ObjectExists", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.versionId = nullValue
    } else {
        request.versionId = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "versionId", request.versionId, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method
*/
func TestPostAddNewParagraph(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    e := initializeTest("PostAddNewParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostAddNewParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testposition := createTestParamValue("PostAddNewParagraph", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid name
*/
func TestPostAddNewParagraphInvalidname(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostAddNewParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid slideIndex
*/
func TestPostAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostAddNewParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid path
*/
func TestPostAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostAddNewParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid shapeIndex
*/
func TestPostAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostAddNewParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid dto
*/
func TestPostAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostAddNewParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PostAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid password
*/
func TestPostAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostAddNewParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid folder
*/
func TestPostAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostAddNewParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid storage
*/
func TestPostAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostAddNewParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid position
*/
func TestPostAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostAddNewParagraph", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method
*/
func TestPostAddNewPortion(t *testing.T) {
    request := createPostAddNewPortionRequest()
    e := initializeTest("PostAddNewPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostAddNewPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testposition := createTestParamValue("PostAddNewPortion", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid name
*/
func TestPostAddNewPortionInvalidname(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostAddNewPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid slideIndex
*/
func TestPostAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostAddNewPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid path
*/
func TestPostAddNewPortionInvalidpath(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostAddNewPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid shapeIndex
*/
func TestPostAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostAddNewPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid paragraphIndex
*/
func TestPostAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PostAddNewPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid dto
*/
func TestPostAddNewPortionInvaliddto(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostAddNewPortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IPortion)
    }

    e := initializeTest("PostAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid password
*/
func TestPostAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostAddNewPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid folder
*/
func TestPostAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostAddNewPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid storage
*/
func TestPostAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostAddNewPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid position
*/
func TestPostAddNewPortionInvalidposition(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostAddNewPortion", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method
*/
func TestPostAddNewShape(t *testing.T) {
    request := createPostAddNewShapeRequest()
    e := initializeTest("PostAddNewShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostAddNewShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testshapeToClone := createTestParamValue("PostAddNewShape", "shapeToClone", "int32")
    switch v := testshapeToClone.(type) { 
    case int32:
        request.shapeToClone = new(int32)
        *request.shapeToClone = v
    }
    testposition := createTestParamValue("PostAddNewShape", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid name
*/
func TestPostAddNewShapeInvalidname(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostAddNewShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid slideIndex
*/
func TestPostAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostAddNewShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid path
*/
func TestPostAddNewShapeInvalidpath(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostAddNewShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid dto
*/
func TestPostAddNewShapeInvaliddto(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostAddNewShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PostAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid password
*/
func TestPostAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostAddNewShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid folder
*/
func TestPostAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostAddNewShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid storage
*/
func TestPostAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostAddNewShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid shapeToClone
*/
func TestPostAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.shapeToClone = new(int32)

    invalidValue := invalidizeTestParamValue(request.shapeToClone, "PostAddNewShape", "shapeToClone", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.shapeToClone = nullValue
    } else {
        *request.shapeToClone = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "shapeToClone", request.shapeToClone, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid position
*/
func TestPostAddNewShapeInvalidposition(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostAddNewShape", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method
*/
func TestPostAddNotesSlide(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    e := initializeTest("PostAddNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostAddNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid name
*/
func TestPostAddNotesSlideInvalidname(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostAddNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid slideIndex
*/
func TestPostAddNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostAddNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid dto
*/
func TestPostAddNotesSlideInvaliddto(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostAddNotesSlide", "dto", "NotesSlide")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(INotesSlide)
    }

    e := initializeTest("PostAddNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid password
*/
func TestPostAddNotesSlideInvalidpassword(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostAddNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid folder
*/
func TestPostAddNotesSlideInvalidfolder(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostAddNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid storage
*/
func TestPostAddNotesSlideInvalidstorage(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostAddNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method
*/
func TestPostCopyLayoutSlideFromSourcePresentation(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidname(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostCopyLayoutSlideFromSourcePresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFrom, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFrom = nullValue
    } else {
        request.cloneFrom = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.cloneFrom, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromPosition, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.cloneFromPosition = nullValue
    } else {
        request.cloneFromPosition = invalidValue.(int32)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromPassword, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFromPassword = nullValue
    } else {
        request.cloneFromPassword = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromStorage, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFromStorage = nullValue
    } else {
        request.cloneFromStorage = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostCopyLayoutSlideFromSourcePresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostCopyLayoutSlideFromSourcePresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostCopyLayoutSlideFromSourcePresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method
*/
func TestPostCopyMasterSlideFromSourcePresentation(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testapplyToAll := createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "applyToAll", "bool")
    switch v := testapplyToAll.(type) { 
    case bool:
        request.applyToAll = new(bool)
        *request.applyToAll = v
    }
    request.password = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "password", "string").(string)
    request.folder = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "folder", "string").(string)
    request.storage = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidname(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostCopyMasterSlideFromSourcePresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFrom, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFrom = nullValue
    } else {
        request.cloneFrom = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.cloneFrom, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromPosition, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.cloneFromPosition = nullValue
    } else {
        request.cloneFromPosition = invalidValue.(int32)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromPassword, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFromPassword = nullValue
    } else {
        request.cloneFromPassword = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.cloneFromStorage, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.cloneFromStorage = nullValue
    } else {
        request.cloneFromStorage = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid applyToAll
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidapplyToAll(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.applyToAll = new(bool)

    invalidValue := invalidizeTestParamValue(request.applyToAll, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.applyToAll = nullValue
    } else {
        *request.applyToAll = invalidValue.(bool)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.applyToAll)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.applyToAll, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostCopyMasterSlideFromSourcePresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostCopyMasterSlideFromSourcePresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostCopyMasterSlideFromSourcePresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method
*/
func TestPostNotesSlideAddNewParagraph(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    e := initializeTest("PostNotesSlideAddNewParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostNotesSlideAddNewParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testposition := createTestParamValue("PostNotesSlideAddNewParagraph", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid name
*/
func TestPostNotesSlideAddNewParagraphInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostNotesSlideAddNewParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid slideIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostNotesSlideAddNewParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid path
*/
func TestPostNotesSlideAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostNotesSlideAddNewParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostNotesSlideAddNewParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid dto
*/
func TestPostNotesSlideAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostNotesSlideAddNewParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid password
*/
func TestPostNotesSlideAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostNotesSlideAddNewParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid folder
*/
func TestPostNotesSlideAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostNotesSlideAddNewParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid storage
*/
func TestPostNotesSlideAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostNotesSlideAddNewParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid position
*/
func TestPostNotesSlideAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostNotesSlideAddNewParagraph", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method
*/
func TestPostNotesSlideAddNewPortion(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    e := initializeTest("PostNotesSlideAddNewPortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostNotesSlideAddNewPortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testposition := createTestParamValue("PostNotesSlideAddNewPortion", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid name
*/
func TestPostNotesSlideAddNewPortionInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostNotesSlideAddNewPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid slideIndex
*/
func TestPostNotesSlideAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostNotesSlideAddNewPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid path
*/
func TestPostNotesSlideAddNewPortionInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostNotesSlideAddNewPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostNotesSlideAddNewPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid paragraphIndex
*/
func TestPostNotesSlideAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PostNotesSlideAddNewPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid dto
*/
func TestPostNotesSlideAddNewPortionInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostNotesSlideAddNewPortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IPortion)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid password
*/
func TestPostNotesSlideAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostNotesSlideAddNewPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid folder
*/
func TestPostNotesSlideAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostNotesSlideAddNewPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid storage
*/
func TestPostNotesSlideAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostNotesSlideAddNewPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid position
*/
func TestPostNotesSlideAddNewPortionInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostNotesSlideAddNewPortion", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method
*/
func TestPostNotesSlideAddNewShape(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    e := initializeTest("PostNotesSlideAddNewShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostNotesSlideAddNewShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    testshapeToClone := createTestParamValue("PostNotesSlideAddNewShape", "shapeToClone", "int32")
    switch v := testshapeToClone.(type) { 
    case int32:
        request.shapeToClone = new(int32)
        *request.shapeToClone = v
    }
    testposition := createTestParamValue("PostNotesSlideAddNewShape", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    return request
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid name
*/
func TestPostNotesSlideAddNewShapeInvalidname(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostNotesSlideAddNewShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid slideIndex
*/
func TestPostNotesSlideAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostNotesSlideAddNewShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid path
*/
func TestPostNotesSlideAddNewShapeInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostNotesSlideAddNewShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid dto
*/
func TestPostNotesSlideAddNewShapeInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PostNotesSlideAddNewShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid password
*/
func TestPostNotesSlideAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostNotesSlideAddNewShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid folder
*/
func TestPostNotesSlideAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostNotesSlideAddNewShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid storage
*/
func TestPostNotesSlideAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostNotesSlideAddNewShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid shapeToClone
*/
func TestPostNotesSlideAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.shapeToClone = new(int32)

    invalidValue := invalidizeTestParamValue(request.shapeToClone, "PostNotesSlideAddNewShape", "shapeToClone", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.shapeToClone = nullValue
    } else {
        *request.shapeToClone = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "shapeToClone", request.shapeToClone, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid position
*/
func TestPostNotesSlideAddNewShapeInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostNotesSlideAddNewShape", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method
*/
func TestPostNotesSlideShapeSaveAs(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    e := initializeTest("PostNotesSlideShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostNotesSlideShapeSaveAs(request)
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
    request.path = createTestParamValue("PostNotesSlideShapeSaveAs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostNotesSlideShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PostNotesSlideShapeSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostNotesSlideShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PostNotesSlideShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostNotesSlideShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostNotesSlideShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PostNotesSlideShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.scaleX = new(float64)
        *request.scaleX = v
    }
    testscaleY := createTestParamValue("PostNotesSlideShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.scaleY = new(float64)
        *request.scaleY = v
    }
    request.bounds = createTestParamValue("PostNotesSlideShapeSaveAs", "bounds", "string").(string)
    request.fontsFolder = createTestParamValue("PostNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid name
*/
func TestPostNotesSlideShapeSaveAsInvalidname(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostNotesSlideShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostNotesSlideShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid path
*/
func TestPostNotesSlideShapeSaveAsInvalidpath(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostNotesSlideShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostNotesSlideShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid format
*/
func TestPostNotesSlideShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostNotesSlideShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid options
*/
func TestPostNotesSlideShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PostNotesSlideShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid password
*/
func TestPostNotesSlideShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostNotesSlideShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid folder
*/
func TestPostNotesSlideShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostNotesSlideShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid storage
*/
func TestPostNotesSlideShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostNotesSlideShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleX, "PostNotesSlideShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleX = nullValue
    } else {
        *request.scaleX = invalidValue.(float64)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleX", request.scaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleY, "PostNotesSlideShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleY = nullValue
    } else {
        *request.scaleY = invalidValue.(float64)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleY", request.scaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPostNotesSlideShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.bounds, "PostNotesSlideShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.bounds = nullValue
    } else {
        request.bounds = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "bounds", request.bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPostNotesSlideShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostNotesSlideShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method
*/
func TestPostPresentationMerge(t *testing.T) {
    request := createPostPresentationMergeRequest()
    e := initializeTest("PostPresentationMerge", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostPresentationMerge(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid name
*/
func TestPostPresentationMergeInvalidname(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostPresentationMerge", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid request
*/
func TestPostPresentationMergeInvalidrequest(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.request, "PostPresentationMerge", "request", "PresentationsMergeRequest")
    if (invalidValue == nil) {
        request.request = nil
    } else {
        request.request = invalidValue.(IPresentationsMergeRequest)
    }

    e := initializeTest("PostPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "request", request.request, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid password
*/
func TestPostPresentationMergeInvalidpassword(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostPresentationMerge", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid storage
*/
func TestPostPresentationMergeInvalidstorage(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostPresentationMerge", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid folder
*/
func TestPostPresentationMergeInvalidfolder(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostPresentationMerge", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method
*/
func TestPostShapeSaveAs(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    e := initializeTest("PostShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostShapeSaveAs(request)
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
    request.path = createTestParamValue("PostShapeSaveAs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PostShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PostShapeSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PostShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PostShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.scaleX = new(float64)
        *request.scaleX = v
    }
    testscaleY := createTestParamValue("PostShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.scaleY = new(float64)
        *request.scaleY = v
    }
    request.bounds = createTestParamValue("PostShapeSaveAs", "bounds", "string").(string)
    request.fontsFolder = createTestParamValue("PostShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid name
*/
func TestPostShapeSaveAsInvalidname(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid slideIndex
*/
func TestPostShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid path
*/
func TestPostShapeSaveAsInvalidpath(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PostShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid shapeIndex
*/
func TestPostShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PostShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid format
*/
func TestPostShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid options
*/
func TestPostShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PostShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PostShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid password
*/
func TestPostShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid folder
*/
func TestPostShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid storage
*/
func TestPostShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleX
*/
func TestPostShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleX, "PostShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleX = nullValue
    } else {
        *request.scaleX = invalidValue.(float64)
    }

    e := initializeTest("PostShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "scaleX", request.scaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleY
*/
func TestPostShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleY, "PostShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleY = nullValue
    } else {
        *request.scaleY = invalidValue.(float64)
    }

    e := initializeTest("PostShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "scaleY", request.scaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid bounds
*/
func TestPostShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.bounds, "PostShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.bounds = nullValue
    } else {
        request.bounds = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "bounds", request.bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid fontsFolder
*/
func TestPostShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method
*/
func TestPostSlideAnimationEffect(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()
    e := initializeTest("PostSlideAnimationEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlideAnimationEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlideAnimationEffectRequest() PostSlideAnimationEffectRequest {
    var request PostSlideAnimationEffectRequest
    request.name = createTestParamValue("PostSlideAnimationEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.effect = createTestParamValue("PostSlideAnimationEffect", "effect", "Effect").(IEffect)
    request.password = createTestParamValue("PostSlideAnimationEffect", "password", "string").(string)
    request.folder = createTestParamValue("PostSlideAnimationEffect", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid name
*/
func TestPostSlideAnimationEffectInvalidname(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid slideIndex
*/
func TestPostSlideAnimationEffectInvalidslideIndex(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid effect
*/
func TestPostSlideAnimationEffectInvalideffect(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effect, "PostSlideAnimationEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.effect = nil
    } else {
        request.effect = invalidValue.(IEffect)
    }

    e := initializeTest("PostSlideAnimationEffect", "effect", request.effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "effect", request.effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid password
*/
func TestPostSlideAnimationEffectInvalidpassword(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid folder
*/
func TestPostSlideAnimationEffectInvalidfolder(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid storage
*/
func TestPostSlideAnimationEffectInvalidstorage(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method
*/
func TestPostSlideAnimationInteractiveSequence(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()
    e := initializeTest("PostSlideAnimationInteractiveSequence", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlideAnimationInteractiveSequence(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlideAnimationInteractiveSequenceRequest() PostSlideAnimationInteractiveSequenceRequest {
    var request PostSlideAnimationInteractiveSequenceRequest
    request.name = createTestParamValue("PostSlideAnimationInteractiveSequence", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlideAnimationInteractiveSequence", "slideIndex", "int32").(int32)
    request.sequence = createTestParamValue("PostSlideAnimationInteractiveSequence", "sequence", "InteractiveSequence").(IInteractiveSequence)
    request.password = createTestParamValue("PostSlideAnimationInteractiveSequence", "password", "string").(string)
    request.folder = createTestParamValue("PostSlideAnimationInteractiveSequence", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlideAnimationInteractiveSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid name
*/
func TestPostSlideAnimationInteractiveSequenceInvalidname(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlideAnimationInteractiveSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid slideIndex
*/
func TestPostSlideAnimationInteractiveSequenceInvalidslideIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlideAnimationInteractiveSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid sequence
*/
func TestPostSlideAnimationInteractiveSequenceInvalidsequence(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.sequence, "PostSlideAnimationInteractiveSequence", "sequence", "InteractiveSequence")
    if (invalidValue == nil) {
        request.sequence = nil
    } else {
        request.sequence = invalidValue.(IInteractiveSequence)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "sequence", request.sequence)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "sequence", request.sequence, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid password
*/
func TestPostSlideAnimationInteractiveSequenceInvalidpassword(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlideAnimationInteractiveSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid folder
*/
func TestPostSlideAnimationInteractiveSequenceInvalidfolder(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlideAnimationInteractiveSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid storage
*/
func TestPostSlideAnimationInteractiveSequenceInvalidstorage(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlideAnimationInteractiveSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method
*/
func TestPostSlideAnimationInteractiveSequenceEffect(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()
    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlideAnimationInteractiveSequenceEffectRequest() PostSlideAnimationInteractiveSequenceEffectRequest {
    var request PostSlideAnimationInteractiveSequenceEffectRequest
    request.name = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.sequenceIndex = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.effect = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "effect", "Effect").(IEffect)
    request.password = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.folder = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidname(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidslideIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidsequenceIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.sequenceIndex, "PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.sequenceIndex = nullValue
    } else {
        request.sequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid effect
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalideffect(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effect, "PostSlideAnimationInteractiveSequenceEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.effect = nil
    } else {
        request.effect = invalidValue.(IEffect)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "effect", request.effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "effect", request.effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidpassword(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidfolder(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidstorage(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method
*/
func TestPostSlideSaveAs(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    e := initializeTest("PostSlideSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlideSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostSlideSaveAsRequest() PostSlideSaveAsRequest {
    var request PostSlideSaveAsRequest
    request.name = createTestParamValue("PostSlideSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlideSaveAs", "slideIndex", "int32").(int32)
    request.format = createTestParamValue("PostSlideSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostSlideSaveAs", "options", "ExportOptions").(IExportOptions)
    testwidth := createTestParamValue("PostSlideSaveAs", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.width = new(int32)
        *request.width = v
    }
    testheight := createTestParamValue("PostSlideSaveAs", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.height = new(int32)
        *request.height = v
    }
    request.password = createTestParamValue("PostSlideSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostSlideSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlideSaveAs", "storage", "string").(string)
    request.fontsFolder = createTestParamValue("PostSlideSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid name
*/
func TestPostSlideSaveAsInvalidname(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlideSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid slideIndex
*/
func TestPostSlideSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlideSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid format
*/
func TestPostSlideSaveAsInvalidformat(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostSlideSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid options
*/
func TestPostSlideSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PostSlideSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlideSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid width
*/
func TestPostSlideSaveAsInvalidwidth(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.width = new(int32)

    invalidValue := invalidizeTestParamValue(request.width, "PostSlideSaveAs", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.width = nullValue
    } else {
        *request.width = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "width", request.width, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid height
*/
func TestPostSlideSaveAsInvalidheight(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.height = new(int32)

    invalidValue := invalidizeTestParamValue(request.height, "PostSlideSaveAs", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.height = nullValue
    } else {
        *request.height = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "height", request.height, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid password
*/
func TestPostSlideSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlideSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid folder
*/
func TestPostSlideSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlideSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid storage
*/
func TestPostSlideSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlideSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid fontsFolder
*/
func TestPostSlideSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostSlideSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method
*/
func TestPostSlidesAdd(t *testing.T) {
    request := createPostSlidesAddRequest()
    e := initializeTest("PostSlidesAdd", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesAdd(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesAddRequest() PostSlidesAddRequest {
    var request PostSlidesAddRequest
    request.name = createTestParamValue("PostSlidesAdd", "name", "string").(string)
    testposition := createTestParamValue("PostSlidesAdd", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    request.password = createTestParamValue("PostSlidesAdd", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesAdd", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesAdd", "storage", "string").(string)
    request.layoutAlias = createTestParamValue("PostSlidesAdd", "layoutAlias", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid name
*/
func TestPostSlidesAddInvalidname(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesAdd", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid position
*/
func TestPostSlidesAddInvalidposition(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostSlidesAdd", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesAdd", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid password
*/
func TestPostSlidesAddInvalidpassword(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesAdd", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid folder
*/
func TestPostSlidesAddInvalidfolder(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesAdd", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid storage
*/
func TestPostSlidesAddInvalidstorage(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesAdd", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid layoutAlias
*/
func TestPostSlidesAddInvalidlayoutAlias(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.layoutAlias, "PostSlidesAdd", "layoutAlias", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.layoutAlias = nullValue
    } else {
        request.layoutAlias = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "layoutAlias", request.layoutAlias)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "layoutAlias", request.layoutAlias, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method
*/
func TestPostSlidesConvert(t *testing.T) {
    request := createPostSlidesConvertRequest()
    e := initializeTest("PostSlidesConvert", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlidesConvert(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostSlidesConvertRequest() PostSlidesConvertRequest {
    var request PostSlidesConvertRequest
    request.format = createTestParamValue("PostSlidesConvert", "format", "string").(string)
    request.document = createTestParamValue("PostSlidesConvert", "document", "[]byte").([]byte)
    request.password = createTestParamValue("PostSlidesConvert", "password", "string").(string)
    request.fontsFolder = createTestParamValue("PostSlidesConvert", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid format
*/
func TestPostSlidesConvertInvalidformat(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostSlidesConvert", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid document
*/
func TestPostSlidesConvertInvaliddocument(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.document, "PostSlidesConvert", "document", "[]byte")
    if (invalidValue == nil) {
        request.document = nil
    } else {
        request.document = invalidValue.([]byte)
    }

    e := initializeTest("PostSlidesConvert", "document", request.document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "document", request.document, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid password
*/
func TestPostSlidesConvertInvalidpassword(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesConvert", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid fontsFolder
*/
func TestPostSlidesConvertInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostSlidesConvert", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method
*/
func TestPostSlidesCopy(t *testing.T) {
    request := createPostSlidesCopyRequest()
    e := initializeTest("PostSlidesCopy", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesCopy(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesCopyRequest() PostSlidesCopyRequest {
    var request PostSlidesCopyRequest
    request.name = createTestParamValue("PostSlidesCopy", "name", "string").(string)
    request.slideToCopy = createTestParamValue("PostSlidesCopy", "slideToCopy", "int32").(int32)
    testposition := createTestParamValue("PostSlidesCopy", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.position = new(int32)
        *request.position = v
    }
    request.source = createTestParamValue("PostSlidesCopy", "source", "string").(string)
    request.sourcePassword = createTestParamValue("PostSlidesCopy", "sourcePassword", "string").(string)
    request.sourceStorage = createTestParamValue("PostSlidesCopy", "sourceStorage", "string").(string)
    request.password = createTestParamValue("PostSlidesCopy", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesCopy", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesCopy", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid name
*/
func TestPostSlidesCopyInvalidname(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesCopy", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid slideToCopy
*/
func TestPostSlidesCopyInvalidslideToCopy(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.slideToCopy, "PostSlidesCopy", "slideToCopy", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideToCopy = nullValue
    } else {
        request.slideToCopy = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesCopy", "slideToCopy", request.slideToCopy)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "slideToCopy", request.slideToCopy, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid position
*/
func TestPostSlidesCopyInvalidposition(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.position = new(int32)

    invalidValue := invalidizeTestParamValue(request.position, "PostSlidesCopy", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.position = nullValue
    } else {
        *request.position = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesCopy", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "position", request.position, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid source
*/
func TestPostSlidesCopyInvalidsource(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.source, "PostSlidesCopy", "source", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.source = nullValue
    } else {
        request.source = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "source", request.source)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "source", request.source, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourcePassword
*/
func TestPostSlidesCopyInvalidsourcePassword(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.sourcePassword, "PostSlidesCopy", "sourcePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sourcePassword = nullValue
    } else {
        request.sourcePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "sourcePassword", request.sourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "sourcePassword", request.sourcePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourceStorage
*/
func TestPostSlidesCopyInvalidsourceStorage(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.sourceStorage, "PostSlidesCopy", "sourceStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sourceStorage = nullValue
    } else {
        request.sourceStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "sourceStorage", request.sourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "sourceStorage", request.sourceStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid password
*/
func TestPostSlidesCopyInvalidpassword(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesCopy", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid folder
*/
func TestPostSlidesCopyInvalidfolder(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesCopy", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid storage
*/
func TestPostSlidesCopyInvalidstorage(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesCopy", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method
*/
func TestPostSlidesDocument(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    e := initializeTest("PostSlidesDocument", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesDocument(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesDocumentRequest() PostSlidesDocumentRequest {
    var request PostSlidesDocumentRequest
    request.name = createTestParamValue("PostSlidesDocument", "name", "string").(string)
    request.data = createTestParamValue("PostSlidesDocument", "data", "[]byte").([]byte)
    request.inputPassword = createTestParamValue("PostSlidesDocument", "inputPassword", "string").(string)
    request.password = createTestParamValue("PostSlidesDocument", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesDocument", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesDocument", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid name
*/
func TestPostSlidesDocumentInvalidname(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesDocument", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid data
*/
func TestPostSlidesDocumentInvaliddata(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.data, "PostSlidesDocument", "data", "[]byte")
    if (invalidValue == nil) {
        request.data = nil
    } else {
        request.data = invalidValue.([]byte)
    }

    e := initializeTest("PostSlidesDocument", "data", request.data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "data", request.data, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid inputPassword
*/
func TestPostSlidesDocumentInvalidinputPassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.inputPassword, "PostSlidesDocument", "inputPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.inputPassword = nullValue
    } else {
        request.inputPassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "inputPassword", request.inputPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "inputPassword", request.inputPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid password
*/
func TestPostSlidesDocumentInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesDocument", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid storage
*/
func TestPostSlidesDocumentInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesDocument", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid folder
*/
func TestPostSlidesDocumentInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesDocument", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method
*/
func TestPostSlidesDocumentFromHtml(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()
    e := initializeTest("PostSlidesDocumentFromHtml", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesDocumentFromHtml(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesDocumentFromHtmlRequest() PostSlidesDocumentFromHtmlRequest {
    var request PostSlidesDocumentFromHtmlRequest
    request.name = createTestParamValue("PostSlidesDocumentFromHtml", "name", "string").(string)
    request.html = createTestParamValue("PostSlidesDocumentFromHtml", "html", "string").(string)
    request.password = createTestParamValue("PostSlidesDocumentFromHtml", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesDocumentFromHtml", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesDocumentFromHtml", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid name
*/
func TestPostSlidesDocumentFromHtmlInvalidname(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesDocumentFromHtml", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid html
*/
func TestPostSlidesDocumentFromHtmlInvalidhtml(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.html, "PostSlidesDocumentFromHtml", "html", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.html = nullValue
    } else {
        request.html = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "html", request.html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "html", request.html, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid password
*/
func TestPostSlidesDocumentFromHtmlInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesDocumentFromHtml", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid storage
*/
func TestPostSlidesDocumentFromHtmlInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesDocumentFromHtml", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid folder
*/
func TestPostSlidesDocumentFromHtmlInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesDocumentFromHtml", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method
*/
func TestPostSlidesDocumentFromSource(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    e := initializeTest("PostSlidesDocumentFromSource", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesDocumentFromSource(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesDocumentFromSourceRequest() PostSlidesDocumentFromSourceRequest {
    var request PostSlidesDocumentFromSourceRequest
    request.name = createTestParamValue("PostSlidesDocumentFromSource", "name", "string").(string)
    request.sourcePath = createTestParamValue("PostSlidesDocumentFromSource", "sourcePath", "string").(string)
    request.sourcePassword = createTestParamValue("PostSlidesDocumentFromSource", "sourcePassword", "string").(string)
    request.sourceStorage = createTestParamValue("PostSlidesDocumentFromSource", "sourceStorage", "string").(string)
    request.password = createTestParamValue("PostSlidesDocumentFromSource", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesDocumentFromSource", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesDocumentFromSource", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid name
*/
func TestPostSlidesDocumentFromSourceInvalidname(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesDocumentFromSource", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePath
*/
func TestPostSlidesDocumentFromSourceInvalidsourcePath(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.sourcePath, "PostSlidesDocumentFromSource", "sourcePath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sourcePath = nullValue
    } else {
        request.sourcePath = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePath", request.sourcePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourcePath", request.sourcePath, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePassword
*/
func TestPostSlidesDocumentFromSourceInvalidsourcePassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.sourcePassword, "PostSlidesDocumentFromSource", "sourcePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sourcePassword = nullValue
    } else {
        request.sourcePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePassword", request.sourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourcePassword", request.sourcePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourceStorage
*/
func TestPostSlidesDocumentFromSourceInvalidsourceStorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.sourceStorage, "PostSlidesDocumentFromSource", "sourceStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sourceStorage = nullValue
    } else {
        request.sourceStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourceStorage", request.sourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourceStorage", request.sourceStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid password
*/
func TestPostSlidesDocumentFromSourceInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesDocumentFromSource", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid storage
*/
func TestPostSlidesDocumentFromSourceInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesDocumentFromSource", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid folder
*/
func TestPostSlidesDocumentFromSourceInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesDocumentFromSource", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method
*/
func TestPostSlidesDocumentFromTemplate(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    e := initializeTest("PostSlidesDocumentFromTemplate", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesDocumentFromTemplate(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesDocumentFromTemplateRequest() PostSlidesDocumentFromTemplateRequest {
    var request PostSlidesDocumentFromTemplateRequest
    request.name = createTestParamValue("PostSlidesDocumentFromTemplate", "name", "string").(string)
    request.templatePath = createTestParamValue("PostSlidesDocumentFromTemplate", "templatePath", "string").(string)
    request.data = createTestParamValue("PostSlidesDocumentFromTemplate", "data", "string").(string)
    request.templatePassword = createTestParamValue("PostSlidesDocumentFromTemplate", "templatePassword", "string").(string)
    request.templateStorage = createTestParamValue("PostSlidesDocumentFromTemplate", "templateStorage", "string").(string)
    testisImageDataEmbedded := createTestParamValue("PostSlidesDocumentFromTemplate", "isImageDataEmbedded", "bool")
    switch v := testisImageDataEmbedded.(type) { 
    case bool:
        request.isImageDataEmbedded = new(bool)
        *request.isImageDataEmbedded = v
    }
    request.password = createTestParamValue("PostSlidesDocumentFromTemplate", "password", "string").(string)
    request.storage = createTestParamValue("PostSlidesDocumentFromTemplate", "storage", "string").(string)
    request.folder = createTestParamValue("PostSlidesDocumentFromTemplate", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid name
*/
func TestPostSlidesDocumentFromTemplateInvalidname(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesDocumentFromTemplate", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePath
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplatePath(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.templatePath, "PostSlidesDocumentFromTemplate", "templatePath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.templatePath = nullValue
    } else {
        request.templatePath = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePath", request.templatePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePath", request.templatePath, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid data
*/
func TestPostSlidesDocumentFromTemplateInvaliddata(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.data, "PostSlidesDocumentFromTemplate", "data", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.data = nullValue
    } else {
        request.data = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "data", request.data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "data", request.data, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePassword
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplatePassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.templatePassword, "PostSlidesDocumentFromTemplate", "templatePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.templatePassword = nullValue
    } else {
        request.templatePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePassword", request.templatePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePassword", request.templatePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templateStorage
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplateStorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.templateStorage, "PostSlidesDocumentFromTemplate", "templateStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.templateStorage = nullValue
    } else {
        request.templateStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templateStorage", request.templateStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templateStorage", request.templateStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid isImageDataEmbedded
*/
func TestPostSlidesDocumentFromTemplateInvalidisImageDataEmbedded(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.isImageDataEmbedded = new(bool)

    invalidValue := invalidizeTestParamValue(request.isImageDataEmbedded, "PostSlidesDocumentFromTemplate", "isImageDataEmbedded", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.isImageDataEmbedded = nullValue
    } else {
        *request.isImageDataEmbedded = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "isImageDataEmbedded", request.isImageDataEmbedded)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "isImageDataEmbedded", request.isImageDataEmbedded, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid password
*/
func TestPostSlidesDocumentFromTemplateInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesDocumentFromTemplate", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid storage
*/
func TestPostSlidesDocumentFromTemplateInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesDocumentFromTemplate", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid folder
*/
func TestPostSlidesDocumentFromTemplateInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesDocumentFromTemplate", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Performs slides pipeline.
   Test for SlidesApi.PostSlidesPipeline method
*/
func TestPostSlidesPipeline(t *testing.T) {
    request := createPostSlidesPipelineRequest()
    e := initializeTest("PostSlidesPipeline", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlidesPipeline(request)
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

/* SlidesApiServiceTests Performs slides pipeline.
   Test for SlidesApi.PostSlidesPipeline method with invalid pipeline
*/
func TestPostSlidesPipelineInvalidpipeline(t *testing.T) {
    request := createPostSlidesPipelineRequest()

    invalidValue := invalidizeTestParamValue(request.pipeline, "PostSlidesPipeline", "pipeline", "Pipeline")
    if (invalidValue == nil) {
        request.pipeline = nil
    } else {
        request.pipeline = invalidValue.(IPipeline)
    }

    e := initializeTest("PostSlidesPipeline", "pipeline", request.pipeline)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPipeline(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPipeline", "pipeline", request.pipeline, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method
*/
func TestPostSlidesPresentationReplaceText(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    e := initializeTest("PostSlidesPresentationReplaceText", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesPresentationReplaceText(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesPresentationReplaceTextRequest() PostSlidesPresentationReplaceTextRequest {
    var request PostSlidesPresentationReplaceTextRequest
    request.name = createTestParamValue("PostSlidesPresentationReplaceText", "name", "string").(string)
    request.oldValue = createTestParamValue("PostSlidesPresentationReplaceText", "oldValue", "string").(string)
    request.newValue = createTestParamValue("PostSlidesPresentationReplaceText", "newValue", "string").(string)
    testignoreCase := createTestParamValue("PostSlidesPresentationReplaceText", "ignoreCase", "bool")
    switch v := testignoreCase.(type) { 
    case bool:
        request.ignoreCase = new(bool)
        *request.ignoreCase = v
    }
    request.password = createTestParamValue("PostSlidesPresentationReplaceText", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesPresentationReplaceText", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesPresentationReplaceText", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid name
*/
func TestPostSlidesPresentationReplaceTextInvalidname(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesPresentationReplaceText", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid oldValue
*/
func TestPostSlidesPresentationReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.oldValue, "PostSlidesPresentationReplaceText", "oldValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.oldValue = nullValue
    } else {
        request.oldValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "oldValue", request.oldValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid newValue
*/
func TestPostSlidesPresentationReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.newValue, "PostSlidesPresentationReplaceText", "newValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.newValue = nullValue
    } else {
        request.newValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "newValue", request.newValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid ignoreCase
*/
func TestPostSlidesPresentationReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.ignoreCase = new(bool)

    invalidValue := invalidizeTestParamValue(request.ignoreCase, "PostSlidesPresentationReplaceText", "ignoreCase", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.ignoreCase = nullValue
    } else {
        *request.ignoreCase = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "ignoreCase", request.ignoreCase, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid password
*/
func TestPostSlidesPresentationReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesPresentationReplaceText", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid folder
*/
func TestPostSlidesPresentationReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesPresentationReplaceText", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid storage
*/
func TestPostSlidesPresentationReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesPresentationReplaceText", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method
*/
func TestPostSlidesReorder(t *testing.T) {
    request := createPostSlidesReorderRequest()
    e := initializeTest("PostSlidesReorder", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesReorder(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesReorderRequest() PostSlidesReorderRequest {
    var request PostSlidesReorderRequest
    request.name = createTestParamValue("PostSlidesReorder", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlidesReorder", "slideIndex", "int32").(int32)
    request.newPosition = createTestParamValue("PostSlidesReorder", "newPosition", "int32").(int32)
    request.password = createTestParamValue("PostSlidesReorder", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesReorder", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesReorder", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid name
*/
func TestPostSlidesReorderInvalidname(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesReorder", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid slideIndex
*/
func TestPostSlidesReorderInvalidslideIndex(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlidesReorder", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesReorder", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid newPosition
*/
func TestPostSlidesReorderInvalidnewPosition(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.newPosition, "PostSlidesReorder", "newPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.newPosition = nullValue
    } else {
        request.newPosition = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesReorder", "newPosition", request.newPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "newPosition", request.newPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid password
*/
func TestPostSlidesReorderInvalidpassword(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesReorder", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid folder
*/
func TestPostSlidesReorderInvalidfolder(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesReorder", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid storage
*/
func TestPostSlidesReorderInvalidstorage(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesReorder", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method
*/
func TestPostSlidesReorderMany(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    e := initializeTest("PostSlidesReorderMany", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesReorderMany(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesReorderManyRequest() PostSlidesReorderManyRequest {
    var request PostSlidesReorderManyRequest
    request.name = createTestParamValue("PostSlidesReorderMany", "name", "string").(string)
    request.oldPositions = createTestParamValue("PostSlidesReorderMany", "oldPositions", "[]int32").([]int32)
    request.newPositions = createTestParamValue("PostSlidesReorderMany", "newPositions", "[]int32").([]int32)
    request.password = createTestParamValue("PostSlidesReorderMany", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesReorderMany", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesReorderMany", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid name
*/
func TestPostSlidesReorderManyInvalidname(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesReorderMany", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid oldPositions
*/
func TestPostSlidesReorderManyInvalidoldPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.oldPositions, "PostSlidesReorderMany", "oldPositions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.oldPositions = nullValue
    } else {
        request.oldPositions = invalidValue.([]int32)
    }

    e := initializeTest("PostSlidesReorderMany", "oldPositions", request.oldPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "oldPositions", request.oldPositions, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid newPositions
*/
func TestPostSlidesReorderManyInvalidnewPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.newPositions, "PostSlidesReorderMany", "newPositions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.newPositions = nullValue
    } else {
        request.newPositions = invalidValue.([]int32)
    }

    e := initializeTest("PostSlidesReorderMany", "newPositions", request.newPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "newPositions", request.newPositions, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid password
*/
func TestPostSlidesReorderManyInvalidpassword(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesReorderMany", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid folder
*/
func TestPostSlidesReorderManyInvalidfolder(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesReorderMany", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid storage
*/
func TestPostSlidesReorderManyInvalidstorage(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesReorderMany", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method
*/
func TestPostSlidesSaveAs(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    e := initializeTest("PostSlidesSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlidesSaveAs(request)
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
    request.fontsFolder = createTestParamValue("PostSlidesSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid name
*/
func TestPostSlidesSaveAsInvalidname(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid format
*/
func TestPostSlidesSaveAsInvalidformat(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostSlidesSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid options
*/
func TestPostSlidesSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PostSlidesSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlidesSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid password
*/
func TestPostSlidesSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid storage
*/
func TestPostSlidesSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid folder
*/
func TestPostSlidesSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid fontsFolder
*/
func TestPostSlidesSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostSlidesSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method
*/
func TestPostSlidesSetDocumentProperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    e := initializeTest("PostSlidesSetDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesSetDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesSetDocumentPropertiesRequest() PostSlidesSetDocumentPropertiesRequest {
    var request PostSlidesSetDocumentPropertiesRequest
    request.name = createTestParamValue("PostSlidesSetDocumentProperties", "name", "string").(string)
    request.properties = createTestParamValue("PostSlidesSetDocumentProperties", "properties", "DocumentProperties").(IDocumentProperties)
    request.password = createTestParamValue("PostSlidesSetDocumentProperties", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesSetDocumentProperties", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesSetDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid name
*/
func TestPostSlidesSetDocumentPropertiesInvalidname(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesSetDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid properties
*/
func TestPostSlidesSetDocumentPropertiesInvalidproperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.properties, "PostSlidesSetDocumentProperties", "properties", "DocumentProperties")
    if (invalidValue == nil) {
        request.properties = nil
    } else {
        request.properties = invalidValue.(IDocumentProperties)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "properties", request.properties)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "properties", request.properties, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid password
*/
func TestPostSlidesSetDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesSetDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid folder
*/
func TestPostSlidesSetDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesSetDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid storage
*/
func TestPostSlidesSetDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesSetDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method
*/
func TestPostSlidesSlideReplaceText(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    e := initializeTest("PostSlidesSlideReplaceText", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesSlideReplaceText(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPostSlidesSlideReplaceTextRequest() PostSlidesSlideReplaceTextRequest {
    var request PostSlidesSlideReplaceTextRequest
    request.name = createTestParamValue("PostSlidesSlideReplaceText", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlidesSlideReplaceText", "slideIndex", "int32").(int32)
    request.oldValue = createTestParamValue("PostSlidesSlideReplaceText", "oldValue", "string").(string)
    request.newValue = createTestParamValue("PostSlidesSlideReplaceText", "newValue", "string").(string)
    testignoreCase := createTestParamValue("PostSlidesSlideReplaceText", "ignoreCase", "bool")
    switch v := testignoreCase.(type) { 
    case bool:
        request.ignoreCase = new(bool)
        *request.ignoreCase = v
    }
    request.password = createTestParamValue("PostSlidesSlideReplaceText", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesSlideReplaceText", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesSlideReplaceText", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid name
*/
func TestPostSlidesSlideReplaceTextInvalidname(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesSlideReplaceText", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid slideIndex
*/
func TestPostSlidesSlideReplaceTextInvalidslideIndex(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PostSlidesSlideReplaceText", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid oldValue
*/
func TestPostSlidesSlideReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.oldValue, "PostSlidesSlideReplaceText", "oldValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.oldValue = nullValue
    } else {
        request.oldValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "oldValue", request.oldValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid newValue
*/
func TestPostSlidesSlideReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.newValue, "PostSlidesSlideReplaceText", "newValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.newValue = nullValue
    } else {
        request.newValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "newValue", request.newValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid ignoreCase
*/
func TestPostSlidesSlideReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.ignoreCase = new(bool)

    invalidValue := invalidizeTestParamValue(request.ignoreCase, "PostSlidesSlideReplaceText", "ignoreCase", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.ignoreCase = nullValue
    } else {
        *request.ignoreCase = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "ignoreCase", request.ignoreCase, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid password
*/
func TestPostSlidesSlideReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesSlideReplaceText", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid folder
*/
func TestPostSlidesSlideReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesSlideReplaceText", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid storage
*/
func TestPostSlidesSlideReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesSlideReplaceText", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method
*/
func TestPostSlidesSplit(t *testing.T) {
    request := createPostSlidesSplitRequest()
    e := initializeTest("PostSlidesSplit", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PostSlidesSplit(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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
    request.fontsFolder = createTestParamValue("PostSlidesSplit", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid name
*/
func TestPostSlidesSplitInvalidname(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PostSlidesSplit", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid options
*/
func TestPostSlidesSplitInvalidoptions(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PostSlidesSplit", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlidesSplit", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid format
*/
func TestPostSlidesSplitInvalidformat(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PostSlidesSplit", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid width
*/
func TestPostSlidesSplitInvalidwidth(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.width = new(int32)

    invalidValue := invalidizeTestParamValue(request.width, "PostSlidesSplit", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.width = nullValue
    } else {
        *request.width = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "width", request.width, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid height
*/
func TestPostSlidesSplitInvalidheight(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.height = new(int32)

    invalidValue := invalidizeTestParamValue(request.height, "PostSlidesSplit", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.height = nullValue
    } else {
        *request.height = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "height", request.height, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid to
*/
func TestPostSlidesSplitInvalidto(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.to = new(int32)

    invalidValue := invalidizeTestParamValue(request.to, "PostSlidesSplit", "to", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.to = nullValue
    } else {
        *request.to = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "to", request.to)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "to", request.to, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid from
*/
func TestPostSlidesSplitInvalidfrom(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.from = new(int32)

    invalidValue := invalidizeTestParamValue(request.from, "PostSlidesSplit", "from", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.from = nullValue
    } else {
        *request.from = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "from", request.from)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "from", request.from, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid destFolder
*/
func TestPostSlidesSplitInvaliddestFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.destFolder, "PostSlidesSplit", "destFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.destFolder = nullValue
    } else {
        request.destFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "destFolder", request.destFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "destFolder", request.destFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid password
*/
func TestPostSlidesSplitInvalidpassword(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PostSlidesSplit", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid storage
*/
func TestPostSlidesSplitInvalidstorage(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PostSlidesSplit", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid folder
*/
func TestPostSlidesSplitInvalidfolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PostSlidesSplit", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid fontsFolder
*/
func TestPostSlidesSplitInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PostSlidesSplit", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method
*/
func TestPutLayoutSlide(t *testing.T) {
    request := createPutLayoutSlideRequest()
    e := initializeTest("PutLayoutSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutLayoutSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid name
*/
func TestPutLayoutSlideInvalidname(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutLayoutSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideIndex
*/
func TestPutLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutLayoutSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideDto
*/
func TestPutLayoutSlideInvalidslideDto(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideDto, "PutLayoutSlide", "slideDto", "LayoutSlide")
    if (invalidValue == nil) {
        request.slideDto = nil
    } else {
        request.slideDto = invalidValue.(ILayoutSlide)
    }

    e := initializeTest("PutLayoutSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "slideDto", request.slideDto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid password
*/
func TestPutLayoutSlideInvalidpassword(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutLayoutSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid folder
*/
func TestPutLayoutSlideInvalidfolder(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutLayoutSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid storage
*/
func TestPutLayoutSlideInvalidstorage(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutLayoutSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method
*/
func TestPutNotesSlideShapeSaveAs(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    e := initializeTest("PutNotesSlideShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.PutNotesSlideShapeSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutNotesSlideShapeSaveAsRequest() PutNotesSlideShapeSaveAsRequest {
    var request PutNotesSlideShapeSaveAsRequest
    request.name = createTestParamValue("PutNotesSlideShapeSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutNotesSlideShapeSaveAs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutNotesSlideShapeSaveAs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutNotesSlideShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PutNotesSlideShapeSaveAs", "format", "string").(string)
    request.outPath = createTestParamValue("PutNotesSlideShapeSaveAs", "outPath", "string").(string)
    request.options = createTestParamValue("PutNotesSlideShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PutNotesSlideShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PutNotesSlideShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PutNotesSlideShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PutNotesSlideShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.scaleX = new(float64)
        *request.scaleX = v
    }
    testscaleY := createTestParamValue("PutNotesSlideShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.scaleY = new(float64)
        *request.scaleY = v
    }
    request.bounds = createTestParamValue("PutNotesSlideShapeSaveAs", "bounds", "string").(string)
    request.fontsFolder = createTestParamValue("PutNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid name
*/
func TestPutNotesSlideShapeSaveAsInvalidname(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutNotesSlideShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutNotesSlideShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid path
*/
func TestPutNotesSlideShapeSaveAsInvalidpath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutNotesSlideShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutNotesSlideShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid format
*/
func TestPutNotesSlideShapeSaveAsInvalidformat(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PutNotesSlideShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid outPath
*/
func TestPutNotesSlideShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.outPath, "PutNotesSlideShapeSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.outPath = nullValue
    } else {
        request.outPath = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "outPath", request.outPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid options
*/
func TestPutNotesSlideShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PutNotesSlideShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid password
*/
func TestPutNotesSlideShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutNotesSlideShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid folder
*/
func TestPutNotesSlideShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutNotesSlideShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid storage
*/
func TestPutNotesSlideShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutNotesSlideShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPutNotesSlideShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.scaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleX, "PutNotesSlideShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleX = nullValue
    } else {
        *request.scaleX = invalidValue.(float64)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleX", request.scaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPutNotesSlideShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.scaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleY, "PutNotesSlideShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleY = nullValue
    } else {
        *request.scaleY = invalidValue.(float64)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleY", request.scaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPutNotesSlideShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.bounds, "PutNotesSlideShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.bounds = nullValue
    } else {
        request.bounds = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "bounds", request.bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPutNotesSlideShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PutNotesSlideShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method
*/
func TestPutPresentationMerge(t *testing.T) {
    request := createPutPresentationMergeRequest()
    e := initializeTest("PutPresentationMerge", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutPresentationMerge(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid name
*/
func TestPutPresentationMergeInvalidname(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutPresentationMerge", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid request
*/
func TestPutPresentationMergeInvalidrequest(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.request, "PutPresentationMerge", "request", "OrderedMergeRequest")
    if (invalidValue == nil) {
        request.request = nil
    } else {
        request.request = invalidValue.(IOrderedMergeRequest)
    }

    e := initializeTest("PutPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "request", request.request, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid password
*/
func TestPutPresentationMergeInvalidpassword(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutPresentationMerge", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid storage
*/
func TestPutPresentationMergeInvalidstorage(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutPresentationMerge", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid folder
*/
func TestPutPresentationMergeInvalidfolder(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutPresentationMerge", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method
*/
func TestPutSetParagraphPortionProperties(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    e := initializeTest("PutSetParagraphPortionProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSetParagraphPortionProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid name
*/
func TestPutSetParagraphPortionPropertiesInvalidname(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSetParagraphPortionProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid slideIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSetParagraphPortionProperties", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid path
*/
func TestPutSetParagraphPortionPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutSetParagraphPortionProperties", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutSetParagraphPortionProperties", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PutSetParagraphPortionProperties", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid portionIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidportionIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "PutSetParagraphPortionProperties", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid dto
*/
func TestPutSetParagraphPortionPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutSetParagraphPortionProperties", "dto", "Portion")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IPortion)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid password
*/
func TestPutSetParagraphPortionPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSetParagraphPortionProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid folder
*/
func TestPutSetParagraphPortionPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSetParagraphPortionProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid storage
*/
func TestPutSetParagraphPortionPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSetParagraphPortionProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method
*/
func TestPutSetParagraphProperties(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    e := initializeTest("PutSetParagraphProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSetParagraphProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid name
*/
func TestPutSetParagraphPropertiesInvalidname(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSetParagraphProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid slideIndex
*/
func TestPutSetParagraphPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSetParagraphProperties", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid path
*/
func TestPutSetParagraphPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutSetParagraphProperties", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutSetParagraphProperties", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PutSetParagraphProperties", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid dto
*/
func TestPutSetParagraphPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutSetParagraphProperties", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PutSetParagraphProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid password
*/
func TestPutSetParagraphPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSetParagraphProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid folder
*/
func TestPutSetParagraphPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSetParagraphProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid storage
*/
func TestPutSetParagraphPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSetParagraphProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method
*/
func TestPutShapeSaveAs(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    e := initializeTest("PutShapeSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.PutShapeSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutShapeSaveAsRequest() PutShapeSaveAsRequest {
    var request PutShapeSaveAsRequest
    request.name = createTestParamValue("PutShapeSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutShapeSaveAs", "slideIndex", "int32").(int32)
    request.path = createTestParamValue("PutShapeSaveAs", "path", "string").(string)
    request.shapeIndex = createTestParamValue("PutShapeSaveAs", "shapeIndex", "int32").(int32)
    request.format = createTestParamValue("PutShapeSaveAs", "format", "string").(string)
    request.outPath = createTestParamValue("PutShapeSaveAs", "outPath", "string").(string)
    request.options = createTestParamValue("PutShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.password = createTestParamValue("PutShapeSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PutShapeSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PutShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PutShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.scaleX = new(float64)
        *request.scaleX = v
    }
    testscaleY := createTestParamValue("PutShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.scaleY = new(float64)
        *request.scaleY = v
    }
    request.bounds = createTestParamValue("PutShapeSaveAs", "bounds", "string").(string)
    request.fontsFolder = createTestParamValue("PutShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid name
*/
func TestPutShapeSaveAsInvalidname(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid slideIndex
*/
func TestPutShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid path
*/
func TestPutShapeSaveAsInvalidpath(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid shapeIndex
*/
func TestPutShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid format
*/
func TestPutShapeSaveAsInvalidformat(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PutShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid outPath
*/
func TestPutShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.outPath, "PutShapeSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.outPath = nullValue
    } else {
        request.outPath = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "outPath", request.outPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid options
*/
func TestPutShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PutShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PutShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid password
*/
func TestPutShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid folder
*/
func TestPutShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid storage
*/
func TestPutShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleX
*/
func TestPutShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.scaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleX, "PutShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleX = nullValue
    } else {
        *request.scaleX = invalidValue.(float64)
    }

    e := initializeTest("PutShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "scaleX", request.scaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleY
*/
func TestPutShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.scaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.scaleY, "PutShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.scaleY = nullValue
    } else {
        *request.scaleY = invalidValue.(float64)
    }

    e := initializeTest("PutShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "scaleY", request.scaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid bounds
*/
func TestPutShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.bounds, "PutShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.bounds = nullValue
    } else {
        request.bounds = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "bounds", request.bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid fontsFolder
*/
func TestPutShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PutShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method
*/
func TestPutSlideAnimation(t *testing.T) {
    request := createPutSlideAnimationRequest()
    e := initializeTest("PutSlideAnimation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlideAnimation(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlideAnimationRequest() PutSlideAnimationRequest {
    var request PutSlideAnimationRequest
    request.name = createTestParamValue("PutSlideAnimation", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlideAnimation", "slideIndex", "int32").(int32)
    request.animation = createTestParamValue("PutSlideAnimation", "animation", "SlideAnimation").(ISlideAnimation)
    request.password = createTestParamValue("PutSlideAnimation", "password", "string").(string)
    request.folder = createTestParamValue("PutSlideAnimation", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid name
*/
func TestPutSlideAnimationInvalidname(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid slideIndex
*/
func TestPutSlideAnimationInvalidslideIndex(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimation", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid animation
*/
func TestPutSlideAnimationInvalidanimation(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.animation, "PutSlideAnimation", "animation", "SlideAnimation")
    if (invalidValue == nil) {
        request.animation = nil
    } else {
        request.animation = invalidValue.(ISlideAnimation)
    }

    e := initializeTest("PutSlideAnimation", "animation", request.animation)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "animation", request.animation, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid password
*/
func TestPutSlideAnimationInvalidpassword(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid folder
*/
func TestPutSlideAnimationInvalidfolder(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid storage
*/
func TestPutSlideAnimationInvalidstorage(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method
*/
func TestPutSlideAnimationEffect(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()
    e := initializeTest("PutSlideAnimationEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlideAnimationEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlideAnimationEffectRequest() PutSlideAnimationEffectRequest {
    var request PutSlideAnimationEffectRequest
    request.name = createTestParamValue("PutSlideAnimationEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.effectIndex = createTestParamValue("PutSlideAnimationEffect", "effectIndex", "int32").(int32)
    request.effect = createTestParamValue("PutSlideAnimationEffect", "effect", "Effect").(IEffect)
    request.password = createTestParamValue("PutSlideAnimationEffect", "password", "string").(string)
    request.folder = createTestParamValue("PutSlideAnimationEffect", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid name
*/
func TestPutSlideAnimationEffectInvalidname(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid slideIndex
*/
func TestPutSlideAnimationEffectInvalidslideIndex(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid effectIndex
*/
func TestPutSlideAnimationEffectInvalideffectIndex(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effectIndex, "PutSlideAnimationEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.effectIndex = nullValue
    } else {
        request.effectIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationEffect", "effectIndex", request.effectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "effectIndex", request.effectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid effect
*/
func TestPutSlideAnimationEffectInvalideffect(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effect, "PutSlideAnimationEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.effect = nil
    } else {
        request.effect = invalidValue.(IEffect)
    }

    e := initializeTest("PutSlideAnimationEffect", "effect", request.effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "effect", request.effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid password
*/
func TestPutSlideAnimationEffectInvalidpassword(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid folder
*/
func TestPutSlideAnimationEffectInvalidfolder(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid storage
*/
func TestPutSlideAnimationEffectInvalidstorage(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method
*/
func TestPutSlideAnimationInteractiveSequenceEffect(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()
    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlideAnimationInteractiveSequenceEffectRequest() PutSlideAnimationInteractiveSequenceEffectRequest {
    var request PutSlideAnimationInteractiveSequenceEffectRequest
    request.name = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.sequenceIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.effectIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32").(int32)
    request.effect = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "effect", "Effect").(IEffect)
    request.password = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.folder = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidname(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidslideIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidsequenceIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.sequenceIndex, "PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.sequenceIndex = nullValue
    } else {
        request.sequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.sequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid effectIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalideffectIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effectIndex, "PutSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.effectIndex = nullValue
    } else {
        request.effectIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "effectIndex", request.effectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "effectIndex", request.effectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid effect
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalideffect(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.effect, "PutSlideAnimationInteractiveSequenceEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.effect = nil
    } else {
        request.effect = invalidValue.(IEffect)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "effect", request.effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "effect", request.effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidpassword(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidfolder(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidstorage(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method
*/
func TestPutSlideSaveAs(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    e := initializeTest("PutSlideSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.PutSlideSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlideSaveAsRequest() PutSlideSaveAsRequest {
    var request PutSlideSaveAsRequest
    request.name = createTestParamValue("PutSlideSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlideSaveAs", "slideIndex", "int32").(int32)
    request.format = createTestParamValue("PutSlideSaveAs", "format", "string").(string)
    request.outPath = createTestParamValue("PutSlideSaveAs", "outPath", "string").(string)
    request.options = createTestParamValue("PutSlideSaveAs", "options", "ExportOptions").(IExportOptions)
    testwidth := createTestParamValue("PutSlideSaveAs", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.width = new(int32)
        *request.width = v
    }
    testheight := createTestParamValue("PutSlideSaveAs", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.height = new(int32)
        *request.height = v
    }
    request.password = createTestParamValue("PutSlideSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PutSlideSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlideSaveAs", "storage", "string").(string)
    request.fontsFolder = createTestParamValue("PutSlideSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid name
*/
func TestPutSlideSaveAsInvalidname(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlideSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid slideIndex
*/
func TestPutSlideSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlideSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid format
*/
func TestPutSlideSaveAsInvalidformat(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PutSlideSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid outPath
*/
func TestPutSlideSaveAsInvalidoutPath(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.outPath, "PutSlideSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.outPath = nullValue
    } else {
        request.outPath = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "outPath", request.outPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid options
*/
func TestPutSlideSaveAsInvalidoptions(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PutSlideSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PutSlideSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid width
*/
func TestPutSlideSaveAsInvalidwidth(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.width = new(int32)

    invalidValue := invalidizeTestParamValue(request.width, "PutSlideSaveAs", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.width = nullValue
    } else {
        *request.width = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "width", request.width, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid height
*/
func TestPutSlideSaveAsInvalidheight(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.height = new(int32)

    invalidValue := invalidizeTestParamValue(request.height, "PutSlideSaveAs", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.height = nullValue
    } else {
        *request.height = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "height", request.height, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid password
*/
func TestPutSlideSaveAsInvalidpassword(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlideSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid folder
*/
func TestPutSlideSaveAsInvalidfolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlideSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid storage
*/
func TestPutSlideSaveAsInvalidstorage(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlideSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid fontsFolder
*/
func TestPutSlideSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PutSlideSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method
*/
func TestPutSlideShapeInfo(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    e := initializeTest("PutSlideShapeInfo", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlideShapeInfo(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid name
*/
func TestPutSlideShapeInfoInvalidname(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlideShapeInfo", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid slideIndex
*/
func TestPutSlideShapeInfoInvalidslideIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlideShapeInfo", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideShapeInfo", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid path
*/
func TestPutSlideShapeInfoInvalidpath(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutSlideShapeInfo", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid shapeIndex
*/
func TestPutSlideShapeInfoInvalidshapeIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutSlideShapeInfo", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideShapeInfo", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid dto
*/
func TestPutSlideShapeInfoInvaliddto(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutSlideShapeInfo", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PutSlideShapeInfo", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid password
*/
func TestPutSlideShapeInfoInvalidpassword(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlideShapeInfo", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid folder
*/
func TestPutSlideShapeInfoInvalidfolder(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlideShapeInfo", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid storage
*/
func TestPutSlideShapeInfoInvalidstorage(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlideShapeInfo", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method
*/
func TestPutSlidesConvert(t *testing.T) {
    request := createPutSlidesConvertRequest()
    e := initializeTest("PutSlidesConvert", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.PutSlidesConvert(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesConvertRequest() PutSlidesConvertRequest {
    var request PutSlidesConvertRequest
    request.format = createTestParamValue("PutSlidesConvert", "format", "string").(string)
    request.outPath = createTestParamValue("PutSlidesConvert", "outPath", "string").(string)
    request.document = createTestParamValue("PutSlidesConvert", "document", "[]byte").([]byte)
    request.password = createTestParamValue("PutSlidesConvert", "password", "string").(string)
    request.fontsFolder = createTestParamValue("PutSlidesConvert", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid format
*/
func TestPutSlidesConvertInvalidformat(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PutSlidesConvert", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid outPath
*/
func TestPutSlidesConvertInvalidoutPath(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.outPath, "PutSlidesConvert", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.outPath = nullValue
    } else {
        request.outPath = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "outPath", request.outPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid document
*/
func TestPutSlidesConvertInvaliddocument(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.document, "PutSlidesConvert", "document", "[]byte")
    if (invalidValue == nil) {
        request.document = nil
    } else {
        request.document = invalidValue.([]byte)
    }

    e := initializeTest("PutSlidesConvert", "document", request.document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "document", request.document, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid password
*/
func TestPutSlidesConvertInvalidpassword(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesConvert", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid fontsFolder
*/
func TestPutSlidesConvertInvalidfontsFolder(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PutSlidesConvert", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method
*/
func TestPutSlidesDocumentFromHtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    e := initializeTest("PutSlidesDocumentFromHtml", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesDocumentFromHtml(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid name
*/
func TestPutSlidesDocumentFromHtmlInvalidname(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesDocumentFromHtml", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid html
*/
func TestPutSlidesDocumentFromHtmlInvalidhtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.html, "PutSlidesDocumentFromHtml", "html", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.html = nullValue
    } else {
        request.html = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "html", request.html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "html", request.html, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid password
*/
func TestPutSlidesDocumentFromHtmlInvalidpassword(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesDocumentFromHtml", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid storage
*/
func TestPutSlidesDocumentFromHtmlInvalidstorage(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesDocumentFromHtml", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid folder
*/
func TestPutSlidesDocumentFromHtmlInvalidfolder(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesDocumentFromHtml", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method
*/
func TestPutSlidesSaveAs(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    e := initializeTest("PutSlidesSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, e = c.SlidesApi.PutSlidesSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesSaveAsRequest() PutSlidesSaveAsRequest {
    var request PutSlidesSaveAsRequest
    request.name = createTestParamValue("PutSlidesSaveAs", "name", "string").(string)
    request.outPath = createTestParamValue("PutSlidesSaveAs", "outPath", "string").(string)
    request.format = createTestParamValue("PutSlidesSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PutSlidesSaveAs", "options", "ExportOptions").(IExportOptions)
    request.password = createTestParamValue("PutSlidesSaveAs", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesSaveAs", "storage", "string").(string)
    request.folder = createTestParamValue("PutSlidesSaveAs", "folder", "string").(string)
    request.fontsFolder = createTestParamValue("PutSlidesSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid name
*/
func TestPutSlidesSaveAsInvalidname(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid outPath
*/
func TestPutSlidesSaveAsInvalidoutPath(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.outPath, "PutSlidesSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.outPath = nullValue
    } else {
        request.outPath = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "outPath", request.outPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid format
*/
func TestPutSlidesSaveAsInvalidformat(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.format, "PutSlidesSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.format = nullValue
    } else {
        request.format = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "format", request.format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid options
*/
func TestPutSlidesSaveAsInvalidoptions(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.options, "PutSlidesSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.options = nil
    } else {
        request.options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PutSlidesSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "options", request.options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid password
*/
func TestPutSlidesSaveAsInvalidpassword(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid storage
*/
func TestPutSlidesSaveAsInvalidstorage(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid folder
*/
func TestPutSlidesSaveAsInvalidfolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid fontsFolder
*/
func TestPutSlidesSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.fontsFolder, "PutSlidesSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.fontsFolder = nullValue
    } else {
        request.fontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "fontsFolder", request.fontsFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method
*/
func TestPutSlidesSetDocumentProperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    e := initializeTest("PutSlidesSetDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesSetDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesSetDocumentPropertyRequest() PutSlidesSetDocumentPropertyRequest {
    var request PutSlidesSetDocumentPropertyRequest
    request.name = createTestParamValue("PutSlidesSetDocumentProperty", "name", "string").(string)
    request.propertyName = createTestParamValue("PutSlidesSetDocumentProperty", "propertyName", "string").(string)
    request.property = createTestParamValue("PutSlidesSetDocumentProperty", "property", "DocumentProperty").(IDocumentProperty)
    request.password = createTestParamValue("PutSlidesSetDocumentProperty", "password", "string").(string)
    request.folder = createTestParamValue("PutSlidesSetDocumentProperty", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlidesSetDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid name
*/
func TestPutSlidesSetDocumentPropertyInvalidname(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSetDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid propertyName
*/
func TestPutSlidesSetDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.propertyName, "PutSlidesSetDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.propertyName = nullValue
    } else {
        request.propertyName = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "propertyName", request.propertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid property
*/
func TestPutSlidesSetDocumentPropertyInvalidproperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.property, "PutSlidesSetDocumentProperty", "property", "DocumentProperty")
    if (invalidValue == nil) {
        request.property = nil
    } else {
        request.property = invalidValue.(IDocumentProperty)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "property", request.property)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "property", request.property, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid password
*/
func TestPutSlidesSetDocumentPropertyInvalidpassword(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSetDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid folder
*/
func TestPutSlidesSetDocumentPropertyInvalidfolder(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSetDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid storage
*/
func TestPutSlidesSetDocumentPropertyInvalidstorage(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSetDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method
*/
func TestPutSlidesSlide(t *testing.T) {
    request := createPutSlidesSlideRequest()
    e := initializeTest("PutSlidesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesSlideRequest() PutSlidesSlideRequest {
    var request PutSlidesSlideRequest
    request.name = createTestParamValue("PutSlidesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlidesSlide", "slideIndex", "int32").(int32)
    request.slideDto = createTestParamValue("PutSlidesSlide", "slideDto", "Slide").(ISlide)
    request.password = createTestParamValue("PutSlidesSlide", "password", "string").(string)
    request.folder = createTestParamValue("PutSlidesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid name
*/
func TestPutSlidesSlideInvalidname(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideIndex
*/
func TestPutSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlidesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideDto
*/
func TestPutSlidesSlideInvalidslideDto(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideDto, "PutSlidesSlide", "slideDto", "Slide")
    if (invalidValue == nil) {
        request.slideDto = nil
    } else {
        request.slideDto = invalidValue.(ISlide)
    }

    e := initializeTest("PutSlidesSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "slideDto", request.slideDto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid password
*/
func TestPutSlidesSlideInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid folder
*/
func TestPutSlidesSlideInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid storage
*/
func TestPutSlidesSlideInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method
*/
func TestPutSlidesSlideBackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    e := initializeTest("PutSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesSlideBackgroundRequest() PutSlidesSlideBackgroundRequest {
    var request PutSlidesSlideBackgroundRequest
    request.name = createTestParamValue("PutSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.background = createTestParamValue("PutSlidesSlideBackground", "background", "SlideBackground").(ISlideBackground)
    request.folder = createTestParamValue("PutSlidesSlideBackground", "folder", "string").(string)
    request.password = createTestParamValue("PutSlidesSlideBackground", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid name
*/
func TestPutSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid background
*/
func TestPutSlidesSlideBackgroundInvalidbackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.background, "PutSlidesSlideBackground", "background", "SlideBackground")
    if (invalidValue == nil) {
        request.background = nil
    } else {
        request.background = invalidValue.(ISlideBackground)
    }

    e := initializeTest("PutSlidesSlideBackground", "background", request.background)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "background", request.background, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid folder
*/
func TestPutSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid password
*/
func TestPutSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid storage
*/
func TestPutSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method
*/
func TestPutSlidesSlideBackgroundColor(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    e := initializeTest("PutSlidesSlideBackgroundColor", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesSlideBackgroundColor(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createPutSlidesSlideBackgroundColorRequest() PutSlidesSlideBackgroundColorRequest {
    var request PutSlidesSlideBackgroundColorRequest
    request.name = createTestParamValue("PutSlidesSlideBackgroundColor", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlidesSlideBackgroundColor", "slideIndex", "int32").(int32)
    request.color = createTestParamValue("PutSlidesSlideBackgroundColor", "color", "string").(string)
    request.folder = createTestParamValue("PutSlidesSlideBackgroundColor", "folder", "string").(string)
    request.password = createTestParamValue("PutSlidesSlideBackgroundColor", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlideBackgroundColor", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid name
*/
func TestPutSlidesSlideBackgroundColorInvalidname(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSlideBackgroundColor", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundColorInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutSlidesSlideBackgroundColor", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid color
*/
func TestPutSlidesSlideBackgroundColorInvalidcolor(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.color, "PutSlidesSlideBackgroundColor", "color", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.color = nullValue
    } else {
        request.color = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "color", request.color)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "color", request.color, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid folder
*/
func TestPutSlidesSlideBackgroundColorInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSlideBackgroundColor", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid password
*/
func TestPutSlidesSlideBackgroundColorInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSlideBackgroundColor", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid storage
*/
func TestPutSlidesSlideBackgroundColorInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSlideBackgroundColor", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method
*/
func TestPutSlidesSlideSize(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    e := initializeTest("PutSlidesSlideSize", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutSlidesSlideSize(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid name
*/
func TestPutSlidesSlideSizeInvalidname(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutSlidesSlideSize", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid password
*/
func TestPutSlidesSlideSizeInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutSlidesSlideSize", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid storage
*/
func TestPutSlidesSlideSizeInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutSlidesSlideSize", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid folder
*/
func TestPutSlidesSlideSizeInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutSlidesSlideSize", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid width
*/
func TestPutSlidesSlideSizeInvalidwidth(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.width = new(int32)

    invalidValue := invalidizeTestParamValue(request.width, "PutSlidesSlideSize", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.width = nullValue
    } else {
        *request.width = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideSize", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "width", request.width, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid height
*/
func TestPutSlidesSlideSizeInvalidheight(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.height = new(int32)

    invalidValue := invalidizeTestParamValue(request.height, "PutSlidesSlideSize", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.height = nullValue
    } else {
        *request.height = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideSize", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "height", request.height, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid sizeType
*/
func TestPutSlidesSlideSizeInvalidsizeType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.sizeType, "PutSlidesSlideSize", "sizeType", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.sizeType = nullValue
    } else {
        request.sizeType = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "sizeType", request.sizeType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "sizeType", request.sizeType, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid scaleType
*/
func TestPutSlidesSlideSizeInvalidscaleType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.scaleType, "PutSlidesSlideSize", "scaleType", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.scaleType = nullValue
    } else {
        request.scaleType = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "scaleType", request.scaleType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "scaleType", request.scaleType, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method
*/
func TestPutUpdateNotesSlide(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    e := initializeTest("PutUpdateNotesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutUpdateNotesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid name
*/
func TestPutUpdateNotesSlideInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutUpdateNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid slideIndex
*/
func TestPutUpdateNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutUpdateNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid dto
*/
func TestPutUpdateNotesSlideInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutUpdateNotesSlide", "dto", "NotesSlide")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(INotesSlide)
    }

    e := initializeTest("PutUpdateNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid password
*/
func TestPutUpdateNotesSlideInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutUpdateNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid folder
*/
func TestPutUpdateNotesSlideInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutUpdateNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid storage
*/
func TestPutUpdateNotesSlideInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutUpdateNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method
*/
func TestPutUpdateNotesSlideShape(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    e := initializeTest("PutUpdateNotesSlideShape", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutUpdateNotesSlideShape(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid name
*/
func TestPutUpdateNotesSlideShapeInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutUpdateNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutUpdateNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid path
*/
func TestPutUpdateNotesSlideShapeInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutUpdateNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutUpdateNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid dto
*/
func TestPutUpdateNotesSlideShapeInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutUpdateNotesSlideShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid password
*/
func TestPutUpdateNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutUpdateNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid folder
*/
func TestPutUpdateNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutUpdateNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid storage
*/
func TestPutUpdateNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutUpdateNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method
*/
func TestPutUpdateNotesSlideShapeParagraph(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid name
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutUpdateNotesSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutUpdateNotesSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid path
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutUpdateNotesSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid dto
*/
func TestPutUpdateNotesSlideShapeParagraphInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutUpdateNotesSlideShapeParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid password
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutUpdateNotesSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid folder
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutUpdateNotesSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid storage
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutUpdateNotesSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method
*/
func TestPutUpdateNotesSlideShapePortion(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    e := initializeTest("PutUpdateNotesSlideShapePortion", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.PutUpdateNotesSlideShapePortion(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
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

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid name
*/
func TestPutUpdateNotesSlideShapePortionInvalidname(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.name, "PutUpdateNotesSlideShapePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.name = nullValue
    } else {
        request.name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "name", request.name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.slideIndex, "PutUpdateNotesSlideShapePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.slideIndex = nullValue
    } else {
        request.slideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "slideIndex", request.slideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid path
*/
func TestPutUpdateNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.path, "PutUpdateNotesSlideShapePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.shapeIndex, "PutUpdateNotesSlideShapePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.shapeIndex = nullValue
    } else {
        request.shapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "shapeIndex", request.shapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.paragraphIndex, "PutUpdateNotesSlideShapePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.paragraphIndex = nullValue
    } else {
        request.paragraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid portionIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.portionIndex, "PutUpdateNotesSlideShapePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.portionIndex = nullValue
    } else {
        request.portionIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "portionIndex", request.portionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid dto
*/
func TestPutUpdateNotesSlideShapePortionInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.dto, "PutUpdateNotesSlideShapePortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.dto = nil
    } else {
        request.dto = invalidValue.(IPortion)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "dto", request.dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid password
*/
func TestPutUpdateNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.password, "PutUpdateNotesSlideShapePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.password = nullValue
    } else {
        request.password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "password", request.password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid folder
*/
func TestPutUpdateNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.folder, "PutUpdateNotesSlideShapePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.folder = nullValue
    } else {
        request.folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "folder", request.folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid storage
*/
func TestPutUpdateNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.storage, "PutUpdateNotesSlideShapePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storage = nullValue
    } else {
        request.storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "storage", request.storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if storage exists
   Test for SlidesApi.StorageExists method
*/
func TestStorageExists(t *testing.T) {
    request := createStorageExistsRequest()
    e := initializeTest("StorageExists", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.StorageExists(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createStorageExistsRequest() StorageExistsRequest {
    var request StorageExistsRequest
    request.storageName = createTestParamValue("StorageExists", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Check if storage exists
   Test for SlidesApi.StorageExists method with invalid storageName
*/
func TestStorageExistsInvalidstorageName(t *testing.T) {
    request := createStorageExistsRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "StorageExists", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("StorageExists", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.StorageExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StorageExists", "storageName", request.storageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method
*/
func TestUploadFile(t *testing.T) {
    request := createUploadFileRequest()
    e := initializeTest("UploadFile", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    _, _, e = c.SlidesApi.UploadFile(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

func createUploadFileRequest() UploadFileRequest {
    var request UploadFileRequest
    request.path = createTestParamValue("UploadFile", "path", "string").(string)
    request.file = createTestParamValue("UploadFile", "file", "[]byte").([]byte)
    request.storageName = createTestParamValue("UploadFile", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid path
*/
func TestUploadFileInvalidpath(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.path, "UploadFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.path = nullValue
    } else {
        request.path = invalidValue.(string)
    }

    e := initializeTest("UploadFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "path", request.path, int32(statusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid file
*/
func TestUploadFileInvalidfile(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.file, "UploadFile", "file", "[]byte")
    if (invalidValue == nil) {
        request.file = nil
    } else {
        request.file = invalidValue.([]byte)
    }

    e := initializeTest("UploadFile", "file", request.file)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "file", request.file, int32(statusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid storageName
*/
func TestUploadFileInvalidstorageName(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.storageName, "UploadFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.storageName = nullValue
    } else {
        request.storageName = invalidValue.(string)
    }

    e := initializeTest("UploadFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "storageName", request.storageName, int32(statusCode), e)
}
