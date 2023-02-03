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
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	GetName() string
	SetName(newValue string)

	// Gets or sets the width.
	GetWidth() float64
	SetWidth(newValue float64)

	// Gets or sets the height.
	GetHeight() float64
	SetHeight(newValue float64)

	// Gets or sets the alternative text.
	GetAlternativeText() string
	SetAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	GetAlternativeTextTitle() string
	SetAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	GetHidden() bool
	SetHidden(newValue bool)

	// Gets or sets the X
	GetX() float64
	SetX(newValue float64)

	// Gets or sets the Y.
	GetY() float64
	SetY(newValue float64)

	// Gets z-order position of shape
	GetZOrderPosition() int32
	SetZOrderPosition(newValue int32)

	// Gets or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Shape type.
	GetType() string
	SetType(newValue string)

	// Gets or sets the type of the chart.
	GetChartType() string
	SetChartType(newValue string)

	// True if data labels over the maximum of the chart shall be shown.
	GetShowDataLabelsOverMaximum() bool
	SetShowDataLabelsOverMaximum(newValue bool)

	// Gets or sets the series of chart data values.
	GetSeries() []ISeries
	SetSeries(newValue []ISeries)

	// Gets or sets the categories for chart data
	GetCategories() []IChartCategory
	SetCategories(newValue []IChartCategory)

	// Data source type for categories.
	GetDataSourceForCategories() IDataSource
	SetDataSourceForCategories(newValue IDataSource)

	// Gets or sets the title.
	GetTitle() IChartTitle
	SetTitle(newValue IChartTitle)

	// Gets or sets the back wall.
	GetBackWall() IChartWall
	SetBackWall(newValue IChartWall)

	// Gets or sets the side wall.
	GetSideWall() IChartWall
	SetSideWall(newValue IChartWall)

	// Gets or sets the floor.
	GetFloor() IChartWall
	SetFloor(newValue IChartWall)

	// Gets or sets the legend.
	GetLegend() ILegend
	SetLegend(newValue ILegend)

	// Gets or sets the axes.
	GetAxes() IAxes
	SetAxes(newValue IAxes)

	// Gets or sets the plot area.
	GetPlotArea() IPlotArea
	SetPlotArea(newValue IPlotArea)

	// Specifies the chart area shall have rounded corners.
	GetHasRoundedCorners() bool
	SetHasRoundedCorners(newValue bool)

	// Gets groups of series. 
	GetSeriesGroups() []IChartSeriesGroup
	SetSeriesGroups(newValue []IChartSeriesGroup)
}

type Chart struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// The title of alternative text associated with the shape.
	AlternativeTextTitle string `json:"AlternativeTextTitle,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden bool `json:"Hidden"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Gets or sets the type of the chart.
	ChartType string `json:"ChartType"`

	// True if data labels over the maximum of the chart shall be shown.
	ShowDataLabelsOverMaximum bool `json:"ShowDataLabelsOverMaximum"`

	// Gets or sets the series of chart data values.
	Series []ISeries `json:"Series,omitempty"`

	// Gets or sets the categories for chart data
	Categories []IChartCategory `json:"Categories,omitempty"`

	// Data source type for categories.
	DataSourceForCategories IDataSource `json:"DataSourceForCategories,omitempty"`

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

	// Specifies the chart area shall have rounded corners.
	HasRoundedCorners bool `json:"HasRoundedCorners"`

	// Gets groups of series. 
	SeriesGroups []IChartSeriesGroup `json:"SeriesGroups,omitempty"`
}

func NewChart() *Chart {
	instance := new(Chart)
	instance.Type_ = "Chart"
	instance.ChartType = "ClusteredColumn"
	return instance
}

func (this *Chart) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Chart) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Chart) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Chart) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Chart) GetName() string {
	return this.Name
}

func (this *Chart) SetName(newValue string) {
	this.Name = newValue
}
func (this *Chart) GetWidth() float64 {
	return this.Width
}

func (this *Chart) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *Chart) GetHeight() float64 {
	return this.Height
}

func (this *Chart) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *Chart) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *Chart) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *Chart) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *Chart) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *Chart) GetHidden() bool {
	return this.Hidden
}

