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


type IMasterSlide interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getName() string
	setName(newValue string)

	getLayoutSlides() []ResourceUriElement
	setLayoutSlides(newValue []ResourceUriElement)

	getDependingSlides() []ResourceUriElement
	setDependingSlides(newValue []ResourceUriElement)
}

type MasterSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	Name string `json:"Name,omitempty"`

	LayoutSlides []ResourceUriElement `json:"LayoutSlides,omitempty"`

	DependingSlides []ResourceUriElement `json:"DependingSlides,omitempty"`
}

func (this MasterSlide) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this MasterSlide) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this MasterSlide) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this MasterSlide) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this MasterSlide) getName() string {
	return this.Name
}

func (this MasterSlide) setName(newValue string) {
	this.Name = newValue
}
func (this MasterSlide) getLayoutSlides() []ResourceUriElement {
	return this.LayoutSlides
}

func (this MasterSlide) setLayoutSlides(newValue []ResourceUriElement) {
	this.LayoutSlides = newValue
}
func (this MasterSlide) getDependingSlides() []ResourceUriElement {
	return this.DependingSlides
}

func (this MasterSlide) setDependingSlides(newValue []ResourceUriElement) {
	this.DependingSlides = newValue
}

func (this *MasterSlide) UnmarshalJSON(b []byte) error {
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
	
	if valName, ok := objMap["Name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valLayoutSlides, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides []ResourceUriElement
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = valueForLayoutSlides
		}
	}
	
	if valDependingSlides, ok := objMap["DependingSlides"]; ok {
		if valDependingSlides != nil {
			var valueForDependingSlides []ResourceUriElement
			err = json.Unmarshal(*valDependingSlides, &valueForDependingSlides)
			if err != nil {
				return err
			}
			this.DependingSlides = valueForDependingSlides
		}
	}

    return nil
}
