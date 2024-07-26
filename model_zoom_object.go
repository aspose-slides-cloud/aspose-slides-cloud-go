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

// Zoom object.
type IZoomObject interface {

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

	GetType() string
	SetType(newValue string)

	// Image type of a zoom object. 
	GetImageType() string
	SetImageType(newValue string)

	// Navigation behavior in slideshow. 
	GetReturnToParent() *bool
	SetReturnToParent(newValue *bool)

	// Specifies whether the Zoom will use the background of the destination slide.
	GetShowBackground() *bool
	SetShowBackground(newValue *bool)

	// Internal image link for zoom object
	GetImage() IResourceUri
	SetImage(newValue IResourceUri)

	// Duration of the transition between Zoom and slide.
	GetTransitionDuration() float64
	SetTransitionDuration(newValue float64)
}

type ZoomObject struct {

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

	Type_ string `json:"Type,omitempty"`

	// Image type of a zoom object. 
	ImageType string `json:"ImageType,omitempty"`

	// Navigation behavior in slideshow. 
	ReturnToParent *bool `json:"ReturnToParent"`

	// Specifies whether the Zoom will use the background of the destination slide.
	ShowBackground *bool `json:"ShowBackground"`

	// Internal image link for zoom object
	Image IResourceUri `json:"Image,omitempty"`

	// Duration of the transition between Zoom and slide.
	TransitionDuration float64 `json:"TransitionDuration,omitempty"`
}

func NewZoomObject() *ZoomObject {
	instance := new(ZoomObject)
	return instance
}

func (this *ZoomObject) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ZoomObject) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ZoomObject) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ZoomObject) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ZoomObject) GetName() string {
	return this.Name
}

func (this *ZoomObject) SetName(newValue string) {
	this.Name = newValue
}
func (this *ZoomObject) GetWidth() float64 {
	return this.Width
}

func (this *ZoomObject) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *ZoomObject) GetHeight() float64 {
	return this.Height
}

func (this *ZoomObject) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *ZoomObject) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *ZoomObject) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *ZoomObject) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *ZoomObject) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *ZoomObject) GetHidden() *bool {
	return this.Hidden
}

func (this *ZoomObject) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *ZoomObject) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *ZoomObject) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *ZoomObject) GetX() float64 {
	return this.X
}

func (this *ZoomObject) SetX(newValue float64) {
	this.X = newValue
}
func (this *ZoomObject) GetY() float64 {
	return this.Y
}

func (this *ZoomObject) SetY(newValue float64) {
	this.Y = newValue
}
func (this *ZoomObject) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *ZoomObject) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *ZoomObject) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ZoomObject) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ZoomObject) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ZoomObject) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ZoomObject) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *ZoomObject) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *ZoomObject) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ZoomObject) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ZoomObject) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *ZoomObject) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *ZoomObject) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *ZoomObject) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *ZoomObject) GetType() string {
	return this.Type_
}

func (this *ZoomObject) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *ZoomObject) GetImageType() string {
	return this.ImageType
}

func (this *ZoomObject) SetImageType(newValue string) {
	this.ImageType = newValue
}
func (this *ZoomObject) GetReturnToParent() *bool {
	return this.ReturnToParent
}

func (this *ZoomObject) SetReturnToParent(newValue *bool) {
	this.ReturnToParent = newValue
}
func (this *ZoomObject) GetShowBackground() *bool {
	return this.ShowBackground
}

func (this *ZoomObject) SetShowBackground(newValue *bool) {
	this.ShowBackground = newValue
}
func (this *ZoomObject) GetImage() IResourceUri {
	return this.Image
}

func (this *ZoomObject) SetImage(newValue IResourceUri) {
	this.Image = newValue
}
func (this *ZoomObject) GetTransitionDuration() float64 {
	return this.TransitionDuration
}

func (this *ZoomObject) SetTransitionDuration(newValue float64) {
	this.TransitionDuration = newValue
}

func (this *ZoomObject) UnmarshalJSON(b []byte) error {
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
	
	if valImageType, ok := GetMapValue(objMap, "imageType"); ok {
		if valImageType != nil {
			var valueForImageType string
			err = json.Unmarshal(*valImageType, &valueForImageType)
			if err != nil {
				var valueForImageTypeInt int32
				err = json.Unmarshal(*valImageType, &valueForImageTypeInt)
				if err != nil {
					return err
				}
				this.ImageType = string(valueForImageTypeInt)
			} else {
				this.ImageType = valueForImageType
			}
		}
	}
	
	if valReturnToParent, ok := GetMapValue(objMap, "returnToParent"); ok {
		if valReturnToParent != nil {
			var valueForReturnToParent *bool
			err = json.Unmarshal(*valReturnToParent, &valueForReturnToParent)
			if err != nil {
				return err
			}
			this.ReturnToParent = valueForReturnToParent
		}
	}
	
	if valShowBackground, ok := GetMapValue(objMap, "showBackground"); ok {
		if valShowBackground != nil {
			var valueForShowBackground *bool
			err = json.Unmarshal(*valShowBackground, &valueForShowBackground)
			if err != nil {
				return err
			}
			this.ShowBackground = valueForShowBackground
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
	
	if valTransitionDuration, ok := GetMapValue(objMap, "transitionDuration"); ok {
		if valTransitionDuration != nil {
			var valueForTransitionDuration float64
			err = json.Unmarshal(*valTransitionDuration, &valueForTransitionDuration)
			if err != nil {
				return err
			}
			this.TransitionDuration = valueForTransitionDuration
		}
	}

	return nil
}