func (this *Chart) SetHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *Chart) GetX() float64 {
	return this.X
}

func (this *Chart) SetX(newValue float64) {
	this.X = newValue
}
func (this *Chart) GetY() float64 {
	return this.Y
}

func (this *Chart) SetY(newValue float64) {
	this.Y = newValue
}
func (this *Chart) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *Chart) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *Chart) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Chart) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Chart) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Chart) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Chart) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *Chart) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *Chart) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Chart) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Chart) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *Chart) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *Chart) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *Chart) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *Chart) GetType() string {
	return this.Type_
}

func (this *Chart) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Chart) GetChartType() string {
	return this.ChartType
}

func (this *Chart) SetChartType(newValue string) {
	this.ChartType = newValue
}
func (this *Chart) GetShowDataLabelsOverMaximum() bool {
	return this.ShowDataLabelsOverMaximum
}

func (this *Chart) SetShowDataLabelsOverMaximum(newValue bool) {
	this.ShowDataLabelsOverMaximum = newValue
}
func (this *Chart) GetSeries() []ISeries {
	return this.Series
}

func (this *Chart) SetSeries(newValue []ISeries) {
	this.Series = newValue
}
func (this *Chart) GetCategories() []IChartCategory {
	return this.Categories
}

func (this *Chart) SetCategories(newValue []IChartCategory) {
	this.Categories = newValue
}
func (this *Chart) GetDataSourceForCategories() IDataSource {
	return this.DataSourceForCategories
}

func (this *Chart) SetDataSourceForCategories(newValue IDataSource) {
	this.DataSourceForCategories = newValue
}
func (this *Chart) GetTitle() IChartTitle {
	return this.Title
}

func (this *Chart) SetTitle(newValue IChartTitle) {
	this.Title = newValue
}
func (this *Chart) GetBackWall() IChartWall {
	return this.BackWall
}

func (this *Chart) SetBackWall(newValue IChartWall) {
	this.BackWall = newValue
}
func (this *Chart) GetSideWall() IChartWall {
	return this.SideWall
}

func (this *Chart) SetSideWall(newValue IChartWall) {
	this.SideWall = newValue
}
func (this *Chart) GetFloor() IChartWall {
	return this.Floor
}

func (this *Chart) SetFloor(newValue IChartWall) {
	this.Floor = newValue
}
func (this *Chart) GetLegend() ILegend {
	return this.Legend
}

func (this *Chart) SetLegend(newValue ILegend) {
	this.Legend = newValue
}
func (this *Chart) GetAxes() IAxes {
	return this.Axes
}

func (this *Chart) SetAxes(newValue IAxes) {
	this.Axes = newValue
}
func (this *Chart) GetPlotArea() IPlotArea {
	return this.PlotArea
}

func (this *Chart) SetPlotArea(newValue IPlotArea) {
	this.PlotArea = newValue
}
func (this *Chart) GetHasRoundedCorners() bool {
	return this.HasRoundedCorners
}

func (this *Chart) SetHasRoundedCorners(newValue bool) {
	this.HasRoundedCorners = newValue
}
func (this *Chart) GetSeriesGroups() []IChartSeriesGroup {
	return this.SeriesGroups
}

func (this *Chart) SetSeriesGroups(newValue []IChartSeriesGroup) {
	this.SeriesGroups = newValue
}

