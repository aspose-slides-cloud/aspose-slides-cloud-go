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

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Link to Document properties.
	getDocumentProperties() IResourceUri
	setDocumentProperties(newValue IResourceUri)

	// Link to Document properties.
	getViewProperties() IResourceUri
	setViewProperties(newValue IResourceUri)

	// Link to slides collection.
	getSlides() IResourceUri
	setSlides(newValue IResourceUri)

	// Link to images collection.
	getImages() IResourceUri
	setImages(newValue IResourceUri)

	// Link to layout slides collection.
	getLayoutSlides() IResourceUri
	setLayoutSlides(newValue IResourceUri)

	// Link to master slides collection.
	getMasterSlides() IResourceUri
	setMasterSlides(newValue IResourceUri)

	// Gets or sets the number of matches 
	getMatches() int32
	setMatches(newValue int32)
}

type DocumentReplaceResult struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Link to Document properties.
	DocumentProperties IResourceUri `json:"DocumentProperties,omitempty"`

	// Link to Document properties.
	ViewProperties IResourceUri `json:"ViewProperties,omitempty"`

	// Link to slides collection.
	Slides IResourceUri `json:"Slides,omitempty"`

	// Link to images collection.
	Images IResourceUri `json:"Images,omitempty"`

	// Link to layout slides collection.
	LayoutSlides IResourceUri `json:"LayoutSlides,omitempty"`

	// Link to master slides collection.
	MasterSlides IResourceUri `json:"MasterSlides,omitempty"`

	// Gets or sets the number of matches 
	Matches int32 `json:"Matches"`
}

func NewDocumentReplaceResult() *DocumentReplaceResult {
	instance := new(DocumentReplaceResult)
	return instance
}

func (this *DocumentReplaceResult) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *DocumentReplaceResult) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *DocumentReplaceResult) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *DocumentReplaceResult) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *DocumentReplaceResult) getDocumentProperties() IResourceUri {
	return this.DocumentProperties
}

func (this *DocumentReplaceResult) setDocumentProperties(newValue IResourceUri) {
	this.DocumentProperties = newValue
}
func (this *DocumentReplaceResult) getViewProperties() IResourceUri {
	return this.ViewProperties
}

func (this *DocumentReplaceResult) setViewProperties(newValue IResourceUri) {
	this.ViewProperties = newValue
}
func (this *DocumentReplaceResult) getSlides() IResourceUri {
	return this.Slides
}

func (this *DocumentReplaceResult) setSlides(newValue IResourceUri) {
	this.Slides = newValue
}
func (this *DocumentReplaceResult) getImages() IResourceUri {
	return this.Images
}

func (this *DocumentReplaceResult) setImages(newValue IResourceUri) {
	this.Images = newValue
}
func (this *DocumentReplaceResult) getLayoutSlides() IResourceUri {
	return this.LayoutSlides
}

func (this *DocumentReplaceResult) setLayoutSlides(newValue IResourceUri) {
	this.LayoutSlides = newValue
}
func (this *DocumentReplaceResult) getMasterSlides() IResourceUri {
	return this.MasterSlides
}

func (this *DocumentReplaceResult) setMasterSlides(newValue IResourceUri) {
	this.MasterSlides = newValue
}
func (this *DocumentReplaceResult) getMatches() int32 {
	return this.Matches
}

func (this *DocumentReplaceResult) setMatches(newValue int32) {
	this.Matches = newValue
}

func (this *DocumentReplaceResult) UnmarshalJSON(b []byte) error {
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
	
	if valDocumentProperties, ok := objMap["documentProperties"]; ok {
		if valDocumentProperties != nil {
			var valueForDocumentProperties ResourceUri
			err = json.Unmarshal(*valDocumentProperties, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			this.DocumentProperties = &valueForDocumentProperties
		}
	}
	if valDocumentPropertiesCap, ok := objMap["DocumentProperties"]; ok {
		if valDocumentPropertiesCap != nil {
			var valueForDocumentProperties ResourceUri
			err = json.Unmarshal(*valDocumentPropertiesCap, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			this.DocumentProperties = &valueForDocumentProperties
		}
	}
	
	if valViewProperties, ok := objMap["viewProperties"]; ok {
		if valViewProperties != nil {
			var valueForViewProperties ResourceUri
			err = json.Unmarshal(*valViewProperties, &valueForViewProperties)
			if err != nil {
				return err
			}
			this.ViewProperties = &valueForViewProperties
		}
	}
	if valViewPropertiesCap, ok := objMap["ViewProperties"]; ok {
		if valViewPropertiesCap != nil {
			var valueForViewProperties ResourceUri
			err = json.Unmarshal(*valViewPropertiesCap, &valueForViewProperties)
			if err != nil {
				return err
			}
			this.ViewProperties = &valueForViewProperties
		}
	}
	
	if valSlides, ok := objMap["slides"]; ok {
		if valSlides != nil {
			var valueForSlides ResourceUri
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = &valueForSlides
		}
	}
	if valSlidesCap, ok := objMap["Slides"]; ok {
		if valSlidesCap != nil {
			var valueForSlides ResourceUri
			err = json.Unmarshal(*valSlidesCap, &valueForSlides)
			if err != nil {
				return err
			}
			this.Slides = &valueForSlides
		}
	}
	
	if valImages, ok := objMap["images"]; ok {
		if valImages != nil {
			var valueForImages ResourceUri
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = &valueForImages
		}
	}
	if valImagesCap, ok := objMap["Images"]; ok {
		if valImagesCap != nil {
			var valueForImages ResourceUri
			err = json.Unmarshal(*valImagesCap, &valueForImages)
			if err != nil {
				return err
			}
			this.Images = &valueForImages
		}
	}
	
	if valLayoutSlides, ok := objMap["layoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides ResourceUri
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = &valueForLayoutSlides
		}
	}
	if valLayoutSlidesCap, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlidesCap != nil {
			var valueForLayoutSlides ResourceUri
			err = json.Unmarshal(*valLayoutSlidesCap, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			this.LayoutSlides = &valueForLayoutSlides
		}
	}
	
	if valMasterSlides, ok := objMap["masterSlides"]; ok {
		if valMasterSlides != nil {
			var valueForMasterSlides ResourceUri
			err = json.Unmarshal(*valMasterSlides, &valueForMasterSlides)
			if err != nil {
				return err
			}
			this.MasterSlides = &valueForMasterSlides
		}
	}
	if valMasterSlidesCap, ok := objMap["MasterSlides"]; ok {
		if valMasterSlidesCap != nil {
			var valueForMasterSlides ResourceUri
			err = json.Unmarshal(*valMasterSlidesCap, &valueForMasterSlides)
			if err != nil {
				return err
			}
			this.MasterSlides = &valueForMasterSlides
		}
	}
	
	if valMatches, ok := objMap["matches"]; ok {
		if valMatches != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatches, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}
	if valMatchesCap, ok := objMap["Matches"]; ok {
		if valMatchesCap != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatchesCap, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}

    return nil
}
