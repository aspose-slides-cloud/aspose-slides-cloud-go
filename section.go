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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Name.
	getName() string
	setName(newValue string)

	// One-based index of slide with which the section starts.
	getFirstSlideIndex() int32
	setFirstSlideIndex(newValue int32)

	// Links to the shapes contained in the section.
	getSlideList() []IResourceUri
	setSlideList(newValue []IResourceUri)
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

func (this *Section) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Section) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Section) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Section) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Section) getName() string {
	return this.Name
}

func (this *Section) setName(newValue string) {
	this.Name = newValue
}
func (this *Section) getFirstSlideIndex() int32 {
	return this.FirstSlideIndex
}

func (this *Section) setFirstSlideIndex(newValue int32) {
	this.FirstSlideIndex = newValue
}
func (this *Section) getSlideList() []IResourceUri {
	return this.SlideList
}

func (this *Section) setSlideList(newValue []IResourceUri) {
	this.SlideList = newValue
}

func (this *Section) UnmarshalJSON(b []byte) error {
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
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []ResourceUri
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				valueForIAlternateLinks[i] = IResourceUri(&v)
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
	
	if valFirstSlideIndex, ok := objMap["firstSlideIndex"]; ok {
		if valFirstSlideIndex != nil {
			var valueForFirstSlideIndex int32
			err = json.Unmarshal(*valFirstSlideIndex, &valueForFirstSlideIndex)
			if err != nil {
				return err
			}
			this.FirstSlideIndex = valueForFirstSlideIndex
		}
	}
	if valFirstSlideIndexCap, ok := objMap["FirstSlideIndex"]; ok {
		if valFirstSlideIndexCap != nil {
			var valueForFirstSlideIndex int32
			err = json.Unmarshal(*valFirstSlideIndexCap, &valueForFirstSlideIndex)
			if err != nil {
				return err
			}
			this.FirstSlideIndex = valueForFirstSlideIndex
		}
	}
	
	if valSlideList, ok := objMap["slideList"]; ok {
		if valSlideList != nil {
			var valueForSlideList []ResourceUri
			err = json.Unmarshal(*valSlideList, &valueForSlideList)
			if err != nil {
				return err
			}
			valueForISlideList := make([]IResourceUri, len(valueForSlideList))
			for i, v := range valueForSlideList {
				valueForISlideList[i] = IResourceUri(&v)
			}
			this.SlideList = valueForISlideList
		}
	}
	if valSlideListCap, ok := objMap["SlideList"]; ok {
		if valSlideListCap != nil {
			var valueForSlideList []ResourceUri
			err = json.Unmarshal(*valSlideListCap, &valueForSlideList)
			if err != nil {
				return err
			}
			valueForISlideList := make([]IResourceUri, len(valueForSlideList))
			for i, v := range valueForSlideList {
				valueForISlideList[i] = IResourceUri(&v)
			}
			this.SlideList = valueForISlideList
		}
	}

    return nil
}
