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

// Represents font info.
type IFontData interface {

	// Font name
	GetFontName() string
	SetFontName(newValue string)

	// Returns true if font is embedded.
	GetIsEmbedded() bool
	SetIsEmbedded(newValue bool)
}

type FontData struct {

	// Font name
	FontName string `json:"FontName,omitempty"`

	// Returns true if font is embedded.
	IsEmbedded bool `json:"IsEmbedded"`
}

func NewFontData() *FontData {
	instance := new(FontData)
	return instance
}

func (this *FontData) GetFontName() string {
	return this.FontName
}

func (this *FontData) SetFontName(newValue string) {
	this.FontName = newValue
}
func (this *FontData) GetIsEmbedded() bool {
	return this.IsEmbedded
}

func (this *FontData) SetIsEmbedded(newValue bool) {
	this.IsEmbedded = newValue
}

func (this *FontData) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valFontName, ok := objMap["fontName"]; ok {
		if valFontName != nil {
			var valueForFontName string
			err = json.Unmarshal(*valFontName, &valueForFontName)
			if err != nil {
				return err
			}
			this.FontName = valueForFontName
		}
	}
	if valFontNameCap, ok := objMap["FontName"]; ok {
		if valFontNameCap != nil {
			var valueForFontName string
			err = json.Unmarshal(*valFontNameCap, &valueForFontName)
			if err != nil {
				return err
			}
			this.FontName = valueForFontName
		}
	}
	
	if valIsEmbedded, ok := objMap["isEmbedded"]; ok {
		if valIsEmbedded != nil {
			var valueForIsEmbedded bool
			err = json.Unmarshal(*valIsEmbedded, &valueForIsEmbedded)
			if err != nil {
				return err
			}
			this.IsEmbedded = valueForIsEmbedded
		}
	}
	if valIsEmbeddedCap, ok := objMap["IsEmbedded"]; ok {
		if valIsEmbeddedCap != nil {
			var valueForIsEmbedded bool
			err = json.Unmarshal(*valIsEmbeddedCap, &valueForIsEmbedded)
			if err != nil {
				return err
			}
			this.IsEmbedded = valueForIsEmbedded
		}
	}

	return nil
}
