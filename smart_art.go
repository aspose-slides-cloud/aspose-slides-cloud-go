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

// Represents SmartArt shape resource.
type ISmartArt interface {

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

	// Shape type.
	getType() string
	setType(newValue string)

	// Combined shape type.
	getShapeType() string
	setShapeType(newValue string)

	// Layout type.
	getLayout() string
	setLayout(newValue string)

	// Quick style.
	getQuickStyle() string
	setQuickStyle(newValue string)

	// Color style.
	getColorStyle() string
	setColorStyle(newValue string)

	// Collection of nodes in SmartArt object.             
	getNodes() []SmartArtNode
	setNodes(newValue []SmartArtNode)

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	getIsReversed() bool
	setIsReversed(newValue bool)
}

type SmartArt struct {

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

	// Shape type.
	Type_ string `json:"Type"`

	// Combined shape type.
	ShapeType string `json:"ShapeType"`

	// Layout type.
	Layout string `json:"Layout"`

	// Quick style.
	QuickStyle string `json:"QuickStyle"`

	// Color style.
	ColorStyle string `json:"ColorStyle"`

	// Collection of nodes in SmartArt object.             
	Nodes []SmartArtNode `json:"Nodes,omitempty"`

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	IsReversed bool `json:"IsReversed"`
}

func (this SmartArt) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this SmartArt) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this SmartArt) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this SmartArt) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this SmartArt) getName() string {
	return this.Name
}

func (this SmartArt) setName(newValue string) {
	this.Name = newValue
}
func (this SmartArt) getWidth() float64 {
	return this.Width
}

func (this SmartArt) setWidth(newValue float64) {
	this.Width = newValue
}
func (this SmartArt) getHeight() float64 {
	return this.Height
}

func (this SmartArt) setHeight(newValue float64) {
	this.Height = newValue
}
func (this SmartArt) getAlternativeText() string {
	return this.AlternativeText
}

func (this SmartArt) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this SmartArt) getHidden() bool {
	return this.Hidden
}

func (this SmartArt) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this SmartArt) getX() float64 {
	return this.X
}

func (this SmartArt) setX(newValue float64) {
	this.X = newValue
}
func (this SmartArt) getY() float64 {
	return this.Y
}

func (this SmartArt) setY(newValue float64) {
	this.Y = newValue
}
func (this SmartArt) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this SmartArt) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this SmartArt) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this SmartArt) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this SmartArt) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this SmartArt) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this SmartArt) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this SmartArt) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this SmartArt) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this SmartArt) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this SmartArt) getType() string {
	return this.Type_
}

func (this SmartArt) setType(newValue string) {
	this.Type_ = newValue
}
func (this SmartArt) getShapeType() string {
	return this.ShapeType
}

func (this SmartArt) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this SmartArt) getLayout() string {
	return this.Layout
}

func (this SmartArt) setLayout(newValue string) {
	this.Layout = newValue
}
func (this SmartArt) getQuickStyle() string {
	return this.QuickStyle
}

func (this SmartArt) setQuickStyle(newValue string) {
	this.QuickStyle = newValue
}
func (this SmartArt) getColorStyle() string {
	return this.ColorStyle
}

func (this SmartArt) setColorStyle(newValue string) {
	this.ColorStyle = newValue
}
func (this SmartArt) getNodes() []SmartArtNode {
	return this.Nodes
}

func (this SmartArt) setNodes(newValue []SmartArtNode) {
	this.Nodes = newValue
}
func (this SmartArt) getIsReversed() bool {
	return this.IsReversed
}

func (this SmartArt) setIsReversed(newValue bool) {
	this.IsReversed = newValue
}

func (this *SmartArt) UnmarshalJSON(b []byte) error {
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
	this.Layout = "AccentProcess"
	if valLayout, ok := objMap["Layout"]; ok {
		if valLayout != nil {
			var valueForLayout string
			err = json.Unmarshal(*valLayout, &valueForLayout)
			if err != nil {
				return err
			}
			this.Layout = valueForLayout
		}
	}
	this.QuickStyle = "SimpleFill"
	if valQuickStyle, ok := objMap["QuickStyle"]; ok {
		if valQuickStyle != nil {
			var valueForQuickStyle string
			err = json.Unmarshal(*valQuickStyle, &valueForQuickStyle)
			if err != nil {
				return err
			}
			this.QuickStyle = valueForQuickStyle
		}
	}
	this.ColorStyle = "Dark1Outline"
	if valColorStyle, ok := objMap["ColorStyle"]; ok {
		if valColorStyle != nil {
			var valueForColorStyle string
			err = json.Unmarshal(*valColorStyle, &valueForColorStyle)
			if err != nil {
				return err
			}
			this.ColorStyle = valueForColorStyle
		}
	}
	
	if valNodes, ok := objMap["Nodes"]; ok {
		if valNodes != nil {
			var valueForNodes []SmartArtNode
			err = json.Unmarshal(*valNodes, &valueForNodes)
			if err != nil {
				return err
			}
			this.Nodes = valueForNodes
		}
	}
	
	if valIsReversed, ok := objMap["IsReversed"]; ok {
		if valIsReversed != nil {
			var valueForIsReversed bool
			err = json.Unmarshal(*valIsReversed, &valueForIsReversed)
			if err != nil {
				return err
			}
			this.IsReversed = valueForIsReversed
		}
	}

    return nil
}
