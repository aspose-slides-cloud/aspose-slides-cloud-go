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

// Zoom frame.
type IZoomFrame interface {

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

	// Image type of a zoom object. 
	getImageType() string
	setImageType(newValue string)

	// Navigation behavior in slideshow. 
	getReturnToParent() bool
	setReturnToParent(newValue bool)

	// Specifies whether the Zoom will use the background of the destination slide.
	getShowBackground() bool
	setShowBackground(newValue bool)

	// Internal image link for zoom object
	getImage() IResourceUri
	setImage(newValue IResourceUri)

	// Duration of the transition between Zoom and slide.
	getTransitionDuration() float64
	setTransitionDuration(newValue float64)

	// Links to the target slide
	getTargetSlideIndex() int32
	setTargetSlideIndex(newValue int32)
}

type ZoomFrame struct {

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

	// Image type of a zoom object. 
	ImageType string `json:"ImageType,omitempty"`

	// Navigation behavior in slideshow. 
	ReturnToParent bool `json:"ReturnToParent"`

	// Specifies whether the Zoom will use the background of the destination slide.
	ShowBackground bool `json:"ShowBackground"`

	// Internal image link for zoom object
	Image IResourceUri `json:"Image,omitempty"`

	// Duration of the transition between Zoom and slide.
	TransitionDuration float64 `json:"TransitionDuration,omitempty"`

	// Links to the target slide
	TargetSlideIndex int32 `json:"TargetSlideIndex,omitempty"`
}

func NewZoomFrame() *ZoomFrame {
	instance := new(ZoomFrame)
	instance.Type_ = "ZoomFrame"
	instance.ImageType = ""
	return instance
}

func (this *ZoomFrame) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ZoomFrame) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ZoomFrame) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ZoomFrame) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ZoomFrame) getName() string {
	return this.Name
}

func (this *ZoomFrame) setName(newValue string) {
	this.Name = newValue
}
func (this *ZoomFrame) getWidth() float64 {
	return this.Width
}

func (this *ZoomFrame) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *ZoomFrame) getHeight() float64 {
	return this.Height
}

func (this *ZoomFrame) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *ZoomFrame) getAlternativeText() string {
	return this.AlternativeText
}

func (this *ZoomFrame) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *ZoomFrame) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *ZoomFrame) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *ZoomFrame) getHidden() bool {
	return this.Hidden
}

func (this *ZoomFrame) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *ZoomFrame) getX() float64 {
	return this.X
}

func (this *ZoomFrame) setX(newValue float64) {
	this.X = newValue
}
func (this *ZoomFrame) getY() float64 {
	return this.Y
}

func (this *ZoomFrame) setY(newValue float64) {
	this.Y = newValue
}
func (this *ZoomFrame) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *ZoomFrame) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *ZoomFrame) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *ZoomFrame) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *ZoomFrame) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *ZoomFrame) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *ZoomFrame) getThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *ZoomFrame) setThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *ZoomFrame) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *ZoomFrame) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *ZoomFrame) getHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *ZoomFrame) setHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *ZoomFrame) getHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *ZoomFrame) setHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *ZoomFrame) getType() string {
	return this.Type_
}

func (this *ZoomFrame) setType(newValue string) {
	this.Type_ = newValue
}
func (this *ZoomFrame) getImageType() string {
	return this.ImageType
}

func (this *ZoomFrame) setImageType(newValue string) {
	this.ImageType = newValue
}
func (this *ZoomFrame) getReturnToParent() bool {
	return this.ReturnToParent
}

func (this *ZoomFrame) setReturnToParent(newValue bool) {
	this.ReturnToParent = newValue
}
func (this *ZoomFrame) getShowBackground() bool {
	return this.ShowBackground
}

func (this *ZoomFrame) setShowBackground(newValue bool) {
	this.ShowBackground = newValue
}
func (this *ZoomFrame) getImage() IResourceUri {
	return this.Image
}

func (this *ZoomFrame) setImage(newValue IResourceUri) {
	this.Image = newValue
}
func (this *ZoomFrame) getTransitionDuration() float64 {
	return this.TransitionDuration
}

