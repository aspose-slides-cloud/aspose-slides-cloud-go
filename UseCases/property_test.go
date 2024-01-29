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

package usecasetests

import (
	"io/ioutil"
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v24"
)

/*
   Test for builtin property
*/
func TestPropertyBuiltin(t *testing.T) {
	propertyName := "Author"
	updatedPropertyValue := "New Value"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}

	property := slidescloud.NewDocumentProperty()
	property.Value = updatedPropertyValue
	result, _, e = c.SlidesApi.SetDocumentProperty(fileName, propertyName, property, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}

	_, _, e = c.SlidesApi.DeleteDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e = c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() == updatedPropertyValue {
		t.Errorf("Wrong property value. Expected not %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if !result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected true.")
		return
	}
}

/*
   Test for custom property
*/
func TestPropertyCustom(t *testing.T) {
	propertyName := "CustomProperty2"
	updatedPropertyValue := "New Value"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	property := slidescloud.NewDocumentProperty()
	property.Value = updatedPropertyValue
	result, _, e := c.SlidesApi.SetDocumentProperty(fileName, propertyName, property, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetName() != propertyName {
		t.Errorf("Wrong property name. Expected %v but was %v.", propertyName, result.GetName())
		return
	}
	if result.GetValue() != updatedPropertyValue {
		t.Errorf("Wrong property value. Expected %v but was %v.", updatedPropertyValue, result.GetValue())
		return
	}
	if result.GetBuiltIn() {
		t.Errorf("Wrong BuiltIn value. Expected false.")
		return
	}

	_, _, e = c.SlidesApi.DeleteDocumentProperty(fileName, propertyName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, response, e := c.SlidesApi.GetDocumentProperty(fileName, propertyName, password, folderName, "")
	if e == nil {
		t.Errorf("The property must have been deleted")
		return
	}
	if response.StatusCode != 404 {
		t.Errorf("Wrong status code. Expected %v but was %v.", 404, response.StatusCode)
		return
	}
}

/*
   Test for properties bulk update
*/
func TestPropertyBulkUpdate(t *testing.T) {
	propertyName := "Author"
	customPropertyName := "CustomProperty2"
	updatedPropertyValue := "New Value"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.GetDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	count := len(result.GetList())

	property1 := slidescloud.NewDocumentProperty()
	property1.Name = propertyName
	property1.Value = updatedPropertyValue
	property2 := slidescloud.NewDocumentProperty()
	property2.Name = customPropertyName
	property2.Value = updatedPropertyValue
	properties := slidescloud.NewDocumentProperties()
	properties.List = []slidescloud.IDocumentProperty{property1, property2}
	result, _, e = c.SlidesApi.SetDocumentProperties(fileName, properties, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetList()) != count+1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count+1, len(result.GetList()))
		return
	}

	result, _, e = c.SlidesApi.DeleteDocumentProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetList()) != count-1 {
		t.Errorf("Wrong property count. Expected %v but was %v.", count-1, len(result.GetList()))
		return
	}
}

/*
   Test for slide properties
*/
func TestPropertySlideProperties(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	GetResult, _, e := c.SlidesApi.GetSlideProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideProperties()
	dto.FirstSlideNumber = GetResult.GetFirstSlideNumber() + 2
	putResult, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.GetOrientation() != GetResult.GetOrientation() {
		t.Errorf("Wrong orientation. Expected %v but was %v.", GetResult.GetOrientation(), putResult.GetOrientation())
		return
	}
	if putResult.GetFirstSlideNumber() == GetResult.GetFirstSlideNumber() {
		t.Errorf("Wrong FirstSlideNumber. Expected not %v but was %v.", GetResult.GetFirstSlideNumber(), putResult.GetFirstSlideNumber())
		return
	}
}

/*
   Test for preset slide size
*/
func TestPropertySlideSizePreset(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideProperties()
	dto.SizeType = "B4IsoPaper"
	result, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetSizeType() != "B4IsoPaper" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "B4IsoPaper", result.GetSizeType())
		return
	}
	if result.GetWidth() != 852 {
		t.Errorf("Wrong width. Expected %v but was %v.", 852, result.GetWidth())
		return
	}
	if result.GetHeight() != 639 {
		t.Errorf("Wrong height. Expected %v but was %v.", 639, result.GetHeight())
		return
	}
}

