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

// Represents placeholder
type IPlaceholder interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Index.
	GetIndex() int32
	SetIndex(newValue int32)

	// Orientation.
	GetOrientation() string
	SetOrientation(newValue string)

	// Size.
	GetSize() string
	SetSize(newValue string)

	// Placeholder type.
	GetType() string
	SetType(newValue string)

	// Shape link.
	GetShape() IResourceUri
	SetShape(newValue IResourceUri)
}

type Placeholder struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Index.
	Index int32 `json:"Index"`

	// Orientation.
	Orientation string `json:"Orientation"`

	// Size.
	Size string `json:"Size"`

	// Placeholder type.
	Type_ string `json:"Type"`

	// Shape link.
	Shape IResourceUri `json:"Shape,omitempty"`
}

func NewPlaceholder() *Placeholder {
	instance := new(Placeholder)
	instance.Orientation = "Horizontal"
	instance.Size = "Full"
	instance.Type_ = "Title"
	return instance
}

func (this *Placeholder) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Placeholder) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Placeholder) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Placeholder) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Placeholder) GetIndex() int32 {
	return this.Index
}

func (this *Placeholder) SetIndex(newValue int32) {
	this.Index = newValue
}
func (this *Placeholder) GetOrientation() string {
	return this.Orientation
}

func (this *Placeholder) SetOrientation(newValue string) {
	this.Orientation = newValue
}
func (this *Placeholder) GetSize() string {
	return this.Size
}

func (this *Placeholder) SetSize(newValue string) {
	this.Size = newValue
}
func (this *Placeholder) GetType() string {
	return this.Type_
}

func (this *Placeholder) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Placeholder) GetShape() IResourceUri {
	return this.Shape
}

func (this *Placeholder) SetShape(newValue IResourceUri) {
	this.Shape = newValue
}

func (this *Placeholder) UnmarshalJSON(b []byte) error {
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
	
	if valIndex, ok := objMap["index"]; ok {
		if valIndex != nil {
			var valueForIndex int32
			err = json.Unmarshal(*valIndex, &valueForIndex)
			if err != nil {
				return err
			}
			this.Index = valueForIndex
		}
	}
	if valIndexCap, ok := objMap["Index"]; ok {
		if valIndexCap != nil {
			var valueForIndex int32
			err = json.Unmarshal(*valIndexCap, &valueForIndex)
			if err != nil {
				return err
			}
			this.Index = valueForIndex
		}
	}
	this.Orientation = "Horizontal"
	if valOrientation, ok := objMap["orientation"]; ok {
		if valOrientation != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientation, &valueForOrientation)
			if err != nil {
				var valueForOrientationInt int32
				err = json.Unmarshal(*valOrientation, &valueForOrientationInt)
				if err != nil {
					return err
				}
				this.Orientation = string(valueForOrientationInt)
			} else {
				this.Orientation = valueForOrientation
			}
		}
	}
	if valOrientationCap, ok := objMap["Orientation"]; ok {
		if valOrientationCap != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientationCap, &valueForOrientation)
			if err != nil {
				var valueForOrientationInt int32
				err = json.Unmarshal(*valOrientationCap, &valueForOrientationInt)
				if err != nil {
					return err
				}
				this.Orientation = string(valueForOrientationInt)
			} else {
				this.Orientation = valueForOrientation
			}
		}
	}
	this.Size = "Full"
	if valSize, ok := objMap["size"]; ok {
		if valSize != nil {
			var valueForSize string
			err = json.Unmarshal(*valSize, &valueForSize)
			if err != nil {
				var valueForSizeInt int32
				err = json.Unmarshal(*valSize, &valueForSizeInt)
				if err != nil {
					return err
				}
				this.Size = string(valueForSizeInt)
			} else {
				this.Size = valueForSize
			}
		}
	}
	if valSizeCap, ok := objMap["Size"]; ok {
		if valSizeCap != nil {
			var valueForSize string
			err = json.Unmarshal(*valSizeCap, &valueForSize)
			if err != nil {
				var valueForSizeInt int32
				err = json.Unmarshal(*valSizeCap, &valueForSizeInt)
				if err != nil {
					return err
				}
				this.Size = string(valueForSizeInt)
			} else {
				this.Size = valueForSize
			}
		}
	}
	this.Type_ = "Title"
	if valType, ok := objMap["type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valShape, ok := objMap["shape"]; ok {
		if valShape != nil {
			var valueForShape ResourceUri
			err = json.Unmarshal(*valShape, &valueForShape)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valShape)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShape, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Shape = vInterfaceObject
			}
		}
	}
	if valShapeCap, ok := objMap["Shape"]; ok {
		if valShapeCap != nil {
			var valueForShape ResourceUri
			err = json.Unmarshal(*valShapeCap, &valueForShape)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valShapeCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShapeCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Shape = vInterfaceObject
			}
		}
	}

	return nil
}
