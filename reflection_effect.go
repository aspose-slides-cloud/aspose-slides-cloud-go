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

// Represents reflection effect 
type IReflectionEffect interface {

	// direction
	getDirection() float64
	setDirection(newValue float64)

	// fade direction
	getFadeDirection() float64
	setFadeDirection(newValue float64)

	// distance
	getDistance() float64
	setDistance(newValue float64)

	// blur radius
	getBlurRadius() float64
	setBlurRadius(newValue float64)

	// scale horizontal
	getScaleHorizontal() float64
	setScaleHorizontal(newValue float64)

	// scale vertical
	getScaleVertical() float64
	setScaleVertical(newValue float64)

	// skew horizontal
	getSkewHorizontal() float64
	setSkewHorizontal(newValue float64)

	// skew vertical
	getSkewVertical() float64
	setSkewVertical(newValue float64)

	// start pos alpha
	getStartPosAlpha() float64
	setStartPosAlpha(newValue float64)

	// end pos alpha
	getEndPosAlpha() float64
	setEndPosAlpha(newValue float64)

	// start reflection opacity
	getStartReflectionOpacity() float64
	setStartReflectionOpacity(newValue float64)

	// end reflection opacity
	getEndReflectionOpacity() float64
	setEndReflectionOpacity(newValue float64)

	// rectangle alignment
	getRectangleAlign() RectangleAlignment
	setRectangleAlign(newValue RectangleAlignment)

	// true if the reflection should rotate with the shape when the shape is rotated
	getRotateShadowWithShape() bool
	setRotateShadowWithShape(newValue bool)
}

type ReflectionEffect struct {

	// direction
	Direction float64 `json:"Direction"`

	// fade direction
	FadeDirection float64 `json:"FadeDirection"`

	// distance
	Distance float64 `json:"Distance"`

	// blur radius
	BlurRadius float64 `json:"BlurRadius"`

	// scale horizontal
	ScaleHorizontal float64 `json:"ScaleHorizontal"`

	// scale vertical
	ScaleVertical float64 `json:"ScaleVertical"`

	// skew horizontal
	SkewHorizontal float64 `json:"SkewHorizontal"`

	// skew vertical
	SkewVertical float64 `json:"SkewVertical"`

	// start pos alpha
	StartPosAlpha float64 `json:"StartPosAlpha"`

	// end pos alpha
	EndPosAlpha float64 `json:"EndPosAlpha"`

	// start reflection opacity
	StartReflectionOpacity float64 `json:"StartReflectionOpacity"`

	// end reflection opacity
	EndReflectionOpacity float64 `json:"EndReflectionOpacity"`

	// rectangle alignment
	RectangleAlign RectangleAlignment `json:"RectangleAlign"`

	// true if the reflection should rotate with the shape when the shape is rotated
	RotateShadowWithShape bool `json:"RotateShadowWithShape"`
}

func (this ReflectionEffect) getDirection() float64 {
	return this.Direction
}

func (this ReflectionEffect) setDirection(newValue float64) {
	this.Direction = newValue
}
func (this ReflectionEffect) getFadeDirection() float64 {
	return this.FadeDirection
}

func (this ReflectionEffect) setFadeDirection(newValue float64) {
	this.FadeDirection = newValue
}
func (this ReflectionEffect) getDistance() float64 {
	return this.Distance
}

func (this ReflectionEffect) setDistance(newValue float64) {
	this.Distance = newValue
}
func (this ReflectionEffect) getBlurRadius() float64 {
	return this.BlurRadius
}

func (this ReflectionEffect) setBlurRadius(newValue float64) {
	this.BlurRadius = newValue
}
func (this ReflectionEffect) getScaleHorizontal() float64 {
	return this.ScaleHorizontal
}

func (this ReflectionEffect) setScaleHorizontal(newValue float64) {
	this.ScaleHorizontal = newValue
}
func (this ReflectionEffect) getScaleVertical() float64 {
	return this.ScaleVertical
}

func (this ReflectionEffect) setScaleVertical(newValue float64) {
	this.ScaleVertical = newValue
}
func (this ReflectionEffect) getSkewHorizontal() float64 {
	return this.SkewHorizontal
}

func (this ReflectionEffect) setSkewHorizontal(newValue float64) {
	this.SkewHorizontal = newValue
}
func (this ReflectionEffect) getSkewVertical() float64 {
	return this.SkewVertical
}

func (this ReflectionEffect) setSkewVertical(newValue float64) {
	this.SkewVertical = newValue
}
func (this ReflectionEffect) getStartPosAlpha() float64 {
	return this.StartPosAlpha
}

func (this ReflectionEffect) setStartPosAlpha(newValue float64) {
	this.StartPosAlpha = newValue
}
func (this ReflectionEffect) getEndPosAlpha() float64 {
	return this.EndPosAlpha
}

func (this ReflectionEffect) setEndPosAlpha(newValue float64) {
	this.EndPosAlpha = newValue
}
func (this ReflectionEffect) getStartReflectionOpacity() float64 {
	return this.StartReflectionOpacity
}

func (this ReflectionEffect) setStartReflectionOpacity(newValue float64) {
	this.StartReflectionOpacity = newValue
}
func (this ReflectionEffect) getEndReflectionOpacity() float64 {
	return this.EndReflectionOpacity
}

func (this ReflectionEffect) setEndReflectionOpacity(newValue float64) {
	this.EndReflectionOpacity = newValue
}
func (this ReflectionEffect) getRectangleAlign() RectangleAlignment {
	return this.RectangleAlign
}

