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

// VBA project
type IVbaProject interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// VBA modules
	GetModules() []IResourceUri
	SetModules(newValue []IResourceUri)

	// VBA references
	GetReferences() []IVbaReference
	SetReferences(newValue []IVbaReference)
}

type VbaProject struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// VBA modules
	Modules []IResourceUri `json:"Modules,omitempty"`

	// VBA references
	References []IVbaReference `json:"References,omitempty"`
}

func NewVbaProject() *VbaProject {
	instance := new(VbaProject)
	return instance
}

func (this *VbaProject) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *VbaProject) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *VbaProject) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *VbaProject) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *VbaProject) GetModules() []IResourceUri {
	return this.Modules
}

func (this *VbaProject) SetModules(newValue []IResourceUri) {
	this.Modules = newValue
}
func (this *VbaProject) GetReferences() []IVbaReference {
	return this.References
}

func (this *VbaProject) SetReferences(newValue []IVbaReference) {
	this.References = newValue
}

func (this *VbaProject) UnmarshalJSON(b []byte) error {
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
	
	if valModules, ok := GetMapValue(objMap, "modules"); ok {
		if valModules != nil {
			var valueForModules []json.RawMessage
			err = json.Unmarshal(*valModules, &valueForModules)
			if err != nil {
				return err
			}
			valueForIModules := make([]IResourceUri, len(valueForModules))
			for i, v := range valueForModules {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIModules[i] = vObject.(IResourceUri)
				}
			}
			this.Modules = valueForIModules
		}
	}
	
	if valReferences, ok := GetMapValue(objMap, "references"); ok {
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

	return nil
}
