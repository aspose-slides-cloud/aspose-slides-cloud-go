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

// Slide properties.
type ISlideProperties interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// First slide number.
	getFirstSlideNumber() int32
	setFirstSlideNumber(newValue int32)

	// Last used view mode.
	getOrientation() string
	setOrientation(newValue string)

	// Scale type.
	getScaleType() string
	setScaleType(newValue string)

	// Size type.
	getSizeType() string
	setSizeType(newValue string)

	// Width.
	getWidth() int32
	setWidth(newValue int32)

	// Height.
	getHeight() int32
	setHeight(newValue int32)
}

type SlideProperties struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// First slide number.
	FirstSlideNumber int32 `json:"FirstSlideNumber,omitempty"`

	// Last used view mode.
	Orientation string `json:"Orientation,omitempty"`

	// Scale type.
	ScaleType string `json:"ScaleType,omitempty"`

	// Size type.
	SizeType string `json:"SizeType,omitempty"`

	// Width.
	Width int32 `json:"Width,omitempty"`

	// Height.
	Height int32 `json:"Height,omitempty"`
}

func NewSlideProperties() *SlideProperties {
	instance := new(SlideProperties)
	instance.Orientation = ""
	instance.ScaleType = ""
	instance.SizeType = ""
	return instance
}

func (this *SlideProperties) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SlideProperties) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SlideProperties) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SlideProperties) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SlideProperties) getFirstSlideNumber() int32 {
	return this.FirstSlideNumber
}

func (this *SlideProperties) setFirstSlideNumber(newValue int32) {
	this.FirstSlideNumber = newValue
}
func (this *SlideProperties) getOrientation() string {
	return this.Orientation
}

func (this *SlideProperties) setOrientation(newValue string) {
	this.Orientation = newValue
}
func (this *SlideProperties) getScaleType() string {
	return this.ScaleType
}

func (this *SlideProperties) setScaleType(newValue string) {
	this.ScaleType = newValue
}
func (this *SlideProperties) getSizeType() string {
	return this.SizeType
}

func (this *SlideProperties) setSizeType(newValue string) {
	this.SizeType = newValue
}
func (this *SlideProperties) getWidth() int32 {
	return this.Width
}

func (this *SlideProperties) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *SlideProperties) getHeight() int32 {
	return this.Height
}

func (this *SlideProperties) setHeight(newValue int32) {
	this.Height = newValue
}

func (this *SlideProperties) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["selfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valFirstSlideNumber, ok := objMap["firstSlideNumber"]; ok {
		if valFirstSlideNumber != nil {
			var valueForFirstSlideNumber int32
			err = json.Unmarshal(*valFirstSlideNumber, &valueForFirstSlideNumber)
			if err != nil {
				return err
			}
			this.FirstSlideNumber = valueForFirstSlideNumber
		}
	}
	if valFirstSlideNumberCap, ok := objMap["FirstSlideNumber"]; ok {
		if valFirstSlideNumberCap != nil {
			var valueForFirstSlideNumber int32
			err = json.Unmarshal(*valFirstSlideNumberCap, &valueForFirstSlideNumber)
			if err != nil {
				return err
			}
			this.FirstSlideNumber = valueForFirstSlideNumber
		}
	}
	this.Orientation = ""
	if valOrientation, ok := objMap["orientation"]; ok {
		if valOrientation != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientation, &valueForOrientation)
			if err != nil {
				var valueForOrientationInt int32
				err = json.Unmarshal(*valOrientation, &valueForOrientationInt)
				if err != nil {
					return err
				}
				this.Orientation = string(valueForOrientationInt)
			} else {
				this.Orientation = valueForOrientation
			}
		}
	}
	if valOrientationCap, ok := objMap["Orientation"]; ok {
		if valOrientationCap != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientationCap, &valueForOrientation)
			if err != nil {
				var valueForOrientationInt int32
				err = json.Unmarshal(*valOrientationCap, &valueForOrientationInt)
				if err != nil {
					return err
				}
				this.Orientation = string(valueForOrientationInt)
			} else {
				this.Orientation = valueForOrientation
			}
		}
	}
	this.ScaleType = ""
	if valScaleType, ok := objMap["scaleType"]; ok {
		if valScaleType != nil {
			var valueForScaleType string
			err = json.Unmarshal(*valScaleType, &valueForScaleType)
			if err != nil {
				var valueForScaleTypeInt int32
				err = json.Unmarshal(*valScaleType, &valueForScaleTypeInt)
				if err != nil {
					return err
				}
				this.ScaleType = string(valueForScaleTypeInt)
			} else {
				this.ScaleType = valueForScaleType
			}
		}
	}
	if valScaleTypeCap, ok := objMap["ScaleType"]; ok {
		if valScaleTypeCap != nil {
			var valueForScaleType string
			err = json.Unmarshal(*valScaleTypeCap, &valueForScaleType)
			if err != nil {
				var valueForScaleTypeInt int32
				err = json.Unmarshal(*valScaleTypeCap, &valueForScaleTypeInt)
				if err != nil {
					return err
				}
				this.ScaleType = string(valueForScaleTypeInt)
			} else {
				this.ScaleType = valueForScaleType
			}
		}
	}
	this.SizeType = ""
	if valSizeType, ok := objMap["sizeType"]; ok {
		if valSizeType != nil {
			var valueForSizeType string
			err = json.Unmarshal(*valSizeType, &valueForSizeType)
			if err != nil {
				var valueForSizeTypeInt int32
				err = json.Unmarshal(*valSizeType, &valueForSizeTypeInt)
				if err != nil {
					return err
				}
				this.SizeType = string(valueForSizeTypeInt)
			} else {
				this.SizeType = valueForSizeType
			}
		}
	}
	if valSizeTypeCap, ok := objMap["SizeType"]; ok {
		if valSizeTypeCap != nil {
			var valueForSizeType string
			err = json.Unmarshal(*valSizeTypeCap, &valueForSizeType)
			if err != nil {
				var valueForSizeTypeInt int32
				err = json.Unmarshal(*valSizeTypeCap, &valueForSizeTypeInt)
				if err != nil {
					return err
				}
				this.SizeType = string(valueForSizeTypeInt)
			} else {
				this.SizeType = valueForSizeType
			}
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

    return nil
}
