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

// Represents list of Links to Paragraphs resources
type IParagraphList interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// A list of links that originate from this document.
	getLinks() []ResourceUri
	setLinks(newValue []ResourceUri)

	getParagraphLinks() []ResourceUriElement
	setParagraphLinks(newValue []ResourceUriElement)
}

type ParagraphList struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	ParagraphLinks []ResourceUriElement `json:"ParagraphLinks,omitempty"`
}

func (this ParagraphList) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this ParagraphList) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this ParagraphList) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this ParagraphList) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this ParagraphList) getLinks() []ResourceUri {
	return this.Links
}

func (this ParagraphList) setLinks(newValue []ResourceUri) {
	this.Links = newValue
}
func (this ParagraphList) getParagraphLinks() []ResourceUriElement {
	return this.ParagraphLinks
}

func (this ParagraphList) setParagraphLinks(newValue []ResourceUriElement) {
	this.ParagraphLinks = newValue
}

func (this *ParagraphList) UnmarshalJSON(b []byte) error {
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

	if valLinks, ok := objMap["Links"]; ok {
		if valLinks != nil {
			var valueForLinks []ResourceUri
			err = json.Unmarshal(*valLinks, &valueForLinks)
			if err != nil {
				return err
			}
			this.Links = valueForLinks
		}
	}

	if valParagraphLinks, ok := objMap["ParagraphLinks"]; ok {
		if valParagraphLinks != nil {
			var valueForParagraphLinks []ResourceUriElement
			err = json.Unmarshal(*valParagraphLinks, &valueForParagraphLinks)
			if err != nil {
				return err
			}
			this.ParagraphLinks = valueForParagraphLinks
		}
	}

    return nil
}
