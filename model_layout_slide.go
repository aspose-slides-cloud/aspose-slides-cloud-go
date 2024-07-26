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
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Name.
	GetName() string
	SetName(newValue string)

	// Layout slide type.
	GetType() string
	SetType(newValue string)

	// Master slide link.
	GetMasterSlide() IResourceUri
	SetMasterSlide(newValue IResourceUri)

	// List of depending slides.
	GetDependingSlides() []IResourceUri
	SetDependingSlides(newValue []IResourceUri)
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

func (this *LayoutSlide) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *LayoutSlide) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *LayoutSlide) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *LayoutSlide) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *LayoutSlide) GetName() string {
	return this.Name
}

func (this *LayoutSlide) SetName(newValue string) {
	this.Name = newValue
}
func (this *LayoutSlide) GetType() string {
	return this.Type_
}

func (this *LayoutSlide) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *LayoutSlide) GetMasterSlide() IResourceUri {
	return this.MasterSlide
}

func (this *LayoutSlide) SetMasterSlide(newValue IResourceUri) {
	this.MasterSlide = newValue
}
func (this *LayoutSlide) GetDependingSlides() []IResourceUri {
	return this.DependingSlides
}

func (this *LayoutSlide) SetDependingSlides(newValue []IResourceUri) {
	this.DependingSlides = newValue
}

func (this *LayoutSlide) UnmarshalJSON(b []byte) error {
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
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	this.Type_ = "Title"
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valMasterSlide, ok := GetMapValue(objMap, "masterSlide"); ok {
		if valMasterSlide != nil {
			var valueForMasterSlide ResourceUri
			err = json.Unmarshal(*valMasterSlide, &valueForMasterSlide)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valMasterSlide)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMasterSlide, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.MasterSlide = vInterfaceObject
			}
		}
	}
	
	if valDependingSlides, ok := GetMapValue(objMap, "dependingSlides"); ok {
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

	return nil
}
