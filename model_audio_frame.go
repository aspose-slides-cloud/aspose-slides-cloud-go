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

// Represents AudioFrame resource.
type IAudioFrame interface {

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

	// Returns or sets a last track index.
	GetAudioCdEndTrack() int32
	SetAudioCdEndTrack(newValue int32)

	// Returns or sets a last track time.
	GetAudioCdEndTrackTime() int32
	SetAudioCdEndTrackTime(newValue int32)

	// Returns or sets a start track index.
	GetAudioCdStartTrack() int32
	SetAudioCdStartTrack(newValue int32)

	// Returns or sets a start track time. 
	GetAudioCdStartTrackTime() int32
	SetAudioCdStartTrackTime(newValue int32)

	// Determines whether a sound is embedded to a presentation.
	GetEmbedded() *bool
	SetEmbedded(newValue *bool)

	// Determines whether an AudioFrame is hidden.
	GetHideAtShowing() *bool
	SetHideAtShowing(newValue *bool)

	// Determines whether an audio is looped. 
	GetPlayLoopMode() *bool
	SetPlayLoopMode(newValue *bool)

	// Returns or sets the audio play mode.
	GetPlayMode() string
	SetPlayMode(newValue string)

	// Returns or sets the audio volume.
	GetVolume() string
	SetVolume(newValue string)

	// Audio data encoded in base64.
	GetBase64Data() string
	SetBase64Data(newValue string)

	// Determines whether an audio is playing across the slides.
	GetPlayAcrossSlides() *bool
	SetPlayAcrossSlides(newValue *bool)

	// Determines whether audio is automatically rewound to start after playing.
	GetRewindAudio() *bool
	SetRewindAudio(newValue *bool)

	// Picture fill format.
	GetPictureFillFormat() IPictureFill
	SetPictureFillFormat(newValue IPictureFill)
}

type AudioFrame struct {

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

	// Returns or sets a last track index.
	AudioCdEndTrack int32 `json:"AudioCdEndTrack,omitempty"`

	// Returns or sets a last track time.
	AudioCdEndTrackTime int32 `json:"AudioCdEndTrackTime,omitempty"`

	// Returns or sets a start track index.
	AudioCdStartTrack int32 `json:"AudioCdStartTrack,omitempty"`

	// Returns or sets a start track time. 
	AudioCdStartTrackTime int32 `json:"AudioCdStartTrackTime,omitempty"`

	// Determines whether a sound is embedded to a presentation.
	Embedded *bool `json:"Embedded"`

	// Determines whether an AudioFrame is hidden.
	HideAtShowing *bool `json:"HideAtShowing"`

	// Determines whether an audio is looped. 
	PlayLoopMode *bool `json:"PlayLoopMode"`

	// Returns or sets the audio play mode.
	PlayMode string `json:"PlayMode,omitempty"`

	// Returns or sets the audio volume.
	Volume string `json:"Volume,omitempty"`

	// Audio data encoded in base64.
	Base64Data string `json:"Base64Data,omitempty"`

	// Determines whether an audio is playing across the slides.
	PlayAcrossSlides *bool `json:"PlayAcrossSlides"`

	// Determines whether audio is automatically rewound to start after playing.
	RewindAudio *bool `json:"RewindAudio"`

	// Picture fill format.
	PictureFillFormat IPictureFill `json:"PictureFillFormat,omitempty"`
}

func NewAudioFrame() *AudioFrame {
	instance := new(AudioFrame)
	instance.Type_ = "AudioFrame"
	instance.ShapeType = "Custom"
	return instance
}

func (this *AudioFrame) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *AudioFrame) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *AudioFrame) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *AudioFrame) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *AudioFrame) GetName() string {
	return this.Name
}

func (this *AudioFrame) SetName(newValue string) {
	this.Name = newValue
}
func (this *AudioFrame) GetWidth() float64 {
	return this.Width
}

func (this *AudioFrame) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *AudioFrame) GetHeight() float64 {
	return this.Height
}

func (this *AudioFrame) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *AudioFrame) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *AudioFrame) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *AudioFrame) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *AudioFrame) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *AudioFrame) GetHidden() *bool {
	return this.Hidden
}

func (this *AudioFrame) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *AudioFrame) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *AudioFrame) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *AudioFrame) GetX() float64 {
	return this.X
}