func (this *ZoomFrame) setTransitionDuration(newValue float64) {
	this.TransitionDuration = newValue
}
func (this *ZoomFrame) getTargetSlideIndex() int32 {
	return this.TargetSlideIndex
}

func (this *ZoomFrame) setTargetSlideIndex(newValue int32) {
	this.TargetSlideIndex = newValue
}

func (this *ZoomFrame) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "ZoomFrame"
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
	this.ImageType = ""
	if valImageType, ok := objMap["imageType"]; ok {
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
	if valImageTypeCap, ok := objMap["ImageType"]; ok {
		if valImageTypeCap != nil {
			var valueForImageType string
			err = json.Unmarshal(*valImageTypeCap, &valueForImageType)
			if err != nil {
				var valueForImageTypeInt int32
				err = json.Unmarshal(*valImageTypeCap, &valueForImageTypeInt)
				if err != nil {
					return err
				}
				this.ImageType = string(valueForImageTypeInt)
			} else {
				this.ImageType = valueForImageType
			}
		}
	}
	
	if valReturnToParent, ok := objMap["returnToParent"]; ok {
		if valReturnToParent != nil {
			var valueForReturnToParent bool
			err = json.Unmarshal(*valReturnToParent, &valueForReturnToParent)
			if err != nil {
				return err
			}
			this.ReturnToParent = valueForReturnToParent
		}
	}
	if valReturnToParentCap, ok := objMap["ReturnToParent"]; ok {
		if valReturnToParentCap != nil {
			var valueForReturnToParent bool
			err = json.Unmarshal(*valReturnToParentCap, &valueForReturnToParent)
			if err != nil {
				return err
			}
			this.ReturnToParent = valueForReturnToParent
		}
	}
	
	if valShowBackground, ok := objMap["showBackground"]; ok {
		if valShowBackground != nil {
			var valueForShowBackground bool
			err = json.Unmarshal(*valShowBackground, &valueForShowBackground)
			if err != nil {
				return err
			}
			this.ShowBackground = valueForShowBackground
		}
	}
	if valShowBackgroundCap, ok := objMap["ShowBackground"]; ok {
		if valShowBackgroundCap != nil {
			var valueForShowBackground bool
			err = json.Unmarshal(*valShowBackgroundCap, &valueForShowBackground)
			if err != nil {
				return err
			}
			this.ShowBackground = valueForShowBackground
		}
	}
	
	if valImage, ok := objMap["image"]; ok {
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
	if valImageCap, ok := objMap["Image"]; ok {
		if valImageCap != nil {
			var valueForImage ResourceUri
			err = json.Unmarshal(*valImageCap, &valueForImage)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valImageCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valImageCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Image = vInterfaceObject
			}
		}
	}
	
	if valTransitionDuration, ok := objMap["transitionDuration"]; ok {
		if valTransitionDuration != nil {
			var valueForTransitionDuration float64
			err = json.Unmarshal(*valTransitionDuration, &valueForTransitionDuration)
			if err != nil {
				return err
			}
			this.TransitionDuration = valueForTransitionDuration
		}
	}
	if valTransitionDurationCap, ok := objMap["TransitionDuration"]; ok {
		if valTransitionDurationCap != nil {
			var valueForTransitionDuration float64
			err = json.Unmarshal(*valTransitionDurationCap, &valueForTransitionDuration)
			if err != nil {
				return err
			}
			this.TransitionDuration = valueForTransitionDuration
		}
	}
	
	if valTargetSlideIndex, ok := objMap["targetSlideIndex"]; ok {
		if valTargetSlideIndex != nil {
			var valueForTargetSlideIndex int32
			err = json.Unmarshal(*valTargetSlideIndex, &valueForTargetSlideIndex)
			if err != nil {
				return err
			}
			this.TargetSlideIndex = valueForTargetSlideIndex
		}
	}
	if valTargetSlideIndexCap, ok := objMap["TargetSlideIndex"]; ok {
		if valTargetSlideIndexCap != nil {
			var valueForTargetSlideIndex int32
			err = json.Unmarshal(*valTargetSlideIndexCap, &valueForTargetSlideIndex)
			if err != nil {
				return err
			}
			this.TargetSlideIndex = valueForTargetSlideIndex
		}
	}

	return nil
}
