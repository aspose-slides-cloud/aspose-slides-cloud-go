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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v23"
)

/*
   Test for Get chart
*/
func TestChartGet(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, 3, 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(chart.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(chart.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(chart.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(chart.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart
*/
func TestChartCreateAutoDataSource(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300
	series1 := slidescloud.NewOneValueSeries()
	series1.Name = "Series1"
	point11 := slidescloud.NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := slidescloud.NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := slidescloud.NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point11, point12, point13}
	series2 := slidescloud.NewOneValueSeries()
	series2.Name = "Series2"
	point21 := slidescloud.NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := slidescloud.NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := slidescloud.NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []slidescloud.IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []slidescloud.ISeries{series1, series2}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Category1"
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Category2"
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart with data in workbook
*/
func TestChartCreateWorksheet(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300

	dataSourceForCategories := slidescloud.NewWorkbook()
	dataSourceForCategories.WorksheetIndex = 0
	dataSourceForCategories.ColumnIndex = 0
	dataSourceForCategories.RowIndex = 1
	chart.DataSourceForCategories = dataSourceForCategories

	series1 := slidescloud.NewOneValueSeries()
	dataSourceForSeries1Name := slidescloud.NewWorkbook()
	dataSourceForSeries1Name.WorksheetIndex = 0
	dataSourceForSeries1Name.ColumnIndex = 1
	dataSourceForSeries1Name.RowIndex = 0
	series1.DataSourceForSeriesName = dataSourceForSeries1Name
	series1.Name = "Series1"

	dataSourceForSeries1Values := slidescloud.NewWorkbook()
	dataSourceForSeries1Values.WorksheetIndex = 0
	dataSourceForSeries1Values.ColumnIndex = 1
	dataSourceForSeries1Values.RowIndex = 1
	series1.DataSourceForValues = dataSourceForSeries1Values
	point11 := slidescloud.NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := slidescloud.NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := slidescloud.NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point11, point12, point13}

	series2 := slidescloud.NewOneValueSeries()
	dataSourceForSeries2Name := slidescloud.NewWorkbook()
	dataSourceForSeries2Name.WorksheetIndex = 0
	dataSourceForSeries2Name.ColumnIndex = 2
	dataSourceForSeries2Name.RowIndex = 0
	series2.DataSourceForSeriesName = dataSourceForSeries2Name
	series2.Name = "Series2"

	dataSourceForSeries2Values := slidescloud.NewWorkbook()
	dataSourceForSeries2Values.WorksheetIndex = 0
	dataSourceForSeries2Values.ColumnIndex = 2
	dataSourceForSeries2Values.RowIndex = 1
	series2.DataSourceForValues = dataSourceForSeries2Values

	point21 := slidescloud.NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := slidescloud.NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := slidescloud.NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []slidescloud.IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []slidescloud.ISeries{series1, series2}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Category1"
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Category2"
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart with literals
*/
func TestChartCreateLiterals(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300

	dataSourceForCategories := slidescloud.NewLiterals()
	chart.DataSourceForCategories = dataSourceForCategories

	series1 := slidescloud.NewOneValueSeries()
	dataSourceForSeries1Name := slidescloud.NewLiterals()
	series1.DataSourceForSeriesName = dataSourceForSeries1Name
	series1.Name = "Series1"

	dataSourceForSeries1Values := slidescloud.NewLiterals()
	series1.DataSourceForValues = dataSourceForSeries1Values
	point11 := slidescloud.NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := slidescloud.NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := slidescloud.NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point11, point12, point13}

	series2 := slidescloud.NewOneValueSeries()
	dataSourceForSeries2Name := slidescloud.NewLiterals()
	series2.DataSourceForSeriesName = dataSourceForSeries2Name
	series2.Name = "Series2"

	dataSourceForSeries2Values := slidescloud.NewLiterals()
	series2.DataSourceForValues = dataSourceForSeries2Values

	point21 := slidescloud.NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := slidescloud.NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := slidescloud.NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []slidescloud.IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []slidescloud.ISeries{series1, series2}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Category1"
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Category2"
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for update chart
*/
func TestChartUpdate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300
	series1 := slidescloud.NewOneValueSeries()
	series1.Name = "Series1"
	point11 := slidescloud.NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := slidescloud.NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := slidescloud.NewOneValueChartDataPoint()
	point13.Value = 70
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point11, point12, point13}
	series2 := slidescloud.NewOneValueSeries()
	series2.Name = "Series2"
	point21 := slidescloud.NewOneValueChartDataPoint()
	point21.Value = 55
	point22 := slidescloud.NewOneValueChartDataPoint()
	point22.Value = 35
	point23 := slidescloud.NewOneValueChartDataPoint()
	point23.Value = 90
	series2.DataPoints = []slidescloud.IOneValueChartDataPoint{point21, point22, point23}
	chart.Series = []slidescloud.ISeries{series1, series2}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Category1"
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Category2"
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.UpdateShape(fileName, 3, 1, chart, "password", folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart series
*/
func TestChartSeriesCreate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	series := slidescloud.NewOneValueSeries()
	series.Name = "Series1"
	point1 := slidescloud.NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := slidescloud.NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := slidescloud.NewOneValueChartDataPoint()
	point3.Value = 14
	point4 := slidescloud.NewOneValueChartDataPoint()
	point4.Value = 70
	series.DataPoints = []slidescloud.IOneValueChartDataPoint{point1, point2, point3, point4}
	result, _, e := c.SlidesApi.CreateChartSeries(fileName, 3, 1, series, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 4 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for update chart series
*/
func TestChartSeriesUpdate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	series := slidescloud.NewOneValueSeries()
	series.Name = "Series1"
	point1 := slidescloud.NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := slidescloud.NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := slidescloud.NewOneValueChartDataPoint()
	point3.Value = 14
	point4 := slidescloud.NewOneValueChartDataPoint()
	point4.Value = 70
	series.DataPoints = []slidescloud.IOneValueChartDataPoint{point1, point2, point3, point4}
	result, _, e := c.SlidesApi.UpdateChartSeries(fileName, 3, 1, 2, series, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for delete chart series
*/
func TestChartSeriesDelete(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartSeries(fileName, 3, 1, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 2 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 2, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for create chart category
*/
func TestChartCategoryCreate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	category := slidescloud.NewChartCategory()
	category.Value = "NewCategory"
	point1 := slidescloud.NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := slidescloud.NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := slidescloud.NewOneValueChartDataPoint()
	point3.Value = 14
	category.DataPoints = []slidescloud.IOneValueChartDataPoint{point1, point2, point3}
	result, _, e := c.SlidesApi.CreateChartCategory(fileName, 3, 1, category, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 5 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 5, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()) != 5 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 5, len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()[4].GetValue() != category.DataPoints[0].GetValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].GetValue(), result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()[4].GetValue())
		return
	}
}

/*
   Test for update chart category
*/
func TestChartCategoryUpdate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	category := slidescloud.NewChartCategory()
	category.Value = "NewCategory"
	point1 := slidescloud.NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := slidescloud.NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := slidescloud.NewOneValueChartDataPoint()
	point3.Value = 14
	category.DataPoints = []slidescloud.IOneValueChartDataPoint{point1, point2, point3}
	result, _, e := c.SlidesApi.UpdateChartCategory(fileName, 3, 1, 2, category, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()[1].GetValue() != category.DataPoints[0].GetValue() {
		t.Errorf("Wrong data point value. Expected %v but was %v.", category.DataPoints[0].GetValue(), result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()[1].GetValue())
		return
	}
}

/*
   Test for delete chart category
*/
func TestChartCategoryDelete(t *testing.T) {

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartCategory(fileName, 3, 1, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 3 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()) != 3 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()))
		return
	}
}

/*
   Test for create chart data point
*/
func TestChartDataPointCreate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	point := slidescloud.NewOneValueChartDataPoint()
	point.Value = 40
	_, response, e := c.SlidesApi.CreateChartDataPoint(fileName, 3, 1, 2, point, password, folderName, "")
	if e == nil {
		t.Errorf("Must have failed because adding data points only works with Scatter & Bubble charts")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for update chart data point
*/
func TestChartDataPointUpdate(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	point := slidescloud.NewOneValueChartDataPoint()
	point.Value = 40
	result, _, e := c.SlidesApi.UpdateChartDataPoint(fileName, 3, 1, 2, 2, point, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()[1].(slidescloud.IOneValueSeries).GetDataPoints()) != 4 {
		t.Errorf("Wrong data point count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetSeries()[1].(slidescloud.IOneValueSeries).GetDataPoints()))
		return
	}
	if result.(slidescloud.IChart).GetSeries()[1].(slidescloud.IOneValueSeries).GetDataPoints()[1].GetValue() != point.Value {
		t.Errorf("Wrong data point value. Expected %v but was %v.", point.Value, result.(slidescloud.IChart).GetSeries()[1].(slidescloud.IOneValueSeries).GetDataPoints()[1].GetValue())
		return
	}
}

/*
   Test for delete chart data point
*/
func TestChartDataPointDelete(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteChartDataPoint(fileName, 3, 1, 2, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 3 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 3, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
	if result.(slidescloud.IChart).GetSeries()[1].(slidescloud.IOneValueSeries).GetDataPoints()[1] != nil {
		t.Errorf("Data point should be nil.")
		return
	}
}

/*
   Test for sunburst chart
*/
func TestChartSunburst(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "Sunburst"
	chart.Width = 400
	chart.Height = 300
	series1 := slidescloud.NewOneValueSeries()
	series1.Name = "Series1"
	point1 := slidescloud.NewOneValueChartDataPoint()
	point1.Value = 40
	point2 := slidescloud.NewOneValueChartDataPoint()
	point2.Value = 50
	point3 := slidescloud.NewOneValueChartDataPoint()
	point3.Value = 70
	point4 := slidescloud.NewOneValueChartDataPoint()
	point4.Value = 60
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point1, point2, point3, point4}
	chart.Series = []slidescloud.ISeries{series1}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Leaf1"
	category1.Level = 3
	category1.ParentCategories = []string{"Branch1", "Stem1"}
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Leaf2"
	category2.Level = 3
	category2.ParentCategories = []string{"Branch1", "Stem1"}
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Branch2"
	category3.Level = 2
	category3.ParentCategories = []string{"Stem1"}
	category4 := slidescloud.NewChartCategory()
	category4.Value = "Stem2"
	category4.Level = 1
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3, category4}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.(slidescloud.IChart).GetSeries()) != 1 {
		t.Errorf("Wrong series count. Expected %v but was %v.", 1, len(result.(slidescloud.IChart).GetSeries()))
		return
	}
	if len(result.(slidescloud.IChart).GetCategories()) != 4 {
		t.Errorf("Wrong category count. Expected %v but was %v.", 4, len(result.(slidescloud.IChart).GetCategories()))
		return
	}
}

/*
   Test for Multi level category axis
*/
func TestMultiLevelCategoryAxis(t *testing.T) {
	var slideIndex int32 = 3

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewChart()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(500)
	dto.SetHeight(400)
	dto.SetChartType("ClusteredColumn")

	series1 := slidescloud.NewOneValueSeries()
	series1.SetType("ClusteredColumn")
	dataPoint1 := slidescloud.NewOneValueChartDataPoint()
	dataPoint1.SetValue(1)
	dataPoint2 := slidescloud.NewOneValueChartDataPoint()
	dataPoint2.SetValue(2)
	dataPoint3 := slidescloud.NewOneValueChartDataPoint()
	dataPoint3.SetValue(3)
	dataPoint4 := slidescloud.NewOneValueChartDataPoint()
	dataPoint4.SetValue(4)
	dataPoint5 := slidescloud.NewOneValueChartDataPoint()
	dataPoint5.SetValue(5)
	dataPoint6 := slidescloud.NewOneValueChartDataPoint()
	dataPoint6.SetValue(6)
	dataPoint7 := slidescloud.NewOneValueChartDataPoint()
	dataPoint7.SetValue(7)
	dataPoint8 := slidescloud.NewOneValueChartDataPoint()
	dataPoint8.SetValue(8)
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{dataPoint1, dataPoint2, dataPoint3, dataPoint4, dataPoint5, dataPoint6, dataPoint7, dataPoint8}
	dto.Series = []slidescloud.ISeries{series1}

	category1 := slidescloud.NewChartCategory()
	category1.SetValue("Category1")
	category1.SetParentCategories([]string{"Sub-category 1", "Root 1"})
	category2 := slidescloud.NewChartCategory()
	category2.SetValue("Category2")
	category3 := slidescloud.NewChartCategory()
	category3.SetValue("Category3")
	category3.SetParentCategories([]string{"Sub-category 2"})
	category4 := slidescloud.NewChartCategory()
	category4.SetValue("Category4")
	category5 := slidescloud.NewChartCategory()
	category5.SetValue("Category5")
	category5.SetParentCategories([]string{"Sub-category 3", "Root 2"})
	category6 := slidescloud.NewChartCategory()
	category6.SetValue("Category6")
	category7 := slidescloud.NewChartCategory()
	category7.SetValue("Category7")
	category7.SetParentCategories([]string{"Sub-category 4"})
	category8 := slidescloud.NewChartCategory()
	category8.SetValue("Category8")
	dto.Categories = []slidescloud.IChartCategory{category1, category2, category3, category4, category5, category6, category7, category8}

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.(slidescloud.IChart).GetChartType() != dto.GetChartType() {
		t.Errorf("Expected %v, but was %v", dto.GetChartType(), response.(slidescloud.IChart).GetChartType())
		return
	}

	if len(response.(slidescloud.IChart).GetSeries()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, response)
		return
	}

	if len(response.(slidescloud.IChart).GetCategories()) != 8 {
		t.Errorf("Expected %v, but was %v", 8, response)
		return
	}
}

/*
   Test for Hide chart legend
*/
func TestTemplate(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart.(slidescloud.IChart).GetLegend().SetHasLegend(false)

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, chart, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.(slidescloud.IChart).GetLegend().GetHasLegend() != false {
		t.Errorf("Expected %v, but was %v", false, true)
		return
	}
}

/*
   Test for chart grid lines format
*/
func TestChartGridLinesFormat(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	horizontalAxis := slidescloud.NewAxis()
	horizontalAxis.MajorGridLinesFormat = slidescloud.NewChartLinesFormat()
	horzMajLineFormat := slidescloud.NewLineFormat()
	horzMajLineFormat.FillFormat = slidescloud.NewNoFill()
	horizontalAxis.MajorGridLinesFormat.SetLineFormat(horzMajLineFormat)

	horizontalAxis.MinorGridLinesFormat = slidescloud.NewChartLinesFormat()
	horzMinorLineFormat := slidescloud.NewLineFormat()
	horzMinorLineFormat.FillFormat = slidescloud.NewSolidFill()
	horzMinorLineFormat.FillFormat.(slidescloud.ISolidFill).SetColor("Black")
	horizontalAxis.MinorGridLinesFormat.SetLineFormat(horzMinorLineFormat)

	verticalAxis := slidescloud.NewAxis()
	verticalAxis.MajorGridLinesFormat = slidescloud.NewChartLinesFormat()
	vertMajorLineFormat := slidescloud.NewLineFormat()
	gradientFill := slidescloud.NewGradientFill()
	gradientFill.Direction = "FromCorner1"
	stop1 := slidescloud.NewGradientFillStop()
	stop1.Color = "White"
	stop1.Position = 0
	stop2 := slidescloud.NewGradientFillStop()
	stop2.Color = "Black"
	stop2.Position = 1
	gradientFill.SetStops([]slidescloud.IGradientFillStop{stop1, stop2})
	vertMajorLineFormat.FillFormat = gradientFill
	verticalAxis.MajorGridLinesFormat.SetLineFormat(vertMajorLineFormat)

	verticalAxis.MinorGridLinesFormat = slidescloud.NewChartLinesFormat()
	vertMinorLineFormat := slidescloud.NewLineFormat()
	vertMinorLineFormat.FillFormat = slidescloud.NewNoFill()
	verticalAxis.MinorGridLinesFormat.SetLineFormat(vertMinorLineFormat)

	axes := slidescloud.NewAxes()
	axes.SetHorizontalAxis(horizontalAxis)
	axes.SetVerticalAxis(verticalAxis)

	chart.(slidescloud.IChart).SetAxes(axes)

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, chart, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	resonseFill := response.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMajorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "NoFill" {
		t.Errorf("Expected %v, but was %v", "NoFill", resonseFill)
		return
	}

	resonseFill = response.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMinorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", resonseFill)
		return
	}

	resonseFill = response.(slidescloud.IChart).GetAxes().GetVerticalAxis().GetMajorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "Gradient" {
		t.Errorf("Expected %v, but was %v", "NoFill", resonseFill)
		return
	}

	resonseFill = response.(slidescloud.IChart).GetAxes().GetVerticalAxis().GetMinorGridLinesFormat().GetLineFormat().GetFillFormat().GetType()
	if resonseFill != "NoFill" {
		t.Errorf("Expected %v, but was %v", "Solid", resonseFill)
		return
	}
}

