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
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Name.
	GetName() string
	SetName(newValue string)

	// Value.
	GetValue() string
	SetValue(newValue string)

	// True for builtin property.
	GetBuiltIn() *bool
	SetBuiltIn(newValue *bool)
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
	BuiltIn *bool `json:"BuiltIn"`
}

func NewDocumentProperty() *DocumentProperty {
	instance := new(DocumentProperty)
	return instance
}

func (this *DocumentProperty) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *DocumentProperty) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *DocumentProperty) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *DocumentProperty) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *DocumentProperty) GetName() string {
	return this.Name
}

func (this *DocumentProperty) SetName(newValue string) {
	this.Name = newValue
}
func (this *DocumentProperty) GetValue() string {
	return this.Value
}

func (this *DocumentProperty) SetValue(newValue string) {
	this.Value = newValue
}
func (this *DocumentProperty) GetBuiltIn() *bool {
	return this.BuiltIn
}

func (this *DocumentProperty) SetBuiltIn(newValue *bool) {
	this.BuiltIn = newValue
}

func (this *DocumentProperty) UnmarshalJSON(b []byte) error {
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
	
	if valValue, ok := GetMapValue(objMap, "value"); ok {
		if valValue != nil {
			var valueForValue string
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}
	
	if valBuiltIn, ok := GetMapValue(objMap, "builtIn"); ok {
		if valBuiltIn != nil {
			var valueForBuiltIn *bool
			err = json.Unmarshal(*valBuiltIn, &valueForBuiltIn)
			if err != nil {
				return err
			}
			this.BuiltIn = valueForBuiltIn
		}
	}

	return nil
}
