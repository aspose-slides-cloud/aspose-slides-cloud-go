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

// Represents Connector resource.
type IConnector interface {

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

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	getThreeDFormat() IThreeDFormat
	setThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	getHyperlinkClick() IHyperlink
	setHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	getHyperlinkMouseOver() IHyperlink
	setHyperlinkMouseOver(newValue IHyperlink)

	// Shape type.
	getType() string
	setType(newValue string)

	// Combined shape type.
	getShapeType() string
	setShapeType(newValue string)

	// Start shape link.
	getStartShapeConnectedTo() IResourceUri
	setStartShapeConnectedTo(newValue IResourceUri)

	// Start shape index.
	getStartShapeConnectedToIndex() int32
	setStartShapeConnectedToIndex(newValue int32)

	// End shape link.
	getEndShapeConnectedTo() IResourceUri
	setEndShapeConnectedTo(newValue IResourceUri)

	// End shape index.
	getEndShapeConnectedToIndex() int32
	setEndShapeConnectedToIndex(newValue int32)
}

type Connector struct {

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

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Combined shape type.
	ShapeType string `json:"ShapeType"`

	// Start shape link.
	StartShapeConnectedTo IResourceUri `json:"StartShapeConnectedTo,omitempty"`

	// Start shape index.
	StartShapeConnectedToIndex int32 `json:"StartShapeConnectedToIndex,omitempty"`

	// End shape link.
	EndShapeConnectedTo IResourceUri `json:"EndShapeConnectedTo,omitempty"`

	// End shape index.
	EndShapeConnectedToIndex int32 `json:"EndShapeConnectedToIndex,omitempty"`
}

func NewConnector() *Connector {
	instance := new(Connector)
	instance.Type_ = "Connector"
	instance.ShapeType = "Custom"
	return instance
}

func (this *Connector) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Connector) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Connector) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Connector) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Connector) getName() string {
	return this.Name
}

func (this *Connector) setName(newValue string) {
	this.Name = newValue
}
func (this *Connector) getWidth() float64 {
	return this.Width
}

func (this *Connector) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *Connector) getHeight() float64 {
	return this.Height
}

func (this *Connector) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *Connector) getAlternativeText() string {
	return this.AlternativeText
}

func (this *Connector) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *Connector) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *Connector) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *Connector) getHidden() bool {
	return this.Hidden
}

func (this *Connector) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *Connector) getX() float64 {
	return this.X
}

func (this *Connector) setX(newValue float64) {
	this.X = newValue
}
func (this *Connector) getY() float64 {
	return this.Y
}

func (this *Connector) setY(newValue float64) {
	this.Y = newValue
}
func (this *Connector) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *Connector) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *Connector) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Connector) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Connector) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Connector) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Connector) getThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *Connector) setThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *Connector) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Connector) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Connector) getHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *Connector) setHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *Connector) getHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *Connector) setHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *Connector) getType() string {
	return this.Type_
}

func (this *Connector) setType(newValue string) {
	this.Type_ = newValue
}
func (this *Connector) getShapeType() string {
	return this.ShapeType
}

func (this *Connector) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *Connector) getStartShapeConnectedTo() IResourceUri {
	return this.StartShapeConnectedTo
}

func (this *Connector) setStartShapeConnectedTo(newValue IResourceUri) {
	this.StartShapeConnectedTo = newValue
}
func (this *Connector) getStartShapeConnectedToIndex() int32 {
	return this.StartShapeConnectedToIndex
}

func (this *Connector) setStartShapeConnectedToIndex(newValue int32) {
	this.StartShapeConnectedToIndex = newValue
}
func (this *Connector) getEndShapeConnectedTo() IResourceUri {
	return this.EndShapeConnectedTo
}

func (this *Connector) setEndShapeConnectedTo(newValue IResourceUri) {
	this.EndShapeConnectedTo = newValue
}
func (this *Connector) getEndShapeConnectedToIndex() int32 {
	return this.EndShapeConnectedToIndex
}

func (this *Connector) setEndShapeConnectedToIndex(newValue int32) {
	this.EndShapeConnectedToIndex = newValue
}

