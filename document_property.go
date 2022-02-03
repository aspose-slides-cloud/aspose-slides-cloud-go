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

// Document property.
type IDocumentProperty interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// Value.
	getValue() string
	setValue(newValue string)

	// True for builtin property.
	getBuiltIn() bool
	setBuiltIn(newValue bool)
}

type DocumentProperty struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// Value.
	Value string `json:"Value,omitempty"`

	// True for builtin property.
	BuiltIn bool `json:"BuiltIn"`
}

func NewDocumentProperty() *DocumentProperty {
	instance := new(DocumentProperty)
	return instance
}

func (this *DocumentProperty) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *DocumentProperty) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *DocumentProperty) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *DocumentProperty) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *DocumentProperty) getName() string {
	return this.Name
}

func (this *DocumentProperty) setName(newValue string) {
	this.Name = newValue
}
func (this *DocumentProperty) getValue() string {
	return this.Value
}

func (this *DocumentProperty) setValue(newValue string) {
	this.Value = newValue
}
func (this *DocumentProperty) getBuiltIn() bool {
	return this.BuiltIn
}

func (this *DocumentProperty) setBuiltIn(newValue bool) {
	this.BuiltIn = newValue
}

func (this *DocumentProperty) UnmarshalJSON(b []byte) error {
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
	
	if valValue, ok := objMap["value"]; ok {
		if valValue != nil {
			var valueForValue string
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	if valValueCap, ok := objMap["Value"]; ok {
		if valValueCap != nil {
			var valueForValue string
			err = json.Unmarshal(*valValueCap, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	
	if valBuiltIn, ok := objMap["builtIn"]; ok {
		if valBuiltIn != nil {
			var valueForBuiltIn bool
			err = json.Unmarshal(*valBuiltIn, &valueForBuiltIn)
			if err != nil {
				return err
			}
			this.BuiltIn = valueForBuiltIn
		}
	}
	if valBuiltInCap, ok := objMap["BuiltIn"]; ok {
		if valBuiltInCap != nil {
			var valueForBuiltIn bool
			err = json.Unmarshal(*valBuiltInCap, &valueForBuiltIn)
			if err != nil {
				return err
			}
			this.BuiltIn = valueForBuiltIn
		}
	}

	return nil
}
