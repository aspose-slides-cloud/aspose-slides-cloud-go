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

// Represents OleObjectFrame resource.
type IOleObjectFrame interface {

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

	// True if an object is visible as icon.
	GetIsObjectIcon() *bool
	SetIsObjectIcon(newValue *bool)

	// The title for OleObject icon.             
	GetSubstitutePictureTitle() string
	SetSubstitutePictureTitle(newValue string)

	// OleObject image fill properties.
	GetSubstitutePictureFormat() IPictureFill
	SetSubstitutePictureFormat(newValue IPictureFill)

	// Returns or sets the name of an object.
	GetObjectName() string
	SetObjectName(newValue string)

	// File data of embedded OLE object. 
	GetEmbeddedFileBase64Data() string
	SetEmbeddedFileBase64Data(newValue string)

	// File extension for the current embedded OLE object
	GetEmbeddedFileExtension() string
	SetEmbeddedFileExtension(newValue string)

	// ProgID of an object.
	GetObjectProgId() string
	SetObjectProgId(newValue string)

	// Full path to a linked file.
	GetLinkPath() string
	SetLinkPath(newValue string)

	// Determines if the linked embedded object is automatically updated when the presentation is opened or printed. Read/write Boolean.
	GetUpdateAutomatic() *bool
	SetUpdateAutomatic(newValue *bool)
}

type OleObjectFrame struct {

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

	// True if an object is visible as icon.
	IsObjectIcon *bool `json:"IsObjectIcon"`

	// The title for OleObject icon.             
	SubstitutePictureTitle string `json:"SubstitutePictureTitle,omitempty"`

	// OleObject image fill properties.
	SubstitutePictureFormat IPictureFill `json:"SubstitutePictureFormat,omitempty"`

	// Returns or sets the name of an object.
	ObjectName string `json:"ObjectName,omitempty"`

	// File data of embedded OLE object. 
	EmbeddedFileBase64Data string `json:"EmbeddedFileBase64Data,omitempty"`

	// File extension for the current embedded OLE object
	EmbeddedFileExtension string `json:"EmbeddedFileExtension,omitempty"`

	// ProgID of an object.
	ObjectProgId string `json:"ObjectProgId,omitempty"`

	// Full path to a linked file.
	LinkPath string `json:"LinkPath,omitempty"`

	// Determines if the linked embedded object is automatically updated when the presentation is opened or printed. Read/write Boolean.
	UpdateAutomatic *bool `json:"UpdateAutomatic"`
}

func NewOleObjectFrame() *OleObjectFrame {
	instance := new(OleObjectFrame)
	instance.Type_ = "OleObjectFrame"
	return instance
}

func (this *OleObjectFrame) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *OleObjectFrame) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *OleObjectFrame) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *OleObjectFrame) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *OleObjectFrame) GetName() string {
	return this.Name
}

func (this *OleObjectFrame) SetName(newValue string) {
	this.Name = newValue
}
func (this *OleObjectFrame) GetWidth() float64 {
	return this.Width
}

func (this *OleObjectFrame) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *OleObjectFrame) GetHeight() float64 {
	return this.Height
}

func (this *OleObjectFrame) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *OleObjectFrame) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *OleObjectFrame) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *OleObjectFrame) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *OleObjectFrame) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *OleObjectFrame) GetHidden() *bool {
	return this.Hidden
}

func (this *OleObjectFrame) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *OleObjectFrame) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *OleObjectFrame) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *OleObjectFrame) GetX() float64 {
	return this.X
}

func (this *OleObjectFrame) SetX(newValue float64) {
	this.X = newValue
}
func (this *OleObjectFrame) GetY() float64 {
	return this.Y
}

func (this *OleObjectFrame) SetY(newValue float64) {
	this.Y = newValue
}
func (this *OleObjectFrame) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *OleObjectFrame) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *OleObjectFrame) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *OleObjectFrame) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *OleObjectFrame) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *OleObjectFrame) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *OleObjectFrame) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *OleObjectFrame) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *OleObjectFrame) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *OleObjectFrame) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *OleObjectFrame) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *OleObjectFrame) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *OleObjectFrame) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *OleObjectFrame) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *OleObjectFrame) GetType() string {
	return this.Type_
}

func (this *OleObjectFrame) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *OleObjectFrame) GetIsObjectIcon() *bool {
	return this.IsObjectIcon
}

func (this *OleObjectFrame) SetIsObjectIcon(newValue *bool) {
	this.IsObjectIcon = newValue
}
func (this *OleObjectFrame) GetSubstitutePictureTitle() string {
	return this.SubstitutePictureTitle
}

func (this *OleObjectFrame) SetSubstitutePictureTitle(newValue string) {
	this.SubstitutePictureTitle = newValue
}
func (this *OleObjectFrame) GetSubstitutePictureFormat() IPictureFill {
	return this.SubstitutePictureFormat
}

