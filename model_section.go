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

// Presentation section.
type ISection interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Name.
	GetName() string
	SetName(newValue string)

	// One-based index of slide with which the section starts.
	GetFirstSlideIndex() int32
	SetFirstSlideIndex(newValue int32)

	// Links to the shapes contained in the section.
	GetSlideList() []IResourceUri
	SetSlideList(newValue []IResourceUri)
}

type Section struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// One-based index of slide with which the section starts.
	FirstSlideIndex int32 `json:"FirstSlideIndex"`

	// Links to the shapes contained in the section.
	SlideList []IResourceUri `json:"SlideList,omitempty"`
}

func NewSection() *Section {
	instance := new(Section)
	return instance
}

func (this *Section) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Section) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Section) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Section) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Section) GetName() string {
	return this.Name
}

func (this *Section) SetName(newValue string) {
	this.Name = newValue
}
func (this *Section) GetFirstSlideIndex() int32 {
	return this.FirstSlideIndex
}

func (this *Section) SetFirstSlideIndex(newValue int32) {
	this.FirstSlideIndex = newValue
}
func (this *Section) GetSlideList() []IResourceUri {
	return this.SlideList
}

func (this *Section) SetSlideList(newValue []IResourceUri) {
	this.SlideList = newValue
}

func (this *Section) UnmarshalJSON(b []byte) error {
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
	
	if valFirstSlideIndex, ok := GetMapValue(objMap, "firstSlideIndex"); ok {
		if valFirstSlideIndex != nil {
			var valueForFirstSlideIndex int32
			err = json.Unmarshal(*valFirstSlideIndex, &valueForFirstSlideIndex)
			if err != nil {
				return err
			}
			this.FirstSlideIndex = valueForFirstSlideIndex
		}
	}
	
	if valSlideList, ok := GetMapValue(objMap, "slideList"); ok {
		if valSlideList != nil {
			var valueForSlideList []json.RawMessage
			err = json.Unmarshal(*valSlideList, &valueForSlideList)
			if err != nil {
				return err
			}
			valueForISlideList := make([]IResourceUri, len(valueForSlideList))
			for i, v := range valueForSlideList {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForISlideList[i] = vObject.(IResourceUri)
				}
			}
			this.SlideList = valueForISlideList
		}
	}

	return nil
}
