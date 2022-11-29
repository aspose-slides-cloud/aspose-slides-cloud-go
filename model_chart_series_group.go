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

// Chart series group. Defines common properties for a group of series.
type IChartSeriesGroup interface {

	// Returns a type of this series group.
	GetType() string
	SetType(newValue string)

	// Specifies the space between bar or column clusters, as a percentage of the bar or column width.
	GetGapWidth() int32
	SetGapWidth(newValue int32)

	// Returns or sets the distance, as a percentage of the marker width, between the data series in a 3D chart.
	GetGapDepth() int32
	SetGapDepth(newValue int32)

	// Gets or sets the angle of the first pie or doughnut chart slice,  in degrees (clockwise from up, from 0 to 360 degrees).
	GetFirstSliceAngle() int32
	SetFirstSliceAngle(newValue int32)

	// Specifies that each data marker in the series has a different color.
	GetIsColorVaried() bool
	SetIsColorVaried(newValue bool)

	// True if chart has series lines. Applied to stacked bar and OfPie charts.
	GetHasSeriesLines() bool
	SetHasSeriesLines(newValue bool)

	// Specifies how much bars and columns shall overlap on 2-D charts (from -100 to 100).
	GetOverlap() int32
	SetOverlap(newValue int32)

	// Specifies the size of the second pie or bar of a pie-of-pie chart or  a bar-of-pie chart, as a percentage of the size of the first pie (can  be between 5 and 200 percents).
	GetSecondPieSize() int32
	SetSecondPieSize(newValue int32)

	// Specifies a value that shall be used to determine which data points  are in the second pie or bar on a pie-of-pie or bar-of-pie chart.  Is used together with PieSplitBy property.
	GetPieSplitPosition() float64
	SetPieSplitPosition(newValue float64)

	// Specifies how to determine which data points are in the second pie or bar  on a pie-of-pie or bar-of-pie chart.
	GetPieSplitBy() string
	SetPieSplitBy(newValue string)

	// Specifies the size of the hole in a doughnut chart (can be between 10 and 90 percents  of the size of the plot area.).
	GetDoughnutHoleSize() int32
	SetDoughnutHoleSize(newValue int32)

	// Specifies the scale factor for the bubble chart (can be  between 0 and 300 percents of the default size). Read/write Int32.
	GetBubbleSizeScale() int32
	SetBubbleSizeScale(newValue int32)

	// Specifies HiLowLines format.  HiLowLines applied with HiLowClose, OpenHiLowClose, VolumeHiLowClose and VolumeOpenHiLowClose chart types.
	GetHiLowLinesFormat() IChartLinesFormat
	SetHiLowLinesFormat(newValue IChartLinesFormat)

	// Specifies how the bubble size values are represented on the bubble chart. Read/write BubbleSizeRepresentationType.
	GetBubbleSizeRepresentation() string
	SetBubbleSizeRepresentation(newValue string)
}