/*
   Test for chart series groups
*/
func TestChartSeriesGroups(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1
	var seriesGroupIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(chart.(slidescloud.IChart).GetSeriesGroups()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(chart.(slidescloud.IChart).GetSeriesGroups()))
		return
	}

	seriesGroup := chart.(slidescloud.IChart).GetSeriesGroups()[0]
	seriesGroup.SetOverlap(10)

	chart, _, e = c.SlidesApi.SetChartSeriesGroup(fileName, slideIndex, shapeIndex, seriesGroupIndex, seriesGroup, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	seriesGroup = chart.(slidescloud.IChart).GetSeriesGroups()[0]

	if seriesGroup.GetOverlap() != 10 {
		t.Errorf("Expected %v, but was %v", 10, seriesGroup.GetOverlap())
		return
	}
}

/*
   Test for chart legend
*/
func TestChartLegend(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	legendDto := slidescloud.NewLegend()
	legendDto.Overlay = true
	solidFill := slidescloud.NewSolidFill()
	solidFill.Color = "#77CEF9"
	legendDto.FillFormat = solidFill

	response, _, e := c.SlidesApi.SetChartLegend(fileName, slideIndex, shapeIndex, legendDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetOverlay() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetOverlay())
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat())
		return
	}
}

/*
   Test for setchart axis
*/
func TestChartAxisSet(t *testing.T) {
	var slideIndex int32 = 3
	var shapeIndex int32 = 1

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	axisDto := slidescloud.NewAxis()
	axisDto.HasTitle = true
	axisDto.IsAutomaticMaxValue = false
	axisDto.MaxValue = 10

	response, _, e := c.SlidesApi.SetChartAxis(fileName, slideIndex, shapeIndex, "VerticalAxis", axisDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetHasTitle() != true {
		t.Errorf("Expected %v, but was %v", true, response.GetHasTitle())
		return
	}

	if response.GetIsAutomaticMaxValue() != false {
		t.Errorf("Expected %v, but was %v", false, response.GetIsAutomaticMaxValue())
		return
	}

	if response.GetMaxValue() != 10 {
		t.Errorf("Expected %v, but was %v", 10, response.GetMaxValue())
		return
	}
}

/*
   Test for chart walls
*/
func TestChartWallSet(t *testing.T) {
	var slideIndex int32 = 8
	var shapeIndex int32 = 2

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	wallDto := slidescloud.NewChartWall()
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.SetColor("#77CEF9")
	wallDto.SetFillFormat(fillFormat)

	response, _, e := c.SlidesApi.SetChartWall(fileName, slideIndex, shapeIndex, "BackWall", wallDto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for data point format
*/
func TestUpdateDataPointFormat(t *testing.T) {
	var slideIndex int32 = 8
	var shapeIndex int32 = 2
	var seriesIndex int32 = 2
	var dataPointIndex int32 = 2

	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	color := "#77CEF9"

	dto := slidescloud.NewOneValueChartDataPoint()
	dto.SetValue(40)
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.SetColor(color)
	lineFormat := slidescloud.NewLineFormat()
	lineFormat.FillFormat = slidescloud.NewSolidFill()
	lineFormat.FillFormat.(slidescloud.ISolidFill).SetColor(color)
	effectFormat := slidescloud.NewEffectFormat()
	effectFormat.Blur = slidescloud.NewBlurEffect()
	effectFormat.Blur.SetGrow(true)
	effectFormat.Blur.SetRadius(5)
	dto.SetFillFormat(fillFormat)
	dto.SetLineFormat(lineFormat)
	dto.SetEffectFormat(effectFormat)

	response, _, e := c.SlidesApi.UpdateChartDataPoint(fileName, slideIndex, shapeIndex, seriesIndex,
		dataPointIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dataPoint := response.GetSeries()[seriesIndex-1].(slidescloud.IOneValueSeries).GetDataPoints()[dataPointIndex-1]

	if dataPoint.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", dataPoint.GetFillFormat().GetType())
		return
	}

	if dataPoint.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", dataPoint.GetFillFormat().GetType())
		return
	}

	if dataPoint.GetLineFormat().GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", dataPoint.GetLineFormat().GetFillFormat().GetType())
		return
	}

	if dataPoint.GetEffectFormat().GetBlur() == nil {
		t.Errorf("Expected %v, but was %v", "not nil", nil)
		return
	}
}

/*
   Test for create chart with data in workbook
*/
func TestChartFormulas(t *testing.T) {
	c := slidescloud.GetTestApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	chart := slidescloud.NewChart()
	chart.ChartType = "ClusteredColumn"
	chart.Width = 400
	chart.Height = 300

	dataSourceForCategories := slidescloud.NewWorkbook()
	dataSourceForCategories.WorksheetIndex = 0
	dataSourceForCategories.ColumnIndex = 0
	dataSourceForCategories.RowIndex = 1
	chart.DataSourceForCategories = dataSourceForCategories

	series1 := slidescloud.NewOneValueSeries()
	dataSourceForSeries1Name := slidescloud.NewWorkbook()
	dataSourceForSeries1Name.WorksheetIndex = 0
	dataSourceForSeries1Name.ColumnIndex = 1
	dataSourceForSeries1Name.RowIndex = 0
	series1.DataSourceForSeriesName = dataSourceForSeries1Name
	series1.Name = "Series1"

	dataSourceForSeries1Values := slidescloud.NewWorkbook()
	dataSourceForSeries1Values.WorksheetIndex = 0
	dataSourceForSeries1Values.ColumnIndex = 1
	dataSourceForSeries1Values.RowIndex = 1
	series1.DataSourceForValues = dataSourceForSeries1Values
	point11 := slidescloud.NewOneValueChartDataPoint()
	point11.Value = 40
	point12 := slidescloud.NewOneValueChartDataPoint()
	point12.Value = 50
	point13 := slidescloud.NewOneValueChartDataPoint()
	point13.ValueFormula = "SUM(B2:B3)"
	series1.DataPoints = []slidescloud.IOneValueChartDataPoint{point11, point12, point13}

	chart.Series = []slidescloud.ISeries{series1}
	category1 := slidescloud.NewChartCategory()
	category1.Value = "Category1"
	category2 := slidescloud.NewChartCategory()
	category2.Value = "Category2"
	category3 := slidescloud.NewChartCategory()
	category3.Value = "Category3"
	chart.Categories = []slidescloud.IChartCategory{category1, category2, category3}
	result, _, e := c.SlidesApi.CreateShape(fileName, 3, chart, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	dataPoint := result.(slidescloud.IChart).GetSeries()[0].(slidescloud.IOneValueSeries).GetDataPoints()[2]

	if dataPoint.GetValue() != 90 {
		t.Errorf("Expected %v, but was %v", 90, dataPoint.GetValue())
		return
	}
}
