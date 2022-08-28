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

// Represents chart axes
type IAxes interface {

	// Gets or sets the horizontal axis.
	GetHorizontalAxis() IAxis
	SetHorizontalAxis(newValue IAxis)

	// Gets or sets the vertical axis.
	GetVerticalAxis() IAxis
	SetVerticalAxis(newValue IAxis)

	// Gets or sets the secondary horizontal axis.
	GetSecondaryHorizontalAxis() IAxis
	SetSecondaryHorizontalAxis(newValue IAxis)

	// Gets or sets the secondary vertical axis.
	GetSecondaryVerticalAxis() IAxis
	SetSecondaryVerticalAxis(newValue IAxis)
}

type Axes struct {

	// Gets or sets the horizontal axis.
	HorizontalAxis IAxis `json:"HorizontalAxis,omitempty"`

	// Gets or sets the vertical axis.
	VerticalAxis IAxis `json:"VerticalAxis,omitempty"`

	// Gets or sets the secondary horizontal axis.
	SecondaryHorizontalAxis IAxis `json:"SecondaryHorizontalAxis,omitempty"`

	// Gets or sets the secondary vertical axis.
	SecondaryVerticalAxis IAxis `json:"SecondaryVerticalAxis,omitempty"`
}

func NewAxes() *Axes {
	instance := new(Axes)
	return instance
}

func (this *Axes) GetHorizontalAxis() IAxis {
	return this.HorizontalAxis
}

func (this *Axes) SetHorizontalAxis(newValue IAxis) {
	this.HorizontalAxis = newValue
}
func (this *Axes) GetVerticalAxis() IAxis {
	return this.VerticalAxis
}

func (this *Axes) SetVerticalAxis(newValue IAxis) {
	this.VerticalAxis = newValue
}
func (this *Axes) GetSecondaryHorizontalAxis() IAxis {
	return this.SecondaryHorizontalAxis
}

func (this *Axes) SetSecondaryHorizontalAxis(newValue IAxis) {
	this.SecondaryHorizontalAxis = newValue
}
func (this *Axes) GetSecondaryVerticalAxis() IAxis {
	return this.SecondaryVerticalAxis
}

func (this *Axes) SetSecondaryVerticalAxis(newValue IAxis) {
	this.SecondaryVerticalAxis = newValue
}

func (this *Axes) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valHorizontalAxis, ok := objMap["horizontalAxis"]; ok {
		if valHorizontalAxis != nil {
			var valueForHorizontalAxis Axis
			err = json.Unmarshal(*valHorizontalAxis, &valueForHorizontalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valHorizontalAxis)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHorizontalAxis, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.HorizontalAxis = vInterfaceObject
			}
		}
	}
	if valHorizontalAxisCap, ok := objMap["HorizontalAxis"]; ok {
		if valHorizontalAxisCap != nil {
			var valueForHorizontalAxis Axis
			err = json.Unmarshal(*valHorizontalAxisCap, &valueForHorizontalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valHorizontalAxisCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHorizontalAxisCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.HorizontalAxis = vInterfaceObject
			}
		}
	}
	
	if valVerticalAxis, ok := objMap["verticalAxis"]; ok {
		if valVerticalAxis != nil {
			var valueForVerticalAxis Axis
			err = json.Unmarshal(*valVerticalAxis, &valueForVerticalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valVerticalAxis)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valVerticalAxis, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.VerticalAxis = vInterfaceObject
			}
		}
	}
	if valVerticalAxisCap, ok := objMap["VerticalAxis"]; ok {
		if valVerticalAxisCap != nil {
			var valueForVerticalAxis Axis
			err = json.Unmarshal(*valVerticalAxisCap, &valueForVerticalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valVerticalAxisCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valVerticalAxisCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.VerticalAxis = vInterfaceObject
			}
		}
	}
	
	if valSecondaryHorizontalAxis, ok := objMap["secondaryHorizontalAxis"]; ok {
		if valSecondaryHorizontalAxis != nil {
			var valueForSecondaryHorizontalAxis Axis
			err = json.Unmarshal(*valSecondaryHorizontalAxis, &valueForSecondaryHorizontalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valSecondaryHorizontalAxis)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSecondaryHorizontalAxis, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.SecondaryHorizontalAxis = vInterfaceObject
			}
		}
	}
	if valSecondaryHorizontalAxisCap, ok := objMap["SecondaryHorizontalAxis"]; ok {
		if valSecondaryHorizontalAxisCap != nil {
			var valueForSecondaryHorizontalAxis Axis
			err = json.Unmarshal(*valSecondaryHorizontalAxisCap, &valueForSecondaryHorizontalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valSecondaryHorizontalAxisCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSecondaryHorizontalAxisCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.SecondaryHorizontalAxis = vInterfaceObject
			}
		}
	}
	
	if valSecondaryVerticalAxis, ok := objMap["secondaryVerticalAxis"]; ok {
		if valSecondaryVerticalAxis != nil {
			var valueForSecondaryVerticalAxis Axis
			err = json.Unmarshal(*valSecondaryVerticalAxis, &valueForSecondaryVerticalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valSecondaryVerticalAxis)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSecondaryVerticalAxis, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.SecondaryVerticalAxis = vInterfaceObject
			}
		}
	}
	if valSecondaryVerticalAxisCap, ok := objMap["SecondaryVerticalAxis"]; ok {
		if valSecondaryVerticalAxisCap != nil {
			var valueForSecondaryVerticalAxis Axis
			err = json.Unmarshal(*valSecondaryVerticalAxisCap, &valueForSecondaryVerticalAxis)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Axis", *valSecondaryVerticalAxisCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSecondaryVerticalAxisCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAxis)
			if ok {
				this.SecondaryVerticalAxis = vInterfaceObject
			}
		}
	}

	return nil
}
