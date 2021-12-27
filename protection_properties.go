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

	// Password for read protection.
	getReadPassword() string
	setReadPassword(newValue string)

	// Password for write protection.
	getWritePassword() string
	setWritePassword(newValue string)

	// Returns true if the presentation protected for editing. 
	getIsWriteProtected() bool
	setIsWriteProtected(newValue bool)

	// Returns true if the presentation protected for reading. 
	getIsEncrypted() bool
	setIsEncrypted(newValue bool)
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

	// Password for read protection.
	ReadPassword string `json:"ReadPassword,omitempty"`

	// Password for write protection.
	WritePassword string `json:"WritePassword,omitempty"`

	// Returns true if the presentation protected for editing. 
	IsWriteProtected bool `json:"IsWriteProtected"`

	// Returns true if the presentation protected for reading. 
	IsEncrypted bool `json:"IsEncrypted"`
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
func (this *ProtectionProperties) getReadPassword() string {
	return this.ReadPassword
}

func (this *ProtectionProperties) setReadPassword(newValue string) {
	this.ReadPassword = newValue
}
func (this *ProtectionProperties) getWritePassword() string {
	return this.WritePassword
}

func (this *ProtectionProperties) setWritePassword(newValue string) {
	this.WritePassword = newValue
}
func (this *ProtectionProperties) getIsWriteProtected() bool {
	return this.IsWriteProtected
}

func (this *ProtectionProperties) setIsWriteProtected(newValue bool) {
	this.IsWriteProtected = newValue
}
func (this *ProtectionProperties) getIsEncrypted() bool {
	return this.IsEncrypted
}

func (this *ProtectionProperties) setIsEncrypted(newValue bool) {
	this.IsEncrypted = newValue
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
	
	if valReadPassword, ok := objMap["readPassword"]; ok {
		if valReadPassword != nil {
			var valueForReadPassword string
			err = json.Unmarshal(*valReadPassword, &valueForReadPassword)
			if err != nil {
				return err
			}
			this.ReadPassword = valueForReadPassword
		}
	}
	if valReadPasswordCap, ok := objMap["ReadPassword"]; ok {
		if valReadPasswordCap != nil {
			var valueForReadPassword string
			err = json.Unmarshal(*valReadPasswordCap, &valueForReadPassword)
			if err != nil {
				return err
			}
			this.ReadPassword = valueForReadPassword
		}
	}
	
	if valWritePassword, ok := objMap["writePassword"]; ok {
		if valWritePassword != nil {
			var valueForWritePassword string
			err = json.Unmarshal(*valWritePassword, &valueForWritePassword)
			if err != nil {
				return err
			}
			this.WritePassword = valueForWritePassword
		}
	}
	if valWritePasswordCap, ok := objMap["WritePassword"]; ok {
		if valWritePasswordCap != nil {
			var valueForWritePassword string
			err = json.Unmarshal(*valWritePasswordCap, &valueForWritePassword)
			if err != nil {
				return err
			}
			this.WritePassword = valueForWritePassword
		}
	}
	
	if valIsWriteProtected, ok := objMap["isWriteProtected"]; ok {
		if valIsWriteProtected != nil {
			var valueForIsWriteProtected bool
			err = json.Unmarshal(*valIsWriteProtected, &valueForIsWriteProtected)
			if err != nil {
				return err
			}
			this.IsWriteProtected = valueForIsWriteProtected
		}
	}
	if valIsWriteProtectedCap, ok := objMap["IsWriteProtected"]; ok {
		if valIsWriteProtectedCap != nil {
			var valueForIsWriteProtected bool
			err = json.Unmarshal(*valIsWriteProtectedCap, &valueForIsWriteProtected)
			if err != nil {
				return err
			}
			this.IsWriteProtected = valueForIsWriteProtected
		}
	}
	
	if valIsEncrypted, ok := objMap["isEncrypted"]; ok {
		if valIsEncrypted != nil {
			var valueForIsEncrypted bool
			err = json.Unmarshal(*valIsEncrypted, &valueForIsEncrypted)
			if err != nil {
				return err
			}
			this.IsEncrypted = valueForIsEncrypted
		}
	}
	if valIsEncryptedCap, ok := objMap["IsEncrypted"]; ok {
		if valIsEncryptedCap != nil {
			var valueForIsEncrypted bool
			err = json.Unmarshal(*valIsEncryptedCap, &valueForIsEncrypted)
			if err != nil {
				return err
			}
			this.IsEncrypted = valueForIsEncrypted
		}
	}

	return nil
}