type ChartSeriesGroup struct {

	// Returns a type of this series group.
	Type_ string `json:"Type,omitempty"`

	// Specifies the space between bar or column clusters, as a percentage of the bar or column width.
	GapWidth int32 `json:"GapWidth,omitempty"`

	// Returns or sets the distance, as a percentage of the marker width, between the data series in a 3D chart.
	GapDepth int32 `json:"GapDepth,omitempty"`

	// Gets or sets the angle of the first pie or doughnut chart slice,  in degrees (clockwise from up, from 0 to 360 degrees).
	FirstSliceAngle int32 `json:"FirstSliceAngle,omitempty"`

	// Specifies that each data marker in the series has a different color.
	IsColorVaried bool `json:"IsColorVaried"`

	// True if chart has series lines. Applied to stacked bar and OfPie charts.
	HasSeriesLines bool `json:"HasSeriesLines"`

	// Specifies how much bars and columns shall overlap on 2-D charts (from -100 to 100).
	Overlap int32 `json:"Overlap,omitempty"`

	// Specifies the size of the second pie or bar of a pie-of-pie chart or  a bar-of-pie chart, as a percentage of the size of the first pie (can  be between 5 and 200 percents).
	SecondPieSize int32 `json:"SecondPieSize,omitempty"`

	// Specifies a value that shall be used to determine which data points  are in the second pie or bar on a pie-of-pie or bar-of-pie chart.  Is used together with PieSplitBy property.
	PieSplitPosition float64 `json:"PieSplitPosition,omitempty"`

	// Specifies how to determine which data points are in the second pie or bar  on a pie-of-pie or bar-of-pie chart.
	PieSplitBy string `json:"PieSplitBy,omitempty"`

	// Specifies the size of the hole in a doughnut chart (can be between 10 and 90 percents  of the size of the plot area.).
	DoughnutHoleSize int32 `json:"DoughnutHoleSize,omitempty"`

	// Specifies the scale factor for the bubble chart (can be  between 0 and 300 percents of the default size). Read/write Int32.
	BubbleSizeScale int32 `json:"BubbleSizeScale,omitempty"`

	// Specifies HiLowLines format.  HiLowLines applied with HiLowClose, OpenHiLowClose, VolumeHiLowClose and VolumeOpenHiLowClose chart types.
	HiLowLinesFormat IChartLinesFormat `json:"HiLowLinesFormat,omitempty"`

	// Specifies how the bubble size values are represented on the bubble chart. Read/write BubbleSizeRepresentationType.
	BubbleSizeRepresentation string `json:"BubbleSizeRepresentation,omitempty"`
}

func NewChartSeriesGroup() *ChartSeriesGroup {
	instance := new(ChartSeriesGroup)
	return instance
}

func (this *ChartSeriesGroup) GetType() string {
	return this.Type_
}

func (this *ChartSeriesGroup) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ChartSeriesGroup) GetGapWidth() int32 {
	return this.GapWidth
}

func (this *ChartSeriesGroup) SetGapWidth(newValue int32) {
	this.GapWidth = newValue
}
func (this *ChartSeriesGroup) GetGapDepth() int32 {
	return this.GapDepth
}

func (this *ChartSeriesGroup) SetGapDepth(newValue int32) {
	this.GapDepth = newValue
}
func (this *ChartSeriesGroup) GetFirstSliceAngle() int32 {
	return this.FirstSliceAngle
}

func (this *ChartSeriesGroup) SetFirstSliceAngle(newValue int32) {
	this.FirstSliceAngle = newValue
}
func (this *ChartSeriesGroup) GetIsColorVaried() bool {
	return this.IsColorVaried
}

func (this *ChartSeriesGroup) SetIsColorVaried(newValue bool) {
	this.IsColorVaried = newValue
}
func (this *ChartSeriesGroup) GetHasSeriesLines() bool {
	return this.HasSeriesLines
}

func (this *ChartSeriesGroup) SetHasSeriesLines(newValue bool) {
	this.HasSeriesLines = newValue
}
func (this *ChartSeriesGroup) GetOverlap() int32 {
	return this.Overlap
}

func (this *ChartSeriesGroup) SetOverlap(newValue int32) {
	this.Overlap = newValue
}
func (this *ChartSeriesGroup) GetSecondPieSize() int32 {
	return this.SecondPieSize
}

func (this *ChartSeriesGroup) SetSecondPieSize(newValue int32) {
	this.SecondPieSize = newValue
}
func (this *ChartSeriesGroup) GetPieSplitPosition() float64 {
	return this.PieSplitPosition
}

func (this *ChartSeriesGroup) SetPieSplitPosition(newValue float64) {
	this.PieSplitPosition = newValue
}
func (this *ChartSeriesGroup) GetPieSplitBy() string {
	return this.PieSplitBy
}

func (this *ChartSeriesGroup) SetPieSplitBy(newValue string) {
	this.PieSplitBy = newValue
}
func (this *ChartSeriesGroup) GetDoughnutHoleSize() int32 {
	return this.DoughnutHoleSize
}

