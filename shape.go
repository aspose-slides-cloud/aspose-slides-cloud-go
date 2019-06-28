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

// Represents AutoShape resource.
type IShape interface {

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

	// Gets or sets the text.
	getText() string
	setText(newValue string)

	// Get or sets list to paragraphs list
	getParagraphs() IResourceUriElement
	setParagraphs(newValue IResourceUriElement)
}

type Shape struct {

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

	Type_ string `json:"Type"`

	ShapeType string `json:"ShapeType"`

	GeometryShapeType string `json:"GeometryShapeType"`

	// Gets or sets the text.
	Text string `json:"Text,omitempty"`

	// Get or sets list to paragraphs list
	Paragraphs IResourceUriElement `json:"Paragraphs,omitempty"`
}

func (this Shape) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Shape) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Shape) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Shape) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Shape) getName() string {
	return this.Name
}

func (this Shape) setName(newValue string) {
	this.Name = newValue
}
func (this Shape) getWidth() float64 {
	return this.Width
}

func (this Shape) setWidth(newValue float64) {
	this.Width = newValue
}
func (this Shape) getHeight() float64 {
	return this.Height
}

func (this Shape) setHeight(newValue float64) {
	this.Height = newValue
}
func (this Shape) getAlternativeText() string {
	return this.AlternativeText
}

func (this Shape) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this Shape) getHidden() bool {
	return this.Hidden
}

func (this Shape) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this Shape) getX() float64 {
	return this.X
}

func (this Shape) setX(newValue float64) {
	this.X = newValue
}
func (this Shape) getY() float64 {
	return this.Y
}

func (this Shape) setY(newValue float64) {
	this.Y = newValue
}
func (this Shape) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this Shape) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this Shape) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this Shape) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this Shape) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this Shape) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this Shape) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this Shape) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this Shape) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this Shape) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this Shape) getType() string {
	return this.Type_
}

func (this Shape) setType(newValue string) {
	this.Type_ = newValue
}
func (this Shape) getShapeType() string {
	return this.ShapeType
}

func (this Shape) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this Shape) getGeometryShapeType() string {
	return this.GeometryShapeType
}

func (this Shape) setGeometryShapeType(newValue string) {
	this.GeometryShapeType = newValue
}
func (this Shape) getText() string {
	return this.Text
}

func (this Shape) setText(newValue string) {
	this.Text = newValue
}
func (this Shape) getParagraphs() IResourceUriElement {
	return this.Paragraphs
}

func (this Shape) setParagraphs(newValue IResourceUriElement) {
	this.Paragraphs = newValue
}

func (this *Shape) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "Shape"
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
	this.ShapeType = "Custom"
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
	this.GeometryShapeType = "Custom"
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
	
	if valText, ok := objMap["Text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	
	if valParagraphs, ok := objMap["Paragraphs"]; ok {
		if valParagraphs != nil {
			var valueForParagraphs ResourceUriElement
			err = json.Unmarshal(*valParagraphs, &valueForParagraphs)
			if err != nil {
				return err
			}
			this.Paragraphs = valueForParagraphs
		}
	}

    return nil
}
