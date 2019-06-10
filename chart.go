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
)

// Represents chart resource
type IChart interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Gets or sets the name.
	getName() string
	setName(newValue string)

	// Gets or sets the width.
	getWidth() float64
	setWidth(newValue float64)

	// Gets or sets the height.
	getHeight() float64
	setHeight(newValue float64)

	// Gets or sets the alternative text.
	getAlternativeText() string
	setAlternativeText(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	getHidden() bool
	setHidden(newValue bool)

	// Gets or sets the X
	getX() float64
	setX(newValue float64)

	// Gets or sets the Y.
	getY() float64
	setY(newValue float64)

	// Gets z-order position of shape
	getZOrderPosition() int32
	setZOrderPosition(newValue int32)

	// Gets or sets the link to shapes.
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	getType() string
	setType(newValue string)

	getShapeType() string
	setShapeType(newValue string)

	// Gets or sets the type of the chart.
	getChartType() string
	setChartType(newValue string)

	// Gets or sets the series of chart data values.
	getSeries() []Series
	setSeries(newValue []Series)

	// Gets or sets the categories for chart data
	getCategories() []string
	setCategories(newValue []string)

	// Gets or sets the title.
	getTitle() IChartTitle
	setTitle(newValue IChartTitle)

	// Gets or sets the back wall.
	getBackWall() IChartWall
	setBackWall(newValue IChartWall)

	// Gets or sets the side wall.
	getSideWall() IChartWall
	setSideWall(newValue IChartWall)

	// Gets or sets the floor.
	getFloor() IChartWall
	setFloor(newValue IChartWall)

	// Gets or sets the legend.
	getLegend() ILegend
	setLegend(newValue ILegend)

	// Gets or sets the axes.
	getAxes() IAxes
	setAxes(newValue IAxes)

	// Gets or sets the plot area.
	getPlotArea() IPlotArea
	setPlotArea(newValue IPlotArea)
}

