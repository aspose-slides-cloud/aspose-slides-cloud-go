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

// Protection properties.
type IProtectionProperties interface {

	// Gets or sets the link to this resource.
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// True if document properties are encrypted. Has effect only for password protected presentations.
	getEncryptDocumentProperties() bool
	setEncryptDocumentProperties(newValue bool)

	// True if the document should be opened as read-only.
	getReadOnlyRecommended() bool
	setReadOnlyRecommended(newValue bool)
}

type ProtectionProperties struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// True if document properties are encrypted. Has effect only for password protected presentations.
	EncryptDocumentProperties bool `json:"EncryptDocumentProperties"`

	// True if the document should be opened as read-only.
	ReadOnlyRecommended bool `json:"ReadOnlyRecommended"`
}

func NewProtectionProperties() *ProtectionProperties {
	instance := new(ProtectionProperties)
	return instance
}

func (this *ProtectionProperties) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ProtectionProperties) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ProtectionProperties) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ProtectionProperties) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ProtectionProperties) getEncryptDocumentProperties() bool {
	return this.EncryptDocumentProperties
}

func (this *ProtectionProperties) setEncryptDocumentProperties(newValue bool) {
	this.EncryptDocumentProperties = newValue
}
func (this *ProtectionProperties) getReadOnlyRecommended() bool {
	return this.ReadOnlyRecommended
}

func (this *ProtectionProperties) setReadOnlyRecommended(newValue bool) {
	this.ReadOnlyRecommended = newValue
}

func (this *ProtectionProperties) UnmarshalJSON(b []byte) error {
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
	
	if valEncryptDocumentProperties, ok := objMap["encryptDocumentProperties"]; ok {
		if valEncryptDocumentProperties != nil {
			var valueForEncryptDocumentProperties bool
			err = json.Unmarshal(*valEncryptDocumentProperties, &valueForEncryptDocumentProperties)
			if err != nil {
				return err
			}
			this.EncryptDocumentProperties = valueForEncryptDocumentProperties
		}
	}
	if valEncryptDocumentPropertiesCap, ok := objMap["EncryptDocumentProperties"]; ok {
		if valEncryptDocumentPropertiesCap != nil {
			var valueForEncryptDocumentProperties bool
			err = json.Unmarshal(*valEncryptDocumentPropertiesCap, &valueForEncryptDocumentProperties)
			if err != nil {
				return err
			}
			this.EncryptDocumentProperties = valueForEncryptDocumentProperties
		}
	}
	
	if valReadOnlyRecommended, ok := objMap["readOnlyRecommended"]; ok {
		if valReadOnlyRecommended != nil {
			var valueForReadOnlyRecommended bool
			err = json.Unmarshal(*valReadOnlyRecommended, &valueForReadOnlyRecommended)
			if err != nil {
				return err
			}
			this.ReadOnlyRecommended = valueForReadOnlyRecommended
		}
	}
	if valReadOnlyRecommendedCap, ok := objMap["ReadOnlyRecommended"]; ok {
		if valReadOnlyRecommendedCap != nil {
			var valueForReadOnlyRecommended bool
			err = json.Unmarshal(*valReadOnlyRecommendedCap, &valueForReadOnlyRecommended)
			if err != nil {
				return err
			}
			this.ReadOnlyRecommended = valueForReadOnlyRecommended
		}
	}

    return nil
}
