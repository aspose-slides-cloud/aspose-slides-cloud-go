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

// Represents text item, referenced by TextItems
type ITextItem interface {

	// Gets or sets the URI to resource that contains text
	GetUri() IResourceUri
	SetUri(newValue IResourceUri)

	// Gets or sets the text.
	GetText() string
	SetText(newValue string)
}

type TextItem struct {

	// Gets or sets the URI to resource that contains text
	Uri IResourceUri `json:"Uri,omitempty"`

	// Gets or sets the text.
	Text string `json:"Text,omitempty"`
}

func NewTextItem() *TextItem {
	instance := new(TextItem)
	return instance
}

func (this *TextItem) GetUri() IResourceUri {
	return this.Uri
}

func (this *TextItem) SetUri(newValue IResourceUri) {
	this.Uri = newValue
}
func (this *TextItem) GetText() string {
	return this.Text
}

func (this *TextItem) SetText(newValue string) {
	this.Text = newValue
}

func (this *TextItem) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valUri, ok := objMap["uri"]; ok {
		if valUri != nil {
			var valueForUri ResourceUri
			err = json.Unmarshal(*valUri, &valueForUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Uri = vInterfaceObject
			}
		}
	}
	if valUriCap, ok := objMap["Uri"]; ok {
		if valUriCap != nil {
			var valueForUri ResourceUri
			err = json.Unmarshal(*valUriCap, &valueForUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Uri = vInterfaceObject
			}
		}
	}
	
	if valText, ok := objMap["text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	if valTextCap, ok := objMap["Text"]; ok {
		if valTextCap != nil {
			var valueForText string
			err = json.Unmarshal(*valTextCap, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}

	return nil
}
