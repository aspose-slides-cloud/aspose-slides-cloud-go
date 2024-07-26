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

// Represents Resource URI
type IResourceUri interface {

	// Gets or sets the href.
	GetHref() string
	SetHref(newValue string)

	// Gets or sets the relation.
	GetRelation() string
	SetRelation(newValue string)

	// Gets or sets the type of link.
	GetLinkType() string
	SetLinkType(newValue string)

	// Gets or sets the title of link.
	GetTitle() string
	SetTitle(newValue string)

	// Resource slide index.
	GetSlideIndex() int32
	SetSlideIndex(newValue int32)

	// Resource shape index.
	GetShapeIndex() int32
	SetShapeIndex(newValue int32)
}

type ResourceUri struct {

	// Gets or sets the href.
	Href string `json:"Href,omitempty"`

	// Gets or sets the relation.
	Relation string `json:"Relation,omitempty"`

	// Gets or sets the type of link.
	LinkType string `json:"LinkType,omitempty"`

	// Gets or sets the title of link.
	Title string `json:"Title,omitempty"`

	// Resource slide index.
	SlideIndex int32 `json:"SlideIndex,omitempty"`

	// Resource shape index.
	ShapeIndex int32 `json:"ShapeIndex,omitempty"`
}

func NewResourceUri() *ResourceUri {
	instance := new(ResourceUri)
	return instance
}

func (this *ResourceUri) GetHref() string {
	return this.Href
}

func (this *ResourceUri) SetHref(newValue string) {
	this.Href = newValue
}
func (this *ResourceUri) GetRelation() string {
	return this.Relation
}

func (this *ResourceUri) SetRelation(newValue string) {
	this.Relation = newValue
}
func (this *ResourceUri) GetLinkType() string {
	return this.LinkType
}

func (this *ResourceUri) SetLinkType(newValue string) {
	this.LinkType = newValue
}
func (this *ResourceUri) GetTitle() string {
	return this.Title
}

func (this *ResourceUri) SetTitle(newValue string) {
	this.Title = newValue
}
func (this *ResourceUri) GetSlideIndex() int32 {
	return this.SlideIndex
}

func (this *ResourceUri) SetSlideIndex(newValue int32) {
	this.SlideIndex = newValue
}
func (this *ResourceUri) GetShapeIndex() int32 {
	return this.ShapeIndex
}

func (this *ResourceUri) SetShapeIndex(newValue int32) {
	this.ShapeIndex = newValue
}

func (this *ResourceUri) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valHref, ok := GetMapValue(objMap, "href"); ok {
		if valHref != nil {
			var valueForHref string
			err = json.Unmarshal(*valHref, &valueForHref)
			if err != nil {
				return err
			}
			this.Href = valueForHref
		}
	}
	
	if valRelation, ok := GetMapValue(objMap, "relation"); ok {
		if valRelation != nil {
			var valueForRelation string
			err = json.Unmarshal(*valRelation, &valueForRelation)
			if err != nil {
				return err
			}
			this.Relation = valueForRelation
		}
	}
	
	if valLinkType, ok := GetMapValue(objMap, "linkType"); ok {
		if valLinkType != nil {
			var valueForLinkType string
			err = json.Unmarshal(*valLinkType, &valueForLinkType)
			if err != nil {
				return err
			}
			this.LinkType = valueForLinkType
		}
	}
	
	if valTitle, ok := GetMapValue(objMap, "title"); ok {
		if valTitle != nil {
			var valueForTitle string
			err = json.Unmarshal(*valTitle, &valueForTitle)
			if err != nil {
				return err
			}
			this.Title = valueForTitle
		}
	}
	
	if valSlideIndex, ok := GetMapValue(objMap, "slideIndex"); ok {
		if valSlideIndex != nil {
			var valueForSlideIndex int32
			err = json.Unmarshal(*valSlideIndex, &valueForSlideIndex)
			if err != nil {
				return err
			}
			this.SlideIndex = valueForSlideIndex
		}
	}
	
	if valShapeIndex, ok := GetMapValue(objMap, "shapeIndex"); ok {
		if valShapeIndex != nil {
			var valueForShapeIndex int32
			err = json.Unmarshal(*valShapeIndex, &valueForShapeIndex)
			if err != nil {
				return err
			}
			this.ShapeIndex = valueForShapeIndex
		}
	}

	return nil
}
