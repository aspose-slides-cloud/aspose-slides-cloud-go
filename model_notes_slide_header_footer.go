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

// Represents header/footer info of notes slide
type INotesSlideHeaderFooter interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// True if date is displayed in the footer
	GetIsDateTimeVisible() *bool
	SetIsDateTimeVisible(newValue *bool)

	// Text to be displayed as date in the footer
	GetDateTimeText() string
	SetDateTimeText(newValue string)

	// True if footer is displayed
	GetIsFooterVisible() *bool
	SetIsFooterVisible(newValue *bool)

	// Text to be displayed in the footer
	GetFooterText() string
	SetFooterText(newValue string)

	// True if header is displayed
	GetIsHeaderVisible() *bool
	SetIsHeaderVisible(newValue *bool)

	// Text to be displayed in the header
	GetHeaderText() string
	SetHeaderText(newValue string)

	// True if slide number is displayed in the footer
	GetIsSlideNumberVisible() *bool
	SetIsSlideNumberVisible(newValue *bool)
}

type NotesSlideHeaderFooter struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// True if date is displayed in the footer
	IsDateTimeVisible *bool `json:"IsDateTimeVisible"`

	// Text to be displayed as date in the footer
	DateTimeText string `json:"DateTimeText,omitempty"`

	// True if footer is displayed
	IsFooterVisible *bool `json:"IsFooterVisible"`

	// Text to be displayed in the footer
	FooterText string `json:"FooterText,omitempty"`

	// True if header is displayed
	IsHeaderVisible *bool `json:"IsHeaderVisible"`

	// Text to be displayed in the header
	HeaderText string `json:"HeaderText,omitempty"`

	// True if slide number is displayed in the footer
	IsSlideNumberVisible *bool `json:"IsSlideNumberVisible"`
}

func NewNotesSlideHeaderFooter() *NotesSlideHeaderFooter {
	instance := new(NotesSlideHeaderFooter)
	return instance
}

func (this *NotesSlideHeaderFooter) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *NotesSlideHeaderFooter) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *NotesSlideHeaderFooter) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *NotesSlideHeaderFooter) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *NotesSlideHeaderFooter) GetIsDateTimeVisible() *bool {
	return this.IsDateTimeVisible
}

func (this *NotesSlideHeaderFooter) SetIsDateTimeVisible(newValue *bool) {
	this.IsDateTimeVisible = newValue
}
func (this *NotesSlideHeaderFooter) GetDateTimeText() string {
	return this.DateTimeText
}

func (this *NotesSlideHeaderFooter) SetDateTimeText(newValue string) {
	this.DateTimeText = newValue
}
func (this *NotesSlideHeaderFooter) GetIsFooterVisible() *bool {
	return this.IsFooterVisible
}

func (this *NotesSlideHeaderFooter) SetIsFooterVisible(newValue *bool) {
	this.IsFooterVisible = newValue
}
func (this *NotesSlideHeaderFooter) GetFooterText() string {
	return this.FooterText
}

func (this *NotesSlideHeaderFooter) SetFooterText(newValue string) {
	this.FooterText = newValue
}
func (this *NotesSlideHeaderFooter) GetIsHeaderVisible() *bool {
	return this.IsHeaderVisible
}

func (this *NotesSlideHeaderFooter) SetIsHeaderVisible(newValue *bool) {
	this.IsHeaderVisible = newValue
}
func (this *NotesSlideHeaderFooter) GetHeaderText() string {
	return this.HeaderText
}

func (this *NotesSlideHeaderFooter) SetHeaderText(newValue string) {
	this.HeaderText = newValue
}
func (this *NotesSlideHeaderFooter) GetIsSlideNumberVisible() *bool {
	return this.IsSlideNumberVisible
}

func (this *NotesSlideHeaderFooter) SetIsSlideNumberVisible(newValue *bool) {
	this.IsSlideNumberVisible = newValue
}

func (this *NotesSlideHeaderFooter) UnmarshalJSON(b []byte) error {
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
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valIsDateTimeVisible, ok := objMap["isDateTimeVisible"]; ok {
		if valIsDateTimeVisible != nil {
			var valueForIsDateTimeVisible *bool
			err = json.Unmarshal(*valIsDateTimeVisible, &valueForIsDateTimeVisible)
			if err != nil {
				return err
			}
			this.IsDateTimeVisible = valueForIsDateTimeVisible
		}
	}
	if valIsDateTimeVisibleCap, ok := objMap["IsDateTimeVisible"]; ok {
		if valIsDateTimeVisibleCap != nil {
			var valueForIsDateTimeVisible *bool
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
			var valueForIsFooterVisible *bool
			err = json.Unmarshal(*valIsFooterVisible, &valueForIsFooterVisible)
			if err != nil {
				return err
			}
			this.IsFooterVisible = valueForIsFooterVisible
		}
	}
	if valIsFooterVisibleCap, ok := objMap["IsFooterVisible"]; ok {
		if valIsFooterVisibleCap != nil {
			var valueForIsFooterVisible *bool
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
	
	if valIsHeaderVisible, ok := objMap["isHeaderVisible"]; ok {
		if valIsHeaderVisible != nil {
			var valueForIsHeaderVisible *bool
			err = json.Unmarshal(*valIsHeaderVisible, &valueForIsHeaderVisible)
			if err != nil {
				return err
			}
			this.IsHeaderVisible = valueForIsHeaderVisible
		}
	}
	if valIsHeaderVisibleCap, ok := objMap["IsHeaderVisible"]; ok {
		if valIsHeaderVisibleCap != nil {
			var valueForIsHeaderVisible *bool
			err = json.Unmarshal(*valIsHeaderVisibleCap, &valueForIsHeaderVisible)
			if err != nil {
				return err
			}
			this.IsHeaderVisible = valueForIsHeaderVisible
		}
	}
	
	if valHeaderText, ok := objMap["headerText"]; ok {
		if valHeaderText != nil {
			var valueForHeaderText string
			err = json.Unmarshal(*valHeaderText, &valueForHeaderText)
			if err != nil {
				return err
			}
			this.HeaderText = valueForHeaderText
		}
	}
	if valHeaderTextCap, ok := objMap["HeaderText"]; ok {
		if valHeaderTextCap != nil {
			var valueForHeaderText string
			err = json.Unmarshal(*valHeaderTextCap, &valueForHeaderText)
			if err != nil {
				return err
			}
			this.HeaderText = valueForHeaderText
		}
	}
	
	if valIsSlideNumberVisible, ok := objMap["isSlideNumberVisible"]; ok {
		if valIsSlideNumberVisible != nil {
			var valueForIsSlideNumberVisible *bool
			err = json.Unmarshal(*valIsSlideNumberVisible, &valueForIsSlideNumberVisible)
			if err != nil {
				return err
			}
			this.IsSlideNumberVisible = valueForIsSlideNumberVisible
		}
	}
	if valIsSlideNumberVisibleCap, ok := objMap["IsSlideNumberVisible"]; ok {
		if valIsSlideNumberVisibleCap != nil {
			var valueForIsSlideNumberVisible *bool
			err = json.Unmarshal(*valIsSlideNumberVisibleCap, &valueForIsSlideNumberVisible)
			if err != nil {
				return err
			}
			this.IsSlideNumberVisible = valueForIsSlideNumberVisible
		}
	}

	return nil
}
