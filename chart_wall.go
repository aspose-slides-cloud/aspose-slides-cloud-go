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

// Represents a chart wall
type IChartWall interface {

	// Get or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Get or sets wall thickness as a percentage of the largest dimension of the plot volume.
	getThickness() int32
	setThickness(newValue int32)

	// Get or sets mode of bar picture filling.
	getPictureType() string
	setPictureType(newValue string)
}

type ChartWall struct {

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Get or sets wall thickness as a percentage of the largest dimension of the plot volume.
	Thickness int32 `json:"Thickness,omitempty"`

	// Get or sets mode of bar picture filling.
	PictureType string `json:"PictureType,omitempty"`
}

func NewChartWall() *ChartWall {
	instance := new(ChartWall)
	instance.PictureType = ""
	return instance
}

func (this *ChartWall) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ChartWall) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ChartWall) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ChartWall) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ChartWall) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ChartWall) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ChartWall) getThickness() int32 {
	return this.Thickness
}

func (this *ChartWall) setThickness(newValue int32) {
	this.Thickness = newValue
}
func (this *ChartWall) getPictureType() string {
	return this.PictureType
}

func (this *ChartWall) setPictureType(newValue string) {
	this.PictureType = newValue
}

func (this *ChartWall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	
	if valThickness, ok := objMap["thickness"]; ok {
		if valThickness != nil {
			var valueForThickness int32
			err = json.Unmarshal(*valThickness, &valueForThickness)
			if err != nil {
				return err
			}
			this.Thickness = valueForThickness
		}
	}
	if valThicknessCap, ok := objMap["Thickness"]; ok {
		if valThicknessCap != nil {
			var valueForThickness int32
			err = json.Unmarshal(*valThicknessCap, &valueForThickness)
			if err != nil {
				return err
			}
			this.Thickness = valueForThickness
		}
	}
	this.PictureType = ""
	if valPictureType, ok := objMap["pictureType"]; ok {
		if valPictureType != nil {
			var valueForPictureType string
			err = json.Unmarshal(*valPictureType, &valueForPictureType)
			if err != nil {
				var valueForPictureTypeInt int32
				err = json.Unmarshal(*valPictureType, &valueForPictureTypeInt)
				if err != nil {
					return err
				}
				this.PictureType = string(valueForPictureTypeInt)
			} else {
				this.PictureType = valueForPictureType
			}
		}
	}
	if valPictureTypeCap, ok := objMap["PictureType"]; ok {
		if valPictureTypeCap != nil {
			var valueForPictureType string
			err = json.Unmarshal(*valPictureTypeCap, &valueForPictureType)
			if err != nil {
				var valueForPictureTypeInt int32
				err = json.Unmarshal(*valPictureTypeCap, &valueForPictureTypeInt)
				if err != nil {
					return err
				}
				this.PictureType = string(valueForPictureTypeInt)
			} else {
				this.PictureType = valueForPictureType
			}
		}
	}

	return nil
}
