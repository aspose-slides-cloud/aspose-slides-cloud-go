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

// Picture fill.
type IPictureFill interface {

	// Fill type.
	getType() string
	setType(newValue string)

	// Percentage of image height that is cropped from the bottom.
	getCropBottom() float64
	setCropBottom(newValue float64)

	// Percentage of image height that is cropped from the left.
	getCropLeft() float64
	setCropLeft(newValue float64)

	// Percentage of image height that is cropped from the right.
	getCropRight() float64
	setCropRight(newValue float64)

	// Percentage of image height that is cropped from the top.
	getCropTop() float64
	setCropTop(newValue float64)

	// Picture resolution.
	getDpi() int32
	setDpi(newValue int32)

	// Internal image link.
	getImage() IResourceUriElement
	setImage(newValue IResourceUriElement)

	// Base 64 image data.
	getBase64Data() string
	setBase64Data(newValue string)

	// SVG image data.
	getSvgData() string
	setSvgData(newValue string)

	// Fill mode.
	getPictureFillMode() string
	setPictureFillMode(newValue string)
}

type PictureFill struct {

	// Fill type.
	Type_ string `json:"Type"`

	// Percentage of image height that is cropped from the bottom.
	CropBottom float64 `json:"CropBottom"`

	// Percentage of image height that is cropped from the left.
	CropLeft float64 `json:"CropLeft"`

	// Percentage of image height that is cropped from the right.
	CropRight float64 `json:"CropRight"`

	// Percentage of image height that is cropped from the top.
	CropTop float64 `json:"CropTop"`

	// Picture resolution.
	Dpi int32 `json:"Dpi"`

	// Internal image link.
	Image IResourceUriElement `json:"Image,omitempty"`

	// Base 64 image data.
	Base64Data string `json:"Base64Data,omitempty"`

	// SVG image data.
	SvgData string `json:"SvgData,omitempty"`

	// Fill mode.
	PictureFillMode string `json:"PictureFillMode"`
}

func NewPictureFill() *PictureFill {
	instance := new(PictureFill)
	instance.Type_ = "Picture"
	instance.PictureFillMode = "Tile"
	return instance
}

func (this *PictureFill) getType() string {
	return this.Type_
}

func (this *PictureFill) setType(newValue string) {
	this.Type_ = newValue
}
func (this *PictureFill) getCropBottom() float64 {
	return this.CropBottom
}

func (this *PictureFill) setCropBottom(newValue float64) {
	this.CropBottom = newValue
}
func (this *PictureFill) getCropLeft() float64 {
	return this.CropLeft
}

func (this *PictureFill) setCropLeft(newValue float64) {
	this.CropLeft = newValue
}
func (this *PictureFill) getCropRight() float64 {
	return this.CropRight
}

func (this *PictureFill) setCropRight(newValue float64) {
	this.CropRight = newValue
}
func (this *PictureFill) getCropTop() float64 {
	return this.CropTop
}

func (this *PictureFill) setCropTop(newValue float64) {
	this.CropTop = newValue
}
func (this *PictureFill) getDpi() int32 {
	return this.Dpi
}

func (this *PictureFill) setDpi(newValue int32) {
	this.Dpi = newValue
}
func (this *PictureFill) getImage() IResourceUriElement {
	return this.Image
}

func (this *PictureFill) setImage(newValue IResourceUriElement) {
	this.Image = newValue
}
func (this *PictureFill) getBase64Data() string {
	return this.Base64Data
}

func (this *PictureFill) setBase64Data(newValue string) {
	this.Base64Data = newValue
}
func (this *PictureFill) getSvgData() string {
	return this.SvgData
}

func (this *PictureFill) setSvgData(newValue string) {
	this.SvgData = newValue
}
func (this *PictureFill) getPictureFillMode() string {
	return this.PictureFillMode
}

func (this *PictureFill) setPictureFillMode(newValue string) {
	this.PictureFillMode = newValue
}

