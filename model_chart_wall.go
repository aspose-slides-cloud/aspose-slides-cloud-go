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
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Get or sets wall thickness as a percentage of the largest dimension of the plot volume.
	GetThickness() int32
	SetThickness(newValue int32)

	// Get or sets mode of bar picture filling.
	GetPictureType() string
	SetPictureType(newValue string)
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
	return instance
}

func (this *ChartWall) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ChartWall) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ChartWall) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ChartWall) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ChartWall) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ChartWall) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ChartWall) GetThickness() int32 {
	return this.Thickness
}

func (this *ChartWall) SetThickness(newValue int32) {
	this.Thickness = newValue
}
func (this *ChartWall) GetPictureType() string {
	return this.PictureType
}

func (this *ChartWall) SetPictureType(newValue string) {
	this.PictureType = newValue
}

func (this *ChartWall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	
	if valThickness, ok := GetMapValue(objMap, "thickness"); ok {
		if valThickness != nil {
			var valueForThickness int32
			err = json.Unmarshal(*valThickness, &valueForThickness)
			if err != nil {
				return err
			}
			this.Thickness = valueForThickness
		}
	}
	
	if valPictureType, ok := GetMapValue(objMap, "pictureType"); ok {
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

	return nil
}
