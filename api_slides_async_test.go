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
   Test for SlidesAsyncApi.StartConvertAndSave method
*/
func TestStartConvertAndSave(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)
    e := InitializeTest("StartConvertAndSave", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid document
*/
func TestStartConvertAndSaveInvalidDocument(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testdocument, "StartConvertAndSave", "document", "[]byte")
    if (invalidValue == nil) {
        testdocument = nil
    } else {
        testdocument, _ = invalidValue.([]byte)
    }

    e := InitializeTest("StartConvertAndSave", "document", testdocument)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "document", "[]byte", testdocument, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid format
*/
func TestStartConvertAndSaveInvalidFormat(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testformat, "StartConvertAndSave", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        testformat = nullValue
    } else {
        testformat, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvertAndSave", "format", testformat)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "format", "string", testformat, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid outPath
*/
func TestStartConvertAndSaveInvalidOutPath(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testoutPath, "StartConvertAndSave", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        testoutPath = nullValue
    } else {
        testoutPath, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvertAndSave", "outPath", testoutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "outPath", "string", testoutPath, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid password
*/
func TestStartConvertAndSaveInvalidPassword(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testpassword, "StartConvertAndSave", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        testpassword = nullValue
    } else {
        testpassword, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvertAndSave", "password", testpassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "password", "string", testpassword, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid storage
*/
func TestStartConvertAndSaveInvalidStorage(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(teststorage, "StartConvertAndSave", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvertAndSave", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid fontsFolder
*/
func TestStartConvertAndSaveInvalidFontsFolder(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testfontsFolder, "StartConvertAndSave", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfontsFolder = nullValue
    } else {
        testfontsFolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartConvertAndSave", "fontsFolder", testfontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "fontsFolder", "string", testfontsFolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid slides
*/
func TestStartConvertAndSaveInvalidSlides(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testslides, "StartConvertAndSave", "slides", "[]int32")
    if (invalidValue == nil) {
        testslides = nil
    } else {
        testslides, _ = invalidValue.([]int32)
    }

    e := InitializeTest("StartConvertAndSave", "slides", testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "slides", "[]int32", testslides, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartConvertAndSave method with invalid options
*/
func TestStartConvertAndSaveInvalidOptions(t *testing.T) {
    testdocument, _ := createTestParamValue("StartConvertAndSave", "document", "[]byte").([]byte)
    testformat, _ := createTestParamValue("StartConvertAndSave", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartConvertAndSave", "outPath", "string").(string)
    testpassword, _ := createTestParamValue("StartConvertAndSave", "password", "string").(string)
    teststorage, _ := createTestParamValue("StartConvertAndSave", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartConvertAndSave", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartConvertAndSave", "slides", "[]int32").([]int32)
    testoptions, _ := createTestParamValue("StartConvertAndSave", "options", "ExportOptions").(IExportOptions)

    invalidValue := invalidizeTestParamValue(testoptions, "StartConvertAndSave", "options", "ExportOptions")
    if (invalidValue == nil) {
        testoptions = nil
    } else {
        testoptions, _ = invalidValue.(IExportOptions)
    }

    e := InitializeTest("StartConvertAndSave", "options", testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartConvertAndSave(testdocument, testformat, testoutPath, testpassword, teststorage, testfontsFolder, testslides, testoptions)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartConvertAndSave", "options", "ExportOptions", testoptions, int32(statusCode), e)
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

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMerge method
*/
func TestStartMerge(t *testing.T) {
    testfiles, _ := createTestParamValue("StartMerge", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMerge", "storage", "string").(string)
    e := InitializeTest("StartMerge", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartMerge(testfiles, testrequest, teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMerge method with invalid files
*/
func TestStartMergeInvalidFiles(t *testing.T) {
    testfiles, _ := createTestParamValue("StartMerge", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMerge", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(testfiles, "StartMerge", "files", "[][]byte")
    if (invalidValue == nil) {
        testfiles = nil
    } else {
        testfiles, _ = invalidValue.([][]byte)
    }

    e := InitializeTest("StartMerge", "files", testfiles)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMerge(testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMerge", "files", "[][]byte", testfiles, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMerge method with invalid request
*/
func TestStartMergeInvalidRequest(t *testing.T) {
    testfiles, _ := createTestParamValue("StartMerge", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMerge", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(testrequest, "StartMerge", "request", "OrderedMergeRequest")
    if (invalidValue == nil) {
        testrequest = nil
    } else {
        testrequest, _ = invalidValue.(IOrderedMergeRequest)
    }

    e := InitializeTest("StartMerge", "request", testrequest)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMerge(testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMerge", "request", "OrderedMergeRequest", testrequest, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMerge method with invalid storage
*/
func TestStartMergeInvalidStorage(t *testing.T) {
    testfiles, _ := createTestParamValue("StartMerge", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMerge", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMerge", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(teststorage, "StartMerge", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartMerge", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMerge(testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMerge", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMergeAndSave method
*/
func TestStartMergeAndSave(t *testing.T) {
    testoutPath, _ := createTestParamValue("StartMergeAndSave", "outPath", "string").(string)
    testfiles, _ := createTestParamValue("StartMergeAndSave", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMergeAndSave", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMergeAndSave", "storage", "string").(string)
    e := InitializeTest("StartMergeAndSave", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartMergeAndSave(testoutPath, testfiles, testrequest, teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMergeAndSave method with invalid outPath
*/
func TestStartMergeAndSaveInvalidOutPath(t *testing.T) {
    testoutPath, _ := createTestParamValue("StartMergeAndSave", "outPath", "string").(string)
    testfiles, _ := createTestParamValue("StartMergeAndSave", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMergeAndSave", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMergeAndSave", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(testoutPath, "StartMergeAndSave", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        testoutPath = nullValue
    } else {
        testoutPath, _ = invalidValue.(string)
    }

    e := InitializeTest("StartMergeAndSave", "outPath", testoutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMergeAndSave(testoutPath, testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMergeAndSave", "outPath", "string", testoutPath, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMergeAndSave method with invalid files
*/
func TestStartMergeAndSaveInvalidFiles(t *testing.T) {
    testoutPath, _ := createTestParamValue("StartMergeAndSave", "outPath", "string").(string)
    testfiles, _ := createTestParamValue("StartMergeAndSave", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMergeAndSave", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMergeAndSave", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(testfiles, "StartMergeAndSave", "files", "[][]byte")
    if (invalidValue == nil) {
        testfiles = nil
    } else {
        testfiles, _ = invalidValue.([][]byte)
    }

    e := InitializeTest("StartMergeAndSave", "files", testfiles)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMergeAndSave(testoutPath, testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMergeAndSave", "files", "[][]byte", testfiles, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMergeAndSave method with invalid request
*/
func TestStartMergeAndSaveInvalidRequest(t *testing.T) {
    testoutPath, _ := createTestParamValue("StartMergeAndSave", "outPath", "string").(string)
    testfiles, _ := createTestParamValue("StartMergeAndSave", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMergeAndSave", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMergeAndSave", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(testrequest, "StartMergeAndSave", "request", "OrderedMergeRequest")
    if (invalidValue == nil) {
        testrequest = nil
    } else {
        testrequest, _ = invalidValue.(IOrderedMergeRequest)
    }

    e := InitializeTest("StartMergeAndSave", "request", testrequest)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMergeAndSave(testoutPath, testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMergeAndSave", "request", "OrderedMergeRequest", testrequest, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartMergeAndSave method with invalid storage
*/
func TestStartMergeAndSaveInvalidStorage(t *testing.T) {
    testoutPath, _ := createTestParamValue("StartMergeAndSave", "outPath", "string").(string)
    testfiles, _ := createTestParamValue("StartMergeAndSave", "files", "[][]byte").([][]byte)
    testrequest, _ := createTestParamValue("StartMergeAndSave", "request", "OrderedMergeRequest").(IOrderedMergeRequest)
    teststorage, _ := createTestParamValue("StartMergeAndSave", "storage", "string").(string)

    invalidValue := invalidizeTestParamValue(teststorage, "StartMergeAndSave", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartMergeAndSave", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartMergeAndSave(testoutPath, testfiles, testrequest, teststorage)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartMergeAndSave", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method
*/
func TestStartSavePresentation(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)
    e := InitializeTest("StartSavePresentation", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := GetTestSlidesAsyncApiClient()
    _, _, e = c.SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid name
*/
func TestStartSavePresentationInvalidName(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testname, "StartSavePresentation", "name", "string")
    if (invalidValue == nil) {
        var nullValue string
        testname = nullValue
    } else {
        testname, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "name", testname)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "name", "string", testname, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid format
*/
func TestStartSavePresentationInvalidFormat(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testformat, "StartSavePresentation", "format", "string")
    if (invalidValue == nil) {
        var nullValue string
        testformat = nullValue
    } else {
        testformat, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "format", testformat)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "format", "string", testformat, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid outPath
*/
func TestStartSavePresentationInvalidOutPath(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testoutPath, "StartSavePresentation", "outPath", "string")
    if (invalidValue == nil) {
        var nullValue string
        testoutPath = nullValue
    } else {
        testoutPath, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "outPath", testoutPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "outPath", "string", testoutPath, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid options
*/
func TestStartSavePresentationInvalidOptions(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testoptions, "StartSavePresentation", "options", "ExportOptions")
    if (invalidValue == nil) {
        testoptions = nil
    } else {
        testoptions, _ = invalidValue.(IExportOptions)
    }

    e := InitializeTest("StartSavePresentation", "options", testoptions)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "options", "ExportOptions", testoptions, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid password
*/
func TestStartSavePresentationInvalidPassword(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testpassword, "StartSavePresentation", "password", "string")
    if (invalidValue == nil) {
        var nullValue string
        testpassword = nullValue
    } else {
        testpassword, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "password", testpassword)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "password", "string", testpassword, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid folder
*/
func TestStartSavePresentationInvalidFolder(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testfolder, "StartSavePresentation", "folder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfolder = nullValue
    } else {
        testfolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "folder", testfolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "folder", "string", testfolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid storage
*/
func TestStartSavePresentationInvalidStorage(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(teststorage, "StartSavePresentation", "storage", "string")
    if (invalidValue == nil) {
        var nullValue string
        teststorage = nullValue
    } else {
        teststorage, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "storage", teststorage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "storage", "string", teststorage, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid fontsFolder
*/
func TestStartSavePresentationInvalidFontsFolder(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testfontsFolder, "StartSavePresentation", "fontsFolder", "string")
    if (invalidValue == nil) {
        var nullValue string
        testfontsFolder = nullValue
    } else {
        testfontsFolder, _ = invalidValue.(string)
    }

    e := InitializeTest("StartSavePresentation", "fontsFolder", testfontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "fontsFolder", "string", testfontsFolder, int32(statusCode), e)
}

/* SlidesAsyncApiServiceTests 
   Test for SlidesAsyncApi.StartSavePresentation method with invalid slides
*/
func TestStartSavePresentationInvalidSlides(t *testing.T) {
    testname, _ := createTestParamValue("StartSavePresentation", "name", "string").(string)
    testformat, _ := createTestParamValue("StartSavePresentation", "format", "string").(string)
    testoutPath, _ := createTestParamValue("StartSavePresentation", "outPath", "string").(string)
    testoptions, _ := createTestParamValue("StartSavePresentation", "options", "ExportOptions").(IExportOptions)
    testpassword, _ := createTestParamValue("StartSavePresentation", "password", "string").(string)
    testfolder, _ := createTestParamValue("StartSavePresentation", "folder", "string").(string)
    teststorage, _ := createTestParamValue("StartSavePresentation", "storage", "string").(string)
    testfontsFolder, _ := createTestParamValue("StartSavePresentation", "fontsFolder", "string").(string)
    testslides, _ := createTestParamValue("StartSavePresentation", "slides", "[]int32").([]int32)

    invalidValue := invalidizeTestParamValue(testslides, "StartSavePresentation", "slides", "[]int32")
    if (invalidValue == nil) {
        testslides = nil
    } else {
        testslides, _ = invalidValue.([]int32)
    }

    e := InitializeTest("StartSavePresentation", "slides", testslides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := GetTestSlidesAsyncApiClient().SlidesAsyncApi.StartSavePresentation(testname, testformat, testoutPath, testoptions, testpassword, testfolder, teststorage, testfontsFolder, testslides)
    statusCode := 400
    if r != nil {
        statusCode = r.StatusCode
    }
    assertError(t, "StartSavePresentation", "slides", "[]int32", testslides, int32(statusCode), e)
}
