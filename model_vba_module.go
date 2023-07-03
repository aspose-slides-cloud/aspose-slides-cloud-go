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

// Represents VBA module.
type IVbaModule interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// VBA module name. 
	GetName() string
	SetName(newValue string)

	// VBA module source code.
	GetSourceCode() string
	SetSourceCode(newValue string)

	// List of references. 
	GetReferences() []IVbaReference
	SetReferences(newValue []IVbaReference)
}

type VbaModule struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// VBA module name. 
	Name string `json:"Name,omitempty"`

	// VBA module source code.
	SourceCode string `json:"SourceCode,omitempty"`

	// List of references. 
	References []IVbaReference `json:"References,omitempty"`
}

func NewVbaModule() *VbaModule {
	instance := new(VbaModule)
	return instance
}

func (this *VbaModule) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *VbaModule) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *VbaModule) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *VbaModule) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *VbaModule) GetName() string {
	return this.Name
}

func (this *VbaModule) SetName(newValue string) {
	this.Name = newValue
}
func (this *VbaModule) GetSourceCode() string {
	return this.SourceCode
}

func (this *VbaModule) SetSourceCode(newValue string) {
	this.SourceCode = newValue
}
func (this *VbaModule) GetReferences() []IVbaReference {
	return this.References
}

func (this *VbaModule) SetReferences(newValue []IVbaReference) {
	this.References = newValue
}

func (this *VbaModule) UnmarshalJSON(b []byte) error {
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
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valSourceCode, ok := objMap["sourceCode"]; ok {
		if valSourceCode != nil {
			var valueForSourceCode string
			err = json.Unmarshal(*valSourceCode, &valueForSourceCode)
			if err != nil {
				return err
			}
			this.SourceCode = valueForSourceCode
		}
	}
	if valSourceCodeCap, ok := objMap["SourceCode"]; ok {
		if valSourceCodeCap != nil {
			var valueForSourceCode string
			err = json.Unmarshal(*valSourceCodeCap, &valueForSourceCode)
			if err != nil {
				return err
			}
			this.SourceCode = valueForSourceCode
		}
	}
	
	if valReferences, ok := objMap["references"]; ok {
		if valReferences != nil {
			var valueForReferences []json.RawMessage
			err = json.Unmarshal(*valReferences, &valueForReferences)
			if err != nil {
				return err
			}
			valueForIReferences := make([]IVbaReference, len(valueForReferences))
			for i, v := range valueForReferences {
				vObject, err := createObjectForType("VbaReference", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIReferences[i] = vObject.(IVbaReference)
				}
			}
			this.References = valueForIReferences
		}
	}
	if valReferencesCap, ok := objMap["References"]; ok {
		if valReferencesCap != nil {
			var valueForReferences []json.RawMessage
			err = json.Unmarshal(*valReferencesCap, &valueForReferences)
			if err != nil {
				return err
			}
			valueForIReferences := make([]IVbaReference, len(valueForReferences))
			for i, v := range valueForReferences {
				vObject, err := createObjectForType("VbaReference", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIReferences[i] = vObject.(IVbaReference)
				}
			}
			this.References = valueForIReferences
		}
	}

	return nil
}
