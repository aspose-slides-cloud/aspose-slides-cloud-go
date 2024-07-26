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

// Summary zoom frame.
type ISummaryZoomFrame interface {

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

	// Zoom layout type
	GetZoomLayout() string
	SetZoomLayout(newValue string)

	// Zoom frame sections
	GetSections() []ISummaryZoomSection
	SetSections(newValue []ISummaryZoomSection)
}

type SummaryZoomFrame struct {

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

	// Zoom layout type
	ZoomLayout string `json:"ZoomLayout,omitempty"`

	// Zoom frame sections
	Sections []ISummaryZoomSection `json:"Sections,omitempty"`
}

func NewSummaryZoomFrame() *SummaryZoomFrame {
	instance := new(SummaryZoomFrame)
	instance.Type_ = "SummaryZoomFrame"
	return instance
}

func (this *SummaryZoomFrame) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SummaryZoomFrame) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SummaryZoomFrame) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SummaryZoomFrame) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SummaryZoomFrame) GetName() string {
	return this.Name
}

func (this *SummaryZoomFrame) SetName(newValue string) {
	this.Name = newValue
}
func (this *SummaryZoomFrame) GetWidth() float64 {
	return this.Width
}

func (this *SummaryZoomFrame) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *SummaryZoomFrame) GetHeight() float64 {
	return this.Height
}

func (this *SummaryZoomFrame) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *SummaryZoomFrame) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *SummaryZoomFrame) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *SummaryZoomFrame) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *SummaryZoomFrame) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *SummaryZoomFrame) GetHidden() *bool {
	return this.Hidden
}

func (this *SummaryZoomFrame) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *SummaryZoomFrame) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *SummaryZoomFrame) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *SummaryZoomFrame) GetX() float64 {
	return this.X
}

func (this *SummaryZoomFrame) SetX(newValue float64) {
	this.X = newValue
}
func (this *SummaryZoomFrame) GetY() float64 {
	return this.Y
}

func (this *SummaryZoomFrame) SetY(newValue float64) {
	this.Y = newValue
}
func (this *SummaryZoomFrame) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *SummaryZoomFrame) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *SummaryZoomFrame) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *SummaryZoomFrame) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *SummaryZoomFrame) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *SummaryZoomFrame) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *SummaryZoomFrame) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *SummaryZoomFrame) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *SummaryZoomFrame) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *SummaryZoomFrame) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *SummaryZoomFrame) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *SummaryZoomFrame) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *SummaryZoomFrame) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *SummaryZoomFrame) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *SummaryZoomFrame) GetType() string {
	return this.Type_
}

func (this *SummaryZoomFrame) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SummaryZoomFrame) GetZoomLayout() string {
	return this.ZoomLayout
}

func (this *SummaryZoomFrame) SetZoomLayout(newValue string) {
	this.ZoomLayout = newValue
}
func (this *SummaryZoomFrame) GetSections() []ISummaryZoomSection {
	return this.Sections
}

func (this *SummaryZoomFrame) SetSections(newValue []ISummaryZoomSection) {
	this.Sections = newValue
}

func (this *SummaryZoomFrame) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "SummaryZoomFrame"
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
	
	if valZoomLayout, ok := GetMapValue(objMap, "zoomLayout"); ok {
		if valZoomLayout != nil {
			var valueForZoomLayout string
			err = json.Unmarshal(*valZoomLayout, &valueForZoomLayout)
			if err != nil {
				var valueForZoomLayoutInt int32
				err = json.Unmarshal(*valZoomLayout, &valueForZoomLayoutInt)
				if err != nil {
					return err
				}
				this.ZoomLayout = string(valueForZoomLayoutInt)
			} else {
				this.ZoomLayout = valueForZoomLayout
			}
		}
	}
	
	if valSections, ok := GetMapValue(objMap, "sections"); ok {
		if valSections != nil {
			var valueForSections []json.RawMessage
			err = json.Unmarshal(*valSections, &valueForSections)
			if err != nil {
				return err
			}
			valueForISections := make([]ISummaryZoomSection, len(valueForSections))
			for i, v := range valueForSections {
				vObject, err := createObjectForType("SummaryZoomSection", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISections[i] = vObject.(ISummaryZoomSection)
				}
			}
			this.Sections = valueForISections
		}
	}

	return nil
}
