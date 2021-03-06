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

// Represents header/footer info of slide
type IHeaderFooter interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// True if date is displayed in the footer
	getIsDateTimeVisible() bool
	setIsDateTimeVisible(newValue bool)

	// Text to be displayed as date in the footer
	getDateTimeText() string
	setDateTimeText(newValue string)

	// True if footer is displayed
	getIsFooterVisible() bool
	setIsFooterVisible(newValue bool)

	// Text to be displayed in the footer
	getFooterText() string
	setFooterText(newValue string)

	// True if slide number is displayed in the footer
	getIsSlideNumberVisible() bool
	setIsSlideNumberVisible(newValue bool)
}

type HeaderFooter struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// True if date is displayed in the footer
	IsDateTimeVisible bool `json:"IsDateTimeVisible"`

	// Text to be displayed as date in the footer
	DateTimeText string `json:"DateTimeText,omitempty"`

	// True if footer is displayed
	IsFooterVisible bool `json:"IsFooterVisible"`

	// Text to be displayed in the footer
	FooterText string `json:"FooterText,omitempty"`

	// True if slide number is displayed in the footer
	IsSlideNumberVisible bool `json:"IsSlideNumberVisible"`
}

func NewHeaderFooter() *HeaderFooter {
	instance := new(HeaderFooter)
	return instance
}

func (this *HeaderFooter) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *HeaderFooter) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *HeaderFooter) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *HeaderFooter) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *HeaderFooter) getIsDateTimeVisible() bool {
	return this.IsDateTimeVisible
}

func (this *HeaderFooter) setIsDateTimeVisible(newValue bool) {
	this.IsDateTimeVisible = newValue
}
func (this *HeaderFooter) getDateTimeText() string {
	return this.DateTimeText
}

func (this *HeaderFooter) setDateTimeText(newValue string) {
	this.DateTimeText = newValue
}
func (this *HeaderFooter) getIsFooterVisible() bool {
	return this.IsFooterVisible
}

func (this *HeaderFooter) setIsFooterVisible(newValue bool) {
	this.IsFooterVisible = newValue
}
func (this *HeaderFooter) getFooterText() string {
	return this.FooterText
}

func (this *HeaderFooter) setFooterText(newValue string) {
	this.FooterText = newValue
}
func (this *HeaderFooter) getIsSlideNumberVisible() bool {
	return this.IsSlideNumberVisible
}

func (this *HeaderFooter) setIsSlideNumberVisible(newValue bool) {
	this.IsSlideNumberVisible = newValue
}

func (this *HeaderFooter) UnmarshalJSON(b []byte) error {
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
	
	if valIsDateTimeVisible, ok := objMap["isDateTimeVisible"]; ok {
		if valIsDateTimeVisible != nil {
			var valueForIsDateTimeVisible bool
			err = json.Unmarshal(*valIsDateTimeVisible, &valueForIsDateTimeVisible)
			if err != nil {
				return err
			}
			this.IsDateTimeVisible = valueForIsDateTimeVisible
		}
	}
	if valIsDateTimeVisibleCap, ok := objMap["IsDateTimeVisible"]; ok {
		if valIsDateTimeVisibleCap != nil {
			var valueForIsDateTimeVisible bool
			err = json.Unmarshal(*valIsDateTimeVisibleCap, &valueForIsDateTimeVisible)
			if err != nil {
				return err
			}
			this.IsDateTimeVisible = valueForIsDateTimeVisible
		}
	}
	
	if valDateTimeText, ok := objMap["dateTimeText"]; ok {
		if valDateTimeText != nil {
			var valueForDateTimeText string
			err = json.Unmarshal(*valDateTimeText, &valueForDateTimeText)
			if err != nil {
				return err
			}
			this.DateTimeText = valueForDateTimeText
		}
	}
	if valDateTimeTextCap, ok := objMap["DateTimeText"]; ok {
		if valDateTimeTextCap != nil {
			var valueForDateTimeText string
			err = json.Unmarshal(*valDateTimeTextCap, &valueForDateTimeText)
			if err != nil {
				return err
			}
			this.DateTimeText = valueForDateTimeText
		}
	}
	
	if valIsFooterVisible, ok := objMap["isFooterVisible"]; ok {
		if valIsFooterVisible != nil {
			var valueForIsFooterVisible bool
			err = json.Unmarshal(*valIsFooterVisible, &valueForIsFooterVisible)
			if err != nil {
				return err
			}
			this.IsFooterVisible = valueForIsFooterVisible
		}
	}
	if valIsFooterVisibleCap, ok := objMap["IsFooterVisible"]; ok {
		if valIsFooterVisibleCap != nil {
			var valueForIsFooterVisible bool
			err = json.Unmarshal(*valIsFooterVisibleCap, &valueForIsFooterVisible)
			if err != nil {
				return err
			}
			this.IsFooterVisible = valueForIsFooterVisible
		}
	}
	
	if valFooterText, ok := objMap["footerText"]; ok {
		if valFooterText != nil {
			var valueForFooterText string
			err = json.Unmarshal(*valFooterText, &valueForFooterText)
			if err != nil {
				return err
			}
			this.FooterText = valueForFooterText
		}
	}
	if valFooterTextCap, ok := objMap["FooterText"]; ok {
		if valFooterTextCap != nil {
			var valueForFooterText string
			err = json.Unmarshal(*valFooterTextCap, &valueForFooterText)
			if err != nil {
				return err
			}
			this.FooterText = valueForFooterText
		}
	}
	
	if valIsSlideNumberVisible, ok := objMap["isSlideNumberVisible"]; ok {
		if valIsSlideNumberVisible != nil {
			var valueForIsSlideNumberVisible bool
			err = json.Unmarshal(*valIsSlideNumberVisible, &valueForIsSlideNumberVisible)
			if err != nil {
				return err
			}
			this.IsSlideNumberVisible = valueForIsSlideNumberVisible
		}
	}
	if valIsSlideNumberVisibleCap, ok := objMap["IsSlideNumberVisible"]; ok {
		if valIsSlideNumberVisibleCap != nil {
			var valueForIsSlideNumberVisible bool
			err = json.Unmarshal(*valIsSlideNumberVisibleCap, &valueForIsSlideNumberVisible)
			if err != nil {
				return err
			}
			this.IsSlideNumberVisible = valueForIsSlideNumberVisible
		}
	}

    return nil
}
