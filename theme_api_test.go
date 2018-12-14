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

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method
*/
func TestGetSlidesTheme(t *testing.T) {
    request := createGetSlidesThemeRequest()
    e := initializeTest("GetSlidesTheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ThemeApi.GetSlidesTheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method with invalid name
*/
func TestGetSlidesThemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesTheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "name", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method with invalid slideIndex
*/
func TestGetSlidesThemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesTheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "slideIndex", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method with invalid password
*/
func TestGetSlidesThemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesTheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "password", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method with invalid folder
*/
func TestGetSlidesThemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesTheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "folder", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme info.
   Test for ThemeApi.GetSlidesTheme method with invalid storage
*/
func TestGetSlidesThemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesTheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesTheme(request)
    assertError(t, "GetSlidesTheme", "storage", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method
*/
func TestGetSlidesThemeColorScheme(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    e := initializeTest("GetSlidesThemeColorScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ThemeApi.GetSlidesThemeColorScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method with invalid name
*/
func TestGetSlidesThemeColorSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesThemeColorScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "name", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method with invalid slideIndex
*/
func TestGetSlidesThemeColorSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesThemeColorScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "slideIndex", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method with invalid password
*/
func TestGetSlidesThemeColorSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesThemeColorScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "password", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method with invalid folder
*/
func TestGetSlidesThemeColorSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesThemeColorScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "folder", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeColorScheme method with invalid storage
*/
func TestGetSlidesThemeColorSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeColorSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesThemeColorScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeColorScheme(request)
    assertError(t, "GetSlidesThemeColorScheme", "storage", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method
*/
func TestGetSlidesThemeFontScheme(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    e := initializeTest("GetSlidesThemeFontScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ThemeApi.GetSlidesThemeFontScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method with invalid name
*/
func TestGetSlidesThemeFontSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesThemeFontScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "name", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFontSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesThemeFontScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "slideIndex", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method with invalid password
*/
func TestGetSlidesThemeFontSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesThemeFontScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "password", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method with invalid folder
*/
func TestGetSlidesThemeFontSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesThemeFontScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "folder", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme font scheme info.
   Test for ThemeApi.GetSlidesThemeFontScheme method with invalid storage
*/
func TestGetSlidesThemeFontSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFontSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesThemeFontScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFontScheme(request)
    assertError(t, "GetSlidesThemeFontScheme", "storage", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method
*/
func TestGetSlidesThemeFormatScheme(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    e := initializeTest("GetSlidesThemeFormatScheme", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.ThemeApi.GetSlidesThemeFormatScheme(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method with invalid name
*/
func TestGetSlidesThemeFormatSchemeInvalidname(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesThemeFormatScheme", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "name", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method with invalid slideIndex
*/
func TestGetSlidesThemeFormatSchemeInvalidslideIndex(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesThemeFormatScheme", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "slideIndex", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method with invalid password
*/
func TestGetSlidesThemeFormatSchemeInvalidpassword(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesThemeFormatScheme", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "password", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method with invalid folder
*/
func TestGetSlidesThemeFormatSchemeInvalidfolder(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesThemeFormatScheme", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "folder", r.Code, e)
}

/* ThemeApiServiceTests Read slide theme color scheme info.
   Test for ThemeApi.GetSlidesThemeFormatScheme method with invalid storage
*/
func TestGetSlidesThemeFormatSchemeInvalidstorage(t *testing.T) {
    request := createGetSlidesThemeFormatSchemeRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesThemeFormatScheme", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().ThemeApi.GetSlidesThemeFormatScheme(request)
    assertError(t, "GetSlidesThemeFormatScheme", "storage", r.Code, e)
}
