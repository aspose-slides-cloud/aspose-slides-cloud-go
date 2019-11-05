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

// Layout slide.
type ILayoutSlide interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// Layout slide type.
	getType() string
	setType(newValue string)

	// Master slide link.
	getMasterSlide() IResourceUriElement
	setMasterSlide(newValue IResourceUriElement)

	// List of depending slides.
	getDependingSlides() []ResourceUriElement
	setDependingSlides(newValue []ResourceUriElement)
}

type LayoutSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// Layout slide type.
	Type_ string `json:"Type"`

	// Master slide link.
	MasterSlide IResourceUriElement `json:"MasterSlide,omitempty"`

	// List of depending slides.
	DependingSlides []ResourceUriElement `json:"DependingSlides,omitempty"`
}

func (this LayoutSlide) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this LayoutSlide) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this LayoutSlide) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this LayoutSlide) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this LayoutSlide) getName() string {
	return this.Name
}

func (this LayoutSlide) setName(newValue string) {
	this.Name = newValue
}
func (this LayoutSlide) getType() string {
	return this.Type_
}

func (this LayoutSlide) setType(newValue string) {
	this.Type_ = newValue
}
func (this LayoutSlide) getMasterSlide() IResourceUriElement {
	return this.MasterSlide
}

func (this LayoutSlide) setMasterSlide(newValue IResourceUriElement) {
	this.MasterSlide = newValue
}
func (this LayoutSlide) getDependingSlides() []ResourceUriElement {
	return this.DependingSlides
}

func (this LayoutSlide) setDependingSlides(newValue []ResourceUriElement) {
	this.DependingSlides = newValue
}

func (this *LayoutSlide) UnmarshalJSON(b []byte) error {
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
			this.SelfUri = valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			this.AlternateLinks = valueForAlternateLinks
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
	this.Type_ = "Title"
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
	
	if valMasterSlide, ok := objMap["masterSlide"]; ok {
		if valMasterSlide != nil {
			var valueForMasterSlide ResourceUriElement
			err = json.Unmarshal(*valMasterSlide, &valueForMasterSlide)
			if err != nil {
				return err
			}
			this.MasterSlide = valueForMasterSlide
		}
	}
	if valMasterSlideCap, ok := objMap["MasterSlide"]; ok {
		if valMasterSlideCap != nil {
			var valueForMasterSlide ResourceUriElement
			err = json.Unmarshal(*valMasterSlideCap, &valueForMasterSlide)
			if err != nil {
				return err
			}
			this.MasterSlide = valueForMasterSlide
		}
	}
	
	if valDependingSlides, ok := objMap["dependingSlides"]; ok {
		if valDependingSlides != nil {
			var valueForDependingSlides []ResourceUriElement
			err = json.Unmarshal(*valDependingSlides, &valueForDependingSlides)
			if err != nil {
				return err
			}
			this.DependingSlides = valueForDependingSlides
		}
	}
	if valDependingSlidesCap, ok := objMap["DependingSlides"]; ok {
		if valDependingSlidesCap != nil {
			var valueForDependingSlides []ResourceUriElement
			err = json.Unmarshal(*valDependingSlidesCap, &valueForDependingSlides)
			if err != nil {
				return err
			}
			this.DependingSlides = valueForDependingSlides
		}
	}

    return nil
}
