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

// Represents Slide's theme 
type ITheme interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Name.
	GetName() string
	SetName(newValue string)

	// Color scheme.
	GetColorScheme() IResourceUri
	SetColorScheme(newValue IResourceUri)

	// Font scheme.
	GetFontScheme() IResourceUri
	SetFontScheme(newValue IResourceUri)

	// Format scheme.
	GetFormatScheme() IResourceUri
	SetFormatScheme(newValue IResourceUri)
}

type Theme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Name.
	Name string `json:"Name,omitempty"`

	// Color scheme.
	ColorScheme IResourceUri `json:"ColorScheme,omitempty"`

	// Font scheme.
	FontScheme IResourceUri `json:"FontScheme,omitempty"`

	// Format scheme.
	FormatScheme IResourceUri `json:"FormatScheme,omitempty"`
}

func NewTheme() *Theme {
	instance := new(Theme)
	return instance
}

func (this *Theme) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Theme) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Theme) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Theme) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Theme) GetName() string {
	return this.Name
}

func (this *Theme) SetName(newValue string) {
	this.Name = newValue
}
func (this *Theme) GetColorScheme() IResourceUri {
	return this.ColorScheme
}

func (this *Theme) SetColorScheme(newValue IResourceUri) {
	this.ColorScheme = newValue
}
func (this *Theme) GetFontScheme() IResourceUri {
	return this.FontScheme
}

func (this *Theme) SetFontScheme(newValue IResourceUri) {
	this.FontScheme = newValue
}
func (this *Theme) GetFormatScheme() IResourceUri {
	return this.FormatScheme
}

func (this *Theme) SetFormatScheme(newValue IResourceUri) {
	this.FormatScheme = newValue
}

func (this *Theme) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := GetMapValue(objMap, "selfUri"); ok {
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
	
	if valAlternateLinks, ok := GetMapValue(objMap, "alternateLinks"); ok {
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
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valColorScheme, ok := GetMapValue(objMap, "colorScheme"); ok {
		if valColorScheme != nil {
			var valueForColorScheme ResourceUri
			err = json.Unmarshal(*valColorScheme, &valueForColorScheme)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valColorScheme)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valColorScheme, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.ColorScheme = vInterfaceObject
			}
		}
	}
	
	if valFontScheme, ok := GetMapValue(objMap, "fontScheme"); ok {
		if valFontScheme != nil {
			var valueForFontScheme ResourceUri
			err = json.Unmarshal(*valFontScheme, &valueForFontScheme)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valFontScheme)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFontScheme, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.FontScheme = vInterfaceObject
			}
		}
	}
	
	if valFormatScheme, ok := GetMapValue(objMap, "formatScheme"); ok {
		if valFormatScheme != nil {
			var valueForFormatScheme ResourceUri
			err = json.Unmarshal(*valFormatScheme, &valueForFormatScheme)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valFormatScheme)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFormatScheme, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.FormatScheme = vInterfaceObject
			}
		}
	}

	return nil
}
