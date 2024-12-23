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

// Slide view properties.
type ICommonSlideViewProperties interface {

	// The view scaling ratio (percentage).
	GetScale() int32
	SetScale(newValue int32)

	// True if the view content should automatically scale to best fit the current window size.
	GetVariableScale() *bool
	SetVariableScale(newValue *bool)

	// Drawing guides
	GetDrawingGuides() []IDrawingGuide
	SetDrawingGuides(newValue []IDrawingGuide)
}

type CommonSlideViewProperties struct {

	// The view scaling ratio (percentage).
	Scale int32 `json:"Scale,omitempty"`

	// True if the view content should automatically scale to best fit the current window size.
	VariableScale *bool `json:"VariableScale"`

	// Drawing guides
	DrawingGuides []IDrawingGuide `json:"DrawingGuides,omitempty"`
}

func NewCommonSlideViewProperties() *CommonSlideViewProperties {
	instance := new(CommonSlideViewProperties)
	return instance
}

func (this *CommonSlideViewProperties) GetScale() int32 {
	return this.Scale
}

func (this *CommonSlideViewProperties) SetScale(newValue int32) {
	this.Scale = newValue
}
func (this *CommonSlideViewProperties) GetVariableScale() *bool {
	return this.VariableScale
}

func (this *CommonSlideViewProperties) SetVariableScale(newValue *bool) {
	this.VariableScale = newValue
}
func (this *CommonSlideViewProperties) GetDrawingGuides() []IDrawingGuide {
	return this.DrawingGuides
}

func (this *CommonSlideViewProperties) SetDrawingGuides(newValue []IDrawingGuide) {
	this.DrawingGuides = newValue
}

func (this *CommonSlideViewProperties) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valScale, ok := GetMapValue(objMap, "scale"); ok {
		if valScale != nil {
			var valueForScale int32
			err = json.Unmarshal(*valScale, &valueForScale)
			if err != nil {
				return err
			}
			this.Scale = valueForScale
		}
	}
	
	if valVariableScale, ok := GetMapValue(objMap, "variableScale"); ok {
		if valVariableScale != nil {
			var valueForVariableScale *bool
			err = json.Unmarshal(*valVariableScale, &valueForVariableScale)
			if err != nil {
				return err
			}
			this.VariableScale = valueForVariableScale
		}
	}
	
	if valDrawingGuides, ok := GetMapValue(objMap, "drawingGuides"); ok {
		if valDrawingGuides != nil {
			var valueForDrawingGuides []json.RawMessage
			err = json.Unmarshal(*valDrawingGuides, &valueForDrawingGuides)
			if err != nil {
				return err
			}
			valueForIDrawingGuides := make([]IDrawingGuide, len(valueForDrawingGuides))
			for i, v := range valueForDrawingGuides {
				vObject, err := createObjectForType("DrawingGuide", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIDrawingGuides[i] = vObject.(IDrawingGuide)
				}
			}
			this.DrawingGuides = valueForIDrawingGuides
		}
	}

	return nil
}