func (this *AudioFrame) SetX(newValue float64) {
	this.X = newValue
}
func (this *AudioFrame) GetY() float64 {
	return this.Y
}

func (this *AudioFrame) SetY(newValue float64) {
	this.Y = newValue
}
func (this *AudioFrame) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *AudioFrame) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *AudioFrame) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *AudioFrame) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *AudioFrame) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *AudioFrame) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *AudioFrame) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *AudioFrame) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *AudioFrame) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *AudioFrame) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *AudioFrame) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *AudioFrame) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *AudioFrame) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *AudioFrame) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *AudioFrame) GetType() string {
	return this.Type_
}

func (this *AudioFrame) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *AudioFrame) GetShapeType() string {
	return this.ShapeType
}

func (this *AudioFrame) SetShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *AudioFrame) GetAudioCdEndTrack() int32 {
	return this.AudioCdEndTrack
}

func (this *AudioFrame) SetAudioCdEndTrack(newValue int32) {
	this.AudioCdEndTrack = newValue
}
func (this *AudioFrame) GetAudioCdEndTrackTime() int32 {
	return this.AudioCdEndTrackTime
}

func (this *AudioFrame) SetAudioCdEndTrackTime(newValue int32) {
	this.AudioCdEndTrackTime = newValue
}
func (this *AudioFrame) GetAudioCdStartTrack() int32 {
	return this.AudioCdStartTrack
}

func (this *AudioFrame) SetAudioCdStartTrack(newValue int32) {
	this.AudioCdStartTrack = newValue
}
func (this *AudioFrame) GetAudioCdStartTrackTime() int32 {
	return this.AudioCdStartTrackTime
}

func (this *AudioFrame) SetAudioCdStartTrackTime(newValue int32) {
	this.AudioCdStartTrackTime = newValue
}
func (this *AudioFrame) GetEmbedded() *bool {
	return this.Embedded
}

func (this *AudioFrame) SetEmbedded(newValue *bool) {
	this.Embedded = newValue
}
func (this *AudioFrame) GetHideAtShowing() *bool {
	return this.HideAtShowing
}

func (this *AudioFrame) SetHideAtShowing(newValue *bool) {
	this.HideAtShowing = newValue
}
func (this *AudioFrame) GetPlayLoopMode() *bool {
	return this.PlayLoopMode
}

func (this *AudioFrame) SetPlayLoopMode(newValue *bool) {
	this.PlayLoopMode = newValue
}
func (this *AudioFrame) GetPlayMode() string {
	return this.PlayMode
}

func (this *AudioFrame) SetPlayMode(newValue string) {
	this.PlayMode = newValue
}
func (this *AudioFrame) GetVolume() string {
	return this.Volume
}

func (this *AudioFrame) SetVolume(newValue string) {
	this.Volume = newValue
}
func (this *AudioFrame) GetBase64Data() string {
	return this.Base64Data
}

func (this *AudioFrame) SetBase64Data(newValue string) {
	this.Base64Data = newValue
}
func (this *AudioFrame) GetPlayAcrossSlides() *bool {
	return this.PlayAcrossSlides
}

func (this *AudioFrame) SetPlayAcrossSlides(newValue *bool) {
	this.PlayAcrossSlides = newValue
}
func (this *AudioFrame) GetRewindAudio() *bool {
	return this.RewindAudio
}

func (this *AudioFrame) SetRewindAudio(newValue *bool) {
	this.RewindAudio = newValue
}
func (this *AudioFrame) GetPictureFillFormat() IPictureFill {
	return this.PictureFillFormat
}

func (this *AudioFrame) SetPictureFillFormat(newValue IPictureFill) {
	this.PictureFillFormat = newValue
}

