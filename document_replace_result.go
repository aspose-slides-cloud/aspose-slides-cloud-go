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

// Represents document replace result DTO.
type IDocumentReplaceResult interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Link to Document properties.
	getDocumentProperties() IResourceUriElement
	setDocumentProperties(newValue IResourceUriElement)

	// Link to slides collection.
	getSlides() IResourceUriElement
	setSlides(newValue IResourceUriElement)

	// Link to images collection.
	getImages() IResourceUriElement
	setImages(newValue IResourceUriElement)

	// Link to layout slides collection.
	getLayoutSlides() IResourceUriElement
	setLayoutSlides(newValue IResourceUriElement)

	// Link to master slides collection.
	getMasterSlides() IResourceUriElement
	setMasterSlides(newValue IResourceUriElement)

	// Gets or sets the number of matches 
	getMatches() int32
	setMatches(newValue int32)
}

type DocumentReplaceResult struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Link to Document properties.
	DocumentProperties IResourceUriElement `json:"DocumentProperties,omitempty"`

	// Link to slides collection.
	Slides IResourceUriElement `json:"Slides,omitempty"`

	// Link to images collection.
	Images IResourceUriElement `json:"Images,omitempty"`

	// Link to layout slides collection.
	LayoutSlides IResourceUriElement `json:"LayoutSlides,omitempty"`

	// Link to master slides collection.
	MasterSlides IResourceUriElement `json:"MasterSlides,omitempty"`

	// Gets or sets the number of matches 
	Matches int32 `json:"Matches"`
}

func (this DocumentReplaceResult) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this DocumentReplaceResult) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this DocumentReplaceResult) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this DocumentReplaceResult) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this DocumentReplaceResult) getDocumentProperties() IResourceUriElement {
	return this.DocumentProperties
}

func (this DocumentReplaceResult) setDocumentProperties(newValue IResourceUriElement) {
	this.DocumentProperties = newValue
}
func (this DocumentReplaceResult) getSlides() IResourceUriElement {
	return this.Slides
}

func (this DocumentReplaceResult) setSlides(newValue IResourceUriElement) {
	this.Slides = newValue
}
func (this DocumentReplaceResult) getImages() IResourceUriElement {
	return this.Images
}

func (this DocumentReplaceResult) setImages(newValue IResourceUriElement) {
	this.Images = newValue
}
func (this DocumentReplaceResult) getLayoutSlides() IResourceUriElement {
	return this.LayoutSlides
}

func (this DocumentReplaceResult) setLayoutSlides(newValue IResourceUriElement) {
	this.LayoutSlides = newValue
}
func (this DocumentReplaceResult) getMasterSlides() IResourceUriElement {
	return this.MasterSlides
}

func (this DocumentReplaceResult) setMasterSlides(newValue IResourceUriElement) {
	this.MasterSlides = newValue
}
func (this DocumentReplaceResult) getMatches() int32 {
	return this.Matches
}

func (this DocumentReplaceResult) setMatches(newValue int32) {
	this.Matches = newValue
}

func (this *DocumentReplaceResult) UnmarshalJSON(b []byte) error {
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
	
	if valDocumentProperties, ok := objMap["DocumentProperties"]; ok {
		if valDocumentProperties != nil {
			var valueForDocumentProperties ResourceUriElement
			err = json.Unmarshal(*valDocumentProperties, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			this.DocumentProperties = valueForDocumentProperties
		}
	}
	
	if valSlides, ok := objMap["Slides"]; ok {
		if valSlides != nil {
			var valueForSlides ResourceUriElement
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = valueForSlides
		}
	}
	
	if valImages, ok := objMap["Images"]; ok {
		if valImages != nil {
			var valueForImages ResourceUriElement
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = valueForImages
		}
	}
	
	if valLayoutSlides, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides ResourceUriElement
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = valueForLayoutSlides
		}
	}
	
	if valMasterSlides, ok := objMap["MasterSlides"]; ok {
		if valMasterSlides != nil {
			var valueForMasterSlides ResourceUriElement
			err = json.Unmarshal(*valMasterSlides, &valueForMasterSlides)
			if err != nil {
				return err
			}
			this.MasterSlides = valueForMasterSlides
		}
	}
	
	if valMatches, ok := objMap["Matches"]; ok {
		if valMatches != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatches, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}

    return nil
}
