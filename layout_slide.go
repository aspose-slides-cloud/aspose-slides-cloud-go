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
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// Layout slide type.
	getType() string
	setType(newValue string)

	// Master slide link.
	getMasterSlide() IResourceUri
	setMasterSlide(newValue IResourceUri)

	// List of depending slides.
	getDependingSlides() []IResourceUri
	setDependingSlides(newValue []IResourceUri)
}

type LayoutSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// Layout slide type.
	Type_ string `json:"Type"`

	// Master slide link.
	MasterSlide IResourceUri `json:"MasterSlide,omitempty"`

	// List of depending slides.
	DependingSlides []IResourceUri `json:"DependingSlides,omitempty"`
}

func NewLayoutSlide() *LayoutSlide {
	instance := new(LayoutSlide)
	instance.Type_ = "Title"
	return instance
}

func (this *LayoutSlide) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *LayoutSlide) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *LayoutSlide) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *LayoutSlide) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *LayoutSlide) getName() string {
	return this.Name
}

func (this *LayoutSlide) setName(newValue string) {
	this.Name = newValue
}
func (this *LayoutSlide) getType() string {
	return this.Type_
}

func (this *LayoutSlide) setType(newValue string) {
	this.Type_ = newValue
}
func (this *LayoutSlide) getMasterSlide() IResourceUri {
	return this.MasterSlide
}

func (this *LayoutSlide) setMasterSlide(newValue IResourceUri) {
	this.MasterSlide = newValue
}
func (this *LayoutSlide) getDependingSlides() []IResourceUri {
	return this.DependingSlides
}

func (this *LayoutSlide) setDependingSlides(newValue []IResourceUri) {
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
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
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
			var valueForMasterSlide ResourceUri
			err = json.Unmarshal(*valMasterSlide, &valueForMasterSlide)
			if err != nil {
				return err
			}
			this.MasterSlide = &valueForMasterSlide
		}
	}
	if valMasterSlideCap, ok := objMap["MasterSlide"]; ok {
		if valMasterSlideCap != nil {
			var valueForMasterSlide ResourceUri
			err = json.Unmarshal(*valMasterSlideCap, &valueForMasterSlide)
			if err != nil {
				return err
			}
			this.MasterSlide = &valueForMasterSlide
		}
	}
	
	if valDependingSlides, ok := objMap["dependingSlides"]; ok {
		if valDependingSlides != nil {
			var valueForDependingSlides []json.RawMessage
			err = json.Unmarshal(*valDependingSlides, &valueForDependingSlides)
			if err != nil {
				return err
			}
			valueForIDependingSlides := make([]IResourceUri, len(valueForDependingSlides))
			for i, v := range valueForDependingSlides {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIDependingSlides[i] = vObject.(IResourceUri)
				}
			}
			this.DependingSlides = valueForIDependingSlides
		}
	}
	if valDependingSlidesCap, ok := objMap["DependingSlides"]; ok {
		if valDependingSlidesCap != nil {
			var valueForDependingSlides []json.RawMessage
			err = json.Unmarshal(*valDependingSlidesCap, &valueForDependingSlides)
			if err != nil {
				return err
			}
			valueForIDependingSlides := make([]IResourceUri, len(valueForDependingSlides))
			for i, v := range valueForDependingSlides {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIDependingSlides[i] = vObject.(IResourceUri)
				}
			}
			this.DependingSlides = valueForIDependingSlides
		}
	}

	return nil
}