/*
   Test for custom slide size
*/
func TestPropertySlideSizeCustom(t *testing.T) {
	var width int32 = 800
	var height int32 = 500
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideProperties()
	dto.Width = width
	dto.Height = height
	result, _, e := c.SlidesApi.SetSlideProperties(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetSizeType() != "Custom" {
		t.Errorf("Wrong size type. Expected %v but was %v.", "Custom", result.GetSizeType())
		return
	}
	if result.GetWidth() != width {
		t.Errorf("Wrong width. Expected %v but was %v.", width, result.GetWidth())
		return
	}
	if result.GetHeight() != height {
		t.Errorf("Wrong height. Expected %v but was %v.", height, result.GetHeight())
		return
	}
}

/*
   Test for protection properties
*/
func TestPropertyProtection(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	GetResult, _, e := c.SlidesApi.GetProtectionProperties(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewProtectionProperties()
	dto.EncryptDocumentProperties = GetResult.GetEncryptDocumentProperties()
	dto.ReadOnlyRecommended = !GetResult.GetReadOnlyRecommended()
	putResult, _, e := c.SlidesApi.SetProtection(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if putResult.GetEncryptDocumentProperties() != GetResult.GetEncryptDocumentProperties() {
		t.Errorf("Wrong EncryptDocumentProperties. Expected %v but was %v.", GetResult.GetEncryptDocumentProperties(), putResult.GetEncryptDocumentProperties())
		return
	}
	if putResult.GetReadOnlyRecommended() == GetResult.GetReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended. Expected not %v but was %v.", GetResult.GetReadOnlyRecommended(), putResult.GetReadOnlyRecommended())
		return
	}
}

/*
   Test for delete protection
*/
func TestPropertyProtectionDelete(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteProtection(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetIsEncrypted() {
		t.Errorf("Wrong IsEncrypted value. Expected false.")
		return
	}
	if result.GetReadOnlyRecommended() {
		t.Errorf("Wrong ReadOnlyRecommended value. Expected false.")
		return
	}
	if result.GetReadPassword() != "" {
		t.Errorf("Wrong ReadPassword value. Expected empty string.")
		return
	}
}

/*
   Test for online protection properties
*/
func TestPropertyProtectionOnline(t *testing.T) {
	source, e := ioutil.ReadFile("../TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	dto := slidescloud.NewProtectionProperties()
	dto.ReadPassword = "newPassword"
	result, _, e := slidescloud.GetTestSlidesApiClient().SlidesApi.SetProtectionOnline(source, dto, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultStat.Size() == int64(len(source)) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), resultStat.Size())
		return
	}
}

/*
   Test for online delete protection
*/
func TestPropertyProtectionUnprotectOnline(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e := slidescloud.GetTestSlidesApiClient().SlidesApi.DeleteProtectionOnline(source, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if resultStat.Size() == int64(len(source)) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), resultStat.Size())
		return
	}
}

/*
   Test for view properties
*/
func TestViewPropertiesGet(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetViewProperties(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetShowComments() != "True" {
		t.Errorf("Expected %v, but was %v", "True", response.GetShowComments())
		return
	}
}

/*
   Test for view properties
*/
func TestViewPropertiesSet(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewViewProperties()
	dto.SetShowComments("False")
	slideViewProperties := slidescloud.NewCommonSlideViewProperties()
	slideViewProperties.SetScale(50)
	dto.SetSlideViewProperties(slideViewProperties)

	response, _, e := c.SlidesApi.SetViewProperties(fileName, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetShowComments() != "False" {
		t.Errorf("Expected %v, but was %v", "False", response.GetShowComments())
		return
	}

	if response.GetSlideViewProperties().GetScale() != 50 {
		t.Errorf("Expected %v, but was %v", 50, response.GetSlideViewProperties().GetScale())
		return
	}
}

/*
   Test for protection check
*/
func TestProtectionCheck(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	protectionProperties, _, e := c.SlidesApi.GetProtectionProperties(fileName, "", folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if protectionProperties.GetIsEncrypted() != true {
		t.Errorf("Expected %v, but was %v", true, protectionProperties.GetIsEncrypted())
		return
	}

	if protectionProperties.GetReadPassword() != "" {
		t.Errorf("Expected %v, but was %v", "", protectionProperties.GetReadPassword())
		return
	}

	protectionProperties, _, e = c.SlidesApi.GetProtectionProperties(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if protectionProperties.GetIsEncrypted() != true {
		t.Errorf("Expected %v, but was %v", true, protectionProperties.GetIsEncrypted())
		return
	}

	if protectionProperties.GetReadPassword() == "" {
		t.Errorf("Expected %v, but was %v", "not null", protectionProperties.GetReadPassword())
		return
	}
}

/*
   Test for slide show properties
*/
func TestSlideShowPropertiesGet(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.GetSlideShowProperties(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetShowAnimation() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetShowAnimation())
		return
	}

	if response.GetShowNarration() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetShowAnimation())
		return
	}
}

/*
   Test for slide show properties
*/
func TestSlideShowPropertiesSet(t *testing.T) {
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewSlideShowProperties()
	dto.SetLoop(true)
	dto.SetUseTimings(true)
	dto.SetSlideShowType("PresentedBySpeaker")

	response, _, e := c.SlidesApi.SetSlideShowProperties(fileName, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetLoop() != dto.GetLoop() {
		t.Errorf("Expected %v, but was %v", dto.GetLoop(), response.GetLoop())
		return
	}

	if response.GetUseTimings() != dto.GetUseTimings() {
		t.Errorf("Expected %v, but was %v", dto.GetUseTimings(), response.GetUseTimings())
		return
	}

	if response.GetSlideShowType() != dto.GetSlideShowType() {
		t.Errorf("Expected %v, but was %v", dto.GetSlideShowType(), response.GetSlideShowType())
		return
	}
}
