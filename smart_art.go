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
	getNodes() []ISmartArtNode
	setNodes(newValue []ISmartArtNode)

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	getIsReversed() bool
	setIsReversed(newValue bool)
}

type SmartArt struct {

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
	Nodes []ISmartArtNode `json:"Nodes,omitempty"`

	// The state of the SmartArt diagram with regard to (left-to-right) LTR or (right-to-left) RTL, if the diagram supports reversal.
	IsReversed bool `json:"IsReversed"`
}

func (this *SmartArt) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SmartArt) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SmartArt) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SmartArt) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SmartArt) getName() string {
	return this.Name
}

func (this *SmartArt) setName(newValue string) {
	this.Name = newValue
}
func (this *SmartArt) getWidth() float64 {
	return this.Width
}

func (this *SmartArt) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *SmartArt) getHeight() float64 {
	return this.Height
}

func (this *SmartArt) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *SmartArt) getAlternativeText() string {
	return this.AlternativeText
}

func (this *SmartArt) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *SmartArt) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *SmartArt) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *SmartArt) getHidden() bool {
	return this.Hidden
}

func (this *SmartArt) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *SmartArt) getX() float64 {
	return this.X
}

func (this *SmartArt) setX(newValue float64) {
	this.X = newValue
}
func (this *SmartArt) getY() float64 {
	return this.Y
}

func (this *SmartArt) setY(newValue float64) {
	this.Y = newValue
}
func (this *SmartArt) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *SmartArt) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *SmartArt) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this *SmartArt) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this *SmartArt) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *SmartArt) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *SmartArt) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *SmartArt) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *SmartArt) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *SmartArt) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *SmartArt) getType() string {
	return this.Type_
}

func (this *SmartArt) setType(newValue string) {
	this.Type_ = newValue
}
func (this *SmartArt) getShapeType() string {
	return this.ShapeType
}

func (this *SmartArt) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *SmartArt) getLayout() string {
	return this.Layout
}

func (this *SmartArt) setLayout(newValue string) {
	this.Layout = newValue
}
func (this *SmartArt) getQuickStyle() string {
	return this.QuickStyle
}

func (this *SmartArt) setQuickStyle(newValue string) {
	this.QuickStyle = newValue
}
func (this *SmartArt) getColorStyle() string {
	return this.ColorStyle
}

func (this *SmartArt) setColorStyle(newValue string) {
	this.ColorStyle = newValue
}
func (this *SmartArt) getNodes() []ISmartArtNode {
	return this.Nodes
}

func (this *SmartArt) setNodes(newValue []ISmartArtNode) {
	this.Nodes = newValue
}
func (this *SmartArt) getIsReversed() bool {
	return this.IsReversed
}

func (this *SmartArt) setIsReversed(newValue bool) {
	this.IsReversed = newValue
}

