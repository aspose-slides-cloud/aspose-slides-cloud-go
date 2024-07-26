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
	GetType() string
	SetType(newValue string)

	// Percentage of image height that is cropped from the bottom.
	GetCropBottom() float64
	SetCropBottom(newValue float64)

	// Percentage of image height that is cropped from the left.
	GetCropLeft() float64
	SetCropLeft(newValue float64)

	// Percentage of image height that is cropped from the right.
	GetCropRight() float64
	SetCropRight(newValue float64)

	// Percentage of image height that is cropped from the top.
	GetCropTop() float64
	SetCropTop(newValue float64)

	// Picture resolution.
	GetDpi() int32
	SetDpi(newValue int32)

	// Internal image link.
	GetImage() IResourceUri
	SetImage(newValue IResourceUri)

	// Base 64 image data.
	GetBase64Data() string
	SetBase64Data(newValue string)

	// SVG image data.
	GetSvgData() string
	SetSvgData(newValue string)

	// Fill mode.
	GetPictureFillMode() string
	SetPictureFillMode(newValue string)

	// Image transform effects.
	GetImageTransformList() []IImageTransformEffect
	SetImageTransformList(newValue []IImageTransformEffect)
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
	Image IResourceUri `json:"Image,omitempty"`

	// Base 64 image data.
	Base64Data string `json:"Base64Data,omitempty"`

	// SVG image data.
	SvgData string `json:"SvgData,omitempty"`

	// Fill mode.
	PictureFillMode string `json:"PictureFillMode"`

	// Image transform effects.
	ImageTransformList []IImageTransformEffect `json:"ImageTransformList,omitempty"`
}

func NewPictureFill() *PictureFill {
	instance := new(PictureFill)
	instance.Type_ = "Picture"
	instance.PictureFillMode = "Tile"
	return instance
}

func (this *PictureFill) GetType() string {
	return this.Type_
}

func (this *PictureFill) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *PictureFill) GetCropBottom() float64 {
	return this.CropBottom
}

func (this *PictureFill) SetCropBottom(newValue float64) {
	this.CropBottom = newValue
}
func (this *PictureFill) GetCropLeft() float64 {
	return this.CropLeft
}

func (this *PictureFill) SetCropLeft(newValue float64) {
	this.CropLeft = newValue
}
func (this *PictureFill) GetCropRight() float64 {
	return this.CropRight
}

func (this *PictureFill) SetCropRight(newValue float64) {
	this.CropRight = newValue
}
func (this *PictureFill) GetCropTop() float64 {
	return this.CropTop
}

func (this *PictureFill) SetCropTop(newValue float64) {
	this.CropTop = newValue
}
func (this *PictureFill) GetDpi() int32 {
	return this.Dpi
}

func (this *PictureFill) SetDpi(newValue int32) {
	this.Dpi = newValue
}
func (this *PictureFill) GetImage() IResourceUri {
	return this.Image
}

func (this *PictureFill) SetImage(newValue IResourceUri) {
	this.Image = newValue
}
func (this *PictureFill) GetBase64Data() string {
	return this.Base64Data
}

func (this *PictureFill) SetBase64Data(newValue string) {
	this.Base64Data = newValue
}
func (this *PictureFill) GetSvgData() string {
	return this.SvgData
}

func (this *PictureFill) SetSvgData(newValue string) {
	this.SvgData = newValue
}
func (this *PictureFill) GetPictureFillMode() string {
	return this.PictureFillMode
}

func (this *PictureFill) SetPictureFillMode(newValue string) {
	this.PictureFillMode = newValue
}
func (this *PictureFill) GetImageTransformList() []IImageTransformEffect {
	return this.ImageTransformList
}

func (this *PictureFill) SetImageTransformList(newValue []IImageTransformEffect) {
	this.ImageTransformList = newValue
}

func (this *PictureFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Picture"
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valCropBottom, ok := GetMapValue(objMap, "cropBottom"); ok {
		if valCropBottom != nil {
			var valueForCropBottom float64
			err = json.Unmarshal(*valCropBottom, &valueForCropBottom)
			if err != nil {
				return err
			}
			this.CropBottom = valueForCropBottom
		}
	}
	
	if valCropLeft, ok := GetMapValue(objMap, "cropLeft"); ok {
		if valCropLeft != nil {
			var valueForCropLeft float64
			err = json.Unmarshal(*valCropLeft, &valueForCropLeft)
			if err != nil {
				return err
			}
			this.CropLeft = valueForCropLeft
		}
	}
	
	if valCropRight, ok := GetMapValue(objMap, "cropRight"); ok {
		if valCropRight != nil {
			var valueForCropRight float64
			err = json.Unmarshal(*valCropRight, &valueForCropRight)
			if err != nil {
				return err
			}
			this.CropRight = valueForCropRight
		}
	}
	
	if valCropTop, ok := GetMapValue(objMap, "cropTop"); ok {
		if valCropTop != nil {
			var valueForCropTop float64
			err = json.Unmarshal(*valCropTop, &valueForCropTop)
			if err != nil {
				return err
			}
			this.CropTop = valueForCropTop
		}
	}
	
	if valDpi, ok := GetMapValue(objMap, "dpi"); ok {
		if valDpi != nil {
			var valueForDpi int32
			err = json.Unmarshal(*valDpi, &valueForDpi)
			if err != nil {
				return err
			}
			this.Dpi = valueForDpi
		}
	}
	
	if valImage, ok := GetMapValue(objMap, "image"); ok {
		if valImage != nil {
			var valueForImage ResourceUri
			err = json.Unmarshal(*valImage, &valueForImage)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valImage)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valImage, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Image = vInterfaceObject
			}
		}
	}
	
	if valBase64Data, ok := GetMapValue(objMap, "base64Data"); ok {
		if valBase64Data != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64Data, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	
	if valSvgData, ok := GetMapValue(objMap, "svgData"); ok {
		if valSvgData != nil {
			var valueForSvgData string
			err = json.Unmarshal(*valSvgData, &valueForSvgData)
			if err != nil {
				return err
			}
			this.SvgData = valueForSvgData
		}
	}
	this.PictureFillMode = "Tile"
	if valPictureFillMode, ok := GetMapValue(objMap, "pictureFillMode"); ok {
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
	
	if valImageTransformList, ok := GetMapValue(objMap, "imageTransformList"); ok {
		if valImageTransformList != nil {
			var valueForImageTransformList []json.RawMessage
			err = json.Unmarshal(*valImageTransformList, &valueForImageTransformList)
			if err != nil {
				return err
			}
			valueForIImageTransformList := make([]IImageTransformEffect, len(valueForImageTransformList))
			for i, v := range valueForImageTransformList {
				vObject, err := createObjectForType("ImageTransformEffect", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIImageTransformList[i] = vObject.(IImageTransformEffect)
				}
			}
			this.ImageTransformList = valueForIImageTransformList
		}
	}

	return nil
}
