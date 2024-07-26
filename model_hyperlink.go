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

// Hyperlink
type IHyperlink interface {

	// If true Hypelink is not applied. 
	GetIsDisabled() *bool
	SetIsDisabled(newValue *bool)

	// Type of HyperLink action             
	GetActionType() string
	SetActionType(newValue string)

	// Specifies the external URL
	GetExternalUrl() string
	SetExternalUrl(newValue string)

	// Index of the target slide
	GetTargetSlideIndex() int32
	SetTargetSlideIndex(newValue int32)

	// Target frame
	GetTargetFrame() string
	SetTargetFrame(newValue string)

	// Hyperlink tooltip
	GetTooltip() string
	SetTooltip(newValue string)

	// Makes hyperlink viewed when it is invoked.             
	GetHistory() *bool
	SetHistory(newValue *bool)

	// Determines whether the hyperlink should be highlighted on click.
	GetHighlightClick() *bool
	SetHighlightClick(newValue *bool)

	// Determines whether the sound should be stopped on hyperlink click
	GetStopSoundOnClick() *bool
	SetStopSoundOnClick(newValue *bool)

	// Represents the source of hyperlink color
	GetColorSource() string
	SetColorSource(newValue string)

	// Audio data encoded in base64. Represents the playing sound of the hyperlink. 
	GetSoundBase64() string
	SetSoundBase64(newValue string)
}

type Hyperlink struct {

	// If true Hypelink is not applied. 
	IsDisabled *bool `json:"IsDisabled"`

	// Type of HyperLink action             
	ActionType string `json:"ActionType"`

	// Specifies the external URL
	ExternalUrl string `json:"ExternalUrl,omitempty"`

	// Index of the target slide
	TargetSlideIndex int32 `json:"TargetSlideIndex,omitempty"`

	// Target frame
	TargetFrame string `json:"TargetFrame,omitempty"`

	// Hyperlink tooltip
	Tooltip string `json:"Tooltip,omitempty"`

	// Makes hyperlink viewed when it is invoked.             
	History *bool `json:"History"`

	// Determines whether the hyperlink should be highlighted on click.
	HighlightClick *bool `json:"HighlightClick"`

	// Determines whether the sound should be stopped on hyperlink click
	StopSoundOnClick *bool `json:"StopSoundOnClick"`

	// Represents the source of hyperlink color
	ColorSource string `json:"ColorSource,omitempty"`

	// Audio data encoded in base64. Represents the playing sound of the hyperlink. 
	SoundBase64 string `json:"SoundBase64,omitempty"`
}

func NewHyperlink() *Hyperlink {
	instance := new(Hyperlink)
	instance.ActionType = "NoAction"
	return instance
}

func (this *Hyperlink) GetIsDisabled() *bool {
	return this.IsDisabled
}

func (this *Hyperlink) SetIsDisabled(newValue *bool) {
	this.IsDisabled = newValue
}
func (this *Hyperlink) GetActionType() string {
	return this.ActionType
}

func (this *Hyperlink) SetActionType(newValue string) {
	this.ActionType = newValue
}
func (this *Hyperlink) GetExternalUrl() string {
	return this.ExternalUrl
}

func (this *Hyperlink) SetExternalUrl(newValue string) {
	this.ExternalUrl = newValue
}
func (this *Hyperlink) GetTargetSlideIndex() int32 {
	return this.TargetSlideIndex
}

func (this *Hyperlink) SetTargetSlideIndex(newValue int32) {
	this.TargetSlideIndex = newValue
}
func (this *Hyperlink) GetTargetFrame() string {
	return this.TargetFrame
}

func (this *Hyperlink) SetTargetFrame(newValue string) {
	this.TargetFrame = newValue
}
func (this *Hyperlink) GetTooltip() string {
	return this.Tooltip
}

func (this *Hyperlink) SetTooltip(newValue string) {
	this.Tooltip = newValue
}
func (this *Hyperlink) GetHistory() *bool {
	return this.History
}

func (this *Hyperlink) SetHistory(newValue *bool) {
	this.History = newValue
}
func (this *Hyperlink) GetHighlightClick() *bool {
	return this.HighlightClick
}

func (this *Hyperlink) SetHighlightClick(newValue *bool) {
	this.HighlightClick = newValue
}
func (this *Hyperlink) GetStopSoundOnClick() *bool {
	return this.StopSoundOnClick
}

func (this *Hyperlink) SetStopSoundOnClick(newValue *bool) {
	this.StopSoundOnClick = newValue
}
func (this *Hyperlink) GetColorSource() string {
	return this.ColorSource
}

