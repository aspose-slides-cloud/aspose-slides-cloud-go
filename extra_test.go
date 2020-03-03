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

/*
 * Test for nullable properties
 */
func TestNullableProperties(t *testing.T) {
        var folderName = "TempSlidesSDK"
        var fileName = "placeholders.pptx"
        var password = "password"
        var min1  = 44.3
        var min2 = 12.0
        var max1 = 104.3
        var max2 = 87.0

	e := initializeTest("NoFunction", "No method", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := getTestApiClient()
	var request CopyFileRequest
	request.SrcPath = "TempTests/" + fileName
	request.DestPath = folderName + "/" + fileName
	_, e = c.SlidesApi.CopyFile(request)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        var dto1 Chart
        dto1.setType("Chart")
        dto1.setShapeType("Chart")
        dto1.setChartType("Line")
        dto1.setWidth(400)
        dto1.setHeight(300)
        var title ChartTitle
        title.setHasTitle(true)
        title.setText("MyTitle")
        dto1.setTitle(&title)
        var series OneValueSeries
        series.setType("ClusteredColumn")
        series.setDataPointType("OneValue")
        series.setName("Series1")
        var point1 OneValueChartDataPoint
        point1.setValue(40)
        var point2 OneValueChartDataPoint
        point2.setValue(50)
        points := make([]IOneValueChartDataPoint, 2)
        points[0] = &point1
        points[1] = &point2
        series.setDataPoints(points)
        serieses := make([]ISeries, 1)
        serieses[0] = &series
        dto1.setSeries(serieses)
        var axes Axes
        var axis1 Axis
        axis1.setIsAutomaticMinValue(false)
        axis1.setMinValue(min1)
        axis1.setIsAutomaticMaxValue(false)
        axis1.setMaxValue(max1)
        axes.setHorizontalAxis(&axis1)
        dto1.setAxes(&axes)
	var postRequest PostAddNewShapeRequest
	postRequest.Name = fileName
	postRequest.SlideIndex = 1
	postRequest.Password = password
	postRequest.Folder = folderName
        dto1.setX(12)
        dto1.setY(14)
	postRequest.Dto = &dto1
	_, _, e = c.SlidesApi.PostAddNewShape(postRequest)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var getRequest GetSlideShapeRequest
	getRequest.Name = fileName
	getRequest.SlideIndex = 1
	getRequest.ShapeIndex = 4
	getRequest.Password = password
	getRequest.Folder = folderName
	r, _, e := c.SlidesApi.GetSlideShape(getRequest)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min1 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

        var dto2 Chart
        dto2.setType("Chart")
        dto2.setShapeType("Chart")
        dto2.setChartType("Line")
        var axis2 Axis
        axis2.setMinValue(min2)
        axes.setHorizontalAxis(&axis2)
        dto2.setAxes(&axes)
	var putRequest PutSlideShapeInfoRequest
	putRequest.Name = fileName
	putRequest.SlideIndex = 1
	putRequest.ShapeIndex = 4
	putRequest.Password = password
	putRequest.Folder = folderName
	putRequest.Dto = &dto2
	_, _, e = c.SlidesApi.PutSlideShapeInfo(putRequest)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetSlideShape(getRequest)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

        var axis3 Axis
        axis3.setMaxValue(max2)
        axes.setHorizontalAxis(&axis3)
	_, _, e = c.SlidesApi.PutSlideShapeInfo(putRequest)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetSlideShape(getRequest)
	if e != nil {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(IChart).getAxes().getHorizontalAxis().getMaxValue() != max2 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}
}