func (this ReflectionEffect) setRectangleAlign(newValue RectangleAlignment) {
	this.RectangleAlign = newValue
}
func (this ReflectionEffect) getRotateShadowWithShape() bool {
	return this.RotateShadowWithShape
}

func (this ReflectionEffect) setRotateShadowWithShape(newValue bool) {
	this.RotateShadowWithShape = newValue
}

func (this *ReflectionEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valDirection, ok := objMap["Direction"]; ok {
		if valDirection != nil {
			var valueForDirection float64
			err = json.Unmarshal(*valDirection, &valueForDirection)
			if err != nil {
				return err
			}
			this.Direction = valueForDirection
		}
	}

	if valFadeDirection, ok := objMap["FadeDirection"]; ok {
		if valFadeDirection != nil {
			var valueForFadeDirection float64
			err = json.Unmarshal(*valFadeDirection, &valueForFadeDirection)
			if err != nil {
				return err
			}
			this.FadeDirection = valueForFadeDirection
		}
	}

	if valDistance, ok := objMap["Distance"]; ok {
		if valDistance != nil {
			var valueForDistance float64
			err = json.Unmarshal(*valDistance, &valueForDistance)
			if err != nil {
				return err
			}
			this.Distance = valueForDistance
		}
	}

	if valBlurRadius, ok := objMap["BlurRadius"]; ok {
		if valBlurRadius != nil {
			var valueForBlurRadius float64
			err = json.Unmarshal(*valBlurRadius, &valueForBlurRadius)
			if err != nil {
				return err
			}
			this.BlurRadius = valueForBlurRadius
		}
	}

	if valScaleHorizontal, ok := objMap["ScaleHorizontal"]; ok {
		if valScaleHorizontal != nil {
			var valueForScaleHorizontal float64
			err = json.Unmarshal(*valScaleHorizontal, &valueForScaleHorizontal)
			if err != nil {
				return err
			}
			this.ScaleHorizontal = valueForScaleHorizontal
		}
	}

	if valScaleVertical, ok := objMap["ScaleVertical"]; ok {
		if valScaleVertical != nil {
			var valueForScaleVertical float64
			err = json.Unmarshal(*valScaleVertical, &valueForScaleVertical)
			if err != nil {
				return err
			}
			this.ScaleVertical = valueForScaleVertical
		}
	}

	if valSkewHorizontal, ok := objMap["SkewHorizontal"]; ok {
		if valSkewHorizontal != nil {
			var valueForSkewHorizontal float64
			err = json.Unmarshal(*valSkewHorizontal, &valueForSkewHorizontal)
			if err != nil {
				return err
			}
			this.SkewHorizontal = valueForSkewHorizontal
		}
	}

	if valSkewVertical, ok := objMap["SkewVertical"]; ok {
		if valSkewVertical != nil {
			var valueForSkewVertical float64
			err = json.Unmarshal(*valSkewVertical, &valueForSkewVertical)
			if err != nil {
				return err
			}
			this.SkewVertical = valueForSkewVertical
		}
	}

	if valStartPosAlpha, ok := objMap["StartPosAlpha"]; ok {
		if valStartPosAlpha != nil {
			var valueForStartPosAlpha float64
			err = json.Unmarshal(*valStartPosAlpha, &valueForStartPosAlpha)
			if err != nil {
				return err
			}
			this.StartPosAlpha = valueForStartPosAlpha
		}
	}

	if valEndPosAlpha, ok := objMap["EndPosAlpha"]; ok {
		if valEndPosAlpha != nil {
			var valueForEndPosAlpha float64
			err = json.Unmarshal(*valEndPosAlpha, &valueForEndPosAlpha)
			if err != nil {
				return err
			}
			this.EndPosAlpha = valueForEndPosAlpha
		}
	}

	if valStartReflectionOpacity, ok := objMap["StartReflectionOpacity"]; ok {
		if valStartReflectionOpacity != nil {
			var valueForStartReflectionOpacity float64
			err = json.Unmarshal(*valStartReflectionOpacity, &valueForStartReflectionOpacity)
			if err != nil {
				return err
			}
			this.StartReflectionOpacity = valueForStartReflectionOpacity
		}
	}

	if valEndReflectionOpacity, ok := objMap["EndReflectionOpacity"]; ok {
		if valEndReflectionOpacity != nil {
			var valueForEndReflectionOpacity float64
			err = json.Unmarshal(*valEndReflectionOpacity, &valueForEndReflectionOpacity)
			if err != nil {
				return err
			}
			this.EndReflectionOpacity = valueForEndReflectionOpacity
		}
	}

	if valRectangleAlign, ok := objMap["RectangleAlign"]; ok {
		if valRectangleAlign != nil {
			var valueForRectangleAlign RectangleAlignment
			err = json.Unmarshal(*valRectangleAlign, &valueForRectangleAlign)
			if err != nil {
				return err
			}
			this.RectangleAlign = valueForRectangleAlign
		}
	}

	if valRotateShadowWithShape, ok := objMap["RotateShadowWithShape"]; ok {
		if valRotateShadowWithShape != nil {
			var valueForRotateShadowWithShape bool
			err = json.Unmarshal(*valRotateShadowWithShape, &valueForRotateShadowWithShape)
			if err != nil {
				return err
			}
			this.RotateShadowWithShape = valueForRotateShadowWithShape
		}
	}

    return nil
}
