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
	"encoding/json"
	"os"
	"testing"
)

/* 
   Test for SlidesApi.PostSlideSaveAs with timeout method
*/
func TestTimeout(t *testing.T) {
	return
	e := initializeTest("PostSlideSaveAs", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var request PostSlideSaveAsRequest
	request.Name = "test.ppt"
	request.SlideIndex = 1
	request.Format = "svg"
	request.Password = "password"
	request.Folder = "TempSlidesSDK"

	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.Timeout = 1
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.PostSlideSaveAs(request)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
 * Test for Shape
 */
func TestShape(t *testing.T) {
	e := initializeTest("GetSlideShape", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var request GetSlideShapeRequest
	request.Name = "test.ppt"
	request.SlideIndex = 1
	request.ShapeIndex = 1
	request.Password = "password"
	request.Folder = "TempSlidesSDK"

	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	r, _, e := testApiClient.SlidesApi.GetSlideShape(request)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IShape).getText() != "1" {
		t.Errorf("Error: Text expected to equal 1.")
		return
	}
}

/* 
   Test for SlidesApi.PostSlideSaveAs with timeout method
*/
func TestChart(t *testing.T) {
	return
	var chart Chart
	if chart.getType() != "Chart" {
		t.Errorf("Unexpected chart type: %v.", chart.getType())
		return
	}
	if chart.getShapeType() != "Chart" {
		t.Errorf("Unexpected chart shape type: %v.", chart.getType())
		return
	}
}
