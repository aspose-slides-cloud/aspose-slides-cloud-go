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

// Represents fill overlay effect 
type IFillOverlayEffect interface {

	// blend mode
	getBlend() string
	setBlend(newValue string)
}

type FillOverlayEffect struct {

	// blend mode
	Blend string `json:"Blend"`
}

func NewFillOverlayEffect() *FillOverlayEffect {
	instance := new(FillOverlayEffect)
	instance.Blend = "Darken"
	return instance
}

func (this *FillOverlayEffect) getBlend() string {
	return this.Blend
}

func (this *FillOverlayEffect) setBlend(newValue string) {
	this.Blend = newValue
}

func (this *FillOverlayEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Blend = "Darken"
	if valBlend, ok := objMap["blend"]; ok {
		if valBlend != nil {
			var valueForBlend string
			err = json.Unmarshal(*valBlend, &valueForBlend)
			if err != nil {
				var valueForBlendInt int32
				err = json.Unmarshal(*valBlend, &valueForBlendInt)
				if err != nil {
					return err
				}
				this.Blend = string(valueForBlendInt)
			} else {
				this.Blend = valueForBlend
			}
		}
	}
	if valBlendCap, ok := objMap["Blend"]; ok {
		if valBlendCap != nil {
			var valueForBlend string
			err = json.Unmarshal(*valBlendCap, &valueForBlend)
			if err != nil {
				var valueForBlendInt int32
				err = json.Unmarshal(*valBlendCap, &valueForBlendInt)
				if err != nil {
					return err
				}
				this.Blend = string(valueForBlendInt)
			} else {
				this.Blend = valueForBlend
			}
		}
	}

	return nil
}
