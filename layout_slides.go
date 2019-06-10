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


type ILayoutSlides interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getSlideList() []ResourceUriElement
	setSlideList(newValue []ResourceUriElement)
}

type LayoutSlides struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	SlideList []ResourceUriElement `json:"SlideList,omitempty"`
}

func (this LayoutSlides) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this LayoutSlides) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this LayoutSlides) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this LayoutSlides) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this LayoutSlides) getSlideList() []ResourceUriElement {
	return this.SlideList
}

func (this LayoutSlides) setSlideList(newValue []ResourceUriElement) {
	this.SlideList = newValue
}

func (this *LayoutSlides) UnmarshalJSON(b []byte) error {
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

	if valSlideList, ok := objMap["SlideList"]; ok {
		if valSlideList != nil {
			var valueForSlideList []ResourceUriElement
			err = json.Unmarshal(*valSlideList, &valueForSlideList)
			if err != nil {
				return err
			}
			this.SlideList = valueForSlideList
		}
	}

    return nil
}
