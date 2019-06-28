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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Main sequence.
	getMainSequence() []Effect
	setMainSequence(newValue []Effect)

	// Interactive sequence list.
	getInteractiveSequences() []InteractiveSequence
	setInteractiveSequences(newValue []InteractiveSequence)
}

type SlideAnimation struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Main sequence.
	MainSequence []Effect `json:"MainSequence,omitempty"`

	// Interactive sequence list.
	InteractiveSequences []InteractiveSequence `json:"InteractiveSequences,omitempty"`
}

func (this SlideAnimation) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this SlideAnimation) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this SlideAnimation) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this SlideAnimation) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this SlideAnimation) getMainSequence() []Effect {
	return this.MainSequence
}

func (this SlideAnimation) setMainSequence(newValue []Effect) {
	this.MainSequence = newValue
}
func (this SlideAnimation) getInteractiveSequences() []InteractiveSequence {
	return this.InteractiveSequences
}

func (this SlideAnimation) setInteractiveSequences(newValue []InteractiveSequence) {
	this.InteractiveSequences = newValue
}

func (this *SlideAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["SelfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}
	
	if valMainSequence, ok := objMap["MainSequence"]; ok {
		if valMainSequence != nil {
			var valueForMainSequence []Effect
			err = json.Unmarshal(*valMainSequence, &valueForMainSequence)
			if err != nil {
				return err
			}
			this.MainSequence = valueForMainSequence
		}
	}
	
	if valInteractiveSequences, ok := objMap["InteractiveSequences"]; ok {
		if valInteractiveSequences != nil {
			var valueForInteractiveSequences []InteractiveSequence
			err = json.Unmarshal(*valInteractiveSequences, &valueForInteractiveSequences)
			if err != nil {
				return err
			}
			this.InteractiveSequences = valueForInteractiveSequences
		}
	}

    return nil
}
