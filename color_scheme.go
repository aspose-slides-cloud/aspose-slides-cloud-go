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

// Slide's color scheme DTO
type IColorScheme interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// First accent color.
	GetAccent1() string
	SetAccent1(newValue string)

	// Second accent color.
	GetAccent2() string
	SetAccent2(newValue string)

	// Third accent color.
	GetAccent3() string
	SetAccent3(newValue string)

	// Fourth accent color.
	GetAccent4() string
	SetAccent4(newValue string)

	// Fifth accent color.
	GetAccent5() string
	SetAccent5(newValue string)

	// Sixth accent color.
	GetAccent6() string
	SetAccent6(newValue string)

	// First dark color.
	GetDark1() string
	SetDark1(newValue string)

	// Second dark color.
	GetDark2() string
	SetDark2(newValue string)

	// Visited hyperlink color.
	GetFollowedHyperlink() string
	SetFollowedHyperlink(newValue string)

	// Hyperlink color/
	GetHyperlink() string
	SetHyperlink(newValue string)

	// First light color.
	GetLight1() string
	SetLight1(newValue string)

	// Second light color.
	GetLight2() string
	SetLight2(newValue string)
}

type ColorScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// First accent color.
	Accent1 string `json:"Accent1,omitempty"`

	// Second accent color.
	Accent2 string `json:"Accent2,omitempty"`

	// Third accent color.
	Accent3 string `json:"Accent3,omitempty"`

	// Fourth accent color.
	Accent4 string `json:"Accent4,omitempty"`

	// Fifth accent color.
	Accent5 string `json:"Accent5,omitempty"`

	// Sixth accent color.
	Accent6 string `json:"Accent6,omitempty"`

	// First dark color.
	Dark1 string `json:"Dark1,omitempty"`

	// Second dark color.
	Dark2 string `json:"Dark2,omitempty"`

	// Visited hyperlink color.
	FollowedHyperlink string `json:"FollowedHyperlink,omitempty"`

	// Hyperlink color/
	Hyperlink string `json:"Hyperlink,omitempty"`

	// First light color.
	Light1 string `json:"Light1,omitempty"`

	// Second light color.
	Light2 string `json:"Light2,omitempty"`
}

func NewColorScheme() *ColorScheme {
	instance := new(ColorScheme)
	return instance
}

func (this *ColorScheme) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ColorScheme) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ColorScheme) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ColorScheme) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ColorScheme) GetAccent1() string {
	return this.Accent1
}

func (this *ColorScheme) SetAccent1(newValue string) {
	this.Accent1 = newValue
}
func (this *ColorScheme) GetAccent2() string {
	return this.Accent2
}

func (this *ColorScheme) SetAccent2(newValue string) {
	this.Accent2 = newValue
}
func (this *ColorScheme) GetAccent3() string {
	return this.Accent3
}

func (this *ColorScheme) SetAccent3(newValue string) {
	this.Accent3 = newValue
}
func (this *ColorScheme) GetAccent4() string {
	return this.Accent4
}

func (this *ColorScheme) SetAccent4(newValue string) {
	this.Accent4 = newValue
}
func (this *ColorScheme) GetAccent5() string {
	return this.Accent5
}

func (this *ColorScheme) SetAccent5(newValue string) {
	this.Accent5 = newValue
}
func (this *ColorScheme) GetAccent6() string {
	return this.Accent6
}

func (this *ColorScheme) SetAccent6(newValue string) {
	this.Accent6 = newValue
}
func (this *ColorScheme) GetDark1() string {
	return this.Dark1
}

func (this *ColorScheme) SetDark1(newValue string) {
	this.Dark1 = newValue
}
func (this *ColorScheme) GetDark2() string {
	return this.Dark2
}

func (this *ColorScheme) SetDark2(newValue string) {
	this.Dark2 = newValue
}
func (this *ColorScheme) GetFollowedHyperlink() string {
	return this.FollowedHyperlink
}

func (this *ColorScheme) SetFollowedHyperlink(newValue string) {
	this.FollowedHyperlink = newValue
}
func (this *ColorScheme) GetHyperlink() string {
	return this.Hyperlink
}