func (this *Chart) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["selfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := objMap["alternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	if valAlternativeTextCap, ok := objMap["AlternativeText"]; ok {
		if valAlternativeTextCap != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeTextCap, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := objMap["alternativeTextTitle"]; ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	if valAlternativeTextTitleCap, ok := objMap["AlternativeTextTitle"]; ok {
		if valAlternativeTextTitleCap != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitleCap, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := objMap["hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	if valHiddenCap, ok := objMap["Hidden"]; ok {
		if valHiddenCap != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHiddenCap, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valX, ok := objMap["x"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	if valXCap, ok := objMap["X"]; ok {
		if valXCap != nil {
			var valueForX float64
			err = json.Unmarshal(*valXCap, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := objMap["y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	if valYCap, ok := objMap["Y"]; ok {
		if valYCap != nil {
			var valueForY float64
			err = json.Unmarshal(*valYCap, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := objMap["zOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	if valZOrderPositionCap, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPositionCap != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPositionCap, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkClick, ok := objMap["hyperlinkClick"]; ok {
		if valHyperlinkClick != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClick, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClick)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClick, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	if valHyperlinkClickCap, ok := objMap["HyperlinkClick"]; ok {
		if valHyperlinkClickCap != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClickCap, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClickCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClickCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkMouseOver, ok := objMap["hyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOver != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOver, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOver)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOver, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	if valHyperlinkMouseOverCap, ok := objMap["HyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOverCap != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOverCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	this.Type_ = "Chart"
	if valType, ok := objMap["type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	this.ChartType = "ClusteredColumn"
	if valChartType, ok := objMap["chartType"]; ok {
		if valChartType != nil {
			var valueForChartType string
			err = json.Unmarshal(*valChartType, &valueForChartType)
			if err != nil {
				var valueForChartTypeInt int32
				err = json.Unmarshal(*valChartType, &valueForChartTypeInt)
				if err != nil {
					return err
				}
				this.ChartType = string(valueForChartTypeInt)
			} else {
				this.ChartType = valueForChartType
			}
		}
	}
	if valChartTypeCap, ok := objMap["ChartType"]; ok {
		if valChartTypeCap != nil {
			var valueForChartType string
			err = json.Unmarshal(*valChartTypeCap, &valueForChartType)
			if err != nil {
				var valueForChartTypeInt int32
				err = json.Unmarshal(*valChartTypeCap, &valueForChartTypeInt)
				if err != nil {
					return err
				}
				this.ChartType = string(valueForChartTypeInt)
			} else {
				this.ChartType = valueForChartType
			}
		}
	}
	
	if valShowDataLabelsOverMaximum, ok := objMap["showDataLabelsOverMaximum"]; ok {
		if valShowDataLabelsOverMaximum != nil {
			var valueForShowDataLabelsOverMaximum bool
			err = json.Unmarshal(*valShowDataLabelsOverMaximum, &valueForShowDataLabelsOverMaximum)
			if err != nil {
				return err
			}
			this.ShowDataLabelsOverMaximum = valueForShowDataLabelsOverMaximum
		}
	}
	if valShowDataLabelsOverMaximumCap, ok := objMap["ShowDataLabelsOverMaximum"]; ok {
		if valShowDataLabelsOverMaximumCap != nil {
			var valueForShowDataLabelsOverMaximum bool
			err = json.Unmarshal(*valShowDataLabelsOverMaximumCap, &valueForShowDataLabelsOverMaximum)
			if err != nil {
				return err
			}
			this.ShowDataLabelsOverMaximum = valueForShowDataLabelsOverMaximum
		}
	}
	
	if valSeries, ok := objMap["series"]; ok {
		if valSeries != nil {
			var valueForSeries []json.RawMessage
			err = json.Unmarshal(*valSeries, &valueForSeries)
			if err != nil {
				return err
			}
			valueForISeries := make([]ISeries, len(valueForSeries))
			for i, v := range valueForSeries {
				vObject, err := createObjectForType("Series", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISeries[i] = vObject.(ISeries)
				}
			}
			this.Series = valueForISeries
		}
	}
	if valSeriesCap, ok := objMap["Series"]; ok {
		if valSeriesCap != nil {
			var valueForSeries []json.RawMessage
			err = json.Unmarshal(*valSeriesCap, &valueForSeries)
			if err != nil {
				return err
			}
			valueForISeries := make([]ISeries, len(valueForSeries))
			for i, v := range valueForSeries {
				vObject, err := createObjectForType("Series", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISeries[i] = vObject.(ISeries)
				}
			}
			this.Series = valueForISeries
		}
	}
	
	if valCategories, ok := objMap["categories"]; ok {
		if valCategories != nil {
			var valueForCategories []json.RawMessage
			err = json.Unmarshal(*valCategories, &valueForCategories)
			if err != nil {
				return err
			}
			valueForICategories := make([]IChartCategory, len(valueForCategories))
			for i, v := range valueForCategories {
				vObject, err := createObjectForType("ChartCategory", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForICategories[i] = vObject.(IChartCategory)
				}
			}
			this.Categories = valueForICategories
		}
	}
	if valCategoriesCap, ok := objMap["Categories"]; ok {
		if valCategoriesCap != nil {
			var valueForCategories []json.RawMessage
			err = json.Unmarshal(*valCategoriesCap, &valueForCategories)
			if err != nil {
				return err
			}
			valueForICategories := make([]IChartCategory, len(valueForCategories))
			for i, v := range valueForCategories {
				vObject, err := createObjectForType("ChartCategory", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForICategories[i] = vObject.(IChartCategory)
				}
			}
			this.Categories = valueForICategories
		}
	}
	
	if valDataSourceForCategories, ok := objMap["dataSourceForCategories"]; ok {
		if valDataSourceForCategories != nil {
			var valueForDataSourceForCategories DataSource
			err = json.Unmarshal(*valDataSourceForCategories, &valueForDataSourceForCategories)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForCategories)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForCategories, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForCategories = vInterfaceObject
			}
		}
	}
	if valDataSourceForCategoriesCap, ok := objMap["DataSourceForCategories"]; ok {
		if valDataSourceForCategoriesCap != nil {
			var valueForDataSourceForCategories DataSource
			err = json.Unmarshal(*valDataSourceForCategoriesCap, &valueForDataSourceForCategories)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("DataSource", *valDataSourceForCategoriesCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDataSourceForCategoriesCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IDataSource)
			if ok {
				this.DataSourceForCategories = vInterfaceObject
			}
		}
	}
	
	if valTitle, ok := objMap["title"]; ok {
		if valTitle != nil {
			var valueForTitle ChartTitle
			err = json.Unmarshal(*valTitle, &valueForTitle)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartTitle", *valTitle)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valTitle, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartTitle)
			if ok {
				this.Title = vInterfaceObject
			}
		}
	}
	if valTitleCap, ok := objMap["Title"]; ok {
		if valTitleCap != nil {
			var valueForTitle ChartTitle
			err = json.Unmarshal(*valTitleCap, &valueForTitle)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartTitle", *valTitleCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valTitleCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartTitle)
			if ok {
				this.Title = vInterfaceObject
			}
		}
	}
	
	if valBackWall, ok := objMap["backWall"]; ok {
		if valBackWall != nil {
			var valueForBackWall ChartWall
			err = json.Unmarshal(*valBackWall, &valueForBackWall)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valBackWall)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBackWall, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.BackWall = vInterfaceObject
			}
		}
	}
	if valBackWallCap, ok := objMap["BackWall"]; ok {
		if valBackWallCap != nil {
			var valueForBackWall ChartWall
			err = json.Unmarshal(*valBackWallCap, &valueForBackWall)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valBackWallCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBackWallCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.BackWall = vInterfaceObject
			}
		}
	}
	
	if valSideWall, ok := objMap["sideWall"]; ok {
		if valSideWall != nil {
			var valueForSideWall ChartWall
			err = json.Unmarshal(*valSideWall, &valueForSideWall)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valSideWall)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSideWall, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.SideWall = vInterfaceObject
			}
		}
	}
	if valSideWallCap, ok := objMap["SideWall"]; ok {
		if valSideWallCap != nil {
			var valueForSideWall ChartWall
			err = json.Unmarshal(*valSideWallCap, &valueForSideWall)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valSideWallCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSideWallCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.SideWall = vInterfaceObject
			}
		}
	}
	
	if valFloor, ok := objMap["floor"]; ok {
		if valFloor != nil {
			var valueForFloor ChartWall
			err = json.Unmarshal(*valFloor, &valueForFloor)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valFloor)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFloor, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.Floor = vInterfaceObject
			}
		}
	}
	if valFloorCap, ok := objMap["Floor"]; ok {
		if valFloorCap != nil {
			var valueForFloor ChartWall
			err = json.Unmarshal(*valFloorCap, &valueForFloor)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartWall", *valFloorCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFloorCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartWall)
			if ok {
				this.Floor = vInterfaceObject
			}
		}
	}
	
	if valLegend, ok := objMap["legend"]; ok {
		if valLegend != nil {
			var valueForLegend Legend
			err = json.Unmarshal(*valLegend, &valueForLegend)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Legend", *valLegend)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLegend, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILegend)
			if ok {
				this.Legend = vInterfaceObject
			}
		}
	}
	if valLegendCap, ok := objMap["Legend"]; ok {
		if valLegendCap != nil {
			var valueForLegend Legend
			err = json.Unmarshal(*valLegendCap, &valueForLegend)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Legend", *valLegendCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLegendCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILegend)
			if ok {
				this.Legend = vInterfaceObject
			}
		}
	}
	
	if valAxes, ok := objMap["axes"]; ok {
		if valAxes != nil {
			var valueForAxes Axes
			err = json.Unmarshal(*valAxes, &valueForAxes)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axes", *valAxes)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valAxes, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxes)
			if ok {
				this.Axes = vInterfaceObject
			}
		}
	}
	if valAxesCap, ok := objMap["Axes"]; ok {
		if valAxesCap != nil {
			var valueForAxes Axes
			err = json.Unmarshal(*valAxesCap, &valueForAxes)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axes", *valAxesCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valAxesCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxes)
			if ok {
				this.Axes = vInterfaceObject
			}
		}
	}
	
	if valPlotArea, ok := objMap["plotArea"]; ok {
		if valPlotArea != nil {
			var valueForPlotArea PlotArea
			err = json.Unmarshal(*valPlotArea, &valueForPlotArea)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PlotArea", *valPlotArea)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPlotArea, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPlotArea)
			if ok {
				this.PlotArea = vInterfaceObject
			}
		}
	}
	if valPlotAreaCap, ok := objMap["PlotArea"]; ok {
		if valPlotAreaCap != nil {
			var valueForPlotArea PlotArea
			err = json.Unmarshal(*valPlotAreaCap, &valueForPlotArea)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PlotArea", *valPlotAreaCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPlotAreaCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPlotArea)
			if ok {
				this.PlotArea = vInterfaceObject
			}
		}
	}
	
	if valHasRoundedCorners, ok := objMap["hasRoundedCorners"]; ok {
		if valHasRoundedCorners != nil {
			var valueForHasRoundedCorners bool
			err = json.Unmarshal(*valHasRoundedCorners, &valueForHasRoundedCorners)
			if err != nil {
				return err
			}
			this.HasRoundedCorners = valueForHasRoundedCorners
		}
	}
	if valHasRoundedCornersCap, ok := objMap["HasRoundedCorners"]; ok {
		if valHasRoundedCornersCap != nil {
			var valueForHasRoundedCorners bool
			err = json.Unmarshal(*valHasRoundedCornersCap, &valueForHasRoundedCorners)
			if err != nil {
				return err
			}
			this.HasRoundedCorners = valueForHasRoundedCorners
		}
	}
	
	if valSeriesGroups, ok := objMap["seriesGroups"]; ok {
		if valSeriesGroups != nil {
			var valueForSeriesGroups []json.RawMessage
			err = json.Unmarshal(*valSeriesGroups, &valueForSeriesGroups)
			if err != nil {
				return err
			}
			valueForISeriesGroups := make([]IChartSeriesGroup, len(valueForSeriesGroups))
			for i, v := range valueForSeriesGroups {
				vObject, err := createObjectForType("ChartSeriesGroup", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISeriesGroups[i] = vObject.(IChartSeriesGroup)
				}
			}
			this.SeriesGroups = valueForISeriesGroups
		}
	}
	if valSeriesGroupsCap, ok := objMap["SeriesGroups"]; ok {
		if valSeriesGroupsCap != nil {
			var valueForSeriesGroups []json.RawMessage
			err = json.Unmarshal(*valSeriesGroupsCap, &valueForSeriesGroups)
			if err != nil {
				return err
			}
			valueForISeriesGroups := make([]IChartSeriesGroup, len(valueForSeriesGroups))
			for i, v := range valueForSeriesGroups {
				vObject, err := createObjectForType("ChartSeriesGroup", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISeriesGroups[i] = vObject.(IChartSeriesGroup)
				}
			}
			this.SeriesGroups = valueForISeriesGroups
		}
	}

	return nil
}