func (this *ChartSeriesGroup) SetDoughnutHoleSize(newValue int32) {
	this.DoughnutHoleSize = newValue
}
func (this *ChartSeriesGroup) GetBubbleSizeScale() int32 {
	return this.BubbleSizeScale
}

func (this *ChartSeriesGroup) SetBubbleSizeScale(newValue int32) {
	this.BubbleSizeScale = newValue
}
func (this *ChartSeriesGroup) GetHiLowLinesFormat() IChartLinesFormat {
	return this.HiLowLinesFormat
}

func (this *ChartSeriesGroup) SetHiLowLinesFormat(newValue IChartLinesFormat) {
	this.HiLowLinesFormat = newValue
}
func (this *ChartSeriesGroup) GetBubbleSizeRepresentation() string {
	return this.BubbleSizeRepresentation
}

func (this *ChartSeriesGroup) SetBubbleSizeRepresentation(newValue string) {
	this.BubbleSizeRepresentation = newValue
}

func (this *ChartSeriesGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
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
	
	if valGapWidth, ok := objMap["gapWidth"]; ok {
		if valGapWidth != nil {
			var valueForGapWidth int32
			err = json.Unmarshal(*valGapWidth, &valueForGapWidth)
			if err != nil {
				return err
			}
			this.GapWidth = valueForGapWidth
		}
	}
	if valGapWidthCap, ok := objMap["GapWidth"]; ok {
		if valGapWidthCap != nil {
			var valueForGapWidth int32
			err = json.Unmarshal(*valGapWidthCap, &valueForGapWidth)
			if err != nil {
				return err
			}
			this.GapWidth = valueForGapWidth
		}
	}
	
	if valGapDepth, ok := objMap["gapDepth"]; ok {
		if valGapDepth != nil {
			var valueForGapDepth int32
			err = json.Unmarshal(*valGapDepth, &valueForGapDepth)
			if err != nil {
				return err
			}
			this.GapDepth = valueForGapDepth
		}
	}
	if valGapDepthCap, ok := objMap["GapDepth"]; ok {
		if valGapDepthCap != nil {
			var valueForGapDepth int32
			err = json.Unmarshal(*valGapDepthCap, &valueForGapDepth)
			if err != nil {
				return err
			}
			this.GapDepth = valueForGapDepth
		}
	}
	
	if valFirstSliceAngle, ok := objMap["firstSliceAngle"]; ok {
		if valFirstSliceAngle != nil {
			var valueForFirstSliceAngle int32
			err = json.Unmarshal(*valFirstSliceAngle, &valueForFirstSliceAngle)
			if err != nil {
				return err
			}
			this.FirstSliceAngle = valueForFirstSliceAngle
		}
	}
	if valFirstSliceAngleCap, ok := objMap["FirstSliceAngle"]; ok {
		if valFirstSliceAngleCap != nil {
			var valueForFirstSliceAngle int32
			err = json.Unmarshal(*valFirstSliceAngleCap, &valueForFirstSliceAngle)
			if err != nil {
				return err
			}
			this.FirstSliceAngle = valueForFirstSliceAngle
		}
	}
	
	if valIsColorVaried, ok := objMap["isColorVaried"]; ok {
		if valIsColorVaried != nil {
			var valueForIsColorVaried bool
			err = json.Unmarshal(*valIsColorVaried, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}
	if valIsColorVariedCap, ok := objMap["IsColorVaried"]; ok {
		if valIsColorVariedCap != nil {
			var valueForIsColorVaried bool
			err = json.Unmarshal(*valIsColorVariedCap, &valueForIsColorVaried)
			if err != nil {
				return err
			}
			this.IsColorVaried = valueForIsColorVaried
		}
	}
	
	if valHasSeriesLines, ok := objMap["hasSeriesLines"]; ok {
		if valHasSeriesLines != nil {
			var valueForHasSeriesLines bool
			err = json.Unmarshal(*valHasSeriesLines, &valueForHasSeriesLines)
			if err != nil {
				return err
			}
			this.HasSeriesLines = valueForHasSeriesLines
		}
	}
	if valHasSeriesLinesCap, ok := objMap["HasSeriesLines"]; ok {
		if valHasSeriesLinesCap != nil {
			var valueForHasSeriesLines bool
			err = json.Unmarshal(*valHasSeriesLinesCap, &valueForHasSeriesLines)
			if err != nil {
				return err
			}
			this.HasSeriesLines = valueForHasSeriesLines
		}
	}
	
	if valOverlap, ok := objMap["overlap"]; ok {
		if valOverlap != nil {
			var valueForOverlap int32
			err = json.Unmarshal(*valOverlap, &valueForOverlap)
			if err != nil {
				return err
			}
			this.Overlap = valueForOverlap
		}
	}
	if valOverlapCap, ok := objMap["Overlap"]; ok {
		if valOverlapCap != nil {
			var valueForOverlap int32
			err = json.Unmarshal(*valOverlapCap, &valueForOverlap)
			if err != nil {
				return err
			}
			this.Overlap = valueForOverlap
		}
	}
	
	if valSecondPieSize, ok := objMap["secondPieSize"]; ok {
		if valSecondPieSize != nil {
			var valueForSecondPieSize int32
			err = json.Unmarshal(*valSecondPieSize, &valueForSecondPieSize)
			if err != nil {
				return err
			}
			this.SecondPieSize = valueForSecondPieSize
		}
	}
	if valSecondPieSizeCap, ok := objMap["SecondPieSize"]; ok {
		if valSecondPieSizeCap != nil {
			var valueForSecondPieSize int32
			err = json.Unmarshal(*valSecondPieSizeCap, &valueForSecondPieSize)
			if err != nil {
				return err
			}
			this.SecondPieSize = valueForSecondPieSize
		}
	}
	
	if valPieSplitPosition, ok := objMap["pieSplitPosition"]; ok {
		if valPieSplitPosition != nil {
			var valueForPieSplitPosition float64
			err = json.Unmarshal(*valPieSplitPosition, &valueForPieSplitPosition)
			if err != nil {
				return err
			}
			this.PieSplitPosition = valueForPieSplitPosition
		}
	}
	if valPieSplitPositionCap, ok := objMap["PieSplitPosition"]; ok {
		if valPieSplitPositionCap != nil {
			var valueForPieSplitPosition float64
			err = json.Unmarshal(*valPieSplitPositionCap, &valueForPieSplitPosition)
			if err != nil {
				return err
			}
			this.PieSplitPosition = valueForPieSplitPosition
		}
	}
	
	if valPieSplitBy, ok := objMap["pieSplitBy"]; ok {
		if valPieSplitBy != nil {
			var valueForPieSplitBy string
			err = json.Unmarshal(*valPieSplitBy, &valueForPieSplitBy)
			if err != nil {
				var valueForPieSplitByInt int32
				err = json.Unmarshal(*valPieSplitBy, &valueForPieSplitByInt)
				if err != nil {
					return err
				}
				this.PieSplitBy = string(valueForPieSplitByInt)
			} else {
				this.PieSplitBy = valueForPieSplitBy
			}
		}
	}
	if valPieSplitByCap, ok := objMap["PieSplitBy"]; ok {
		if valPieSplitByCap != nil {
			var valueForPieSplitBy string
			err = json.Unmarshal(*valPieSplitByCap, &valueForPieSplitBy)
			if err != nil {
				var valueForPieSplitByInt int32
				err = json.Unmarshal(*valPieSplitByCap, &valueForPieSplitByInt)
				if err != nil {
					return err
				}
				this.PieSplitBy = string(valueForPieSplitByInt)
			} else {
				this.PieSplitBy = valueForPieSplitBy
			}
		}
	}
	
	if valDoughnutHoleSize, ok := objMap["doughnutHoleSize"]; ok {
		if valDoughnutHoleSize != nil {
			var valueForDoughnutHoleSize int32
			err = json.Unmarshal(*valDoughnutHoleSize, &valueForDoughnutHoleSize)
			if err != nil {
				return err
			}
			this.DoughnutHoleSize = valueForDoughnutHoleSize
		}
	}
	if valDoughnutHoleSizeCap, ok := objMap["DoughnutHoleSize"]; ok {
		if valDoughnutHoleSizeCap != nil {
			var valueForDoughnutHoleSize int32
			err = json.Unmarshal(*valDoughnutHoleSizeCap, &valueForDoughnutHoleSize)
			if err != nil {
				return err
			}
			this.DoughnutHoleSize = valueForDoughnutHoleSize
		}
	}
	
	if valBubbleSizeScale, ok := objMap["bubbleSizeScale"]; ok {
		if valBubbleSizeScale != nil {
			var valueForBubbleSizeScale int32
			err = json.Unmarshal(*valBubbleSizeScale, &valueForBubbleSizeScale)
			if err != nil {
				return err
			}
			this.BubbleSizeScale = valueForBubbleSizeScale
		}
	}
	if valBubbleSizeScaleCap, ok := objMap["BubbleSizeScale"]; ok {
		if valBubbleSizeScaleCap != nil {
			var valueForBubbleSizeScale int32
			err = json.Unmarshal(*valBubbleSizeScaleCap, &valueForBubbleSizeScale)
			if err != nil {
				return err
			}
			this.BubbleSizeScale = valueForBubbleSizeScale
		}
	}
	
	if valHiLowLinesFormat, ok := objMap["hiLowLinesFormat"]; ok {
		if valHiLowLinesFormat != nil {
			var valueForHiLowLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valHiLowLinesFormat, &valueForHiLowLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valHiLowLinesFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHiLowLinesFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.HiLowLinesFormat = vInterfaceObject
			}
		}
	}
	if valHiLowLinesFormatCap, ok := objMap["HiLowLinesFormat"]; ok {
		if valHiLowLinesFormatCap != nil {
			var valueForHiLowLinesFormat ChartLinesFormat
			err = json.Unmarshal(*valHiLowLinesFormatCap, &valueForHiLowLinesFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ChartLinesFormat", *valHiLowLinesFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHiLowLinesFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IChartLinesFormat)
			if ok {
				this.HiLowLinesFormat = vInterfaceObject
			}
		}
	}
	
	if valBubbleSizeRepresentation, ok := objMap["bubbleSizeRepresentation"]; ok {
		if valBubbleSizeRepresentation != nil {
			var valueForBubbleSizeRepresentation string
			err = json.Unmarshal(*valBubbleSizeRepresentation, &valueForBubbleSizeRepresentation)
			if err != nil {
				var valueForBubbleSizeRepresentationInt int32
				err = json.Unmarshal(*valBubbleSizeRepresentation, &valueForBubbleSizeRepresentationInt)
				if err != nil {
					return err
				}
				this.BubbleSizeRepresentation = string(valueForBubbleSizeRepresentationInt)
			} else {
				this.BubbleSizeRepresentation = valueForBubbleSizeRepresentation
			}
		}
	}
	if valBubbleSizeRepresentationCap, ok := objMap["BubbleSizeRepresentation"]; ok {
		if valBubbleSizeRepresentationCap != nil {
			var valueForBubbleSizeRepresentation string
			err = json.Unmarshal(*valBubbleSizeRepresentationCap, &valueForBubbleSizeRepresentation)
			if err != nil {
				var valueForBubbleSizeRepresentationInt int32
				err = json.Unmarshal(*valBubbleSizeRepresentationCap, &valueForBubbleSizeRepresentationInt)
				if err != nil {
					return err
				}
				this.BubbleSizeRepresentation = string(valueForBubbleSizeRepresentationInt)
			} else {
				this.BubbleSizeRepresentation = valueForBubbleSizeRepresentation
			}
		}
	}

	return nil
}
