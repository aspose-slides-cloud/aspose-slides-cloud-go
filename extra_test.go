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
	"io/ioutil"
	"os"
	"testing"
)

/* 
   Test for SlidesApi.PostSlideSaveAs with timeout method
*/
func TestPipelineTemplate(t *testing.T) {
        file1 := NewRequestInputFile()
        file1.Index = 0

        file2 := NewRequestInputFile()
        file2.Index = 1

        input := NewInput()
        input.TemplateData = file1
        input.Template = file2

        output := NewOutputFile()

        task := NewSave()
        task.Format = "pptx"
        task.Output = output

        pipeline := NewPipeline()
        pipeline.Input = input
        pipeline.Tasks = []ITask { task }

	data1, e := ioutil.ReadFile("TestData/TemplatingCVDataWithBase64.xml")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data2, e := ioutil.ReadFile("TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
        files := [][]byte { data1, data2 }

	_, _, e = getTestApiClient().SlidesApi.Pipeline(pipeline, files)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

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

	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.Timeout = 1
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.DownloadSlide("test.pptx", 1, "svg", nil, nil, nil, "password", "TempSlidesSDK", "", "")
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

	r, _, e := getTestApiClient().SlidesApi.GetShape("test.pptx", 1, 1, "password", "TempSlidesSDK", "")
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
	chart := NewChart()
	if chart.getType() != "Chart" {
		t.Errorf("Unexpected chart type: %v.", chart.getType())
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
	_, e = c.SlidesApi.CopyFile("TempTests/" + fileName, folderName + "/" + fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        var dto1 Chart
        dto1.setType("Chart")
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
        dto1.setX(12)
        dto1.setY(14)
	_, _, e = c.SlidesApi.CreateShape(fileName, 1, &dto1, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e := c.SlidesApi.GetShape(fileName, 1, 4, password, folderName, "")
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
        dto2.setChartType("Line")
        var axis2 Axis
        axis2.setMinValue(min2)
        axes.setHorizontalAxis(&axis2)
        dto2.setAxes(&axes)
	_, _, e = c.SlidesApi.UpdateShape(fileName, 1, 4, &dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, 1, 4, password, folderName, "")
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
	_, _, e = c.SlidesApi.UpdateShape(fileName, 1, 4, &dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, 1, 4, password, folderName, "")
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

/* 
   Test for call with valid auth data
*/
func TestGoodAuth(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/* 
   Test for call with valid auth data
*/
func TestBadAuth(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.AppSid = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, r, e := testApiClient.SlidesApi.GetApiInfo()
	if e == nil {
		t.Errorf("Must have failed.")
		return
	}
	if r == nil {
		t.Errorf("Null response not expected.")
		return
	}
	statusCode := r.StatusCode
	if statusCode != 401 {
		t.Errorf("Unexpected error code: %v.", statusCode)
		return
	}
}

/* 
   Test for call with valid auth data
*/
func TestGoodAuthToken(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.AppSid = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/* 
   Test for call with valid auth data
*/
func TestBadAuthToken(t *testing.T) {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.OAuthToken = "invalid"
	testApiClient = NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
