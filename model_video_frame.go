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

// Represents VideoFrame resource.
type IVideoFrame interface {

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

	// Determines whether a video is shown in full screen mode.
	GetFullScreenMode() *bool
	SetFullScreenMode(newValue *bool)

	// Determines whether a VideoFrame is hidden. 
	GetHideAtShowing() *bool
	SetHideAtShowing(newValue *bool)

	// Determines whether a video is looped.
	GetPlayLoopMode() *bool
	SetPlayLoopMode(newValue *bool)

	// Returns or sets the video play mode.  
	GetPlayMode() string
	SetPlayMode(newValue string)

	// Determines whether a video is automatically rewinded to start as soon as the movie has finished playing
	GetRewindVideo() *bool
	SetRewindVideo(newValue *bool)

	// Returns or sets the audio volume.
	GetVolume() string
	SetVolume(newValue string)

	// Video data encoded in base64.
	GetBase64Data() string
	SetBase64Data(newValue string)

	// Picture fill format.
	GetPictureFillFormat() IPictureFill
	SetPictureFillFormat(newValue IPictureFill)

	// Trim start [ms]
	GetTrimFromStart() float64
	SetTrimFromStart(newValue float64)

	// Trim end [ms]
	GetTrimFromEnd() float64
	SetTrimFromEnd(newValue float64)
}

type VideoFrame struct {

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
	ZOrderPosition int32 `json:"ZOrderPosition,omitempty"`

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

	// Determines whether a video is shown in full screen mode.
	FullScreenMode *bool `json:"FullScreenMode"`

	// Determines whether a VideoFrame is hidden. 
	HideAtShowing *bool `json:"HideAtShowing"`

	// Determines whether a video is looped.
	PlayLoopMode *bool `json:"PlayLoopMode"`

	// Returns or sets the video play mode.  
	PlayMode string `json:"PlayMode,omitempty"`

	// Determines whether a video is automatically rewinded to start as soon as the movie has finished playing
	RewindVideo *bool `json:"RewindVideo"`

	// Returns or sets the audio volume.
	Volume string `json:"Volume,omitempty"`

	// Video data encoded in base64.
	Base64Data string `json:"Base64Data,omitempty"`

	// Picture fill format.
	PictureFillFormat IPictureFill `json:"PictureFillFormat,omitempty"`

	// Trim start [ms]
	TrimFromStart float64 `json:"TrimFromStart,omitempty"`

	// Trim end [ms]
	TrimFromEnd float64 `json:"TrimFromEnd,omitempty"`
}

func NewVideoFrame() *VideoFrame {
	instance := new(VideoFrame)
	instance.Type_ = "VideoFrame"
	instance.ShapeType = "Custom"
	return instance
}

func (this *VideoFrame) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *VideoFrame) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *VideoFrame) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *VideoFrame) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *VideoFrame) GetName() string {
	return this.Name
}

func (this *VideoFrame) SetName(newValue string) {
	this.Name = newValue
}
func (this *VideoFrame) GetWidth() float64 {
	return this.Width
}

func (this *VideoFrame) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *VideoFrame) GetHeight() float64 {
	return this.Height
}

func (this *VideoFrame) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *VideoFrame) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *VideoFrame) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *VideoFrame) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *VideoFrame) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *VideoFrame) GetHidden() *bool {
	return this.Hidden
}

func (this *VideoFrame) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *VideoFrame) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *VideoFrame) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *VideoFrame) GetX() float64 {
	return this.X
}

func (this *VideoFrame) SetX(newValue float64) {
	this.X = newValue
}
func (this *VideoFrame) GetY() float64 {
	return this.Y
}

func (this *VideoFrame) SetY(newValue float64) {
	this.Y = newValue
}
func (this *VideoFrame) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *VideoFrame) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *VideoFrame) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *VideoFrame) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *VideoFrame) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *VideoFrame) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *VideoFrame) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *VideoFrame) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *VideoFrame) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *VideoFrame) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *VideoFrame) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *VideoFrame) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *VideoFrame) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *VideoFrame) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *VideoFrame) GetType() string {
	return this.Type_
}

func (this *VideoFrame) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *VideoFrame) GetShapeType() string {
	return this.ShapeType
}

func (this *VideoFrame) SetShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *VideoFrame) GetFullScreenMode() *bool {
	return this.FullScreenMode
}

func (this *VideoFrame) SetFullScreenMode(newValue *bool) {
	this.FullScreenMode = newValue
}
func (this *VideoFrame) GetHideAtShowing() *bool {
	return this.HideAtShowing
}

func (this *VideoFrame) SetHideAtShowing(newValue *bool) {
	this.HideAtShowing = newValue
}
func (this *VideoFrame) GetPlayLoopMode() *bool {
	return this.PlayLoopMode
}

func (this *VideoFrame) SetPlayLoopMode(newValue *bool) {
	this.PlayLoopMode = newValue
}
func (this *VideoFrame) GetPlayMode() string {
	return this.PlayMode
}

func (this *VideoFrame) SetPlayMode(newValue string) {
	this.PlayMode = newValue
}
func (this *VideoFrame) GetRewindVideo() *bool {
	return this.RewindVideo
}