func (this *SmartArt) UnmarshalJSON(b []byte) error {
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
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUriElement
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
	this.Type_ = "SmartArt"
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
	this.ShapeType = "Diagram"
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
	this.Layout = "AccentProcess"
	if valLayout, ok := objMap["layout"]; ok {
		if valLayout != nil {
			var valueForLayout string
			err = json.Unmarshal(*valLayout, &valueForLayout)
			if err != nil {
				var valueForLayoutInt int32
				err = json.Unmarshal(*valLayout, &valueForLayoutInt)
				if err != nil {
					return err
				}
				this.Layout = string(valueForLayoutInt)
			} else {
				this.Layout = valueForLayout
			}
		}
	}
	if valLayoutCap, ok := objMap["Layout"]; ok {
		if valLayoutCap != nil {
			var valueForLayout string
			err = json.Unmarshal(*valLayoutCap, &valueForLayout)
			if err != nil {
				var valueForLayoutInt int32
				err = json.Unmarshal(*valLayoutCap, &valueForLayoutInt)
				if err != nil {
					return err
				}
				this.Layout = string(valueForLayoutInt)
			} else {
				this.Layout = valueForLayout
			}
		}
	}
	this.QuickStyle = "SimpleFill"
	if valQuickStyle, ok := objMap["quickStyle"]; ok {
		if valQuickStyle != nil {
			var valueForQuickStyle string
			err = json.Unmarshal(*valQuickStyle, &valueForQuickStyle)
			if err != nil {
				var valueForQuickStyleInt int32
				err = json.Unmarshal(*valQuickStyle, &valueForQuickStyleInt)
				if err != nil {
					return err
				}
				this.QuickStyle = string(valueForQuickStyleInt)
			} else {
				this.QuickStyle = valueForQuickStyle
			}
		}
	}
	if valQuickStyleCap, ok := objMap["QuickStyle"]; ok {
		if valQuickStyleCap != nil {
			var valueForQuickStyle string
			err = json.Unmarshal(*valQuickStyleCap, &valueForQuickStyle)
			if err != nil {
				var valueForQuickStyleInt int32
				err = json.Unmarshal(*valQuickStyleCap, &valueForQuickStyleInt)
				if err != nil {
					return err
				}
				this.QuickStyle = string(valueForQuickStyleInt)
			} else {
				this.QuickStyle = valueForQuickStyle
			}
		}
	}
	this.ColorStyle = "Dark1Outline"
	if valColorStyle, ok := objMap["colorStyle"]; ok {
		if valColorStyle != nil {
			var valueForColorStyle string
			err = json.Unmarshal(*valColorStyle, &valueForColorStyle)
			if err != nil {
				var valueForColorStyleInt int32
				err = json.Unmarshal(*valColorStyle, &valueForColorStyleInt)
				if err != nil {
					return err
				}
				this.ColorStyle = string(valueForColorStyleInt)
			} else {
				this.ColorStyle = valueForColorStyle
			}
		}
	}
	if valColorStyleCap, ok := objMap["ColorStyle"]; ok {
		if valColorStyleCap != nil {
			var valueForColorStyle string
			err = json.Unmarshal(*valColorStyleCap, &valueForColorStyle)
			if err != nil {
				var valueForColorStyleInt int32
				err = json.Unmarshal(*valColorStyleCap, &valueForColorStyleInt)
				if err != nil {
					return err
				}
				this.ColorStyle = string(valueForColorStyleInt)
			} else {
				this.ColorStyle = valueForColorStyle
			}
		}
	}
	
	if valNodes, ok := objMap["nodes"]; ok {
		if valNodes != nil {
			var valueForNodes []SmartArtNode
			err = json.Unmarshal(*valNodes, &valueForNodes)
			if err != nil {
				return err
			}
			valueForINodes := make([]ISmartArtNode, len(valueForNodes))
			for i, v := range valueForNodes {
				valueForINodes[i] = ISmartArtNode(&v)
			}
			this.Nodes = valueForINodes
		}
	}
	if valNodesCap, ok := objMap["Nodes"]; ok {
		if valNodesCap != nil {
			var valueForNodes []SmartArtNode
			err = json.Unmarshal(*valNodesCap, &valueForNodes)
			if err != nil {
				return err
			}
			valueForINodes := make([]ISmartArtNode, len(valueForNodes))
			for i, v := range valueForNodes {
				valueForINodes[i] = ISmartArtNode(&v)
			}
			this.Nodes = valueForINodes
		}
	}
	
	if valIsReversed, ok := objMap["isReversed"]; ok {
		if valIsReversed != nil {
			var valueForIsReversed bool
			err = json.Unmarshal(*valIsReversed, &valueForIsReversed)
			if err != nil {
				return err
			}
			this.IsReversed = valueForIsReversed
		}
	}
	if valIsReversedCap, ok := objMap["IsReversed"]; ok {
		if valIsReversedCap != nil {
			var valueForIsReversed bool
			err = json.Unmarshal(*valIsReversedCap, &valueForIsReversed)
			if err != nil {
				return err
			}
			this.IsReversed = valueForIsReversed
		}
	}

    return nil
}
