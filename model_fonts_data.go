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

// List of fonts data
type IFontsData interface {

	// Fonts data list.
	GetList() []IFontData
	SetList(newValue []IFontData)
}

type FontsData struct {

	// Fonts data list.
	List []IFontData `json:"List,omitempty"`
}

func NewFontsData() *FontsData {
	instance := new(FontsData)
	return instance
}

func (this *FontsData) GetList() []IFontData {
	return this.List
}

func (this *FontsData) SetList(newValue []IFontData) {
	this.List = newValue
}

func (this *FontsData) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valList, ok := GetMapValue(objMap, "list"); ok {
		if valList != nil {
			var valueForList []json.RawMessage
			err = json.Unmarshal(*valList, &valueForList)
			if err != nil {
				return err
			}
			valueForIList := make([]IFontData, len(valueForList))
			for i, v := range valueForList {
				vObject, err := createObjectForType("FontData", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIList[i] = vObject.(IFontData)
				}
			}
			this.List = valueForIList
		}
	}

	return nil
}
