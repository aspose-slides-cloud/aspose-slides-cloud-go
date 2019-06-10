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

// Represents list of Links to Shapes resources
type IShapes interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getShapesLinks() []ResourceUriElement
	setShapesLinks(newValue []ResourceUriElement)
}

type Shapes struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	ShapesLinks []ResourceUriElement `json:"ShapesLinks,omitempty"`
}

func (this Shapes) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Shapes) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Shapes) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Shapes) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Shapes) getShapesLinks() []ResourceUriElement {
	return this.ShapesLinks
}

func (this Shapes) setShapesLinks(newValue []ResourceUriElement) {
	this.ShapesLinks = newValue
}

func (this *Shapes) UnmarshalJSON(b []byte) error {
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

	if valShapesLinks, ok := objMap["ShapesLinks"]; ok {
		if valShapesLinks != nil {
			var valueForShapesLinks []ResourceUriElement
			err = json.Unmarshal(*valShapesLinks, &valueForShapesLinks)
			if err != nil {
				return err
			}
			this.ShapesLinks = valueForShapesLinks
		}
	}

    return nil
}
