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

// Represents outer shadow effect 
type IOuterShadowEffect interface {

	// direction
	GetDirection() float64
	SetDirection(newValue float64)

	// distance
	GetDistance() float64
	SetDistance(newValue float64)

	// blur radius
	GetBlurRadius() float64
	SetBlurRadius(newValue float64)

	// shadow color
	GetShadowColor() string
	SetShadowColor(newValue string)
}

type OuterShadowEffect struct {

	// direction
	Direction float64 `json:"Direction"`

	// distance
	Distance float64 `json:"Distance"`

	// blur radius
	BlurRadius float64 `json:"BlurRadius"`

	// shadow color
	ShadowColor string `json:"ShadowColor,omitempty"`
}

func NewOuterShadowEffect() *OuterShadowEffect {
	instance := new(OuterShadowEffect)
	return instance
}

func (this *OuterShadowEffect) GetDirection() float64 {
	return this.Direction
}

func (this *OuterShadowEffect) SetDirection(newValue float64) {
	this.Direction = newValue
}
func (this *OuterShadowEffect) GetDistance() float64 {
	return this.Distance
}

func (this *OuterShadowEffect) SetDistance(newValue float64) {
	this.Distance = newValue
}
func (this *OuterShadowEffect) GetBlurRadius() float64 {
	return this.BlurRadius
}

func (this *OuterShadowEffect) SetBlurRadius(newValue float64) {
	this.BlurRadius = newValue
}
func (this *OuterShadowEffect) GetShadowColor() string {
	return this.ShadowColor
}

func (this *OuterShadowEffect) SetShadowColor(newValue string) {
	this.ShadowColor = newValue
}

func (this *OuterShadowEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDirection, ok := objMap["direction"]; ok {
		if valDirection != nil {
			var valueForDirection float64
			err = json.Unmarshal(*valDirection, &valueForDirection)
			if err != nil {
				return err
			}
			this.Direction = valueForDirection
		}
	}
	if valDirectionCap, ok := objMap["Direction"]; ok {
		if valDirectionCap != nil {
			var valueForDirection float64
			err = json.Unmarshal(*valDirectionCap, &valueForDirection)
			if err != nil {
				return err
			}
			this.Direction = valueForDirection
		}
	}
	
	if valDistance, ok := objMap["distance"]; ok {
		if valDistance != nil {
			var valueForDistance float64
			err = json.Unmarshal(*valDistance, &valueForDistance)
			if err != nil {
				return err
			}
			this.Distance = valueForDistance
		}
	}
	if valDistanceCap, ok := objMap["Distance"]; ok {
		if valDistanceCap != nil {
			var valueForDistance float64
			err = json.Unmarshal(*valDistanceCap, &valueForDistance)
			if err != nil {
				return err
			}
			this.Distance = valueForDistance
		}
	}
	
	if valBlurRadius, ok := objMap["blurRadius"]; ok {
		if valBlurRadius != nil {
			var valueForBlurRadius float64
			err = json.Unmarshal(*valBlurRadius, &valueForBlurRadius)
			if err != nil {
				return err
			}
			this.BlurRadius = valueForBlurRadius
		}
	}
	if valBlurRadiusCap, ok := objMap["BlurRadius"]; ok {
		if valBlurRadiusCap != nil {
			var valueForBlurRadius float64
			err = json.Unmarshal(*valBlurRadiusCap, &valueForBlurRadius)
			if err != nil {
				return err
			}
			this.BlurRadius = valueForBlurRadius
		}
	}
	
	if valShadowColor, ok := objMap["shadowColor"]; ok {
		if valShadowColor != nil {
			var valueForShadowColor string
			err = json.Unmarshal(*valShadowColor, &valueForShadowColor)
			if err != nil {
				return err
			}
			this.ShadowColor = valueForShadowColor
		}
	}
	if valShadowColorCap, ok := objMap["ShadowColor"]; ok {
		if valShadowColorCap != nil {
			var valueForShadowColor string
			err = json.Unmarshal(*valShadowColorCap, &valueForShadowColor)
			if err != nil {
				return err
			}
			this.ShadowColor = valueForShadowColor
		}
	}

	return nil
}