func (this *PictureFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Picture"
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
	
	if valCropBottom, ok := objMap["cropBottom"]; ok {
		if valCropBottom != nil {
			var valueForCropBottom float64
			err = json.Unmarshal(*valCropBottom, &valueForCropBottom)
			if err != nil {
				return err
			}
			this.CropBottom = valueForCropBottom
		}
	}
	if valCropBottomCap, ok := objMap["CropBottom"]; ok {
		if valCropBottomCap != nil {
			var valueForCropBottom float64
			err = json.Unmarshal(*valCropBottomCap, &valueForCropBottom)
			if err != nil {
				return err
			}
			this.CropBottom = valueForCropBottom
		}
	}
	
	if valCropLeft, ok := objMap["cropLeft"]; ok {
		if valCropLeft != nil {
			var valueForCropLeft float64
			err = json.Unmarshal(*valCropLeft, &valueForCropLeft)
			if err != nil {
				return err
			}
			this.CropLeft = valueForCropLeft
		}
	}
	if valCropLeftCap, ok := objMap["CropLeft"]; ok {
		if valCropLeftCap != nil {
			var valueForCropLeft float64
			err = json.Unmarshal(*valCropLeftCap, &valueForCropLeft)
			if err != nil {
				return err
			}
			this.CropLeft = valueForCropLeft
		}
	}
	
	if valCropRight, ok := objMap["cropRight"]; ok {
		if valCropRight != nil {
			var valueForCropRight float64
			err = json.Unmarshal(*valCropRight, &valueForCropRight)
			if err != nil {
				return err
			}
			this.CropRight = valueForCropRight
		}
	}
	if valCropRightCap, ok := objMap["CropRight"]; ok {
		if valCropRightCap != nil {
			var valueForCropRight float64
			err = json.Unmarshal(*valCropRightCap, &valueForCropRight)
			if err != nil {
				return err
			}
			this.CropRight = valueForCropRight
		}
	}
	
	if valCropTop, ok := objMap["cropTop"]; ok {
		if valCropTop != nil {
			var valueForCropTop float64
			err = json.Unmarshal(*valCropTop, &valueForCropTop)
			if err != nil {
				return err
			}
			this.CropTop = valueForCropTop
		}
	}
	if valCropTopCap, ok := objMap["CropTop"]; ok {
		if valCropTopCap != nil {
			var valueForCropTop float64
			err = json.Unmarshal(*valCropTopCap, &valueForCropTop)
			if err != nil {
				return err
			}
			this.CropTop = valueForCropTop
		}
	}
	
	if valDpi, ok := objMap["dpi"]; ok {
		if valDpi != nil {
			var valueForDpi int32
			err = json.Unmarshal(*valDpi, &valueForDpi)
			if err != nil {
				return err
			}
			this.Dpi = valueForDpi
		}
	}
	if valDpiCap, ok := objMap["Dpi"]; ok {
		if valDpiCap != nil {
			var valueForDpi int32
			err = json.Unmarshal(*valDpiCap, &valueForDpi)
			if err != nil {
				return err
			}
			this.Dpi = valueForDpi
		}
	}
	
	if valImage, ok := objMap["image"]; ok {
		if valImage != nil {
			var valueForImage ResourceUriElement
			err = json.Unmarshal(*valImage, &valueForImage)
			if err != nil {
				return err
			}
			this.Image = &valueForImage
		}
	}
	if valImageCap, ok := objMap["Image"]; ok {
		if valImageCap != nil {
			var valueForImage ResourceUriElement
			err = json.Unmarshal(*valImageCap, &valueForImage)
			if err != nil {
				return err
			}
			this.Image = &valueForImage
		}
	}
	
	if valBase64Data, ok := objMap["base64Data"]; ok {
		if valBase64Data != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64Data, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	if valBase64DataCap, ok := objMap["Base64Data"]; ok {
		if valBase64DataCap != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64DataCap, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	
	if valSvgData, ok := objMap["svgData"]; ok {
		if valSvgData != nil {
			var valueForSvgData string
			err = json.Unmarshal(*valSvgData, &valueForSvgData)
			if err != nil {
				return err
			}
			this.SvgData = valueForSvgData
		}
	}
	if valSvgDataCap, ok := objMap["SvgData"]; ok {
		if valSvgDataCap != nil {
			var valueForSvgData string
			err = json.Unmarshal(*valSvgDataCap, &valueForSvgData)
			if err != nil {
				return err
			}
			this.SvgData = valueForSvgData
		}
	}
	this.PictureFillMode = "Tile"
	if valPictureFillMode, ok := objMap["pictureFillMode"]; ok {
		if valPictureFillMode != nil {
			var valueForPictureFillMode string
			err = json.Unmarshal(*valPictureFillMode, &valueForPictureFillMode)
			if err != nil {
				var valueForPictureFillModeInt int32
				err = json.Unmarshal(*valPictureFillMode, &valueForPictureFillModeInt)
				if err != nil {
					return err
				}
				this.PictureFillMode = string(valueForPictureFillModeInt)
			} else {
				this.PictureFillMode = valueForPictureFillMode
			}
		}
	}
	if valPictureFillModeCap, ok := objMap["PictureFillMode"]; ok {
		if valPictureFillModeCap != nil {
			var valueForPictureFillMode string
			err = json.Unmarshal(*valPictureFillModeCap, &valueForPictureFillMode)
			if err != nil {
				var valueForPictureFillModeInt int32
				err = json.Unmarshal(*valPictureFillModeCap, &valueForPictureFillModeInt)
				if err != nil {
					return err
				}
				this.PictureFillMode = string(valueForPictureFillModeInt)
			} else {
				this.PictureFillMode = valueForPictureFillMode
			}
		}
	}

    return nil
}