func (this *VideoFrame) SetRewindVideo(newValue *bool) {
	this.RewindVideo = newValue
}
func (this *VideoFrame) GetVolume() string {
	return this.Volume
}

func (this *VideoFrame) SetVolume(newValue string) {
	this.Volume = newValue
}
func (this *VideoFrame) GetBase64Data() string {
	return this.Base64Data
}

func (this *VideoFrame) SetBase64Data(newValue string) {
	this.Base64Data = newValue
}
func (this *VideoFrame) GetPictureFillFormat() IPictureFill {
	return this.PictureFillFormat
}

func (this *VideoFrame) SetPictureFillFormat(newValue IPictureFill) {
	this.PictureFillFormat = newValue
}
func (this *VideoFrame) GetTrimFromStart() float64 {
	return this.TrimFromStart
}

func (this *VideoFrame) SetTrimFromStart(newValue float64) {
	this.TrimFromStart = newValue
}
func (this *VideoFrame) GetTrimFromEnd() float64 {
	return this.TrimFromEnd
}

func (this *VideoFrame) SetTrimFromEnd(newValue float64) {
	this.TrimFromEnd = newValue
}

func (this *VideoFrame) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "VideoFrame"
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
	
	if valFullScreenMode, ok := GetMapValue(objMap, "fullScreenMode"); ok {
		if valFullScreenMode != nil {
			var valueForFullScreenMode *bool
			err = json.Unmarshal(*valFullScreenMode, &valueForFullScreenMode)
			if err != nil {
				return err
			}
			this.FullScreenMode = valueForFullScreenMode
		}
	}
	
	if valHideAtShowing, ok := GetMapValue(objMap, "hideAtShowing"); ok {
		if valHideAtShowing != nil {
			var valueForHideAtShowing *bool
			err = json.Unmarshal(*valHideAtShowing, &valueForHideAtShowing)
			if err != nil {
				return err
			}
			this.HideAtShowing = valueForHideAtShowing
		}
	}
	
	if valPlayLoopMode, ok := GetMapValue(objMap, "playLoopMode"); ok {
		if valPlayLoopMode != nil {
			var valueForPlayLoopMode *bool
			err = json.Unmarshal(*valPlayLoopMode, &valueForPlayLoopMode)
			if err != nil {
				return err
			}
			this.PlayLoopMode = valueForPlayLoopMode
		}
	}
	
	if valPlayMode, ok := GetMapValue(objMap, "playMode"); ok {
		if valPlayMode != nil {
			var valueForPlayMode string
			err = json.Unmarshal(*valPlayMode, &valueForPlayMode)
			if err != nil {
				var valueForPlayModeInt int32
				err = json.Unmarshal(*valPlayMode, &valueForPlayModeInt)
				if err != nil {
					return err
				}
				this.PlayMode = string(valueForPlayModeInt)
			} else {
				this.PlayMode = valueForPlayMode
			}
		}
	}
	
	if valRewindVideo, ok := GetMapValue(objMap, "rewindVideo"); ok {
		if valRewindVideo != nil {
			var valueForRewindVideo *bool
			err = json.Unmarshal(*valRewindVideo, &valueForRewindVideo)
			if err != nil {
				return err
			}
			this.RewindVideo = valueForRewindVideo
		}
	}
	
	if valVolume, ok := GetMapValue(objMap, "volume"); ok {
		if valVolume != nil {
			var valueForVolume string
			err = json.Unmarshal(*valVolume, &valueForVolume)
			if err != nil {
				var valueForVolumeInt int32
				err = json.Unmarshal(*valVolume, &valueForVolumeInt)
				if err != nil {
					return err
				}
				this.Volume = string(valueForVolumeInt)
			} else {
				this.Volume = valueForVolume
			}
		}
	}
	
	if valBase64Data, ok := GetMapValue(objMap, "base64Data"); ok {
		if valBase64Data != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64Data, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	
	if valPictureFillFormat, ok := GetMapValue(objMap, "pictureFillFormat"); ok {
		if valPictureFillFormat != nil {
			var valueForPictureFillFormat PictureFill
			err = json.Unmarshal(*valPictureFillFormat, &valueForPictureFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PictureFill", *valPictureFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPictureFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPictureFill)
			if ok {
				this.PictureFillFormat = vInterfaceObject
			}
		}
	}
	
	if valTrimFromStart, ok := GetMapValue(objMap, "trimFromStart"); ok {
		if valTrimFromStart != nil {
			var valueForTrimFromStart float64
			err = json.Unmarshal(*valTrimFromStart, &valueForTrimFromStart)
			if err != nil {
				return err
			}
			this.TrimFromStart = valueForTrimFromStart
		}
	}
	
	if valTrimFromEnd, ok := GetMapValue(objMap, "trimFromEnd"); ok {
		if valTrimFromEnd != nil {
			var valueForTrimFromEnd float64
			err = json.Unmarshal(*valTrimFromEnd, &valueForTrimFromEnd)
			if err != nil {
				return err
			}
			this.TrimFromEnd = valueForTrimFromEnd
		}
	}

	return nil
}
