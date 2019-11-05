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
	getIsVisible() bool
	setIsVisible(newValue bool)

	// True if the axis has a visible title
	getHasTitle() bool
	setHasTitle(newValue bool)

	// Axis position
	getPosition() string
	setPosition(newValue string)

	// The scaling value of the display units for the value axis
	getDisplayUnit() string
	setDisplayUnit(newValue string)

	// The smallest time unit that is represented on the date axis
	getBaseUnitScale() string
	setBaseUnitScale(newValue string)

	// True the major unit of the axis is automatically assigned
	getIsAutomaticMajorUnit() bool
	setIsAutomaticMajorUnit(newValue bool)

	// The major units for the date or value axis
	getMajorUnit() float64
	setMajorUnit(newValue float64)

	// The major unit scale for the date axis
	getMajorUnitScale() string
	setMajorUnitScale(newValue string)

	// The type of major tick mark for the specified axis
	getMajorTickMark() string
	setMajorTickMark(newValue string)

	// True the minor unit of the axis is automatically assigned
	getIsAutomaticMinorUnit() bool
	setIsAutomaticMinorUnit(newValue bool)

	// The minor units for the date or value axis
	getMinorUnit() float64
	setMinorUnit(newValue float64)

	// The minor unit scale for the date axis
	getMinorUnitScale() string
	setMinorUnitScale(newValue string)

	// The type of minor tick mark for the specified axis
	getMinorTickMark() string
	setMinorTickMark(newValue string)

	// True if the max value is automatically assigned
	getIsAutomaticMaxValue() bool
	setIsAutomaticMaxValue(newValue bool)

	// The maximum value on the value axis
	getMaxValue() float64
	setMaxValue(newValue float64)

	// True if the min value is automatically assigned
	getIsAutomaticMinValue() bool
	setIsAutomaticMinValue(newValue bool)

	// The minimum value on the value axis
	getMinValue() float64
	setMinValue(newValue float64)

	// True if the value axis scale type is logarithmic
	getIsLogarithmic() bool
	setIsLogarithmic(newValue bool)

	// The logarithmic base. Default value is 10
	getLogBase() float64
	setLogBase(newValue float64)

	// The type of the category axis
	getCategoryAxisType() string
	setCategoryAxisType(newValue string)

	// True if the value axis crosses the category axis between categories. This property applies only to category axes, and it doesn't apply to 3-D charts
	getAxisBetweenCategories() bool
	setAxisBetweenCategories(newValue bool)

	// The distance of labels from the axis. Applied to category or date axis. Value must be between 0% and 1000%.             
	getLabelOffset() int32
	setLabelOffset(newValue int32)

	// True if MS PowerPoint plots data points from last to first
	getIsPlotOrderReversed() bool
	setIsPlotOrderReversed(newValue bool)

	// True if the format is linked to source data
	getIsNumberFormatLinkedToSource() bool
	setIsNumberFormatLinkedToSource(newValue bool)

	// the format string for the Axis Labels
	getNumberFormat() string
	setNumberFormat(newValue string)

	// The CrossType on the specified axis where the other axis crosses
	getCrossType() string
	setCrossType(newValue string)

	// The point on the axis where the perpendicular axis crosses it
	getCrossAt() float64
	setCrossAt(newValue float64)

	// True for automatic tick marks spacing value
	getIsAutomaticTickMarksSpacing() bool
	setIsAutomaticTickMarksSpacing(newValue bool)

	// Specifies how many tick marks shall be skipped before the next one shall be drawn. Applied to category or series axis.
	getTickMarksSpacing() int32
	setTickMarksSpacing(newValue int32)

	// True for automatic tick label spacing value
	getIsAutomaticTickLabelSpacing() bool
	setIsAutomaticTickLabelSpacing(newValue bool)

	// Specifies how many tick labels to skip between label that is drawn.
	getTickLabelSpacing() int32
	setTickLabelSpacing(newValue int32)

	// The position of tick-mark labels on the specified axis.
	getTickLabelPosition() string
	setTickLabelPosition(newValue string)

	// Represents the rotation angle of tick labels.
	getTickLabelRotationAngle() float64
	setTickLabelRotationAngle(newValue float64)

	// Get or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)
}

