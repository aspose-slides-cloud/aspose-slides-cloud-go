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

// Represents soft edge effect 
type ISoftEdgeEffect interface {

	// radius
	GetRadius() float64
	SetRadius(newValue float64)
}

type SoftEdgeEffect struct {

	// radius
	Radius float64 `json:"Radius"`
}

func NewSoftEdgeEffect() *SoftEdgeEffect {
	instance := new(SoftEdgeEffect)
	return instance
}

func (this *SoftEdgeEffect) GetRadius() float64 {
	return this.Radius
}

func (this *SoftEdgeEffect) SetRadius(newValue float64) {
	this.Radius = newValue
}

func (this *SoftEdgeEffect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valRadius, ok := GetMapValue(objMap, "radius"); ok {
		if valRadius != nil {
			var valueForRadius float64
			err = json.Unmarshal(*valRadius, &valueForRadius)
			if err != nil {
				return err
			}
			this.Radius = valueForRadius
		}
	}

	return nil
}
