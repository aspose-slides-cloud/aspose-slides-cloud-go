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
	getHorizontalAxis() IAxis
	setHorizontalAxis(newValue IAxis)

	// Gets or sets the vertical axis.
	getVerticalAxis() IAxis
	setVerticalAxis(newValue IAxis)

	// Gets or sets the secondary horizontal axis.
	getSecondaryHorizontalAxis() IAxis
	setSecondaryHorizontalAxis(newValue IAxis)

	// Gets or sets the secondary vertical axis.
	getSecondaryVerticalAxis() IAxis
	setSecondaryVerticalAxis(newValue IAxis)
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

func (this Axes) getHorizontalAxis() IAxis {
	return this.HorizontalAxis
}

func (this Axes) setHorizontalAxis(newValue IAxis) {
	this.HorizontalAxis = newValue
}
func (this Axes) getVerticalAxis() IAxis {
	return this.VerticalAxis
}

func (this Axes) setVerticalAxis(newValue IAxis) {
	this.VerticalAxis = newValue
}
func (this Axes) getSecondaryHorizontalAxis() IAxis {
	return this.SecondaryHorizontalAxis
}

func (this Axes) setSecondaryHorizontalAxis(newValue IAxis) {
	this.SecondaryHorizontalAxis = newValue
}
func (this Axes) getSecondaryVerticalAxis() IAxis {
	return this.SecondaryVerticalAxis
}

func (this Axes) setSecondaryVerticalAxis(newValue IAxis) {
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
			this.HorizontalAxis = valueForHorizontalAxis
		}
	}
	if valHorizontalAxisCap, ok := objMap["HorizontalAxis"]; ok {
		if valHorizontalAxisCap != nil {
			var valueForHorizontalAxis Axis
			err = json.Unmarshal(*valHorizontalAxisCap, &valueForHorizontalAxis)
			if err != nil {
				return err
			}
			this.HorizontalAxis = valueForHorizontalAxis
		}
	}
	
	if valVerticalAxis, ok := objMap["verticalAxis"]; ok {
		if valVerticalAxis != nil {
			var valueForVerticalAxis Axis
			err = json.Unmarshal(*valVerticalAxis, &valueForVerticalAxis)
			if err != nil {
				return err
			}
			this.VerticalAxis = valueForVerticalAxis
		}
	}
	if valVerticalAxisCap, ok := objMap["VerticalAxis"]; ok {
		if valVerticalAxisCap != nil {
			var valueForVerticalAxis Axis
			err = json.Unmarshal(*valVerticalAxisCap, &valueForVerticalAxis)
			if err != nil {
				return err
			}
			this.VerticalAxis = valueForVerticalAxis
		}
	}
	
	if valSecondaryHorizontalAxis, ok := objMap["secondaryHorizontalAxis"]; ok {
		if valSecondaryHorizontalAxis != nil {
			var valueForSecondaryHorizontalAxis Axis
			err = json.Unmarshal(*valSecondaryHorizontalAxis, &valueForSecondaryHorizontalAxis)
			if err != nil {
				return err
			}
			this.SecondaryHorizontalAxis = valueForSecondaryHorizontalAxis
		}
	}
	if valSecondaryHorizontalAxisCap, ok := objMap["SecondaryHorizontalAxis"]; ok {
		if valSecondaryHorizontalAxisCap != nil {
			var valueForSecondaryHorizontalAxis Axis
			err = json.Unmarshal(*valSecondaryHorizontalAxisCap, &valueForSecondaryHorizontalAxis)
			if err != nil {
				return err
			}
			this.SecondaryHorizontalAxis = valueForSecondaryHorizontalAxis
		}
	}
	
	if valSecondaryVerticalAxis, ok := objMap["secondaryVerticalAxis"]; ok {
		if valSecondaryVerticalAxis != nil {
			var valueForSecondaryVerticalAxis Axis
			err = json.Unmarshal(*valSecondaryVerticalAxis, &valueForSecondaryVerticalAxis)
			if err != nil {
				return err
			}
			this.SecondaryVerticalAxis = valueForSecondaryVerticalAxis
		}
	}
	if valSecondaryVerticalAxisCap, ok := objMap["SecondaryVerticalAxis"]; ok {
		if valSecondaryVerticalAxisCap != nil {
			var valueForSecondaryVerticalAxis Axis
			err = json.Unmarshal(*valSecondaryVerticalAxisCap, &valueForSecondaryVerticalAxis)
			if err != nil {
				return err
			}
			this.SecondaryVerticalAxis = valueForSecondaryVerticalAxis
		}
	}

    return nil
}
