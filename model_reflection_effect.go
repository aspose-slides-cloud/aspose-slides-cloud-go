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
	GetDirection() float64
	SetDirection(newValue float64)

	// fade direction
	GetFadeDirection() float64
	SetFadeDirection(newValue float64)

	// distance
	GetDistance() float64
	SetDistance(newValue float64)

	// blur radius
	GetBlurRadius() float64
	SetBlurRadius(newValue float64)

	// scale horizontal
	GetScaleHorizontal() float64
	SetScaleHorizontal(newValue float64)

	// scale vertical
	GetScaleVertical() float64
	SetScaleVertical(newValue float64)

	// skew horizontal
	GetSkewHorizontal() float64
	SetSkewHorizontal(newValue float64)

	// skew vertical
	GetSkewVertical() float64
	SetSkewVertical(newValue float64)

	// start pos alpha
	GetStartPosAlpha() float64
	SetStartPosAlpha(newValue float64)

	// end pos alpha
	GetEndPosAlpha() float64
	SetEndPosAlpha(newValue float64)

	// start reflection opacity
	GetStartReflectionOpacity() float64
	SetStartReflectionOpacity(newValue float64)

	// end reflection opacity
	GetEndReflectionOpacity() float64
	SetEndReflectionOpacity(newValue float64)

	// rectangle alignment
	GetRectangleAlign() string
	SetRectangleAlign(newValue string)

	// true if the reflection should rotate with the shape when the shape is rotated
	GetRotateShadowWithShape() *bool
	SetRotateShadowWithShape(newValue *bool)
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
	RectangleAlign string `json:"RectangleAlign"`

	// true if the reflection should rotate with the shape when the shape is rotated
	RotateShadowWithShape *bool `json:"RotateShadowWithShape"`
}

func NewReflectionEffect() *ReflectionEffect {
	instance := new(ReflectionEffect)
	instance.RectangleAlign = "TopLeft"
	return instance
}

func (this *ReflectionEffect) GetDirection() float64 {
	return this.Direction
}

func (this *ReflectionEffect) SetDirection(newValue float64) {
	this.Direction = newValue
}
func (this *ReflectionEffect) GetFadeDirection() float64 {
	return this.FadeDirection
}

func (this *ReflectionEffect) SetFadeDirection(newValue float64) {
	this.FadeDirection = newValue
}
func (this *ReflectionEffect) GetDistance() float64 {
	return this.Distance
}

func (this *ReflectionEffect) SetDistance(newValue float64) {
	this.Distance = newValue
}
func (this *ReflectionEffect) GetBlurRadius() float64 {
	return this.BlurRadius
}

func (this *ReflectionEffect) SetBlurRadius(newValue float64) {
	this.BlurRadius = newValue
}
func (this *ReflectionEffect) GetScaleHorizontal() float64 {
	return this.ScaleHorizontal
}

func (this *ReflectionEffect) SetScaleHorizontal(newValue float64) {
	this.ScaleHorizontal = newValue
}
func (this *ReflectionEffect) GetScaleVertical() float64 {
	return this.ScaleVertical
}

func (this *ReflectionEffect) SetScaleVertical(newValue float64) {
	this.ScaleVertical = newValue
}
func (this *ReflectionEffect) GetSkewHorizontal() float64 {
	return this.SkewHorizontal
}

func (this *ReflectionEffect) SetSkewHorizontal(newValue float64) {
	this.SkewHorizontal = newValue
}
func (this *ReflectionEffect) GetSkewVertical() float64 {
	return this.SkewVertical
}

func (this *ReflectionEffect) SetSkewVertical(newValue float64) {
	this.SkewVertical = newValue
}
func (this *ReflectionEffect) GetStartPosAlpha() float64 {
	return this.StartPosAlpha
}

func (this *ReflectionEffect) SetStartPosAlpha(newValue float64) {
	this.StartPosAlpha = newValue
}
func (this *ReflectionEffect) GetEndPosAlpha() float64 {
	return this.EndPosAlpha
}

func (this *ReflectionEffect) SetEndPosAlpha(newValue float64) {
	this.EndPosAlpha = newValue
}
func (this *ReflectionEffect) GetStartReflectionOpacity() float64 {
	return this.StartReflectionOpacity
}

func (this *ReflectionEffect) SetStartReflectionOpacity(newValue float64) {
	this.StartReflectionOpacity = newValue
}
func (this *ReflectionEffect) GetEndReflectionOpacity() float64 {
	return this.EndReflectionOpacity
}

func (this *ReflectionEffect) SetEndReflectionOpacity(newValue float64) {
	this.EndReflectionOpacity = newValue
}
func (this *ReflectionEffect) GetRectangleAlign() string {
	return this.RectangleAlign
}

func (this *ReflectionEffect) SetRectangleAlign(newValue string) {
	this.RectangleAlign = newValue
}
func (this *ReflectionEffect) GetRotateShadowWithShape() *bool {
	return this.RotateShadowWithShape
}

func (this *ReflectionEffect) SetRotateShadowWithShape(newValue *bool) {
	this.RotateShadowWithShape = newValue
}

