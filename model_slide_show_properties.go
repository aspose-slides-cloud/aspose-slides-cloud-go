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

// Slide show properties.
type ISlideShowProperties interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Loop slide show.
	GetLoop() *bool
	SetLoop(newValue *bool)

	// Start slide in the slide show.
	GetStartSlide() int32
	SetStartSlide(newValue int32)

	// End slides in the slide show.
	GetEndSlide() int32
	SetEndSlide(newValue int32)

	// Pen color.
	GetPenColor() string
	SetPenColor(newValue string)

	// Show animation.
	GetShowAnimation() *bool
	SetShowAnimation(newValue *bool)

	// Show narrration.
	GetShowNarration() *bool
	SetShowNarration(newValue *bool)

	// Show media controls.
	GetShowMediaControls() *bool
	SetShowMediaControls(newValue *bool)

	// Use timings.
	GetUseTimings() *bool
	SetUseTimings(newValue *bool)

	// Slide show type.
	GetSlideShowType() string
	SetSlideShowType(newValue string)

	// Show scroll bar. Applied with BrowsedByIndividual slide show type.
	GetShowScrollbar() *bool
	SetShowScrollbar(newValue *bool)
}

type SlideShowProperties struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Loop slide show.
	Loop *bool `json:"Loop"`

	// Start slide in the slide show.
	StartSlide int32 `json:"StartSlide,omitempty"`

	// End slides in the slide show.
	EndSlide int32 `json:"EndSlide,omitempty"`

	// Pen color.
	PenColor string `json:"PenColor,omitempty"`

	// Show animation.
	ShowAnimation *bool `json:"ShowAnimation"`

	// Show narrration.
	ShowNarration *bool `json:"ShowNarration"`

	// Show media controls.
	ShowMediaControls *bool `json:"ShowMediaControls"`

	// Use timings.
	UseTimings *bool `json:"UseTimings"`

	// Slide show type.
	SlideShowType string `json:"SlideShowType,omitempty"`

	// Show scroll bar. Applied with BrowsedByIndividual slide show type.
	ShowScrollbar *bool `json:"ShowScrollbar"`
}

func NewSlideShowProperties() *SlideShowProperties {
	instance := new(SlideShowProperties)
	return instance
}

func (this *SlideShowProperties) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SlideShowProperties) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SlideShowProperties) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SlideShowProperties) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SlideShowProperties) GetLoop() *bool {
	return this.Loop
}

func (this *SlideShowProperties) SetLoop(newValue *bool) {
	this.Loop = newValue
}
func (this *SlideShowProperties) GetStartSlide() int32 {
	return this.StartSlide
}

func (this *SlideShowProperties) SetStartSlide(newValue int32) {
	this.StartSlide = newValue
}
func (this *SlideShowProperties) GetEndSlide() int32 {
	return this.EndSlide
}

func (this *SlideShowProperties) SetEndSlide(newValue int32) {
	this.EndSlide = newValue
}
func (this *SlideShowProperties) GetPenColor() string {
	return this.PenColor
}

func (this *SlideShowProperties) SetPenColor(newValue string) {
	this.PenColor = newValue
}
func (this *SlideShowProperties) GetShowAnimation() *bool {
	return this.ShowAnimation
}

func (this *SlideShowProperties) SetShowAnimation(newValue *bool) {
	this.ShowAnimation = newValue
}
func (this *SlideShowProperties) GetShowNarration() *bool {
	return this.ShowNarration
}

func (this *SlideShowProperties) SetShowNarration(newValue *bool) {
	this.ShowNarration = newValue
}
func (this *SlideShowProperties) GetShowMediaControls() *bool {
	return this.ShowMediaControls
}

func (this *SlideShowProperties) SetShowMediaControls(newValue *bool) {
	this.ShowMediaControls = newValue
}
func (this *SlideShowProperties) GetUseTimings() *bool {
	return this.UseTimings
}

func (this *SlideShowProperties) SetUseTimings(newValue *bool) {
	this.UseTimings = newValue
}
func (this *SlideShowProperties) GetSlideShowType() string {
	return this.SlideShowType
}

