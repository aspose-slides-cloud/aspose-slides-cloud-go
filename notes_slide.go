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

// Represents notes slide DTO.
type INotesSlide interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// A list of links that originate from this document.
	getLinks() []ResourceUri
	setLinks(newValue []ResourceUri)

	// Get or sets text of notes slide.
	getText() string
	setText(newValue string)

	// Get or sets the  link to list notes slide shapes.
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)
}

type NotesSlide struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	// Get or sets text of notes slide.
	Text string `json:"Text,omitempty"`

	// Get or sets the  link to list notes slide shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`
}

func (this NotesSlide) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this NotesSlide) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this NotesSlide) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this NotesSlide) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this NotesSlide) getLinks() []ResourceUri {
	return this.Links
}

func (this NotesSlide) setLinks(newValue []ResourceUri) {
	this.Links = newValue
}
func (this NotesSlide) getText() string {
	return this.Text
}

func (this NotesSlide) setText(newValue string) {
	this.Text = newValue
}
func (this NotesSlide) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this NotesSlide) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}

func (this *NotesSlide) UnmarshalJSON(b []byte) error {
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

	if valText, ok := objMap["Text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}

	if valShapes, ok := objMap["Shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}

    return nil
}
