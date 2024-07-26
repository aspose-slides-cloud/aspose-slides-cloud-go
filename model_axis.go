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
	
	if valIsVisible, ok := GetMapValue(objMap, "isVisible"); ok {
		if valIsVisible != nil {
			var valueForIsVisible *bool
			err = json.Unmarshal(*valIsVisible, &valueForIsVisible)
			if err != nil {
				return err
			}
			this.IsVisible = valueForIsVisible
		}
	}
	
	if valHasTitle, ok := GetMapValue(objMap, "hasTitle"); ok {
		if valHasTitle != nil {
			var valueForHasTitle *bool
			err = json.Unmarshal(*valHasTitle, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}
	
	if valTitle, ok := GetMapValue(objMap, "title"); ok {
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
	
	if valPosition, ok := GetMapValue(objMap, "position"); ok {
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
	
	if valDisplayUnit, ok := GetMapValue(objMap, "displayUnit"); ok {
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
	
	if valBaseUnitScale, ok := GetMapValue(objMap, "baseUnitScale"); ok {
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
	
	if valIsAutomaticMajorUnit, ok := GetMapValue(objMap, "isAutomaticMajorUnit"); ok {
		if valIsAutomaticMajorUnit != nil {
			var valueForIsAutomaticMajorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMajorUnit, &valueForIsAutomaticMajorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMajorUnit = valueForIsAutomaticMajorUnit
		}
	}
	
	if valMajorUnit, ok := GetMapValue(objMap, "majorUnit"); ok {
		if valMajorUnit != nil {
			var valueForMajorUnit float64
			err = json.Unmarshal(*valMajorUnit, &valueForMajorUnit)
			if err != nil {
				return err
			}
			this.MajorUnit = valueForMajorUnit
		}
	}
	
	if valMajorUnitScale, ok := GetMapValue(objMap, "majorUnitScale"); ok {
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
	
	if valMajorTickMark, ok := GetMapValue(objMap, "majorTickMark"); ok {
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
	
	if valIsAutomaticMinorUnit, ok := GetMapValue(objMap, "isAutomaticMinorUnit"); ok {
		if valIsAutomaticMinorUnit != nil {
			var valueForIsAutomaticMinorUnit *bool
			err = json.Unmarshal(*valIsAutomaticMinorUnit, &valueForIsAutomaticMinorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMinorUnit = valueForIsAutomaticMinorUnit
		}
	}
	
	if valMinorUnit, ok := GetMapValue(objMap, "minorUnit"); ok {
		if valMinorUnit != nil {
			var valueForMinorUnit float64
			err = json.Unmarshal(*valMinorUnit, &valueForMinorUnit)
			if err != nil {
				return err
			}
			this.MinorUnit = valueForMinorUnit
		}
	}
	
	if valMinorUnitScale, ok := GetMapValue(objMap, "minorUnitScale"); ok {
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
	
	if valMinorTickMark, ok := GetMapValue(objMap, "minorTickMark"); ok {
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
	
	if valIsAutomaticMaxValue, ok := GetMapValue(objMap, "isAutomaticMaxValue"); ok {
		if valIsAutomaticMaxValue != nil {
			var valueForIsAutomaticMaxValue *bool
			err = json.Unmarshal(*valIsAutomaticMaxValue, &valueForIsAutomaticMaxValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMaxValue = valueForIsAutomaticMaxValue
		}
	}
	
	if valMaxValue, ok := GetMapValue(objMap, "maxValue"); ok {
		if valMaxValue != nil {
			var valueForMaxValue float64
			err = json.Unmarshal(*valMaxValue, &valueForMaxValue)
			if err != nil {
				return err
			}
			this.MaxValue = valueForMaxValue
		}
	}
	
	if valIsAutomaticMinValue, ok := GetMapValue(objMap, "isAutomaticMinValue"); ok {
		if valIsAutomaticMinValue != nil {
			var valueForIsAutomaticMinValue *bool
			err = json.Unmarshal(*valIsAutomaticMinValue, &valueForIsAutomaticMinValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMinValue = valueForIsAutomaticMinValue
		}
	}
	
	if valMinValue, ok := GetMapValue(objMap, "minValue"); ok {
		if valMinValue != nil {
			var valueForMinValue float64
			err = json.Unmarshal(*valMinValue, &valueForMinValue)
			if err != nil {
				return err
			}
			this.MinValue = valueForMinValue
		}
	}
	
	if valIsLogarithmic, ok := GetMapValue(objMap, "isLogarithmic"); ok {
		if valIsLogarithmic != nil {
			var valueForIsLogarithmic *bool
			err = json.Unmarshal(*valIsLogarithmic, &valueForIsLogarithmic)
			if err != nil {
				return err
			}
			this.IsLogarithmic = valueForIsLogarithmic
		}
	}
	
	if valLogBase, ok := GetMapValue(objMap, "logBase"); ok {
		if valLogBase != nil {
			var valueForLogBase float64
			err = json.Unmarshal(*valLogBase, &valueForLogBase)
			if err != nil {
				return err
			}
			this.LogBase = valueForLogBase
		}
	}
	
	if valCategoryAxisType, ok := GetMapValue(objMap, "categoryAxisType"); ok {
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
	
	if valAxisBetweenCategories, ok := GetMapValue(objMap, "axisBetweenCategories"); ok {
		if valAxisBetweenCategories != nil {
			var valueForAxisBetweenCategories *bool
			err = json.Unmarshal(*valAxisBetweenCategories, &valueForAxisBetweenCategories)
			if err != nil {
				return err
			}
			this.AxisBetweenCategories = valueForAxisBetweenCategories
		}
	}
	
	if valLabelOffset, ok := GetMapValue(objMap, "labelOffset"); ok {
		if valLabelOffset != nil {
			var valueForLabelOffset int32
			err = json.Unmarshal(*valLabelOffset, &valueForLabelOffset)
			if err != nil {
				return err
			}
			this.LabelOffset = valueForLabelOffset
		}
	}
	
	if valIsPlotOrderReversed, ok := GetMapValue(objMap, "isPlotOrderReversed"); ok {
		if valIsPlotOrderReversed != nil {
			var valueForIsPlotOrderReversed *bool
			err = json.Unmarshal(*valIsPlotOrderReversed, &valueForIsPlotOrderReversed)
			if err != nil {
				return err
			}
			this.IsPlotOrderReversed = valueForIsPlotOrderReversed
		}
	}
	
	if valIsNumberFormatLinkedToSource, ok := GetMapValue(objMap, "isNumberFormatLinkedToSource"); ok {
		if valIsNumberFormatLinkedToSource != nil {
			var valueForIsNumberFormatLinkedToSource *bool
			err = json.Unmarshal(*valIsNumberFormatLinkedToSource, &valueForIsNumberFormatLinkedToSource)
			if err != nil {
				return err
			}
			this.IsNumberFormatLinkedToSource = valueForIsNumberFormatLinkedToSource
		}
	}
	
	if valNumberFormat, ok := GetMapValue(objMap, "numberFormat"); ok {
		if valNumberFormat != nil {
			var valueForNumberFormat string
			err = json.Unmarshal(*valNumberFormat, &valueForNumberFormat)
			if err != nil {
				return err
			}
			this.NumberFormat = valueForNumberFormat
		}
	}
	
	if valCrossType, ok := GetMapValue(objMap, "crossType"); ok {
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
	
	if valCrossAt, ok := GetMapValue(objMap, "crossAt"); ok {
		if valCrossAt != nil {
			var valueForCrossAt float64
			err = json.Unmarshal(*valCrossAt, &valueForCrossAt)
			if err != nil {
				return err
			}
			this.CrossAt = valueForCrossAt
		}
	}
	
	if valIsAutomaticTickMarksSpacing, ok := GetMapValue(objMap, "isAutomaticTickMarksSpacing"); ok {
		if valIsAutomaticTickMarksSpacing != nil {
			var valueForIsAutomaticTickMarksSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickMarksSpacing, &valueForIsAutomaticTickMarksSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickMarksSpacing = valueForIsAutomaticTickMarksSpacing
		}
	}
	
	if valTickMarksSpacing, ok := GetMapValue(objMap, "tickMarksSpacing"); ok {
		if valTickMarksSpacing != nil {
			var valueForTickMarksSpacing int32
			err = json.Unmarshal(*valTickMarksSpacing, &valueForTickMarksSpacing)
			if err != nil {
				return err
			}
			this.TickMarksSpacing = valueForTickMarksSpacing
		}
	}
	
	if valIsAutomaticTickLabelSpacing, ok := GetMapValue(objMap, "isAutomaticTickLabelSpacing"); ok {
		if valIsAutomaticTickLabelSpacing != nil {
			var valueForIsAutomaticTickLabelSpacing *bool
			err = json.Unmarshal(*valIsAutomaticTickLabelSpacing, &valueForIsAutomaticTickLabelSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickLabelSpacing = valueForIsAutomaticTickLabelSpacing
		}
	}
	
	if valTickLabelSpacing, ok := GetMapValue(objMap, "tickLabelSpacing"); ok {
		if valTickLabelSpacing != nil {
			var valueForTickLabelSpacing int32
			err = json.Unmarshal(*valTickLabelSpacing, &valueForTickLabelSpacing)
			if err != nil {
				return err
			}
			this.TickLabelSpacing = valueForTickLabelSpacing
		}
	}
	
	if valTickLabelPosition, ok := GetMapValue(objMap, "tickLabelPosition"); ok {
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
	
	if valTickLabelRotationAngle, ok := GetMapValue(objMap, "tickLabelRotationAngle"); ok {
		if valTickLabelRotationAngle != nil {
			var valueForTickLabelRotationAngle float64
			err = json.Unmarshal(*valTickLabelRotationAngle, &valueForTickLabelRotationAngle)
			if err != nil {
				return err
			}
			this.TickLabelRotationAngle = valueForTickLabelRotationAngle
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
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
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
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
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
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
	
	if valMajorGridLinesFormat, ok := GetMapValue(objMap, "majorGridLinesFormat"); ok {
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
	
	if valMinorGridLinesFormat, ok := GetMapValue(objMap, "minorGridLinesFormat"); ok {
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

	return nil
}