type Axis struct {

	// True if the axis is visible
	IsVisible bool `json:"IsVisible"`

	// True if the axis has a visible title
	HasTitle bool `json:"HasTitle"`

	// Axis position
	Position string `json:"Position"`

	// The scaling value of the display units for the value axis
	DisplayUnit string `json:"DisplayUnit"`

	// The smallest time unit that is represented on the date axis
	BaseUnitScale string `json:"BaseUnitScale"`

	// True the major unit of the axis is automatically assigned
	IsAutomaticMajorUnit bool `json:"IsAutomaticMajorUnit"`

	// The major units for the date or value axis
	MajorUnit float64 `json:"MajorUnit"`

	// The major unit scale for the date axis
	MajorUnitScale string `json:"MajorUnitScale"`

	// The type of major tick mark for the specified axis
	MajorTickMark string `json:"MajorTickMark"`

	// True the minor unit of the axis is automatically assigned
	IsAutomaticMinorUnit bool `json:"IsAutomaticMinorUnit"`

	// The minor units for the date or value axis
	MinorUnit float64 `json:"MinorUnit"`

	// The minor unit scale for the date axis
	MinorUnitScale string `json:"MinorUnitScale"`

	// The type of minor tick mark for the specified axis
	MinorTickMark string `json:"MinorTickMark"`

	// True if the max value is automatically assigned
	IsAutomaticMaxValue bool `json:"IsAutomaticMaxValue"`

	// The maximum value on the value axis
	MaxValue float64 `json:"MaxValue"`

	// True if the min value is automatically assigned
	IsAutomaticMinValue bool `json:"IsAutomaticMinValue"`

	// The minimum value on the value axis
	MinValue float64 `json:"MinValue"`

	// True if the value axis scale type is logarithmic
	IsLogarithmic bool `json:"IsLogarithmic"`

	// The logarithmic base. Default value is 10
	LogBase float64 `json:"LogBase"`

	// The type of the category axis
	CategoryAxisType string `json:"CategoryAxisType"`

	// True if the value axis crosses the category axis between categories. This property applies only to category axes, and it doesn't apply to 3-D charts
	AxisBetweenCategories bool `json:"AxisBetweenCategories"`

	// The distance of labels from the axis. Applied to category or date axis. Value must be between 0% and 1000%.             
	LabelOffset int32 `json:"LabelOffset"`

	// True if MS PowerPoint plots data points from last to first
	IsPlotOrderReversed bool `json:"IsPlotOrderReversed"`

	// True if the format is linked to source data
	IsNumberFormatLinkedToSource bool `json:"IsNumberFormatLinkedToSource"`

	// the format string for the Axis Labels
	NumberFormat string `json:"NumberFormat,omitempty"`

	// The CrossType on the specified axis where the other axis crosses
	CrossType string `json:"CrossType"`

	// The point on the axis where the perpendicular axis crosses it
	CrossAt float64 `json:"CrossAt"`

	// True for automatic tick marks spacing value
	IsAutomaticTickMarksSpacing bool `json:"IsAutomaticTickMarksSpacing"`

	// Specifies how many tick marks shall be skipped before the next one shall be drawn. Applied to category or series axis.
	TickMarksSpacing int32 `json:"TickMarksSpacing"`

	// True for automatic tick label spacing value
	IsAutomaticTickLabelSpacing bool `json:"IsAutomaticTickLabelSpacing"`

	// Specifies how many tick labels to skip between label that is drawn.
	TickLabelSpacing int32 `json:"TickLabelSpacing"`

	// The position of tick-mark labels on the specified axis.
	TickLabelPosition string `json:"TickLabelPosition"`

	// Represents the rotation angle of tick labels.
	TickLabelRotationAngle float64 `json:"TickLabelRotationAngle"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func (this Axis) getIsVisible() bool {
	return this.IsVisible
}

func (this Axis) setIsVisible(newValue bool) {
	this.IsVisible = newValue
}
func (this Axis) getHasTitle() bool {
	return this.HasTitle
}

func (this Axis) setHasTitle(newValue bool) {
	this.HasTitle = newValue
}
func (this Axis) getPosition() string {
	return this.Position
}

func (this Axis) setPosition(newValue string) {
	this.Position = newValue
}
func (this Axis) getDisplayUnit() string {
	return this.DisplayUnit
}

func (this Axis) setDisplayUnit(newValue string) {
	this.DisplayUnit = newValue
}
func (this Axis) getBaseUnitScale() string {
	return this.BaseUnitScale
}

func (this Axis) setBaseUnitScale(newValue string) {
	this.BaseUnitScale = newValue
}
func (this Axis) getIsAutomaticMajorUnit() bool {
	return this.IsAutomaticMajorUnit
}

func (this Axis) setIsAutomaticMajorUnit(newValue bool) {
	this.IsAutomaticMajorUnit = newValue
}
func (this Axis) getMajorUnit() float64 {
	return this.MajorUnit
}

func (this Axis) setMajorUnit(newValue float64) {
	this.MajorUnit = newValue
}
func (this Axis) getMajorUnitScale() string {
	return this.MajorUnitScale
}

func (this Axis) setMajorUnitScale(newValue string) {
	this.MajorUnitScale = newValue
}
func (this Axis) getMajorTickMark() string {
	return this.MajorTickMark
}

func (this Axis) setMajorTickMark(newValue string) {
	this.MajorTickMark = newValue
}
func (this Axis) getIsAutomaticMinorUnit() bool {
	return this.IsAutomaticMinorUnit
}

func (this Axis) setIsAutomaticMinorUnit(newValue bool) {
	this.IsAutomaticMinorUnit = newValue
}
func (this Axis) getMinorUnit() float64 {
	return this.MinorUnit
}

func (this Axis) setMinorUnit(newValue float64) {
	this.MinorUnit = newValue
}
func (this Axis) getMinorUnitScale() string {
	return this.MinorUnitScale
}

func (this Axis) setMinorUnitScale(newValue string) {
	this.MinorUnitScale = newValue
}
func (this Axis) getMinorTickMark() string {
	return this.MinorTickMark
}

func (this Axis) setMinorTickMark(newValue string) {
	this.MinorTickMark = newValue
}
func (this Axis) getIsAutomaticMaxValue() bool {
	return this.IsAutomaticMaxValue
}

func (this Axis) setIsAutomaticMaxValue(newValue bool) {
	this.IsAutomaticMaxValue = newValue
}
func (this Axis) getMaxValue() float64 {
	return this.MaxValue
}

func (this Axis) setMaxValue(newValue float64) {
	this.MaxValue = newValue
}
func (this Axis) getIsAutomaticMinValue() bool {
	return this.IsAutomaticMinValue
}

func (this Axis) setIsAutomaticMinValue(newValue bool) {
	this.IsAutomaticMinValue = newValue
}
func (this Axis) getMinValue() float64 {
	return this.MinValue
}

func (this Axis) setMinValue(newValue float64) {
	this.MinValue = newValue
}
func (this Axis) getIsLogarithmic() bool {
	return this.IsLogarithmic
}

func (this Axis) setIsLogarithmic(newValue bool) {
	this.IsLogarithmic = newValue
}
func (this Axis) getLogBase() float64 {
	return this.LogBase
}

func (this Axis) setLogBase(newValue float64) {
	this.LogBase = newValue
}
func (this Axis) getCategoryAxisType() string {
	return this.CategoryAxisType
}

func (this Axis) setCategoryAxisType(newValue string) {
	this.CategoryAxisType = newValue
}
func (this Axis) getAxisBetweenCategories() bool {
	return this.AxisBetweenCategories
}

func (this Axis) setAxisBetweenCategories(newValue bool) {
	this.AxisBetweenCategories = newValue
}
func (this Axis) getLabelOffset() int32 {
	return this.LabelOffset
}

func (this Axis) setLabelOffset(newValue int32) {
	this.LabelOffset = newValue
}
func (this Axis) getIsPlotOrderReversed() bool {
	return this.IsPlotOrderReversed
}

func (this Axis) setIsPlotOrderReversed(newValue bool) {
	this.IsPlotOrderReversed = newValue
}
func (this Axis) getIsNumberFormatLinkedToSource() bool {
	return this.IsNumberFormatLinkedToSource
}

func (this Axis) setIsNumberFormatLinkedToSource(newValue bool) {
	this.IsNumberFormatLinkedToSource = newValue
}
func (this Axis) getNumberFormat() string {
	return this.NumberFormat
}

func (this Axis) setNumberFormat(newValue string) {
	this.NumberFormat = newValue
}
func (this Axis) getCrossType() string {
	return this.CrossType
}

func (this Axis) setCrossType(newValue string) {
	this.CrossType = newValue
}
func (this Axis) getCrossAt() float64 {
	return this.CrossAt
}

func (this Axis) setCrossAt(newValue float64) {
	this.CrossAt = newValue
}
func (this Axis) getIsAutomaticTickMarksSpacing() bool {
	return this.IsAutomaticTickMarksSpacing
}

func (this Axis) setIsAutomaticTickMarksSpacing(newValue bool) {
	this.IsAutomaticTickMarksSpacing = newValue
}
func (this Axis) getTickMarksSpacing() int32 {
	return this.TickMarksSpacing
}

func (this Axis) setTickMarksSpacing(newValue int32) {
	this.TickMarksSpacing = newValue
}
func (this Axis) getIsAutomaticTickLabelSpacing() bool {
	return this.IsAutomaticTickLabelSpacing
}

func (this Axis) setIsAutomaticTickLabelSpacing(newValue bool) {
	this.IsAutomaticTickLabelSpacing = newValue
}
func (this Axis) getTickLabelSpacing() int32 {
	return this.TickLabelSpacing
}

func (this Axis) setTickLabelSpacing(newValue int32) {
	this.TickLabelSpacing = newValue
}
func (this Axis) getTickLabelPosition() string {
	return this.TickLabelPosition
}

func (this Axis) setTickLabelPosition(newValue string) {
	this.TickLabelPosition = newValue
}
func (this Axis) getTickLabelRotationAngle() float64 {
	return this.TickLabelRotationAngle
}

func (this Axis) setTickLabelRotationAngle(newValue float64) {
	this.TickLabelRotationAngle = newValue
}
func (this Axis) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Axis) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Axis) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Axis) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Axis) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Axis) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *Axis) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valIsVisible, ok := objMap["isVisible"]; ok {
		if valIsVisible != nil {
			var valueForIsVisible bool
			err = json.Unmarshal(*valIsVisible, &valueForIsVisible)
			if err != nil {
				return err
			}
			this.IsVisible = valueForIsVisible
		}
	}
	if valIsVisibleCap, ok := objMap["IsVisible"]; ok {
		if valIsVisibleCap != nil {
			var valueForIsVisible bool
			err = json.Unmarshal(*valIsVisibleCap, &valueForIsVisible)
			if err != nil {
				return err
			}
			this.IsVisible = valueForIsVisible
		}
	}
	
	if valHasTitle, ok := objMap["hasTitle"]; ok {
		if valHasTitle != nil {
			var valueForHasTitle bool
			err = json.Unmarshal(*valHasTitle, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}
	if valHasTitleCap, ok := objMap["HasTitle"]; ok {
		if valHasTitleCap != nil {
			var valueForHasTitle bool
			err = json.Unmarshal(*valHasTitleCap, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}
	this.Position = "Bottom"
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
	this.DisplayUnit = "None"
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
	this.BaseUnitScale = "Days"
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
			var valueForIsAutomaticMajorUnit bool
			err = json.Unmarshal(*valIsAutomaticMajorUnit, &valueForIsAutomaticMajorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMajorUnit = valueForIsAutomaticMajorUnit
		}
	}
	if valIsAutomaticMajorUnitCap, ok := objMap["IsAutomaticMajorUnit"]; ok {
		if valIsAutomaticMajorUnitCap != nil {
			var valueForIsAutomaticMajorUnit bool
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
	this.MajorUnitScale = "Days"
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
	this.MajorTickMark = "Cross"
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
			var valueForIsAutomaticMinorUnit bool
			err = json.Unmarshal(*valIsAutomaticMinorUnit, &valueForIsAutomaticMinorUnit)
			if err != nil {
				return err
			}
			this.IsAutomaticMinorUnit = valueForIsAutomaticMinorUnit
		}
	}
	if valIsAutomaticMinorUnitCap, ok := objMap["IsAutomaticMinorUnit"]; ok {
		if valIsAutomaticMinorUnitCap != nil {
			var valueForIsAutomaticMinorUnit bool
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
	this.MinorUnitScale = "Days"
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
	this.MinorTickMark = "Cross"
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
			var valueForIsAutomaticMaxValue bool
			err = json.Unmarshal(*valIsAutomaticMaxValue, &valueForIsAutomaticMaxValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMaxValue = valueForIsAutomaticMaxValue
		}
	}
	if valIsAutomaticMaxValueCap, ok := objMap["IsAutomaticMaxValue"]; ok {
		if valIsAutomaticMaxValueCap != nil {
			var valueForIsAutomaticMaxValue bool
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
			var valueForIsAutomaticMinValue bool
			err = json.Unmarshal(*valIsAutomaticMinValue, &valueForIsAutomaticMinValue)
			if err != nil {
				return err
			}
			this.IsAutomaticMinValue = valueForIsAutomaticMinValue
		}
	}
	if valIsAutomaticMinValueCap, ok := objMap["IsAutomaticMinValue"]; ok {
		if valIsAutomaticMinValueCap != nil {
			var valueForIsAutomaticMinValue bool
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
			var valueForIsLogarithmic bool
			err = json.Unmarshal(*valIsLogarithmic, &valueForIsLogarithmic)
			if err != nil {
				return err
			}
			this.IsLogarithmic = valueForIsLogarithmic
		}
	}
	if valIsLogarithmicCap, ok := objMap["IsLogarithmic"]; ok {
		if valIsLogarithmicCap != nil {
			var valueForIsLogarithmic bool
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
	this.CategoryAxisType = "Text"
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
			var valueForAxisBetweenCategories bool
			err = json.Unmarshal(*valAxisBetweenCategories, &valueForAxisBetweenCategories)
			if err != nil {
				return err
			}
			this.AxisBetweenCategories = valueForAxisBetweenCategories
		}
	}
	if valAxisBetweenCategoriesCap, ok := objMap["AxisBetweenCategories"]; ok {
		if valAxisBetweenCategoriesCap != nil {
			var valueForAxisBetweenCategories bool
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
			var valueForIsPlotOrderReversed bool
			err = json.Unmarshal(*valIsPlotOrderReversed, &valueForIsPlotOrderReversed)
			if err != nil {
				return err
			}
			this.IsPlotOrderReversed = valueForIsPlotOrderReversed
		}
	}
	if valIsPlotOrderReversedCap, ok := objMap["IsPlotOrderReversed"]; ok {
		if valIsPlotOrderReversedCap != nil {
			var valueForIsPlotOrderReversed bool
			err = json.Unmarshal(*valIsPlotOrderReversedCap, &valueForIsPlotOrderReversed)
			if err != nil {
				return err
			}
			this.IsPlotOrderReversed = valueForIsPlotOrderReversed
		}
	}
	
	if valIsNumberFormatLinkedToSource, ok := objMap["isNumberFormatLinkedToSource"]; ok {
		if valIsNumberFormatLinkedToSource != nil {
			var valueForIsNumberFormatLinkedToSource bool
			err = json.Unmarshal(*valIsNumberFormatLinkedToSource, &valueForIsNumberFormatLinkedToSource)
			if err != nil {
				return err
			}
			this.IsNumberFormatLinkedToSource = valueForIsNumberFormatLinkedToSource
		}
	}
	if valIsNumberFormatLinkedToSourceCap, ok := objMap["IsNumberFormatLinkedToSource"]; ok {
		if valIsNumberFormatLinkedToSourceCap != nil {
			var valueForIsNumberFormatLinkedToSource bool
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
	this.CrossType = "AxisCrossesAtZero"
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
			var valueForIsAutomaticTickMarksSpacing bool
			err = json.Unmarshal(*valIsAutomaticTickMarksSpacing, &valueForIsAutomaticTickMarksSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickMarksSpacing = valueForIsAutomaticTickMarksSpacing
		}
	}
	if valIsAutomaticTickMarksSpacingCap, ok := objMap["IsAutomaticTickMarksSpacing"]; ok {
		if valIsAutomaticTickMarksSpacingCap != nil {
			var valueForIsAutomaticTickMarksSpacing bool
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
			var valueForIsAutomaticTickLabelSpacing bool
			err = json.Unmarshal(*valIsAutomaticTickLabelSpacing, &valueForIsAutomaticTickLabelSpacing)
			if err != nil {
				return err
			}
			this.IsAutomaticTickLabelSpacing = valueForIsAutomaticTickLabelSpacing
		}
	}
	if valIsAutomaticTickLabelSpacingCap, ok := objMap["IsAutomaticTickLabelSpacing"]; ok {
		if valIsAutomaticTickLabelSpacingCap != nil {
			var valueForIsAutomaticTickLabelSpacing bool
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
	this.TickLabelPosition = "High"
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
			this.FillFormat = valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

    return nil
}
