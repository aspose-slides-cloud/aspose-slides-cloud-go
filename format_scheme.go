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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Background style links.
	getBackgroundStyles() []IResourceUri
	setBackgroundStyles(newValue []IResourceUri)

	// Effect style links.
	getEffectStyles() []IResourceUri
	setEffectStyles(newValue []IResourceUri)

	// Fill style links.
	getFillStyles() []IResourceUri
	setFillStyles(newValue []IResourceUri)

	// Line style links.
	getLineStyles() []IResourceUri
	setLineStyles(newValue []IResourceUri)
}

type FormatScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Background style links.
	BackgroundStyles []IResourceUri `json:"BackgroundStyles,omitempty"`

	// Effect style links.
	EffectStyles []IResourceUri `json:"EffectStyles,omitempty"`

	// Fill style links.
	FillStyles []IResourceUri `json:"FillStyles,omitempty"`

	// Line style links.
	LineStyles []IResourceUri `json:"LineStyles,omitempty"`
}

func NewFormatScheme() *FormatScheme {
	instance := new(FormatScheme)
	return instance
}

func (this *FormatScheme) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *FormatScheme) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *FormatScheme) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *FormatScheme) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *FormatScheme) getBackgroundStyles() []IResourceUri {
	return this.BackgroundStyles
}

func (this *FormatScheme) setBackgroundStyles(newValue []IResourceUri) {
	this.BackgroundStyles = newValue
}
func (this *FormatScheme) getEffectStyles() []IResourceUri {
	return this.EffectStyles
}

func (this *FormatScheme) setEffectStyles(newValue []IResourceUri) {
	this.EffectStyles = newValue
}
func (this *FormatScheme) getFillStyles() []IResourceUri {
	return this.FillStyles
}

func (this *FormatScheme) setFillStyles(newValue []IResourceUri) {
	this.FillStyles = newValue
}
func (this *FormatScheme) getLineStyles() []IResourceUri {
	return this.LineStyles
}

func (this *FormatScheme) setLineStyles(newValue []IResourceUri) {
	this.LineStyles = newValue
}

func (this *FormatScheme) UnmarshalJSON(b []byte) error {
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
	
	if valBackgroundStyles, ok := objMap["backgroundStyles"]; ok {
		if valBackgroundStyles != nil {
			var valueForBackgroundStyles []json.RawMessage
			err = json.Unmarshal(*valBackgroundStyles, &valueForBackgroundStyles)
			if err != nil {
				return err
			}
			valueForIBackgroundStyles := make([]IResourceUri, len(valueForBackgroundStyles))
			for i, v := range valueForBackgroundStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIBackgroundStyles[i] = vObject.(IResourceUri)
				}
			}
			this.BackgroundStyles = valueForIBackgroundStyles
		}
	}
	if valBackgroundStylesCap, ok := objMap["BackgroundStyles"]; ok {
		if valBackgroundStylesCap != nil {
			var valueForBackgroundStyles []json.RawMessage
			err = json.Unmarshal(*valBackgroundStylesCap, &valueForBackgroundStyles)
			if err != nil {
				return err
			}
			valueForIBackgroundStyles := make([]IResourceUri, len(valueForBackgroundStyles))
			for i, v := range valueForBackgroundStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIBackgroundStyles[i] = vObject.(IResourceUri)
				}
			}
			this.BackgroundStyles = valueForIBackgroundStyles
		}
	}
	
	if valEffectStyles, ok := objMap["effectStyles"]; ok {
		if valEffectStyles != nil {
			var valueForEffectStyles []json.RawMessage
			err = json.Unmarshal(*valEffectStyles, &valueForEffectStyles)
			if err != nil {
				return err
			}
			valueForIEffectStyles := make([]IResourceUri, len(valueForEffectStyles))
			for i, v := range valueForEffectStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIEffectStyles[i] = vObject.(IResourceUri)
				}
			}
			this.EffectStyles = valueForIEffectStyles
		}
	}
	if valEffectStylesCap, ok := objMap["EffectStyles"]; ok {
		if valEffectStylesCap != nil {
			var valueForEffectStyles []json.RawMessage
			err = json.Unmarshal(*valEffectStylesCap, &valueForEffectStyles)
			if err != nil {
				return err
			}
			valueForIEffectStyles := make([]IResourceUri, len(valueForEffectStyles))
			for i, v := range valueForEffectStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIEffectStyles[i] = vObject.(IResourceUri)
				}
			}
			this.EffectStyles = valueForIEffectStyles
		}
	}
	
	if valFillStyles, ok := objMap["fillStyles"]; ok {
		if valFillStyles != nil {
			var valueForFillStyles []json.RawMessage
			err = json.Unmarshal(*valFillStyles, &valueForFillStyles)
			if err != nil {
				return err
			}
			valueForIFillStyles := make([]IResourceUri, len(valueForFillStyles))
			for i, v := range valueForFillStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFillStyles[i] = vObject.(IResourceUri)
				}
			}
			this.FillStyles = valueForIFillStyles
		}
	}
	if valFillStylesCap, ok := objMap["FillStyles"]; ok {
		if valFillStylesCap != nil {
			var valueForFillStyles []json.RawMessage
			err = json.Unmarshal(*valFillStylesCap, &valueForFillStyles)
			if err != nil {
				return err
			}
			valueForIFillStyles := make([]IResourceUri, len(valueForFillStyles))
			for i, v := range valueForFillStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFillStyles[i] = vObject.(IResourceUri)
				}
			}
			this.FillStyles = valueForIFillStyles
		}
	}
	
	if valLineStyles, ok := objMap["lineStyles"]; ok {
		if valLineStyles != nil {
			var valueForLineStyles []json.RawMessage
			err = json.Unmarshal(*valLineStyles, &valueForLineStyles)
			if err != nil {
				return err
			}
			valueForILineStyles := make([]IResourceUri, len(valueForLineStyles))
			for i, v := range valueForLineStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForILineStyles[i] = vObject.(IResourceUri)
				}
			}
			this.LineStyles = valueForILineStyles
		}
	}
	if valLineStylesCap, ok := objMap["LineStyles"]; ok {
		if valLineStylesCap != nil {
			var valueForLineStyles []json.RawMessage
			err = json.Unmarshal(*valLineStylesCap, &valueForLineStyles)
			if err != nil {
				return err
			}
			valueForILineStyles := make([]IResourceUri, len(valueForLineStyles))
			for i, v := range valueForLineStyles {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForILineStyles[i] = vObject.(IResourceUri)
				}
			}
			this.LineStyles = valueForILineStyles
		}
	}

	return nil
}
