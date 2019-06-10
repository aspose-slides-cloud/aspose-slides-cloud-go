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
    request.srcPath = invalidizeTestParamValue(request.srcPath, "srcPath", "string").(string)

    e := initializeTest("CopyFile", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    assertError(t, "CopyFile", "srcPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destPath
*/
func TestCopyFileInvaliddestPath(t *testing.T) {
    request := createCopyFileRequest()
    request.destPath = invalidizeTestParamValue(request.destPath, "destPath", "string").(string)

    e := initializeTest("CopyFile", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    assertError(t, "CopyFile", "destPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid srcStorageName
*/
func TestCopyFileInvalidsrcStorageName(t *testing.T) {
    request := createCopyFileRequest()
    request.srcStorageName = invalidizeTestParamValue(request.srcStorageName, "srcStorageName", "string").(string)

    e := initializeTest("CopyFile", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    assertError(t, "CopyFile", "srcStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid destStorageName
*/
func TestCopyFileInvaliddestStorageName(t *testing.T) {
    request := createCopyFileRequest()
    request.destStorageName = invalidizeTestParamValue(request.destStorageName, "destStorageName", "string").(string)

    e := initializeTest("CopyFile", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    assertError(t, "CopyFile", "destStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy file
   Test for SlidesApi.CopyFile method with invalid versionId
*/
func TestCopyFileInvalidversionId(t *testing.T) {
    request := createCopyFileRequest()
    request.versionId = invalidizeTestParamValue(request.versionId, "versionId", "string").(string)

    e := initializeTest("CopyFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFile(request)
    assertError(t, "CopyFile", "versionId", int32(r.StatusCode), e)
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
    request.srcPath = invalidizeTestParamValue(request.srcPath, "srcPath", "string").(string)

    e := initializeTest("CopyFolder", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    assertError(t, "CopyFolder", "srcPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destPath
*/
func TestCopyFolderInvaliddestPath(t *testing.T) {
    request := createCopyFolderRequest()
    request.destPath = invalidizeTestParamValue(request.destPath, "destPath", "string").(string)

    e := initializeTest("CopyFolder", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    assertError(t, "CopyFolder", "destPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid srcStorageName
*/
func TestCopyFolderInvalidsrcStorageName(t *testing.T) {
    request := createCopyFolderRequest()
    request.srcStorageName = invalidizeTestParamValue(request.srcStorageName, "srcStorageName", "string").(string)

    e := initializeTest("CopyFolder", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    assertError(t, "CopyFolder", "srcStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy folder
   Test for SlidesApi.CopyFolder method with invalid destStorageName
*/
func TestCopyFolderInvaliddestStorageName(t *testing.T) {
    request := createCopyFolderRequest()
    request.destStorageName = invalidizeTestParamValue(request.destStorageName, "destStorageName", "string").(string)

    e := initializeTest("CopyFolder", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CopyFolder(request)
    assertError(t, "CopyFolder", "destStorageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("CreateFolder", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    assertError(t, "CreateFolder", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create the folder
   Test for SlidesApi.CreateFolder method with invalid storageName
*/
func TestCreateFolderInvalidstorageName(t *testing.T) {
    request := createCreateFolderRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("CreateFolder", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.CreateFolder(request)
    assertError(t, "CreateFolder", "storageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    assertError(t, "DeleteFile", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid storageName
*/
func TestDeleteFileInvalidstorageName(t *testing.T) {
    request := createDeleteFileRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("DeleteFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    assertError(t, "DeleteFile", "storageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete file
   Test for SlidesApi.DeleteFile method with invalid versionId
*/
func TestDeleteFileInvalidversionId(t *testing.T) {
    request := createDeleteFileRequest()
    request.versionId = invalidizeTestParamValue(request.versionId, "versionId", "string").(string)

    e := initializeTest("DeleteFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFile(request)
    assertError(t, "DeleteFile", "versionId", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteFolder", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    assertError(t, "DeleteFolder", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid storageName
*/
func TestDeleteFolderInvalidstorageName(t *testing.T) {
    request := createDeleteFolderRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("DeleteFolder", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    assertError(t, "DeleteFolder", "storageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete folder
   Test for SlidesApi.DeleteFolder method with invalid recursive
*/
func TestDeleteFolderInvalidrecursive(t *testing.T) {
    request := createDeleteFolderRequest()
    request.recursive = new(bool)
    *request.recursive = invalidizeTestParamValue(request.recursive, "recursive", "bool").(bool)

    e := initializeTest("DeleteFolder", "recursive", request.recursive)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.DeleteFolder(request)
    assertError(t, "DeleteFolder", "recursive", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid slideIndex
*/
func TestDeleteNotesSlideInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid password
*/
func TestDeleteNotesSlideInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid folder
*/
func TestDeleteNotesSlideInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove notes slide.
   Test for SlidesApi.DeleteNotesSlide method with invalid storage
*/
func TestDeleteNotesSlideInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlide(request)
    assertError(t, "DeleteNotesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid path
*/
func TestDeleteNotesSlideParagraphInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid paragraphIndex
*/
func TestDeleteNotesSlideParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid password
*/
func TestDeleteNotesSlideParagraphInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid folder
*/
func TestDeleteNotesSlideParagraphInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteNotesSlideParagraph method with invalid storage
*/
func TestDeleteNotesSlideParagraphInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraph(request)
    assertError(t, "DeleteNotesSlideParagraph", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid slideIndex
*/
func TestDeleteNotesSlideParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid path
*/
func TestDeleteNotesSlideParagraphsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid shapeIndex
*/
func TestDeleteNotesSlideParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid paragraphs
*/
func TestDeleteNotesSlideParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.paragraphs = invalidizeTestParamValue(request.paragraphs, "paragraphs", "[]int32").([]int32)

    e := initializeTest("DeleteNotesSlideParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "paragraphs", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid password
*/
func TestDeleteNotesSlideParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid folder
*/
func TestDeleteNotesSlideParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteNotesSlideParagraphs method with invalid storage
*/
func TestDeleteNotesSlideParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlideParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideParagraphs(request)
    assertError(t, "DeleteNotesSlideParagraphs", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlidePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid path
*/
func TestDeleteNotesSlidePortionInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlidePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid portionIndex
*/
func TestDeleteNotesSlidePortionInvalidportionIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid password
*/
func TestDeleteNotesSlidePortionInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlidePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid folder
*/
func TestDeleteNotesSlidePortionInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlidePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeleteNotesSlidePortion method with invalid storage
*/
func TestDeleteNotesSlidePortionInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlidePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortion(request)
    assertError(t, "DeleteNotesSlidePortion", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlidePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid slideIndex
*/
func TestDeleteNotesSlidePortionsInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid path
*/
func TestDeleteNotesSlidePortionsInvalidpath(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlidePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid shapeIndex
*/
func TestDeleteNotesSlidePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid paragraphIndex
*/
func TestDeleteNotesSlidePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlidePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid portions
*/
func TestDeleteNotesSlidePortionsInvalidportions(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.portions = invalidizeTestParamValue(request.portions, "portions", "[]int32").([]int32)

    e := initializeTest("DeleteNotesSlidePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "portions", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid password
*/
func TestDeleteNotesSlidePortionsInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlidePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid folder
*/
func TestDeleteNotesSlidePortionsInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlidePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeleteNotesSlidePortions method with invalid storage
*/
func TestDeleteNotesSlidePortionsInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlidePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlidePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlidePortions(request)
    assertError(t, "DeleteNotesSlidePortions", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid slideIndex
*/
func TestDeleteNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid path
*/
func TestDeleteNotesSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid shapeIndex
*/
func TestDeleteNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid password
*/
func TestDeleteNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid folder
*/
func TestDeleteNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteNotesSlideShape method with invalid storage
*/
func TestDeleteNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShape(request)
    assertError(t, "DeleteNotesSlideShape", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid slideIndex
*/
func TestDeleteNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid path
*/
func TestDeleteNotesSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid shapes
*/
func TestDeleteNotesSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.shapes = invalidizeTestParamValue(request.shapes, "shapes", "[]int32").([]int32)

    e := initializeTest("DeleteNotesSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "shapes", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid password
*/
func TestDeleteNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid folder
*/
func TestDeleteNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteNotesSlideShapes method with invalid storage
*/
func TestDeleteNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteNotesSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteNotesSlideShapes(request)
    assertError(t, "DeleteNotesSlideShapes", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid slideIndex
*/
func TestDeleteParagraphInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid path
*/
func TestDeleteParagraphInvalidpath(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid shapeIndex
*/
func TestDeleteParagraphInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid paragraphIndex
*/
func TestDeleteParagraphInvalidparagraphIndex(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeleteParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid password
*/
func TestDeleteParagraphInvalidpassword(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid folder
*/
func TestDeleteParagraphInvalidfolder(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a paragraph.
   Test for SlidesApi.DeleteParagraph method with invalid storage
*/
func TestDeleteParagraphInvalidstorage(t *testing.T) {
    request := createDeleteParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraph(request)
    assertError(t, "DeleteParagraph", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid slideIndex
*/
func TestDeleteParagraphsInvalidslideIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid path
*/
func TestDeleteParagraphsInvalidpath(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid shapeIndex
*/
func TestDeleteParagraphsInvalidshapeIndex(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid paragraphs
*/
func TestDeleteParagraphsInvalidparagraphs(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.paragraphs = invalidizeTestParamValue(request.paragraphs, "paragraphs", "[]int32").([]int32)

    e := initializeTest("DeleteParagraphs", "paragraphs", request.paragraphs)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "paragraphs", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid password
*/
func TestDeleteParagraphsInvalidpassword(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid folder
*/
func TestDeleteParagraphsInvalidfolder(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of paragraphs.
   Test for SlidesApi.DeleteParagraphs method with invalid storage
*/
func TestDeleteParagraphsInvalidstorage(t *testing.T) {
    request := createDeleteParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteParagraphs(request)
    assertError(t, "DeleteParagraphs", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeletePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid slideIndex
*/
func TestDeletePortionInvalidslideIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeletePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid path
*/
func TestDeletePortionInvalidpath(t *testing.T) {
    request := createDeletePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeletePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid shapeIndex
*/
func TestDeletePortionInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeletePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid paragraphIndex
*/
func TestDeletePortionInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeletePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid portionIndex
*/
func TestDeletePortionInvalidportionIndex(t *testing.T) {
    request := createDeletePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("DeletePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid password
*/
func TestDeletePortionInvalidpassword(t *testing.T) {
    request := createDeletePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeletePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid folder
*/
func TestDeletePortionInvalidfolder(t *testing.T) {
    request := createDeletePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeletePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a portion.
   Test for SlidesApi.DeletePortion method with invalid storage
*/
func TestDeletePortionInvalidstorage(t *testing.T) {
    request := createDeletePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeletePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortion(request)
    assertError(t, "DeletePortion", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeletePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid slideIndex
*/
func TestDeletePortionsInvalidslideIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeletePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid path
*/
func TestDeletePortionsInvalidpath(t *testing.T) {
    request := createDeletePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeletePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid shapeIndex
*/
func TestDeletePortionsInvalidshapeIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeletePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid paragraphIndex
*/
func TestDeletePortionsInvalidparagraphIndex(t *testing.T) {
    request := createDeletePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("DeletePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid portions
*/
func TestDeletePortionsInvalidportions(t *testing.T) {
    request := createDeletePortionsRequest()
    request.portions = invalidizeTestParamValue(request.portions, "portions", "[]int32").([]int32)

    e := initializeTest("DeletePortions", "portions", request.portions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "portions", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid password
*/
func TestDeletePortionsInvalidpassword(t *testing.T) {
    request := createDeletePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeletePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid folder
*/
func TestDeletePortionsInvalidfolder(t *testing.T) {
    request := createDeletePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeletePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of portions.
   Test for SlidesApi.DeletePortions method with invalid storage
*/
func TestDeletePortionsInvalidstorage(t *testing.T) {
    request := createDeletePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeletePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeletePortions(request)
    assertError(t, "DeletePortions", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlideByIndex", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid slideIndex
*/
func TestDeleteSlideByIndexInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteSlideByIndex", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid password
*/
func TestDeleteSlideByIndexInvalidpassword(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlideByIndex", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid folder
*/
func TestDeleteSlideByIndexInvalidfolder(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlideByIndex", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete a presentation slide by index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid storage
*/
func TestDeleteSlideByIndexInvalidstorage(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlideByIndex", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid slideIndex
*/
func TestDeleteSlideShapeInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid path
*/
func TestDeleteSlideShapeInvalidpath(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid shapeIndex
*/
func TestDeleteSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("DeleteSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid password
*/
func TestDeleteSlideShapeInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid folder
*/
func TestDeleteSlideShapeInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a shape.
   Test for SlidesApi.DeleteSlideShape method with invalid storage
*/
func TestDeleteSlideShapeInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShape(request)
    assertError(t, "DeleteSlideShape", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid slideIndex
*/
func TestDeleteSlideShapesInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid path
*/
func TestDeleteSlideShapesInvalidpath(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DeleteSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid shapes
*/
func TestDeleteSlideShapesInvalidshapes(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.shapes = invalidizeTestParamValue(request.shapes, "shapes", "[]int32").([]int32)

    e := initializeTest("DeleteSlideShapes", "shapes", request.shapes)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "shapes", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid password
*/
func TestDeleteSlideShapesInvalidpassword(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid folder
*/
func TestDeleteSlideShapesInvalidfolder(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove a range of shapes.
   Test for SlidesApi.DeleteSlideShapes method with invalid storage
*/
func TestDeleteSlideShapesInvalidstorage(t *testing.T) {
    request := createDeleteSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlideShapes(request)
    assertError(t, "DeleteSlideShapes", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesCleanSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid slides
*/
func TestDeleteSlidesCleanSlidesListInvalidslides(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.slides = invalidizeTestParamValue(request.slides, "slides", "[]int32").([]int32)

    e := initializeTest("DeleteSlidesCleanSlidesList", "slides", request.slides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "slides", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid password
*/
func TestDeleteSlidesCleanSlidesListInvalidpassword(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesCleanSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid folder
*/
func TestDeleteSlidesCleanSlidesListInvalidfolder(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesCleanSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid storage
*/
func TestDeleteSlidesCleanSlidesListInvalidstorage(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesCleanSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid password
*/
func TestDeleteSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid folder
*/
func TestDeleteSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Clean document properties.
   Test for SlidesApi.DeleteSlidesDocumentProperties method with invalid storage
*/
func TestDeleteSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid propertyName
*/
func TestDeleteSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "propertyName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid password
*/
func TestDeleteSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid folder
*/
func TestDeleteSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Delete document property.
   Test for SlidesApi.DeleteSlidesDocumentProperty method with invalid storage
*/
func TestDeleteSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid slideIndex
*/
func TestDeleteSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("DeleteSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid password
*/
func TestDeleteSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid folder
*/
func TestDeleteSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Remove background from a slide.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid storage
*/
func TestDeleteSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "storage", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("DownloadFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    assertError(t, "DownloadFile", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid storageName
*/
func TestDownloadFileInvalidstorageName(t *testing.T) {
    request := createDownloadFileRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("DownloadFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    assertError(t, "DownloadFile", "storageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Download file
   Test for SlidesApi.DownloadFile method with invalid versionId
*/
func TestDownloadFileInvalidversionId(t *testing.T) {
    request := createDownloadFileRequest()
    request.versionId = invalidizeTestParamValue(request.versionId, "versionId", "string").(string)

    e := initializeTest("DownloadFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.DownloadFile(request)
    assertError(t, "DownloadFile", "versionId", int32(r.StatusCode), e)
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
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("GetDiscUsage", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetDiscUsage(request)
    assertError(t, "GetDiscUsage", "storageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetFileVersions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    assertError(t, "GetFileVersions", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get file versions
   Test for SlidesApi.GetFileVersions method with invalid storageName
*/
func TestGetFileVersionsInvalidstorageName(t *testing.T) {
    request := createGetFileVersionsRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("GetFileVersions", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFileVersions(request)
    assertError(t, "GetFileVersions", "storageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetFilesList", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    assertError(t, "GetFilesList", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get all files and folders within a folder
   Test for SlidesApi.GetFilesList method with invalid storageName
*/
func TestGetFilesListInvalidstorageName(t *testing.T) {
    request := createGetFilesListRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("GetFilesList", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetFilesList(request)
    assertError(t, "GetFilesList", "storageName", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid slideIndex
*/
func TestGetLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid password
*/
func TestGetLayoutSlideInvalidpassword(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid folder
*/
func TestGetLayoutSlideInvalidfolder(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlide info.
   Test for SlidesApi.GetLayoutSlide method with invalid storage
*/
func TestGetLayoutSlideInvalidstorage(t *testing.T) {
    request := createGetLayoutSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlide(request)
    assertError(t, "GetLayoutSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetLayoutSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid password
*/
func TestGetLayoutSlidesListInvalidpassword(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetLayoutSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid folder
*/
func TestGetLayoutSlidesListInvalidfolder(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetLayoutSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation layoutSlides info.
   Test for SlidesApi.GetLayoutSlidesList method with invalid storage
*/
func TestGetLayoutSlidesListInvalidstorage(t *testing.T) {
    request := createGetLayoutSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetLayoutSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetLayoutSlidesList(request)
    assertError(t, "GetLayoutSlidesList", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetMasterSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid slideIndex
*/
func TestGetMasterSlideInvalidslideIndex(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetMasterSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid password
*/
func TestGetMasterSlideInvalidpassword(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetMasterSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid folder
*/
func TestGetMasterSlideInvalidfolder(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetMasterSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlide info.
   Test for SlidesApi.GetMasterSlide method with invalid storage
*/
func TestGetMasterSlideInvalidstorage(t *testing.T) {
    request := createGetMasterSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetMasterSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlide(request)
    assertError(t, "GetMasterSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetMasterSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid password
*/
func TestGetMasterSlidesListInvalidpassword(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetMasterSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid folder
*/
func TestGetMasterSlidesListInvalidfolder(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetMasterSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation masterSlides info.
   Test for SlidesApi.GetMasterSlidesList method with invalid storage
*/
func TestGetMasterSlidesListInvalidstorage(t *testing.T) {
    request := createGetMasterSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetMasterSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetMasterSlidesList(request)
    assertError(t, "GetMasterSlidesList", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid slideIndex
*/
func TestGetNotesSlideInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid password
*/
func TestGetNotesSlideInvalidpassword(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid folder
*/
func TestGetNotesSlideInvalidfolder(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read notes slide info.
   Test for SlidesApi.GetNotesSlide method with invalid storage
*/
func TestGetNotesSlideInvalidstorage(t *testing.T) {
    request := createGetNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlide(request)
    assertError(t, "GetNotesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid slideIndex
*/
func TestGetNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid path
*/
func TestGetNotesSlideShapeInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid shapeIndex
*/
func TestGetNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid password
*/
func TestGetNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid folder
*/
func TestGetNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetNotesSlideShape method with invalid storage
*/
func TestGetNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShape(request)
    assertError(t, "GetNotesSlideShape", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid path
*/
func TestGetNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid password
*/
func TestGetNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid folder
*/
func TestGetNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetNotesSlideShapeParagraph method with invalid storage
*/
func TestGetNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraph(request)
    assertError(t, "GetNotesSlideShapeParagraph", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid path
*/
func TestGetNotesSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetNotesSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid password
*/
func TestGetNotesSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid folder
*/
func TestGetNotesSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetNotesSlideShapeParagraphs method with invalid storage
*/
func TestGetNotesSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapeParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapeParagraphs(request)
    assertError(t, "GetNotesSlideShapeParagraphs", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid path
*/
func TestGetNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid portionIndex
*/
func TestGetNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid password
*/
func TestGetNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid folder
*/
func TestGetNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetNotesSlideShapePortion method with invalid storage
*/
func TestGetNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortion(request)
    assertError(t, "GetNotesSlideShapePortion", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShapePortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid slideIndex
*/
func TestGetNotesSlideShapePortionsInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid path
*/
func TestGetNotesSlideShapePortionsInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShapePortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid shapeIndex
*/
func TestGetNotesSlideShapePortionsInvalidshapeIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid paragraphIndex
*/
func TestGetNotesSlideShapePortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapePortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid password
*/
func TestGetNotesSlideShapePortionsInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShapePortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid folder
*/
func TestGetNotesSlideShapePortionsInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShapePortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetNotesSlideShapePortions method with invalid storage
*/
func TestGetNotesSlideShapePortionsInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapePortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShapePortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapePortions(request)
    assertError(t, "GetNotesSlideShapePortions", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid slideIndex
*/
func TestGetNotesSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid path
*/
func TestGetNotesSlideShapesInvalidpath(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetNotesSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid password
*/
func TestGetNotesSlideShapesInvalidpassword(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid folder
*/
func TestGetNotesSlideShapesInvalidfolder(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetNotesSlideShapes method with invalid storage
*/
func TestGetNotesSlideShapesInvalidstorage(t *testing.T) {
    request := createGetNotesSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideShapes(request)
    assertError(t, "GetNotesSlideShapes", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("GetNotesSlideWithFormat", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid slideIndex
*/
func TestGetNotesSlideWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetNotesSlideWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid format
*/
func TestGetNotesSlideWithFormatInvalidformat(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("GetNotesSlideWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid width
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
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid height
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
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid password
*/
func TestGetNotesSlideWithFormatInvalidpassword(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid folder
*/
func TestGetNotesSlideWithFormatInvalidfolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid storage
*/
func TestGetNotesSlideWithFormatInvalidstorage(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert notes slide to the specified image format.
   Test for SlidesApi.GetNotesSlideWithFormat method with invalid fontsFolder
*/
func TestGetNotesSlideWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetNotesSlideWithFormatRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("GetNotesSlideWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetNotesSlideWithFormat(request)
    assertError(t, "GetNotesSlideWithFormat", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetParagraphPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid slideIndex
*/
func TestGetParagraphPortionInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid path
*/
func TestGetParagraphPortionInvalidpath(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetParagraphPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid shapeIndex
*/
func TestGetParagraphPortionInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid paragraphIndex
*/
func TestGetParagraphPortionInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid portionIndex
*/
func TestGetParagraphPortionInvalidportionIndex(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid password
*/
func TestGetParagraphPortionInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetParagraphPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid folder
*/
func TestGetParagraphPortionInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetParagraphPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portion info.
   Test for SlidesApi.GetParagraphPortion method with invalid storage
*/
func TestGetParagraphPortionInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetParagraphPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortion(request)
    assertError(t, "GetParagraphPortion", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetParagraphPortions", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid slideIndex
*/
func TestGetParagraphPortionsInvalidslideIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortions", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid path
*/
func TestGetParagraphPortionsInvalidpath(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetParagraphPortions", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid shapeIndex
*/
func TestGetParagraphPortionsInvalidshapeIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortions", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid paragraphIndex
*/
func TestGetParagraphPortionsInvalidparagraphIndex(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetParagraphPortions", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid password
*/
func TestGetParagraphPortionsInvalidpassword(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetParagraphPortions", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid folder
*/
func TestGetParagraphPortionsInvalidfolder(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetParagraphPortions", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read paragraph portions info.
   Test for SlidesApi.GetParagraphPortions method with invalid storage
*/
func TestGetParagraphPortionsInvalidstorage(t *testing.T) {
    request := createGetParagraphPortionsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetParagraphPortions", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetParagraphPortions(request)
    assertError(t, "GetParagraphPortions", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid slideIndex
*/
func TestGetSlideShapeInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid path
*/
func TestGetSlideShapeInvalidpath(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid shapeIndex
*/
func TestGetSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid password
*/
func TestGetSlideShapeInvalidpassword(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid folder
*/
func TestGetSlideShapeInvalidfolder(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shape info.
   Test for SlidesApi.GetSlideShape method with invalid storage
*/
func TestGetSlideShapeInvalidstorage(t *testing.T) {
    request := createGetSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShape(request)
    assertError(t, "GetSlideShape", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid slideIndex
*/
func TestGetSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid path
*/
func TestGetSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid paragraphIndex
*/
func TestGetSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("GetSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid password
*/
func TestGetSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid folder
*/
func TestGetSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraph info.
   Test for SlidesApi.GetSlideShapeParagraph method with invalid storage
*/
func TestGetSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createGetSlideShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraph(request)
    assertError(t, "GetSlideShapeParagraph", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlideShapeParagraphs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid slideIndex
*/
func TestGetSlideShapeParagraphsInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlideShapeParagraphs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid path
*/
func TestGetSlideShapeParagraphsInvalidpath(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetSlideShapeParagraphs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid shapeIndex
*/
func TestGetSlideShapeParagraphsInvalidshapeIndex(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("GetSlideShapeParagraphs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid password
*/
func TestGetSlideShapeParagraphsInvalidpassword(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlideShapeParagraphs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid folder
*/
func TestGetSlideShapeParagraphsInvalidfolder(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlideShapeParagraphs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read shape paragraphs info.
   Test for SlidesApi.GetSlideShapeParagraphs method with invalid storage
*/
func TestGetSlideShapeParagraphsInvalidstorage(t *testing.T) {
    request := createGetSlideShapeParagraphsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlideShapeParagraphs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapeParagraphs(request)
    assertError(t, "GetSlideShapeParagraphs", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlideShapes", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid slideIndex
*/
func TestGetSlideShapesInvalidslideIndex(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlideShapes", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid path
*/
func TestGetSlideShapesInvalidpath(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("GetSlideShapes", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid password
*/
func TestGetSlideShapesInvalidpassword(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlideShapes", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid folder
*/
func TestGetSlideShapesInvalidfolder(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlideShapes", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide shapes info.
   Test for SlidesApi.GetSlideShapes method with invalid storage
*/
func TestGetSlideShapesInvalidstorage(t *testing.T) {
    request := createGetSlideShapesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlideShapes", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideShapes(request)
    assertError(t, "GetSlideShapes", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid password
*/
func TestGetSlidesDocumentInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid storage
*/
func TestGetSlidesDocumentInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation info.
   Test for SlidesApi.GetSlidesDocument method with invalid folder
*/
func TestGetSlidesDocumentInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocument(request)
    assertError(t, "GetSlidesDocument", "folder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid password
*/
func TestGetSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid folder
*/
func TestGetSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document properties.
   Test for SlidesApi.GetSlidesDocumentProperties method with invalid storage
*/
func TestGetSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid propertyName
*/
func TestGetSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "propertyName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid password
*/
func TestGetSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid folder
*/
func TestGetSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation document property.
   Test for SlidesApi.GetSlidesDocumentProperty method with invalid storage
*/
func TestGetSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesImageWithDefaultFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    assertError(t, "GetSlidesImageWithDefaultFormat", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid index
*/
func TestGetSlidesImageWithDefaultFormatInvalidindex(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()
    request.index = invalidizeTestParamValue(request.index, "index", "int32").(int32)

    e := initializeTest("GetSlidesImageWithDefaultFormat", "index", request.index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    assertError(t, "GetSlidesImageWithDefaultFormat", "index", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid password
*/
func TestGetSlidesImageWithDefaultFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesImageWithDefaultFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    assertError(t, "GetSlidesImageWithDefaultFormat", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid folder
*/
func TestGetSlidesImageWithDefaultFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesImageWithDefaultFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    assertError(t, "GetSlidesImageWithDefaultFormat", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image binary data.
   Test for SlidesApi.GetSlidesImageWithDefaultFormat method with invalid storage
*/
func TestGetSlidesImageWithDefaultFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesImageWithDefaultFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesImageWithDefaultFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithDefaultFormat(request)
    assertError(t, "GetSlidesImageWithDefaultFormat", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("GetSlidesImageWithFormat", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesImageWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid index
*/
func TestGetSlidesImageWithFormatInvalidindex(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.index = invalidizeTestParamValue(request.index, "index", "int32").(int32)

    e := initializeTest("GetSlidesImageWithFormat", "index", request.index)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "index", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid format
*/
func TestGetSlidesImageWithFormatInvalidformat(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("GetSlidesImageWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid password
*/
func TestGetSlidesImageWithFormatInvalidpassword(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesImageWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid folder
*/
func TestGetSlidesImageWithFormatInvalidfolder(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesImageWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Get image in specified format.
   Test for SlidesApi.GetSlidesImageWithFormat method with invalid storage
*/
func TestGetSlidesImageWithFormatInvalidstorage(t *testing.T) {
    request := createGetSlidesImageWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesImageWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImageWithFormat(request)
    assertError(t, "GetSlidesImageWithFormat", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid password
*/
func TestGetSlidesImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid folder
*/
func TestGetSlidesImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation images info.
   Test for SlidesApi.GetSlidesImages method with invalid storage
*/
func TestGetSlidesImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesImagesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesImages(request)
    assertError(t, "GetSlidesImages", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesPlaceholder", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid slideIndex
*/
func TestGetSlidesPlaceholderInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesPlaceholder", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid placeholderIndex
*/
func TestGetSlidesPlaceholderInvalidplaceholderIndex(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.placeholderIndex = invalidizeTestParamValue(request.placeholderIndex, "placeholderIndex", "int32").(int32)

    e := initializeTest("GetSlidesPlaceholder", "placeholderIndex", request.placeholderIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "placeholderIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid password
*/
func TestGetSlidesPlaceholderInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesPlaceholder", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid folder
*/
func TestGetSlidesPlaceholderInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesPlaceholder", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholder info.
   Test for SlidesApi.GetSlidesPlaceholder method with invalid storage
*/
func TestGetSlidesPlaceholderInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholderRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesPlaceholder", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholder(request)
    assertError(t, "GetSlidesPlaceholder", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesPlaceholders", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid slideIndex
*/
func TestGetSlidesPlaceholdersInvalidslideIndex(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesPlaceholders", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid password
*/
func TestGetSlidesPlaceholdersInvalidpassword(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesPlaceholders", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid folder
*/
func TestGetSlidesPlaceholdersInvalidfolder(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesPlaceholders", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide placeholders info.
   Test for SlidesApi.GetSlidesPlaceholders method with invalid storage
*/
func TestGetSlidesPlaceholdersInvalidstorage(t *testing.T) {
    request := createGetSlidesPlaceholdersRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesPlaceholders", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPlaceholders(request)
    assertError(t, "GetSlidesPlaceholders", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesPresentationTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid withEmpty
*/
func TestGetSlidesPresentationTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.withEmpty = new(bool)
    *request.withEmpty = invalidizeTestParamValue(request.withEmpty, "withEmpty", "bool").(bool)

    e := initializeTest("GetSlidesPresentationTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "withEmpty", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid password
*/
func TestGetSlidesPresentationTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesPresentationTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid folder
*/
func TestGetSlidesPresentationTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesPresentationTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract presentation text items.
   Test for SlidesApi.GetSlidesPresentationTextItems method with invalid storage
*/
func TestGetSlidesPresentationTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesPresentationTextItemsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesPresentationTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesPresentationTextItems(request)
    assertError(t, "GetSlidesPresentationTextItems", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid slideIndex
*/
func TestGetSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid password
*/
func TestGetSlidesSlideInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid folder
*/
func TestGetSlidesSlideInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid storage
*/
func TestGetSlidesSlideInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid slideIndex
*/
func TestGetSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid password
*/
func TestGetSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid folder
*/
func TestGetSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide background info.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid storage
*/
func TestGetSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlideComments", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid slideIndex
*/
func TestGetSlidesSlideCommentsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesSlideComments", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid password
*/
func TestGetSlidesSlideCommentsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlideComments", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid folder
*/
func TestGetSlidesSlideCommentsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlideComments", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid storage
*/
func TestGetSlidesSlideCommentsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlideComments", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlideImages", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid slideIndex
*/
func TestGetSlidesSlideImagesInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesSlideImages", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid password
*/
func TestGetSlidesSlideImagesInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlideImages", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid folder
*/
func TestGetSlidesSlideImagesInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlideImages", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide images info.
   Test for SlidesApi.GetSlidesSlideImages method with invalid storage
*/
func TestGetSlidesSlideImagesInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideImagesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlideImages", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideImages(request)
    assertError(t, "GetSlidesSlideImages", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlideTextItems", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid slideIndex
*/
func TestGetSlidesSlideTextItemsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesSlideTextItems", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid withEmpty
*/
func TestGetSlidesSlideTextItemsInvalidwithEmpty(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.withEmpty = new(bool)
    *request.withEmpty = invalidizeTestParamValue(request.withEmpty, "withEmpty", "bool").(bool)

    e := initializeTest("GetSlidesSlideTextItems", "withEmpty", request.withEmpty)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "withEmpty", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid password
*/
func TestGetSlidesSlideTextItemsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlideTextItems", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid folder
*/
func TestGetSlidesSlideTextItemsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlideTextItems", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Extract slide text items.
   Test for SlidesApi.GetSlidesSlideTextItems method with invalid storage
*/
func TestGetSlidesSlideTextItemsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideTextItemsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlideTextItems", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlideTextItems(request)
    assertError(t, "GetSlidesSlideTextItems", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid password
*/
func TestGetSlidesSlidesListInvalidpassword(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid folder
*/
func TestGetSlidesSlidesListInvalidfolder(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid storage
*/
func TestGetSlidesSlidesListInvalidstorage(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesTheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid slideIndex
*/
func TestGetSlidesThemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesTheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid password
*/
func TestGetSlidesThemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesTheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid folder
*/
func TestGetSlidesThemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesTheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme info.
   Test for SlidesApi.GetSlidesTheme method with invalid storage
*/
func TestGetSlidesThemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesTheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesThemeColorScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid slideIndex
*/
func TestGetSlidesThemeColorSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesThemeColorScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid password
*/
func TestGetSlidesThemeColorSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesThemeColorScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid folder
*/
func TestGetSlidesThemeColorSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesThemeColorScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme color scheme info.
   Test for SlidesApi.GetSlidesThemeColorScheme method with invalid storage
*/
func TestGetSlidesThemeColorSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesThemeColorScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesThemeFontScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFontSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesThemeFontScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid password
*/
func TestGetSlidesThemeFontSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesThemeFontScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid folder
*/
func TestGetSlidesThemeFontSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesThemeFontScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme font scheme info.
   Test for SlidesApi.GetSlidesThemeFontScheme method with invalid storage
*/
func TestGetSlidesThemeFontSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesThemeFontScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesThemeFormatScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFormatSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("GetSlidesThemeFormatScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid password
*/
func TestGetSlidesThemeFormatSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesThemeFormatScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid folder
*/
func TestGetSlidesThemeFormatSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesThemeFormatScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide theme format scheme info.
   Test for SlidesApi.GetSlidesThemeFormatScheme method with invalid storage
*/
func TestGetSlidesThemeFormatSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesThemeFormatScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "storage", int32(r.StatusCode), e)
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
    request.srcPath = invalidizeTestParamValue(request.srcPath, "srcPath", "string").(string)

    e := initializeTest("MoveFile", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    assertError(t, "MoveFile", "srcPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destPath
*/
func TestMoveFileInvaliddestPath(t *testing.T) {
    request := createMoveFileRequest()
    request.destPath = invalidizeTestParamValue(request.destPath, "destPath", "string").(string)

    e := initializeTest("MoveFile", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    assertError(t, "MoveFile", "destPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid srcStorageName
*/
func TestMoveFileInvalidsrcStorageName(t *testing.T) {
    request := createMoveFileRequest()
    request.srcStorageName = invalidizeTestParamValue(request.srcStorageName, "srcStorageName", "string").(string)

    e := initializeTest("MoveFile", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    assertError(t, "MoveFile", "srcStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid destStorageName
*/
func TestMoveFileInvaliddestStorageName(t *testing.T) {
    request := createMoveFileRequest()
    request.destStorageName = invalidizeTestParamValue(request.destStorageName, "destStorageName", "string").(string)

    e := initializeTest("MoveFile", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    assertError(t, "MoveFile", "destStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move file
   Test for SlidesApi.MoveFile method with invalid versionId
*/
func TestMoveFileInvalidversionId(t *testing.T) {
    request := createMoveFileRequest()
    request.versionId = invalidizeTestParamValue(request.versionId, "versionId", "string").(string)

    e := initializeTest("MoveFile", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFile(request)
    assertError(t, "MoveFile", "versionId", int32(r.StatusCode), e)
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
    request.srcPath = invalidizeTestParamValue(request.srcPath, "srcPath", "string").(string)

    e := initializeTest("MoveFolder", "srcPath", request.srcPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    assertError(t, "MoveFolder", "srcPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destPath
*/
func TestMoveFolderInvaliddestPath(t *testing.T) {
    request := createMoveFolderRequest()
    request.destPath = invalidizeTestParamValue(request.destPath, "destPath", "string").(string)

    e := initializeTest("MoveFolder", "destPath", request.destPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    assertError(t, "MoveFolder", "destPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid srcStorageName
*/
func TestMoveFolderInvalidsrcStorageName(t *testing.T) {
    request := createMoveFolderRequest()
    request.srcStorageName = invalidizeTestParamValue(request.srcStorageName, "srcStorageName", "string").(string)

    e := initializeTest("MoveFolder", "srcStorageName", request.srcStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    assertError(t, "MoveFolder", "srcStorageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Move folder
   Test for SlidesApi.MoveFolder method with invalid destStorageName
*/
func TestMoveFolderInvaliddestStorageName(t *testing.T) {
    request := createMoveFolderRequest()
    request.destStorageName = invalidizeTestParamValue(request.destStorageName, "destStorageName", "string").(string)

    e := initializeTest("MoveFolder", "destStorageName", request.destStorageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.MoveFolder(request)
    assertError(t, "MoveFolder", "destStorageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("ObjectExists", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    assertError(t, "ObjectExists", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid storageName
*/
func TestObjectExistsInvalidstorageName(t *testing.T) {
    request := createObjectExistsRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("ObjectExists", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    assertError(t, "ObjectExists", "storageName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Check if file or folder exists
   Test for SlidesApi.ObjectExists method with invalid versionId
*/
func TestObjectExistsInvalidversionId(t *testing.T) {
    request := createObjectExistsRequest()
    request.versionId = invalidizeTestParamValue(request.versionId, "versionId", "string").(string)

    e := initializeTest("ObjectExists", "versionId", request.versionId)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.ObjectExists(request)
    assertError(t, "ObjectExists", "versionId", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid slideIndex
*/
func TestPostAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid path
*/
func TestPostAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid shapeIndex
*/
func TestPostAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid dto
*/
func TestPostAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)

    e := initializeTest("PostAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid password
*/
func TestPostAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid folder
*/
func TestPostAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid storage
*/
func TestPostAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostAddNewParagraph method with invalid position
*/
func TestPostAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostAddNewParagraphRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewParagraph(request)
    assertError(t, "PostAddNewParagraph", "position", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid slideIndex
*/
func TestPostAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid path
*/
func TestPostAddNewPortionInvalidpath(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid shapeIndex
*/
func TestPostAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid paragraphIndex
*/
func TestPostAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PostAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid dto
*/
func TestPostAddNewPortionInvaliddto(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)

    e := initializeTest("PostAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid password
*/
func TestPostAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid folder
*/
func TestPostAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid storage
*/
func TestPostAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostAddNewPortion method with invalid position
*/
func TestPostAddNewPortionInvalidposition(t *testing.T) {
    request := createPostAddNewPortionRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewPortion(request)
    assertError(t, "PostAddNewPortion", "position", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid slideIndex
*/
func TestPostAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid path
*/
func TestPostAddNewShapeInvalidpath(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid dto
*/
func TestPostAddNewShapeInvaliddto(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)

    e := initializeTest("PostAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid password
*/
func TestPostAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid folder
*/
func TestPostAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid storage
*/
func TestPostAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid shapeToClone
*/
func TestPostAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.shapeToClone = new(int32)
    *request.shapeToClone = invalidizeTestParamValue(request.shapeToClone, "shapeToClone", "int32").(int32)

    e := initializeTest("PostAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "shapeToClone", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostAddNewShape method with invalid position
*/
func TestPostAddNewShapeInvalidposition(t *testing.T) {
    request := createPostAddNewShapeRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNewShape(request)
    assertError(t, "PostAddNewShape", "position", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostAddNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid slideIndex
*/
func TestPostAddNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostAddNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid dto
*/
func TestPostAddNotesSlideInvaliddto(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "NotesSlide").(INotesSlide)

    e := initializeTest("PostAddNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid password
*/
func TestPostAddNotesSlideInvalidpassword(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostAddNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid folder
*/
func TestPostAddNotesSlideInvalidfolder(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostAddNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Add new notes slide.
   Test for SlidesApi.PostAddNotesSlide method with invalid storage
*/
func TestPostAddNotesSlideInvalidstorage(t *testing.T) {
    request := createPostAddNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostAddNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostAddNotesSlide(request)
    assertError(t, "PostAddNotesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFrom = invalidizeTestParamValue(request.cloneFrom, "cloneFrom", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFrom", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromPosition = invalidizeTestParamValue(request.cloneFromPosition, "cloneFromPosition", "int32").(int32)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPosition", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromPassword = invalidizeTestParamValue(request.cloneFromPassword, "cloneFromPassword", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromPassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.cloneFromStorage = invalidizeTestParamValue(request.cloneFromStorage, "cloneFromStorage", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "cloneFromStorage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy layoutSlide from source presentation.
   Test for SlidesApi.PostCopyLayoutSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyLayoutSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyLayoutSlideFromSourcePresentationRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostCopyLayoutSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyLayoutSlideFromSourcePresentation(request)
    assertError(t, "PostCopyLayoutSlideFromSourcePresentation", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFrom
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFrom(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFrom = invalidizeTestParamValue(request.cloneFrom, "cloneFrom", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFrom", request.cloneFrom)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFrom", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPosition
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPosition(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromPosition = invalidizeTestParamValue(request.cloneFromPosition, "cloneFromPosition", "int32").(int32)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", request.cloneFromPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPosition", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromPassword
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromPassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromPassword = invalidizeTestParamValue(request.cloneFromPassword, "cloneFromPassword", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", request.cloneFromPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromPassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid cloneFromStorage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidcloneFromStorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.cloneFromStorage = invalidizeTestParamValue(request.cloneFromStorage, "cloneFromStorage", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", request.cloneFromStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "cloneFromStorage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid applyToAll
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidapplyToAll(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.applyToAll = new(bool)
    *request.applyToAll = invalidizeTestParamValue(request.applyToAll, "applyToAll", "bool").(bool)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "applyToAll", request.applyToAll)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "applyToAll", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid password
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidpassword(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid folder
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidfolder(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy masterSlide from source presentation.
   Test for SlidesApi.PostCopyMasterSlideFromSourcePresentation method with invalid storage
*/
func TestPostCopyMasterSlideFromSourcePresentationInvalidstorage(t *testing.T) {
    request := createPostCopyMasterSlideFromSourcePresentationRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostCopyMasterSlideFromSourcePresentation", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostCopyMasterSlideFromSourcePresentation(request)
    assertError(t, "PostCopyMasterSlideFromSourcePresentation", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostNotesSlideAddNewParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid slideIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid path
*/
func TestPostNotesSlideAddNewParagraphInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostNotesSlideAddNewParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewParagraphInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid dto
*/
func TestPostNotesSlideAddNewParagraphInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)

    e := initializeTest("PostNotesSlideAddNewParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid password
*/
func TestPostNotesSlideAddNewParagraphInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostNotesSlideAddNewParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid folder
*/
func TestPostNotesSlideAddNewParagraphInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostNotesSlideAddNewParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid storage
*/
func TestPostNotesSlideAddNewParagraphInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostNotesSlideAddNewParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new paragraph.
   Test for SlidesApi.PostNotesSlideAddNewParagraph method with invalid position
*/
func TestPostNotesSlideAddNewParagraphInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewParagraphRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewParagraph", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewParagraph(request)
    assertError(t, "PostNotesSlideAddNewParagraph", "position", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostNotesSlideAddNewPortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid slideIndex
*/
func TestPostNotesSlideAddNewPortionInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewPortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid path
*/
func TestPostNotesSlideAddNewPortionInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostNotesSlideAddNewPortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid shapeIndex
*/
func TestPostNotesSlideAddNewPortionInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewPortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid paragraphIndex
*/
func TestPostNotesSlideAddNewPortionInvalidparagraphIndex(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewPortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid dto
*/
func TestPostNotesSlideAddNewPortionInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)

    e := initializeTest("PostNotesSlideAddNewPortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid password
*/
func TestPostNotesSlideAddNewPortionInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostNotesSlideAddNewPortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid folder
*/
func TestPostNotesSlideAddNewPortionInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostNotesSlideAddNewPortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid storage
*/
func TestPostNotesSlideAddNewPortionInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostNotesSlideAddNewPortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Creates new portion.
   Test for SlidesApi.PostNotesSlideAddNewPortion method with invalid position
*/
func TestPostNotesSlideAddNewPortionInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewPortionRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewPortion", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewPortion(request)
    assertError(t, "PostNotesSlideAddNewPortion", "position", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostNotesSlideAddNewShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid slideIndex
*/
func TestPostNotesSlideAddNewShapeInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid path
*/
func TestPostNotesSlideAddNewShapeInvalidpath(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostNotesSlideAddNewShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid dto
*/
func TestPostNotesSlideAddNewShapeInvaliddto(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)

    e := initializeTest("PostNotesSlideAddNewShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid password
*/
func TestPostNotesSlideAddNewShapeInvalidpassword(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostNotesSlideAddNewShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid folder
*/
func TestPostNotesSlideAddNewShapeInvalidfolder(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostNotesSlideAddNewShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid storage
*/
func TestPostNotesSlideAddNewShapeInvalidstorage(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostNotesSlideAddNewShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid shapeToClone
*/
func TestPostNotesSlideAddNewShapeInvalidshapeToClone(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.shapeToClone = new(int32)
    *request.shapeToClone = invalidizeTestParamValue(request.shapeToClone, "shapeToClone", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewShape", "shapeToClone", request.shapeToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "shapeToClone", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create new shape.
   Test for SlidesApi.PostNotesSlideAddNewShape method with invalid position
*/
func TestPostNotesSlideAddNewShapeInvalidposition(t *testing.T) {
    request := createPostNotesSlideAddNewShapeRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostNotesSlideAddNewShape", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideAddNewShape(request)
    assertError(t, "PostNotesSlideAddNewShape", "position", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PostNotesSlideShapeSaveAs", "format", "int32").(int32)
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
    testbounds := createTestParamValue("PostNotesSlideShapeSaveAs", "bounds", "int32")
    switch v := testbounds.(type) { 
    case int32:
        request.bounds = new(int32)
        *request.bounds = v
    }
    request.fontsFolder = createTestParamValue("PostNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid name
*/
func TestPostNotesSlideShapeSaveAsInvalidname(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid path
*/
func TestPostNotesSlideShapeSaveAsInvalidpath(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPostNotesSlideShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid format
*/
func TestPostNotesSlideShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostNotesSlideShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid options
*/
func TestPostNotesSlideShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)

    e := initializeTest("PostNotesSlideShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid password
*/
func TestPostNotesSlideShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid folder
*/
func TestPostNotesSlideShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid storage
*/
func TestPostNotesSlideShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleX = new(float64)
    *request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPostNotesSlideShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.scaleY = new(float64)
    *request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)

    e := initializeTest("PostNotesSlideShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPostNotesSlideShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.bounds = new(int32)
    *request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "int32").(int32)

    e := initializeTest("PostNotesSlideShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPostNotesSlideShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostNotesSlideShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostNotesSlideShapeSaveAs(request)
    assertError(t, "PostNotesSlideShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid request
*/
func TestPostPresentationMergeInvalidrequest(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.request = invalidizeTestParamValue(request.request, "request", "PresentationsMergeRequest").(IPresentationsMergeRequest)

    e := initializeTest("PostPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "request", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid password
*/
func TestPostPresentationMergeInvalidpassword(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid storage
*/
func TestPostPresentationMergeInvalidstorage(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations specified in the request parameter.
   Test for SlidesApi.PostPresentationMerge method with invalid folder
*/
func TestPostPresentationMergeInvalidfolder(t *testing.T) {
    request := createPostPresentationMergeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostPresentationMerge(request)
    assertError(t, "PostPresentationMerge", "folder", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PostShapeSaveAs", "format", "int32").(int32)
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
    testbounds := createTestParamValue("PostShapeSaveAs", "bounds", "int32")
    switch v := testbounds.(type) { 
    case int32:
        request.bounds = new(int32)
        *request.bounds = v
    }
    request.fontsFolder = createTestParamValue("PostShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid name
*/
func TestPostShapeSaveAsInvalidname(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid slideIndex
*/
func TestPostShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid path
*/
func TestPostShapeSaveAsInvalidpath(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PostShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid shapeIndex
*/
func TestPostShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PostShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid format
*/
func TestPostShapeSaveAsInvalidformat(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid options
*/
func TestPostShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)

    e := initializeTest("PostShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid password
*/
func TestPostShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid folder
*/
func TestPostShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid storage
*/
func TestPostShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleX
*/
func TestPostShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleX = new(float64)
    *request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)

    e := initializeTest("PostShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid scaleY
*/
func TestPostShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.scaleY = new(float64)
    *request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)

    e := initializeTest("PostShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid bounds
*/
func TestPostShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.bounds = new(int32)
    *request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "int32").(int32)

    e := initializeTest("PostShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PostShapeSaveAs method with invalid fontsFolder
*/
func TestPostShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostShapeSaveAs(request)
    assertError(t, "PostShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PostSlideSaveAs", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlideSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid slideIndex
*/
func TestPostSlideSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostSlideSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid format
*/
func TestPostSlideSaveAsInvalidformat(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostSlideSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid options
*/
func TestPostSlideSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PostSlideSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid width
*/
func TestPostSlideSaveAsInvalidwidth(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.width = new(int32)
    *request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)

    e := initializeTest("PostSlideSaveAs", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid height
*/
func TestPostSlideSaveAsInvalidheight(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.height = new(int32)
    *request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)

    e := initializeTest("PostSlideSaveAs", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid password
*/
func TestPostSlideSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlideSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid folder
*/
func TestPostSlideSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlideSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid storage
*/
func TestPostSlideSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlideSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PostSlideSaveAs method with invalid fontsFolder
*/
func TestPostSlideSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostSlideSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesAdd", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid position
*/
func TestPostSlidesAddInvalidposition(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostSlidesAdd", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "position", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid password
*/
func TestPostSlidesAddInvalidpassword(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesAdd", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid folder
*/
func TestPostSlidesAddInvalidfolder(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesAdd", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid storage
*/
func TestPostSlidesAddInvalidstorage(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesAdd", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a slide.
   Test for SlidesApi.PostSlidesAdd method with invalid layoutAlias
*/
func TestPostSlidesAddInvalidlayoutAlias(t *testing.T) {
    request := createPostSlidesAddRequest()
    request.layoutAlias = invalidizeTestParamValue(request.layoutAlias, "layoutAlias", "string").(string)

    e := initializeTest("PostSlidesAdd", "layoutAlias", request.layoutAlias)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesAdd(request)
    assertError(t, "PostSlidesAdd", "layoutAlias", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PostSlidesConvert", "format", "int32").(int32)
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
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostSlidesConvert", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    assertError(t, "PostSlidesConvert", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid document
*/
func TestPostSlidesConvertInvaliddocument(t *testing.T) {
    request := createPostSlidesConvertRequest()
    request.document = invalidizeTestParamValue(request.document, "document", "[]byte").([]byte)

    e := initializeTest("PostSlidesConvert", "document", request.document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    assertError(t, "PostSlidesConvert", "document", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid password
*/
func TestPostSlidesConvertInvalidpassword(t *testing.T) {
    request := createPostSlidesConvertRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesConvert", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    assertError(t, "PostSlidesConvert", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PostSlidesConvert method with invalid fontsFolder
*/
func TestPostSlidesConvertInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesConvertRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostSlidesConvert", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesConvert(request)
    assertError(t, "PostSlidesConvert", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesCopy", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid slideToCopy
*/
func TestPostSlidesCopyInvalidslideToCopy(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.slideToCopy = invalidizeTestParamValue(request.slideToCopy, "slideToCopy", "int32").(int32)

    e := initializeTest("PostSlidesCopy", "slideToCopy", request.slideToCopy)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "slideToCopy", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid position
*/
func TestPostSlidesCopyInvalidposition(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.position = new(int32)
    *request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)

    e := initializeTest("PostSlidesCopy", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "position", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid source
*/
func TestPostSlidesCopyInvalidsource(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.source = invalidizeTestParamValue(request.source, "source", "string").(string)

    e := initializeTest("PostSlidesCopy", "source", request.source)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "source", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourcePassword
*/
func TestPostSlidesCopyInvalidsourcePassword(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.sourcePassword = invalidizeTestParamValue(request.sourcePassword, "sourcePassword", "string").(string)

    e := initializeTest("PostSlidesCopy", "sourcePassword", request.sourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "sourcePassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid sourceStorage
*/
func TestPostSlidesCopyInvalidsourceStorage(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.sourceStorage = invalidizeTestParamValue(request.sourceStorage, "sourceStorage", "string").(string)

    e := initializeTest("PostSlidesCopy", "sourceStorage", request.sourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "sourceStorage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid password
*/
func TestPostSlidesCopyInvalidpassword(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesCopy", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid folder
*/
func TestPostSlidesCopyInvalidfolder(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesCopy", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Copy a slide from the current or another presentation.
   Test for SlidesApi.PostSlidesCopy method with invalid storage
*/
func TestPostSlidesCopyInvalidstorage(t *testing.T) {
    request := createPostSlidesCopyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesCopy", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesCopy(request)
    assertError(t, "PostSlidesCopy", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesDocument", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid data
*/
func TestPostSlidesDocumentInvaliddata(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.data = invalidizeTestParamValue(request.data, "data", "[]byte").([]byte)

    e := initializeTest("PostSlidesDocument", "data", request.data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "data", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid inputPassword
*/
func TestPostSlidesDocumentInvalidinputPassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.inputPassword = invalidizeTestParamValue(request.inputPassword, "inputPassword", "string").(string)

    e := initializeTest("PostSlidesDocument", "inputPassword", request.inputPassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "inputPassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid password
*/
func TestPostSlidesDocumentInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesDocument", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid storage
*/
func TestPostSlidesDocumentInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesDocument", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocument method with invalid folder
*/
func TestPostSlidesDocumentInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesDocument", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocument(request)
    assertError(t, "PostSlidesDocument", "folder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesDocumentFromHtml", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    assertError(t, "PostSlidesDocumentFromHtml", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid html
*/
func TestPostSlidesDocumentFromHtmlInvalidhtml(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()
    request.html = invalidizeTestParamValue(request.html, "html", "string").(string)

    e := initializeTest("PostSlidesDocumentFromHtml", "html", request.html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    assertError(t, "PostSlidesDocumentFromHtml", "html", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid password
*/
func TestPostSlidesDocumentFromHtmlInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesDocumentFromHtml", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    assertError(t, "PostSlidesDocumentFromHtml", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid storage
*/
func TestPostSlidesDocumentFromHtmlInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesDocumentFromHtml", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    assertError(t, "PostSlidesDocumentFromHtml", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create presentation document from html.
   Test for SlidesApi.PostSlidesDocumentFromHtml method with invalid folder
*/
func TestPostSlidesDocumentFromHtmlInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromHtmlRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesDocumentFromHtml", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromHtml(request)
    assertError(t, "PostSlidesDocumentFromHtml", "folder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePath
*/
func TestPostSlidesDocumentFromSourceInvalidsourcePath(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.sourcePath = invalidizeTestParamValue(request.sourcePath, "sourcePath", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePath", request.sourcePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "sourcePath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourcePassword
*/
func TestPostSlidesDocumentFromSourceInvalidsourcePassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.sourcePassword = invalidizeTestParamValue(request.sourcePassword, "sourcePassword", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "sourcePassword", request.sourcePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "sourcePassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid sourceStorage
*/
func TestPostSlidesDocumentFromSourceInvalidsourceStorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.sourceStorage = invalidizeTestParamValue(request.sourceStorage, "sourceStorage", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "sourceStorage", request.sourceStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "sourceStorage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid password
*/
func TestPostSlidesDocumentFromSourceInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid storage
*/
func TestPostSlidesDocumentFromSourceInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation from an existing source.
   Test for SlidesApi.PostSlidesDocumentFromSource method with invalid folder
*/
func TestPostSlidesDocumentFromSourceInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromSourceRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesDocumentFromSource", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromSource(request)
    assertError(t, "PostSlidesDocumentFromSource", "folder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePath
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplatePath(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.templatePath = invalidizeTestParamValue(request.templatePath, "templatePath", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePath", request.templatePath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid data
*/
func TestPostSlidesDocumentFromTemplateInvaliddata(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.data = invalidizeTestParamValue(request.data, "data", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "data", request.data)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "data", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templatePassword
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplatePassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.templatePassword = invalidizeTestParamValue(request.templatePassword, "templatePassword", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "templatePassword", request.templatePassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "templatePassword", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid templateStorage
*/
func TestPostSlidesDocumentFromTemplateInvalidtemplateStorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.templateStorage = invalidizeTestParamValue(request.templateStorage, "templateStorage", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "templateStorage", request.templateStorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "templateStorage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid isImageDataEmbedded
*/
func TestPostSlidesDocumentFromTemplateInvalidisImageDataEmbedded(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.isImageDataEmbedded = new(bool)
    *request.isImageDataEmbedded = invalidizeTestParamValue(request.isImageDataEmbedded, "isImageDataEmbedded", "bool").(bool)

    e := initializeTest("PostSlidesDocumentFromTemplate", "isImageDataEmbedded", request.isImageDataEmbedded)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "isImageDataEmbedded", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid password
*/
func TestPostSlidesDocumentFromTemplateInvalidpassword(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid storage
*/
func TestPostSlidesDocumentFromTemplateInvalidstorage(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Create a presentation.
   Test for SlidesApi.PostSlidesDocumentFromTemplate method with invalid folder
*/
func TestPostSlidesDocumentFromTemplateInvalidfolder(t *testing.T) {
    request := createPostSlidesDocumentFromTemplateRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesDocumentFromTemplate", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesDocumentFromTemplate(request)
    assertError(t, "PostSlidesDocumentFromTemplate", "folder", int32(r.StatusCode), e)
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
    request.pipeline = invalidizeTestParamValue(request.pipeline, "pipeline", "Pipeline").(IPipeline)

    e := initializeTest("PostSlidesPipeline", "pipeline", request.pipeline)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPipeline(request)
    assertError(t, "PostSlidesPipeline", "pipeline", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid oldValue
*/
func TestPostSlidesPresentationReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.oldValue = invalidizeTestParamValue(request.oldValue, "oldValue", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "oldValue", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid newValue
*/
func TestPostSlidesPresentationReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.newValue = invalidizeTestParamValue(request.newValue, "newValue", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "newValue", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid ignoreCase
*/
func TestPostSlidesPresentationReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.ignoreCase = new(bool)
    *request.ignoreCase = invalidizeTestParamValue(request.ignoreCase, "ignoreCase", "bool").(bool)

    e := initializeTest("PostSlidesPresentationReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "ignoreCase", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid password
*/
func TestPostSlidesPresentationReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid folder
*/
func TestPostSlidesPresentationReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesPresentationReplaceText method with invalid storage
*/
func TestPostSlidesPresentationReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesPresentationReplaceTextRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesPresentationReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesPresentationReplaceText(request)
    assertError(t, "PostSlidesPresentationReplaceText", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesReorder", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid slideIndex
*/
func TestPostSlidesReorderInvalidslideIndex(t *testing.T) {
    request := createPostSlidesReorderRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostSlidesReorder", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid newPosition
*/
func TestPostSlidesReorderInvalidnewPosition(t *testing.T) {
    request := createPostSlidesReorderRequest()
    request.newPosition = invalidizeTestParamValue(request.newPosition, "newPosition", "int32").(int32)

    e := initializeTest("PostSlidesReorder", "newPosition", request.newPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "newPosition", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid password
*/
func TestPostSlidesReorderInvalidpassword(t *testing.T) {
    request := createPostSlidesReorderRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesReorder", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid folder
*/
func TestPostSlidesReorderInvalidfolder(t *testing.T) {
    request := createPostSlidesReorderRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesReorder", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position.
   Test for SlidesApi.PostSlidesReorder method with invalid storage
*/
func TestPostSlidesReorderInvalidstorage(t *testing.T) {
    request := createPostSlidesReorderRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesReorder", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorder(request)
    assertError(t, "PostSlidesReorder", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesReorderMany", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid oldPositions
*/
func TestPostSlidesReorderManyInvalidoldPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    request.oldPositions = invalidizeTestParamValue(request.oldPositions, "oldPositions", "[]int32").([]int32)

    e := initializeTest("PostSlidesReorderMany", "oldPositions", request.oldPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "oldPositions", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid newPositions
*/
func TestPostSlidesReorderManyInvalidnewPositions(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    request.newPositions = invalidizeTestParamValue(request.newPositions, "newPositions", "[]int32").([]int32)

    e := initializeTest("PostSlidesReorderMany", "newPositions", request.newPositions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "newPositions", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid password
*/
func TestPostSlidesReorderManyInvalidpassword(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesReorderMany", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid folder
*/
func TestPostSlidesReorderManyInvalidfolder(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesReorderMany", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slides positions.
   Test for SlidesApi.PostSlidesReorderMany method with invalid storage
*/
func TestPostSlidesReorderManyInvalidstorage(t *testing.T) {
    request := createPostSlidesReorderManyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesReorderMany", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesReorderMany(request)
    assertError(t, "PostSlidesReorderMany", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PostSlidesSaveAs", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid format
*/
func TestPostSlidesSaveAsInvalidformat(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostSlidesSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid options
*/
func TestPostSlidesSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PostSlidesSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid password
*/
func TestPostSlidesSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid storage
*/
func TestPostSlidesSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid folder
*/
func TestPostSlidesSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PostSlidesSaveAs method with invalid fontsFolder
*/
func TestPostSlidesSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostSlidesSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSaveAs(request)
    assertError(t, "PostSlidesSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid properties
*/
func TestPostSlidesSetDocumentPropertiesInvalidproperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.properties = invalidizeTestParamValue(request.properties, "properties", "DocumentProperties").(IDocumentProperties)

    e := initializeTest("PostSlidesSetDocumentProperties", "properties", request.properties)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "properties", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid password
*/
func TestPostSlidesSetDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid folder
*/
func TestPostSlidesSetDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document properties.
   Test for SlidesApi.PostSlidesSetDocumentProperties method with invalid storage
*/
func TestPostSlidesSetDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid slideIndex
*/
func TestPostSlidesSlideReplaceTextInvalidslideIndex(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PostSlidesSlideReplaceText", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid oldValue
*/
func TestPostSlidesSlideReplaceTextInvalidoldValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.oldValue = invalidizeTestParamValue(request.oldValue, "oldValue", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "oldValue", request.oldValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "oldValue", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid newValue
*/
func TestPostSlidesSlideReplaceTextInvalidnewValue(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.newValue = invalidizeTestParamValue(request.newValue, "newValue", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "newValue", request.newValue)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "newValue", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid ignoreCase
*/
func TestPostSlidesSlideReplaceTextInvalidignoreCase(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.ignoreCase = new(bool)
    *request.ignoreCase = invalidizeTestParamValue(request.ignoreCase, "ignoreCase", "bool").(bool)

    e := initializeTest("PostSlidesSlideReplaceText", "ignoreCase", request.ignoreCase)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "ignoreCase", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid password
*/
func TestPostSlidesSlideReplaceTextInvalidpassword(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid folder
*/
func TestPostSlidesSlideReplaceTextInvalidfolder(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Replace text with a new value.
   Test for SlidesApi.PostSlidesSlideReplaceText method with invalid storage
*/
func TestPostSlidesSlideReplaceTextInvalidstorage(t *testing.T) {
    request := createPostSlidesSlideReplaceTextRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSlideReplaceText", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSlideReplaceText(request)
    assertError(t, "PostSlidesSlideReplaceText", "storage", int32(r.StatusCode), e)
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
    testformat := createTestParamValue("PostSlidesSplit", "format", "int32")
    switch v := testformat.(type) { 
    case int32:
        request.format = new(int32)
        *request.format = v
    }
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSplit", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid options
*/
func TestPostSlidesSplitInvalidoptions(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PostSlidesSplit", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid format
*/
func TestPostSlidesSplitInvalidformat(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.format = new(int32)
    *request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PostSlidesSplit", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid width
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
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid height
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
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid to
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
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "to", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid from
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
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "from", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid destFolder
*/
func TestPostSlidesSplitInvaliddestFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.destFolder = invalidizeTestParamValue(request.destFolder, "destFolder", "string").(string)

    e := initializeTest("PostSlidesSplit", "destFolder", request.destFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "destFolder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid password
*/
func TestPostSlidesSplitInvalidpassword(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSplit", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid storage
*/
func TestPostSlidesSplitInvalidstorage(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSplit", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid folder
*/
func TestPostSlidesSplitInvalidfolder(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSplit", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Splitting presentations. Create one image per slide.
   Test for SlidesApi.PostSlidesSplit method with invalid fontsFolder
*/
func TestPostSlidesSplitInvalidfontsFolder(t *testing.T) {
    request := createPostSlidesSplitRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PostSlidesSplit", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlidesSplit(request)
    assertError(t, "PostSlidesSplit", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutLayoutSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideIndex
*/
func TestPutLayoutSlideInvalidslideIndex(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutLayoutSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid slideDto
*/
func TestPutLayoutSlideInvalidslideDto(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.slideDto = invalidizeTestParamValue(request.slideDto, "slideDto", "LayoutSlide").(ILayoutSlide)

    e := initializeTest("PutLayoutSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "slideDto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid password
*/
func TestPutLayoutSlideInvalidpassword(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutLayoutSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid folder
*/
func TestPutLayoutSlideInvalidfolder(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutLayoutSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a layoutSlide.
   Test for SlidesApi.PutLayoutSlide method with invalid storage
*/
func TestPutLayoutSlideInvalidstorage(t *testing.T) {
    request := createPutLayoutSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutLayoutSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutLayoutSlide(request)
    assertError(t, "PutLayoutSlide", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PutNotesSlideShapeSaveAs", "format", "int32").(int32)
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
    testbounds := createTestParamValue("PutNotesSlideShapeSaveAs", "bounds", "int32")
    switch v := testbounds.(type) { 
    case int32:
        request.bounds = new(int32)
        *request.bounds = v
    }
    request.fontsFolder = createTestParamValue("PutNotesSlideShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid name
*/
func TestPutNotesSlideShapeSaveAsInvalidname(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid slideIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutNotesSlideShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid path
*/
func TestPutNotesSlideShapeSaveAsInvalidpath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid shapeIndex
*/
func TestPutNotesSlideShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutNotesSlideShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid format
*/
func TestPutNotesSlideShapeSaveAsInvalidformat(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PutNotesSlideShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid outPath
*/
func TestPutNotesSlideShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid options
*/
func TestPutNotesSlideShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)

    e := initializeTest("PutNotesSlideShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid password
*/
func TestPutNotesSlideShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid folder
*/
func TestPutNotesSlideShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid storage
*/
func TestPutNotesSlideShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleX
*/
func TestPutNotesSlideShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.scaleX = new(float64)
    *request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid scaleY
*/
func TestPutNotesSlideShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.scaleY = new(float64)
    *request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)

    e := initializeTest("PutNotesSlideShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid bounds
*/
func TestPutNotesSlideShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.bounds = new(int32)
    *request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "int32").(int32)

    e := initializeTest("PutNotesSlideShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutNotesSlideShapeSaveAs method with invalid fontsFolder
*/
func TestPutNotesSlideShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutNotesSlideShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutNotesSlideShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutNotesSlideShapeSaveAs(request)
    assertError(t, "PutNotesSlideShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutPresentationMerge", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid request
*/
func TestPutPresentationMergeInvalidrequest(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.request = invalidizeTestParamValue(request.request, "request", "OrderedMergeRequest").(IOrderedMergeRequest)

    e := initializeTest("PutPresentationMerge", "request", request.request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "request", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid password
*/
func TestPutPresentationMergeInvalidpassword(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutPresentationMerge", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid storage
*/
func TestPutPresentationMergeInvalidstorage(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutPresentationMerge", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Merge the presentation with other presentations or some of their slides specified in the request parameter.
   Test for SlidesApi.PutPresentationMerge method with invalid folder
*/
func TestPutPresentationMergeInvalidfolder(t *testing.T) {
    request := createPutPresentationMergeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutPresentationMerge", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutPresentationMerge(request)
    assertError(t, "PutPresentationMerge", "folder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSetParagraphPortionProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid slideIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphPortionProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid path
*/
func TestPutSetParagraphPortionPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutSetParagraphPortionProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphPortionProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphPortionProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid portionIndex
*/
func TestPutSetParagraphPortionPropertiesInvalidportionIndex(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphPortionProperties", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid dto
*/
func TestPutSetParagraphPortionPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)

    e := initializeTest("PutSetParagraphPortionProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid password
*/
func TestPutSetParagraphPortionPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSetParagraphPortionProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid folder
*/
func TestPutSetParagraphPortionPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSetParagraphPortionProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutSetParagraphPortionProperties method with invalid storage
*/
func TestPutSetParagraphPortionPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPortionPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSetParagraphPortionProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphPortionProperties(request)
    assertError(t, "PutSetParagraphPortionProperties", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSetParagraphProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid slideIndex
*/
func TestPutSetParagraphPropertiesInvalidslideIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphProperties", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid path
*/
func TestPutSetParagraphPropertiesInvalidpath(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutSetParagraphProperties", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid shapeIndex
*/
func TestPutSetParagraphPropertiesInvalidshapeIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphProperties", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid paragraphIndex
*/
func TestPutSetParagraphPropertiesInvalidparagraphIndex(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PutSetParagraphProperties", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid dto
*/
func TestPutSetParagraphPropertiesInvaliddto(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)

    e := initializeTest("PutSetParagraphProperties", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid password
*/
func TestPutSetParagraphPropertiesInvalidpassword(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSetParagraphProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid folder
*/
func TestPutSetParagraphPropertiesInvalidfolder(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSetParagraphProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutSetParagraphProperties method with invalid storage
*/
func TestPutSetParagraphPropertiesInvalidstorage(t *testing.T) {
    request := createPutSetParagraphPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSetParagraphProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSetParagraphProperties(request)
    assertError(t, "PutSetParagraphProperties", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PutShapeSaveAs", "format", "int32").(int32)
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
    testbounds := createTestParamValue("PutShapeSaveAs", "bounds", "int32")
    switch v := testbounds.(type) { 
    case int32:
        request.bounds = new(int32)
        *request.bounds = v
    }
    request.fontsFolder = createTestParamValue("PutShapeSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid name
*/
func TestPutShapeSaveAsInvalidname(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutShapeSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid slideIndex
*/
func TestPutShapeSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutShapeSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid path
*/
func TestPutShapeSaveAsInvalidpath(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutShapeSaveAs", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid shapeIndex
*/
func TestPutShapeSaveAsInvalidshapeIndex(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutShapeSaveAs", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid format
*/
func TestPutShapeSaveAsInvalidformat(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PutShapeSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid outPath
*/
func TestPutShapeSaveAsInvalidoutPath(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutShapeSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid options
*/
func TestPutShapeSaveAsInvalidoptions(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "IShapeExportOptions").(IIShapeExportOptions)

    e := initializeTest("PutShapeSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid password
*/
func TestPutShapeSaveAsInvalidpassword(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutShapeSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid folder
*/
func TestPutShapeSaveAsInvalidfolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutShapeSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid storage
*/
func TestPutShapeSaveAsInvalidstorage(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutShapeSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleX
*/
func TestPutShapeSaveAsInvalidscaleX(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.scaleX = new(float64)
    *request.scaleX = invalidizeTestParamValue(request.scaleX, "scaleX", "float64").(float64)

    e := initializeTest("PutShapeSaveAs", "scaleX", request.scaleX)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "scaleX", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid scaleY
*/
func TestPutShapeSaveAsInvalidscaleY(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.scaleY = new(float64)
    *request.scaleY = invalidizeTestParamValue(request.scaleY, "scaleY", "float64").(float64)

    e := initializeTest("PutShapeSaveAs", "scaleY", request.scaleY)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "scaleY", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid bounds
*/
func TestPutShapeSaveAsInvalidbounds(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.bounds = new(int32)
    *request.bounds = invalidizeTestParamValue(request.bounds, "bounds", "int32").(int32)

    e := initializeTest("PutShapeSaveAs", "bounds", request.bounds)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "bounds", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Render shape to specified picture format.
   Test for SlidesApi.PutShapeSaveAs method with invalid fontsFolder
*/
func TestPutShapeSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutShapeSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutShapeSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutShapeSaveAs(request)
    assertError(t, "PutShapeSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PutSlideSaveAs", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlideSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid slideIndex
*/
func TestPutSlideSaveAsInvalidslideIndex(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSlideSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid format
*/
func TestPutSlideSaveAsInvalidformat(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PutSlideSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid outPath
*/
func TestPutSlideSaveAsInvalidoutPath(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutSlideSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid options
*/
func TestPutSlideSaveAsInvalidoptions(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PutSlideSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid width
*/
func TestPutSlideSaveAsInvalidwidth(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.width = new(int32)
    *request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)

    e := initializeTest("PutSlideSaveAs", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid height
*/
func TestPutSlideSaveAsInvalidheight(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.height = new(int32)
    *request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)

    e := initializeTest("PutSlideSaveAs", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid password
*/
func TestPutSlideSaveAsInvalidpassword(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlideSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid folder
*/
func TestPutSlideSaveAsInvalidfolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlideSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid storage
*/
func TestPutSlideSaveAsInvalidstorage(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlideSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a slide to a specified format.
   Test for SlidesApi.PutSlideSaveAs method with invalid fontsFolder
*/
func TestPutSlideSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutSlideSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutSlideSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlideSaveAs(request)
    assertError(t, "PutSlideSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlideShapeInfo", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid slideIndex
*/
func TestPutSlideShapeInfoInvalidslideIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSlideShapeInfo", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid path
*/
func TestPutSlideShapeInfoInvalidpath(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutSlideShapeInfo", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid shapeIndex
*/
func TestPutSlideShapeInfoInvalidshapeIndex(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutSlideShapeInfo", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid dto
*/
func TestPutSlideShapeInfoInvaliddto(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)

    e := initializeTest("PutSlideShapeInfo", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid password
*/
func TestPutSlideShapeInfoInvalidpassword(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlideShapeInfo", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid folder
*/
func TestPutSlideShapeInfoInvalidfolder(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlideShapeInfo", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutSlideShapeInfo method with invalid storage
*/
func TestPutSlideShapeInfoInvalidstorage(t *testing.T) {
    request := createPutSlideShapeInfoRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlideShapeInfo", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlideShapeInfo(request)
    assertError(t, "PutSlideShapeInfo", "storage", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PutSlidesConvert", "format", "int32").(int32)
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
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PutSlidesConvert", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid outPath
*/
func TestPutSlidesConvertInvalidoutPath(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutSlidesConvert", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid document
*/
func TestPutSlidesConvertInvaliddocument(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.document = invalidizeTestParamValue(request.document, "document", "[]byte").([]byte)

    e := initializeTest("PutSlidesConvert", "document", request.document)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "document", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid password
*/
func TestPutSlidesConvertInvalidpassword(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesConvert", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert presentation from request content to format specified.
   Test for SlidesApi.PutSlidesConvert method with invalid fontsFolder
*/
func TestPutSlidesConvertInvalidfontsFolder(t *testing.T) {
    request := createPutSlidesConvertRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutSlidesConvert", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesConvert(request)
    assertError(t, "PutSlidesConvert", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid html
*/
func TestPutSlidesDocumentFromHtmlInvalidhtml(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.html = invalidizeTestParamValue(request.html, "html", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "html", request.html)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "html", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid password
*/
func TestPutSlidesDocumentFromHtmlInvalidpassword(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid storage
*/
func TestPutSlidesDocumentFromHtmlInvalidstorage(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update presentation document from html.
   Test for SlidesApi.PutSlidesDocumentFromHtml method with invalid folder
*/
func TestPutSlidesDocumentFromHtmlInvalidfolder(t *testing.T) {
    request := createPutSlidesDocumentFromHtmlRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesDocumentFromHtml", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesDocumentFromHtml(request)
    assertError(t, "PutSlidesDocumentFromHtml", "folder", int32(r.StatusCode), e)
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
    request.format = createTestParamValue("PutSlidesSaveAs", "format", "int32").(int32)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid outPath
*/
func TestPutSlidesSaveAsInvalidoutPath(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid format
*/
func TestPutSlidesSaveAsInvalidformat(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "int32").(int32)

    e := initializeTest("PutSlidesSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid options
*/
func TestPutSlidesSaveAsInvalidoptions(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "ExportOptions").(IExportOptions)

    e := initializeTest("PutSlidesSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid password
*/
func TestPutSlidesSaveAsInvalidpassword(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid storage
*/
func TestPutSlidesSaveAsInvalidstorage(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid folder
*/
func TestPutSlidesSaveAsInvalidfolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Save a presentation to a specified format.
   Test for SlidesApi.PutSlidesSaveAs method with invalid fontsFolder
*/
func TestPutSlidesSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPutSlidesSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)

    e := initializeTest("PutSlidesSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, e := getTestApiClient().SlidesApi.PutSlidesSaveAs(request)
    assertError(t, "PutSlidesSaveAs", "fontsFolder", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid propertyName
*/
func TestPutSlidesSetDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "propertyName", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid property
*/
func TestPutSlidesSetDocumentPropertyInvalidproperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.property = invalidizeTestParamValue(request.property, "property", "DocumentProperty").(IDocumentProperty)

    e := initializeTest("PutSlidesSetDocumentProperty", "property", request.property)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "property", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid password
*/
func TestPutSlidesSetDocumentPropertyInvalidpassword(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid folder
*/
func TestPutSlidesSetDocumentPropertyInvalidfolder(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set document property.
   Test for SlidesApi.PutSlidesSetDocumentProperty method with invalid storage
*/
func TestPutSlidesSetDocumentPropertyInvalidstorage(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideIndex
*/
func TestPutSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid slideDto
*/
func TestPutSlidesSlideInvalidslideDto(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.slideDto = invalidizeTestParamValue(request.slideDto, "slideDto", "Slide").(ISlide)

    e := initializeTest("PutSlidesSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "slideDto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid password
*/
func TestPutSlidesSlideInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid folder
*/
func TestPutSlidesSlideInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update a slide.
   Test for SlidesApi.PutSlidesSlide method with invalid storage
*/
func TestPutSlidesSlideInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid background
*/
func TestPutSlidesSlideBackgroundInvalidbackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.background = invalidizeTestParamValue(request.background, "background", "SlideBackground").(ISlideBackground)

    e := initializeTest("PutSlidesSlideBackground", "background", request.background)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "background", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid folder
*/
func TestPutSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid password
*/
func TestPutSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background for a slide.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid storage
*/
func TestPutSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSlideBackgroundColor", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundColorInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutSlidesSlideBackgroundColor", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid color
*/
func TestPutSlidesSlideBackgroundColorInvalidcolor(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    request.color = invalidizeTestParamValue(request.color, "color", "string").(string)

    e := initializeTest("PutSlidesSlideBackgroundColor", "color", request.color)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "color", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid folder
*/
func TestPutSlidesSlideBackgroundColorInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSlideBackgroundColor", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid password
*/
func TestPutSlidesSlideBackgroundColorInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSlideBackgroundColor", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set background color for a slide.
   Test for SlidesApi.PutSlidesSlideBackgroundColor method with invalid storage
*/
func TestPutSlidesSlideBackgroundColorInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundColorRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSlideBackgroundColor", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideBackgroundColor(request)
    assertError(t, "PutSlidesSlideBackgroundColor", "storage", int32(r.StatusCode), e)
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
    testsizeType := createTestParamValue("PutSlidesSlideSize", "sizeType", "int32")
    switch v := testsizeType.(type) { 
    case int32:
        request.sizeType = new(int32)
        *request.sizeType = v
    }
    testscaleType := createTestParamValue("PutSlidesSlideSize", "scaleType", "int32")
    switch v := testscaleType.(type) { 
    case int32:
        request.scaleType = new(int32)
        *request.scaleType = v
    }
    return request
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid name
*/
func TestPutSlidesSlideSizeInvalidname(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid password
*/
func TestPutSlidesSlideSizeInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid storage
*/
func TestPutSlidesSlideSizeInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid folder
*/
func TestPutSlidesSlideSizeInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSlideSize", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid width
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
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid height
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
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid sizeType
*/
func TestPutSlidesSlideSizeInvalidsizeType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.sizeType = new(int32)
    *request.sizeType = invalidizeTestParamValue(request.sizeType, "sizeType", "int32").(int32)

    e := initializeTest("PutSlidesSlideSize", "sizeType", request.sizeType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "sizeType", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Set slide size for a presentation.
   Test for SlidesApi.PutSlidesSlideSize method with invalid scaleType
*/
func TestPutSlidesSlideSizeInvalidscaleType(t *testing.T) {
    request := createPutSlidesSlideSizeRequest()
    request.scaleType = new(int32)
    *request.scaleType = invalidizeTestParamValue(request.scaleType, "scaleType", "int32").(int32)

    e := initializeTest("PutSlidesSlideSize", "scaleType", request.scaleType)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutSlidesSlideSize(request)
    assertError(t, "PutSlidesSlideSize", "scaleType", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid slideIndex
*/
func TestPutUpdateNotesSlideInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid dto
*/
func TestPutUpdateNotesSlideInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "NotesSlide").(INotesSlide)

    e := initializeTest("PutUpdateNotesSlide", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid password
*/
func TestPutUpdateNotesSlideInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid folder
*/
func TestPutUpdateNotesSlideInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update notes slide properties.
   Test for SlidesApi.PutUpdateNotesSlide method with invalid storage
*/
func TestPutUpdateNotesSlideInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutUpdateNotesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlide(request)
    assertError(t, "PutUpdateNotesSlide", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShape", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShape", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid path
*/
func TestPutUpdateNotesSlideShapeInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShape", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShape", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid dto
*/
func TestPutUpdateNotesSlideShapeInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "ShapeBase").(IShapeBase)

    e := initializeTest("PutUpdateNotesSlideShape", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid password
*/
func TestPutUpdateNotesSlideShapeInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShape", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid folder
*/
func TestPutUpdateNotesSlideShapeInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShape", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update shape properties.
   Test for SlidesApi.PutUpdateNotesSlideShape method with invalid storage
*/
func TestPutUpdateNotesSlideShapeInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShape", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShape(request)
    assertError(t, "PutUpdateNotesSlideShape", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid path
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid dto
*/
func TestPutUpdateNotesSlideShapeParagraphInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Paragraph").(IParagraph)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid password
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid folder
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update paragraph properties.
   Test for SlidesApi.PutUpdateNotesSlideShapeParagraph method with invalid storage
*/
func TestPutUpdateNotesSlideShapeParagraphInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapeParagraphRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapeParagraph", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapeParagraph(request)
    assertError(t, "PutUpdateNotesSlideShapeParagraph", "storage", int32(r.StatusCode), e)
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
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid slideIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidslideIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid path
*/
func TestPutUpdateNotesSlideShapePortionInvalidpath(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid shapeIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidshapeIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.shapeIndex = invalidizeTestParamValue(request.shapeIndex, "shapeIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "shapeIndex", request.shapeIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "shapeIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid paragraphIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidparagraphIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.paragraphIndex = invalidizeTestParamValue(request.paragraphIndex, "paragraphIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "paragraphIndex", request.paragraphIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "paragraphIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid portionIndex
*/
func TestPutUpdateNotesSlideShapePortionInvalidportionIndex(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.portionIndex = invalidizeTestParamValue(request.portionIndex, "portionIndex", "int32").(int32)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "portionIndex", request.portionIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "portionIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid dto
*/
func TestPutUpdateNotesSlideShapePortionInvaliddto(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.dto = invalidizeTestParamValue(request.dto, "dto", "Portion").(IPortion)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "dto", request.dto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "dto", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid password
*/
func TestPutUpdateNotesSlideShapePortionInvalidpassword(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid folder
*/
func TestPutUpdateNotesSlideShapePortionInvalidfolder(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Update portion properties.
   Test for SlidesApi.PutUpdateNotesSlideShapePortion method with invalid storage
*/
func TestPutUpdateNotesSlideShapePortionInvalidstorage(t *testing.T) {
    request := createPutUpdateNotesSlideShapePortionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutUpdateNotesSlideShapePortion", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PutUpdateNotesSlideShapePortion(request)
    assertError(t, "PutUpdateNotesSlideShapePortion", "storage", int32(r.StatusCode), e)
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
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("StorageExists", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.StorageExists(request)
    assertError(t, "StorageExists", "storageName", int32(r.StatusCode), e)
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
    request.path = invalidizeTestParamValue(request.path, "path", "string").(string)

    e := initializeTest("UploadFile", "path", request.path)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    assertError(t, "UploadFile", "path", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid file
*/
func TestUploadFileInvalidfile(t *testing.T) {
    request := createUploadFileRequest()
    request.file = invalidizeTestParamValue(request.file, "file", "[]byte").([]byte)

    e := initializeTest("UploadFile", "file", request.file)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    assertError(t, "UploadFile", "file", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Upload file
   Test for SlidesApi.UploadFile method with invalid storageName
*/
func TestUploadFileInvalidstorageName(t *testing.T) {
    request := createUploadFileRequest()
    request.storageName = invalidizeTestParamValue(request.storageName, "storageName", "string").(string)

    e := initializeTest("UploadFile", "storageName", request.storageName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.UploadFile(request)
    assertError(t, "UploadFile", "storageName", int32(r.StatusCode), e)
}
