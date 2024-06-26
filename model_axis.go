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

// Represents a chart axis
type IAxis interface {

	// True if the axis is visible
	GetIsVisible() *bool
	SetIsVisible(newValue *bool)

	// True if the axis has a visible title
	GetHasTitle() *bool
	SetHasTitle(newValue *bool)

	// Axis title
	GetTitle() IChartTitle
	SetTitle(newValue IChartTitle)

	// Axis position
	GetPosition() string
	SetPosition(newValue string)

	// The scaling value of the display units for the value axis
	GetDisplayUnit() string
	SetDisplayUnit(newValue string)

	// The smallest time unit that is represented on the date axis
	GetBaseUnitScale() string
	SetBaseUnitScale(newValue string)

	// True the major unit of the axis is automatically assigned
	GetIsAutomaticMajorUnit() *bool
	SetIsAutomaticMajorUnit(newValue *bool)

	// The major units for the date or value axis
	GetMajorUnit() float64
	SetMajorUnit(newValue float64)

	// The major unit scale for the date axis
	GetMajorUnitScale() string
	SetMajorUnitScale(newValue string)

	// The type of major tick mark for the specified axis
	GetMajorTickMark() string
	SetMajorTickMark(newValue string)

	// True the minor unit of the axis is automatically assigned
	GetIsAutomaticMinorUnit() *bool
	SetIsAutomaticMinorUnit(newValue *bool)

	// The minor units for the date or value axis
	GetMinorUnit() float64
	SetMinorUnit(newValue float64)

	// The minor unit scale for the date axis
	GetMinorUnitScale() string
	SetMinorUnitScale(newValue string)

	// The type of minor tick mark for the specified axis
	GetMinorTickMark() string
	SetMinorTickMark(newValue string)

	// True if the max value is automatically assigned
	GetIsAutomaticMaxValue() *bool
	SetIsAutomaticMaxValue(newValue *bool)

	// The maximum value on the value axis
	GetMaxValue() float64
	SetMaxValue(newValue float64)

	// True if the min value is automatically assigned
	GetIsAutomaticMinValue() *bool
	SetIsAutomaticMinValue(newValue *bool)

	// The minimum value on the value axis
	GetMinValue() float64
	SetMinValue(newValue float64)

	// True if the value axis scale type is logarithmic
	GetIsLogarithmic() *bool
	SetIsLogarithmic(newValue *bool)

	// The logarithmic base. Default value is 10
	GetLogBase() float64
	SetLogBase(newValue float64)

	// The type of the category axis
	GetCategoryAxisType() string
	SetCategoryAxisType(newValue string)

	// True if the value axis crosses the category axis between categories. This property applies only to category axes, and it doesn't apply to 3-D charts
	GetAxisBetweenCategories() *bool
	SetAxisBetweenCategories(newValue *bool)

	// The distance of labels from the axis. Applied to category or date axis. Value must be between 0% and 1000%.             
	GetLabelOffset() int32
	SetLabelOffset(newValue int32)

	// True if MS PowerPoint plots data points from last to first
	GetIsPlotOrderReversed() *bool
	SetIsPlotOrderReversed(newValue *bool)

	// True if the format is linked to source data
	GetIsNumberFormatLinkedToSource() *bool
	SetIsNumberFormatLinkedToSource(newValue *bool)

	// the format string for the Axis Labels
	GetNumberFormat() string
	SetNumberFormat(newValue string)

	// The CrossType on the specified axis where the other axis crosses
	GetCrossType() string
	SetCrossType(newValue string)

	// The point on the axis where the perpendicular axis crosses it
	GetCrossAt() float64
	SetCrossAt(newValue float64)

	// True for automatic tick marks spacing value
	GetIsAutomaticTickMarksSpacing() *bool
	SetIsAutomaticTickMarksSpacing(newValue *bool)

	// Specifies how many tick marks shall be skipped before the next one shall be drawn. Applied to category or series axis.
	GetTickMarksSpacing() int32
	SetTickMarksSpacing(newValue int32)

	// True for automatic tick label spacing value
	GetIsAutomaticTickLabelSpacing() *bool
	SetIsAutomaticTickLabelSpacing(newValue *bool)

	// Specifies how many tick labels to skip between label that is drawn.
	GetTickLabelSpacing() int32
	SetTickLabelSpacing(newValue int32)

	// The position of tick-mark labels on the specified axis.
	GetTickLabelPosition() string
	SetTickLabelPosition(newValue string)

	// Represents the rotation angle of tick labels.
	GetTickLabelRotationAngle() float64
	SetTickLabelRotationAngle(newValue float64)

	// Get or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Get or sets the format of major grid lines.
	GetMajorGridLinesFormat() IChartLinesFormat
	SetMajorGridLinesFormat(newValue IChartLinesFormat)

	// Get or sets the format of major grid lines.
	GetMinorGridLinesFormat() IChartLinesFormat
	SetMinorGridLinesFormat(newValue IChartLinesFormat)
}

