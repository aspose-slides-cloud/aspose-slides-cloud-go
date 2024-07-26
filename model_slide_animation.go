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

// Represents comments collection of slide
type ISlideAnimation interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Main sequence.
	GetMainSequence() []IEffect
	SetMainSequence(newValue []IEffect)

	// Interactive sequence list.
	GetInteractiveSequences() []IInteractiveSequence
	SetInteractiveSequences(newValue []IInteractiveSequence)
}

type SlideAnimation struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Main sequence.
	MainSequence []IEffect `json:"MainSequence,omitempty"`

	// Interactive sequence list.
	InteractiveSequences []IInteractiveSequence `json:"InteractiveSequences,omitempty"`
}

func NewSlideAnimation() *SlideAnimation {
	instance := new(SlideAnimation)
	return instance
}

func (this *SlideAnimation) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SlideAnimation) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SlideAnimation) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SlideAnimation) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SlideAnimation) GetMainSequence() []IEffect {
	return this.MainSequence
}

func (this *SlideAnimation) SetMainSequence(newValue []IEffect) {
	this.MainSequence = newValue
}
func (this *SlideAnimation) GetInteractiveSequences() []IInteractiveSequence {
	return this.InteractiveSequences
}

func (this *SlideAnimation) SetInteractiveSequences(newValue []IInteractiveSequence) {
	this.InteractiveSequences = newValue
}

func (this *SlideAnimation) UnmarshalJSON(b []byte) error {
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
	
	if valMainSequence, ok := GetMapValue(objMap, "mainSequence"); ok {
		if valMainSequence != nil {
			var valueForMainSequence []json.RawMessage
			err = json.Unmarshal(*valMainSequence, &valueForMainSequence)
			if err != nil {
				return err
			}
			valueForIMainSequence := make([]IEffect, len(valueForMainSequence))
			for i, v := range valueForMainSequence {
				vObject, err := createObjectForType("Effect", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIMainSequence[i] = vObject.(IEffect)
				}
			}
			this.MainSequence = valueForIMainSequence
		}
	}
	
	if valInteractiveSequences, ok := GetMapValue(objMap, "interactiveSequences"); ok {
		if valInteractiveSequences != nil {
			var valueForInteractiveSequences []json.RawMessage
			err = json.Unmarshal(*valInteractiveSequences, &valueForInteractiveSequences)
			if err != nil {
				return err
			}
			valueForIInteractiveSequences := make([]IInteractiveSequence, len(valueForInteractiveSequences))
			for i, v := range valueForInteractiveSequences {
				vObject, err := createObjectForType("InteractiveSequence", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIInteractiveSequences[i] = vObject.(IInteractiveSequence)
				}
			}
			this.InteractiveSequences = valueForIInteractiveSequences
		}
	}

	return nil
}