func (this *SlideShowProperties) SetSlideShowType(newValue string) {
	this.SlideShowType = newValue
}
func (this *SlideShowProperties) GetShowScrollbar() *bool {
	return this.ShowScrollbar
}

func (this *SlideShowProperties) SetShowScrollbar(newValue *bool) {
	this.ShowScrollbar = newValue
}

func (this *SlideShowProperties) UnmarshalJSON(b []byte) error {
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
	
	if valLoop, ok := GetMapValue(objMap, "loop"); ok {
		if valLoop != nil {
			var valueForLoop *bool
			err = json.Unmarshal(*valLoop, &valueForLoop)
			if err != nil {
				return err
			}
			this.Loop = valueForLoop
		}
	}
	
	if valStartSlide, ok := GetMapValue(objMap, "startSlide"); ok {
		if valStartSlide != nil {
			var valueForStartSlide int32
			err = json.Unmarshal(*valStartSlide, &valueForStartSlide)
			if err != nil {
				return err
			}
			this.StartSlide = valueForStartSlide
		}
	}
	
	if valEndSlide, ok := GetMapValue(objMap, "endSlide"); ok {
		if valEndSlide != nil {
			var valueForEndSlide int32
			err = json.Unmarshal(*valEndSlide, &valueForEndSlide)
			if err != nil {
				return err
			}
			this.EndSlide = valueForEndSlide
		}
	}
	
	if valPenColor, ok := GetMapValue(objMap, "penColor"); ok {
		if valPenColor != nil {
			var valueForPenColor string
			err = json.Unmarshal(*valPenColor, &valueForPenColor)
			if err != nil {
				return err
			}
			this.PenColor = valueForPenColor
		}
	}
	
	if valShowAnimation, ok := GetMapValue(objMap, "showAnimation"); ok {
		if valShowAnimation != nil {
			var valueForShowAnimation *bool
			err = json.Unmarshal(*valShowAnimation, &valueForShowAnimation)
			if err != nil {
				return err
			}
			this.ShowAnimation = valueForShowAnimation
		}
	}
	
	if valShowNarration, ok := GetMapValue(objMap, "showNarration"); ok {
		if valShowNarration != nil {
			var valueForShowNarration *bool
			err = json.Unmarshal(*valShowNarration, &valueForShowNarration)
			if err != nil {
				return err
			}
			this.ShowNarration = valueForShowNarration
		}
	}
	
	if valShowMediaControls, ok := GetMapValue(objMap, "showMediaControls"); ok {
		if valShowMediaControls != nil {
			var valueForShowMediaControls *bool
			err = json.Unmarshal(*valShowMediaControls, &valueForShowMediaControls)
			if err != nil {
				return err
			}
			this.ShowMediaControls = valueForShowMediaControls
		}
	}
	
	if valUseTimings, ok := GetMapValue(objMap, "useTimings"); ok {
		if valUseTimings != nil {
			var valueForUseTimings *bool
			err = json.Unmarshal(*valUseTimings, &valueForUseTimings)
			if err != nil {
				return err
			}
			this.UseTimings = valueForUseTimings
		}
	}
	
	if valSlideShowType, ok := GetMapValue(objMap, "slideShowType"); ok {
		if valSlideShowType != nil {
			var valueForSlideShowType string
			err = json.Unmarshal(*valSlideShowType, &valueForSlideShowType)
			if err != nil {
				var valueForSlideShowTypeInt int32
				err = json.Unmarshal(*valSlideShowType, &valueForSlideShowTypeInt)
				if err != nil {
					return err
				}
				this.SlideShowType = string(valueForSlideShowTypeInt)
			} else {
				this.SlideShowType = valueForSlideShowType
			}
		}
	}
	
	if valShowScrollbar, ok := GetMapValue(objMap, "showScrollbar"); ok {
		if valShowScrollbar != nil {
			var valueForShowScrollbar *bool
			err = json.Unmarshal(*valShowScrollbar, &valueForShowScrollbar)
			if err != nil {
				return err
			}
			this.ShowScrollbar = valueForShowScrollbar
		}
	}

	return nil
}
