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
	getShapes() IResourceUri
	setShapes(newValue IResourceUri)

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

	// Returns or sets a last track index.
	getAudioCdEndTrack() int32
	setAudioCdEndTrack(newValue int32)

	// Returns or sets a last track time.
	getAudioCdEndTrackTime() int32
	setAudioCdEndTrackTime(newValue int32)

	// Returns or sets a start track index.
	getAudioCdStartTrack() int32
	setAudioCdStartTrack(newValue int32)

	// Returns or sets a start track time. 
	getAudioCdStartTrackTime() int32
	setAudioCdStartTrackTime(newValue int32)

	// Determines whether a sound is embedded to a presentation.
	getEmbedded() bool
	setEmbedded(newValue bool)

	// Determines whether an AudioFrame is hidden.
	getHideAtShowing() bool
	setHideAtShowing(newValue bool)

	// Determines whether an audio is looped. 
	getPlayLoopMode() bool
	setPlayLoopMode(newValue bool)

	// Returns or sets the audio play mode.
	getPlayMode() string
	setPlayMode(newValue string)

	// Returns or sets the audio volume.
	getVolume() string
	setVolume(newValue string)

	// Audio data encoded in base64.
	getBase64Data() string
	setBase64Data(newValue string)
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
	Hidden bool `json:"Hidden"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUri `json:"Shapes,omitempty"`

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

	// Returns or sets a last track index.
	AudioCdEndTrack int32 `json:"AudioCdEndTrack,omitempty"`

	// Returns or sets a last track time.
	AudioCdEndTrackTime int32 `json:"AudioCdEndTrackTime,omitempty"`

	// Returns or sets a start track index.
	AudioCdStartTrack int32 `json:"AudioCdStartTrack,omitempty"`

	// Returns or sets a start track time. 
	AudioCdStartTrackTime int32 `json:"AudioCdStartTrackTime,omitempty"`

	// Determines whether a sound is embedded to a presentation.
	Embedded bool `json:"Embedded"`

	// Determines whether an AudioFrame is hidden.
	HideAtShowing bool `json:"HideAtShowing"`

	// Determines whether an audio is looped. 
	PlayLoopMode bool `json:"PlayLoopMode"`

	// Returns or sets the audio play mode.
	PlayMode string `json:"PlayMode,omitempty"`

	// Returns or sets the audio volume.
	Volume string `json:"Volume,omitempty"`

	// Audio data encoded in base64.
	Base64Data string `json:"Base64Data,omitempty"`
}

func NewAudioFrame() *AudioFrame {
	instance := new(AudioFrame)
	instance.Type_ = "AudioFrame"
	instance.ShapeType = "Custom"
	instance.PlayMode = ""
	instance.Volume = ""
	return instance
}

func (this *AudioFrame) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *AudioFrame) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *AudioFrame) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *AudioFrame) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *AudioFrame) getName() string {
	return this.Name
}

func (this *AudioFrame) setName(newValue string) {
	this.Name = newValue
}
func (this *AudioFrame) getWidth() float64 {
	return this.Width
}

func (this *AudioFrame) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *AudioFrame) getHeight() float64 {
	return this.Height
}

func (this *AudioFrame) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *AudioFrame) getAlternativeText() string {
	return this.AlternativeText
}

func (this *AudioFrame) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *AudioFrame) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *AudioFrame) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *AudioFrame) getHidden() bool {
	return this.Hidden
}

func (this *AudioFrame) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *AudioFrame) getX() float64 {
	return this.X
}

func (this *AudioFrame) setX(newValue float64) {
	this.X = newValue
}
func (this *AudioFrame) getY() float64 {
	return this.Y
}

func (this *AudioFrame) setY(newValue float64) {
	this.Y = newValue
}
func (this *AudioFrame) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *AudioFrame) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *AudioFrame) getShapes() IResourceUri {
	return this.Shapes
}

func (this *AudioFrame) setShapes(newValue IResourceUri) {
	this.Shapes = newValue
}
func (this *AudioFrame) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *AudioFrame) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *AudioFrame) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *AudioFrame) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *AudioFrame) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *AudioFrame) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *AudioFrame) getType() string {
	return this.Type_
}

func (this *AudioFrame) setType(newValue string) {
	this.Type_ = newValue
}
func (this *AudioFrame) getShapeType() string {
	return this.ShapeType
}

func (this *AudioFrame) setShapeType(newValue string) {
	this.ShapeType = newValue
}
func (this *AudioFrame) getAudioCdEndTrack() int32 {
	return this.AudioCdEndTrack
}

func (this *AudioFrame) setAudioCdEndTrack(newValue int32) {
	this.AudioCdEndTrack = newValue
}
func (this *AudioFrame) getAudioCdEndTrackTime() int32 {
	return this.AudioCdEndTrackTime
}

func (this *AudioFrame) setAudioCdEndTrackTime(newValue int32) {
	this.AudioCdEndTrackTime = newValue
}
func (this *AudioFrame) getAudioCdStartTrack() int32 {
	return this.AudioCdStartTrack
}

func (this *AudioFrame) setAudioCdStartTrack(newValue int32) {
	this.AudioCdStartTrack = newValue
}
func (this *AudioFrame) getAudioCdStartTrackTime() int32 {
	return this.AudioCdStartTrackTime
}

func (this *AudioFrame) setAudioCdStartTrackTime(newValue int32) {
	this.AudioCdStartTrackTime = newValue
}
func (this *AudioFrame) getEmbedded() bool {
	return this.Embedded
}

func (this *AudioFrame) setEmbedded(newValue bool) {
	this.Embedded = newValue
}
func (this *AudioFrame) getHideAtShowing() bool {
	return this.HideAtShowing
}

func (this *AudioFrame) setHideAtShowing(newValue bool) {
	this.HideAtShowing = newValue
}
func (this *AudioFrame) getPlayLoopMode() bool {
	return this.PlayLoopMode
}

func (this *AudioFrame) setPlayLoopMode(newValue bool) {
	this.PlayLoopMode = newValue
}
func (this *AudioFrame) getPlayMode() string {
	return this.PlayMode
}

func (this *AudioFrame) setPlayMode(newValue string) {
	this.PlayMode = newValue
}
func (this *AudioFrame) getVolume() string {
	return this.Volume
}

func (this *AudioFrame) setVolume(newValue string) {
	this.Volume = newValue
}
func (this *AudioFrame) getBase64Data() string {
	return this.Base64Data
}

func (this *AudioFrame) setBase64Data(newValue string) {
	this.Base64Data = newValue
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
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUri
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
			var valueForEmbedded bool
			err = json.Unmarshal(*valEmbedded, &valueForEmbedded)
			if err != nil {
				return err
			}
			this.Embedded = valueForEmbedded
		}
	}
	if valEmbeddedCap, ok := objMap["Embedded"]; ok {
		if valEmbeddedCap != nil {
			var valueForEmbedded bool
			err = json.Unmarshal(*valEmbeddedCap, &valueForEmbedded)
			if err != nil {
				return err
			}
			this.Embedded = valueForEmbedded
		}
	}
	
	if valHideAtShowing, ok := objMap["hideAtShowing"]; ok {
		if valHideAtShowing != nil {
			var valueForHideAtShowing bool
			err = json.Unmarshal(*valHideAtShowing, &valueForHideAtShowing)
			if err != nil {
				return err
			}
			this.HideAtShowing = valueForHideAtShowing
		}
	}
	if valHideAtShowingCap, ok := objMap["HideAtShowing"]; ok {
		if valHideAtShowingCap != nil {
			var valueForHideAtShowing bool
			err = json.Unmarshal(*valHideAtShowingCap, &valueForHideAtShowing)
			if err != nil {
				return err
			}
			this.HideAtShowing = valueForHideAtShowing
		}
	}
	
	if valPlayLoopMode, ok := objMap["playLoopMode"]; ok {
		if valPlayLoopMode != nil {
			var valueForPlayLoopMode bool
			err = json.Unmarshal(*valPlayLoopMode, &valueForPlayLoopMode)
			if err != nil {
				return err
			}
			this.PlayLoopMode = valueForPlayLoopMode
		}
	}
	if valPlayLoopModeCap, ok := objMap["PlayLoopMode"]; ok {
		if valPlayLoopModeCap != nil {
			var valueForPlayLoopMode bool
			err = json.Unmarshal(*valPlayLoopModeCap, &valueForPlayLoopMode)
			if err != nil {
				return err
			}
			this.PlayLoopMode = valueForPlayLoopMode
		}
	}
	this.PlayMode = ""
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
	this.Volume = ""
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

    return nil
}
