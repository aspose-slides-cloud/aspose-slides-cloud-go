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
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Name.
	GetName() string
	SetName(newValue string)

	// List of layout slide links.
	GetLayoutSlides() []IResourceUri
	SetLayoutSlides(newValue []IResourceUri)

	// List of depending slide links.
	GetDependingSlides() []IResourceUri
	SetDependingSlides(newValue []IResourceUri)
}

type MasterSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// List of layout slide links.
	LayoutSlides []IResourceUri `json:"LayoutSlides,omitempty"`

	// List of depending slide links.
	DependingSlides []IResourceUri `json:"DependingSlides,omitempty"`
}

func NewMasterSlide() *MasterSlide {
	instance := new(MasterSlide)
	return instance
}

func (this *MasterSlide) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *MasterSlide) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *MasterSlide) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *MasterSlide) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *MasterSlide) GetName() string {
	return this.Name
}

func (this *MasterSlide) SetName(newValue string) {
	this.Name = newValue
}
func (this *MasterSlide) GetLayoutSlides() []IResourceUri {
	return this.LayoutSlides
}

func (this *MasterSlide) SetLayoutSlides(newValue []IResourceUri) {
	this.LayoutSlides = newValue
}
func (this *MasterSlide) GetDependingSlides() []IResourceUri {
	return this.DependingSlides
}

func (this *MasterSlide) SetDependingSlides(newValue []IResourceUri) {
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
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
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
	
	if valLayoutSlides, ok := objMap["layoutSlides"]; ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides []json.RawMessage
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			valueForILayoutSlides := make([]IResourceUri, len(valueForLayoutSlides))
			for i, v := range valueForLayoutSlides {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForILayoutSlides[i] = vObject.(IResourceUri)
				}
			}
			this.LayoutSlides = valueForILayoutSlides
		}
	}
	if valLayoutSlidesCap, ok := objMap["LayoutSlides"]; ok {
		if valLayoutSlidesCap != nil {
			var valueForLayoutSlides []json.RawMessage
			err = json.Unmarshal(*valLayoutSlidesCap, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			valueForILayoutSlides := make([]IResourceUri, len(valueForLayoutSlides))
			for i, v := range valueForLayoutSlides {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForILayoutSlides[i] = vObject.(IResourceUri)
				}
			}
			this.LayoutSlides = valueForILayoutSlides
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
