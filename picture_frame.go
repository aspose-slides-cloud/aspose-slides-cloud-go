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

// Represents PictureFrame resource.
type IPictureFrame interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	getName() string
	setName(newValue string)

	// Gets or sets the width.
	getWidth() float64
	setWidth(newValue float64)

	// Gets or sets the height.
	getHeight() float64
	setHeight(newValue float64)

	// Gets or sets the alternative text.
	getAlternativeText() string
	setAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	getAlternativeTextTitle() string
	setAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	getHidden() bool
	setHidden(newValue bool)

	// Gets or sets the X
	getX() float64
	setX(newValue float64)

	// Gets or sets the Y.
	getY() float64
	setY(newValue float64)

	// Gets z-order position of shape
	getZOrderPosition() int32
	setZOrderPosition(newValue int32)

	// Gets or sets the link to shapes.
	getShapes() IResourceUri
	setShapes(newValue IResourceUri)

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Shape type.
	getType() string
	setType(newValue string)

	// Combined shape type.
	getShapeType() string
	setShapeType(newValue string)

	// Fill format.
	getPictureFillFormat() IPictureFill
	setPictureFillFormat(newValue IPictureFill)
}

type PictureFrame struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// The title of alternative text associated with the shape.
	AlternativeTextTitle string `json:"AlternativeTextTitle,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden bool `json:"Hidden"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUri `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Combined shape type.
	ShapeType string `json:"ShapeType"`

	// Fill format.
	PictureFillFormat IPictureFill `json:"PictureFillFormat,omitempty"`
}

func NewPictureFrame() *PictureFrame {
	instance := new(PictureFrame)
	instance.Type_ = "PictureFrame"
	instance.ShapeType = "Custom"
	return instance
}

func (this *PictureFrame) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *PictureFrame) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *PictureFrame) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *PictureFrame) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *PictureFrame) getName() string {
	return this.Name
}

func (this *PictureFrame) setName(newValue string) {
	this.Name = newValue
}
func (this *PictureFrame) getWidth() float64 {
	return this.Width
}

func (this *PictureFrame) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *PictureFrame) getHeight() float64 {
	return this.Height
}

func (this *PictureFrame) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *PictureFrame) getAlternativeText() string {
	return this.AlternativeText
}

func (this *PictureFrame) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *PictureFrame) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *PictureFrame) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *PictureFrame) getHidden() bool {
	return this.Hidden
}

func (this *PictureFrame) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *PictureFrame) getX() float64 {
	return this.X
}

func (this *PictureFrame) setX(newValue float64) {
	this.X = newValue
}
func (this *PictureFrame) getY() float64 {
	return this.Y
}

func (this *PictureFrame) setY(newValue float64) {
	this.Y = newValue
}
func (this *PictureFrame) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *PictureFrame) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *PictureFrame) getShapes() IResourceUri {
	return this.Shapes
}

func (this *PictureFrame) setShapes(newValue IResourceUri) {
	this.Shapes = newValue
}
func (this *PictureFrame) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *PictureFrame) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *PictureFrame) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *PictureFrame) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *PictureFrame) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *PictureFrame) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *PictureFrame) getType() string {
	return this.Type_
}

func (this *PictureFrame) setType(newValue string) {
	this.Type_ = newValue
}
func (this *PictureFrame) getShapeType() string {
	return this.ShapeType
}

func (this *PictureFrame) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *PictureFrame) getPictureFillFormat() IPictureFill {
	return this.PictureFillFormat
}

func (this *PictureFrame) setPictureFillFormat(newValue IPictureFill) {
	this.PictureFillFormat = newValue
}

func (this *PictureFrame) UnmarshalJSON(b []byte) error {
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
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := objMap["alternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	if valAlternativeTextCap, ok := objMap["AlternativeText"]; ok {
		if valAlternativeTextCap != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeTextCap, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := objMap["alternativeTextTitle"]; ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	if valAlternativeTextTitleCap, ok := objMap["AlternativeTextTitle"]; ok {
		if valAlternativeTextTitleCap != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitleCap, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := objMap["hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	if valHiddenCap, ok := objMap["Hidden"]; ok {
		if valHiddenCap != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHiddenCap, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valX, ok := objMap["x"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	if valXCap, ok := objMap["X"]; ok {
		if valXCap != nil {
			var valueForX float64
			err = json.Unmarshal(*valXCap, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := objMap["y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	if valYCap, ok := objMap["Y"]; ok {
		if valYCap != nil {
			var valueForY float64
			err = json.Unmarshal(*valYCap, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := objMap["zOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	if valZOrderPositionCap, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPositionCap != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPositionCap, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valShapes, ok := objMap["shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapesCap, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
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
	this.Type_ = "PictureFrame"
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
	this.ShapeType = "Custom"
	if valShapeType, ok := objMap["shapeType"]; ok {
		if valShapeType != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeType, &valueForShapeType)
			if err != nil {
				var valueForShapeTypeInt int32
				err = json.Unmarshal(*valShapeType, &valueForShapeTypeInt)
				if err != nil {
					return err
				}
				this.ShapeType = string(valueForShapeTypeInt)
			} else {
				this.ShapeType = valueForShapeType
			}
		}
	}
	if valShapeTypeCap, ok := objMap["ShapeType"]; ok {
		if valShapeTypeCap != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeTypeCap, &valueForShapeType)
			if err != nil {
				var valueForShapeTypeInt int32
				err = json.Unmarshal(*valShapeTypeCap, &valueForShapeTypeInt)
				if err != nil {
					return err
				}
				this.ShapeType = string(valueForShapeTypeInt)
			} else {
				this.ShapeType = valueForShapeType
			}
		}
	}
	
	if valPictureFillFormat, ok := objMap["pictureFillFormat"]; ok {
		if valPictureFillFormat != nil {
			var valueForPictureFillFormat PictureFill
			err = json.Unmarshal(*valPictureFillFormat, &valueForPictureFillFormat)
			if err != nil {
				return err
			}
			this.PictureFillFormat = &valueForPictureFillFormat
		}
	}
	if valPictureFillFormatCap, ok := objMap["PictureFillFormat"]; ok {
		if valPictureFillFormatCap != nil {
			var valueForPictureFillFormat PictureFill
			err = json.Unmarshal(*valPictureFillFormatCap, &valueForPictureFillFormat)
			if err != nil {
				return err
			}
			this.PictureFillFormat = &valueForPictureFillFormat
		}
	}

    return nil
}
