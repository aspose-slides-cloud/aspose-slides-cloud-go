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

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// List of shape links.
	getShapesLinks() []IResourceUriElement
	setShapesLinks(newValue []IResourceUriElement)
}

type Shapes struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// List of shape links.
	ShapesLinks []IResourceUriElement `json:"ShapesLinks,omitempty"`
}

func NewShapes() *Shapes {
	instance := new(Shapes)
	return instance
}

func (this *Shapes) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Shapes) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Shapes) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Shapes) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Shapes) getShapesLinks() []IResourceUriElement {
	return this.ShapesLinks
}

func (this *Shapes) setShapesLinks(newValue []IResourceUriElement) {
	this.ShapesLinks = newValue
}

func (this *Shapes) UnmarshalJSON(b []byte) error {
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
	
	if valShapesLinks, ok := objMap["shapesLinks"]; ok {
		if valShapesLinks != nil {
			var valueForShapesLinks []ResourceUriElement
			err = json.Unmarshal(*valShapesLinks, &valueForShapesLinks)
			if err != nil {
				return err
			}
			valueForIShapesLinks := make([]IResourceUriElement, len(valueForShapesLinks))
			for i, v := range valueForShapesLinks {
				valueForIShapesLinks[i] = IResourceUriElement(&v)
			}
			this.ShapesLinks = valueForIShapesLinks
		}
	}
	if valShapesLinksCap, ok := objMap["ShapesLinks"]; ok {
		if valShapesLinksCap != nil {
			var valueForShapesLinks []ResourceUriElement
			err = json.Unmarshal(*valShapesLinksCap, &valueForShapesLinks)
			if err != nil {
				return err
			}
			valueForIShapesLinks := make([]IResourceUriElement, len(valueForShapesLinks))
			for i, v := range valueForShapesLinks {
				valueForIShapesLinks[i] = IResourceUriElement(&v)
			}
			this.ShapesLinks = valueForIShapesLinks
		}
	}

    return nil
}