type Axis struct {

	// True if the axis is visible
	IsVisible *bool `json:"IsVisible"`

	// True if the axis has a visible title
	HasTitle *bool `json:"HasTitle"`

	// Axis title
	Title IChartTitle `json:"Title,omitempty"`

	// Axis position
	Position string `json:"Position,omitempty"`

	// The scaling value of the display units for the value axis
	DisplayUnit string `json:"DisplayUnit,omitempty"`

	// The smallest time unit that is represented on the date axis
	BaseUnitScale string `json:"BaseUnitScale,omitempty"`

	// True the major unit of the axis is automatically assigned
	IsAutomaticMajorUnit *bool `json:"IsAutomaticMajorUnit"`

	// The major units for the date or value axis
	MajorUnit float64 `json:"MajorUnit,omitempty"`

	// The major unit scale for the date axis
	MajorUnitScale string `json:"MajorUnitScale,omitempty"`

	// The type of major tick mark for the specified axis
	MajorTickMark string `json:"MajorTickMark,omitempty"`

	// True the minor unit of the axis is automatically assigned
	IsAutomaticMinorUnit *bool `json:"IsAutomaticMinorUnit"`

	// The minor units for the date or value axis
	MinorUnit float64 `json:"MinorUnit,omitempty"`

	// The minor unit scale for the date axis
	MinorUnitScale string `json:"MinorUnitScale,omitempty"`

	// The type of minor tick mark for the specified axis
	MinorTickMark string `json:"MinorTickMark,omitempty"`

	// True if the max value is automatically assigned
	IsAutomaticMaxValue *bool `json:"IsAutomaticMaxValue"`

	// The maximum value on the value axis
	MaxValue float64 `json:"MaxValue,omitempty"`

	// True if the min value is automatically assigned
	IsAutomaticMinValue *bool `json:"IsAutomaticMinValue"`

	// The minimum value on the value axis
	MinValue float64 `json:"MinValue,omitempty"`

	// True if the value axis scale type is logarithmic
	IsLogarithmic *bool `json:"IsLogarithmic"`

	// The logarithmic base. Default value is 10
	LogBase float64 `json:"LogBase,omitempty"`

	// The type of the category axis
	CategoryAxisType string `json:"CategoryAxisType,omitempty"`

	// True if the value axis crosses the category axis between categories. This property applies only to category axes, and it doesn't apply to 3-D charts
	AxisBetweenCategories *bool `json:"AxisBetweenCategories"`

	// The distance of labels from the axis. Applied to category or date axis. Value must be between 0% and 1000%.             
	LabelOffset int32 `json:"LabelOffset,omitempty"`

	// True if MS PowerPoint plots data points from last to first
	IsPlotOrderReversed *bool `json:"IsPlotOrderReversed"`

	// True if the format is linked to source data
	IsNumberFormatLinkedToSource *bool `json:"IsNumberFormatLinkedToSource"`

	// the format string for the Axis Labels
	NumberFormat string `json:"NumberFormat,omitempty"`

	// The CrossType on the specified axis where the other axis crosses
	CrossType string `json:"CrossType,omitempty"`

	// The point on the axis where the perpendicular axis crosses it
	CrossAt float64 `json:"CrossAt,omitempty"`

	// True for automatic tick marks spacing value
	IsAutomaticTickMarksSpacing *bool `json:"IsAutomaticTickMarksSpacing"`

	// Specifies how many tick marks shall be skipped before the next one shall be drawn. Applied to category or series axis.
	TickMarksSpacing int32 `json:"TickMarksSpacing,omitempty"`

	// True for automatic tick label spacing value
	IsAutomaticTickLabelSpacing *bool `json:"IsAutomaticTickLabelSpacing"`

	// Specifies how many tick labels to skip between label that is drawn.
	TickLabelSpacing int32 `json:"TickLabelSpacing,omitempty"`

	// The position of tick-mark labels on the specified axis.
	TickLabelPosition string `json:"TickLabelPosition,omitempty"`

	// Represents the rotation angle of tick labels.
	TickLabelRotationAngle float64 `json:"TickLabelRotationAngle,omitempty"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Get or sets the format of major grid lines.
	MajorGridLinesFormat IChartLinesFormat `json:"MajorGridLinesFormat,omitempty"`

	// Get or sets the format of major grid lines.
	MinorGridLinesFormat IChartLinesFormat `json:"MinorGridLinesFormat,omitempty"`
}

func NewAxis() *Axis {
	instance := new(Axis)
	return instance
}

func (this *Axis) GetIsVisible() *bool {
	return this.IsVisible
}

func (this *Axis) SetIsVisible(newValue *bool) {
	this.IsVisible = newValue
}
func (this *Axis) GetHasTitle() *bool {
	return this.HasTitle
}

func (this *Axis) SetHasTitle(newValue *bool) {
	this.HasTitle = newValue
}
func (this *Axis) GetTitle() IChartTitle {
	return this.Title
}

func (this *Axis) SetTitle(newValue IChartTitle) {
	this.Title = newValue
}
func (this *Axis) GetPosition() string {
	return this.Position
}

func (this *Axis) SetPosition(newValue string) {
	this.Position = newValue
}
func (this *Axis) GetDisplayUnit() string {
	return this.DisplayUnit
}

func (this *Axis) SetDisplayUnit(newValue string) {
	this.DisplayUnit = newValue
}
func (this *Axis) GetBaseUnitScale() string {
	return this.BaseUnitScale
}

func (this *Axis) SetBaseUnitScale(newValue string) {
	this.BaseUnitScale = newValue
}
func (this *Axis) GetIsAutomaticMajorUnit() *bool {
	return this.IsAutomaticMajorUnit
}

func (this *Axis) SetIsAutomaticMajorUnit(newValue *bool) {
	this.IsAutomaticMajorUnit = newValue
}
func (this *Axis) GetMajorUnit() float64 {
	return this.MajorUnit
}

func (this *Axis) SetMajorUnit(newValue float64) {
	this.MajorUnit = newValue
}
func (this *Axis) GetMajorUnitScale() string {
	return this.MajorUnitScale
}

func (this *Axis) SetMajorUnitScale(newValue string) {
	this.MajorUnitScale = newValue
}
func (this *Axis) GetMajorTickMark() string {
	return this.MajorTickMark
}

func (this *Axis) SetMajorTickMark(newValue string) {
	this.MajorTickMark = newValue
}
func (this *Axis) GetIsAutomaticMinorUnit() *bool {
	return this.IsAutomaticMinorUnit
}

func (this *Axis) SetIsAutomaticMinorUnit(newValue *bool) {
	this.IsAutomaticMinorUnit = newValue
}
func (this *Axis) GetMinorUnit() float64 {
	return this.MinorUnit
}

func (this *Axis) SetMinorUnit(newValue float64) {
	this.MinorUnit = newValue
}
func (this *Axis) GetMinorUnitScale() string {
	return this.MinorUnitScale
}

func (this *Axis) SetMinorUnitScale(newValue string) {
	this.MinorUnitScale = newValue
}
func (this *Axis) GetMinorTickMark() string {
	return this.MinorTickMark
}

func (this *Axis) SetMinorTickMark(newValue string) {
	this.MinorTickMark = newValue
}
func (this *Axis) GetIsAutomaticMaxValue() *bool {
	return this.IsAutomaticMaxValue
}

func (this *Axis) SetIsAutomaticMaxValue(newValue *bool) {
	this.IsAutomaticMaxValue = newValue
}
func (this *Axis) GetMaxValue() float64 {
	return this.MaxValue
}

func (this *Axis) SetMaxValue(newValue float64) {
	this.MaxValue = newValue
}
func (this *Axis) GetIsAutomaticMinValue() *bool {
	return this.IsAutomaticMinValue
}

func (this *Axis) SetIsAutomaticMinValue(newValue *bool) {
	this.IsAutomaticMinValue = newValue
}
func (this *Axis) GetMinValue() float64 {
	return this.MinValue
}

func (this *Axis) SetMinValue(newValue float64) {
	this.MinValue = newValue
}
func (this *Axis) GetIsLogarithmic() *bool {
	return this.IsLogarithmic
}

func (this *Axis) SetIsLogarithmic(newValue *bool) {
	this.IsLogarithmic = newValue
}
func (this *Axis) GetLogBase() float64 {
	return this.LogBase
}

func (this *Axis) SetLogBase(newValue float64) {
	this.LogBase = newValue
}
func (this *Axis) GetCategoryAxisType() string {
	return this.CategoryAxisType
}

func (this *Axis) SetCategoryAxisType(newValue string) {
	this.CategoryAxisType = newValue
}
func (this *Axis) GetAxisBetweenCategories() *bool {
	return this.AxisBetweenCategories
}

func (this *Axis) SetAxisBetweenCategories(newValue *bool) {
	this.AxisBetweenCategories = newValue
}
func (this *Axis) GetLabelOffset() int32 {
	return this.LabelOffset
}

func (this *Axis) SetLabelOffset(newValue int32) {
	this.LabelOffset = newValue
}
func (this *Axis) GetIsPlotOrderReversed() *bool {
	return this.IsPlotOrderReversed
}

func (this *Axis) SetIsPlotOrderReversed(newValue *bool) {
	this.IsPlotOrderReversed = newValue
}
func (this *Axis) GetIsNumberFormatLinkedToSource() *bool {
	return this.IsNumberFormatLinkedToSource
}

func (this *Axis) SetIsNumberFormatLinkedToSource(newValue *bool) {
	this.IsNumberFormatLinkedToSource = newValue
}
func (this *Axis) GetNumberFormat() string {
	return this.NumberFormat
}

func (this *Axis) SetNumberFormat(newValue string) {
	this.NumberFormat = newValue
}
func (this *Axis) GetCrossType() string {
	return this.CrossType
}

func (this *Axis) SetCrossType(newValue string) {
	this.CrossType = newValue
}
func (this *Axis) GetCrossAt() float64 {
	return this.CrossAt
}

func (this *Axis) SetCrossAt(newValue float64) {
	this.CrossAt = newValue
}
func (this *Axis) GetIsAutomaticTickMarksSpacing() *bool {
	return this.IsAutomaticTickMarksSpacing
}

func (this *Axis) SetIsAutomaticTickMarksSpacing(newValue *bool) {
	this.IsAutomaticTickMarksSpacing = newValue
}
func (this *Axis) GetTickMarksSpacing() int32 {
	return this.TickMarksSpacing
}

func (this *Axis) SetTickMarksSpacing(newValue int32) {
	this.TickMarksSpacing = newValue
}
func (this *Axis) GetIsAutomaticTickLabelSpacing() *bool {
	return this.IsAutomaticTickLabelSpacing
}

func (this *Axis) SetIsAutomaticTickLabelSpacing(newValue *bool) {
	this.IsAutomaticTickLabelSpacing = newValue
}
func (this *Axis) GetTickLabelSpacing() int32 {
	return this.TickLabelSpacing
}

func (this *Axis) SetTickLabelSpacing(newValue int32) {
	this.TickLabelSpacing = newValue
}
func (this *Axis) GetTickLabelPosition() string {
	return this.TickLabelPosition
}

func (this *Axis) SetTickLabelPosition(newValue string) {
	this.TickLabelPosition = newValue
}
func (this *Axis) GetTickLabelRotationAngle() float64 {
	return this.TickLabelRotationAngle
}

func (this *Axis) SetTickLabelRotationAngle(newValue float64) {
	this.TickLabelRotationAngle = newValue
}
func (this *Axis) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Axis) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Axis) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Axis) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Axis) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Axis) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Axis) GetMajorGridLinesFormat() IChartLinesFormat {
	return this.MajorGridLinesFormat
}

func (this *Axis) SetMajorGridLinesFormat(newValue IChartLinesFormat) {
	this.MajorGridLinesFormat = newValue
}
func (this *Axis) GetMinorGridLinesFormat() IChartLinesFormat {
	return this.MinorGridLinesFormat
}

func (this *Axis) SetMinorGridLinesFormat(newValue IChartLinesFormat) {
	this.MinorGridLinesFormat = newValue
}

func (this *Axis) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valIsVisible, ok := objMap["isVisible"]; ok {
		if valIsVisible != nil {
			var valueForIsVisible *bool
			err = json.Unmarshal(*valIsVisible, &valueForIsVisible)
			if err != nil {
				return err
			}
			this.IsVisible = valueForIsVisible
		}
	}
	if valIsVisibleCap, ok := objMap["IsVisible"]; ok {
		if valIsVisibleCap != nil {
			var valueForIsVisible *bool
			err = json.Unmarshal(*valIsVisibleCap, &valueForIsVisible)
			if err != nil {
				return err
			}
			this.IsVisible = valueForIsVisible
		}
	}
	
	if valHasTitle, ok := objMap["hasTitle"]; ok {
		if valHasTitle != nil {
			var valueForHasTitle *bool
			err = json.Unmarshal(*valHasTitle, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}
	if valHasTitleCap, ok := objMap["HasTitle"]; ok {
		if valHasTitleCap != nil {
			var valueForHasTitle *bool
			err = json.Unmarshal(*valHasTitleCap, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
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
	
	if valPosition, ok := objMap["position"]; ok {
		if valPosition != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPosition, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPosition, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
	}
	if valPositionCap, ok := objMap["Position"]; ok {
		if valPositionCap != nil {
			var valueForPosition string
			err = json.Unmarshal(*valPositionCap, &valueForPosition)
			if err != nil {
				var valueForPositionInt int32
				err = json.Unmarshal(*valPositionCap, &valueForPositionInt)
				if err != nil {
					return err
				}
				this.Position = string(valueForPositionInt)
			} else {
				this.Position = valueForPosition
			}
		}
	}
	
	if valDisplayUnit, ok := objMap["displayUnit"]; ok {
		if valDisplayUnit != nil {
			var valueForDisplayUnit string
			err = json.Unmarshal(*valDisplayUnit, &valueForDisplayUnit)
			if err != nil {
				var valueForDisplayUnitInt int32
				err = json.Unmarshal(*valDisplayUnit, &valueForDisplayUnitInt)
				if err != nil {
					return err
				}
				this.DisplayUnit = string(valueForDisplayUnitInt)
			} else {
				this.DisplayUnit = valueForDisplayUnit
			}
		}
	}
	if valDisplayUnitCap, ok := objMap["DisplayUnit"]; ok {
		if valDisplayUnitCap != nil {
			var valueForDisplayUnit string
			err = json.Unmarshal(*valDisplayUnitCap, &valueForDisplayUnit)
			if err != nil {
				var valueForDisplayUnitInt int32
				err = json.Unmarshal(*valDisplayUnitCap, &valueForDisplayUnitInt)
				if err != nil {
					return err
				}
				this.DisplayUnit = string(valueForDisplayUnitInt)
			} else {
				this.DisplayUnit = valueForDisplayUnit
			}
		}
	}
	
	if valBaseUnitScale, ok := objMap["baseUnitScale"]; ok {
		if valBaseUnitScale != nil {
			var valueForBaseUnitScale string
			err = json.Unmarshal(*valBaseUnitScale, &valueForBaseUnitScale)
			if err != nil {
				var valueForBaseUnitScaleInt int32
				err = json.Unmarshal(*valBaseUnitScale, &valueForBaseUnitScaleInt)
				if err != nil {
					return err
				}
				this.BaseUnitScale = string(valueForBaseUnitScaleInt)
			} else {
				this.BaseUnitScale = valueForBaseUnitScale
			}
		}
	}
	if valBaseUnitScaleCap, ok := objMap["BaseUnitScale"]; ok {
		if valBaseUnitScaleCap != nil {
			var valueForBaseUnitScale string
			err = json.Unmarshal(*valBaseUnitScaleCap, &valueForBaseUnitScale)
			if err != nil {
				var valueForBaseUnitScaleInt int32
				err = json.Unmarshal(*valBaseUnitScaleCap, &valueForBaseUnitScaleInt)
				if err != nil {
					return err
				}
				this.BaseUnitScale = string(valueForBaseUnitScaleInt)
			} else {
				this.BaseUnitScale = valueForBaseUnitScale
			}
		}
	}
	
	if valIsAutomaticMajorUnit, ok := objMap["isAutomaticMajorUnit"]; ok {
		if valIsAutomaticMajorUnit != nil {
			var valueForIsAutomaticMajorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMajorUnit, &valueForIsAutomaticMajorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMajorUnit = valueForIsAutomaticMajorUnit
		}
	}
	if valIsAutomaticMajorUnitCap, ok := objMap["IsAutomaticMajorUnit"]; ok {
		if valIsAutomaticMajorUnitCap != nil {
			var valueForIsAutomaticMajorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMajorUnitCap, &valueForIsAutomaticMajorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMajorUnit = valueForIsAutomaticMajorUnit
		}
	}
	
	if valMajorUnit, ok := objMap["majorUnit"]; ok {
		if valMajorUnit != nil {
			var valueForMajorUnit float64
			err = json.Unmarshal(*valMajorUnit, &valueForMajorUnit)
			if err != nil {
				return err
			}
			this.MajorUnit = valueForMajorUnit
		}
	}
	if valMajorUnitCap, ok := objMap["MajorUnit"]; ok {
		if valMajorUnitCap != nil {
			var valueForMajorUnit float64
			err = json.Unmarshal(*valMajorUnitCap, &valueForMajorUnit)
			if err != nil {
				return err
			}
			this.MajorUnit = valueForMajorUnit
		}
	}
	
	if valMajorUnitScale, ok := objMap["majorUnitScale"]; ok {
		if valMajorUnitScale != nil {
			var valueForMajorUnitScale string
			err = json.Unmarshal(*valMajorUnitScale, &valueForMajorUnitScale)
			if err != nil {
				var valueForMajorUnitScaleInt int32
				err = json.Unmarshal(*valMajorUnitScale, &valueForMajorUnitScaleInt)
				if err != nil {
					return err
				}
				this.MajorUnitScale = string(valueForMajorUnitScaleInt)
			} else {
				this.MajorUnitScale = valueForMajorUnitScale
			}
		}
	}
	if valMajorUnitScaleCap, ok := objMap["MajorUnitScale"]; ok {
		if valMajorUnitScaleCap != nil {
			var valueForMajorUnitScale string
			err = json.Unmarshal(*valMajorUnitScaleCap, &valueForMajorUnitScale)
			if err != nil {
				var valueForMajorUnitScaleInt int32
				err = json.Unmarshal(*valMajorUnitScaleCap, &valueForMajorUnitScaleInt)
				if err != nil {
					return err
				}
				this.MajorUnitScale = string(valueForMajorUnitScaleInt)
			} else {
				this.MajorUnitScale = valueForMajorUnitScale
			}
		}
	}
	
	if valMajorTickMark, ok := objMap["majorTickMark"]; ok {
		if valMajorTickMark != nil {
			var valueForMajorTickMark string
			err = json.Unmarshal(*valMajorTickMark, &valueForMajorTickMark)
			if err != nil {
				var valueForMajorTickMarkInt int32
				err = json.Unmarshal(*valMajorTickMark, &valueForMajorTickMarkInt)
				if err != nil {
					return err
				}
				this.MajorTickMark = string(valueForMajorTickMarkInt)
			} else {
				this.MajorTickMark = valueForMajorTickMark
			}
		}
	}
	if valMajorTickMarkCap, ok := objMap["MajorTickMark"]; ok {
		if valMajorTickMarkCap != nil {
			var valueForMajorTickMark string
			err = json.Unmarshal(*valMajorTickMarkCap, &valueForMajorTickMark)
			if err != nil {
				var valueForMajorTickMarkInt int32
				err = json.Unmarshal(*valMajorTickMarkCap, &valueForMajorTickMarkInt)
				if err != nil {
					return err
				}
				this.MajorTickMark = string(valueForMajorTickMarkInt)
			} else {
				this.MajorTickMark = valueForMajorTickMark
			}
		}
	}
	
	if valIsAutomaticMinorUnit, ok := objMap["isAutomaticMinorUnit"]; ok {
		if valIsAutomaticMinorUnit != nil {
			var valueForIsAutomaticMinorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMinorUnit, &valueForIsAutomaticMinorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMinorUnit = valueForIsAutomaticMinorUnit
		}
	}
	if valIsAutomaticMinorUnitCap, ok := objMap["IsAutomaticMinorUnit"]; ok {
		if valIsAutomaticMinorUnitCap != nil {
			var valueForIsAutomaticMinorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMinorUnitCap, &valueForIsAutomaticMinorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMinorUnit = valueForIsAutomaticMinorUnit
		}
	}
	
	if valMinorUnit, ok := objMap["minorUnit"]; ok {
		if valMinorUnit != nil {
			var valueForMinorUnit float64
			err = json.Unmarshal(*valMinorUnit, &valueForMinorUnit)
			if err != nil {
				return err
			}
			this.MinorUnit = valueForMinorUnit
		}
	}
	if valMinorUnitCap, ok := objMap["MinorUnit"]; ok {
		if valMinorUnitCap != nil {
			var valueForMinorUnit float64
			err = json.Unmarshal(*valMinorUnitCap, &valueForMinorUnit)
			if err != nil {
				return err
			}
			this.MinorUnit = valueForMinorUnit
		}
	}
	
	if valMinorUnitScale, ok := objMap["minorUnitScale"]; ok {
		if valMinorUnitScale != nil {
			var valueForMinorUnitScale string
			err = json.Unmarshal(*valMinorUnitScale, &valueForMinorUnitScale)
			if err != nil {
				var valueForMinorUnitScaleInt int32
				err = json.Unmarshal(*valMinorUnitScale, &valueForMinorUnitScaleInt)
				if err != nil {
					return err
				}
				this.MinorUnitScale = string(valueForMinorUnitScaleInt)
			} else {
				this.MinorUnitScale = valueForMinorUnitScale
			}
		}
	}
	if valMinorUnitScaleCap, ok := objMap["MinorUnitScale"]; ok {
		if valMinorUnitScaleCap != nil {
			var valueForMinorUnitScale string
			err = json.Unmarshal(*valMinorUnitScaleCap, &valueForMinorUnitScale)
			if err != nil {
				var valueForMinorUnitScaleInt int32
				err = json.Unmarshal(*valMinorUnitScaleCap, &valueForMinorUnitScaleInt)
				if err != nil {
					return err
				}
				this.MinorUnitScale = string(valueForMinorUnitScaleInt)
			} else {
				this.MinorUnitScale = valueForMinorUnitScale
			}
		}
	}
	
	if valMinorTickMark, ok := objMap["minorTickMark"]; ok {
		if valMinorTickMark != nil {
			var valueForMinorTickMark string
			err = json.Unmarshal(*valMinorTickMark, &valueForMinorTickMark)
			if err != nil {
				var valueForMinorTickMarkInt int32
				err = json.Unmarshal(*valMinorTickMark, &valueForMinorTickMarkInt)
				if err != nil {
					return err
				}
				this.MinorTickMark = string(valueForMinorTickMarkInt)
			} else {
				this.MinorTickMark = valueForMinorTickMark
			}
		}
	}
	if valMinorTickMarkCap, ok := objMap["MinorTickMark"]; ok {
		if valMinorTickMarkCap != nil {
			var valueForMinorTickMark string
			err = json.Unmarshal(*valMinorTickMarkCap, &valueForMinorTickMark)
			if err != nil {
				var valueForMinorTickMarkInt int32
				err = json.Unmarshal(*valMinorTickMarkCap, &valueForMinorTickMarkInt)
				if err != nil {
					return err
				}
				this.MinorTickMark = string(valueForMinorTickMarkInt)
			} else {
				this.MinorTickMark = valueForMinorTickMark
			}
		}
	}
	
	if valIsAutomaticMaxValue, ok := objMap["isAutomaticMaxValue"]; ok {
		if valIsAutomaticMaxValue != nil {
			var valueForIsAutomaticMaxValue *bool
			err = json.Unmarshal(*valIsAutomaticMaxValue, &valueForIsAutomaticMaxValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMaxValue = valueForIsAutomaticMaxValue
		}
	}
	if valIsAutomaticMaxValueCap, ok := objMap["IsAutomaticMaxValue"]; ok {
		if valIsAutomaticMaxValueCap != nil {
			var valueForIsAutomaticMaxValue *bool
			err = json.Unmarshal(*valIsAutomaticMaxValueCap, &valueForIsAutomaticMaxValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMaxValue = valueForIsAutomaticMaxValue
		}
	}
	
	if valMaxValue, ok := objMap["maxValue"]; ok {
		if valMaxValue != nil {
			var valueForMaxValue float64
			err = json.Unmarshal(*valMaxValue, &valueForMaxValue)
			if err != nil {
				return err
			}
			this.MaxValue = valueForMaxValue
		}
	}
	if valMaxValueCap, ok := objMap["MaxValue"]; ok {
		if valMaxValueCap != nil {
			var valueForMaxValue float64
			err = json.Unmarshal(*valMaxValueCap, &valueForMaxValue)
			if err != nil {
				return err
			}
			this.MaxValue = valueForMaxValue
		}
	}
	
	if valIsAutomaticMinValue, ok := objMap["isAutomaticMinValue"]; ok {
		if valIsAutomaticMinValue != nil {
			var valueForIsAutomaticMinValue *bool
			err = json.Unmarshal(*valIsAutomaticMinValue, &valueForIsAutomaticMinValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMinValue = valueForIsAutomaticMinValue
		}
	}
	if valIsAutomaticMinValueCap, ok := objMap["IsAutomaticMinValue"]; ok {
		if valIsAutomaticMinValueCap != nil {
			var valueForIsAutomaticMinValue *bool
			err = json.Unmarshal(*valIsAutomaticMinValueCap, &valueForIsAutomaticMinValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMinValue = valueForIsAutomaticMinValue
		}
	}
	
	if valMinValue, ok := objMap["minValue"]; ok {
		if valMinValue != nil {
			var valueForMinValue float64
			err = json.Unmarshal(*valMinValue, &valueForMinValue)
			if err != nil {
				return err
			}
			this.MinValue = valueForMinValue
		}
	}
	if valMinValueCap, ok := objMap["MinValue"]; ok {
		if valMinValueCap != nil {
			var valueForMinValue float64
			err = json.Unmarshal(*valMinValueCap, &valueForMinValue)
			if err != nil {
				return err
			}
			this.MinValue = valueForMinValue
		}
	}
	
	if valIsLogarithmic, ok := objMap["isLogarithmic"]; ok {
		if valIsLogarithmic != nil {
			var valueForIsLogarithmic *bool
			err = json.Unmarshal(*valIsLogarithmic, &valueForIsLogarithmic)
			if err != nil {
				return err
			}
			this.IsLogarithmic = valueForIsLogarithmic
		}
	}
	if valIsLogarithmicCap, ok := objMap["IsLogarithmic"]; ok {
		if valIsLogarithmicCap != nil {
			var valueForIsLogarithmic *bool
			err = json.Unmarshal(*valIsLogarithmicCap, &valueForIsLogarithmic)
			if err != nil {
				return err
			}
			this.IsLogarithmic = valueForIsLogarithmic
		}
	}
	
	if valLogBase, ok := objMap["logBase"]; ok {
		if valLogBase != nil {
			var valueForLogBase float64
			err = json.Unmarshal(*valLogBase, &valueForLogBase)
			if err != nil {
				return err
			}
			this.LogBase = valueForLogBase
		}
	}
	if valLogBaseCap, ok := objMap["LogBase"]; ok {
		if valLogBaseCap != nil {
			var valueForLogBase float64
			err = json.Unmarshal(*valLogBaseCap, &valueForLogBase)
			if err != nil {
				return err
			}
			this.LogBase = valueForLogBase
		}
	}
	
	if valCategoryAxisType, ok := objMap["categoryAxisType"]; ok {
		if valCategoryAxisType != nil {
			var valueForCategoryAxisType string
			err = json.Unmarshal(*valCategoryAxisType, &valueForCategoryAxisType)
			if err != nil {
				var valueForCategoryAxisTypeInt int32
				err = json.Unmarshal(*valCategoryAxisType, &valueForCategoryAxisTypeInt)
				if err != nil {
					return err
				}
				this.CategoryAxisType = string(valueForCategoryAxisTypeInt)
			} else {
				this.CategoryAxisType = valueForCategoryAxisType
			}
		}
	}
	if valCategoryAxisTypeCap, ok := objMap["CategoryAxisType"]; ok {
		if valCategoryAxisTypeCap != nil {
			var valueForCategoryAxisType string
			err = json.Unmarshal(*valCategoryAxisTypeCap, &valueForCategoryAxisType)
			if err != nil {
				var valueForCategoryAxisTypeInt int32
				err = json.Unmarshal(*valCategoryAxisTypeCap, &valueForCategoryAxisTypeInt)
				if err != nil {
					return err
				}
				this.CategoryAxisType = string(valueForCategoryAxisTypeInt)
			} else {
				this.CategoryAxisType = valueForCategoryAxisType
			}
		}
	}
	
	if valAxisBetweenCategories, ok := objMap["axisBetweenCategories"]; ok {
		if valAxisBetweenCategories != nil {
			var valueForAxisBetweenCategories *bool
			err = json.Unmarshal(*valAxisBetweenCategories, &valueForAxisBetweenCategories)
			if err != nil {
				return err
			}
			this.AxisBetweenCategories = valueForAxisBetweenCategories
		}
	}
	if valAxisBetweenCategoriesCap, ok := objMap["AxisBetweenCategories"]; ok {
		if valAxisBetweenCategoriesCap != nil {
			var valueForAxisBetweenCategories *bool
			err = json.Unmarshal(*valAxisBetweenCategoriesCap, &valueForAxisBetweenCategories)
			if err != nil {
				return err
			}
			this.AxisBetweenCategories = valueForAxisBetweenCategories
		}
	}
	
	if valLabelOffset, ok := objMap["labelOffset"]; ok {
		if valLabelOffset != nil {
			var valueForLabelOffset int32
			err = json.Unmarshal(*valLabelOffset, &valueForLabelOffset)
			if err != nil {
				return err
			}
			this.LabelOffset = valueForLabelOffset
		}
	}
	if valLabelOffsetCap, ok := objMap["LabelOffset"]; ok {
		if valLabelOffsetCap != nil {
			var valueForLabelOffset int32
			err = json.Unmarshal(*valLabelOffsetCap, &valueForLabelOffset)
			if err != nil {
				return err
			}
			this.LabelOffset = valueForLabelOffset
		}
	}
	
	if valIsPlotOrderReversed, ok := objMap["isPlotOrderReversed"]; ok {
		if valIsPlotOrderReversed != nil {
			var valueForIsPlotOrderReversed *bool
			err = json.Unmarshal(*valIsPlotOrderReversed, &valueForIsPlotOrderReversed)
			if err != nil {
				return err
			}
			this.IsPlotOrderReversed = valueForIsPlotOrderReversed
		}
	}
	if valIsPlotOrderReversedCap, ok := objMap["IsPlotOrderReversed"]; ok {
		if valIsPlotOrderReversedCap != nil {
			var valueForIsPlotOrderReversed *bool
			err = json.Unmarshal(*valIsPlotOrderReversedCap, &valueForIsPlotOrderReversed)
			if err != nil {
				return err
			}
			this.IsPlotOrderReversed = valueForIsPlotOrderReversed
		}
	}
	
	if valIsNumberFormatLinkedToSource, ok := objMap["isNumberFormatLinkedToSource"]; ok {
		if valIsNumberFormatLinkedToSource != nil {
			var valueForIsNumberFormatLinkedToSource *bool
			err = json.Unmarshal(*valIsNumberFormatLinkedToSource, &valueForIsNumberFormatLinkedToSource)
			if err != nil {
				return err
			}
			this.IsNumberFormatLinkedToSource = valueForIsNumberFormatLinkedToSource
		}
	}
	if valIsNumberFormatLinkedToSourceCap, ok := objMap["IsNumberFormatLinkedToSource"]; ok {
		if valIsNumberFormatLinkedToSourceCap != nil {
			var valueForIsNumberFormatLinkedToSource *bool
			err = json.Unmarshal(*valIsNumberFormatLinkedToSourceCap, &valueForIsNumberFormatLinkedToSource)
			if err != nil {
				return err
			}
			this.IsNumberFormatLinkedToSource = valueForIsNumberFormatLinkedToSource
		}
	}
	
	if valNumberFormat, ok := objMap["numberFormat"]; ok {
		if valNumberFormat != nil {
			var valueForNumberFormat string
			err = json.Unmarshal(*valNumberFormat, &valueForNumberFormat)
			if err != nil {
				return err
			}
			this.NumberFormat = valueForNumberFormat
		}
	}
	if valNumberFormatCap, ok := objMap["NumberFormat"]; ok {
		if valNumberFormatCap != nil {
			var valueForNumberFormat string
			err = json.Unmarshal(*valNumberFormatCap, &valueForNumberFormat)
			if err != nil {
				return err
			}
			this.NumberFormat = valueForNumberFormat
		}
	}
	
	if valCrossType, ok := objMap["crossType"]; ok {
		if valCrossType != nil {
			var valueForCrossType string
			err = json.Unmarshal(*valCrossType, &valueForCrossType)
			if err != nil {
				var valueForCrossTypeInt int32
				err = json.Unmarshal(*valCrossType, &valueForCrossTypeInt)
				if err != nil {
					return err
				}
				this.CrossType = string(valueForCrossTypeInt)
			} else {
				this.CrossType = valueForCrossType
			}
		}
	}
	if valCrossTypeCap, ok := objMap["CrossType"]; ok {
		if valCrossTypeCap != nil {
			var valueForCrossType string
			err = json.Unmarshal(*valCrossTypeCap, &valueForCrossType)
			if err != nil {
				var valueForCrossTypeInt int32
				err = json.Unmarshal(*valCrossTypeCap, &valueForCrossTypeInt)
				if err != nil {
					return err
				}
				this.CrossType = string(valueForCrossTypeInt)
			} else {
				this.CrossType = valueForCrossType
			}
		}
	}
	
	if valCrossAt, ok := objMap["crossAt"]; ok {
		if valCrossAt != nil {
			var valueForCrossAt float64
			err = json.Unmarshal(*valCrossAt, &valueForCrossAt)
			if err != nil {
				return err
			}
			this.CrossAt = valueForCrossAt
		}
	}
	if valCrossAtCap, ok := objMap["CrossAt"]; ok {
		if valCrossAtCap != nil {
			var valueForCrossAt float64
			err = json.Unmarshal(*valCrossAtCap, &valueForCrossAt)
			if err != nil {
				return err
			}
			this.CrossAt = valueForCrossAt
		}
	}
	
	if valIsAutomaticTickMarksSpacing, ok := objMap["isAutomaticTickMarksSpacing"]; ok {
		if valIsAutomaticTickMarksSpacing != nil {
			var valueForIsAutomaticTickMarksSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickMarksSpacing, &valueForIsAutomaticTickMarksSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickMarksSpacing = valueForIsAutomaticTickMarksSpacing
		}
	}
	if valIsAutomaticTickMarksSpacingCap, ok := objMap["IsAutomaticTickMarksSpacing"]; ok {
		if valIsAutomaticTickMarksSpacingCap != nil {
			var valueForIsAutomaticTickMarksSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickMarksSpacingCap, &valueForIsAutomaticTickMarksSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickMarksSpacing = valueForIsAutomaticTickMarksSpacing
		}
	}
	
	if valTickMarksSpacing, ok := objMap["tickMarksSpacing"]; ok {
		if valTickMarksSpacing != nil {
			var valueForTickMarksSpacing int32
			err = json.Unmarshal(*valTickMarksSpacing, &valueForTickMarksSpacing)
			if err != nil {
				return err
			}
			this.TickMarksSpacing = valueForTickMarksSpacing
		}
	}
	if valTickMarksSpacingCap, ok := objMap["TickMarksSpacing"]; ok {
		if valTickMarksSpacingCap != nil {
			var valueForTickMarksSpacing int32
			err = json.Unmarshal(*valTickMarksSpacingCap, &valueForTickMarksSpacing)
			if err != nil {
				return err
			}
			this.TickMarksSpacing = valueForTickMarksSpacing
		}
	}
	
	if valIsAutomaticTickLabelSpacing, ok := objMap["isAutomaticTickLabelSpacing"]; ok {
		if valIsAutomaticTickLabelSpacing != nil {
			var valueForIsAutomaticTickLabelSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickLabelSpacing, &valueForIsAutomaticTickLabelSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickLabelSpacing = valueForIsAutomaticTickLabelSpacing
		}
	}
	if valIsAutomaticTickLabelSpacingCap, ok := objMap["IsAutomaticTickLabelSpacing"]; ok {
		if valIsAutomaticTickLabelSpacingCap != nil {
			var valueForIsAutomaticTickLabelSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickLabelSpacingCap, &valueForIsAutomaticTickLabelSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickLabelSpacing = valueForIsAutomaticTickLabelSpacing
		}
	}
	
	if valTickLabelSpacing, ok := objMap["tickLabelSpacing"]; ok {
		if valTickLabelSpacing != nil {
			var valueForTickLabelSpacing int32
			err = json.Unmarshal(*valTickLabelSpacing, &valueForTickLabelSpacing)
			if err != nil {
				return err
			}
			this.TickLabelSpacing = valueForTickLabelSpacing
		}
	}
	if valTickLabelSpacingCap, ok := objMap["TickLabelSpacing"]; ok {
		if valTickLabelSpacingCap != nil {
			var valueForTickLabelSpacing int32
			err = json.Unmarshal(*valTickLabelSpacingCap, &valueForTickLabelSpacing)
			if err != nil {
				return err
			}
			this.TickLabelSpacing = valueForTickLabelSpacing
		}
	}
	
	if valTickLabelPosition, ok := objMap["tickLabelPosition"]; ok {
		if valTickLabelPosition != nil {
			var valueForTickLabelPosition string
			err = json.Unmarshal(*valTickLabelPosition, &valueForTickLabelPosition)
			if err != nil {
				var valueForTickLabelPositionInt int32
				err = json.Unmarshal(*valTickLabelPosition, &valueForTickLabelPositionInt)
				if err != nil {
					return err
				}
				this.TickLabelPosition = string(valueForTickLabelPositionInt)
			} else {
				this.TickLabelPosition = valueForTickLabelPosition
			}
		}
	}
	if valTickLabelPositionCap, ok := objMap["TickLabelPosition"]; ok {
		if valTickLabelPositionCap != nil {
			var valueForTickLabelPosition string
			err = json.Unmarshal(*valTickLabelPositionCap, &valueForTickLabelPosition)
			if err != nil {
				var valueForTickLabelPositionInt int32
				err = json.Unmarshal(*valTickLabelPositionCap, &valueForTickLabelPositionInt)
				if err != nil {
					return err
				}
				this.TickLabelPosition = string(valueForTickLabelPositionInt)
			} else {
				this.TickLabelPosition = valueForTickLabelPosition
			}
		}
	}
	
	if valTickLabelRotationAngle, ok := objMap["tickLabelRotationAngle"]; ok {
		if valTickLabelRotationAngle != nil {
			var valueForTickLabelRotationAngle float64
			err = json.Unmarshal(*valTickLabelRotationAngle, &valueForTickLabelRotationAngle)
			if err != nil {
				return err
			}
			this.TickLabelRotationAngle = valueForTickLabelRotationAngle
		}
	}
	if valTickLabelRotationAngleCap, ok := objMap["TickLabelRotationAngle"]; ok {
		if valTickLabelRotationAngleCap != nil {
			var valueForTickLabelRotationAngle float64
			err = json.Unmarshal(*valTickLabelRotationAngleCap, &valueForTickLabelRotationAngle)
			if err != nil {
				return err
			}
			this.TickLabelRotationAngle = valueForTickLabelRotationAngle
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
	
	if valMajorGridLinesFormat, ok := objMap["majorGridLinesFormat"]; ok {
		if valMajorGridLinesFormat != nil {
			var valueForMajorGridLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valMajorGridLinesFormat, &valueForMajorGridLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valMajorGridLinesFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMajorGridLinesFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.MajorGridLinesFormat = vInterfaceObject
			}
		}
	}
	if valMajorGridLinesFormatCap, ok := objMap["MajorGridLinesFormat"]; ok {
		if valMajorGridLinesFormatCap != nil {
			var valueForMajorGridLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valMajorGridLinesFormatCap, &valueForMajorGridLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valMajorGridLinesFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMajorGridLinesFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.MajorGridLinesFormat = vInterfaceObject
			}
		}
	}
	
	if valMinorGridLinesFormat, ok := objMap["minorGridLinesFormat"]; ok {
		if valMinorGridLinesFormat != nil {
			var valueForMinorGridLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valMinorGridLinesFormat, &valueForMinorGridLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valMinorGridLinesFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMinorGridLinesFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.MinorGridLinesFormat = vInterfaceObject
			}
		}
	}
	if valMinorGridLinesFormatCap, ok := objMap["MinorGridLinesFormat"]; ok {
		if valMinorGridLinesFormatCap != nil {
			var valueForMinorGridLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valMinorGridLinesFormatCap, &valueForMinorGridLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valMinorGridLinesFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMinorGridLinesFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.MinorGridLinesFormat = vInterfaceObject
			}
		}
	}

	return nil
}
