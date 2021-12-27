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

// Provides options that control how a shape is saved in thumbnail.
type IShapeImageExportOptions interface {

	// Get or sets scaling ratio by X axis.
	getScaleX() float64
	setScaleX(newValue float64)

	// Get or sets scaling ratio by Y axis.
	getScaleY() float64
	setScaleY(newValue float64)

	// Get or sets thumbnail bounds
	getThumbnailBounds() string
	setThumbnailBounds(newValue string)

	// Gets export format.
	getFormat() string
	setFormat(newValue string)
}

type ShapeImageExportOptions struct {

	// Get or sets scaling ratio by X axis.
	ScaleX float64 `json:"ScaleX,omitempty"`

	// Get or sets scaling ratio by Y axis.
	ScaleY float64 `json:"ScaleY,omitempty"`

	// Get or sets thumbnail bounds
	ThumbnailBounds string `json:"ThumbnailBounds,omitempty"`

	// Gets export format.
	Format string `json:"Format,omitempty"`
}

func NewShapeImageExportOptions() *ShapeImageExportOptions {
	instance := new(ShapeImageExportOptions)
	instance.ThumbnailBounds = ""
	return instance
}

func (this *ShapeImageExportOptions) getScaleX() float64 {
	return this.ScaleX
}

func (this *ShapeImageExportOptions) setScaleX(newValue float64) {
	this.ScaleX = newValue
}
func (this *ShapeImageExportOptions) getScaleY() float64 {
	return this.ScaleY
}

func (this *ShapeImageExportOptions) setScaleY(newValue float64) {
	this.ScaleY = newValue
}
func (this *ShapeImageExportOptions) getThumbnailBounds() string {
	return this.ThumbnailBounds
}

func (this *ShapeImageExportOptions) setThumbnailBounds(newValue string) {
	this.ThumbnailBounds = newValue
}
func (this *ShapeImageExportOptions) getFormat() string {
	return this.Format
}

func (this *ShapeImageExportOptions) setFormat(newValue string) {
	this.Format = newValue
}

func (this *ShapeImageExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valScaleX, ok := objMap["scaleX"]; ok {
		if valScaleX != nil {
			var valueForScaleX float64
			err = json.Unmarshal(*valScaleX, &valueForScaleX)
			if err != nil {
				return err
			}
			this.ScaleX = valueForScaleX
		}
	}
	if valScaleXCap, ok := objMap["ScaleX"]; ok {
		if valScaleXCap != nil {
			var valueForScaleX float64
			err = json.Unmarshal(*valScaleXCap, &valueForScaleX)
			if err != nil {
				return err
			}
			this.ScaleX = valueForScaleX
		}
	}
	
	if valScaleY, ok := objMap["scaleY"]; ok {
		if valScaleY != nil {
			var valueForScaleY float64
			err = json.Unmarshal(*valScaleY, &valueForScaleY)
			if err != nil {
				return err
			}
			this.ScaleY = valueForScaleY
		}
	}
	if valScaleYCap, ok := objMap["ScaleY"]; ok {
		if valScaleYCap != nil {
			var valueForScaleY float64
			err = json.Unmarshal(*valScaleYCap, &valueForScaleY)
			if err != nil {
				return err
			}
			this.ScaleY = valueForScaleY
		}
	}
	this.ThumbnailBounds = ""
	if valThumbnailBounds, ok := objMap["thumbnailBounds"]; ok {
		if valThumbnailBounds != nil {
			var valueForThumbnailBounds string
			err = json.Unmarshal(*valThumbnailBounds, &valueForThumbnailBounds)
			if err != nil {
				var valueForThumbnailBoundsInt int32
				err = json.Unmarshal(*valThumbnailBounds, &valueForThumbnailBoundsInt)
				if err != nil {
					return err
				}
				this.ThumbnailBounds = string(valueForThumbnailBoundsInt)
			} else {
				this.ThumbnailBounds = valueForThumbnailBounds
			}
		}
	}
	if valThumbnailBoundsCap, ok := objMap["ThumbnailBounds"]; ok {
		if valThumbnailBoundsCap != nil {
			var valueForThumbnailBounds string
			err = json.Unmarshal(*valThumbnailBoundsCap, &valueForThumbnailBounds)
			if err != nil {
				var valueForThumbnailBoundsInt int32
				err = json.Unmarshal(*valThumbnailBoundsCap, &valueForThumbnailBoundsInt)
				if err != nil {
					return err
				}
				this.ThumbnailBounds = string(valueForThumbnailBoundsInt)
			} else {
				this.ThumbnailBounds = valueForThumbnailBounds
			}
		}
	}
	
	if valFormat, ok := objMap["format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}

	return nil
}
