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


type IDocumentProperty interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getName() string
	setName(newValue string)

	getValue() string
	setValue(newValue string)

	getBuiltIn() bool
	setBuiltIn(newValue bool)
}

type DocumentProperty struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	Name string `json:"Name,omitempty"`

	Value string `json:"Value,omitempty"`

	BuiltIn bool `json:"BuiltIn"`
}

func (this DocumentProperty) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this DocumentProperty) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this DocumentProperty) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this DocumentProperty) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this DocumentProperty) getName() string {
	return this.Name
}

func (this DocumentProperty) setName(newValue string) {
	this.Name = newValue
}
func (this DocumentProperty) getValue() string {
	return this.Value
}

func (this DocumentProperty) setValue(newValue string) {
	this.Value = newValue
}
func (this DocumentProperty) getBuiltIn() bool {
	return this.BuiltIn
}

func (this DocumentProperty) setBuiltIn(newValue bool) {
	this.BuiltIn = newValue
}

func (this *DocumentProperty) UnmarshalJSON(b []byte) error {
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

	if valValue, ok := objMap["Value"]; ok {
		if valValue != nil {
			var valueForValue string
			err = json.Unmarshal(*valValue, &valueForValue)
			if err != nil {
				return err
			}
			this.Value = valueForValue
		}
	}

	if valBuiltIn, ok := objMap["BuiltIn"]; ok {
		if valBuiltIn != nil {
			var valueForBuiltIn bool
			err = json.Unmarshal(*valBuiltIn, &valueForBuiltIn)
			if err != nil {
				return err
			}
			this.BuiltIn = valueForBuiltIn
		}
	}

    return nil
}