func (this *OleObjectFrame) SetSubstitutePictureFormat(newValue IPictureFill) {
	this.SubstitutePictureFormat = newValue
}
func (this *OleObjectFrame) GetObjectName() string {
	return this.ObjectName
}

func (this *OleObjectFrame) SetObjectName(newValue string) {
	this.ObjectName = newValue
}
func (this *OleObjectFrame) GetEmbeddedFileBase64Data() string {
	return this.EmbeddedFileBase64Data
}

func (this *OleObjectFrame) SetEmbeddedFileBase64Data(newValue string) {
	this.EmbeddedFileBase64Data = newValue
}
func (this *OleObjectFrame) GetEmbeddedFileExtension() string {
	return this.EmbeddedFileExtension
}

func (this *OleObjectFrame) SetEmbeddedFileExtension(newValue string) {
	this.EmbeddedFileExtension = newValue
}
func (this *OleObjectFrame) GetObjectProgId() string {
	return this.ObjectProgId
}

func (this *OleObjectFrame) SetObjectProgId(newValue string) {
	this.ObjectProgId = newValue
}
func (this *OleObjectFrame) GetLinkPath() string {
	return this.LinkPath
}

func (this *OleObjectFrame) SetLinkPath(newValue string) {
	this.LinkPath = newValue
}
func (this *OleObjectFrame) GetUpdateAutomatic() *bool {
	return this.UpdateAutomatic
}

func (this *OleObjectFrame) SetUpdateAutomatic(newValue *bool) {
	this.UpdateAutomatic = newValue
}