func (this *ReflectionEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDirection, ok := GetMapValue(objMap, "direction"); ok {
		if valDirection != nil {
			var valueForDirection float64
			err = json.Unmarshal(*valDirection, &valueForDirection)
			if err != nil {
				return err
			}
			this.Direction = valueForDirection
		}
	}
	
	if valFadeDirection, ok := GetMapValue(objMap, "fadeDirection"); ok {
		if valFadeDirection != nil {
			var valueForFadeDirection float64
			err = json.Unmarshal(*valFadeDirection, &valueForFadeDirection)
			if err != nil {
				return err
			}
			this.FadeDirection = valueForFadeDirection
		}
	}
	
	if valDistance, ok := GetMapValue(objMap, "distance"); ok {
		if valDistance != nil {
			var valueForDistance float64
			err = json.Unmarshal(*valDistance, &valueForDistance)
			if err != nil {
				return err
			}
			this.Distance = valueForDistance
		}
	}
	
	if valBlurRadius, ok := GetMapValue(objMap, "blurRadius"); ok {
		if valBlurRadius != nil {
			var valueForBlurRadius float64
			err = json.Unmarshal(*valBlurRadius, &valueForBlurRadius)
			if err != nil {
				return err
			}
			this.BlurRadius = valueForBlurRadius
		}
	}
	
	if valScaleHorizontal, ok := GetMapValue(objMap, "scaleHorizontal"); ok {
		if valScaleHorizontal != nil {
			var valueForScaleHorizontal float64
			err = json.Unmarshal(*valScaleHorizontal, &valueForScaleHorizontal)
			if err != nil {
				return err
			}
			this.ScaleHorizontal = valueForScaleHorizontal
		}
	}
	
	if valScaleVertical, ok := GetMapValue(objMap, "scaleVertical"); ok {
		if valScaleVertical != nil {
			var valueForScaleVertical float64
			err = json.Unmarshal(*valScaleVertical, &valueForScaleVertical)
			if err != nil {
				return err
			}
			this.ScaleVertical = valueForScaleVertical
		}
	}
	
	if valSkewHorizontal, ok := GetMapValue(objMap, "skewHorizontal"); ok {
		if valSkewHorizontal != nil {
			var valueForSkewHorizontal float64
			err = json.Unmarshal(*valSkewHorizontal, &valueForSkewHorizontal)
			if err != nil {
				return err
			}
			this.SkewHorizontal = valueForSkewHorizontal
		}
	}
	
	if valSkewVertical, ok := GetMapValue(objMap, "skewVertical"); ok {
		if valSkewVertical != nil {
			var valueForSkewVertical float64
			err = json.Unmarshal(*valSkewVertical, &valueForSkewVertical)
			if err != nil {
				return err
			}
			this.SkewVertical = valueForSkewVertical
		}
	}
	
	if valStartPosAlpha, ok := GetMapValue(objMap, "startPosAlpha"); ok {
		if valStartPosAlpha != nil {
			var valueForStartPosAlpha float64
			err = json.Unmarshal(*valStartPosAlpha, &valueForStartPosAlpha)
			if err != nil {
				return err
			}
			this.StartPosAlpha = valueForStartPosAlpha
		}
	}
	
	if valEndPosAlpha, ok := GetMapValue(objMap, "endPosAlpha"); ok {
		if valEndPosAlpha != nil {
			var valueForEndPosAlpha float64
			err = json.Unmarshal(*valEndPosAlpha, &valueForEndPosAlpha)
			if err != nil {
				return err
			}
			this.EndPosAlpha = valueForEndPosAlpha
		}
	}
	
	if valStartReflectionOpacity, ok := GetMapValue(objMap, "startReflectionOpacity"); ok {
		if valStartReflectionOpacity != nil {
			var valueForStartReflectionOpacity float64
			err = json.Unmarshal(*valStartReflectionOpacity, &valueForStartReflectionOpacity)
			if err != nil {
				return err
			}
			this.StartReflectionOpacity = valueForStartReflectionOpacity
		}
	}
	
	if valEndReflectionOpacity, ok := GetMapValue(objMap, "endReflectionOpacity"); ok {
		if valEndReflectionOpacity != nil {
			var valueForEndReflectionOpacity float64
			err = json.Unmarshal(*valEndReflectionOpacity, &valueForEndReflectionOpacity)
			if err != nil {
				return err
			}
			this.EndReflectionOpacity = valueForEndReflectionOpacity
		}
	}
	this.RectangleAlign = "TopLeft"
	if valRectangleAlign, ok := GetMapValue(objMap, "rectangleAlign"); ok {
		if valRectangleAlign != nil {
			var valueForRectangleAlign string
			err = json.Unmarshal(*valRectangleAlign, &valueForRectangleAlign)
			if err != nil {
				var valueForRectangleAlignInt int32
				err = json.Unmarshal(*valRectangleAlign, &valueForRectangleAlignInt)
				if err != nil {
					return err
				}
				this.RectangleAlign = string(valueForRectangleAlignInt)
			} else {
				this.RectangleAlign = valueForRectangleAlign
			}
		}
	}
	
	if valRotateShadowWithShape, ok := GetMapValue(objMap, "rotateShadowWithShape"); ok {
		if valRotateShadowWithShape != nil {
			var valueForRotateShadowWithShape *bool
			err = json.Unmarshal(*valRotateShadowWithShape, &valueForRotateShadowWithShape)
			if err != nil {
				return err
			}
			this.RotateShadowWithShape = valueForRotateShadowWithShape
		}
	}

	return nil
}
