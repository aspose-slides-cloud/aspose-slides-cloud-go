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

// Represents placeholder
type IPlaceholder interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	getIndex() int32
	setIndex(newValue int32)

	getOrientation() string
	setOrientation(newValue string)

	getSize() string
	setSize(newValue string)

	getType() string
	setType(newValue string)

	getShape() IResourceUriElement
	setShape(newValue IResourceUriElement)
}

type Placeholder struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	Index int32 `json:"Index"`

	Orientation string `json:"Orientation"`

	Size string `json:"Size"`

	Type_ string `json:"Type"`

	Shape IResourceUriElement `json:"Shape,omitempty"`
}

func (this Placeholder) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this Placeholder) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this Placeholder) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this Placeholder) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this Placeholder) getIndex() int32 {
	return this.Index
}

func (this Placeholder) setIndex(newValue int32) {
	this.Index = newValue
}
func (this Placeholder) getOrientation() string {
	return this.Orientation
}

func (this Placeholder) setOrientation(newValue string) {
	this.Orientation = newValue
}
func (this Placeholder) getSize() string {
	return this.Size
}

func (this Placeholder) setSize(newValue string) {
	this.Size = newValue
}
func (this Placeholder) getType() string {
	return this.Type_
}

func (this Placeholder) setType(newValue string) {
	this.Type_ = newValue
}
func (this Placeholder) getShape() IResourceUriElement {
	return this.Shape
}

func (this Placeholder) setShape(newValue IResourceUriElement) {
	this.Shape = newValue
}

func (this *Placeholder) UnmarshalJSON(b []byte) error {
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

	if valIndex, ok := objMap["Index"]; ok {
		if valIndex != nil {
			var valueForIndex int32
			err = json.Unmarshal(*valIndex, &valueForIndex)
			if err != nil {
				return err
			}
			this.Index = valueForIndex
		}
	}

	if valOrientation, ok := objMap["Orientation"]; ok {
		if valOrientation != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientation, &valueForOrientation)
			if err != nil {
				return err
			}
			this.Orientation = valueForOrientation
		}
	}

	if valSize, ok := objMap["Size"]; ok {
		if valSize != nil {
			var valueForSize string
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				return err
			}
			this.Size = valueForSize
		}
	}

	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}

	if valShape, ok := objMap["Shape"]; ok {
		if valShape != nil {
			var valueForShape ResourceUriElement
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				return err
			}
			this.Shape = valueForShape
		}
	}

    return nil
}
