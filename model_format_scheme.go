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

// Represents Format Scheme for slide's theme
type IFormatScheme interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Background styles.
	GetBackgroundStyles() []IFillFormat
	SetBackgroundStyles(newValue []IFillFormat)

	// Effect styles.
	GetEffectStyles() []IEffectFormat
	SetEffectStyles(newValue []IEffectFormat)

	// Fill styles.
	GetFillStyles() []IFillFormat
	SetFillStyles(newValue []IFillFormat)

	// Line style.
	GetLineStyles() []ILineFormat
	SetLineStyles(newValue []ILineFormat)
}

type FormatScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Background styles.
	BackgroundStyles []IFillFormat `json:"BackgroundStyles,omitempty"`

	// Effect styles.
	EffectStyles []IEffectFormat `json:"EffectStyles,omitempty"`

	// Fill styles.
	FillStyles []IFillFormat `json:"FillStyles,omitempty"`

	// Line style.
	LineStyles []ILineFormat `json:"LineStyles,omitempty"`
}

func NewFormatScheme() *FormatScheme {
	instance := new(FormatScheme)
	return instance
}

func (this *FormatScheme) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *FormatScheme) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *FormatScheme) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *FormatScheme) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *FormatScheme) GetBackgroundStyles() []IFillFormat {
	return this.BackgroundStyles
}

func (this *FormatScheme) SetBackgroundStyles(newValue []IFillFormat) {
	this.BackgroundStyles = newValue
}
func (this *FormatScheme) GetEffectStyles() []IEffectFormat {
	return this.EffectStyles
}

func (this *FormatScheme) SetEffectStyles(newValue []IEffectFormat) {
	this.EffectStyles = newValue
}
func (this *FormatScheme) GetFillStyles() []IFillFormat {
	return this.FillStyles
}

func (this *FormatScheme) SetFillStyles(newValue []IFillFormat) {
	this.FillStyles = newValue
}
func (this *FormatScheme) GetLineStyles() []ILineFormat {
	return this.LineStyles
}

func (this *FormatScheme) SetLineStyles(newValue []ILineFormat) {
	this.LineStyles = newValue
}

func (this *FormatScheme) UnmarshalJSON(b []byte) error {
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
	
	if valBackgroundStyles, ok := GetMapValue(objMap, "backgroundStyles"); ok {
		if valBackgroundStyles != nil {
			var valueForBackgroundStyles []json.RawMessage
			err = json.Unmarshal(*valBackgroundStyles, &valueForBackgroundStyles)
			if err != nil {
				return err
			}
			valueForIBackgroundStyles := make([]IFillFormat, len(valueForBackgroundStyles))
			for i, v := range valueForBackgroundStyles {
				vObject, err := createObjectForType("FillFormat", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIBackgroundStyles[i] = vObject.(IFillFormat)
				}
			}
			this.BackgroundStyles = valueForIBackgroundStyles
		}
	}
	
	if valEffectStyles, ok := GetMapValue(objMap, "effectStyles"); ok {
		if valEffectStyles != nil {
			var valueForEffectStyles []json.RawMessage
			err = json.Unmarshal(*valEffectStyles, &valueForEffectStyles)
			if err != nil {
				return err
			}
			valueForIEffectStyles := make([]IEffectFormat, len(valueForEffectStyles))
			for i, v := range valueForEffectStyles {
				vObject, err := createObjectForType("EffectFormat", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIEffectStyles[i] = vObject.(IEffectFormat)
				}
			}
			this.EffectStyles = valueForIEffectStyles
		}
	}
	
	if valFillStyles, ok := GetMapValue(objMap, "fillStyles"); ok {
		if valFillStyles != nil {
			var valueForFillStyles []json.RawMessage
			err = json.Unmarshal(*valFillStyles, &valueForFillStyles)
			if err != nil {
				return err
			}
			valueForIFillStyles := make([]IFillFormat, len(valueForFillStyles))
			for i, v := range valueForFillStyles {
				vObject, err := createObjectForType("FillFormat", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFillStyles[i] = vObject.(IFillFormat)
				}
			}
			this.FillStyles = valueForIFillStyles
		}
	}
	
	if valLineStyles, ok := GetMapValue(objMap, "lineStyles"); ok {
		if valLineStyles != nil {
			var valueForLineStyles []json.RawMessage
			err = json.Unmarshal(*valLineStyles, &valueForLineStyles)
			if err != nil {
				return err
			}
			valueForILineStyles := make([]ILineFormat, len(valueForLineStyles))
			for i, v := range valueForLineStyles {
				vObject, err := createObjectForType("LineFormat", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForILineStyles[i] = vObject.(ILineFormat)
				}
			}
			this.LineStyles = valueForILineStyles
		}
	}

	return nil
}
