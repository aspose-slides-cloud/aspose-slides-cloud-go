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

// Represents gradient fill format
type IGradientFill interface {

	// Type of fill.
	GetType() string
	SetType(newValue string)

	// Gradient style.
	GetDirection() string
	SetDirection(newValue string)

	// Gradient shape.
	GetShape() string
	SetShape(newValue string)

	// Gradient stops.
	GetStops() []IGradientFillStop
	SetStops(newValue []IGradientFillStop)

	// Gradient angle.
	GetLinearAngle() float64
	SetLinearAngle(newValue float64)

	// True if the gradient is scaled.
	GetIsScaled() *bool
	SetIsScaled(newValue *bool)

	// Gradient flipping mode.
	GetTileFlip() string
	SetTileFlip(newValue string)
}

type GradientFill struct {

	// Type of fill.
	Type_ string `json:"Type"`

	// Gradient style.
	Direction string `json:"Direction,omitempty"`

	// Gradient shape.
	Shape string `json:"Shape,omitempty"`

	// Gradient stops.
	Stops []IGradientFillStop `json:"Stops,omitempty"`

	// Gradient angle.
	LinearAngle float64 `json:"LinearAngle,omitempty"`

	// True if the gradient is scaled.
	IsScaled *bool `json:"IsScaled"`

	// Gradient flipping mode.
	TileFlip string `json:"TileFlip,omitempty"`
}

func NewGradientFill() *GradientFill {
	instance := new(GradientFill)
	instance.Type_ = "Gradient"
	return instance
}

func (this *GradientFill) GetType() string {
	return this.Type_
}

func (this *GradientFill) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *GradientFill) GetDirection() string {
	return this.Direction
}

func (this *GradientFill) SetDirection(newValue string) {
	this.Direction = newValue
}
func (this *GradientFill) GetShape() string {
	return this.Shape
}

func (this *GradientFill) SetShape(newValue string) {
	this.Shape = newValue
}
func (this *GradientFill) GetStops() []IGradientFillStop {
	return this.Stops
}

func (this *GradientFill) SetStops(newValue []IGradientFillStop) {
	this.Stops = newValue
}
func (this *GradientFill) GetLinearAngle() float64 {
	return this.LinearAngle
}

func (this *GradientFill) SetLinearAngle(newValue float64) {
	this.LinearAngle = newValue
}
func (this *GradientFill) GetIsScaled() *bool {
	return this.IsScaled
}

func (this *GradientFill) SetIsScaled(newValue *bool) {
	this.IsScaled = newValue
}
func (this *GradientFill) GetTileFlip() string {
	return this.TileFlip
}

func (this *GradientFill) SetTileFlip(newValue string) {
	this.TileFlip = newValue
}

func (this *GradientFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Gradient"
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valDirection, ok := GetMapValue(objMap, "direction"); ok {
		if valDirection != nil {
			var valueForDirection string
			err = json.Unmarshal(*valDirection, &valueForDirection)
			if err != nil {
				var valueForDirectionInt int32
				err = json.Unmarshal(*valDirection, &valueForDirectionInt)
				if err != nil {
					return err
				}
				this.Direction = string(valueForDirectionInt)
			} else {
				this.Direction = valueForDirection
			}
		}
	}
	
	if valShape, ok := GetMapValue(objMap, "shape"); ok {
		if valShape != nil {
			var valueForShape string
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				var valueForShapeInt int32
				err = json.Unmarshal(*valShape, &valueForShapeInt)
				if err != nil {
					return err
				}
				this.Shape = string(valueForShapeInt)
			} else {
				this.Shape = valueForShape
			}
		}
	}
	
	if valStops, ok := GetMapValue(objMap, "stops"); ok {
		if valStops != nil {
			var valueForStops []json.RawMessage
			err = json.Unmarshal(*valStops, &valueForStops)
			if err != nil {
				return err
			}
			valueForIStops := make([]IGradientFillStop, len(valueForStops))
			for i, v := range valueForStops {
				vObject, err := createObjectForType("GradientFillStop", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIStops[i] = vObject.(IGradientFillStop)
				}
			}
			this.Stops = valueForIStops
		}
	}
	
	if valLinearAngle, ok := GetMapValue(objMap, "linearAngle"); ok {
		if valLinearAngle != nil {
			var valueForLinearAngle float64
			err = json.Unmarshal(*valLinearAngle, &valueForLinearAngle)
			if err != nil {
				return err
			}
			this.LinearAngle = valueForLinearAngle
		}
	}
	
	if valIsScaled, ok := GetMapValue(objMap, "isScaled"); ok {
		if valIsScaled != nil {
			var valueForIsScaled *bool
			err = json.Unmarshal(*valIsScaled, &valueForIsScaled)
			if err != nil {
				return err
			}
			this.IsScaled = valueForIsScaled
		}
	}
	
	if valTileFlip, ok := GetMapValue(objMap, "tileFlip"); ok {
		if valTileFlip != nil {
			var valueForTileFlip string
			err = json.Unmarshal(*valTileFlip, &valueForTileFlip)
			if err != nil {
				var valueForTileFlipInt int32
				err = json.Unmarshal(*valTileFlip, &valueForTileFlipInt)
				if err != nil {
					return err
				}
				this.TileFlip = string(valueForTileFlipInt)
			} else {
				this.TileFlip = valueForTileFlip
			}
		}
	}

	return nil
}
