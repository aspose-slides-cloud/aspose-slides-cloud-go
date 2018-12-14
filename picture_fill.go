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


type IPictureFill interface {

	getType() FillType
	setType(newValue FillType)

	getCropBottom() float64
	setCropBottom(newValue float64)

	getCropLeft() float64
	setCropLeft(newValue float64)

	getCropRight() float64
	setCropRight(newValue float64)

	getCropTop() float64
	setCropTop(newValue float64)

	getDpi() int32
	setDpi(newValue int32)

	getImage() IResourceUriElement
	setImage(newValue IResourceUriElement)

	getBase64Data() string
	setBase64Data(newValue string)

	getSvgData() string
	setSvgData(newValue string)

	getPictureFillMode() PictureFillMode
	setPictureFillMode(newValue PictureFillMode)
}

type PictureFill struct {

	Type_ FillType `json:"Type"`

	CropBottom float64 `json:"CropBottom,omitempty"`

	CropLeft float64 `json:"CropLeft,omitempty"`

	CropRight float64 `json:"CropRight,omitempty"`

	CropTop float64 `json:"CropTop,omitempty"`

	Dpi int32 `json:"Dpi,omitempty"`

	Image IResourceUriElement `json:"Image,omitempty"`

	Base64Data string `json:"Base64Data,omitempty"`

	SvgData string `json:"SvgData,omitempty"`

	PictureFillMode PictureFillMode `json:"PictureFillMode,omitempty"`
}

func (this PictureFill) getType() FillType {
	return this.Type_
}

func (this PictureFill) setType(newValue FillType) {
	this.Type_ = newValue
}
func (this PictureFill) getCropBottom() float64 {
	return this.CropBottom
}

func (this PictureFill) setCropBottom(newValue float64) {
	this.CropBottom = newValue
}
func (this PictureFill) getCropLeft() float64 {
	return this.CropLeft
}

func (this PictureFill) setCropLeft(newValue float64) {
	this.CropLeft = newValue
}
func (this PictureFill) getCropRight() float64 {
	return this.CropRight
}

func (this PictureFill) setCropRight(newValue float64) {
	this.CropRight = newValue
}
func (this PictureFill) getCropTop() float64 {
	return this.CropTop
}

func (this PictureFill) setCropTop(newValue float64) {
	this.CropTop = newValue
}
func (this PictureFill) getDpi() int32 {
	return this.Dpi
}

func (this PictureFill) setDpi(newValue int32) {
	this.Dpi = newValue
}
func (this PictureFill) getImage() IResourceUriElement {
	return this.Image
}

func (this PictureFill) setImage(newValue IResourceUriElement) {
	this.Image = newValue
}
func (this PictureFill) getBase64Data() string {
	return this.Base64Data
}

func (this PictureFill) setBase64Data(newValue string) {
	this.Base64Data = newValue
}
func (this PictureFill) getSvgData() string {
	return this.SvgData
}

func (this PictureFill) setSvgData(newValue string) {
	this.SvgData = newValue
}
func (this PictureFill) getPictureFillMode() PictureFillMode {
	return this.PictureFillMode
}

func (this PictureFill) setPictureFillMode(newValue PictureFillMode) {
	this.PictureFillMode = newValue
}

func (this *PictureFill) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType FillType
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valCropBottom, ok := objMap["CropBottom"]; ok {
		if valCropBottom != nil {
			var valueForCropBottom float64
			err = json.Unmarshal(*valCropBottom, &valueForCropBottom)
			if err != nil {
				return err
			}
			this.CropBottom = valueForCropBottom
		}
	}

	if valCropLeft, ok := objMap["CropLeft"]; ok {
		if valCropLeft != nil {
			var valueForCropLeft float64
			err = json.Unmarshal(*valCropLeft, &valueForCropLeft)
			if err != nil {
				return err
			}
			this.CropLeft = valueForCropLeft
		}
	}

	if valCropRight, ok := objMap["CropRight"]; ok {
		if valCropRight != nil {
			var valueForCropRight float64
			err = json.Unmarshal(*valCropRight, &valueForCropRight)
			if err != nil {
				return err
			}
			this.CropRight = valueForCropRight
		}
	}

	if valCropTop, ok := objMap["CropTop"]; ok {
		if valCropTop != nil {
			var valueForCropTop float64
			err = json.Unmarshal(*valCropTop, &valueForCropTop)
			if err != nil {
				return err
			}
			this.CropTop = valueForCropTop
		}
	}

	if valDpi, ok := objMap["Dpi"]; ok {
		if valDpi != nil {
			var valueForDpi int32
			err = json.Unmarshal(*valDpi, &valueForDpi)
			if err != nil {
				return err
			}
			this.Dpi = valueForDpi
		}
	}

	if valImage, ok := objMap["Image"]; ok {
		if valImage != nil {
			var valueForImage ResourceUriElement
			err = json.Unmarshal(*valImage, &valueForImage)
			if err != nil {
				return err
			}
			this.Image = valueForImage
		}
	}

	if valBase64Data, ok := objMap["Base64Data"]; ok {
		if valBase64Data != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64Data, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}

	if valSvgData, ok := objMap["SvgData"]; ok {
		if valSvgData != nil {
			var valueForSvgData string
			err = json.Unmarshal(*valSvgData, &valueForSvgData)
			if err != nil {
				return err
			}
			this.SvgData = valueForSvgData
		}
	}

	if valPictureFillMode, ok := objMap["PictureFillMode"]; ok {
		if valPictureFillMode != nil {
			var valueForPictureFillMode PictureFillMode
			err = json.Unmarshal(*valPictureFillMode, &valueForPictureFillMode)
			if err != nil {
				return err
			}
			this.PictureFillMode = valueForPictureFillMode
		}
	}

    return nil
}
