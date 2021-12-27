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

// Represents preset shadow effect 
type IPresetShadowEffect interface {

	// direction
	getDirection() float64
	setDirection(newValue float64)

	// distance
	getDistance() float64
	setDistance(newValue float64)

	// preset
	getPreset() string
	setPreset(newValue string)

	// shadow color
	getShadowColor() string
	setShadowColor(newValue string)
}

type PresetShadowEffect struct {

	// direction
	Direction float64 `json:"Direction"`

	// distance
	Distance float64 `json:"Distance"`

	// preset
	Preset string `json:"Preset"`

	// shadow color
	ShadowColor string `json:"ShadowColor,omitempty"`
}

func NewPresetShadowEffect() *PresetShadowEffect {
	instance := new(PresetShadowEffect)
	instance.Preset = "TopLeftDropShadow"
	return instance
}

func (this *PresetShadowEffect) getDirection() float64 {
	return this.Direction
}

func (this *PresetShadowEffect) setDirection(newValue float64) {
	this.Direction = newValue
}
func (this *PresetShadowEffect) getDistance() float64 {
	return this.Distance
}

func (this *PresetShadowEffect) setDistance(newValue float64) {
	this.Distance = newValue
}
func (this *PresetShadowEffect) getPreset() string {
	return this.Preset
}

func (this *PresetShadowEffect) setPreset(newValue string) {
	this.Preset = newValue
}
func (this *PresetShadowEffect) getShadowColor() string {
	return this.ShadowColor
}

func (this *PresetShadowEffect) setShadowColor(newValue string) {
	this.ShadowColor = newValue
}

func (this *PresetShadowEffect) UnmarshalJSON(b []byte) error {
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
	this.Preset = "TopLeftDropShadow"
	if valPreset, ok := objMap["preset"]; ok {
		if valPreset != nil {
			var valueForPreset string
			err = json.Unmarshal(*valPreset, &valueForPreset)
			if err != nil {
				var valueForPresetInt int32
				err = json.Unmarshal(*valPreset, &valueForPresetInt)
				if err != nil {
					return err
				}
				this.Preset = string(valueForPresetInt)
			} else {
				this.Preset = valueForPreset
			}
		}
	}
	if valPresetCap, ok := objMap["Preset"]; ok {
		if valPresetCap != nil {
			var valueForPreset string
			err = json.Unmarshal(*valPresetCap, &valueForPreset)
			if err != nil {
				var valueForPresetInt int32
				err = json.Unmarshal(*valPresetCap, &valueForPresetInt)
				if err != nil {
					return err
				}
				this.Preset = string(valueForPresetInt)
			} else {
				this.Preset = valueForPreset
			}
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