func (this *Hyperlink) SetColorSource(newValue string) {
	this.ColorSource = newValue
}
func (this *Hyperlink) GetSoundBase64() string {
	return this.SoundBase64
}

func (this *Hyperlink) SetSoundBase64(newValue string) {
	this.SoundBase64 = newValue
}

func (this *Hyperlink) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valIsDisabled, ok := GetMapValue(objMap, "isDisabled"); ok {
		if valIsDisabled != nil {
			var valueForIsDisabled *bool
			err = json.Unmarshal(*valIsDisabled, &valueForIsDisabled)
			if err != nil {
				return err
			}
			this.IsDisabled = valueForIsDisabled
		}
	}
	this.ActionType = "NoAction"
	if valActionType, ok := GetMapValue(objMap, "actionType"); ok {
		if valActionType != nil {
			var valueForActionType string
			err = json.Unmarshal(*valActionType, &valueForActionType)
			if err != nil {
				var valueForActionTypeInt int32
				err = json.Unmarshal(*valActionType, &valueForActionTypeInt)
				if err != nil {
					return err
				}
				this.ActionType = string(valueForActionTypeInt)
			} else {
				this.ActionType = valueForActionType
			}
		}
	}
	
	if valExternalUrl, ok := GetMapValue(objMap, "externalUrl"); ok {
		if valExternalUrl != nil {
			var valueForExternalUrl string
			err = json.Unmarshal(*valExternalUrl, &valueForExternalUrl)
			if err != nil {
				return err
			}
			this.ExternalUrl = valueForExternalUrl
		}
	}
	
	if valTargetSlideIndex, ok := GetMapValue(objMap, "targetSlideIndex"); ok {
		if valTargetSlideIndex != nil {
			var valueForTargetSlideIndex int32
			err = json.Unmarshal(*valTargetSlideIndex, &valueForTargetSlideIndex)
			if err != nil {
				return err
			}
			this.TargetSlideIndex = valueForTargetSlideIndex
		}
	}
	
	if valTargetFrame, ok := GetMapValue(objMap, "targetFrame"); ok {
		if valTargetFrame != nil {
			var valueForTargetFrame string
			err = json.Unmarshal(*valTargetFrame, &valueForTargetFrame)
			if err != nil {
				return err
			}
			this.TargetFrame = valueForTargetFrame
		}
	}
	
	if valTooltip, ok := GetMapValue(objMap, "tooltip"); ok {
		if valTooltip != nil {
			var valueForTooltip string
			err = json.Unmarshal(*valTooltip, &valueForTooltip)
			if err != nil {
				return err
			}
			this.Tooltip = valueForTooltip
		}
	}
	
	if valHistory, ok := GetMapValue(objMap, "history"); ok {
		if valHistory != nil {
			var valueForHistory *bool
			err = json.Unmarshal(*valHistory, &valueForHistory)
			if err != nil {
				return err
			}
			this.History = valueForHistory
		}
	}
	
	if valHighlightClick, ok := GetMapValue(objMap, "highlightClick"); ok {
		if valHighlightClick != nil {
			var valueForHighlightClick *bool
			err = json.Unmarshal(*valHighlightClick, &valueForHighlightClick)
			if err != nil {
				return err
			}
			this.HighlightClick = valueForHighlightClick
		}
	}
	
	if valStopSoundOnClick, ok := GetMapValue(objMap, "stopSoundOnClick"); ok {
		if valStopSoundOnClick != nil {
			var valueForStopSoundOnClick *bool
			err = json.Unmarshal(*valStopSoundOnClick, &valueForStopSoundOnClick)
			if err != nil {
				return err
			}
			this.StopSoundOnClick = valueForStopSoundOnClick
		}
	}
	
	if valColorSource, ok := GetMapValue(objMap, "colorSource"); ok {
		if valColorSource != nil {
			var valueForColorSource string
			err = json.Unmarshal(*valColorSource, &valueForColorSource)
			if err != nil {
				var valueForColorSourceInt int32
				err = json.Unmarshal(*valColorSource, &valueForColorSourceInt)
				if err != nil {
					return err
				}
				this.ColorSource = string(valueForColorSourceInt)
			} else {
				this.ColorSource = valueForColorSource
			}
		}
	}
	
	if valSoundBase64, ok := GetMapValue(objMap, "soundBase64"); ok {
		if valSoundBase64 != nil {
			var valueForSoundBase64 string
			err = json.Unmarshal(*valSoundBase64, &valueForSoundBase64)
			if err != nil {
				return err
			}
			this.SoundBase64 = valueForSoundBase64
		}
	}

	return nil
}
