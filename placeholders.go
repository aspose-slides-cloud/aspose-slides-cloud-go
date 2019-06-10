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


type IPlaceholders interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getPlaceholderLinks() []ResourceUri
	setPlaceholderLinks(newValue []ResourceUri)
}

type Placeholders struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	PlaceholderLinks []ResourceUri `json:"PlaceholderLinks,omitempty"`
}

func (this Placeholders) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Placeholders) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Placeholders) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Placeholders) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Placeholders) getPlaceholderLinks() []ResourceUri {
	return this.PlaceholderLinks
}

func (this Placeholders) setPlaceholderLinks(newValue []ResourceUri) {
	this.PlaceholderLinks = newValue
}

func (this *Placeholders) UnmarshalJSON(b []byte) error {
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

	if valPlaceholderLinks, ok := objMap["PlaceholderLinks"]; ok {
		if valPlaceholderLinks != nil {
			var valueForPlaceholderLinks []ResourceUri
			err = json.Unmarshal(*valPlaceholderLinks, &valueForPlaceholderLinks)
			if err != nil {
				return err
			}
			this.PlaceholderLinks = valueForPlaceholderLinks
		}
	}

    return nil
}
