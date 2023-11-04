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

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.GetOperationResult method
*/
func TestGetOperationResult(t *testing.T) {
    testid, _ := createTestParamValue("GetOperationResult", "id", "string").(string)
    e := InitializeTest("GetOperationResult", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    r, _, e := c.SlidesAsyncApi.GetOperationResult(testid)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.GetOperationResult method with invalid id
*/
func TestGetOperationResultInvalidId(t *testing.T) {
    testid, _ := createTestParamValue("GetOperationResult", "id", "string").(string)

    invalidValue := invalidizeTestParamValue(testid, "GetOperationResult", "id", "string")
    if (invalidValue == nil) {
        var nullValue string
        testid = nullValue
    } else {
        testid, _ = invalidValue.(string)
    }

    e := InitializeTest("GetOperationResult", "id", testid)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.GetOperationResult(testid)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetOperationResult", "id", "string", testid, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.GetOperationStatus method
*/
func TestGetOperationStatus(t *testing.T) {
    testid, _ := createTestParamValue("GetOperationStatus", "id", "string").(string)
    e := InitializeTest("GetOperationStatus", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.GetOperationStatus(testid)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.GetOperationStatus method with invalid id
*/
func TestGetOperationStatusInvalidId(t *testing.T) {
    testid, _ := createTestParamValue("GetOperationStatus", "id", "string").(string)

    invalidValue := invalidizeTestParamValue(testid, "GetOperationStatus", "id", "string")
    if (invalidValue == nil) {
        var nullValue string
        testid = nullValue
    } else {
        testid, _ = invalidValue.(string)
    }

    e := InitializeTest("GetOperationStatus", "id", testid)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.GetOperationStatus(testid)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "GetOperationStatus", "id", "string", testid, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method
*/
func TestStartConvert(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)
    e := InitializeTest("StartConvert", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid document
*/
func TestStartConvertInvalidDocument(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testdocument, "StartConvert", "document", "[]byte")
    if (invalidValue == nil) {
        testdocument = nil
    } else {
        testdocument, _ = invalidValue.([]byte)
    }

    e := InitializeTest("StartConvert", "document", testdocument)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "document", "[]byte", testdocument, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid format
*/
func TestStartConvertInvalidFormat(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testformat, "StartConvert", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        testformat = nullValue
    } else {
        testformat, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvert", "format", testformat)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "format", "string", testformat, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid password
*/
func TestStartConvertInvalidPassword(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testpassword, "StartConvert", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        testpassword = nullValue
    } else {
        testpassword, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvert", "password", testpassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "password", "string", testpassword, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid storage
*/
func TestStartConvertInvalidStorage(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(teststorage, "StartConvert", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvert", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid fontsFolder
*/
func TestStartConvertInvalidFontsFolder(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testfontsFolder, "StartConvert", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfontsFolder = nullValue
    } else {
        testfontsFolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvert", "fontsFolder", testfontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "fontsFolder", "string", testfontsFolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid slides
*/
func TestStartConvertInvalidSlides(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testslides, "StartConvert", "slides", "[]int32")
    if (invalidValue == nil) {
        testslides = nil
    } else {
        testslides, _ = invalidValue.([]int32)
    }

    e := InitializeTest("StartConvert", "slides", testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "slides", "[]int32", testslides, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvert method with invalid options
*/
func TestStartConvertInvalidOptions(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvert", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvert", "format", "string").(string)
    testpassword, _ := createTestParamValue("StartConvert", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvert", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvert", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvert", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvert", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testoptions, "StartConvert", "options", "ExportOptions")
    if (invalidValue == nil) {
        testoptions = nil
    } else {
        testoptions, _ = invalidValue.(IExportOptions)
    }

    e := InitializeTest("StartConvert", "options", testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvert(testdocument, testformat, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvert", "options", "ExportOptions", testoptions, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method
*/
func TestStartDownloadPresentation(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)
    e := InitializeTest("StartDownloadPresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid name
*/
func TestStartDownloadPresentationInvalidName(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testname, "StartDownloadPresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        testname = nullValue
    } else {
        testname, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "name", testname)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "name", "string", testname, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid format
*/
func TestStartDownloadPresentationInvalidFormat(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testformat, "StartDownloadPresentation", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        testformat = nullValue
    } else {
        testformat, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "format", testformat)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "format", "string", testformat, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid options
*/
func TestStartDownloadPresentationInvalidOptions(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testoptions, "StartDownloadPresentation", "options", "ExportOptions")
    if (invalidValue == nil) {
        testoptions = nil
    } else {
        testoptions, _ = invalidValue.(IExportOptions)
    }

    e := InitializeTest("StartDownloadPresentation", "options", testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "options", "ExportOptions", testoptions, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid password
*/
func TestStartDownloadPresentationInvalidPassword(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testpassword, "StartDownloadPresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        testpassword = nullValue
    } else {
        testpassword, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "password", testpassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "password", "string", testpassword, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid folder
*/
func TestStartDownloadPresentationInvalidFolder(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testfolder, "StartDownloadPresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfolder = nullValue
    } else {
        testfolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "folder", testfolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "folder", "string", testfolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid storage
*/
func TestStartDownloadPresentationInvalidStorage(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(teststorage, "StartDownloadPresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid fontsFolder
*/
func TestStartDownloadPresentationInvalidFontsFolder(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testfontsFolder, "StartDownloadPresentation", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfontsFolder = nullValue
    } else {
        testfontsFolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartDownloadPresentation", "fontsFolder", testfontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "fontsFolder", "string", testfontsFolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartDownloadPresentation method with invalid slides
*/
func TestStartDownloadPresentationInvalidSlides(t *testing.T) {
    testname, _ := createTestParamValue("StartDownloadPresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartDownloadPresentation", "format", "string").(string)
    testoptions, _ := createTestParamValue("StartDownloadPresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartDownloadPresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartDownloadPresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartDownloadPresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartDownloadPresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartDownloadPresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testslides, "StartDownloadPresentation", "slides", "[]int32")
    if (invalidValue == nil) {
        testslides = nil
    } else {
        testslides, _ = invalidValue.([]int32)
    }

    e := InitializeTest("StartDownloadPresentation", "slides", testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartDownloadPresentation(testname, testformat, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartDownloadPresentation", "slides", "[]int32", testslides, int32(statusCode), e)
}
