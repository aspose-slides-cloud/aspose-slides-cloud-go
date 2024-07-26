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

// Represents document DTO.
type IDocument interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Link to Document properties.
	GetDocumentProperties() IResourceUri
	SetDocumentProperties(newValue IResourceUri)

	// Link to Document properties.
	GetViewProperties() IResourceUri
	SetViewProperties(newValue IResourceUri)

	// Link to slides collection.
	GetSlides() IResourceUri
	SetSlides(newValue IResourceUri)

	// Link to images collection.
	GetImages() IResourceUri
	SetImages(newValue IResourceUri)

	// Link to layout slides collection.
	GetLayoutSlides() IResourceUri
	SetLayoutSlides(newValue IResourceUri)

	// Link to master slides collection.
	GetMasterSlides() IResourceUri
	SetMasterSlides(newValue IResourceUri)
}

type Document struct {

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
}

func NewDocument() *Document {
	instance := new(Document)
	return instance
}

func (this *Document) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Document) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Document) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Document) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Document) GetDocumentProperties() IResourceUri {
	return this.DocumentProperties
}

func (this *Document) SetDocumentProperties(newValue IResourceUri) {
	this.DocumentProperties = newValue
}
func (this *Document) GetViewProperties() IResourceUri {
	return this.ViewProperties
}

func (this *Document) SetViewProperties(newValue IResourceUri) {
	this.ViewProperties = newValue
}
func (this *Document) GetSlides() IResourceUri {
	return this.Slides
}

func (this *Document) SetSlides(newValue IResourceUri) {
	this.Slides = newValue
}
func (this *Document) GetImages() IResourceUri {
	return this.Images
}

func (this *Document) SetImages(newValue IResourceUri) {
	this.Images = newValue
}
func (this *Document) GetLayoutSlides() IResourceUri {
	return this.LayoutSlides
}

func (this *Document) SetLayoutSlides(newValue IResourceUri) {
	this.LayoutSlides = newValue
}
func (this *Document) GetMasterSlides() IResourceUri {
	return this.MasterSlides
}

func (this *Document) SetMasterSlides(newValue IResourceUri) {
	this.MasterSlides = newValue
}

func (this *Document) UnmarshalJSON(b []byte) error {
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
	
	if valDocumentProperties, ok := GetMapValue(objMap, "documentProperties"); ok {
		if valDocumentProperties != nil {
			var valueForDocumentProperties ResourceUri
			err = json.Unmarshal(*valDocumentProperties, &valueForDocumentProperties)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valDocumentProperties)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDocumentProperties, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.DocumentProperties = vInterfaceObject
			}
		}
	}
	
	if valViewProperties, ok := GetMapValue(objMap, "viewProperties"); ok {
		if valViewProperties != nil {
			var valueForViewProperties ResourceUri
			err = json.Unmarshal(*valViewProperties, &valueForViewProperties)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valViewProperties)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valViewProperties, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.ViewProperties = vInterfaceObject
			}
		}
	}
	
	if valSlides, ok := GetMapValue(objMap, "slides"); ok {
		if valSlides != nil {
			var valueForSlides ResourceUri
			err = json.Unmarshal(*valSlides, &valueForSlides)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSlides)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlides, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Slides = vInterfaceObject
			}
		}
	}
	
	if valImages, ok := GetMapValue(objMap, "images"); ok {
		if valImages != nil {
			var valueForImages ResourceUri
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valImages)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valImages, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Images = vInterfaceObject
			}
		}
	}
	
	if valLayoutSlides, ok := GetMapValue(objMap, "layoutSlides"); ok {
		if valLayoutSlides != nil {
			var valueForLayoutSlides ResourceUri
			err = json.Unmarshal(*valLayoutSlides, &valueForLayoutSlides)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valLayoutSlides)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLayoutSlides, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.LayoutSlides = vInterfaceObject
			}
		}
	}
	
	if valMasterSlides, ok := GetMapValue(objMap, "masterSlides"); ok {
		if valMasterSlides != nil {
			var valueForMasterSlides ResourceUri
			err = json.Unmarshal(*valMasterSlides, &valueForMasterSlides)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valMasterSlides)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMasterSlides, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.MasterSlides = vInterfaceObject
			}
		}
	}

	return nil
}
