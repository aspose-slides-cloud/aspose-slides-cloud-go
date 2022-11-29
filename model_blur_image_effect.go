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

// Represents a Blur effect that is applied to the entire shape, including its fill. All color channels, including alpha, are affected.
type IBlurImageEffect interface {

	// Image transform effect type
	GetType() string
	SetType(newValue string)

	// Returns or sets blur radius.
	GetRadius() float64
	SetRadius(newValue float64)

	// Determines whether the bounds of the object should be grown as a result of the blurring. True indicates the bounds are grown while false indicates that they are not.
	GetGrow() bool
	SetGrow(newValue bool)
}

type BlurImageEffect struct {

	// Image transform effect type
	Type_ string `json:"Type"`

	// Returns or sets blur radius.
	Radius float64 `json:"Radius"`

	// Determines whether the bounds of the object should be grown as a result of the blurring. True indicates the bounds are grown while false indicates that they are not.
	Grow bool `json:"Grow"`
}

func NewBlurImageEffect() *BlurImageEffect {
	instance := new(BlurImageEffect)
	instance.Type_ = "Blur"
	return instance
}

func (this *BlurImageEffect) GetType() string {
	return this.Type_
}

func (this *BlurImageEffect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *BlurImageEffect) GetRadius() float64 {
	return this.Radius
}

func (this *BlurImageEffect) SetRadius(newValue float64) {
	this.Radius = newValue
}
func (this *BlurImageEffect) GetGrow() bool {
	return this.Grow
}

func (this *BlurImageEffect) SetGrow(newValue bool) {
	this.Grow = newValue
}

func (this *BlurImageEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Blur"
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
	
	if valRadius, ok := objMap["radius"]; ok {
		if valRadius != nil {
			var valueForRadius float64
			err = json.Unmarshal(*valRadius, &valueForRadius)
			if err != nil {
				return err
			}
			this.Radius = valueForRadius
		}
	}
	if valRadiusCap, ok := objMap["Radius"]; ok {
		if valRadiusCap != nil {
			var valueForRadius float64
			err = json.Unmarshal(*valRadiusCap, &valueForRadius)
			if err != nil {
				return err
			}
			this.Radius = valueForRadius
		}
	}
	
	if valGrow, ok := objMap["grow"]; ok {
		if valGrow != nil {
			var valueForGrow bool
			err = json.Unmarshal(*valGrow, &valueForGrow)
			if err != nil {
				return err
			}
			this.Grow = valueForGrow
		}
	}
	if valGrowCap, ok := objMap["Grow"]; ok {
		if valGrowCap != nil {
			var valueForGrow bool
			err = json.Unmarshal(*valGrowCap, &valueForGrow)
			if err != nil {
				return err
			}
			this.Grow = valueForGrow
		}
	}

	return nil
}
