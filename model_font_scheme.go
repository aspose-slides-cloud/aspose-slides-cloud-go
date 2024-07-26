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

// Represents font scheme
type IFontScheme interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets fonts collection for a \"heading\" part of the slide.
	GetMajor() IFontSet
	SetMajor(newValue IFontSet)

	// Gets or sets  the fonts collection for a \"body\" part of the slide.
	GetMinor() IFontSet
	SetMinor(newValue IFontSet)

	// Gets or sets the name.
	GetName() string
	SetName(newValue string)
}

type FontScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets fonts collection for a \"heading\" part of the slide.
	Major IFontSet `json:"Major,omitempty"`

	// Gets or sets  the fonts collection for a \"body\" part of the slide.
	Minor IFontSet `json:"Minor,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`
}

func NewFontScheme() *FontScheme {
	instance := new(FontScheme)
	return instance
}

func (this *FontScheme) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *FontScheme) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *FontScheme) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *FontScheme) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *FontScheme) GetMajor() IFontSet {
	return this.Major
}

func (this *FontScheme) SetMajor(newValue IFontSet) {
	this.Major = newValue
}
func (this *FontScheme) GetMinor() IFontSet {
	return this.Minor
}

func (this *FontScheme) SetMinor(newValue IFontSet) {
	this.Minor = newValue
}
func (this *FontScheme) GetName() string {
	return this.Name
}

func (this *FontScheme) SetName(newValue string) {
	this.Name = newValue
}

func (this *FontScheme) UnmarshalJSON(b []byte) error {
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
	
	if valMajor, ok := GetMapValue(objMap, "major"); ok {
		if valMajor != nil {
			var valueForMajor FontSet
			err = json.Unmarshal(*valMajor, &valueForMajor)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FontSet", *valMajor)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMajor, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFontSet)
			if ok {
				this.Major = vInterfaceObject
			}
		}
	}
	
	if valMinor, ok := GetMapValue(objMap, "minor"); ok {
		if valMinor != nil {
			var valueForMinor FontSet
			err = json.Unmarshal(*valMinor, &valueForMinor)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FontSet", *valMinor)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valMinor, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFontSet)
			if ok {
				this.Minor = vInterfaceObject
			}
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

	return nil
}