func (this *ColorScheme) SetHyperlink(newValue string) {
	this.Hyperlink = newValue
}
func (this *ColorScheme) GetLight1() string {
	return this.Light1
}

func (this *ColorScheme) SetLight1(newValue string) {
	this.Light1 = newValue
}
func (this *ColorScheme) GetLight2() string {
	return this.Light2
}

func (this *ColorScheme) SetLight2(newValue string) {
	this.Light2 = newValue
}

func (this *ColorScheme) UnmarshalJSON(b []byte) error {
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
	
	if valAccent1, ok := objMap["accent1"]; ok {
		if valAccent1 != nil {
			var valueForAccent1 string
			err = json.Unmarshal(*valAccent1, &valueForAccent1)
			if err != nil {
				return err
			}
			this.Accent1 = valueForAccent1
		}
	}
	if valAccent1Cap, ok := objMap["Accent1"]; ok {
		if valAccent1Cap != nil {
			var valueForAccent1 string
			err = json.Unmarshal(*valAccent1Cap, &valueForAccent1)
			if err != nil {
				return err
			}
			this.Accent1 = valueForAccent1
		}
	}
	
	if valAccent2, ok := objMap["accent2"]; ok {
		if valAccent2 != nil {
			var valueForAccent2 string
			err = json.Unmarshal(*valAccent2, &valueForAccent2)
			if err != nil {
				return err
			}
			this.Accent2 = valueForAccent2
		}
	}
	if valAccent2Cap, ok := objMap["Accent2"]; ok {
		if valAccent2Cap != nil {
			var valueForAccent2 string
			err = json.Unmarshal(*valAccent2Cap, &valueForAccent2)
			if err != nil {
				return err
			}
			this.Accent2 = valueForAccent2
		}
	}
	
	if valAccent3, ok := objMap["accent3"]; ok {
		if valAccent3 != nil {
			var valueForAccent3 string
			err = json.Unmarshal(*valAccent3, &valueForAccent3)
			if err != nil {
				return err
			}
			this.Accent3 = valueForAccent3
		}
	}
	if valAccent3Cap, ok := objMap["Accent3"]; ok {
		if valAccent3Cap != nil {
			var valueForAccent3 string
			err = json.Unmarshal(*valAccent3Cap, &valueForAccent3)
			if err != nil {
				return err
			}
			this.Accent3 = valueForAccent3
		}
	}
	
	if valAccent4, ok := objMap["accent4"]; ok {
		if valAccent4 != nil {
			var valueForAccent4 string
			err = json.Unmarshal(*valAccent4, &valueForAccent4)
			if err != nil {
				return err
			}
			this.Accent4 = valueForAccent4
		}
	}
	if valAccent4Cap, ok := objMap["Accent4"]; ok {
		if valAccent4Cap != nil {
			var valueForAccent4 string
			err = json.Unmarshal(*valAccent4Cap, &valueForAccent4)
			if err != nil {
				return err
			}
			this.Accent4 = valueForAccent4
		}
	}
	
	if valAccent5, ok := objMap["accent5"]; ok {
		if valAccent5 != nil {
			var valueForAccent5 string
			err = json.Unmarshal(*valAccent5, &valueForAccent5)
			if err != nil {
				return err
			}
			this.Accent5 = valueForAccent5
		}
	}
	if valAccent5Cap, ok := objMap["Accent5"]; ok {
		if valAccent5Cap != nil {
			var valueForAccent5 string
			err = json.Unmarshal(*valAccent5Cap, &valueForAccent5)
			if err != nil {
				return err
			}
			this.Accent5 = valueForAccent5
		}
	}
	
	if valAccent6, ok := objMap["accent6"]; ok {
		if valAccent6 != nil {
			var valueForAccent6 string
			err = json.Unmarshal(*valAccent6, &valueForAccent6)
			if err != nil {
				return err
			}
			this.Accent6 = valueForAccent6
		}
	}
	if valAccent6Cap, ok := objMap["Accent6"]; ok {
		if valAccent6Cap != nil {
			var valueForAccent6 string
			err = json.Unmarshal(*valAccent6Cap, &valueForAccent6)
			if err != nil {
				return err
			}
			this.Accent6 = valueForAccent6
		}
	}
	
	if valDark1, ok := objMap["dark1"]; ok {
		if valDark1 != nil {
			var valueForDark1 string
			err = json.Unmarshal(*valDark1, &valueForDark1)
			if err != nil {
				return err
			}
			this.Dark1 = valueForDark1
		}
	}
	if valDark1Cap, ok := objMap["Dark1"]; ok {
		if valDark1Cap != nil {
			var valueForDark1 string
			err = json.Unmarshal(*valDark1Cap, &valueForDark1)
			if err != nil {
				return err
			}
			this.Dark1 = valueForDark1
		}
	}
	
	if valDark2, ok := objMap["dark2"]; ok {
		if valDark2 != nil {
			var valueForDark2 string
			err = json.Unmarshal(*valDark2, &valueForDark2)
			if err != nil {
				return err
			}
			this.Dark2 = valueForDark2
		}
	}
	if valDark2Cap, ok := objMap["Dark2"]; ok {
		if valDark2Cap != nil {
			var valueForDark2 string
			err = json.Unmarshal(*valDark2Cap, &valueForDark2)
			if err != nil {
				return err
			}
			this.Dark2 = valueForDark2
		}
	}
	
	if valFollowedHyperlink, ok := objMap["followedHyperlink"]; ok {
		if valFollowedHyperlink != nil {
			var valueForFollowedHyperlink string
			err = json.Unmarshal(*valFollowedHyperlink, &valueForFollowedHyperlink)
			if err != nil {
				return err
			}
			this.FollowedHyperlink = valueForFollowedHyperlink
		}
	}
	if valFollowedHyperlinkCap, ok := objMap["FollowedHyperlink"]; ok {
		if valFollowedHyperlinkCap != nil {
			var valueForFollowedHyperlink string
			err = json.Unmarshal(*valFollowedHyperlinkCap, &valueForFollowedHyperlink)
			if err != nil {
				return err
			}
			this.FollowedHyperlink = valueForFollowedHyperlink
		}
	}
	
	if valHyperlink, ok := objMap["hyperlink"]; ok {
		if valHyperlink != nil {
			var valueForHyperlink string
			err = json.Unmarshal(*valHyperlink, &valueForHyperlink)
			if err != nil {
				return err
			}
			this.Hyperlink = valueForHyperlink
		}
	}
	if valHyperlinkCap, ok := objMap["Hyperlink"]; ok {
		if valHyperlinkCap != nil {
			var valueForHyperlink string
			err = json.Unmarshal(*valHyperlinkCap, &valueForHyperlink)
			if err != nil {
				return err
			}
			this.Hyperlink = valueForHyperlink
		}
	}
	
	if valLight1, ok := objMap["light1"]; ok {
		if valLight1 != nil {
			var valueForLight1 string
			err = json.Unmarshal(*valLight1, &valueForLight1)
			if err != nil {
				return err
			}
			this.Light1 = valueForLight1
		}
	}
	if valLight1Cap, ok := objMap["Light1"]; ok {
		if valLight1Cap != nil {
			var valueForLight1 string
			err = json.Unmarshal(*valLight1Cap, &valueForLight1)
			if err != nil {
				return err
			}
			this.Light1 = valueForLight1
		}
	}
	
	if valLight2, ok := objMap["light2"]; ok {
		if valLight2 != nil {
			var valueForLight2 string
			err = json.Unmarshal(*valLight2, &valueForLight2)
			if err != nil {
				return err
			}
			this.Light2 = valueForLight2
		}
	}
	if valLight2Cap, ok := objMap["Light2"]; ok {
		if valLight2Cap != nil {
			var valueForLight2 string
			err = json.Unmarshal(*valLight2Cap, &valueForLight2)
			if err != nil {
				return err
			}
			this.Light2 = valueForLight2
		}
	}

	return nil
}
