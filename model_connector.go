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
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	GetName() string
	SetName(newValue string)

	// Gets or sets the width.
	GetWidth() float64
	SetWidth(newValue float64)

	// Gets or sets the height.
	GetHeight() float64
	SetHeight(newValue float64)

	// Gets or sets the alternative text.
	GetAlternativeText() string
	SetAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	GetAlternativeTextTitle() string
	SetAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	GetHidden() *bool
	SetHidden(newValue *bool)

	// Gets or sets 'Mark as decorative' option.
	GetIsDecorative() *bool
	SetIsDecorative(newValue *bool)

	// Gets or sets the X
	GetX() float64
	SetX(newValue float64)

	// Gets or sets the Y.
	GetY() float64
	SetY(newValue float64)

	// Gets z-order position of shape
	GetZOrderPosition() int32
	SetZOrderPosition(newValue int32)

	// Gets or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Shape type.
	GetType() string
	SetType(newValue string)

	// Combined shape type.
	GetShapeType() string
	SetShapeType(newValue string)

	// Start shape link.
	GetStartShapeConnectedTo() IResourceUri
	SetStartShapeConnectedTo(newValue IResourceUri)

	// Start shape index.
	GetStartShapeConnectedToIndex() int32
	SetStartShapeConnectedToIndex(newValue int32)

	// End shape link.
	GetEndShapeConnectedTo() IResourceUri
	SetEndShapeConnectedTo(newValue IResourceUri)

	// End shape index.
	GetEndShapeConnectedToIndex() int32
	SetEndShapeConnectedToIndex(newValue int32)
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
	Hidden *bool `json:"Hidden"`

	// Gets or sets 'Mark as decorative' option.
	IsDecorative *bool `json:"IsDecorative"`

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

func (this *Connector) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Connector) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Connector) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Connector) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Connector) GetName() string {
	return this.Name
}

func (this *Connector) SetName(newValue string) {
	this.Name = newValue
}
func (this *Connector) GetWidth() float64 {
	return this.Width
}

func (this *Connector) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *Connector) GetHeight() float64 {
	return this.Height
}

func (this *Connector) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *Connector) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *Connector) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *Connector) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *Connector) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *Connector) GetHidden() *bool {
	return this.Hidden
}

func (this *Connector) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *Connector) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *Connector) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *Connector) GetX() float64 {
	return this.X
}

func (this *Connector) SetX(newValue float64) {
	this.X = newValue
}
func (this *Connector) GetY() float64 {
	return this.Y
}

func (this *Connector) SetY(newValue float64) {
	this.Y = newValue
}
func (this *Connector) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *Connector) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *Connector) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Connector) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Connector) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Connector) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Connector) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *Connector) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *Connector) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Connector) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Connector) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *Connector) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *Connector) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *Connector) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *Connector) GetType() string {
	return this.Type_
}

func (this *Connector) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Connector) GetShapeType() string {
	return this.ShapeType
}

func (this *Connector) SetShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *Connector) GetStartShapeConnectedTo() IResourceUri {
	return this.StartShapeConnectedTo
}

func (this *Connector) SetStartShapeConnectedTo(newValue IResourceUri) {
	this.StartShapeConnectedTo = newValue
}
func (this *Connector) GetStartShapeConnectedToIndex() int32 {
	return this.StartShapeConnectedToIndex
}

func (this *Connector) SetStartShapeConnectedToIndex(newValue int32) {
	this.StartShapeConnectedToIndex = newValue
}
func (this *Connector) GetEndShapeConnectedTo() IResourceUri {
	return this.EndShapeConnectedTo
}

func (this *Connector) SetEndShapeConnectedTo(newValue IResourceUri) {
	this.EndShapeConnectedTo = newValue
}
func (this *Connector) GetEndShapeConnectedToIndex() int32 {
	return this.EndShapeConnectedToIndex
}

func (this *Connector) SetEndShapeConnectedToIndex(newValue int32) {
	this.EndShapeConnectedToIndex = newValue
}

func (this *Connector) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := GetMapValue(objMap, "selfUri"); ok {
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
	
	if valAlternateLinks, ok := GetMapValue(objMap, "alternateLinks"); ok {
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
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := GetMapValue(objMap, "alternativeText"); ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := GetMapValue(objMap, "alternativeTextTitle"); ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := GetMapValue(objMap, "hidden"); ok {
		if valHidden != nil {
			var valueForHidden *bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valIsDecorative, ok := GetMapValue(objMap, "isDecorative"); ok {
		if valIsDecorative != nil {
			var valueForIsDecorative *bool
			err = json.Unmarshal(*valIsDecorative, &valueForIsDecorative)
			if err != nil {
				return err
			}
			this.IsDecorative = valueForIsDecorative
		}
	}
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := GetMapValue(objMap, "zOrderPosition"); ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
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
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
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
	
	if valThreeDFormat, ok := GetMapValue(objMap, "threeDFormat"); ok {
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
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
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
	
	if valHyperlinkClick, ok := GetMapValue(objMap, "hyperlinkClick"); ok {
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
	
	if valHyperlinkMouseOver, ok := GetMapValue(objMap, "hyperlinkMouseOver"); ok {
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
	this.Type_ = "Connector"
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
	this.ShapeType = "Custom"
	if valShapeType, ok := GetMapValue(objMap, "shapeType"); ok {
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
	
	if valStartShapeConnectedTo, ok := GetMapValue(objMap, "startShapeConnectedTo"); ok {
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
	
	if valStartShapeConnectedToIndex, ok := GetMapValue(objMap, "startShapeConnectedToIndex"); ok {
		if valStartShapeConnectedToIndex != nil {
			var valueForStartShapeConnectedToIndex int32
			err = json.Unmarshal(*valStartShapeConnectedToIndex, &valueForStartShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.StartShapeConnectedToIndex = valueForStartShapeConnectedToIndex
		}
	}
	
	if valEndShapeConnectedTo, ok := GetMapValue(objMap, "endShapeConnectedTo"); ok {
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
	
	if valEndShapeConnectedToIndex, ok := GetMapValue(objMap, "endShapeConnectedToIndex"); ok {
		if valEndShapeConnectedToIndex != nil {
			var valueForEndShapeConnectedToIndex int32
			err = json.Unmarshal(*valEndShapeConnectedToIndex, &valueForEndShapeConnectedToIndex)
			if err != nil {
				return err
			}
			this.EndShapeConnectedToIndex = valueForEndShapeConnectedToIndex
		}
	}

	return nil
}