type Chart struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden bool `json:"Hidden,omitempty"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	Type_ string `json:"Type"`

	ShapeType string `json:"ShapeType"`

	// Gets or sets the type of the chart.
	ChartType string `json:"ChartType"`

	// Gets or sets the series of chart data values.
	Series []Series `json:"Series,omitempty"`

	// Gets or sets the categories for chart data
	Categories []string `json:"Categories,omitempty"`

	// Gets or sets the title.
	Title IChartTitle `json:"Title,omitempty"`

	// Gets or sets the back wall.
	BackWall IChartWall `json:"BackWall,omitempty"`

	// Gets or sets the side wall.
	SideWall IChartWall `json:"SideWall,omitempty"`

	// Gets or sets the floor.
	Floor IChartWall `json:"Floor,omitempty"`

	// Gets or sets the legend.
	Legend ILegend `json:"Legend,omitempty"`

	// Gets or sets the axes.
	Axes IAxes `json:"Axes,omitempty"`

	// Gets or sets the plot area.
	PlotArea IPlotArea `json:"PlotArea,omitempty"`
}

func (this Chart) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Chart) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Chart) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Chart) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Chart) getName() string {
	return this.Name
}

func (this Chart) setName(newValue string) {
	this.Name = newValue
}
func (this Chart) getWidth() float64 {
	return this.Width
}

func (this Chart) setWidth(newValue float64) {
	this.Width = newValue
}
func (this Chart) getHeight() float64 {
	return this.Height
}

func (this Chart) setHeight(newValue float64) {
	this.Height = newValue
}
func (this Chart) getAlternativeText() string {
	return this.AlternativeText
}

func (this Chart) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this Chart) getHidden() bool {
	return this.Hidden
}

func (this Chart) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this Chart) getX() float64 {
	return this.X
}

func (this Chart) setX(newValue float64) {
	this.X = newValue
}
func (this Chart) getY() float64 {
	return this.Y
}

func (this Chart) setY(newValue float64) {
	this.Y = newValue
}
func (this Chart) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this Chart) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this Chart) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this Chart) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this Chart) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Chart) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Chart) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Chart) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Chart) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Chart) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this Chart) getType() string {
	return this.Type_
}

func (this Chart) setType(newValue string) {
	this.Type_ = newValue
}
func (this Chart) getShapeType() string {
	return this.ShapeType
}

func (this Chart) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this Chart) getChartType() string {
	return this.ChartType
}

func (this Chart) setChartType(newValue string) {
	this.ChartType = newValue
}
func (this Chart) getSeries() []Series {
	return this.Series
}

func (this Chart) setSeries(newValue []Series) {
	this.Series = newValue
}
func (this Chart) getCategories() []string {
	return this.Categories
}

func (this Chart) setCategories(newValue []string) {
	this.Categories = newValue
}
func (this Chart) getTitle() IChartTitle {
	return this.Title
}

func (this Chart) setTitle(newValue IChartTitle) {
	this.Title = newValue
}
func (this Chart) getBackWall() IChartWall {
	return this.BackWall
}

func (this Chart) setBackWall(newValue IChartWall) {
	this.BackWall = newValue
}
func (this Chart) getSideWall() IChartWall {
	return this.SideWall
}

func (this Chart) setSideWall(newValue IChartWall) {
	this.SideWall = newValue
}
func (this Chart) getFloor() IChartWall {
	return this.Floor
}

func (this Chart) setFloor(newValue IChartWall) {
	this.Floor = newValue
}
func (this Chart) getLegend() ILegend {
	return this.Legend
}

func (this Chart) setLegend(newValue ILegend) {
	this.Legend = newValue
}
func (this Chart) getAxes() IAxes {
	return this.Axes
}

func (this Chart) setAxes(newValue IAxes) {
	this.Axes = newValue
}
func (this Chart) getPlotArea() IPlotArea {
	return this.PlotArea
}

func (this Chart) setPlotArea(newValue IPlotArea) {
	this.PlotArea = newValue
}

func (this *Chart) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valSelfUri, ok := objMap["SelfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}

	if valAlternateLinks, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}

	if valName, ok := objMap["Name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

	if valHeight, ok := objMap["Height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	if valAlternativeText, ok := objMap["AlternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}

	if valHidden, ok := objMap["Hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}

	if valX, ok := objMap["X"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}

	if valY, ok := objMap["Y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}

	if valZOrderPosition, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}

	if valShapes, ok := objMap["Shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}

	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}

	if valEffectFormat, ok := objMap["EffectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}

	if valLineFormat, ok := objMap["LineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valShapeType, ok := objMap["ShapeType"]; ok {
		if valShapeType != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeType, &valueForShapeType)
			if err != nil {
				return err
			}
			this.ShapeType = valueForShapeType
		}
	}

	if valChartType, ok := objMap["ChartType"]; ok {
		if valChartType != nil {
			var valueForChartType string
			err = json.Unmarshal(*valChartType, &valueForChartType)
			if err != nil {
				return err
			}
			this.ChartType = valueForChartType
		}
	}

	if valSeries, ok := objMap["Series"]; ok {
		if valSeries != nil {
			var valueForSeries []Series
			err = json.Unmarshal(*valSeries, &valueForSeries)
			if err != nil {
				return err
			}
			this.Series = valueForSeries
		}
	}

	if valCategories, ok := objMap["Categories"]; ok {
		if valCategories != nil {
			var valueForCategories []string
			err = json.Unmarshal(*valCategories, &valueForCategories)
			if err != nil {
				return err
			}
			this.Categories = valueForCategories
		}
	}

	if valTitle, ok := objMap["Title"]; ok {
		if valTitle != nil {
			var valueForTitle ChartTitle
			err = json.Unmarshal(*valTitle, &valueForTitle)
			if err != nil {
				return err
			}
			this.Title = valueForTitle
		}
	}

	if valBackWall, ok := objMap["BackWall"]; ok {
		if valBackWall != nil {
			var valueForBackWall ChartWall
			err = json.Unmarshal(*valBackWall, &valueForBackWall)
			if err != nil {
				return err
			}
			this.BackWall = valueForBackWall
		}
	}

	if valSideWall, ok := objMap["SideWall"]; ok {
		if valSideWall != nil {
			var valueForSideWall ChartWall
			err = json.Unmarshal(*valSideWall, &valueForSideWall)
			if err != nil {
				return err
			}
			this.SideWall = valueForSideWall
		}
	}

	if valFloor, ok := objMap["Floor"]; ok {
		if valFloor != nil {
			var valueForFloor ChartWall
			err = json.Unmarshal(*valFloor, &valueForFloor)
			if err != nil {
				return err
			}
			this.Floor = valueForFloor
		}
	}

	if valLegend, ok := objMap["Legend"]; ok {
		if valLegend != nil {
			var valueForLegend Legend
			err = json.Unmarshal(*valLegend, &valueForLegend)
			if err != nil {
				return err
			}
			this.Legend = valueForLegend
		}
	}

	if valAxes, ok := objMap["Axes"]; ok {
		if valAxes != nil {
			var valueForAxes Axes
			err = json.Unmarshal(*valAxes, &valueForAxes)
			if err != nil {
				return err
			}
			this.Axes = valueForAxes
		}
	}

	if valPlotArea, ok := objMap["PlotArea"]; ok {
		if valPlotArea != nil {
			var valueForPlotArea PlotArea
			err = json.Unmarshal(*valPlotArea, &valueForPlotArea)
			if err != nil {
				return err
			}
			this.PlotArea = valueForPlotArea
		}
	}

    return nil
}