func (this *Connector) UnmarshalJSON(b []byte) error {
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
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
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
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkClick, ok := objMap["hyperlinkClick"]; ok {
		if valHyperlinkClick != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClick, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClick)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClick, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	if valHyperlinkClickCap, ok := objMap["HyperlinkClick"]; ok {
		if valHyperlinkClickCap != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClickCap, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClickCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClickCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkMouseOver, ok := objMap["hyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOver != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOver, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOver)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOver, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	if valHyperlinkMouseOverCap, ok := objMap["HyperlinkMouseOver"]; ok {
		if valHyperlinkMouseOverCap != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOverCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOverCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	this.Type_ = "Connector"
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
	
	if valStartShapeConnectedTo, ok := objMap["startShapeConnectedTo"]; ok {
		if valStartShapeConnectedTo != nil {
			var valueForStartShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valStartShapeConnectedTo, &valueForStartShapeConnectedTo)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valStartShapeConnectedTo)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valStartShapeConnectedTo, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.StartShapeConnectedTo = vInterfaceObject
			}
		}
	}
	if valStartShapeConnectedToCap, ok := objMap["StartShapeConnectedTo"]; ok {
		if valStartShapeConnectedToCap != nil {
			var valueForStartShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valStartShapeConnectedToCap, &valueForStartShapeConnectedTo)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valStartShapeConnectedToCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valStartShapeConnectedToCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.StartShapeConnectedTo = vInterfaceObject
			}
		}
	}
	
	if valStartShapeConnectedToIndex, ok := objMap["startShapeConnectedToIndex"]; ok {
		if valStartShapeConnectedToIndex != nil {
			var valueForStartShapeConnectedToIndex int32
			err = json.Unmarshal(*valStartShapeConnectedToIndex, &valueForStartShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.StartShapeConnectedToIndex = valueForStartShapeConnectedToIndex
		}
	}
	if valStartShapeConnectedToIndexCap, ok := objMap["StartShapeConnectedToIndex"]; ok {
		if valStartShapeConnectedToIndexCap != nil {
			var valueForStartShapeConnectedToIndex int32
			err = json.Unmarshal(*valStartShapeConnectedToIndexCap, &valueForStartShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.StartShapeConnectedToIndex = valueForStartShapeConnectedToIndex
		}
	}
	
	if valEndShapeConnectedTo, ok := objMap["endShapeConnectedTo"]; ok {
		if valEndShapeConnectedTo != nil {
			var valueForEndShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valEndShapeConnectedTo, &valueForEndShapeConnectedTo)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valEndShapeConnectedTo)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEndShapeConnectedTo, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.EndShapeConnectedTo = vInterfaceObject
			}
		}
	}
	if valEndShapeConnectedToCap, ok := objMap["EndShapeConnectedTo"]; ok {
		if valEndShapeConnectedToCap != nil {
			var valueForEndShapeConnectedTo ResourceUri
			err = json.Unmarshal(*valEndShapeConnectedToCap, &valueForEndShapeConnectedTo)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valEndShapeConnectedToCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEndShapeConnectedToCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.EndShapeConnectedTo = vInterfaceObject
			}
		}
	}
	
	if valEndShapeConnectedToIndex, ok := objMap["endShapeConnectedToIndex"]; ok {
		if valEndShapeConnectedToIndex != nil {
			var valueForEndShapeConnectedToIndex int32
			err = json.Unmarshal(*valEndShapeConnectedToIndex, &valueForEndShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.EndShapeConnectedToIndex = valueForEndShapeConnectedToIndex
		}
	}
	if valEndShapeConnectedToIndexCap, ok := objMap["EndShapeConnectedToIndex"]; ok {
		if valEndShapeConnectedToIndexCap != nil {
			var valueForEndShapeConnectedToIndex int32
			err = json.Unmarshal(*valEndShapeConnectedToIndexCap, &valueForEndShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.EndShapeConnectedToIndex = valueForEndShapeConnectedToIndex
		}
	}

	return nil
}
