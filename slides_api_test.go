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
    request.SrcPath = createTestParamValue("CopyFile", "srcPath", "string").(string)
    request.DestPath = createTestParamValue("CopyFile", "destPath", "string").(string)
    request.SrcStorageName = createTestParamValue("CopyFile", "srcStorageName", "string").(string)
    request.DestStorageName = createTestParamValue("CopyFile", "destStorageName", "string").(string)
    request.VersionId = createTestParamValue("CopyFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid srcPath
*/
func TestCopyFileInvalidSrcPath(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.SrcPath, "CopyFile", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcPath = nullValue
    } else {
        request.SrcPath = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "srcPath", request.SrcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "srcPath", request.SrcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destPath
*/
func TestCopyFileInvalidDestPath(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.DestPath, "CopyFile", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestPath = nullValue
    } else {
        request.DestPath = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "destPath", request.DestPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "destPath", request.DestPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid srcStorageName
*/
func TestCopyFileInvalidSrcStorageName(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.SrcStorageName, "CopyFile", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcStorageName = nullValue
    } else {
        request.SrcStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "srcStorageName", request.SrcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "srcStorageName", request.SrcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destStorageName
*/
func TestCopyFileInvalidDestStorageName(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.DestStorageName, "CopyFile", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestStorageName = nullValue
    } else {
        request.DestStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "destStorageName", request.DestStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "destStorageName", request.DestStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid versionId
*/
func TestCopyFileInvalidVersionId(t *testing.T) {
    request := createCopyFileRequest()

    invalidValue := invalidizeTestParamValue(request.VersionId, "CopyFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.VersionId = nullValue
    } else {
        request.VersionId = invalidValue.(string)
    }

    e := initializeTest("CopyFile", "versionId", request.VersionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFile", "versionId", request.VersionId, int32(statusCode), e)
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
    request.SrcPath = createTestParamValue("CopyFolder", "srcPath", "string").(string)
    request.DestPath = createTestParamValue("CopyFolder", "destPath", "string").(string)
    request.SrcStorageName = createTestParamValue("CopyFolder", "srcStorageName", "string").(string)
    request.DestStorageName = createTestParamValue("CopyFolder", "destStorageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid srcPath
*/
func TestCopyFolderInvalidSrcPath(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.SrcPath, "CopyFolder", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcPath = nullValue
    } else {
        request.SrcPath = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "srcPath", request.SrcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "srcPath", request.SrcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destPath
*/
func TestCopyFolderInvalidDestPath(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.DestPath, "CopyFolder", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestPath = nullValue
    } else {
        request.DestPath = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "destPath", request.DestPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "destPath", request.DestPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid srcStorageName
*/
func TestCopyFolderInvalidSrcStorageName(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.SrcStorageName, "CopyFolder", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcStorageName = nullValue
    } else {
        request.SrcStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "srcStorageName", request.SrcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "srcStorageName", request.SrcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destStorageName
*/
func TestCopyFolderInvalidDestStorageName(t *testing.T) {
    request := createCopyFolderRequest()

    invalidValue := invalidizeTestParamValue(request.DestStorageName, "CopyFolder", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestStorageName = nullValue
    } else {
        request.DestStorageName = invalidValue.(string)
    }

    e := initializeTest("CopyFolder", "destStorageName", request.DestStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CopyFolder", "destStorageName", request.DestStorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("CreateFolder", "path", "string").(string)
    request.StorageName = createTestParamValue("CreateFolder", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method with invalid path
*/
func TestCreateFolderInvalidPath(t *testing.T) {
    request := createCreateFolderRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "CreateFolder", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("CreateFolder", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CreateFolder", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method with invalid storageName
*/
func TestCreateFolderInvalidStorageName(t *testing.T) {
    request := createCreateFolderRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "CreateFolder", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("CreateFolder", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "CreateFolder", "storageName", request.StorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("DeleteFile", "path", "string").(string)
    request.StorageName = createTestParamValue("DeleteFile", "storageName", "string").(string)
    request.VersionId = createTestParamValue("DeleteFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid path
*/
func TestDeleteFileInvalidPath(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid storageName
*/
func TestDeleteFileInvalidStorageName(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "DeleteFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "storageName", request.StorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid versionId
*/
func TestDeleteFileInvalidVersionId(t *testing.T) {
    request := createDeleteFileRequest()

    invalidValue := invalidizeTestParamValue(request.VersionId, "DeleteFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.VersionId = nullValue
    } else {
        request.VersionId = invalidValue.(string)
    }

    e := initializeTest("DeleteFile", "versionId", request.VersionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFile", "versionId", request.VersionId, int32(statusCode), e)
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
    request.Path = createTestParamValue("DeleteFolder", "path", "string").(string)
    request.StorageName = createTestParamValue("DeleteFolder", "storageName", "string").(string)
    testrecursive := createTestParamValue("DeleteFolder", "recursive", "bool")
    switch v := testrecursive.(type) { 
    case bool:
        request.Recursive = new(bool)
        *request.Recursive = v
    }
    return request
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid path
*/
func TestDeleteFolderInvalidPath(t *testing.T) {
    request := createDeleteFolderRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteFolder", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteFolder", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid storageName
*/
func TestDeleteFolderInvalidStorageName(t *testing.T) {
    request := createDeleteFolderRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "DeleteFolder", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("DeleteFolder", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "storageName", request.StorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid recursive
*/
func TestDeleteFolderInvalidRecursive(t *testing.T) {
    request := createDeleteFolderRequest()
    request.Recursive = new(bool)

    invalidValue := invalidizeTestParamValue(request.Recursive, "DeleteFolder", "recursive", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.Recursive = nullValue
    } else {
        *request.Recursive = invalidValue.(bool)
    }

    e := initializeTest("DeleteFolder", "recursive", request.Recursive)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteFolder", "recursive", request.Recursive, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlide", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteNotesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid name
*/
func TestDeleteNotesSlideInvalidName(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid slideIndex
*/
func TestDeleteNotesSlideInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid password
*/
func TestDeleteNotesSlideInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid folder
*/
func TestDeleteNotesSlideInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid storage
*/
func TestDeleteNotesSlideInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlideParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlideParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlideParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteNotesSlideParagraph", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeleteNotesSlideParagraph", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteNotesSlideParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlideParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlideParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid name
*/
func TestDeleteNotesSlideParagraphInvalidName(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlideParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlideParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid path
*/
func TestDeleteNotesSlideParagraphInvalidPath(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlideParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphInvalidShapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteNotesSlideParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid paragraphIndex
*/
func TestDeleteNotesSlideParagraphInvalidParagraphIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeleteNotesSlideParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid password
*/
func TestDeleteNotesSlideParagraphInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlideParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid folder
*/
func TestDeleteNotesSlideParagraphInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlideParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid storage
*/
func TestDeleteNotesSlideParagraphInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlideParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraph", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlideParagraphs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlideParagraphs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlideParagraphs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteNotesSlideParagraphs", "shapeIndex", "int32").(int32)
    request.Paragraphs = createTestParamValue("DeleteNotesSlideParagraphs", "paragraphs", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteNotesSlideParagraphs", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlideParagraphs", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlideParagraphs", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid name
*/
func TestDeleteNotesSlideParagraphsInvalidName(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlideParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphsInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlideParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid path
*/
func TestDeleteNotesSlideParagraphsInvalidPath(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlideParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphsInvalidShapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteNotesSlideParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid paragraphs
*/
func TestDeleteNotesSlideParagraphsInvalidParagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Paragraphs, "DeleteNotesSlideParagraphs", "paragraphs", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Paragraphs = nullValue
    } else {
        request.Paragraphs = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "paragraphs", request.Paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "paragraphs", request.Paragraphs, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid password
*/
func TestDeleteNotesSlideParagraphsInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlideParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid folder
*/
func TestDeleteNotesSlideParagraphsInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlideParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid storage
*/
func TestDeleteNotesSlideParagraphsInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlideParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideParagraphs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideParagraphs", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlidePortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlidePortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlidePortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteNotesSlidePortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeleteNotesSlidePortion", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("DeleteNotesSlidePortion", "portionIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteNotesSlidePortion", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlidePortion", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlidePortion", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid name
*/
func TestDeleteNotesSlidePortionInvalidName(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlidePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlidePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid path
*/
func TestDeleteNotesSlidePortionInvalidPath(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlidePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionInvalidShapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteNotesSlidePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionInvalidParagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeleteNotesSlidePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid portionIndex
*/
func TestDeleteNotesSlidePortionInvalidPortionIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "DeleteNotesSlidePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortion", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid password
*/
func TestDeleteNotesSlidePortionInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlidePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid folder
*/
func TestDeleteNotesSlidePortionInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlidePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid storage
*/
func TestDeleteNotesSlidePortionInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlidePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortion", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlidePortions", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlidePortions", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlidePortions", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteNotesSlidePortions", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeleteNotesSlidePortions", "paragraphIndex", "int32").(int32)
    request.Portions = createTestParamValue("DeleteNotesSlidePortions", "portions", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteNotesSlidePortions", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlidePortions", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlidePortions", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid name
*/
func TestDeleteNotesSlidePortionsInvalidName(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlidePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionsInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlidePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid path
*/
func TestDeleteNotesSlidePortionsInvalidPath(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlidePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionsInvalidShapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteNotesSlidePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionsInvalidParagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeleteNotesSlidePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid portions
*/
func TestDeleteNotesSlidePortionsInvalidPortions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Portions, "DeleteNotesSlidePortions", "portions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Portions = nullValue
    } else {
        request.Portions = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlidePortions", "portions", request.Portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "portions", request.Portions, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid password
*/
func TestDeleteNotesSlidePortionsInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlidePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid folder
*/
func TestDeleteNotesSlidePortionsInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlidePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid storage
*/
func TestDeleteNotesSlidePortionsInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlidePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlidePortions", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlidePortions", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlideShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlideShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlideShape", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteNotesSlideShape", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteNotesSlideShape", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlideShape", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlideShape", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid name
*/
func TestDeleteNotesSlideShapeInvalidName(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid slideIndex
*/
func TestDeleteNotesSlideShapeInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid path
*/
func TestDeleteNotesSlideShapeInvalidPath(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid shapeIndex
*/
func TestDeleteNotesSlideShapeInvalidShapeIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShape", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid password
*/
func TestDeleteNotesSlideShapeInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid folder
*/
func TestDeleteNotesSlideShapeInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid storage
*/
func TestDeleteNotesSlideShapeInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShape", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteNotesSlideShapes", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteNotesSlideShapes", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteNotesSlideShapes", "path", "string").(string)
    request.Shapes = createTestParamValue("DeleteNotesSlideShapes", "shapes", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteNotesSlideShapes", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteNotesSlideShapes", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteNotesSlideShapes", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid name
*/
func TestDeleteNotesSlideShapesInvalidName(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteNotesSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid slideIndex
*/
func TestDeleteNotesSlideShapesInvalidSlideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteNotesSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteNotesSlideShapes", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid path
*/
func TestDeleteNotesSlideShapesInvalidPath(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteNotesSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid shapes
*/
func TestDeleteNotesSlideShapesInvalidShapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Shapes, "DeleteNotesSlideShapes", "shapes", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Shapes = nullValue
    } else {
        request.Shapes = invalidValue.([]int32)
    }

    e := initializeTest("DeleteNotesSlideShapes", "shapes", request.Shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "shapes", request.Shapes, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid password
*/
func TestDeleteNotesSlideShapesInvalidPassword(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteNotesSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid folder
*/
func TestDeleteNotesSlideShapesInvalidFolder(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteNotesSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid storage
*/
func TestDeleteNotesSlideShapesInvalidStorage(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteNotesSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteNotesSlideShapes", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteNotesSlideShapes", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteParagraph", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeleteParagraph", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid name
*/
func TestDeleteParagraphInvalidName(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid slideIndex
*/
func TestDeleteParagraphInvalidSlideIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid path
*/
func TestDeleteParagraphInvalidPath(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid shapeIndex
*/
func TestDeleteParagraphInvalidShapeIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid paragraphIndex
*/
func TestDeleteParagraphInvalidParagraphIndex(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeleteParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraph", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid password
*/
func TestDeleteParagraphInvalidPassword(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid folder
*/
func TestDeleteParagraphInvalidFolder(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid storage
*/
func TestDeleteParagraphInvalidStorage(t *testing.T) {
    request := createDeleteParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraph", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteParagraphs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteParagraphs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteParagraphs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteParagraphs", "shapeIndex", "int32").(int32)
    request.Paragraphs = createTestParamValue("DeleteParagraphs", "paragraphs", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteParagraphs", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteParagraphs", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteParagraphs", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid name
*/
func TestDeleteParagraphsInvalidName(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid slideIndex
*/
func TestDeleteParagraphsInvalidSlideIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraphs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid path
*/
func TestDeleteParagraphsInvalidPath(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid shapeIndex
*/
func TestDeleteParagraphsInvalidShapeIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteParagraphs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid paragraphs
*/
func TestDeleteParagraphsInvalidParagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Paragraphs, "DeleteParagraphs", "paragraphs", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Paragraphs = nullValue
    } else {
        request.Paragraphs = invalidValue.([]int32)
    }

    e := initializeTest("DeleteParagraphs", "paragraphs", request.Paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "paragraphs", request.Paragraphs, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid password
*/
func TestDeleteParagraphsInvalidPassword(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid folder
*/
func TestDeleteParagraphsInvalidFolder(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid storage
*/
func TestDeleteParagraphsInvalidStorage(t *testing.T) {
    request := createDeleteParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteParagraphs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteParagraphs", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeletePortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeletePortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeletePortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeletePortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeletePortion", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("DeletePortion", "portionIndex", "int32").(int32)
    request.Password = createTestParamValue("DeletePortion", "password", "string").(string)
    request.Folder = createTestParamValue("DeletePortion", "folder", "string").(string)
    request.Storage = createTestParamValue("DeletePortion", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid name
*/
func TestDeletePortionInvalidName(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeletePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid slideIndex
*/
func TestDeletePortionInvalidSlideIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeletePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid path
*/
func TestDeletePortionInvalidPath(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeletePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid shapeIndex
*/
func TestDeletePortionInvalidShapeIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeletePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid paragraphIndex
*/
func TestDeletePortionInvalidParagraphIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeletePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid portionIndex
*/
func TestDeletePortionInvalidPortionIndex(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "DeletePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortion", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid password
*/
func TestDeletePortionInvalidPassword(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeletePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid folder
*/
func TestDeletePortionInvalidFolder(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeletePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid storage
*/
func TestDeletePortionInvalidStorage(t *testing.T) {
    request := createDeletePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeletePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeletePortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortion", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeletePortions", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeletePortions", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeletePortions", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeletePortions", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("DeletePortions", "paragraphIndex", "int32").(int32)
    request.Portions = createTestParamValue("DeletePortions", "portions", "[]int32").([]int32)
    request.Password = createTestParamValue("DeletePortions", "password", "string").(string)
    request.Folder = createTestParamValue("DeletePortions", "folder", "string").(string)
    request.Storage = createTestParamValue("DeletePortions", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid name
*/
func TestDeletePortionsInvalidName(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeletePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid slideIndex
*/
func TestDeletePortionsInvalidSlideIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeletePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid path
*/
func TestDeletePortionsInvalidPath(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeletePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid shapeIndex
*/
func TestDeletePortionsInvalidShapeIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeletePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid paragraphIndex
*/
func TestDeletePortionsInvalidParagraphIndex(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "DeletePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("DeletePortions", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid portions
*/
func TestDeletePortionsInvalidPortions(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Portions, "DeletePortions", "portions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Portions = nullValue
    } else {
        request.Portions = invalidValue.([]int32)
    }

    e := initializeTest("DeletePortions", "portions", request.Portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "portions", request.Portions, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid password
*/
func TestDeletePortionsInvalidPassword(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeletePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid folder
*/
func TestDeletePortionsInvalidFolder(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeletePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid storage
*/
func TestDeletePortionsInvalidStorage(t *testing.T) {
    request := createDeletePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeletePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeletePortions", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeletePortions", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimation", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimation", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimation", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimation", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid name
*/
func TestDeleteSlideAnimationInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid slideIndex
*/
func TestDeleteSlideAnimationInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimation", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid password
*/
func TestDeleteSlideAnimationInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid folder
*/
func TestDeleteSlideAnimationInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove animation from a slide.
   Test for SlidesApi.DeleteSlideAnimation method with invalid storage
*/
func TestDeleteSlideAnimationInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimation", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimation", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimationEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.EffectIndex = createTestParamValue("DeleteSlideAnimationEffect", "effectIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimationEffect", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimationEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid name
*/
func TestDeleteSlideAnimationEffectInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid slideIndex
*/
func TestDeleteSlideAnimationEffectInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid effectIndex
*/
func TestDeleteSlideAnimationEffectInvalidEffectIndex(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.EffectIndex, "DeleteSlideAnimationEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.EffectIndex = nullValue
    } else {
        request.EffectIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "effectIndex", request.EffectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "effectIndex", request.EffectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid password
*/
func TestDeleteSlideAnimationEffectInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid folder
*/
func TestDeleteSlideAnimationEffectInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation.
   Test for SlidesApi.DeleteSlideAnimationEffect method with invalid storage
*/
func TestDeleteSlideAnimationEffectInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "slideIndex", "int32").(int32)
    request.SequenceIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "sequenceIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimationInteractiveSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimationInteractiveSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimationInteractiveSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid sequenceIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidSequenceIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.SequenceIndex, "DeleteSlideAnimationInteractiveSequence", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SequenceIndex = nullValue
    } else {
        request.SequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "sequenceIndex", request.SequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "sequenceIndex", request.SequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimationInteractiveSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimationInteractiveSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an interactive sequence from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequence method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequenceInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimationInteractiveSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequence", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequence", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.SequenceIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.EffectIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidSequenceIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SequenceIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SequenceIndex = nullValue
    } else {
        request.SequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid effectIndex
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidEffectIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.EffectIndex, "DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.EffectIndex = nullValue
    } else {
        request.EffectIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", request.EffectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "effectIndex", request.EffectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove an effect from slide animation interactive sequence.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequenceEffectInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequenceEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequenceEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimationInteractiveSequences", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid name
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimationInteractiveSequences", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid slideIndex
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimationInteractiveSequences", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid password
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimationInteractiveSequences", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid folder
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimationInteractiveSequences", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear all interactive sequences from slide animation.
   Test for SlidesApi.DeleteSlideAnimationInteractiveSequences method with invalid storage
*/
func TestDeleteSlideAnimationInteractiveSequencesInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationInteractiveSequencesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimationInteractiveSequences", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationInteractiveSequences", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationInteractiveSequences(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationInteractiveSequences", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideAnimationMainSequence", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideAnimationMainSequence", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideAnimationMainSequence", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideAnimationMainSequence", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideAnimationMainSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid name
*/
func TestDeleteSlideAnimationMainSequenceInvalidName(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideAnimationMainSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid slideIndex
*/
func TestDeleteSlideAnimationMainSequenceInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideAnimationMainSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid password
*/
func TestDeleteSlideAnimationMainSequenceInvalidPassword(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideAnimationMainSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid folder
*/
func TestDeleteSlideAnimationMainSequenceInvalidFolder(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideAnimationMainSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clear main sequence in slide animation.
   Test for SlidesApi.DeleteSlideAnimationMainSequence method with invalid storage
*/
func TestDeleteSlideAnimationMainSequenceInvalidStorage(t *testing.T) {
    request := createDeleteSlideAnimationMainSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideAnimationMainSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideAnimationMainSequence", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideAnimationMainSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideAnimationMainSequence", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideByIndex", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideByIndex", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideByIndex", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideByIndex", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideByIndex", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid name
*/
func TestDeleteSlideByIndexInvalidName(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideByIndex", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid slideIndex
*/
func TestDeleteSlideByIndexInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideByIndex", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideByIndex", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid password
*/
func TestDeleteSlideByIndexInvalidPassword(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideByIndex", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid folder
*/
func TestDeleteSlideByIndexInvalidFolder(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideByIndex", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid storage
*/
func TestDeleteSlideByIndexInvalidStorage(t *testing.T) {
    request := createDeleteSlideByIndexRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideByIndex", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideByIndex", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideByIndex", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteSlideShape", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("DeleteSlideShape", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlideShape", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideShape", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideShape", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid name
*/
func TestDeleteSlideShapeInvalidName(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid slideIndex
*/
func TestDeleteSlideShapeInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid path
*/
func TestDeleteSlideShapeInvalidPath(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid shapeIndex
*/
func TestDeleteSlideShapeInvalidShapeIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "DeleteSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShape", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid password
*/
func TestDeleteSlideShapeInvalidPassword(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid folder
*/
func TestDeleteSlideShapeInvalidFolder(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid storage
*/
func TestDeleteSlideShapeInvalidStorage(t *testing.T) {
    request := createDeleteSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShape", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlideShapes", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlideShapes", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("DeleteSlideShapes", "path", "string").(string)
    request.Shapes = createTestParamValue("DeleteSlideShapes", "shapes", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteSlideShapes", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlideShapes", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlideShapes", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid name
*/
func TestDeleteSlideShapesInvalidName(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid slideIndex
*/
func TestDeleteSlideShapesInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlideShapes", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid path
*/
func TestDeleteSlideShapesInvalidPath(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DeleteSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid shapes
*/
func TestDeleteSlideShapesInvalidShapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Shapes, "DeleteSlideShapes", "shapes", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Shapes = nullValue
    } else {
        request.Shapes = invalidValue.([]int32)
    }

    e := initializeTest("DeleteSlideShapes", "shapes", request.Shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "shapes", request.Shapes, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid password
*/
func TestDeleteSlideShapesInvalidPassword(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid folder
*/
func TestDeleteSlideShapesInvalidFolder(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid storage
*/
func TestDeleteSlideShapesInvalidStorage(t *testing.T) {
    request := createDeleteSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlideShapes", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlideShapes", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlidesCleanSlidesList", "name", "string").(string)
    request.Slides = createTestParamValue("DeleteSlidesCleanSlidesList", "slides", "[]int32").([]int32)
    request.Password = createTestParamValue("DeleteSlidesCleanSlidesList", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlidesCleanSlidesList", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlidesCleanSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid name
*/
func TestDeleteSlidesCleanSlidesListInvalidName(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlidesCleanSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid slides
*/
func TestDeleteSlidesCleanSlidesListInvalidSlides(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Slides, "DeleteSlidesCleanSlidesList", "slides", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.Slides = nullValue
    } else {
        request.Slides = invalidValue.([]int32)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "slides", request.Slides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "slides", request.Slides, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid password
*/
func TestDeleteSlidesCleanSlidesListInvalidPassword(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlidesCleanSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid folder
*/
func TestDeleteSlidesCleanSlidesListInvalidFolder(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlidesCleanSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid storage
*/
func TestDeleteSlidesCleanSlidesListInvalidStorage(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlidesCleanSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesCleanSlidesList", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesCleanSlidesList", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlidesDocumentProperties", "name", "string").(string)
    request.Password = createTestParamValue("DeleteSlidesDocumentProperties", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlidesDocumentProperties", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlidesDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid name
*/
func TestDeleteSlidesDocumentPropertiesInvalidName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlidesDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid password
*/
func TestDeleteSlidesDocumentPropertiesInvalidPassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlidesDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid folder
*/
func TestDeleteSlidesDocumentPropertiesInvalidFolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlidesDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid storage
*/
func TestDeleteSlidesDocumentPropertiesInvalidStorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlidesDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperties", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperties", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlidesDocumentProperty", "name", "string").(string)
    request.PropertyName = createTestParamValue("DeleteSlidesDocumentProperty", "propertyName", "string").(string)
    request.Password = createTestParamValue("DeleteSlidesDocumentProperty", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlidesDocumentProperty", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlidesDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid name
*/
func TestDeleteSlidesDocumentPropertyInvalidName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlidesDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid propertyName
*/
func TestDeleteSlidesDocumentPropertyInvalidPropertyName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.PropertyName, "DeleteSlidesDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.PropertyName = nullValue
    } else {
        request.PropertyName = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "propertyName", request.PropertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "propertyName", request.PropertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid password
*/
func TestDeleteSlidesDocumentPropertyInvalidPassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlidesDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid folder
*/
func TestDeleteSlidesDocumentPropertyInvalidFolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlidesDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid storage
*/
func TestDeleteSlidesDocumentPropertyInvalidStorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlidesDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesDocumentProperty", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesDocumentProperty", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("DeleteSlidesSlideBackground", "name", "string").(string)
    request.SlideIndex = createTestParamValue("DeleteSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("DeleteSlidesSlideBackground", "password", "string").(string)
    request.Folder = createTestParamValue("DeleteSlidesSlideBackground", "folder", "string").(string)
    request.Storage = createTestParamValue("DeleteSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid name
*/
func TestDeleteSlidesSlideBackgroundInvalidName(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "DeleteSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid slideIndex
*/
func TestDeleteSlidesSlideBackgroundInvalidSlideIndex(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "DeleteSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid password
*/
func TestDeleteSlidesSlideBackgroundInvalidPassword(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "DeleteSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid folder
*/
func TestDeleteSlidesSlideBackgroundInvalidFolder(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "DeleteSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid storage
*/
func TestDeleteSlidesSlideBackgroundInvalidStorage(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "DeleteSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("DeleteSlidesSlideBackground", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DeleteSlidesSlideBackground", "storage", request.Storage, int32(statusCode), e)
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
    request.Path = createTestParamValue("DownloadFile", "path", "string").(string)
    request.StorageName = createTestParamValue("DownloadFile", "storageName", "string").(string)
    request.VersionId = createTestParamValue("DownloadFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid path
*/
func TestDownloadFileInvalidPath(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "DownloadFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid storageName
*/
func TestDownloadFileInvalidStorageName(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "DownloadFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "storageName", request.StorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid versionId
*/
func TestDownloadFileInvalidVersionId(t *testing.T) {
    request := createDownloadFileRequest()

    invalidValue := invalidizeTestParamValue(request.VersionId, "DownloadFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.VersionId = nullValue
    } else {
        request.VersionId = invalidValue.(string)
    }

    e := initializeTest("DownloadFile", "versionId", request.VersionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "DownloadFile", "versionId", request.VersionId, int32(statusCode), e)
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
    request.StorageName = createTestParamValue("GetDiscUsage", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get disc usage
   Test for SlidesApi.GetDiscUsage method with invalid storageName
*/
func TestGetDiscUsageInvalidStorageName(t *testing.T) {
    request := createGetDiscUsageRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "GetDiscUsage", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("GetDiscUsage", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetDiscUsage(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetDiscUsage", "storageName", request.StorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("GetFileVersions", "path", "string").(string)
    request.StorageName = createTestParamValue("GetFileVersions", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method with invalid path
*/
func TestGetFileVersionsInvalidPath(t *testing.T) {
    request := createGetFileVersionsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetFileVersions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetFileVersions", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFileVersions", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method with invalid storageName
*/
func TestGetFileVersionsInvalidStorageName(t *testing.T) {
    request := createGetFileVersionsRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "GetFileVersions", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("GetFileVersions", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFileVersions", "storageName", request.StorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("GetFilesList", "path", "string").(string)
    request.StorageName = createTestParamValue("GetFilesList", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method with invalid path
*/
func TestGetFilesListInvalidPath(t *testing.T) {
    request := createGetFilesListRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetFilesList", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetFilesList", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFilesList", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method with invalid storageName
*/
func TestGetFilesListInvalidStorageName(t *testing.T) {
    request := createGetFilesListRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "GetFilesList", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("GetFilesList", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetFilesList", "storageName", request.StorageName, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetLayoutSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetLayoutSlide", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetLayoutSlide", "password", "string").(string)
    request.Folder = createTestParamValue("GetLayoutSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("GetLayoutSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid name
*/
func TestGetLayoutSlideInvalidName(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetLayoutSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid slideIndex
*/
func TestGetLayoutSlideInvalidSlideIndex(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetLayoutSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetLayoutSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid password
*/
func TestGetLayoutSlideInvalidPassword(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetLayoutSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid folder
*/
func TestGetLayoutSlideInvalidFolder(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetLayoutSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid storage
*/
func TestGetLayoutSlideInvalidStorage(t *testing.T) {
    request := createGetLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetLayoutSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetLayoutSlidesList", "name", "string").(string)
    request.Password = createTestParamValue("GetLayoutSlidesList", "password", "string").(string)
    request.Folder = createTestParamValue("GetLayoutSlidesList", "folder", "string").(string)
    request.Storage = createTestParamValue("GetLayoutSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid name
*/
func TestGetLayoutSlidesListInvalidName(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetLayoutSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid password
*/
func TestGetLayoutSlidesListInvalidPassword(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetLayoutSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid folder
*/
func TestGetLayoutSlidesListInvalidFolder(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetLayoutSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid storage
*/
func TestGetLayoutSlidesListInvalidStorage(t *testing.T) {
    request := createGetLayoutSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetLayoutSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetLayoutSlidesList", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetLayoutSlidesList", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetMasterSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetMasterSlide", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetMasterSlide", "password", "string").(string)
    request.Folder = createTestParamValue("GetMasterSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("GetMasterSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid name
*/
func TestGetMasterSlideInvalidName(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetMasterSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid slideIndex
*/
func TestGetMasterSlideInvalidSlideIndex(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetMasterSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetMasterSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid password
*/
func TestGetMasterSlideInvalidPassword(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetMasterSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid folder
*/
func TestGetMasterSlideInvalidFolder(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetMasterSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid storage
*/
func TestGetMasterSlideInvalidStorage(t *testing.T) {
    request := createGetMasterSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetMasterSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetMasterSlidesList", "name", "string").(string)
    request.Password = createTestParamValue("GetMasterSlidesList", "password", "string").(string)
    request.Folder = createTestParamValue("GetMasterSlidesList", "folder", "string").(string)
    request.Storage = createTestParamValue("GetMasterSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid name
*/
func TestGetMasterSlidesListInvalidName(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetMasterSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid password
*/
func TestGetMasterSlidesListInvalidPassword(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetMasterSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid folder
*/
func TestGetMasterSlidesListInvalidFolder(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetMasterSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid storage
*/
func TestGetMasterSlidesListInvalidStorage(t *testing.T) {
    request := createGetMasterSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetMasterSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetMasterSlidesList", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetMasterSlidesList", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlide", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid name
*/
func TestGetNotesSlideInvalidName(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid slideIndex
*/
func TestGetNotesSlideInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid password
*/
func TestGetNotesSlideInvalidPassword(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid folder
*/
func TestGetNotesSlideInvalidFolder(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid storage
*/
func TestGetNotesSlideInvalidStorage(t *testing.T) {
    request := createGetNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShape", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetNotesSlideShape", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlideShape", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShape", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShape", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid name
*/
func TestGetNotesSlideShapeInvalidName(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid slideIndex
*/
func TestGetNotesSlideShapeInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid path
*/
func TestGetNotesSlideShapeInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid shapeIndex
*/
func TestGetNotesSlideShapeInvalidShapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShape", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid password
*/
func TestGetNotesSlideShapeInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid folder
*/
func TestGetNotesSlideShapeInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid storage
*/
func TestGetNotesSlideShapeInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShape", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShapeParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShapeParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetNotesSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetNotesSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlideShapeParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShapeParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid name
*/
func TestGetNotesSlideShapeParagraphInvalidName(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid path
*/
func TestGetNotesSlideShapeParagraphInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphInvalidShapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetNotesSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetNotesSlideShapeParagraphInvalidParagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetNotesSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid password
*/
func TestGetNotesSlideShapeParagraphInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid folder
*/
func TestGetNotesSlideShapeParagraphInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid storage
*/
func TestGetNotesSlideShapeParagraphInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraph", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShapeParagraphs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShapeParagraphs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShapeParagraphs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetNotesSlideShapeParagraphs", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlideShapeParagraphs", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShapeParagraphs", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShapeParagraphs", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid name
*/
func TestGetNotesSlideShapeParagraphsInvalidName(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShapeParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShapeParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid path
*/
func TestGetNotesSlideShapeParagraphsInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShapeParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidShapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetNotesSlideShapeParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid password
*/
func TestGetNotesSlideShapeParagraphsInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShapeParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid folder
*/
func TestGetNotesSlideShapeParagraphsInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShapeParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid storage
*/
func TestGetNotesSlideShapeParagraphsInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShapeParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapeParagraphs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapeParagraphs", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShapePortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShapePortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShapePortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetNotesSlideShapePortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetNotesSlideShapePortion", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("GetNotesSlideShapePortion", "portionIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlideShapePortion", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShapePortion", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShapePortion", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid name
*/
func TestGetNotesSlideShapePortionInvalidName(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShapePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShapePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid path
*/
func TestGetNotesSlideShapePortionInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShapePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionInvalidShapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetNotesSlideShapePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionInvalidParagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetNotesSlideShapePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid portionIndex
*/
func TestGetNotesSlideShapePortionInvalidPortionIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "GetNotesSlideShapePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortion", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid password
*/
func TestGetNotesSlideShapePortionInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShapePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid folder
*/
func TestGetNotesSlideShapePortionInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShapePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid storage
*/
func TestGetNotesSlideShapePortionInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShapePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortion", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShapePortions", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShapePortions", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShapePortions", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetNotesSlideShapePortions", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetNotesSlideShapePortions", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("GetNotesSlideShapePortions", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShapePortions", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShapePortions", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid name
*/
func TestGetNotesSlideShapePortionsInvalidName(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShapePortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionsInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShapePortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid path
*/
func TestGetNotesSlideShapePortionsInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShapePortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionsInvalidShapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetNotesSlideShapePortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionsInvalidParagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetNotesSlideShapePortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapePortions", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid password
*/
func TestGetNotesSlideShapePortionsInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShapePortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid folder
*/
func TestGetNotesSlideShapePortionsInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShapePortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid storage
*/
func TestGetNotesSlideShapePortionsInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShapePortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapePortions", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapePortions", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideShapes", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideShapes", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetNotesSlideShapes", "path", "string").(string)
    request.Password = createTestParamValue("GetNotesSlideShapes", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideShapes", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideShapes", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid name
*/
func TestGetNotesSlideShapesInvalidName(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid slideIndex
*/
func TestGetNotesSlideShapesInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideShapes", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid path
*/
func TestGetNotesSlideShapesInvalidPath(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetNotesSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid password
*/
func TestGetNotesSlideShapesInvalidPassword(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid folder
*/
func TestGetNotesSlideShapesInvalidFolder(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid storage
*/
func TestGetNotesSlideShapesInvalidStorage(t *testing.T) {
    request := createGetNotesSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideShapes", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideShapes", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetNotesSlideWithFormat", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetNotesSlideWithFormat", "slideIndex", "int32").(int32)
    request.Format = createTestParamValue("GetNotesSlideWithFormat", "format", "string").(string)
    testwidth := createTestParamValue("GetNotesSlideWithFormat", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.Width = new(int32)
        *request.Width = v
    }
    testheight := createTestParamValue("GetNotesSlideWithFormat", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.Height = new(int32)
        *request.Height = v
    }
    request.Password = createTestParamValue("GetNotesSlideWithFormat", "password", "string").(string)
    request.Folder = createTestParamValue("GetNotesSlideWithFormat", "folder", "string").(string)
    request.Storage = createTestParamValue("GetNotesSlideWithFormat", "storage", "string").(string)
    request.FontsFolder = createTestParamValue("GetNotesSlideWithFormat", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid name
*/
func TestGetNotesSlideWithFormatInvalidName(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetNotesSlideWithFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid slideIndex
*/
func TestGetNotesSlideWithFormatInvalidSlideIndex(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetNotesSlideWithFormat", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid format
*/
func TestGetNotesSlideWithFormatInvalidFormat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "GetNotesSlideWithFormat", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid width
*/
func TestGetNotesSlideWithFormatInvalidWidth(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.Width = new(int32)

    invalidValue := invalidizeTestParamValue(request.Width, "GetNotesSlideWithFormat", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Width = nullValue
    } else {
        *request.Width = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "width", request.Width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "width", request.Width, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid height
*/
func TestGetNotesSlideWithFormatInvalidHeight(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.Height = new(int32)

    invalidValue := invalidizeTestParamValue(request.Height, "GetNotesSlideWithFormat", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Height = nullValue
    } else {
        *request.Height = invalidValue.(int32)
    }

    e := initializeTest("GetNotesSlideWithFormat", "height", request.Height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "height", request.Height, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid password
*/
func TestGetNotesSlideWithFormatInvalidPassword(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetNotesSlideWithFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid folder
*/
func TestGetNotesSlideWithFormatInvalidFolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetNotesSlideWithFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid storage
*/
func TestGetNotesSlideWithFormatInvalidStorage(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetNotesSlideWithFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid fontsFolder
*/
func TestGetNotesSlideWithFormatInvalidFontsFolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "GetNotesSlideWithFormat", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("GetNotesSlideWithFormat", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetNotesSlideWithFormat", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetParagraphPortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetParagraphPortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetParagraphPortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetParagraphPortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetParagraphPortion", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("GetParagraphPortion", "portionIndex", "int32").(int32)
    request.Password = createTestParamValue("GetParagraphPortion", "password", "string").(string)
    request.Folder = createTestParamValue("GetParagraphPortion", "folder", "string").(string)
    request.Storage = createTestParamValue("GetParagraphPortion", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid name
*/
func TestGetParagraphPortionInvalidName(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetParagraphPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid slideIndex
*/
func TestGetParagraphPortionInvalidSlideIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetParagraphPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid path
*/
func TestGetParagraphPortionInvalidPath(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetParagraphPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid shapeIndex
*/
func TestGetParagraphPortionInvalidShapeIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetParagraphPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid paragraphIndex
*/
func TestGetParagraphPortionInvalidParagraphIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetParagraphPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid portionIndex
*/
func TestGetParagraphPortionInvalidPortionIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "GetParagraphPortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortion", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid password
*/
func TestGetParagraphPortionInvalidPassword(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetParagraphPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid folder
*/
func TestGetParagraphPortionInvalidFolder(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetParagraphPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid storage
*/
func TestGetParagraphPortionInvalidStorage(t *testing.T) {
    request := createGetParagraphPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetParagraphPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortion", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetParagraphPortions", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetParagraphPortions", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetParagraphPortions", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetParagraphPortions", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetParagraphPortions", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("GetParagraphPortions", "password", "string").(string)
    request.Folder = createTestParamValue("GetParagraphPortions", "folder", "string").(string)
    request.Storage = createTestParamValue("GetParagraphPortions", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid name
*/
func TestGetParagraphPortionsInvalidName(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetParagraphPortions", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid slideIndex
*/
func TestGetParagraphPortionsInvalidSlideIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetParagraphPortions", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid path
*/
func TestGetParagraphPortionsInvalidPath(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetParagraphPortions", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid shapeIndex
*/
func TestGetParagraphPortionsInvalidShapeIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetParagraphPortions", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid paragraphIndex
*/
func TestGetParagraphPortionsInvalidParagraphIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetParagraphPortions", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetParagraphPortions", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid password
*/
func TestGetParagraphPortionsInvalidPassword(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetParagraphPortions", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid folder
*/
func TestGetParagraphPortionsInvalidFolder(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetParagraphPortions", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid storage
*/
func TestGetParagraphPortionsInvalidStorage(t *testing.T) {
    request := createGetParagraphPortionsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetParagraphPortions", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetParagraphPortions", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetParagraphPortions", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlideAnimation", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlideAnimation", "slideIndex", "int32").(int32)
    request.ShapeIndex = createTestParamValue("GetSlideAnimation", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlideAnimation", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlideAnimation", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid name
*/
func TestGetSlideAnimationInvalidName(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid slideIndex
*/
func TestGetSlideAnimationInvalidSlideIndex(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideAnimation", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid shapeIndex
*/
func TestGetSlideAnimationInvalidShapeIndex(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetSlideAnimation", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideAnimation", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid password
*/
func TestGetSlideAnimationInvalidPassword(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid folder
*/
func TestGetSlideAnimationInvalidFolder(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide animation effects.
   Test for SlidesApi.GetSlideAnimation method with invalid storage
*/
func TestGetSlideAnimationInvalidStorage(t *testing.T) {
    request := createGetSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideAnimation", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideAnimation", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlideShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlideShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetSlideShape", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetSlideShape", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlideShape", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlideShape", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlideShape", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid name
*/
func TestGetSlideShapeInvalidName(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid slideIndex
*/
func TestGetSlideShapeInvalidSlideIndex(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid path
*/
func TestGetSlideShapeInvalidPath(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid shapeIndex
*/
func TestGetSlideShapeInvalidShapeIndex(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShape", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid password
*/
func TestGetSlideShapeInvalidPassword(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid folder
*/
func TestGetSlideShapeInvalidFolder(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid storage
*/
func TestGetSlideShapeInvalidStorage(t *testing.T) {
    request := createGetSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShape", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlideShapeParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetSlideShapeParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("GetSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlideShapeParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlideShapeParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid name
*/
func TestGetSlideShapeParagraphInvalidName(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid slideIndex
*/
func TestGetSlideShapeParagraphInvalidSlideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid path
*/
func TestGetSlideShapeParagraphInvalidPath(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphInvalidShapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetSlideShapeParagraphInvalidParagraphIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "GetSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid password
*/
func TestGetSlideShapeParagraphInvalidPassword(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid folder
*/
func TestGetSlideShapeParagraphInvalidFolder(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid storage
*/
func TestGetSlideShapeParagraphInvalidStorage(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraph", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlideShapeParagraphs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlideShapeParagraphs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetSlideShapeParagraphs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("GetSlideShapeParagraphs", "shapeIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlideShapeParagraphs", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlideShapeParagraphs", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlideShapeParagraphs", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid name
*/
func TestGetSlideShapeParagraphsInvalidName(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlideShapeParagraphs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetSlideShapeParagraphsInvalidSlideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlideShapeParagraphs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraphs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid path
*/
func TestGetSlideShapeParagraphsInvalidPath(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetSlideShapeParagraphs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphsInvalidShapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "GetSlideShapeParagraphs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapeParagraphs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid password
*/
func TestGetSlideShapeParagraphsInvalidPassword(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlideShapeParagraphs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid folder
*/
func TestGetSlideShapeParagraphsInvalidFolder(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlideShapeParagraphs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid storage
*/
func TestGetSlideShapeParagraphsInvalidStorage(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlideShapeParagraphs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapeParagraphs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapeParagraphs", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlideShapes", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlideShapes", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("GetSlideShapes", "path", "string").(string)
    request.Password = createTestParamValue("GetSlideShapes", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlideShapes", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlideShapes", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid name
*/
func TestGetSlideShapesInvalidName(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlideShapes", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid slideIndex
*/
func TestGetSlideShapesInvalidSlideIndex(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlideShapes", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlideShapes", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid path
*/
func TestGetSlideShapesInvalidPath(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "GetSlideShapes", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid password
*/
func TestGetSlideShapesInvalidPassword(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlideShapes", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid folder
*/
func TestGetSlideShapesInvalidFolder(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlideShapes", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid storage
*/
func TestGetSlideShapesInvalidStorage(t *testing.T) {
    request := createGetSlideShapesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlideShapes", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlideShapes", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlideShapes", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesDocument", "name", "string").(string)
    request.Password = createTestParamValue("GetSlidesDocument", "password", "string").(string)
    request.Storage = createTestParamValue("GetSlidesDocument", "storage", "string").(string)
    request.Folder = createTestParamValue("GetSlidesDocument", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid name
*/
func TestGetSlidesDocumentInvalidName(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesDocument", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid password
*/
func TestGetSlidesDocumentInvalidPassword(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesDocument", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid storage
*/
func TestGetSlidesDocumentInvalidStorage(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesDocument", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid folder
*/
func TestGetSlidesDocumentInvalidFolder(t *testing.T) {
    request := createGetSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesDocument", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocument", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocument", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesDocumentProperties", "name", "string").(string)
    request.Password = createTestParamValue("GetSlidesDocumentProperties", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesDocumentProperties", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid name
*/
func TestGetSlidesDocumentPropertiesInvalidName(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid password
*/
func TestGetSlidesDocumentPropertiesInvalidPassword(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid folder
*/
func TestGetSlidesDocumentPropertiesInvalidFolder(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid storage
*/
func TestGetSlidesDocumentPropertiesInvalidStorage(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperties", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperties", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesDocumentProperty", "name", "string").(string)
    request.PropertyName = createTestParamValue("GetSlidesDocumentProperty", "propertyName", "string").(string)
    request.Password = createTestParamValue("GetSlidesDocumentProperty", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesDocumentProperty", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid name
*/
func TestGetSlidesDocumentPropertyInvalidName(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid propertyName
*/
func TestGetSlidesDocumentPropertyInvalidPropertyName(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.PropertyName, "GetSlidesDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.PropertyName = nullValue
    } else {
        request.PropertyName = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "propertyName", request.PropertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "propertyName", request.PropertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid password
*/
func TestGetSlidesDocumentPropertyInvalidPassword(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid folder
*/
func TestGetSlidesDocumentPropertyInvalidFolder(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid storage
*/
func TestGetSlidesDocumentPropertyInvalidStorage(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesDocumentProperty", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesDocumentProperty", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesImageWithDefaultFormat", "name", "string").(string)
    request.Index = createTestParamValue("GetSlidesImageWithDefaultFormat", "index", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesImageWithDefaultFormat", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesImageWithDefaultFormat", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesImageWithDefaultFormat", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid name
*/
func TestGetSlidesImageWithDefaultFormatInvalidName(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesImageWithDefaultFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid index
*/
func TestGetSlidesImageWithDefaultFormatInvalidIndex(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Index, "GetSlidesImageWithDefaultFormat", "index", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.Index = nullValue
    } else {
        request.Index = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "index", request.Index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "index", request.Index, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid password
*/
func TestGetSlidesImageWithDefaultFormatInvalidPassword(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesImageWithDefaultFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid folder
*/
func TestGetSlidesImageWithDefaultFormatInvalidFolder(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesImageWithDefaultFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid storage
*/
func TestGetSlidesImageWithDefaultFormatInvalidStorage(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesImageWithDefaultFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithDefaultFormat", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithDefaultFormat", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesImageWithFormat", "name", "string").(string)
    request.Index = createTestParamValue("GetSlidesImageWithFormat", "index", "int32").(int32)
    request.Format = createTestParamValue("GetSlidesImageWithFormat", "format", "string").(string)
    request.Password = createTestParamValue("GetSlidesImageWithFormat", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesImageWithFormat", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesImageWithFormat", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid name
*/
func TestGetSlidesImageWithFormatInvalidName(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesImageWithFormat", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid index
*/
func TestGetSlidesImageWithFormatInvalidIndex(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Index, "GetSlidesImageWithFormat", "index", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.Index = nullValue
    } else {
        request.Index = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesImageWithFormat", "index", request.Index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "index", request.Index, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid format
*/
func TestGetSlidesImageWithFormatInvalidFormat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "GetSlidesImageWithFormat", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid password
*/
func TestGetSlidesImageWithFormatInvalidPassword(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesImageWithFormat", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid folder
*/
func TestGetSlidesImageWithFormatInvalidFolder(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesImageWithFormat", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid storage
*/
func TestGetSlidesImageWithFormatInvalidStorage(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesImageWithFormat", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImageWithFormat", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImageWithFormat", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesImages", "name", "string").(string)
    request.Password = createTestParamValue("GetSlidesImages", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesImages", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesImages", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid name
*/
func TestGetSlidesImagesInvalidName(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesImages", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid password
*/
func TestGetSlidesImagesInvalidPassword(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesImages", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid folder
*/
func TestGetSlidesImagesInvalidFolder(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesImages", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid storage
*/
func TestGetSlidesImagesInvalidStorage(t *testing.T) {
    request := createGetSlidesImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesImages", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesImages", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesImages", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesPlaceholder", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesPlaceholder", "slideIndex", "int32").(int32)
    request.PlaceholderIndex = createTestParamValue("GetSlidesPlaceholder", "placeholderIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesPlaceholder", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesPlaceholder", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesPlaceholder", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid name
*/
func TestGetSlidesPlaceholderInvalidName(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesPlaceholder", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid slideIndex
*/
func TestGetSlidesPlaceholderInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesPlaceholder", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholder", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid placeholderIndex
*/
func TestGetSlidesPlaceholderInvalidPlaceholderIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.PlaceholderIndex, "GetSlidesPlaceholder", "placeholderIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PlaceholderIndex = nullValue
    } else {
        request.PlaceholderIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholder", "placeholderIndex", request.PlaceholderIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "placeholderIndex", request.PlaceholderIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid password
*/
func TestGetSlidesPlaceholderInvalidPassword(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesPlaceholder", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid folder
*/
func TestGetSlidesPlaceholderInvalidFolder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesPlaceholder", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid storage
*/
func TestGetSlidesPlaceholderInvalidStorage(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesPlaceholder", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholder", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholder", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesPlaceholders", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesPlaceholders", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesPlaceholders", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesPlaceholders", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesPlaceholders", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid name
*/
func TestGetSlidesPlaceholdersInvalidName(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesPlaceholders", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid slideIndex
*/
func TestGetSlidesPlaceholdersInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesPlaceholders", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesPlaceholders", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid password
*/
func TestGetSlidesPlaceholdersInvalidPassword(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesPlaceholders", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid folder
*/
func TestGetSlidesPlaceholdersInvalidFolder(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesPlaceholders", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid storage
*/
func TestGetSlidesPlaceholdersInvalidStorage(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesPlaceholders", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPlaceholders", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPlaceholders", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesPresentationTextItems", "name", "string").(string)
    testwithEmpty := createTestParamValue("GetSlidesPresentationTextItems", "withEmpty", "bool")
    switch v := testwithEmpty.(type) { 
    case bool:
        request.WithEmpty = new(bool)
        *request.WithEmpty = v
    }
    request.Password = createTestParamValue("GetSlidesPresentationTextItems", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesPresentationTextItems", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesPresentationTextItems", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid name
*/
func TestGetSlidesPresentationTextItemsInvalidName(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesPresentationTextItems", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid withEmpty
*/
func TestGetSlidesPresentationTextItemsInvalidWithEmpty(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.WithEmpty = new(bool)

    invalidValue := invalidizeTestParamValue(request.WithEmpty, "GetSlidesPresentationTextItems", "withEmpty", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.WithEmpty = nullValue
    } else {
        *request.WithEmpty = invalidValue.(bool)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "withEmpty", request.WithEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "withEmpty", request.WithEmpty, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid password
*/
func TestGetSlidesPresentationTextItemsInvalidPassword(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesPresentationTextItems", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid folder
*/
func TestGetSlidesPresentationTextItemsInvalidFolder(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesPresentationTextItems", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid storage
*/
func TestGetSlidesPresentationTextItemsInvalidStorage(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesPresentationTextItems", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesPresentationTextItems", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesPresentationTextItems", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesSlide", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid name
*/
func TestGetSlidesSlideInvalidName(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid slideIndex
*/
func TestGetSlidesSlideInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid password
*/
func TestGetSlidesSlideInvalidPassword(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid folder
*/
func TestGetSlidesSlideInvalidFolder(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid storage
*/
func TestGetSlidesSlideInvalidStorage(t *testing.T) {
    request := createGetSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlideBackground", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesSlideBackground", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlideBackground", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid name
*/
func TestGetSlidesSlideBackgroundInvalidName(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid slideIndex
*/
func TestGetSlidesSlideBackgroundInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideBackground", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid password
*/
func TestGetSlidesSlideBackgroundInvalidPassword(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid folder
*/
func TestGetSlidesSlideBackgroundInvalidFolder(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid storage
*/
func TestGetSlidesSlideBackgroundInvalidStorage(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideBackground", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideBackground", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlideComments", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesSlideComments", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesSlideComments", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlideComments", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlideComments", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid name
*/
func TestGetSlidesSlideCommentsInvalidName(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlideComments", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid slideIndex
*/
func TestGetSlidesSlideCommentsInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesSlideComments", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideComments", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid password
*/
func TestGetSlidesSlideCommentsInvalidPassword(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlideComments", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid folder
*/
func TestGetSlidesSlideCommentsInvalidFolder(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlideComments", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid storage
*/
func TestGetSlidesSlideCommentsInvalidStorage(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlideComments", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideComments", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideComments", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlideImages", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesSlideImages", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesSlideImages", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlideImages", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlideImages", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid name
*/
func TestGetSlidesSlideImagesInvalidName(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlideImages", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid slideIndex
*/
func TestGetSlidesSlideImagesInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesSlideImages", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideImages", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid password
*/
func TestGetSlidesSlideImagesInvalidPassword(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlideImages", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid folder
*/
func TestGetSlidesSlideImagesInvalidFolder(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlideImages", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid storage
*/
func TestGetSlidesSlideImagesInvalidStorage(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlideImages", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideImages", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideImages", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlideTextItems", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesSlideTextItems", "slideIndex", "int32").(int32)
    testwithEmpty := createTestParamValue("GetSlidesSlideTextItems", "withEmpty", "bool")
    switch v := testwithEmpty.(type) { 
    case bool:
        request.WithEmpty = new(bool)
        *request.WithEmpty = v
    }
    request.Password = createTestParamValue("GetSlidesSlideTextItems", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlideTextItems", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlideTextItems", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid name
*/
func TestGetSlidesSlideTextItemsInvalidName(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlideTextItems", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid slideIndex
*/
func TestGetSlidesSlideTextItemsInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesSlideTextItems", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesSlideTextItems", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid withEmpty
*/
func TestGetSlidesSlideTextItemsInvalidWithEmpty(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.WithEmpty = new(bool)

    invalidValue := invalidizeTestParamValue(request.WithEmpty, "GetSlidesSlideTextItems", "withEmpty", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.WithEmpty = nullValue
    } else {
        *request.WithEmpty = invalidValue.(bool)
    }

    e := initializeTest("GetSlidesSlideTextItems", "withEmpty", request.WithEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "withEmpty", request.WithEmpty, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid password
*/
func TestGetSlidesSlideTextItemsInvalidPassword(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlideTextItems", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid folder
*/
func TestGetSlidesSlideTextItemsInvalidFolder(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlideTextItems", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid storage
*/
func TestGetSlidesSlideTextItemsInvalidStorage(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlideTextItems", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlideTextItems", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlideTextItems", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesSlidesList", "name", "string").(string)
    request.Password = createTestParamValue("GetSlidesSlidesList", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesSlidesList", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid name
*/
func TestGetSlidesSlidesListInvalidName(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesSlidesList", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid password
*/
func TestGetSlidesSlidesListInvalidPassword(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesSlidesList", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid folder
*/
func TestGetSlidesSlidesListInvalidFolder(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesSlidesList", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid storage
*/
func TestGetSlidesSlidesListInvalidStorage(t *testing.T) {
    request := createGetSlidesSlidesListRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesSlidesList", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesSlidesList", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesSlidesList", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesTheme", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesTheme", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesTheme", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesTheme", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesTheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid name
*/
func TestGetSlidesThemeInvalidName(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesTheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid slideIndex
*/
func TestGetSlidesThemeInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesTheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesTheme", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid password
*/
func TestGetSlidesThemeInvalidPassword(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesTheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid folder
*/
func TestGetSlidesThemeInvalidFolder(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesTheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid storage
*/
func TestGetSlidesThemeInvalidStorage(t *testing.T) {
    request := createGetSlidesThemeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesTheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesTheme", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesTheme", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesThemeColorScheme", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesThemeColorScheme", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesThemeColorScheme", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesThemeColorScheme", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesThemeColorScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid name
*/
func TestGetSlidesThemeColorSchemeInvalidName(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesThemeColorScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid slideIndex
*/
func TestGetSlidesThemeColorSchemeInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesThemeColorScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid password
*/
func TestGetSlidesThemeColorSchemeInvalidPassword(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesThemeColorScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid folder
*/
func TestGetSlidesThemeColorSchemeInvalidFolder(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesThemeColorScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid storage
*/
func TestGetSlidesThemeColorSchemeInvalidStorage(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesThemeColorScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeColorScheme", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeColorScheme", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesThemeFontScheme", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesThemeFontScheme", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesThemeFontScheme", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesThemeFontScheme", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesThemeFontScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid name
*/
func TestGetSlidesThemeFontSchemeInvalidName(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesThemeFontScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFontSchemeInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesThemeFontScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid password
*/
func TestGetSlidesThemeFontSchemeInvalidPassword(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesThemeFontScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid folder
*/
func TestGetSlidesThemeFontSchemeInvalidFolder(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesThemeFontScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid storage
*/
func TestGetSlidesThemeFontSchemeInvalidStorage(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesThemeFontScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFontScheme", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFontScheme", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("GetSlidesThemeFormatScheme", "name", "string").(string)
    request.SlideIndex = createTestParamValue("GetSlidesThemeFormatScheme", "slideIndex", "int32").(int32)
    request.Password = createTestParamValue("GetSlidesThemeFormatScheme", "password", "string").(string)
    request.Folder = createTestParamValue("GetSlidesThemeFormatScheme", "folder", "string").(string)
    request.Storage = createTestParamValue("GetSlidesThemeFormatScheme", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid name
*/
func TestGetSlidesThemeFormatSchemeInvalidName(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "GetSlidesThemeFormatScheme", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFormatSchemeInvalidSlideIndex(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "GetSlidesThemeFormatScheme", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid password
*/
func TestGetSlidesThemeFormatSchemeInvalidPassword(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "GetSlidesThemeFormatScheme", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid folder
*/
func TestGetSlidesThemeFormatSchemeInvalidFolder(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "GetSlidesThemeFormatScheme", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid storage
*/
func TestGetSlidesThemeFormatSchemeInvalidStorage(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "GetSlidesThemeFormatScheme", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("GetSlidesThemeFormatScheme", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetSlidesThemeFormatScheme", "storage", request.Storage, int32(statusCode), e)
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
    request.SrcPath = createTestParamValue("MoveFile", "srcPath", "string").(string)
    request.DestPath = createTestParamValue("MoveFile", "destPath", "string").(string)
    request.SrcStorageName = createTestParamValue("MoveFile", "srcStorageName", "string").(string)
    request.DestStorageName = createTestParamValue("MoveFile", "destStorageName", "string").(string)
    request.VersionId = createTestParamValue("MoveFile", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid srcPath
*/
func TestMoveFileInvalidSrcPath(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.SrcPath, "MoveFile", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcPath = nullValue
    } else {
        request.SrcPath = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "srcPath", request.SrcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "srcPath", request.SrcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destPath
*/
func TestMoveFileInvalidDestPath(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.DestPath, "MoveFile", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestPath = nullValue
    } else {
        request.DestPath = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "destPath", request.DestPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "destPath", request.DestPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid srcStorageName
*/
func TestMoveFileInvalidSrcStorageName(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.SrcStorageName, "MoveFile", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcStorageName = nullValue
    } else {
        request.SrcStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "srcStorageName", request.SrcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "srcStorageName", request.SrcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destStorageName
*/
func TestMoveFileInvalidDestStorageName(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.DestStorageName, "MoveFile", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestStorageName = nullValue
    } else {
        request.DestStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "destStorageName", request.DestStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "destStorageName", request.DestStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid versionId
*/
func TestMoveFileInvalidVersionId(t *testing.T) {
    request := createMoveFileRequest()

    invalidValue := invalidizeTestParamValue(request.VersionId, "MoveFile", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.VersionId = nullValue
    } else {
        request.VersionId = invalidValue.(string)
    }

    e := initializeTest("MoveFile", "versionId", request.VersionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFile", "versionId", request.VersionId, int32(statusCode), e)
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
    request.SrcPath = createTestParamValue("MoveFolder", "srcPath", "string").(string)
    request.DestPath = createTestParamValue("MoveFolder", "destPath", "string").(string)
    request.SrcStorageName = createTestParamValue("MoveFolder", "srcStorageName", "string").(string)
    request.DestStorageName = createTestParamValue("MoveFolder", "destStorageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid srcPath
*/
func TestMoveFolderInvalidSrcPath(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.SrcPath, "MoveFolder", "srcPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcPath = nullValue
    } else {
        request.SrcPath = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "srcPath", request.SrcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "srcPath", request.SrcPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destPath
*/
func TestMoveFolderInvalidDestPath(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.DestPath, "MoveFolder", "destPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestPath = nullValue
    } else {
        request.DestPath = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "destPath", request.DestPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "destPath", request.DestPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid srcStorageName
*/
func TestMoveFolderInvalidSrcStorageName(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.SrcStorageName, "MoveFolder", "srcStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SrcStorageName = nullValue
    } else {
        request.SrcStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "srcStorageName", request.SrcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "srcStorageName", request.SrcStorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destStorageName
*/
func TestMoveFolderInvalidDestStorageName(t *testing.T) {
    request := createMoveFolderRequest()

    invalidValue := invalidizeTestParamValue(request.DestStorageName, "MoveFolder", "destStorageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestStorageName = nullValue
    } else {
        request.DestStorageName = invalidValue.(string)
    }

    e := initializeTest("MoveFolder", "destStorageName", request.DestStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "MoveFolder", "destStorageName", request.DestStorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("ObjectExists", "path", "string").(string)
    request.StorageName = createTestParamValue("ObjectExists", "storageName", "string").(string)
    request.VersionId = createTestParamValue("ObjectExists", "versionId", "string").(string)
    return request
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid path
*/
func TestObjectExistsInvalidPath(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "ObjectExists", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid storageName
*/
func TestObjectExistsInvalidStorageName(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "ObjectExists", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "storageName", request.StorageName, int32(statusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid versionId
*/
func TestObjectExistsInvalidVersionId(t *testing.T) {
    request := createObjectExistsRequest()

    invalidValue := invalidizeTestParamValue(request.VersionId, "ObjectExists", "versionId", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.VersionId = nullValue
    } else {
        request.VersionId = invalidValue.(string)
    }

    e := initializeTest("ObjectExists", "versionId", request.VersionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "ObjectExists", "versionId", request.VersionId, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostAddNewParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostAddNewParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostAddNewParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostAddNewParagraph", "shapeIndex", "int32").(int32)
    request.Dto = createTestParamValue("PostAddNewParagraph", "dto", "Paragraph").(IParagraph)
    request.Password = createTestParamValue("PostAddNewParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("PostAddNewParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("PostAddNewParagraph", "storage", "string").(string)
    testposition := createTestParamValue("PostAddNewParagraph", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid name
*/
func TestPostAddNewParagraphInvalidName(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostAddNewParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid slideIndex
*/
func TestPostAddNewParagraphInvalidSlideIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostAddNewParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid path
*/
func TestPostAddNewParagraphInvalidPath(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostAddNewParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid shapeIndex
*/
func TestPostAddNewParagraphInvalidShapeIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostAddNewParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid dto
*/
func TestPostAddNewParagraphInvalidDto(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostAddNewParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PostAddNewParagraph", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid password
*/
func TestPostAddNewParagraphInvalidPassword(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostAddNewParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid folder
*/
func TestPostAddNewParagraphInvalidFolder(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostAddNewParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid storage
*/
func TestPostAddNewParagraphInvalidStorage(t *testing.T) {
    request := createPostAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostAddNewParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid position
*/
func TestPostAddNewParagraphInvalidPosition(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostAddNewParagraph", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewParagraph", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewParagraph", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostAddNewPortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostAddNewPortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostAddNewPortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostAddNewPortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PostAddNewPortion", "paragraphIndex", "int32").(int32)
    request.Dto = createTestParamValue("PostAddNewPortion", "dto", "Portion").(IPortion)
    request.Password = createTestParamValue("PostAddNewPortion", "password", "string").(string)
    request.Folder = createTestParamValue("PostAddNewPortion", "folder", "string").(string)
    request.Storage = createTestParamValue("PostAddNewPortion", "storage", "string").(string)
    testposition := createTestParamValue("PostAddNewPortion", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid name
*/
func TestPostAddNewPortionInvalidName(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostAddNewPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid slideIndex
*/
func TestPostAddNewPortionInvalidSlideIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostAddNewPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid path
*/
func TestPostAddNewPortionInvalidPath(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostAddNewPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid shapeIndex
*/
func TestPostAddNewPortionInvalidShapeIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostAddNewPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid paragraphIndex
*/
func TestPostAddNewPortionInvalidParagraphIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PostAddNewPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid dto
*/
func TestPostAddNewPortionInvalidDto(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostAddNewPortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IPortion)
    }

    e := initializeTest("PostAddNewPortion", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid password
*/
func TestPostAddNewPortionInvalidPassword(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostAddNewPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid folder
*/
func TestPostAddNewPortionInvalidFolder(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostAddNewPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid storage
*/
func TestPostAddNewPortionInvalidStorage(t *testing.T) {
    request := createPostAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostAddNewPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewPortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid position
*/
func TestPostAddNewPortionInvalidPosition(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostAddNewPortion", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewPortion", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewPortion", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostAddNewShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostAddNewShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostAddNewShape", "path", "string").(string)
    request.Dto = createTestParamValue("PostAddNewShape", "dto", "ShapeBase").(IShapeBase)
    request.Password = createTestParamValue("PostAddNewShape", "password", "string").(string)
    request.Folder = createTestParamValue("PostAddNewShape", "folder", "string").(string)
    request.Storage = createTestParamValue("PostAddNewShape", "storage", "string").(string)
    testshapeToClone := createTestParamValue("PostAddNewShape", "shapeToClone", "int32")
    switch v := testshapeToClone.(type) { 
    case int32:
        request.ShapeToClone = new(int32)
        *request.ShapeToClone = v
    }
    testposition := createTestParamValue("PostAddNewShape", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid name
*/
func TestPostAddNewShapeInvalidName(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostAddNewShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid slideIndex
*/
func TestPostAddNewShapeInvalidSlideIndex(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostAddNewShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid path
*/
func TestPostAddNewShapeInvalidPath(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostAddNewShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid dto
*/
func TestPostAddNewShapeInvalidDto(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostAddNewShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PostAddNewShape", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid password
*/
func TestPostAddNewShapeInvalidPassword(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostAddNewShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid folder
*/
func TestPostAddNewShapeInvalidFolder(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostAddNewShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid storage
*/
func TestPostAddNewShapeInvalidStorage(t *testing.T) {
    request := createPostAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostAddNewShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNewShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid shapeToClone
*/
func TestPostAddNewShapeInvalidShapeToClone(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.ShapeToClone = new(int32)

    invalidValue := invalidizeTestParamValue(request.ShapeToClone, "PostAddNewShape", "shapeToClone", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.ShapeToClone = nullValue
    } else {
        *request.ShapeToClone = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "shapeToClone", request.ShapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "shapeToClone", request.ShapeToClone, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid position
*/
func TestPostAddNewShapeInvalidPosition(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostAddNewShape", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostAddNewShape", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNewShape", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostAddNotesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostAddNotesSlide", "slideIndex", "int32").(int32)
    request.Dto = createTestParamValue("PostAddNotesSlide", "dto", "NotesSlide").(INotesSlide)
    request.Password = createTestParamValue("PostAddNotesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("PostAddNotesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("PostAddNotesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid name
*/
func TestPostAddNotesSlideInvalidName(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostAddNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid slideIndex
*/
func TestPostAddNotesSlideInvalidSlideIndex(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostAddNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostAddNotesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid dto
*/
func TestPostAddNotesSlideInvalidDto(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostAddNotesSlide", "dto", "NotesSlide")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(INotesSlide)
    }

    e := initializeTest("PostAddNotesSlide", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid password
*/
func TestPostAddNotesSlideInvalidPassword(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostAddNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid folder
*/
func TestPostAddNotesSlideInvalidFolder(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostAddNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid storage
*/
func TestPostAddNotesSlideInvalidStorage(t *testing.T) {
    request := createPostAddNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostAddNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostAddNotesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostAddNotesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "name", "string").(string)
    request.CloneFrom = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", "string").(string)
    request.CloneFromPosition = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", "int32").(int32)
    request.CloneFromPassword = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", "string").(string)
    request.CloneFromStorage = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", "string").(string)
    request.Password = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "password", "string").(string)
    request.Folder = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "folder", "string").(string)
    request.Storage = createTestParamValue("PostCopyLayoutSlideFromSourcePresentation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidName(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostCopyLayoutSlideFromSourcePresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidCloneFrom(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFrom, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFrom = nullValue
    } else {
        request.CloneFrom = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.CloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.CloneFrom, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidCloneFromPosition(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromPosition, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.CloneFromPosition = nullValue
    } else {
        request.CloneFromPosition = invalidValue.(int32)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.CloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.CloneFromPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidCloneFromPassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromPassword, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFromPassword = nullValue
    } else {
        request.CloneFromPassword = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.CloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.CloneFromPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidCloneFromStorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromStorage, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFromStorage = nullValue
    } else {
        request.CloneFromStorage = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.CloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.CloneFromStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidPassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostCopyLayoutSlideFromSourcePresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidFolder(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostCopyLayoutSlideFromSourcePresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidStorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostCopyLayoutSlideFromSourcePresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "name", "string").(string)
    request.CloneFrom = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", "string").(string)
    request.CloneFromPosition = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", "int32").(int32)
    request.CloneFromPassword = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", "string").(string)
    request.CloneFromStorage = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", "string").(string)
    testapplyToAll := createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "applyToAll", "bool")
    switch v := testapplyToAll.(type) { 
    case bool:
        request.ApplyToAll = new(bool)
        *request.ApplyToAll = v
    }
    request.Password = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "password", "string").(string)
    request.Folder = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "folder", "string").(string)
    request.Storage = createTestParamValue("PostCopyMasterSlideFromSourcePresentation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid name
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidName(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostCopyMasterSlideFromSourcePresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidCloneFrom(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFrom, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFrom = nullValue
    } else {
        request.CloneFrom = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.CloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.CloneFrom, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidCloneFromPosition(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromPosition, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.CloneFromPosition = nullValue
    } else {
        request.CloneFromPosition = invalidValue.(int32)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.CloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.CloneFromPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidCloneFromPassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromPassword, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFromPassword = nullValue
    } else {
        request.CloneFromPassword = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.CloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.CloneFromPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidCloneFromStorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.CloneFromStorage, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.CloneFromStorage = nullValue
    } else {
        request.CloneFromStorage = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.CloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.CloneFromStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid applyToAll
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidApplyToAll(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.ApplyToAll = new(bool)

    invalidValue := invalidizeTestParamValue(request.ApplyToAll, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.ApplyToAll = nullValue
    } else {
        *request.ApplyToAll = invalidValue.(bool)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.ApplyToAll)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.ApplyToAll, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidPassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostCopyMasterSlideFromSourcePresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidFolder(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostCopyMasterSlideFromSourcePresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidStorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostCopyMasterSlideFromSourcePresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostNotesSlideAddNewParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostNotesSlideAddNewParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostNotesSlideAddNewParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostNotesSlideAddNewParagraph", "shapeIndex", "int32").(int32)
    request.Dto = createTestParamValue("PostNotesSlideAddNewParagraph", "dto", "Paragraph").(IParagraph)
    request.Password = createTestParamValue("PostNotesSlideAddNewParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("PostNotesSlideAddNewParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("PostNotesSlideAddNewParagraph", "storage", "string").(string)
    testposition := createTestParamValue("PostNotesSlideAddNewParagraph", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid name
*/
func TestPostNotesSlideAddNewParagraphInvalidName(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostNotesSlideAddNewParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid slideIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidSlideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostNotesSlideAddNewParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid path
*/
func TestPostNotesSlideAddNewParagraphInvalidPath(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostNotesSlideAddNewParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidShapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostNotesSlideAddNewParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid dto
*/
func TestPostNotesSlideAddNewParagraphInvalidDto(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostNotesSlideAddNewParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid password
*/
func TestPostNotesSlideAddNewParagraphInvalidPassword(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostNotesSlideAddNewParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid folder
*/
func TestPostNotesSlideAddNewParagraphInvalidFolder(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostNotesSlideAddNewParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid storage
*/
func TestPostNotesSlideAddNewParagraphInvalidStorage(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostNotesSlideAddNewParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid position
*/
func TestPostNotesSlideAddNewParagraphInvalidPosition(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostNotesSlideAddNewParagraph", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewParagraph", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewParagraph", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostNotesSlideAddNewPortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostNotesSlideAddNewPortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostNotesSlideAddNewPortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostNotesSlideAddNewPortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PostNotesSlideAddNewPortion", "paragraphIndex", "int32").(int32)
    request.Dto = createTestParamValue("PostNotesSlideAddNewPortion", "dto", "Portion").(IPortion)
    request.Password = createTestParamValue("PostNotesSlideAddNewPortion", "password", "string").(string)
    request.Folder = createTestParamValue("PostNotesSlideAddNewPortion", "folder", "string").(string)
    request.Storage = createTestParamValue("PostNotesSlideAddNewPortion", "storage", "string").(string)
    testposition := createTestParamValue("PostNotesSlideAddNewPortion", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid name
*/
func TestPostNotesSlideAddNewPortionInvalidName(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostNotesSlideAddNewPortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid slideIndex
*/
func TestPostNotesSlideAddNewPortionInvalidSlideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostNotesSlideAddNewPortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid path
*/
func TestPostNotesSlideAddNewPortionInvalidPath(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostNotesSlideAddNewPortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewPortionInvalidShapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostNotesSlideAddNewPortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid paragraphIndex
*/
func TestPostNotesSlideAddNewPortionInvalidParagraphIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PostNotesSlideAddNewPortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid dto
*/
func TestPostNotesSlideAddNewPortionInvalidDto(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostNotesSlideAddNewPortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IPortion)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid password
*/
func TestPostNotesSlideAddNewPortionInvalidPassword(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostNotesSlideAddNewPortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid folder
*/
func TestPostNotesSlideAddNewPortionInvalidFolder(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostNotesSlideAddNewPortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid storage
*/
func TestPostNotesSlideAddNewPortionInvalidStorage(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostNotesSlideAddNewPortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid position
*/
func TestPostNotesSlideAddNewPortionInvalidPosition(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostNotesSlideAddNewPortion", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewPortion", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewPortion", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostNotesSlideAddNewShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostNotesSlideAddNewShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostNotesSlideAddNewShape", "path", "string").(string)
    request.Dto = createTestParamValue("PostNotesSlideAddNewShape", "dto", "ShapeBase").(IShapeBase)
    request.Password = createTestParamValue("PostNotesSlideAddNewShape", "password", "string").(string)
    request.Folder = createTestParamValue("PostNotesSlideAddNewShape", "folder", "string").(string)
    request.Storage = createTestParamValue("PostNotesSlideAddNewShape", "storage", "string").(string)
    testshapeToClone := createTestParamValue("PostNotesSlideAddNewShape", "shapeToClone", "int32")
    switch v := testshapeToClone.(type) { 
    case int32:
        request.ShapeToClone = new(int32)
        *request.ShapeToClone = v
    }
    testposition := createTestParamValue("PostNotesSlideAddNewShape", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    return request
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid name
*/
func TestPostNotesSlideAddNewShapeInvalidName(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostNotesSlideAddNewShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid slideIndex
*/
func TestPostNotesSlideAddNewShapeInvalidSlideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostNotesSlideAddNewShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid path
*/
func TestPostNotesSlideAddNewShapeInvalidPath(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostNotesSlideAddNewShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid dto
*/
func TestPostNotesSlideAddNewShapeInvalidDto(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PostNotesSlideAddNewShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid password
*/
func TestPostNotesSlideAddNewShapeInvalidPassword(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostNotesSlideAddNewShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid folder
*/
func TestPostNotesSlideAddNewShapeInvalidFolder(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostNotesSlideAddNewShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid storage
*/
func TestPostNotesSlideAddNewShapeInvalidStorage(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostNotesSlideAddNewShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid shapeToClone
*/
func TestPostNotesSlideAddNewShapeInvalidShapeToClone(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.ShapeToClone = new(int32)

    invalidValue := invalidizeTestParamValue(request.ShapeToClone, "PostNotesSlideAddNewShape", "shapeToClone", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.ShapeToClone = nullValue
    } else {
        *request.ShapeToClone = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "shapeToClone", request.ShapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "shapeToClone", request.ShapeToClone, int32(statusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid position
*/
func TestPostNotesSlideAddNewShapeInvalidPosition(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostNotesSlideAddNewShape", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideAddNewShape", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideAddNewShape", "position", request.Position, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostNotesSlideShapeSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostNotesSlideShapeSaveAs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostNotesSlideShapeSaveAs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostNotesSlideShapeSaveAs", "shapeIndex", "int32").(int32)
    request.Format = createTestParamValue("PostNotesSlideShapeSaveAs", "format", "string").(string)
    request.Options = createTestParamValue("PostNotesSlideShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.Password = createTestParamValue("PostNotesSlideShapeSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PostNotesSlideShapeSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PostNotesSlideShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PostNotesSlideShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.ScaleX = new(float64)
        *request.ScaleX = v
    }
    testscaleY := createTestParamValue("PostNotesSlideShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.ScaleY = new(float64)
        *request.ScaleY = v
    }
    request.Bounds = createTestParamValue("PostNotesSlideShapeSaveAs", "bounds", "string").(string)
    request.FontsFolder = createTestParamValue("PostNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid name
*/
func TestPostNotesSlideShapeSaveAsInvalidName(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostNotesSlideShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostNotesSlideShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid path
*/
func TestPostNotesSlideShapeSaveAsInvalidPath(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostNotesSlideShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidShapeIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostNotesSlideShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid format
*/
func TestPostNotesSlideShapeSaveAsInvalidFormat(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostNotesSlideShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid options
*/
func TestPostNotesSlideShapeSaveAsInvalidOptions(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PostNotesSlideShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid password
*/
func TestPostNotesSlideShapeSaveAsInvalidPassword(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostNotesSlideShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid folder
*/
func TestPostNotesSlideShapeSaveAsInvalidFolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostNotesSlideShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid storage
*/
func TestPostNotesSlideShapeSaveAsInvalidStorage(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostNotesSlideShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPostNotesSlideShapeSaveAsInvalidScaleX(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.ScaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleX, "PostNotesSlideShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleX = nullValue
    } else {
        *request.ScaleX = invalidValue.(float64)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleX", request.ScaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleX", request.ScaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPostNotesSlideShapeSaveAsInvalidScaleY(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.ScaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleY, "PostNotesSlideShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleY = nullValue
    } else {
        *request.ScaleY = invalidValue.(float64)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleY", request.ScaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleY", request.ScaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPostNotesSlideShapeSaveAsInvalidBounds(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Bounds, "PostNotesSlideShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Bounds = nullValue
    } else {
        request.Bounds = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "bounds", request.Bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "bounds", request.Bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPostNotesSlideShapeSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostNotesSlideShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostNotesSlideShapeSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostNotesSlideShapeSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostPresentationMerge", "name", "string").(string)
    request.Request = createTestParamValue("PostPresentationMerge", "request", "PresentationsMergeRequest").(IPresentationsMergeRequest)
    request.Password = createTestParamValue("PostPresentationMerge", "password", "string").(string)
    request.Storage = createTestParamValue("PostPresentationMerge", "storage", "string").(string)
    request.Folder = createTestParamValue("PostPresentationMerge", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid name
*/
func TestPostPresentationMergeInvalidName(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostPresentationMerge", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid request
*/
func TestPostPresentationMergeInvalidRequest(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Request, "PostPresentationMerge", "request", "PresentationsMergeRequest")
    if (invalidValue == nil) {
        request.Request = nil
    } else {
        request.Request = invalidValue.(IPresentationsMergeRequest)
    }

    e := initializeTest("PostPresentationMerge", "request", request.Request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "request", request.Request, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid password
*/
func TestPostPresentationMergeInvalidPassword(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostPresentationMerge", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid storage
*/
func TestPostPresentationMergeInvalidStorage(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostPresentationMerge", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid folder
*/
func TestPostPresentationMergeInvalidFolder(t *testing.T) {
    request := createPostPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostPresentationMerge", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostPresentationMerge", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostPresentationMerge", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostShapeSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostShapeSaveAs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PostShapeSaveAs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PostShapeSaveAs", "shapeIndex", "int32").(int32)
    request.Format = createTestParamValue("PostShapeSaveAs", "format", "string").(string)
    request.Options = createTestParamValue("PostShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.Password = createTestParamValue("PostShapeSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PostShapeSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PostShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PostShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.ScaleX = new(float64)
        *request.ScaleX = v
    }
    testscaleY := createTestParamValue("PostShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.ScaleY = new(float64)
        *request.ScaleY = v
    }
    request.Bounds = createTestParamValue("PostShapeSaveAs", "bounds", "string").(string)
    request.FontsFolder = createTestParamValue("PostShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid name
*/
func TestPostShapeSaveAsInvalidName(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid slideIndex
*/
func TestPostShapeSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostShapeSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid path
*/
func TestPostShapeSaveAsInvalidPath(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PostShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid shapeIndex
*/
func TestPostShapeSaveAsInvalidShapeIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PostShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PostShapeSaveAs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid format
*/
func TestPostShapeSaveAsInvalidFormat(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid options
*/
func TestPostShapeSaveAsInvalidOptions(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PostShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PostShapeSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid password
*/
func TestPostShapeSaveAsInvalidPassword(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid folder
*/
func TestPostShapeSaveAsInvalidFolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid storage
*/
func TestPostShapeSaveAsInvalidStorage(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleX
*/
func TestPostShapeSaveAsInvalidScaleX(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.ScaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleX, "PostShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleX = nullValue
    } else {
        *request.ScaleX = invalidValue.(float64)
    }

    e := initializeTest("PostShapeSaveAs", "scaleX", request.ScaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "scaleX", request.ScaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleY
*/
func TestPostShapeSaveAsInvalidScaleY(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.ScaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleY, "PostShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleY = nullValue
    } else {
        *request.ScaleY = invalidValue.(float64)
    }

    e := initializeTest("PostShapeSaveAs", "scaleY", request.ScaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "scaleY", request.ScaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid bounds
*/
func TestPostShapeSaveAsInvalidBounds(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Bounds, "PostShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Bounds = nullValue
    } else {
        request.Bounds = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "bounds", request.Bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "bounds", request.Bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid fontsFolder
*/
func TestPostShapeSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostShapeSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostShapeSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlideAnimationEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.Effect = createTestParamValue("PostSlideAnimationEffect", "effect", "Effect").(IEffect)
    request.Password = createTestParamValue("PostSlideAnimationEffect", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlideAnimationEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid name
*/
func TestPostSlideAnimationEffectInvalidName(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid slideIndex
*/
func TestPostSlideAnimationEffectInvalidSlideIndex(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid effect
*/
func TestPostSlideAnimationEffectInvalidEffect(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Effect, "PostSlideAnimationEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.Effect = nil
    } else {
        request.Effect = invalidValue.(IEffect)
    }

    e := initializeTest("PostSlideAnimationEffect", "effect", request.Effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "effect", request.Effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid password
*/
func TestPostSlideAnimationEffectInvalidPassword(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid folder
*/
func TestPostSlideAnimationEffectInvalidFolder(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an effect to slide animation.
   Test for SlidesApi.PostSlideAnimationEffect method with invalid storage
*/
func TestPostSlideAnimationEffectInvalidStorage(t *testing.T) {
    request := createPostSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlideAnimationInteractiveSequence", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlideAnimationInteractiveSequence", "slideIndex", "int32").(int32)
    request.Sequence = createTestParamValue("PostSlideAnimationInteractiveSequence", "sequence", "InteractiveSequence").(IInteractiveSequence)
    request.Password = createTestParamValue("PostSlideAnimationInteractiveSequence", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlideAnimationInteractiveSequence", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlideAnimationInteractiveSequence", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid name
*/
func TestPostSlideAnimationInteractiveSequenceInvalidName(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlideAnimationInteractiveSequence", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid slideIndex
*/
func TestPostSlideAnimationInteractiveSequenceInvalidSlideIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlideAnimationInteractiveSequence", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid sequence
*/
func TestPostSlideAnimationInteractiveSequenceInvalidSequence(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Sequence, "PostSlideAnimationInteractiveSequence", "sequence", "InteractiveSequence")
    if (invalidValue == nil) {
        request.Sequence = nil
    } else {
        request.Sequence = invalidValue.(IInteractiveSequence)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "sequence", request.Sequence)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "sequence", request.Sequence, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid password
*/
func TestPostSlideAnimationInteractiveSequenceInvalidPassword(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlideAnimationInteractiveSequence", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid folder
*/
func TestPostSlideAnimationInteractiveSequenceInvalidFolder(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlideAnimationInteractiveSequence", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PostSlideAnimationInteractiveSequence method with invalid storage
*/
func TestPostSlideAnimationInteractiveSequenceInvalidStorage(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlideAnimationInteractiveSequence", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequence", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequence(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequence", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.SequenceIndex = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.Effect = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "effect", "Effect").(IEffect)
    request.Password = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidName(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidSlideIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidSequenceIndex(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SequenceIndex, "PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SequenceIndex = nullValue
    } else {
        request.SequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid effect
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidEffect(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Effect, "PostSlideAnimationInteractiveSequenceEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.Effect = nil
    } else {
        request.Effect = invalidValue.(IEffect)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "effect", request.Effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "effect", request.Effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidPassword(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidFolder(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Add an animation effect to a slide interactive sequence.
   Test for SlidesApi.PostSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestPostSlideAnimationInteractiveSequenceEffectInvalidStorage(t *testing.T) {
    request := createPostSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideAnimationInteractiveSequenceEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideAnimationInteractiveSequenceEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlideSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlideSaveAs", "slideIndex", "int32").(int32)
    request.Format = createTestParamValue("PostSlideSaveAs", "format", "string").(string)
    request.Options = createTestParamValue("PostSlideSaveAs", "options", "ExportOptions").(IExportOptions)
    testwidth := createTestParamValue("PostSlideSaveAs", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.Width = new(int32)
        *request.Width = v
    }
    testheight := createTestParamValue("PostSlideSaveAs", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.Height = new(int32)
        *request.Height = v
    }
    request.Password = createTestParamValue("PostSlideSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlideSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlideSaveAs", "storage", "string").(string)
    request.FontsFolder = createTestParamValue("PostSlideSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid name
*/
func TestPostSlideSaveAsInvalidName(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlideSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid slideIndex
*/
func TestPostSlideSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlideSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid format
*/
func TestPostSlideSaveAsInvalidFormat(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostSlideSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid options
*/
func TestPostSlideSaveAsInvalidOptions(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PostSlideSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlideSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid width
*/
func TestPostSlideSaveAsInvalidWidth(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.Width = new(int32)

    invalidValue := invalidizeTestParamValue(request.Width, "PostSlideSaveAs", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Width = nullValue
    } else {
        *request.Width = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "width", request.Width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "width", request.Width, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid height
*/
func TestPostSlideSaveAsInvalidHeight(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.Height = new(int32)

    invalidValue := invalidizeTestParamValue(request.Height, "PostSlideSaveAs", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Height = nullValue
    } else {
        *request.Height = invalidValue.(int32)
    }

    e := initializeTest("PostSlideSaveAs", "height", request.Height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "height", request.Height, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid password
*/
func TestPostSlideSaveAsInvalidPassword(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlideSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid folder
*/
func TestPostSlideSaveAsInvalidFolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlideSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid storage
*/
func TestPostSlideSaveAsInvalidStorage(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlideSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid fontsFolder
*/
func TestPostSlideSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostSlideSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlideSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlideSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesAdd", "name", "string").(string)
    testposition := createTestParamValue("PostSlidesAdd", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    request.Password = createTestParamValue("PostSlidesAdd", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesAdd", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesAdd", "storage", "string").(string)
    request.LayoutAlias = createTestParamValue("PostSlidesAdd", "layoutAlias", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid name
*/
func TestPostSlidesAddInvalidName(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesAdd", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid position
*/
func TestPostSlidesAddInvalidPosition(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostSlidesAdd", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesAdd", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "position", request.Position, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid password
*/
func TestPostSlidesAddInvalidPassword(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesAdd", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid folder
*/
func TestPostSlidesAddInvalidFolder(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesAdd", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid storage
*/
func TestPostSlidesAddInvalidStorage(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesAdd", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid layoutAlias
*/
func TestPostSlidesAddInvalidLayoutAlias(t *testing.T) {
    request := createPostSlidesAddRequest()

    invalidValue := invalidizeTestParamValue(request.LayoutAlias, "PostSlidesAdd", "layoutAlias", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.LayoutAlias = nullValue
    } else {
        request.LayoutAlias = invalidValue.(string)
    }

    e := initializeTest("PostSlidesAdd", "layoutAlias", request.LayoutAlias)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesAdd", "layoutAlias", request.LayoutAlias, int32(statusCode), e)
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
    request.Format = createTestParamValue("PostSlidesConvert", "format", "string").(string)
    request.Document = createTestParamValue("PostSlidesConvert", "document", "[]byte").([]byte)
    request.Password = createTestParamValue("PostSlidesConvert", "password", "string").(string)
    request.FontsFolder = createTestParamValue("PostSlidesConvert", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid format
*/
func TestPostSlidesConvertInvalidFormat(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostSlidesConvert", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid document
*/
func TestPostSlidesConvertInvalidDocument(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Document, "PostSlidesConvert", "document", "[]byte")
    if (invalidValue == nil) {
        request.Document = nil
    } else {
        request.Document = invalidValue.([]byte)
    }

    e := initializeTest("PostSlidesConvert", "document", request.Document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "document", request.Document, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid password
*/
func TestPostSlidesConvertInvalidPassword(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesConvert", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid fontsFolder
*/
func TestPostSlidesConvertInvalidFontsFolder(t *testing.T) {
    request := createPostSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostSlidesConvert", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesConvert", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesConvert", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesCopy", "name", "string").(string)
    request.SlideToCopy = createTestParamValue("PostSlidesCopy", "slideToCopy", "int32").(int32)
    testposition := createTestParamValue("PostSlidesCopy", "position", "int32")
    switch v := testposition.(type) { 
    case int32:
        request.Position = new(int32)
        *request.Position = v
    }
    request.Source = createTestParamValue("PostSlidesCopy", "source", "string").(string)
    request.SourcePassword = createTestParamValue("PostSlidesCopy", "sourcePassword", "string").(string)
    request.SourceStorage = createTestParamValue("PostSlidesCopy", "sourceStorage", "string").(string)
    request.Password = createTestParamValue("PostSlidesCopy", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesCopy", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesCopy", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid name
*/
func TestPostSlidesCopyInvalidName(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesCopy", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid slideToCopy
*/
func TestPostSlidesCopyInvalidSlideToCopy(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.SlideToCopy, "PostSlidesCopy", "slideToCopy", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideToCopy = nullValue
    } else {
        request.SlideToCopy = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesCopy", "slideToCopy", request.SlideToCopy)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "slideToCopy", request.SlideToCopy, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid position
*/
func TestPostSlidesCopyInvalidPosition(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.Position = new(int32)

    invalidValue := invalidizeTestParamValue(request.Position, "PostSlidesCopy", "position", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Position = nullValue
    } else {
        *request.Position = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesCopy", "position", request.Position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "position", request.Position, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid source
*/
func TestPostSlidesCopyInvalidSource(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.Source, "PostSlidesCopy", "source", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Source = nullValue
    } else {
        request.Source = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "source", request.Source)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "source", request.Source, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourcePassword
*/
func TestPostSlidesCopyInvalidSourcePassword(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.SourcePassword, "PostSlidesCopy", "sourcePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SourcePassword = nullValue
    } else {
        request.SourcePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "sourcePassword", request.SourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "sourcePassword", request.SourcePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourceStorage
*/
func TestPostSlidesCopyInvalidSourceStorage(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.SourceStorage, "PostSlidesCopy", "sourceStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SourceStorage = nullValue
    } else {
        request.SourceStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "sourceStorage", request.SourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "sourceStorage", request.SourceStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid password
*/
func TestPostSlidesCopyInvalidPassword(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesCopy", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid folder
*/
func TestPostSlidesCopyInvalidFolder(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesCopy", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid storage
*/
func TestPostSlidesCopyInvalidStorage(t *testing.T) {
    request := createPostSlidesCopyRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesCopy", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesCopy", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesCopy", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesDocument", "name", "string").(string)
    request.Data = createTestParamValue("PostSlidesDocument", "data", "[]byte").([]byte)
    request.InputPassword = createTestParamValue("PostSlidesDocument", "inputPassword", "string").(string)
    request.Password = createTestParamValue("PostSlidesDocument", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesDocument", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesDocument", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid name
*/
func TestPostSlidesDocumentInvalidName(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesDocument", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid data
*/
func TestPostSlidesDocumentInvalidData(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Data, "PostSlidesDocument", "data", "[]byte")
    if (invalidValue == nil) {
        request.Data = nil
    } else {
        request.Data = invalidValue.([]byte)
    }

    e := initializeTest("PostSlidesDocument", "data", request.Data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "data", request.Data, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid inputPassword
*/
func TestPostSlidesDocumentInvalidInputPassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.InputPassword, "PostSlidesDocument", "inputPassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.InputPassword = nullValue
    } else {
        request.InputPassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "inputPassword", request.InputPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "inputPassword", request.InputPassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid password
*/
func TestPostSlidesDocumentInvalidPassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesDocument", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid storage
*/
func TestPostSlidesDocumentInvalidStorage(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesDocument", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid folder
*/
func TestPostSlidesDocumentInvalidFolder(t *testing.T) {
    request := createPostSlidesDocumentRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesDocument", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocument", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocument", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesDocumentFromHtml", "name", "string").(string)
    request.Html = createTestParamValue("PostSlidesDocumentFromHtml", "html", "string").(string)
    request.Password = createTestParamValue("PostSlidesDocumentFromHtml", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesDocumentFromHtml", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesDocumentFromHtml", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid name
*/
func TestPostSlidesDocumentFromHtmlInvalidName(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesDocumentFromHtml", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid html
*/
func TestPostSlidesDocumentFromHtmlInvalidHtml(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Html, "PostSlidesDocumentFromHtml", "html", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Html = nullValue
    } else {
        request.Html = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "html", request.Html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "html", request.Html, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid password
*/
func TestPostSlidesDocumentFromHtmlInvalidPassword(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesDocumentFromHtml", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid storage
*/
func TestPostSlidesDocumentFromHtmlInvalidStorage(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesDocumentFromHtml", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid folder
*/
func TestPostSlidesDocumentFromHtmlInvalidFolder(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesDocumentFromHtml", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromHtml", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromHtml", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesDocumentFromSource", "name", "string").(string)
    request.SourcePath = createTestParamValue("PostSlidesDocumentFromSource", "sourcePath", "string").(string)
    request.SourcePassword = createTestParamValue("PostSlidesDocumentFromSource", "sourcePassword", "string").(string)
    request.SourceStorage = createTestParamValue("PostSlidesDocumentFromSource", "sourceStorage", "string").(string)
    request.Password = createTestParamValue("PostSlidesDocumentFromSource", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesDocumentFromSource", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesDocumentFromSource", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid name
*/
func TestPostSlidesDocumentFromSourceInvalidName(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesDocumentFromSource", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePath
*/
func TestPostSlidesDocumentFromSourceInvalidSourcePath(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.SourcePath, "PostSlidesDocumentFromSource", "sourcePath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SourcePath = nullValue
    } else {
        request.SourcePath = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePath", request.SourcePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourcePath", request.SourcePath, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePassword
*/
func TestPostSlidesDocumentFromSourceInvalidSourcePassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.SourcePassword, "PostSlidesDocumentFromSource", "sourcePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SourcePassword = nullValue
    } else {
        request.SourcePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePassword", request.SourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourcePassword", request.SourcePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourceStorage
*/
func TestPostSlidesDocumentFromSourceInvalidSourceStorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.SourceStorage, "PostSlidesDocumentFromSource", "sourceStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SourceStorage = nullValue
    } else {
        request.SourceStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "sourceStorage", request.SourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "sourceStorage", request.SourceStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid password
*/
func TestPostSlidesDocumentFromSourceInvalidPassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesDocumentFromSource", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid storage
*/
func TestPostSlidesDocumentFromSourceInvalidStorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesDocumentFromSource", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid folder
*/
func TestPostSlidesDocumentFromSourceInvalidFolder(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesDocumentFromSource", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromSource", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromSource", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesDocumentFromTemplate", "name", "string").(string)
    request.TemplatePath = createTestParamValue("PostSlidesDocumentFromTemplate", "templatePath", "string").(string)
    request.Data = createTestParamValue("PostSlidesDocumentFromTemplate", "data", "string").(string)
    request.TemplatePassword = createTestParamValue("PostSlidesDocumentFromTemplate", "templatePassword", "string").(string)
    request.TemplateStorage = createTestParamValue("PostSlidesDocumentFromTemplate", "templateStorage", "string").(string)
    testisImageDataEmbedded := createTestParamValue("PostSlidesDocumentFromTemplate", "isImageDataEmbedded", "bool")
    switch v := testisImageDataEmbedded.(type) { 
    case bool:
        request.IsImageDataEmbedded = new(bool)
        *request.IsImageDataEmbedded = v
    }
    request.Password = createTestParamValue("PostSlidesDocumentFromTemplate", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesDocumentFromTemplate", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesDocumentFromTemplate", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid name
*/
func TestPostSlidesDocumentFromTemplateInvalidName(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesDocumentFromTemplate", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePath
*/
func TestPostSlidesDocumentFromTemplateInvalidTemplatePath(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.TemplatePath, "PostSlidesDocumentFromTemplate", "templatePath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.TemplatePath = nullValue
    } else {
        request.TemplatePath = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePath", request.TemplatePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePath", request.TemplatePath, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid data
*/
func TestPostSlidesDocumentFromTemplateInvalidData(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.Data, "PostSlidesDocumentFromTemplate", "data", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Data = nullValue
    } else {
        request.Data = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "data", request.Data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "data", request.Data, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePassword
*/
func TestPostSlidesDocumentFromTemplateInvalidTemplatePassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.TemplatePassword, "PostSlidesDocumentFromTemplate", "templatePassword", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.TemplatePassword = nullValue
    } else {
        request.TemplatePassword = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePassword", request.TemplatePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePassword", request.TemplatePassword, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templateStorage
*/
func TestPostSlidesDocumentFromTemplateInvalidTemplateStorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.TemplateStorage, "PostSlidesDocumentFromTemplate", "templateStorage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.TemplateStorage = nullValue
    } else {
        request.TemplateStorage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "templateStorage", request.TemplateStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "templateStorage", request.TemplateStorage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid isImageDataEmbedded
*/
func TestPostSlidesDocumentFromTemplateInvalidIsImageDataEmbedded(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.IsImageDataEmbedded = new(bool)

    invalidValue := invalidizeTestParamValue(request.IsImageDataEmbedded, "PostSlidesDocumentFromTemplate", "isImageDataEmbedded", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.IsImageDataEmbedded = nullValue
    } else {
        *request.IsImageDataEmbedded = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "isImageDataEmbedded", request.IsImageDataEmbedded)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "isImageDataEmbedded", request.IsImageDataEmbedded, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid password
*/
func TestPostSlidesDocumentFromTemplateInvalidPassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesDocumentFromTemplate", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid storage
*/
func TestPostSlidesDocumentFromTemplateInvalidStorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesDocumentFromTemplate", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid folder
*/
func TestPostSlidesDocumentFromTemplateInvalidFolder(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesDocumentFromTemplate", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesDocumentFromTemplate", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesDocumentFromTemplate", "folder", request.Folder, int32(statusCode), e)
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
    request.Pipeline = createTestParamValue("PostSlidesPipeline", "pipeline", "Pipeline").(IPipeline)
    return request
}

/* SlidesApiServiceTests Performs slides pipeline.
   Test for SlidesApi.PostSlidesPipeline method with invalid pipeline
*/
func TestPostSlidesPipelineInvalidPipeline(t *testing.T) {
    request := createPostSlidesPipelineRequest()

    invalidValue := invalidizeTestParamValue(request.Pipeline, "PostSlidesPipeline", "pipeline", "Pipeline")
    if (invalidValue == nil) {
        request.Pipeline = nil
    } else {
        request.Pipeline = invalidValue.(IPipeline)
    }

    e := initializeTest("PostSlidesPipeline", "pipeline", request.Pipeline)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPipeline(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPipeline", "pipeline", request.Pipeline, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesPresentationReplaceText", "name", "string").(string)
    request.OldValue = createTestParamValue("PostSlidesPresentationReplaceText", "oldValue", "string").(string)
    request.NewValue = createTestParamValue("PostSlidesPresentationReplaceText", "newValue", "string").(string)
    testignoreCase := createTestParamValue("PostSlidesPresentationReplaceText", "ignoreCase", "bool")
    switch v := testignoreCase.(type) { 
    case bool:
        request.IgnoreCase = new(bool)
        *request.IgnoreCase = v
    }
    request.Password = createTestParamValue("PostSlidesPresentationReplaceText", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesPresentationReplaceText", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesPresentationReplaceText", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid name
*/
func TestPostSlidesPresentationReplaceTextInvalidName(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesPresentationReplaceText", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid oldValue
*/
func TestPostSlidesPresentationReplaceTextInvalidOldValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.OldValue, "PostSlidesPresentationReplaceText", "oldValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OldValue = nullValue
    } else {
        request.OldValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "oldValue", request.OldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "oldValue", request.OldValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid newValue
*/
func TestPostSlidesPresentationReplaceTextInvalidNewValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.NewValue, "PostSlidesPresentationReplaceText", "newValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.NewValue = nullValue
    } else {
        request.NewValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "newValue", request.NewValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "newValue", request.NewValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid ignoreCase
*/
func TestPostSlidesPresentationReplaceTextInvalidIgnoreCase(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.IgnoreCase = new(bool)

    invalidValue := invalidizeTestParamValue(request.IgnoreCase, "PostSlidesPresentationReplaceText", "ignoreCase", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.IgnoreCase = nullValue
    } else {
        *request.IgnoreCase = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "ignoreCase", request.IgnoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "ignoreCase", request.IgnoreCase, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid password
*/
func TestPostSlidesPresentationReplaceTextInvalidPassword(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesPresentationReplaceText", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid folder
*/
func TestPostSlidesPresentationReplaceTextInvalidFolder(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesPresentationReplaceText", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid storage
*/
func TestPostSlidesPresentationReplaceTextInvalidStorage(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesPresentationReplaceText", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesPresentationReplaceText", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesPresentationReplaceText", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesReorder", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlidesReorder", "slideIndex", "int32").(int32)
    request.NewPosition = createTestParamValue("PostSlidesReorder", "newPosition", "int32").(int32)
    request.Password = createTestParamValue("PostSlidesReorder", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesReorder", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesReorder", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid name
*/
func TestPostSlidesReorderInvalidName(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesReorder", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid slideIndex
*/
func TestPostSlidesReorderInvalidSlideIndex(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlidesReorder", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesReorder", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid newPosition
*/
func TestPostSlidesReorderInvalidNewPosition(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.NewPosition, "PostSlidesReorder", "newPosition", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.NewPosition = nullValue
    } else {
        request.NewPosition = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesReorder", "newPosition", request.NewPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "newPosition", request.NewPosition, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid password
*/
func TestPostSlidesReorderInvalidPassword(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesReorder", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid folder
*/
func TestPostSlidesReorderInvalidFolder(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesReorder", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid storage
*/
func TestPostSlidesReorderInvalidStorage(t *testing.T) {
    request := createPostSlidesReorderRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesReorder", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorder", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorder", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesReorderMany", "name", "string").(string)
    request.OldPositions = createTestParamValue("PostSlidesReorderMany", "oldPositions", "[]int32").([]int32)
    request.NewPositions = createTestParamValue("PostSlidesReorderMany", "newPositions", "[]int32").([]int32)
    request.Password = createTestParamValue("PostSlidesReorderMany", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesReorderMany", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesReorderMany", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid name
*/
func TestPostSlidesReorderManyInvalidName(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesReorderMany", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid oldPositions
*/
func TestPostSlidesReorderManyInvalidOldPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.OldPositions, "PostSlidesReorderMany", "oldPositions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.OldPositions = nullValue
    } else {
        request.OldPositions = invalidValue.([]int32)
    }

    e := initializeTest("PostSlidesReorderMany", "oldPositions", request.OldPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "oldPositions", request.OldPositions, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid newPositions
*/
func TestPostSlidesReorderManyInvalidNewPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.NewPositions, "PostSlidesReorderMany", "newPositions", "[]int32")
    if (invalidValue == nil) {
        var nullValue []int32
        request.NewPositions = nullValue
    } else {
        request.NewPositions = invalidValue.([]int32)
    }

    e := initializeTest("PostSlidesReorderMany", "newPositions", request.NewPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "newPositions", request.NewPositions, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid password
*/
func TestPostSlidesReorderManyInvalidPassword(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesReorderMany", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid folder
*/
func TestPostSlidesReorderManyInvalidFolder(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesReorderMany", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid storage
*/
func TestPostSlidesReorderManyInvalidStorage(t *testing.T) {
    request := createPostSlidesReorderManyRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesReorderMany", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesReorderMany", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesReorderMany", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesSaveAs", "name", "string").(string)
    request.Format = createTestParamValue("PostSlidesSaveAs", "format", "string").(string)
    request.Options = createTestParamValue("PostSlidesSaveAs", "options", "ExportOptions").(IExportOptions)
    request.Password = createTestParamValue("PostSlidesSaveAs", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesSaveAs", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesSaveAs", "folder", "string").(string)
    request.FontsFolder = createTestParamValue("PostSlidesSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid name
*/
func TestPostSlidesSaveAsInvalidName(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid format
*/
func TestPostSlidesSaveAsInvalidFormat(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostSlidesSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid options
*/
func TestPostSlidesSaveAsInvalidOptions(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PostSlidesSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlidesSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid password
*/
func TestPostSlidesSaveAsInvalidPassword(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid storage
*/
func TestPostSlidesSaveAsInvalidStorage(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid folder
*/
func TestPostSlidesSaveAsInvalidFolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid fontsFolder
*/
func TestPostSlidesSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostSlidesSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesSetDocumentProperties", "name", "string").(string)
    request.Properties = createTestParamValue("PostSlidesSetDocumentProperties", "properties", "DocumentProperties").(IDocumentProperties)
    request.Password = createTestParamValue("PostSlidesSetDocumentProperties", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesSetDocumentProperties", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesSetDocumentProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid name
*/
func TestPostSlidesSetDocumentPropertiesInvalidName(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesSetDocumentProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid properties
*/
func TestPostSlidesSetDocumentPropertiesInvalidProperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Properties, "PostSlidesSetDocumentProperties", "properties", "DocumentProperties")
    if (invalidValue == nil) {
        request.Properties = nil
    } else {
        request.Properties = invalidValue.(IDocumentProperties)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "properties", request.Properties)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "properties", request.Properties, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid password
*/
func TestPostSlidesSetDocumentPropertiesInvalidPassword(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesSetDocumentProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid folder
*/
func TestPostSlidesSetDocumentPropertiesInvalidFolder(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesSetDocumentProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid storage
*/
func TestPostSlidesSetDocumentPropertiesInvalidStorage(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesSetDocumentProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSetDocumentProperties", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSetDocumentProperties", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesSlideReplaceText", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PostSlidesSlideReplaceText", "slideIndex", "int32").(int32)
    request.OldValue = createTestParamValue("PostSlidesSlideReplaceText", "oldValue", "string").(string)
    request.NewValue = createTestParamValue("PostSlidesSlideReplaceText", "newValue", "string").(string)
    testignoreCase := createTestParamValue("PostSlidesSlideReplaceText", "ignoreCase", "bool")
    switch v := testignoreCase.(type) { 
    case bool:
        request.IgnoreCase = new(bool)
        *request.IgnoreCase = v
    }
    request.Password = createTestParamValue("PostSlidesSlideReplaceText", "password", "string").(string)
    request.Folder = createTestParamValue("PostSlidesSlideReplaceText", "folder", "string").(string)
    request.Storage = createTestParamValue("PostSlidesSlideReplaceText", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid name
*/
func TestPostSlidesSlideReplaceTextInvalidName(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesSlideReplaceText", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid slideIndex
*/
func TestPostSlidesSlideReplaceTextInvalidSlideIndex(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PostSlidesSlideReplaceText", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid oldValue
*/
func TestPostSlidesSlideReplaceTextInvalidOldValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.OldValue, "PostSlidesSlideReplaceText", "oldValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OldValue = nullValue
    } else {
        request.OldValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "oldValue", request.OldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "oldValue", request.OldValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid newValue
*/
func TestPostSlidesSlideReplaceTextInvalidNewValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.NewValue, "PostSlidesSlideReplaceText", "newValue", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.NewValue = nullValue
    } else {
        request.NewValue = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "newValue", request.NewValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "newValue", request.NewValue, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid ignoreCase
*/
func TestPostSlidesSlideReplaceTextInvalidIgnoreCase(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.IgnoreCase = new(bool)

    invalidValue := invalidizeTestParamValue(request.IgnoreCase, "PostSlidesSlideReplaceText", "ignoreCase", "bool")
    if (invalidValue == nil) {
        var nullValue *bool
        request.IgnoreCase = nullValue
    } else {
        *request.IgnoreCase = invalidValue.(bool)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "ignoreCase", request.IgnoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "ignoreCase", request.IgnoreCase, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid password
*/
func TestPostSlidesSlideReplaceTextInvalidPassword(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesSlideReplaceText", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid folder
*/
func TestPostSlidesSlideReplaceTextInvalidFolder(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesSlideReplaceText", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid storage
*/
func TestPostSlidesSlideReplaceTextInvalidStorage(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesSlideReplaceText", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSlideReplaceText", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSlideReplaceText", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PostSlidesSplit", "name", "string").(string)
    request.Options = createTestParamValue("PostSlidesSplit", "options", "ExportOptions").(IExportOptions)
    request.Format = createTestParamValue("PostSlidesSplit", "format", "string").(string)
    testwidth := createTestParamValue("PostSlidesSplit", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.Width = new(int32)
        *request.Width = v
    }
    testheight := createTestParamValue("PostSlidesSplit", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.Height = new(int32)
        *request.Height = v
    }
    testto := createTestParamValue("PostSlidesSplit", "to", "int32")
    switch v := testto.(type) { 
    case int32:
        request.To = new(int32)
        *request.To = v
    }
    testfrom := createTestParamValue("PostSlidesSplit", "from", "int32")
    switch v := testfrom.(type) { 
    case int32:
        request.From = new(int32)
        *request.From = v
    }
    request.DestFolder = createTestParamValue("PostSlidesSplit", "destFolder", "string").(string)
    request.Password = createTestParamValue("PostSlidesSplit", "password", "string").(string)
    request.Storage = createTestParamValue("PostSlidesSplit", "storage", "string").(string)
    request.Folder = createTestParamValue("PostSlidesSplit", "folder", "string").(string)
    request.FontsFolder = createTestParamValue("PostSlidesSplit", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid name
*/
func TestPostSlidesSplitInvalidName(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PostSlidesSplit", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid options
*/
func TestPostSlidesSplitInvalidOptions(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PostSlidesSplit", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PostSlidesSplit", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid format
*/
func TestPostSlidesSplitInvalidFormat(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PostSlidesSplit", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid width
*/
func TestPostSlidesSplitInvalidWidth(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.Width = new(int32)

    invalidValue := invalidizeTestParamValue(request.Width, "PostSlidesSplit", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Width = nullValue
    } else {
        *request.Width = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "width", request.Width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "width", request.Width, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid height
*/
func TestPostSlidesSplitInvalidHeight(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.Height = new(int32)

    invalidValue := invalidizeTestParamValue(request.Height, "PostSlidesSplit", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Height = nullValue
    } else {
        *request.Height = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "height", request.Height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "height", request.Height, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid to
*/
func TestPostSlidesSplitInvalidTo(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.To = new(int32)

    invalidValue := invalidizeTestParamValue(request.To, "PostSlidesSplit", "to", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.To = nullValue
    } else {
        *request.To = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "to", request.To)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "to", request.To, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid from
*/
func TestPostSlidesSplitInvalidFrom(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.From = new(int32)

    invalidValue := invalidizeTestParamValue(request.From, "PostSlidesSplit", "from", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.From = nullValue
    } else {
        *request.From = invalidValue.(int32)
    }

    e := initializeTest("PostSlidesSplit", "from", request.From)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "from", request.From, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid destFolder
*/
func TestPostSlidesSplitInvalidDestFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.DestFolder, "PostSlidesSplit", "destFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.DestFolder = nullValue
    } else {
        request.DestFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "destFolder", request.DestFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "destFolder", request.DestFolder, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid password
*/
func TestPostSlidesSplitInvalidPassword(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PostSlidesSplit", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid storage
*/
func TestPostSlidesSplitInvalidStorage(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PostSlidesSplit", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid folder
*/
func TestPostSlidesSplitInvalidFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PostSlidesSplit", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid fontsFolder
*/
func TestPostSlidesSplitInvalidFontsFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PostSlidesSplit", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PostSlidesSplit", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PostSlidesSplit", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutLayoutSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutLayoutSlide", "slideIndex", "int32").(int32)
    request.SlideDto = createTestParamValue("PutLayoutSlide", "slideDto", "LayoutSlide").(ILayoutSlide)
    request.Password = createTestParamValue("PutLayoutSlide", "password", "string").(string)
    request.Folder = createTestParamValue("PutLayoutSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("PutLayoutSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid name
*/
func TestPutLayoutSlideInvalidName(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutLayoutSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideIndex
*/
func TestPutLayoutSlideInvalidSlideIndex(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutLayoutSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutLayoutSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideDto
*/
func TestPutLayoutSlideInvalidSlideDto(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideDto, "PutLayoutSlide", "slideDto", "LayoutSlide")
    if (invalidValue == nil) {
        request.SlideDto = nil
    } else {
        request.SlideDto = invalidValue.(ILayoutSlide)
    }

    e := initializeTest("PutLayoutSlide", "slideDto", request.SlideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "slideDto", request.SlideDto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid password
*/
func TestPutLayoutSlideInvalidPassword(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutLayoutSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid folder
*/
func TestPutLayoutSlideInvalidFolder(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutLayoutSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid storage
*/
func TestPutLayoutSlideInvalidStorage(t *testing.T) {
    request := createPutLayoutSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutLayoutSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutLayoutSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutLayoutSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutNotesSlideShapeSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutNotesSlideShapeSaveAs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutNotesSlideShapeSaveAs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutNotesSlideShapeSaveAs", "shapeIndex", "int32").(int32)
    request.Format = createTestParamValue("PutNotesSlideShapeSaveAs", "format", "string").(string)
    request.OutPath = createTestParamValue("PutNotesSlideShapeSaveAs", "outPath", "string").(string)
    request.Options = createTestParamValue("PutNotesSlideShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.Password = createTestParamValue("PutNotesSlideShapeSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PutNotesSlideShapeSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PutNotesSlideShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PutNotesSlideShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.ScaleX = new(float64)
        *request.ScaleX = v
    }
    testscaleY := createTestParamValue("PutNotesSlideShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.ScaleY = new(float64)
        *request.ScaleY = v
    }
    request.Bounds = createTestParamValue("PutNotesSlideShapeSaveAs", "bounds", "string").(string)
    request.FontsFolder = createTestParamValue("PutNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid name
*/
func TestPutNotesSlideShapeSaveAsInvalidName(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutNotesSlideShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutNotesSlideShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid path
*/
func TestPutNotesSlideShapeSaveAsInvalidPath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutNotesSlideShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidShapeIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutNotesSlideShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid format
*/
func TestPutNotesSlideShapeSaveAsInvalidFormat(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PutNotesSlideShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid outPath
*/
func TestPutNotesSlideShapeSaveAsInvalidOutPath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.OutPath, "PutNotesSlideShapeSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OutPath = nullValue
    } else {
        request.OutPath = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "outPath", request.OutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "outPath", request.OutPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid options
*/
func TestPutNotesSlideShapeSaveAsInvalidOptions(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PutNotesSlideShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid password
*/
func TestPutNotesSlideShapeSaveAsInvalidPassword(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutNotesSlideShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid folder
*/
func TestPutNotesSlideShapeSaveAsInvalidFolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutNotesSlideShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid storage
*/
func TestPutNotesSlideShapeSaveAsInvalidStorage(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutNotesSlideShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPutNotesSlideShapeSaveAsInvalidScaleX(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.ScaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleX, "PutNotesSlideShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleX = nullValue
    } else {
        *request.ScaleX = invalidValue.(float64)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleX", request.ScaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleX", request.ScaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPutNotesSlideShapeSaveAsInvalidScaleY(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.ScaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleY, "PutNotesSlideShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleY = nullValue
    } else {
        *request.ScaleY = invalidValue.(float64)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleY", request.ScaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleY", request.ScaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPutNotesSlideShapeSaveAsInvalidBounds(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Bounds, "PutNotesSlideShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Bounds = nullValue
    } else {
        request.Bounds = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "bounds", request.Bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "bounds", request.Bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPutNotesSlideShapeSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PutNotesSlideShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutNotesSlideShapeSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutNotesSlideShapeSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutPresentationMerge", "name", "string").(string)
    request.Request = createTestParamValue("PutPresentationMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    request.Password = createTestParamValue("PutPresentationMerge", "password", "string").(string)
    request.Storage = createTestParamValue("PutPresentationMerge", "storage", "string").(string)
    request.Folder = createTestParamValue("PutPresentationMerge", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid name
*/
func TestPutPresentationMergeInvalidName(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutPresentationMerge", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid request
*/
func TestPutPresentationMergeInvalidRequest(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Request, "PutPresentationMerge", "request", "OrderedMergeRequest")
    if (invalidValue == nil) {
        request.Request = nil
    } else {
        request.Request = invalidValue.(IOrderedMergeRequest)
    }

    e := initializeTest("PutPresentationMerge", "request", request.Request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "request", request.Request, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid password
*/
func TestPutPresentationMergeInvalidPassword(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutPresentationMerge", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid storage
*/
func TestPutPresentationMergeInvalidStorage(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutPresentationMerge", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid folder
*/
func TestPutPresentationMergeInvalidFolder(t *testing.T) {
    request := createPutPresentationMergeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutPresentationMerge", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutPresentationMerge", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutPresentationMerge", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSetParagraphPortionProperties", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSetParagraphPortionProperties", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutSetParagraphPortionProperties", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutSetParagraphPortionProperties", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PutSetParagraphPortionProperties", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("PutSetParagraphPortionProperties", "portionIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutSetParagraphPortionProperties", "dto", "Portion").(IPortion)
    request.Password = createTestParamValue("PutSetParagraphPortionProperties", "password", "string").(string)
    request.Folder = createTestParamValue("PutSetParagraphPortionProperties", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSetParagraphPortionProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid name
*/
func TestPutSetParagraphPortionPropertiesInvalidName(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSetParagraphPortionProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid slideIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidSlideIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSetParagraphPortionProperties", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid path
*/
func TestPutSetParagraphPortionPropertiesInvalidPath(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutSetParagraphPortionProperties", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidShapeIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutSetParagraphPortionProperties", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidParagraphIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PutSetParagraphPortionProperties", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid portionIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidPortionIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "PutSetParagraphPortionProperties", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid dto
*/
func TestPutSetParagraphPortionPropertiesInvalidDto(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutSetParagraphPortionProperties", "dto", "Portion")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IPortion)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid password
*/
func TestPutSetParagraphPortionPropertiesInvalidPassword(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSetParagraphPortionProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid folder
*/
func TestPutSetParagraphPortionPropertiesInvalidFolder(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSetParagraphPortionProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid storage
*/
func TestPutSetParagraphPortionPropertiesInvalidStorage(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSetParagraphPortionProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphPortionProperties", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphPortionProperties", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSetParagraphProperties", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSetParagraphProperties", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutSetParagraphProperties", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutSetParagraphProperties", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PutSetParagraphProperties", "paragraphIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutSetParagraphProperties", "dto", "Paragraph").(IParagraph)
    request.Password = createTestParamValue("PutSetParagraphProperties", "password", "string").(string)
    request.Folder = createTestParamValue("PutSetParagraphProperties", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSetParagraphProperties", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid name
*/
func TestPutSetParagraphPropertiesInvalidName(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSetParagraphProperties", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid slideIndex
*/
func TestPutSetParagraphPropertiesInvalidSlideIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSetParagraphProperties", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid path
*/
func TestPutSetParagraphPropertiesInvalidPath(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutSetParagraphProperties", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPropertiesInvalidShapeIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutSetParagraphProperties", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPropertiesInvalidParagraphIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PutSetParagraphProperties", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSetParagraphProperties", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid dto
*/
func TestPutSetParagraphPropertiesInvalidDto(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutSetParagraphProperties", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PutSetParagraphProperties", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid password
*/
func TestPutSetParagraphPropertiesInvalidPassword(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSetParagraphProperties", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid folder
*/
func TestPutSetParagraphPropertiesInvalidFolder(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSetParagraphProperties", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid storage
*/
func TestPutSetParagraphPropertiesInvalidStorage(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSetParagraphProperties", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSetParagraphProperties", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSetParagraphProperties", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutShapeSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutShapeSaveAs", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutShapeSaveAs", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutShapeSaveAs", "shapeIndex", "int32").(int32)
    request.Format = createTestParamValue("PutShapeSaveAs", "format", "string").(string)
    request.OutPath = createTestParamValue("PutShapeSaveAs", "outPath", "string").(string)
    request.Options = createTestParamValue("PutShapeSaveAs", "options", "IShapeExportOptions").(IIShapeExportOptions)
    request.Password = createTestParamValue("PutShapeSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PutShapeSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PutShapeSaveAs", "storage", "string").(string)
    testscaleX := createTestParamValue("PutShapeSaveAs", "scaleX", "float64")
    switch v := testscaleX.(type) { 
    case float64:
        request.ScaleX = new(float64)
        *request.ScaleX = v
    }
    testscaleY := createTestParamValue("PutShapeSaveAs", "scaleY", "float64")
    switch v := testscaleY.(type) { 
    case float64:
        request.ScaleY = new(float64)
        *request.ScaleY = v
    }
    request.Bounds = createTestParamValue("PutShapeSaveAs", "bounds", "string").(string)
    request.FontsFolder = createTestParamValue("PutShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid name
*/
func TestPutShapeSaveAsInvalidName(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutShapeSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid slideIndex
*/
func TestPutShapeSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutShapeSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutShapeSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid path
*/
func TestPutShapeSaveAsInvalidPath(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutShapeSaveAs", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid shapeIndex
*/
func TestPutShapeSaveAsInvalidShapeIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutShapeSaveAs", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutShapeSaveAs", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid format
*/
func TestPutShapeSaveAsInvalidFormat(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PutShapeSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid outPath
*/
func TestPutShapeSaveAsInvalidOutPath(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.OutPath, "PutShapeSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OutPath = nullValue
    } else {
        request.OutPath = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "outPath", request.OutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "outPath", request.OutPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid options
*/
func TestPutShapeSaveAsInvalidOptions(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PutShapeSaveAs", "options", "IShapeExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IIShapeExportOptions)
    }

    e := initializeTest("PutShapeSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid password
*/
func TestPutShapeSaveAsInvalidPassword(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutShapeSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid folder
*/
func TestPutShapeSaveAsInvalidFolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutShapeSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid storage
*/
func TestPutShapeSaveAsInvalidStorage(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutShapeSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleX
*/
func TestPutShapeSaveAsInvalidScaleX(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.ScaleX = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleX, "PutShapeSaveAs", "scaleX", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleX = nullValue
    } else {
        *request.ScaleX = invalidValue.(float64)
    }

    e := initializeTest("PutShapeSaveAs", "scaleX", request.ScaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "scaleX", request.ScaleX, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleY
*/
func TestPutShapeSaveAsInvalidScaleY(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.ScaleY = new(float64)

    invalidValue := invalidizeTestParamValue(request.ScaleY, "PutShapeSaveAs", "scaleY", "float64")
    if (invalidValue == nil) {
        var nullValue *float64
        request.ScaleY = nullValue
    } else {
        *request.ScaleY = invalidValue.(float64)
    }

    e := initializeTest("PutShapeSaveAs", "scaleY", request.ScaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "scaleY", request.ScaleY, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid bounds
*/
func TestPutShapeSaveAsInvalidBounds(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Bounds, "PutShapeSaveAs", "bounds", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Bounds = nullValue
    } else {
        request.Bounds = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "bounds", request.Bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "bounds", request.Bounds, int32(statusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid fontsFolder
*/
func TestPutShapeSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PutShapeSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutShapeSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutShapeSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlideAnimation", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlideAnimation", "slideIndex", "int32").(int32)
    request.Animation = createTestParamValue("PutSlideAnimation", "animation", "SlideAnimation").(ISlideAnimation)
    request.Password = createTestParamValue("PutSlideAnimation", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlideAnimation", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlideAnimation", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid name
*/
func TestPutSlideAnimationInvalidName(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlideAnimation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid slideIndex
*/
func TestPutSlideAnimationInvalidSlideIndex(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlideAnimation", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimation", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid animation
*/
func TestPutSlideAnimationInvalidAnimation(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Animation, "PutSlideAnimation", "animation", "SlideAnimation")
    if (invalidValue == nil) {
        request.Animation = nil
    } else {
        request.Animation = invalidValue.(ISlideAnimation)
    }

    e := initializeTest("PutSlideAnimation", "animation", request.Animation)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "animation", request.Animation, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid password
*/
func TestPutSlideAnimationInvalidPassword(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlideAnimation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid folder
*/
func TestPutSlideAnimationInvalidFolder(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlideAnimation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide animation.
   Test for SlidesApi.PutSlideAnimation method with invalid storage
*/
func TestPutSlideAnimationInvalidStorage(t *testing.T) {
    request := createPutSlideAnimationRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlideAnimation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimation", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimation(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimation", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlideAnimationEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlideAnimationEffect", "slideIndex", "int32").(int32)
    request.EffectIndex = createTestParamValue("PutSlideAnimationEffect", "effectIndex", "int32").(int32)
    request.Effect = createTestParamValue("PutSlideAnimationEffect", "effect", "Effect").(IEffect)
    request.Password = createTestParamValue("PutSlideAnimationEffect", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlideAnimationEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlideAnimationEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid name
*/
func TestPutSlideAnimationEffectInvalidName(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlideAnimationEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid slideIndex
*/
func TestPutSlideAnimationEffectInvalidSlideIndex(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlideAnimationEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid effectIndex
*/
func TestPutSlideAnimationEffectInvalidEffectIndex(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.EffectIndex, "PutSlideAnimationEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.EffectIndex = nullValue
    } else {
        request.EffectIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationEffect", "effectIndex", request.EffectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "effectIndex", request.EffectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid effect
*/
func TestPutSlideAnimationEffectInvalidEffect(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Effect, "PutSlideAnimationEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.Effect = nil
    } else {
        request.Effect = invalidValue.(IEffect)
    }

    e := initializeTest("PutSlideAnimationEffect", "effect", request.Effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "effect", request.Effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid password
*/
func TestPutSlideAnimationEffectInvalidPassword(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlideAnimationEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid folder
*/
func TestPutSlideAnimationEffectInvalidFolder(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlideAnimationEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide.
   Test for SlidesApi.PutSlideAnimationEffect method with invalid storage
*/
func TestPutSlideAnimationEffectInvalidStorage(t *testing.T) {
    request := createPutSlideAnimationEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlideAnimationEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32").(int32)
    request.SequenceIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32").(int32)
    request.EffectIndex = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32").(int32)
    request.Effect = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "effect", "Effect").(IEffect)
    request.Password = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlideAnimationInteractiveSequenceEffect", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid name
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidName(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlideAnimationInteractiveSequenceEffect", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid slideIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidSlideIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlideAnimationInteractiveSequenceEffect", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid sequenceIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidSequenceIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.SequenceIndex, "PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SequenceIndex = nullValue
    } else {
        request.SequenceIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "sequenceIndex", request.SequenceIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid effectIndex
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidEffectIndex(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.EffectIndex, "PutSlideAnimationInteractiveSequenceEffect", "effectIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.EffectIndex = nullValue
    } else {
        request.EffectIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "effectIndex", request.EffectIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "effectIndex", request.EffectIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid effect
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidEffect(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Effect, "PutSlideAnimationInteractiveSequenceEffect", "effect", "Effect")
    if (invalidValue == nil) {
        request.Effect = nil
    } else {
        request.Effect = invalidValue.(IEffect)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "effect", request.Effect)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "effect", request.Effect, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid password
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidPassword(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlideAnimationInteractiveSequenceEffect", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid folder
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidFolder(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlideAnimationInteractiveSequenceEffect", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Modify an animation effect for a slide interactive sequence.
   Test for SlidesApi.PutSlideAnimationInteractiveSequenceEffect method with invalid storage
*/
func TestPutSlideAnimationInteractiveSequenceEffectInvalidStorage(t *testing.T) {
    request := createPutSlideAnimationInteractiveSequenceEffectRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlideAnimationInteractiveSequenceEffect", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideAnimationInteractiveSequenceEffect", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideAnimationInteractiveSequenceEffect(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideAnimationInteractiveSequenceEffect", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlideSaveAs", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlideSaveAs", "slideIndex", "int32").(int32)
    request.Format = createTestParamValue("PutSlideSaveAs", "format", "string").(string)
    request.OutPath = createTestParamValue("PutSlideSaveAs", "outPath", "string").(string)
    request.Options = createTestParamValue("PutSlideSaveAs", "options", "ExportOptions").(IExportOptions)
    testwidth := createTestParamValue("PutSlideSaveAs", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.Width = new(int32)
        *request.Width = v
    }
    testheight := createTestParamValue("PutSlideSaveAs", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.Height = new(int32)
        *request.Height = v
    }
    request.Password = createTestParamValue("PutSlideSaveAs", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlideSaveAs", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlideSaveAs", "storage", "string").(string)
    request.FontsFolder = createTestParamValue("PutSlideSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid name
*/
func TestPutSlideSaveAsInvalidName(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlideSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid slideIndex
*/
func TestPutSlideSaveAsInvalidSlideIndex(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlideSaveAs", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid format
*/
func TestPutSlideSaveAsInvalidFormat(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PutSlideSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid outPath
*/
func TestPutSlideSaveAsInvalidOutPath(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.OutPath, "PutSlideSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OutPath = nullValue
    } else {
        request.OutPath = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "outPath", request.OutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "outPath", request.OutPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid options
*/
func TestPutSlideSaveAsInvalidOptions(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PutSlideSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PutSlideSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid width
*/
func TestPutSlideSaveAsInvalidWidth(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.Width = new(int32)

    invalidValue := invalidizeTestParamValue(request.Width, "PutSlideSaveAs", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Width = nullValue
    } else {
        *request.Width = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "width", request.Width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "width", request.Width, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid height
*/
func TestPutSlideSaveAsInvalidHeight(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.Height = new(int32)

    invalidValue := invalidizeTestParamValue(request.Height, "PutSlideSaveAs", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Height = nullValue
    } else {
        *request.Height = invalidValue.(int32)
    }

    e := initializeTest("PutSlideSaveAs", "height", request.Height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "height", request.Height, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid password
*/
func TestPutSlideSaveAsInvalidPassword(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlideSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid folder
*/
func TestPutSlideSaveAsInvalidFolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlideSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid storage
*/
func TestPutSlideSaveAsInvalidStorage(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlideSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid fontsFolder
*/
func TestPutSlideSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PutSlideSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlideSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlideShapeInfo", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlideShapeInfo", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutSlideShapeInfo", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutSlideShapeInfo", "shapeIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutSlideShapeInfo", "dto", "ShapeBase").(IShapeBase)
    request.Password = createTestParamValue("PutSlideShapeInfo", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlideShapeInfo", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlideShapeInfo", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid name
*/
func TestPutSlideShapeInfoInvalidName(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlideShapeInfo", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid slideIndex
*/
func TestPutSlideShapeInfoInvalidSlideIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlideShapeInfo", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideShapeInfo", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid path
*/
func TestPutSlideShapeInfoInvalidPath(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutSlideShapeInfo", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid shapeIndex
*/
func TestPutSlideShapeInfoInvalidShapeIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutSlideShapeInfo", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlideShapeInfo", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid dto
*/
func TestPutSlideShapeInfoInvalidDto(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutSlideShapeInfo", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PutSlideShapeInfo", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid password
*/
func TestPutSlideShapeInfoInvalidPassword(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlideShapeInfo", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid folder
*/
func TestPutSlideShapeInfoInvalidFolder(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlideShapeInfo", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid storage
*/
func TestPutSlideShapeInfoInvalidStorage(t *testing.T) {
    request := createPutSlideShapeInfoRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlideShapeInfo", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlideShapeInfo", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlideShapeInfo", "storage", request.Storage, int32(statusCode), e)
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
    request.Format = createTestParamValue("PutSlidesConvert", "format", "string").(string)
    request.OutPath = createTestParamValue("PutSlidesConvert", "outPath", "string").(string)
    request.Document = createTestParamValue("PutSlidesConvert", "document", "[]byte").([]byte)
    request.Password = createTestParamValue("PutSlidesConvert", "password", "string").(string)
    request.FontsFolder = createTestParamValue("PutSlidesConvert", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid format
*/
func TestPutSlidesConvertInvalidFormat(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PutSlidesConvert", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid outPath
*/
func TestPutSlidesConvertInvalidOutPath(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.OutPath, "PutSlidesConvert", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OutPath = nullValue
    } else {
        request.OutPath = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "outPath", request.OutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "outPath", request.OutPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid document
*/
func TestPutSlidesConvertInvalidDocument(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Document, "PutSlidesConvert", "document", "[]byte")
    if (invalidValue == nil) {
        request.Document = nil
    } else {
        request.Document = invalidValue.([]byte)
    }

    e := initializeTest("PutSlidesConvert", "document", request.Document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "document", request.Document, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid password
*/
func TestPutSlidesConvertInvalidPassword(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesConvert", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid fontsFolder
*/
func TestPutSlidesConvertInvalidFontsFolder(t *testing.T) {
    request := createPutSlidesConvertRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PutSlidesConvert", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesConvert", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesConvert", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesDocumentFromHtml", "name", "string").(string)
    request.Html = createTestParamValue("PutSlidesDocumentFromHtml", "html", "string").(string)
    request.Password = createTestParamValue("PutSlidesDocumentFromHtml", "password", "string").(string)
    request.Storage = createTestParamValue("PutSlidesDocumentFromHtml", "storage", "string").(string)
    request.Folder = createTestParamValue("PutSlidesDocumentFromHtml", "folder", "string").(string)
    return request
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid name
*/
func TestPutSlidesDocumentFromHtmlInvalidName(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesDocumentFromHtml", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid html
*/
func TestPutSlidesDocumentFromHtmlInvalidHtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Html, "PutSlidesDocumentFromHtml", "html", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Html = nullValue
    } else {
        request.Html = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "html", request.Html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "html", request.Html, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid password
*/
func TestPutSlidesDocumentFromHtmlInvalidPassword(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesDocumentFromHtml", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid storage
*/
func TestPutSlidesDocumentFromHtmlInvalidStorage(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesDocumentFromHtml", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid folder
*/
func TestPutSlidesDocumentFromHtmlInvalidFolder(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesDocumentFromHtml", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesDocumentFromHtml", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesDocumentFromHtml", "folder", request.Folder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSaveAs", "name", "string").(string)
    request.OutPath = createTestParamValue("PutSlidesSaveAs", "outPath", "string").(string)
    request.Format = createTestParamValue("PutSlidesSaveAs", "format", "string").(string)
    request.Options = createTestParamValue("PutSlidesSaveAs", "options", "ExportOptions").(IExportOptions)
    request.Password = createTestParamValue("PutSlidesSaveAs", "password", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSaveAs", "storage", "string").(string)
    request.Folder = createTestParamValue("PutSlidesSaveAs", "folder", "string").(string)
    request.FontsFolder = createTestParamValue("PutSlidesSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid name
*/
func TestPutSlidesSaveAsInvalidName(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSaveAs", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid outPath
*/
func TestPutSlidesSaveAsInvalidOutPath(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.OutPath, "PutSlidesSaveAs", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.OutPath = nullValue
    } else {
        request.OutPath = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "outPath", request.OutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "outPath", request.OutPath, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid format
*/
func TestPutSlidesSaveAsInvalidFormat(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Format, "PutSlidesSaveAs", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Format = nullValue
    } else {
        request.Format = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "format", request.Format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "format", request.Format, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid options
*/
func TestPutSlidesSaveAsInvalidOptions(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Options, "PutSlidesSaveAs", "options", "ExportOptions")
    if (invalidValue == nil) {
        request.Options = nil
    } else {
        request.Options = invalidValue.(IExportOptions)
    }

    e := initializeTest("PutSlidesSaveAs", "options", request.Options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "options", request.Options, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid password
*/
func TestPutSlidesSaveAsInvalidPassword(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSaveAs", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid storage
*/
func TestPutSlidesSaveAsInvalidStorage(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSaveAs", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid folder
*/
func TestPutSlidesSaveAsInvalidFolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSaveAs", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid fontsFolder
*/
func TestPutSlidesSaveAsInvalidFontsFolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()

    invalidValue := invalidizeTestParamValue(request.FontsFolder, "PutSlidesSaveAs", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.FontsFolder = nullValue
    } else {
        request.FontsFolder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSaveAs", "fontsFolder", request.FontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSaveAs", "fontsFolder", request.FontsFolder, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSetDocumentProperty", "name", "string").(string)
    request.PropertyName = createTestParamValue("PutSlidesSetDocumentProperty", "propertyName", "string").(string)
    request.Property = createTestParamValue("PutSlidesSetDocumentProperty", "property", "DocumentProperty").(IDocumentProperty)
    request.Password = createTestParamValue("PutSlidesSetDocumentProperty", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlidesSetDocumentProperty", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSetDocumentProperty", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid name
*/
func TestPutSlidesSetDocumentPropertyInvalidName(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSetDocumentProperty", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid propertyName
*/
func TestPutSlidesSetDocumentPropertyInvalidPropertyName(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.PropertyName, "PutSlidesSetDocumentProperty", "propertyName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.PropertyName = nullValue
    } else {
        request.PropertyName = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "propertyName", request.PropertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "propertyName", request.PropertyName, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid property
*/
func TestPutSlidesSetDocumentPropertyInvalidProperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Property, "PutSlidesSetDocumentProperty", "property", "DocumentProperty")
    if (invalidValue == nil) {
        request.Property = nil
    } else {
        request.Property = invalidValue.(IDocumentProperty)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "property", request.Property)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "property", request.Property, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid password
*/
func TestPutSlidesSetDocumentPropertyInvalidPassword(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSetDocumentProperty", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid folder
*/
func TestPutSlidesSetDocumentPropertyInvalidFolder(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSetDocumentProperty", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid storage
*/
func TestPutSlidesSetDocumentPropertyInvalidStorage(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSetDocumentProperty", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSetDocumentProperty", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSetDocumentProperty", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlidesSlide", "slideIndex", "int32").(int32)
    request.SlideDto = createTestParamValue("PutSlidesSlide", "slideDto", "Slide").(ISlide)
    request.Password = createTestParamValue("PutSlidesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("PutSlidesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid name
*/
func TestPutSlidesSlideInvalidName(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideIndex
*/
func TestPutSlidesSlideInvalidSlideIndex(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlidesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideDto
*/
func TestPutSlidesSlideInvalidSlideDto(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideDto, "PutSlidesSlide", "slideDto", "Slide")
    if (invalidValue == nil) {
        request.SlideDto = nil
    } else {
        request.SlideDto = invalidValue.(ISlide)
    }

    e := initializeTest("PutSlidesSlide", "slideDto", request.SlideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "slideDto", request.SlideDto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid password
*/
func TestPutSlidesSlideInvalidPassword(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid folder
*/
func TestPutSlidesSlideInvalidFolder(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid storage
*/
func TestPutSlidesSlideInvalidStorage(t *testing.T) {
    request := createPutSlidesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSlideBackground", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.Background = createTestParamValue("PutSlidesSlideBackground", "background", "SlideBackground").(ISlideBackground)
    request.Folder = createTestParamValue("PutSlidesSlideBackground", "folder", "string").(string)
    request.Password = createTestParamValue("PutSlidesSlideBackground", "password", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid name
*/
func TestPutSlidesSlideBackgroundInvalidName(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSlideBackground", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundInvalidSlideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlidesSlideBackground", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideBackground", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid background
*/
func TestPutSlidesSlideBackgroundInvalidBackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Background, "PutSlidesSlideBackground", "background", "SlideBackground")
    if (invalidValue == nil) {
        request.Background = nil
    } else {
        request.Background = invalidValue.(ISlideBackground)
    }

    e := initializeTest("PutSlidesSlideBackground", "background", request.Background)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "background", request.Background, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid folder
*/
func TestPutSlidesSlideBackgroundInvalidFolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSlideBackground", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid password
*/
func TestPutSlidesSlideBackgroundInvalidPassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSlideBackground", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid storage
*/
func TestPutSlidesSlideBackgroundInvalidStorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSlideBackground", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackground", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackground", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSlideBackgroundColor", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutSlidesSlideBackgroundColor", "slideIndex", "int32").(int32)
    request.Color = createTestParamValue("PutSlidesSlideBackgroundColor", "color", "string").(string)
    request.Folder = createTestParamValue("PutSlidesSlideBackgroundColor", "folder", "string").(string)
    request.Password = createTestParamValue("PutSlidesSlideBackgroundColor", "password", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSlideBackgroundColor", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid name
*/
func TestPutSlidesSlideBackgroundColorInvalidName(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSlideBackgroundColor", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundColorInvalidSlideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutSlidesSlideBackgroundColor", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid color
*/
func TestPutSlidesSlideBackgroundColorInvalidColor(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.Color, "PutSlidesSlideBackgroundColor", "color", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Color = nullValue
    } else {
        request.Color = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "color", request.Color)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "color", request.Color, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid folder
*/
func TestPutSlidesSlideBackgroundColorInvalidFolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSlideBackgroundColor", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid password
*/
func TestPutSlidesSlideBackgroundColorInvalidPassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSlideBackgroundColor", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid storage
*/
func TestPutSlidesSlideBackgroundColorInvalidStorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSlideBackgroundColor", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideBackgroundColor", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideBackgroundColor", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutSlidesSlideSize", "name", "string").(string)
    request.Password = createTestParamValue("PutSlidesSlideSize", "password", "string").(string)
    request.Storage = createTestParamValue("PutSlidesSlideSize", "storage", "string").(string)
    request.Folder = createTestParamValue("PutSlidesSlideSize", "folder", "string").(string)
    testwidth := createTestParamValue("PutSlidesSlideSize", "width", "int32")
    switch v := testwidth.(type) { 
    case int32:
        request.Width = new(int32)
        *request.Width = v
    }
    testheight := createTestParamValue("PutSlidesSlideSize", "height", "int32")
    switch v := testheight.(type) { 
    case int32:
        request.Height = new(int32)
        *request.Height = v
    }
    request.SizeType = createTestParamValue("PutSlidesSlideSize", "sizeType", "string").(string)
    request.ScaleType = createTestParamValue("PutSlidesSlideSize", "scaleType", "string").(string)
    return request
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid name
*/
func TestPutSlidesSlideSizeInvalidName(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutSlidesSlideSize", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid password
*/
func TestPutSlidesSlideSizeInvalidPassword(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutSlidesSlideSize", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid storage
*/
func TestPutSlidesSlideSizeInvalidStorage(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutSlidesSlideSize", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "storage", request.Storage, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid folder
*/
func TestPutSlidesSlideSizeInvalidFolder(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutSlidesSlideSize", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid width
*/
func TestPutSlidesSlideSizeInvalidWidth(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.Width = new(int32)

    invalidValue := invalidizeTestParamValue(request.Width, "PutSlidesSlideSize", "width", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Width = nullValue
    } else {
        *request.Width = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideSize", "width", request.Width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "width", request.Width, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid height
*/
func TestPutSlidesSlideSizeInvalidHeight(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.Height = new(int32)

    invalidValue := invalidizeTestParamValue(request.Height, "PutSlidesSlideSize", "height", "int32")
    if (invalidValue == nil) {
        var nullValue *int32
        request.Height = nullValue
    } else {
        *request.Height = invalidValue.(int32)
    }

    e := initializeTest("PutSlidesSlideSize", "height", request.Height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "height", request.Height, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid sizeType
*/
func TestPutSlidesSlideSizeInvalidSizeType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.SizeType, "PutSlidesSlideSize", "sizeType", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.SizeType = nullValue
    } else {
        request.SizeType = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "sizeType", request.SizeType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "sizeType", request.SizeType, int32(statusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid scaleType
*/
func TestPutSlidesSlideSizeInvalidScaleType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()

    invalidValue := invalidizeTestParamValue(request.ScaleType, "PutSlidesSlideSize", "scaleType", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.ScaleType = nullValue
    } else {
        request.ScaleType = invalidValue.(string)
    }

    e := initializeTest("PutSlidesSlideSize", "scaleType", request.ScaleType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutSlidesSlideSize", "scaleType", request.ScaleType, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutUpdateNotesSlide", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutUpdateNotesSlide", "slideIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutUpdateNotesSlide", "dto", "NotesSlide").(INotesSlide)
    request.Password = createTestParamValue("PutUpdateNotesSlide", "password", "string").(string)
    request.Folder = createTestParamValue("PutUpdateNotesSlide", "folder", "string").(string)
    request.Storage = createTestParamValue("PutUpdateNotesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid name
*/
func TestPutUpdateNotesSlideInvalidName(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutUpdateNotesSlide", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid slideIndex
*/
func TestPutUpdateNotesSlideInvalidSlideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutUpdateNotesSlide", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlide", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid dto
*/
func TestPutUpdateNotesSlideInvalidDto(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutUpdateNotesSlide", "dto", "NotesSlide")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(INotesSlide)
    }

    e := initializeTest("PutUpdateNotesSlide", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid password
*/
func TestPutUpdateNotesSlideInvalidPassword(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutUpdateNotesSlide", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid folder
*/
func TestPutUpdateNotesSlideInvalidFolder(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutUpdateNotesSlide", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid storage
*/
func TestPutUpdateNotesSlideInvalidStorage(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutUpdateNotesSlide", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlide", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlide", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutUpdateNotesSlideShape", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutUpdateNotesSlideShape", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutUpdateNotesSlideShape", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutUpdateNotesSlideShape", "shapeIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutUpdateNotesSlideShape", "dto", "ShapeBase").(IShapeBase)
    request.Password = createTestParamValue("PutUpdateNotesSlideShape", "password", "string").(string)
    request.Folder = createTestParamValue("PutUpdateNotesSlideShape", "folder", "string").(string)
    request.Storage = createTestParamValue("PutUpdateNotesSlideShape", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid name
*/
func TestPutUpdateNotesSlideShapeInvalidName(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutUpdateNotesSlideShape", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeInvalidSlideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutUpdateNotesSlideShape", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid path
*/
func TestPutUpdateNotesSlideShapeInvalidPath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutUpdateNotesSlideShape", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeInvalidShapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutUpdateNotesSlideShape", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid dto
*/
func TestPutUpdateNotesSlideShapeInvalidDto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutUpdateNotesSlideShape", "dto", "ShapeBase")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IShapeBase)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid password
*/
func TestPutUpdateNotesSlideShapeInvalidPassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutUpdateNotesSlideShape", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid folder
*/
func TestPutUpdateNotesSlideShapeInvalidFolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutUpdateNotesSlideShape", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid storage
*/
func TestPutUpdateNotesSlideShapeInvalidStorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutUpdateNotesSlideShape", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShape", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShape", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "dto", "Paragraph").(IParagraph)
    request.Password = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "password", "string").(string)
    request.Folder = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "folder", "string").(string)
    request.Storage = createTestParamValue("PutUpdateNotesSlideShapeParagraph", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid name
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidName(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutUpdateNotesSlideShapeParagraph", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidSlideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutUpdateNotesSlideShapeParagraph", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid path
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidPath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutUpdateNotesSlideShapeParagraph", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidShapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidParagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid dto
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidDto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutUpdateNotesSlideShapeParagraph", "dto", "Paragraph")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IParagraph)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid password
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidPassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutUpdateNotesSlideShapeParagraph", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid folder
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidFolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutUpdateNotesSlideShapeParagraph", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid storage
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidStorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutUpdateNotesSlideShapeParagraph", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "storage", request.Storage, int32(statusCode), e)
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
    request.Name = createTestParamValue("PutUpdateNotesSlideShapePortion", "name", "string").(string)
    request.SlideIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "slideIndex", "int32").(int32)
    request.Path = createTestParamValue("PutUpdateNotesSlideShapePortion", "path", "string").(string)
    request.ShapeIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "shapeIndex", "int32").(int32)
    request.ParagraphIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "paragraphIndex", "int32").(int32)
    request.PortionIndex = createTestParamValue("PutUpdateNotesSlideShapePortion", "portionIndex", "int32").(int32)
    request.Dto = createTestParamValue("PutUpdateNotesSlideShapePortion", "dto", "Portion").(IPortion)
    request.Password = createTestParamValue("PutUpdateNotesSlideShapePortion", "password", "string").(string)
    request.Folder = createTestParamValue("PutUpdateNotesSlideShapePortion", "folder", "string").(string)
    request.Storage = createTestParamValue("PutUpdateNotesSlideShapePortion", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid name
*/
func TestPutUpdateNotesSlideShapePortionInvalidName(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Name, "PutUpdateNotesSlideShapePortion", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Name = nullValue
    } else {
        request.Name = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "name", request.Name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "name", request.Name, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidSlideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.SlideIndex, "PutUpdateNotesSlideShapePortion", "slideIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.SlideIndex = nullValue
    } else {
        request.SlideIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "slideIndex", request.SlideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "slideIndex", request.SlideIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid path
*/
func TestPutUpdateNotesSlideShapePortionInvalidPath(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "PutUpdateNotesSlideShapePortion", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidShapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ShapeIndex, "PutUpdateNotesSlideShapePortion", "shapeIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ShapeIndex = nullValue
    } else {
        request.ShapeIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "shapeIndex", request.ShapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "shapeIndex", request.ShapeIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidParagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.ParagraphIndex, "PutUpdateNotesSlideShapePortion", "paragraphIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.ParagraphIndex = nullValue
    } else {
        request.ParagraphIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "paragraphIndex", request.ParagraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "paragraphIndex", request.ParagraphIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid portionIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidPortionIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.PortionIndex, "PutUpdateNotesSlideShapePortion", "portionIndex", "int32")
    if (invalidValue == nil) {
        var nullValue int32
        request.PortionIndex = nullValue
    } else {
        request.PortionIndex = invalidValue.(int32)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "portionIndex", request.PortionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "portionIndex", request.PortionIndex, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid dto
*/
func TestPutUpdateNotesSlideShapePortionInvalidDto(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Dto, "PutUpdateNotesSlideShapePortion", "dto", "Portion")
    if (invalidValue == nil) {
        request.Dto = nil
    } else {
        request.Dto = invalidValue.(IPortion)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "dto", request.Dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "dto", request.Dto, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid password
*/
func TestPutUpdateNotesSlideShapePortionInvalidPassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Password, "PutUpdateNotesSlideShapePortion", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Password = nullValue
    } else {
        request.Password = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "password", request.Password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "password", request.Password, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid folder
*/
func TestPutUpdateNotesSlideShapePortionInvalidFolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Folder, "PutUpdateNotesSlideShapePortion", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Folder = nullValue
    } else {
        request.Folder = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "folder", request.Folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "folder", request.Folder, int32(statusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid storage
*/
func TestPutUpdateNotesSlideShapePortionInvalidStorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()

    invalidValue := invalidizeTestParamValue(request.Storage, "PutUpdateNotesSlideShapePortion", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Storage = nullValue
    } else {
        request.Storage = invalidValue.(string)
    }

    e := initializeTest("PutUpdateNotesSlideShapePortion", "storage", request.Storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "PutUpdateNotesSlideShapePortion", "storage", request.Storage, int32(statusCode), e)
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
    request.StorageName = createTestParamValue("StorageExists", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Check if storage exists
   Test for SlidesApi.StorageExists method with invalid storageName
*/
func TestStorageExistsInvalidStorageName(t *testing.T) {
    request := createStorageExistsRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "StorageExists", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("StorageExists", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.StorageExists(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StorageExists", "storageName", request.StorageName, int32(statusCode), e)
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
    request.Path = createTestParamValue("UploadFile", "path", "string").(string)
    request.File = createTestParamValue("UploadFile", "file", "[]byte").([]byte)
    request.StorageName = createTestParamValue("UploadFile", "storageName", "string").(string)
    return request
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid path
*/
func TestUploadFileInvalidPath(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.Path, "UploadFile", "path", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.Path = nullValue
    } else {
        request.Path = invalidValue.(string)
    }

    e := initializeTest("UploadFile", "path", request.Path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "path", request.Path, int32(statusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid file
*/
func TestUploadFileInvalidFile(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.File, "UploadFile", "file", "[]byte")
    if (invalidValue == nil) {
        request.File = nil
    } else {
        request.File = invalidValue.([]byte)
    }

    e := initializeTest("UploadFile", "file", request.File)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "file", request.File, int32(statusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid storageName
*/
func TestUploadFileInvalidStorageName(t *testing.T) {
    request := createUploadFileRequest()

    invalidValue := invalidizeTestParamValue(request.StorageName, "UploadFile", "storageName", "string")
    if (invalidValue == nil) {
        var nullValue string
        request.StorageName = nullValue
    } else {
        request.StorageName = invalidValue.(string)
    }

    e := initializeTest("UploadFile", "storageName", request.StorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    statusCode := 0
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "UploadFile", "storageName", request.StorageName, int32(statusCode), e)
}