func (this *OleObjectFrame) UnmarshalJSON(b []byte) error {
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
			var valueForHidden *bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	if valHiddenCap, ok := objMap["Hidden"]; ok {
		if valHiddenCap != nil {
			var valueForHidden *bool
			err = json.Unmarshal(*valHiddenCap, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valIsDecorative, ok := objMap["isDecorative"]; ok {
		if valIsDecorative != nil {
			var valueForIsDecorative *bool
			err = json.Unmarshal(*valIsDecorative, &valueForIsDecorative)
			if err != nil {
				return err
			}
			this.IsDecorative = valueForIsDecorative
		}
	}
	if valIsDecorativeCap, ok := objMap["IsDecorative"]; ok {
		if valIsDecorativeCap != nil {
			var valueForIsDecorative *bool
			err = json.Unmarshal(*valIsDecorativeCap, &valueForIsDecorative)
			if err != nil {
				return err
			}
			this.IsDecorative = valueForIsDecorative
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
	this.Type_ = "OleObjectFrame"
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
	
	if valIsObjectIcon, ok := objMap["isObjectIcon"]; ok {
		if valIsObjectIcon != nil {
			var valueForIsObjectIcon *bool
			err = json.Unmarshal(*valIsObjectIcon, &valueForIsObjectIcon)
			if err != nil {
				return err
			}
			this.IsObjectIcon = valueForIsObjectIcon
		}
	}
	if valIsObjectIconCap, ok := objMap["IsObjectIcon"]; ok {
		if valIsObjectIconCap != nil {
			var valueForIsObjectIcon *bool
			err = json.Unmarshal(*valIsObjectIconCap, &valueForIsObjectIcon)
			if err != nil {
				return err
			}
			this.IsObjectIcon = valueForIsObjectIcon
		}
	}
	
	if valSubstitutePictureTitle, ok := objMap["substitutePictureTitle"]; ok {
		if valSubstitutePictureTitle != nil {
			var valueForSubstitutePictureTitle string
			err = json.Unmarshal(*valSubstitutePictureTitle, &valueForSubstitutePictureTitle)
			if err != nil {
				return err
			}
			this.SubstitutePictureTitle = valueForSubstitutePictureTitle
		}
	}
	if valSubstitutePictureTitleCap, ok := objMap["SubstitutePictureTitle"]; ok {
		if valSubstitutePictureTitleCap != nil {
			var valueForSubstitutePictureTitle string
			err = json.Unmarshal(*valSubstitutePictureTitleCap, &valueForSubstitutePictureTitle)
			if err != nil {
				return err
			}
			this.SubstitutePictureTitle = valueForSubstitutePictureTitle
		}
	}
	
	if valSubstitutePictureFormat, ok := objMap["substitutePictureFormat"]; ok {
		if valSubstitutePictureFormat != nil {
			var valueForSubstitutePictureFormat PictureFill
			err = json.Unmarshal(*valSubstitutePictureFormat, &valueForSubstitutePictureFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PictureFill", *valSubstitutePictureFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSubstitutePictureFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPictureFill)
			if ok {
				this.SubstitutePictureFormat = vInterfaceObject
			}
		}
	}
	if valSubstitutePictureFormatCap, ok := objMap["SubstitutePictureFormat"]; ok {
		if valSubstitutePictureFormatCap != nil {
			var valueForSubstitutePictureFormat PictureFill
			err = json.Unmarshal(*valSubstitutePictureFormatCap, &valueForSubstitutePictureFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PictureFill", *valSubstitutePictureFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSubstitutePictureFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPictureFill)
			if ok {
				this.SubstitutePictureFormat = vInterfaceObject
			}
		}
	}
	
	if valObjectName, ok := objMap["objectName"]; ok {
		if valObjectName != nil {
			var valueForObjectName string
			err = json.Unmarshal(*valObjectName, &valueForObjectName)
			if err != nil {
				return err
			}
			this.ObjectName = valueForObjectName
		}
	}
	if valObjectNameCap, ok := objMap["ObjectName"]; ok {
		if valObjectNameCap != nil {
			var valueForObjectName string
			err = json.Unmarshal(*valObjectNameCap, &valueForObjectName)
			if err != nil {
				return err
			}
			this.ObjectName = valueForObjectName
		}
	}
	
	if valEmbeddedFileBase64Data, ok := objMap["embeddedFileBase64Data"]; ok {
		if valEmbeddedFileBase64Data != nil {
			var valueForEmbeddedFileBase64Data string
			err = json.Unmarshal(*valEmbeddedFileBase64Data, &valueForEmbeddedFileBase64Data)
			if err != nil {
				return err
			}
			this.EmbeddedFileBase64Data = valueForEmbeddedFileBase64Data
		}
	}
	if valEmbeddedFileBase64DataCap, ok := objMap["EmbeddedFileBase64Data"]; ok {
		if valEmbeddedFileBase64DataCap != nil {
			var valueForEmbeddedFileBase64Data string
			err = json.Unmarshal(*valEmbeddedFileBase64DataCap, &valueForEmbeddedFileBase64Data)
			if err != nil {
				return err
			}
			this.EmbeddedFileBase64Data = valueForEmbeddedFileBase64Data
		}
	}
	
	if valEmbeddedFileExtension, ok := objMap["embeddedFileExtension"]; ok {
		if valEmbeddedFileExtension != nil {
			var valueForEmbeddedFileExtension string
			err = json.Unmarshal(*valEmbeddedFileExtension, &valueForEmbeddedFileExtension)
			if err != nil {
				return err
			}
			this.EmbeddedFileExtension = valueForEmbeddedFileExtension
		}
	}
	if valEmbeddedFileExtensionCap, ok := objMap["EmbeddedFileExtension"]; ok {
		if valEmbeddedFileExtensionCap != nil {
			var valueForEmbeddedFileExtension string
			err = json.Unmarshal(*valEmbeddedFileExtensionCap, &valueForEmbeddedFileExtension)
			if err != nil {
				return err
			}
			this.EmbeddedFileExtension = valueForEmbeddedFileExtension
		}
	}
	
	if valObjectProgId, ok := objMap["objectProgId"]; ok {
		if valObjectProgId != nil {
			var valueForObjectProgId string
			err = json.Unmarshal(*valObjectProgId, &valueForObjectProgId)
			if err != nil {
				return err
			}
			this.ObjectProgId = valueForObjectProgId
		}
	}
	if valObjectProgIdCap, ok := objMap["ObjectProgId"]; ok {
		if valObjectProgIdCap != nil {
			var valueForObjectProgId string
			err = json.Unmarshal(*valObjectProgIdCap, &valueForObjectProgId)
			if err != nil {
				return err
			}
			this.ObjectProgId = valueForObjectProgId
		}
	}
	
	if valLinkPath, ok := objMap["linkPath"]; ok {
		if valLinkPath != nil {
			var valueForLinkPath string
			err = json.Unmarshal(*valLinkPath, &valueForLinkPath)
			if err != nil {
				return err
			}
			this.LinkPath = valueForLinkPath
		}
	}
	if valLinkPathCap, ok := objMap["LinkPath"]; ok {
		if valLinkPathCap != nil {
			var valueForLinkPath string
			err = json.Unmarshal(*valLinkPathCap, &valueForLinkPath)
			if err != nil {
				return err
			}
			this.LinkPath = valueForLinkPath
		}
	}
	
	if valUpdateAutomatic, ok := objMap["updateAutomatic"]; ok {
		if valUpdateAutomatic != nil {
			var valueForUpdateAutomatic *bool
			err = json.Unmarshal(*valUpdateAutomatic, &valueForUpdateAutomatic)
			if err != nil {
				return err
			}
			this.UpdateAutomatic = valueForUpdateAutomatic
		}
	}
	if valUpdateAutomaticCap, ok := objMap["UpdateAutomatic"]; ok {
		if valUpdateAutomaticCap != nil {
			var valueForUpdateAutomatic *bool
			err = json.Unmarshal(*valUpdateAutomaticCap, &valueForUpdateAutomatic)
			if err != nil {
				return err
			}
			this.UpdateAutomatic = valueForUpdateAutomatic
		}
	}

	return nil
}
