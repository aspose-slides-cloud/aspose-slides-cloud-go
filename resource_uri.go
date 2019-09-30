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
	getHref() string
	setHref(newValue string)

	// Gets or sets the relation.
	getRelation() string
	setRelation(newValue string)

	// Gets or sets the type of link.
	getLinkType() string
	setLinkType(newValue string)

	// Gets or sets the title of link.
	getTitle() string
	setTitle(newValue string)
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
}

func (this ResourceUri) getHref() string {
	return this.Href
}

func (this ResourceUri) setHref(newValue string) {
	this.Href = newValue
}
func (this ResourceUri) getRelation() string {
	return this.Relation
}

func (this ResourceUri) setRelation(newValue string) {
	this.Relation = newValue
}
func (this ResourceUri) getLinkType() string {
	return this.LinkType
}

func (this ResourceUri) setLinkType(newValue string) {
	this.LinkType = newValue
}
func (this ResourceUri) getTitle() string {
	return this.Title
}

func (this ResourceUri) setTitle(newValue string) {
	this.Title = newValue
}

func (this *ResourceUri) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valHref, ok := objMap["href"]; ok {
		if valHref != nil {
			var valueForHref string
			err = json.Unmarshal(*valHref, &valueForHref)
			if err != nil {
				return err
			}
			this.Href = valueForHref
		}
	}
	if valHrefCap, ok := objMap["Href"]; ok {
		if valHrefCap != nil {
			var valueForHref string
			err = json.Unmarshal(*valHrefCap, &valueForHref)
			if err != nil {
				return err
			}
			this.Href = valueForHref
		}
	}
	
	if valRelation, ok := objMap["relation"]; ok {
		if valRelation != nil {
			var valueForRelation string
			err = json.Unmarshal(*valRelation, &valueForRelation)
			if err != nil {
				return err
			}
			this.Relation = valueForRelation
		}
	}
	if valRelationCap, ok := objMap["Relation"]; ok {
		if valRelationCap != nil {
			var valueForRelation string
			err = json.Unmarshal(*valRelationCap, &valueForRelation)
			if err != nil {
				return err
			}
			this.Relation = valueForRelation
		}
	}
	
	if valLinkType, ok := objMap["linkType"]; ok {
		if valLinkType != nil {
			var valueForLinkType string
			err = json.Unmarshal(*valLinkType, &valueForLinkType)
			if err != nil {
				return err
			}
			this.LinkType = valueForLinkType
		}
	}
	if valLinkTypeCap, ok := objMap["LinkType"]; ok {
		if valLinkTypeCap != nil {
			var valueForLinkType string
			err = json.Unmarshal(*valLinkTypeCap, &valueForLinkType)
			if err != nil {
				return err
			}
			this.LinkType = valueForLinkType
		}
	}
	
	if valTitle, ok := objMap["title"]; ok {
		if valTitle != nil {
			var valueForTitle string
			err = json.Unmarshal(*valTitle, &valueForTitle)
			if err != nil {
				return err
			}
			this.Title = valueForTitle
		}
	}
	if valTitleCap, ok := objMap["Title"]; ok {
		if valTitleCap != nil {
			var valueForTitle string
			err = json.Unmarshal(*valTitleCap, &valueForTitle)
			if err != nil {
				return err
			}
			this.Title = valueForTitle
		}
	}

    return nil
}
