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

// Master slide.
type IMasterSlide interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// List of layout slide links.
	getLayoutSlides() []IResourceUriElement
	setLayoutSlides(newValue []IResourceUriElement)

	// List of depending slide links.
	getDependingSlides() []IResourceUriElement
	setDependingSlides(newValue []IResourceUriElement)
}

type MasterSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// List of layout slide links.
	LayoutSlides []IResourceUriElement `json:"LayoutSlides,omitempty"`

	// List of depending slide links.
	DependingSlides []IResourceUriElement `json:"DependingSlides,omitempty"`
}

func (this *MasterSlide) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *MasterSlide) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *MasterSlide) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *MasterSlide) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *MasterSlide) getName() string {
	return this.Name
}

func (this *MasterSlide) setName(newValue string) {
	this.Name = newValue
}
func (this *MasterSlide) getLayoutSlides() []IResourceUriElement {
	return this.LayoutSlides
}

func (this *MasterSlide) setLayoutSlides(newValue []IResourceUriElement) {
	this.LayoutSlides = newValue
}
func (this *MasterSlide) getDependingSlides() []IResourceUriElement {
	return this.DependingSlides
}

func (this *MasterSlide) setDependingSlides(newValue []IResourceUriElement) {
	this.DependingSlides = newValue
}

func (this *MasterSlide) UnmarshalJSON(b []byte) error {
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
	
	if valLayoutSlides, ok := objMap["layoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides []ResourceUriElement
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			valueForILayoutSlides := make([]IResourceUriElement, len(valueForLayoutSlides))
			for i, v := range valueForLayoutSlides {
				valueForILayoutSlides[i] = IResourceUriElement(&v)
			}
			this.LayoutSlides = valueForILayoutSlides
		}
	}
	if valLayoutSlidesCap, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlidesCap != nil {
			var valueForLayoutSlides []ResourceUriElement
			err = json.Unmarshal(*valLayoutSlidesCap, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			valueForILayoutSlides := make([]IResourceUriElement, len(valueForLayoutSlides))
			for i, v := range valueForLayoutSlides {
				valueForILayoutSlides[i] = IResourceUriElement(&v)
			}
			this.LayoutSlides = valueForILayoutSlides
		}
	}
	
	if valDependingSlides, ok := objMap["dependingSlides"]; ok {
		if valDependingSlides != nil {
			var valueForDependingSlides []ResourceUriElement
			err = json.Unmarshal(*valDependingSlides, &valueForDependingSlides)
			if err != nil {
				return err
			}
			valueForIDependingSlides := make([]IResourceUriElement, len(valueForDependingSlides))
			for i, v := range valueForDependingSlides {
				valueForIDependingSlides[i] = IResourceUriElement(&v)
			}
			this.DependingSlides = valueForIDependingSlides
		}
	}
	if valDependingSlidesCap, ok := objMap["DependingSlides"]; ok {
		if valDependingSlidesCap != nil {
			var valueForDependingSlides []ResourceUriElement
			err = json.Unmarshal(*valDependingSlidesCap, &valueForDependingSlides)
			if err != nil {
				return err
			}
			valueForIDependingSlides := make([]IResourceUriElement, len(valueForDependingSlides))
			for i, v := range valueForDependingSlides {
				valueForIDependingSlides[i] = IResourceUriElement(&v)
			}
			this.DependingSlides = valueForIDependingSlides
		}
	}

    return nil
}
