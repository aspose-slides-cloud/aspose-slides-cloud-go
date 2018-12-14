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

	getAlternateLinks() []ResourceUri
	setAlternateLinks(newValue []ResourceUri)

	// A list of links that originate from this document.
	getLinks() []ResourceUri
	setLinks(newValue []ResourceUri)

	getBackgroundStyles() []ResourceUri
	setBackgroundStyles(newValue []ResourceUri)

	getEffectStyles() []ResourceUri
	setEffectStyles(newValue []ResourceUri)

	getFillStyles() []ResourceUri
	setFillStyles(newValue []ResourceUri)

	getLineStyles() []ResourceUri
	setLineStyles(newValue []ResourceUri)
}

type FormatScheme struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	BackgroundStyles []ResourceUri `json:"BackgroundStyles,omitempty"`

	EffectStyles []ResourceUri `json:"EffectStyles,omitempty"`

	FillStyles []ResourceUri `json:"FillStyles,omitempty"`

	LineStyles []ResourceUri `json:"LineStyles,omitempty"`
}

func (this FormatScheme) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this FormatScheme) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this FormatScheme) getAlternateLinks() []ResourceUri {
	return this.AlternateLinks
}

func (this FormatScheme) setAlternateLinks(newValue []ResourceUri) {
	this.AlternateLinks = newValue
}
func (this FormatScheme) getLinks() []ResourceUri {
	return this.Links
}

func (this FormatScheme) setLinks(newValue []ResourceUri) {
	this.Links = newValue
}
func (this FormatScheme) getBackgroundStyles() []ResourceUri {
	return this.BackgroundStyles
}

func (this FormatScheme) setBackgroundStyles(newValue []ResourceUri) {
	this.BackgroundStyles = newValue
}
func (this FormatScheme) getEffectStyles() []ResourceUri {
	return this.EffectStyles
}

func (this FormatScheme) setEffectStyles(newValue []ResourceUri) {
	this.EffectStyles = newValue
}
func (this FormatScheme) getFillStyles() []ResourceUri {
	return this.FillStyles
}

func (this FormatScheme) setFillStyles(newValue []ResourceUri) {
	this.FillStyles = newValue
}
func (this FormatScheme) getLineStyles() []ResourceUri {
	return this.LineStyles
}

func (this FormatScheme) setLineStyles(newValue []ResourceUri) {
	this.LineStyles = newValue
}

func (this *FormatScheme) UnmarshalJSON(b []byte) error {
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

	if valLinks, ok := objMap["Links"]; ok {
		if valLinks != nil {
			var valueForLinks []ResourceUri
			err = json.Unmarshal(*valLinks, &valueForLinks)
			if err != nil {
				return err
			}
			this.Links = valueForLinks
		}
	}

	if valBackgroundStyles, ok := objMap["BackgroundStyles"]; ok {
		if valBackgroundStyles != nil {
			var valueForBackgroundStyles []ResourceUri
			err = json.Unmarshal(*valBackgroundStyles, &valueForBackgroundStyles)
			if err != nil {
				return err
			}
			this.BackgroundStyles = valueForBackgroundStyles
		}
	}

	if valEffectStyles, ok := objMap["EffectStyles"]; ok {
		if valEffectStyles != nil {
			var valueForEffectStyles []ResourceUri
			err = json.Unmarshal(*valEffectStyles, &valueForEffectStyles)
			if err != nil {
				return err
			}
			this.EffectStyles = valueForEffectStyles
		}
	}

	if valFillStyles, ok := objMap["FillStyles"]; ok {
		if valFillStyles != nil {
			var valueForFillStyles []ResourceUri
			err = json.Unmarshal(*valFillStyles, &valueForFillStyles)
			if err != nil {
				return err
			}
			this.FillStyles = valueForFillStyles
		}
	}

	if valLineStyles, ok := objMap["LineStyles"]; ok {
		if valLineStyles != nil {
			var valueForLineStyles []ResourceUri
			err = json.Unmarshal(*valLineStyles, &valueForLineStyles)
			if err != nil {
				return err
			}
			this.LineStyles = valueForLineStyles
		}
	}

    return nil
}
