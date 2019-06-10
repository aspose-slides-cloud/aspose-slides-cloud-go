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

// Represents GeometryShape resource.
type IGeometryShape interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

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
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	getType() string
	setType(newValue string)

	getShapeType() string
	setShapeType(newValue string)

	getGeometryShapeType() string
	setGeometryShapeType(newValue string)
}

type GeometryShape struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden bool `json:"Hidden,omitempty"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	Type_ string `json:"Type,omitempty"`

	ShapeType string `json:"ShapeType"`

	GeometryShapeType string `json:"GeometryShapeType"`
}

func (this GeometryShape) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this GeometryShape) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this GeometryShape) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this GeometryShape) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this GeometryShape) getName() string {
	return this.Name
}

func (this GeometryShape) setName(newValue string) {
	this.Name = newValue
}
func (this GeometryShape) getWidth() float64 {
	return this.Width
}

func (this GeometryShape) setWidth(newValue float64) {
	this.Width = newValue
}
func (this GeometryShape) getHeight() float64 {
	return this.Height
}

func (this GeometryShape) setHeight(newValue float64) {
	this.Height = newValue
}
func (this GeometryShape) getAlternativeText() string {
	return this.AlternativeText
}

func (this GeometryShape) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this GeometryShape) getHidden() bool {
	return this.Hidden
}

func (this GeometryShape) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this GeometryShape) getX() float64 {
	return this.X
}

func (this GeometryShape) setX(newValue float64) {
	this.X = newValue
}
func (this GeometryShape) getY() float64 {
	return this.Y
}

func (this GeometryShape) setY(newValue float64) {
	this.Y = newValue
}
func (this GeometryShape) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this GeometryShape) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this GeometryShape) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this GeometryShape) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this GeometryShape) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this GeometryShape) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this GeometryShape) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this GeometryShape) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this GeometryShape) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this GeometryShape) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this GeometryShape) getType() string {
	return this.Type_
}

func (this GeometryShape) setType(newValue string) {
	this.Type_ = newValue
}
func (this GeometryShape) getShapeType() string {
	return this.ShapeType
}

func (this GeometryShape) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this GeometryShape) getGeometryShapeType() string {
	return this.GeometryShapeType
}

func (this GeometryShape) setGeometryShapeType(newValue string) {
	this.GeometryShapeType = newValue
}

func (this *GeometryShape) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valSelfUri, ok := objMap["SelfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}

	if valAlternateLinks, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}

	if valName, ok := objMap["Name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

	if valHeight, ok := objMap["Height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	if valAlternativeText, ok := objMap["AlternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}

	if valHidden, ok := objMap["Hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}

	if valX, ok := objMap["X"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}

	if valY, ok := objMap["Y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}

	if valZOrderPosition, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}

	if valShapes, ok := objMap["Shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}

	if valFillFormat, ok := objMap["FillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = valueForFillFormat
		}
	}

	if valEffectFormat, ok := objMap["EffectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = valueForEffectFormat
		}
	}

	if valLineFormat, ok := objMap["LineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = valueForLineFormat
		}
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valShapeType, ok := objMap["ShapeType"]; ok {
		if valShapeType != nil {
			var valueForShapeType string
			err = json.Unmarshal(*valShapeType, &valueForShapeType)
			if err != nil {
				return err
			}
			this.ShapeType = valueForShapeType
		}
	}

	if valGeometryShapeType, ok := objMap["GeometryShapeType"]; ok {
		if valGeometryShapeType != nil {
			var valueForGeometryShapeType string
			err = json.Unmarshal(*valGeometryShapeType, &valueForGeometryShapeType)
			if err != nil {
				return err
			}
			this.GeometryShapeType = valueForGeometryShapeType
		}
	}

    return nil
}
