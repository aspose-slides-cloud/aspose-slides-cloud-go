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

// Represents a series marker
type ISeriesMarker interface {

	// size
	getSize() int32
	setSize(newValue int32)

	// symbol
	getSymbol() string
	setSymbol(newValue string)

	// Get or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Get or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Get or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)
}

type SeriesMarker struct {

	// size
	Size int32 `json:"Size,omitempty"`

	// symbol
	Symbol string `json:"Symbol,omitempty"`

	// Get or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`
}

func NewSeriesMarker() *SeriesMarker {
	instance := new(SeriesMarker)
	instance.Symbol = ""
	return instance
}

func (this *SeriesMarker) getSize() int32 {
	return this.Size
}

func (this *SeriesMarker) setSize(newValue int32) {
	this.Size = newValue
}
func (this *SeriesMarker) getSymbol() string {
	return this.Symbol
}

func (this *SeriesMarker) setSymbol(newValue string) {
	this.Symbol = newValue
}
func (this *SeriesMarker) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *SeriesMarker) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *SeriesMarker) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *SeriesMarker) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *SeriesMarker) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *SeriesMarker) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}

func (this *SeriesMarker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSize, ok := objMap["size"]; ok {
		if valSize != nil {
			var valueForSize int32
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}
	if valSizeCap, ok := objMap["Size"]; ok {
		if valSizeCap != nil {
			var valueForSize int32
			err = json.Unmarshal(*valSizeCap, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}
	this.Symbol = ""
	if valSymbol, ok := objMap["symbol"]; ok {
		if valSymbol != nil {
			var valueForSymbol string
			err = json.Unmarshal(*valSymbol, &valueForSymbol)
			if err != nil {
				var valueForSymbolInt int32
				err = json.Unmarshal(*valSymbol, &valueForSymbolInt)
				if err != nil {
					return err
				}
				this.Symbol = string(valueForSymbolInt)
			} else {
				this.Symbol = valueForSymbol
			}
		}
	}
	if valSymbolCap, ok := objMap["Symbol"]; ok {
		if valSymbolCap != nil {
			var valueForSymbol string
			err = json.Unmarshal(*valSymbolCap, &valueForSymbol)
			if err != nil {
				var valueForSymbolInt int32
				err = json.Unmarshal(*valSymbolCap, &valueForSymbolInt)
				if err != nil {
					return err
				}
				this.Symbol = string(valueForSymbolInt)
			} else {
				this.Symbol = valueForSymbol
			}
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}

	return nil
}