func (this *AudioFrame) UnmarshalJSON(b []byte) error {
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
	this.Type_ = "AudioFrame"
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
	
	if valAudioCdEndTrack, ok := objMap["audioCdEndTrack"]; ok {
		if valAudioCdEndTrack != nil {
			var valueForAudioCdEndTrack int32
			err = json.Unmarshal(*valAudioCdEndTrack, &valueForAudioCdEndTrack)
			if err != nil {
				return err
			}
			this.AudioCdEndTrack = valueForAudioCdEndTrack
		}
	}
	if valAudioCdEndTrackCap, ok := objMap["AudioCdEndTrack"]; ok {
		if valAudioCdEndTrackCap != nil {
			var valueForAudioCdEndTrack int32
			err = json.Unmarshal(*valAudioCdEndTrackCap, &valueForAudioCdEndTrack)
			if err != nil {
				return err
			}
			this.AudioCdEndTrack = valueForAudioCdEndTrack
		}
	}
	
	if valAudioCdEndTrackTime, ok := objMap["audioCdEndTrackTime"]; ok {
		if valAudioCdEndTrackTime != nil {
			var valueForAudioCdEndTrackTime int32
			err = json.Unmarshal(*valAudioCdEndTrackTime, &valueForAudioCdEndTrackTime)
			if err != nil {
				return err
			}
			this.AudioCdEndTrackTime = valueForAudioCdEndTrackTime
		}
	}
	if valAudioCdEndTrackTimeCap, ok := objMap["AudioCdEndTrackTime"]; ok {
		if valAudioCdEndTrackTimeCap != nil {
			var valueForAudioCdEndTrackTime int32
			err = json.Unmarshal(*valAudioCdEndTrackTimeCap, &valueForAudioCdEndTrackTime)
			if err != nil {
				return err
			}
			this.AudioCdEndTrackTime = valueForAudioCdEndTrackTime
		}
	}
	
	if valAudioCdStartTrack, ok := objMap["audioCdStartTrack"]; ok {
		if valAudioCdStartTrack != nil {
			var valueForAudioCdStartTrack int32
			err = json.Unmarshal(*valAudioCdStartTrack, &valueForAudioCdStartTrack)
			if err != nil {
				return err
			}
			this.AudioCdStartTrack = valueForAudioCdStartTrack
		}
	}
	if valAudioCdStartTrackCap, ok := objMap["AudioCdStartTrack"]; ok {
		if valAudioCdStartTrackCap != nil {
			var valueForAudioCdStartTrack int32
			err = json.Unmarshal(*valAudioCdStartTrackCap, &valueForAudioCdStartTrack)
			if err != nil {
				return err
			}
			this.AudioCdStartTrack = valueForAudioCdStartTrack
		}
	}
	
	if valAudioCdStartTrackTime, ok := objMap["audioCdStartTrackTime"]; ok {
		if valAudioCdStartTrackTime != nil {
			var valueForAudioCdStartTrackTime int32
			err = json.Unmarshal(*valAudioCdStartTrackTime, &valueForAudioCdStartTrackTime)
			if err != nil {
				return err
			}
			this.AudioCdStartTrackTime = valueForAudioCdStartTrackTime
		}
	}
	if valAudioCdStartTrackTimeCap, ok := objMap["AudioCdStartTrackTime"]; ok {
		if valAudioCdStartTrackTimeCap != nil {
			var valueForAudioCdStartTrackTime int32
			err = json.Unmarshal(*valAudioCdStartTrackTimeCap, &valueForAudioCdStartTrackTime)
			if err != nil {
				return err
			}
			this.AudioCdStartTrackTime = valueForAudioCdStartTrackTime
		}
	}
	
	if valEmbedded, ok := objMap["embedded"]; ok {
		if valEmbedded != nil {
			var valueForEmbedded *bool
			err = json.Unmarshal(*valEmbedded, &valueForEmbedded)
			if err != nil {
				return err
			}
			this.Embedded = valueForEmbedded
		}
	}
	if valEmbeddedCap, ok := objMap["Embedded"]; ok {
		if valEmbeddedCap != nil {
			var valueForEmbedded *bool
			err = json.Unmarshal(*valEmbeddedCap, &valueForEmbedded)
			if err != nil {
				return err
			}
			this.Embedded = valueForEmbedded
		}
	}
	
	if valHideAtShowing, ok := objMap["hideAtShowing"]; ok {
		if valHideAtShowing != nil {
			var valueForHideAtShowing *bool
			err = json.Unmarshal(*valHideAtShowing, &valueForHideAtShowing)
			if err != nil {
				return err
			}
			this.HideAtShowing = valueForHideAtShowing
		}
	}
	if valHideAtShowingCap, ok := objMap["HideAtShowing"]; ok {
		if valHideAtShowingCap != nil {
			var valueForHideAtShowing *bool
			err = json.Unmarshal(*valHideAtShowingCap, &valueForHideAtShowing)
			if err != nil {
				return err
			}
			this.HideAtShowing = valueForHideAtShowing
		}
	}
	
	if valPlayLoopMode, ok := objMap["playLoopMode"]; ok {
		if valPlayLoopMode != nil {
			var valueForPlayLoopMode *bool
			err = json.Unmarshal(*valPlayLoopMode, &valueForPlayLoopMode)
			if err != nil {
				return err
			}
			this.PlayLoopMode = valueForPlayLoopMode
		}
	}
	if valPlayLoopModeCap, ok := objMap["PlayLoopMode"]; ok {
		if valPlayLoopModeCap != nil {
			var valueForPlayLoopMode *bool
			err = json.Unmarshal(*valPlayLoopModeCap, &valueForPlayLoopMode)
			if err != nil {
				return err
			}
			this.PlayLoopMode = valueForPlayLoopMode
		}
	}
	
	if valPlayMode, ok := objMap["playMode"]; ok {
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
	if valPlayModeCap, ok := objMap["PlayMode"]; ok {
		if valPlayModeCap != nil {
			var valueForPlayMode string
			err = json.Unmarshal(*valPlayModeCap, &valueForPlayMode)
			if err != nil {
				var valueForPlayModeInt int32
				err = json.Unmarshal(*valPlayModeCap, &valueForPlayModeInt)
				if err != nil {
					return err
				}
				this.PlayMode = string(valueForPlayModeInt)
			} else {
				this.PlayMode = valueForPlayMode
			}
		}
	}
	
	if valVolume, ok := objMap["volume"]; ok {
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
	if valVolumeCap, ok := objMap["Volume"]; ok {
		if valVolumeCap != nil {
			var valueForVolume string
			err = json.Unmarshal(*valVolumeCap, &valueForVolume)
			if err != nil {
				var valueForVolumeInt int32
				err = json.Unmarshal(*valVolumeCap, &valueForVolumeInt)
				if err != nil {
					return err
				}
				this.Volume = string(valueForVolumeInt)
			} else {
				this.Volume = valueForVolume
			}
		}
	}
	
	if valBase64Data, ok := objMap["base64Data"]; ok {
		if valBase64Data != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64Data, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	if valBase64DataCap, ok := objMap["Base64Data"]; ok {
		if valBase64DataCap != nil {
			var valueForBase64Data string
			err = json.Unmarshal(*valBase64DataCap, &valueForBase64Data)
			if err != nil {
				return err
			}
			this.Base64Data = valueForBase64Data
		}
	}
	
	if valPlayAcrossSlides, ok := objMap["playAcrossSlides"]; ok {
		if valPlayAcrossSlides != nil {
			var valueForPlayAcrossSlides *bool
			err = json.Unmarshal(*valPlayAcrossSlides, &valueForPlayAcrossSlides)
			if err != nil {
				return err
			}
			this.PlayAcrossSlides = valueForPlayAcrossSlides
		}
	}
	if valPlayAcrossSlidesCap, ok := objMap["PlayAcrossSlides"]; ok {
		if valPlayAcrossSlidesCap != nil {
			var valueForPlayAcrossSlides *bool
			err = json.Unmarshal(*valPlayAcrossSlidesCap, &valueForPlayAcrossSlides)
			if err != nil {
				return err
			}
			this.PlayAcrossSlides = valueForPlayAcrossSlides
		}
	}
	
	if valRewindAudio, ok := objMap["rewindAudio"]; ok {
		if valRewindAudio != nil {
			var valueForRewindAudio *bool
			err = json.Unmarshal(*valRewindAudio, &valueForRewindAudio)
			if err != nil {
				return err
			}
			this.RewindAudio = valueForRewindAudio
		}
	}
	if valRewindAudioCap, ok := objMap["RewindAudio"]; ok {
		if valRewindAudioCap != nil {
			var valueForRewindAudio *bool
			err = json.Unmarshal(*valRewindAudioCap, &valueForRewindAudio)
			if err != nil {
				return err
			}
			this.RewindAudio = valueForRewindAudio
		}
	}
	
	if valPictureFillFormat, ok := objMap["pictureFillFormat"]; ok {
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
	if valPictureFillFormatCap, ok := objMap["PictureFillFormat"]; ok {
		if valPictureFillFormatCap != nil {
			var valueForPictureFillFormat PictureFill
			err = json.Unmarshal(*valPictureFillFormatCap, &valueForPictureFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PictureFill", *valPictureFillFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPictureFillFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPictureFill)
			if ok {
				this.PictureFillFormat = vInterfaceObject
			}
		}
	}

	return nil
}
