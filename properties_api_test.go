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

/* PropertiesApiServiceTests Clean document properties.
   Test for PropertiesApi.DeleteSlidesDocumentProperties method
*/
func TestDeleteSlidesDocumentProperties(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    e := initializeTest("DeleteSlidesDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.DeleteSlidesDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Clean document properties.
   Test for PropertiesApi.DeleteSlidesDocumentProperties method with invalid name
*/
func TestDeleteSlidesDocumentPropertiesInvalidname(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "name", r.Code, e)
}

/* PropertiesApiServiceTests Clean document properties.
   Test for PropertiesApi.DeleteSlidesDocumentProperties method with invalid password
*/
func TestDeleteSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "password", r.Code, e)
}

/* PropertiesApiServiceTests Clean document properties.
   Test for PropertiesApi.DeleteSlidesDocumentProperties method with invalid folder
*/
func TestDeleteSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Clean document properties.
   Test for PropertiesApi.DeleteSlidesDocumentProperties method with invalid storage
*/
func TestDeleteSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperties(request)
    assertError(t, "DeleteSlidesDocumentProperties", "storage", r.Code, e)
}

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method
*/
func TestDeleteSlidesDocumentProperty(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    e := initializeTest("DeleteSlidesDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.DeleteSlidesDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method with invalid name
*/
func TestDeleteSlidesDocumentPropertyInvalidname(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "name", r.Code, e)
}

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method with invalid propertyName
*/
func TestDeleteSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "propertyName", r.Code, e)
}

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method with invalid password
*/
func TestDeleteSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "password", r.Code, e)
}

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method with invalid folder
*/
func TestDeleteSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Delete document property.
   Test for PropertiesApi.DeleteSlidesDocumentProperty method with invalid storage
*/
func TestDeleteSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createDeleteSlidesDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("DeleteSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.DeleteSlidesDocumentProperty(request)
    assertError(t, "DeleteSlidesDocumentProperty", "storage", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document properties.
   Test for PropertiesApi.GetSlidesDocumentProperties method
*/
func TestGetSlidesDocumentProperties(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    e := initializeTest("GetSlidesDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.GetSlidesDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Read presentation document properties.
   Test for PropertiesApi.GetSlidesDocumentProperties method with invalid name
*/
func TestGetSlidesDocumentPropertiesInvalidname(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "name", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document properties.
   Test for PropertiesApi.GetSlidesDocumentProperties method with invalid password
*/
func TestGetSlidesDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "password", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document properties.
   Test for PropertiesApi.GetSlidesDocumentProperties method with invalid folder
*/
func TestGetSlidesDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document properties.
   Test for PropertiesApi.GetSlidesDocumentProperties method with invalid storage
*/
func TestGetSlidesDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperties(request)
    assertError(t, "GetSlidesDocumentProperties", "storage", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method
*/
func TestGetSlidesDocumentProperty(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    e := initializeTest("GetSlidesDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.GetSlidesDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method with invalid name
*/
func TestGetSlidesDocumentPropertyInvalidname(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "name", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method with invalid propertyName
*/
func TestGetSlidesDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "propertyName", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method with invalid password
*/
func TestGetSlidesDocumentPropertyInvalidpassword(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "password", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method with invalid folder
*/
func TestGetSlidesDocumentPropertyInvalidfolder(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Read presentation document property.
   Test for PropertiesApi.GetSlidesDocumentProperty method with invalid storage
*/
func TestGetSlidesDocumentPropertyInvalidstorage(t *testing.T) {
    request := createGetSlidesDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("GetSlidesDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.GetSlidesDocumentProperty(request)
    assertError(t, "GetSlidesDocumentProperty", "storage", r.Code, e)
}

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method
*/
func TestPostSlidesSetDocumentProperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    e := initializeTest("PostSlidesSetDocumentProperties", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.PostSlidesSetDocumentProperties(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method with invalid name
*/
func TestPostSlidesSetDocumentPropertiesInvalidname(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "name", r.Code, e)
}

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method with invalid properties
*/
func TestPostSlidesSetDocumentPropertiesInvalidproperties(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.properties = invalidizeTestParamValue(request.properties, "properties", "DocumentProperties").(IDocumentProperties)

    e := initializeTest("PostSlidesSetDocumentProperties", "properties", request.properties)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "properties", r.Code, e)
}

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method with invalid password
*/
func TestPostSlidesSetDocumentPropertiesInvalidpassword(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "password", r.Code, e)
}

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method with invalid folder
*/
func TestPostSlidesSetDocumentPropertiesInvalidfolder(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Set document properties.
   Test for PropertiesApi.PostSlidesSetDocumentProperties method with invalid storage
*/
func TestPostSlidesSetDocumentPropertiesInvalidstorage(t *testing.T) {
    request := createPostSlidesSetDocumentPropertiesRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PostSlidesSetDocumentProperties", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PostSlidesSetDocumentProperties(request)
    assertError(t, "PostSlidesSetDocumentProperties", "storage", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method
*/
func TestPutSlidesSetDocumentProperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    e := initializeTest("PutSlidesSetDocumentProperty", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.PropertiesApi.PutSlidesSetDocumentProperty(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
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

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid name
*/
func TestPutSlidesSetDocumentPropertyInvalidname(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "name", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid propertyName
*/
func TestPutSlidesSetDocumentPropertyInvalidpropertyName(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.propertyName = invalidizeTestParamValue(request.propertyName, "propertyName", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "propertyName", request.propertyName)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "propertyName", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid property
*/
func TestPutSlidesSetDocumentPropertyInvalidproperty(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.property = invalidizeTestParamValue(request.property, "property", "DocumentProperty").(IDocumentProperty)

    e := initializeTest("PutSlidesSetDocumentProperty", "property", request.property)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "property", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid password
*/
func TestPutSlidesSetDocumentPropertyInvalidpassword(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "password", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid folder
*/
func TestPutSlidesSetDocumentPropertyInvalidfolder(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "folder", r.Code, e)
}

/* PropertiesApiServiceTests Set document property.
   Test for PropertiesApi.PutSlidesSetDocumentProperty method with invalid storage
*/
func TestPutSlidesSetDocumentPropertyInvalidstorage(t *testing.T) {
    request := createPutSlidesSetDocumentPropertyRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)

    e := initializeTest("PutSlidesSetDocumentProperty", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().PropertiesApi.PutSlidesSetDocumentProperty(request)
    assertError(t, "PutSlidesSetDocumentProperty", "storage", r.Code, e)
}
