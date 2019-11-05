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
	getType() string
	setType(newValue string)

	// Gradient style.
	getDirection() string
	setDirection(newValue string)

	// Gradient shape.
	getShape() string
	setShape(newValue string)

	// Gradient stops.
	getStops() []GradientFillStop
	setStops(newValue []GradientFillStop)

	// Gradient angle.
	getLinearAngle() float64
	setLinearAngle(newValue float64)

	// True if the gradient is scaled.
	getIsScaled() bool
	setIsScaled(newValue bool)

	// Gradient flipping mode.
	getTileFlip() string
	setTileFlip(newValue string)
}

type GradientFill struct {

	// Type of fill.
	Type_ string `json:"Type"`

	// Gradient style.
	Direction string `json:"Direction"`

	// Gradient shape.
	Shape string `json:"Shape"`

	// Gradient stops.
	Stops []GradientFillStop `json:"Stops,omitempty"`

	// Gradient angle.
	LinearAngle float64 `json:"LinearAngle"`

	// True if the gradient is scaled.
	IsScaled bool `json:"IsScaled"`

	// Gradient flipping mode.
	TileFlip string `json:"TileFlip"`
}

func (this GradientFill) getType() string {
	return this.Type_
}

func (this GradientFill) setType(newValue string) {
	this.Type_ = newValue
}
func (this GradientFill) getDirection() string {
	return this.Direction
}

func (this GradientFill) setDirection(newValue string) {
	this.Direction = newValue
}
func (this GradientFill) getShape() string {
	return this.Shape
}

func (this GradientFill) setShape(newValue string) {
	this.Shape = newValue
}
func (this GradientFill) getStops() []GradientFillStop {
	return this.Stops
}

func (this GradientFill) setStops(newValue []GradientFillStop) {
	this.Stops = newValue
}
func (this GradientFill) getLinearAngle() float64 {
	return this.LinearAngle
}

func (this GradientFill) setLinearAngle(newValue float64) {
	this.LinearAngle = newValue
}
func (this GradientFill) getIsScaled() bool {
	return this.IsScaled
}

func (this GradientFill) setIsScaled(newValue bool) {
	this.IsScaled = newValue
}
func (this GradientFill) getTileFlip() string {
	return this.TileFlip
}

func (this GradientFill) setTileFlip(newValue string) {
	this.TileFlip = newValue
}

func (this *GradientFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Gradient"
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
	this.Direction = "FromCorner1"
	if valDirection, ok := objMap["direction"]; ok {
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
	if valDirectionCap, ok := objMap["Direction"]; ok {
		if valDirectionCap != nil {
			var valueForDirection string
			err = json.Unmarshal(*valDirectionCap, &valueForDirection)
			if err != nil {
				var valueForDirectionInt int32
				err = json.Unmarshal(*valDirectionCap, &valueForDirectionInt)
				if err != nil {
					return err
				}
				this.Direction = string(valueForDirectionInt)
			} else {
				this.Direction = valueForDirection
			}
		}
	}
	this.Shape = "Linear"
	if valShape, ok := objMap["shape"]; ok {
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
	if valShapeCap, ok := objMap["Shape"]; ok {
		if valShapeCap != nil {
			var valueForShape string
			err = json.Unmarshal(*valShapeCap, &valueForShape)
			if err != nil {
				var valueForShapeInt int32
				err = json.Unmarshal(*valShapeCap, &valueForShapeInt)
				if err != nil {
					return err
				}
				this.Shape = string(valueForShapeInt)
			} else {
				this.Shape = valueForShape
			}
		}
	}
	
	if valStops, ok := objMap["stops"]; ok {
		if valStops != nil {
			var valueForStops []GradientFillStop
			err = json.Unmarshal(*valStops, &valueForStops)
			if err != nil {
				return err
			}
			this.Stops = valueForStops
		}
	}
	if valStopsCap, ok := objMap["Stops"]; ok {
		if valStopsCap != nil {
			var valueForStops []GradientFillStop
			err = json.Unmarshal(*valStopsCap, &valueForStops)
			if err != nil {
				return err
			}
			this.Stops = valueForStops
		}
	}
	
	if valLinearAngle, ok := objMap["linearAngle"]; ok {
		if valLinearAngle != nil {
			var valueForLinearAngle float64
			err = json.Unmarshal(*valLinearAngle, &valueForLinearAngle)
			if err != nil {
				return err
			}
			this.LinearAngle = valueForLinearAngle
		}
	}
	if valLinearAngleCap, ok := objMap["LinearAngle"]; ok {
		if valLinearAngleCap != nil {
			var valueForLinearAngle float64
			err = json.Unmarshal(*valLinearAngleCap, &valueForLinearAngle)
			if err != nil {
				return err
			}
			this.LinearAngle = valueForLinearAngle
		}
	}
	
	if valIsScaled, ok := objMap["isScaled"]; ok {
		if valIsScaled != nil {
			var valueForIsScaled bool
			err = json.Unmarshal(*valIsScaled, &valueForIsScaled)
			if err != nil {
				return err
			}
			this.IsScaled = valueForIsScaled
		}
	}
	if valIsScaledCap, ok := objMap["IsScaled"]; ok {
		if valIsScaledCap != nil {
			var valueForIsScaled bool
			err = json.Unmarshal(*valIsScaledCap, &valueForIsScaled)
			if err != nil {
				return err
			}
			this.IsScaled = valueForIsScaled
		}
	}
	this.TileFlip = "NoFlip"
	if valTileFlip, ok := objMap["tileFlip"]; ok {
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
	if valTileFlipCap, ok := objMap["TileFlip"]; ok {
		if valTileFlipCap != nil {
			var valueForTileFlip string
			err = json.Unmarshal(*valTileFlipCap, &valueForTileFlip)
			if err != nil {
				var valueForTileFlipInt int32
				err = json.Unmarshal(*valTileFlipCap, &valueForTileFlipInt)
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
