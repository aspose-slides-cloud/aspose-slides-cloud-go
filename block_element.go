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

// Specifies an instance of mathematical text that contained within a MathParagraph and starts on its own line.
type IBlockElement interface {

	// Element type
	getType() string
	setType(newValue string)

	// List of math elements.
	getMathElementList() []IMathElement
	setMathElementList(newValue []IMathElement)
}

type BlockElement struct {

	// Element type
	Type_ string `json:"Type"`

	// List of math elements.
	MathElementList []IMathElement `json:"MathElementList,omitempty"`
}

func NewBlockElement() *BlockElement {
	instance := new(BlockElement)
	instance.Type_ = "Block"
	return instance
}

func (this *BlockElement) getType() string {
	return this.Type_
}

func (this *BlockElement) setType(newValue string) {
	this.Type_ = newValue
}
func (this *BlockElement) getMathElementList() []IMathElement {
	return this.MathElementList
}

func (this *BlockElement) setMathElementList(newValue []IMathElement) {
	this.MathElementList = newValue
}

func (this *BlockElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Block"
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
	
	if valMathElementList, ok := objMap["mathElementList"]; ok {
		if valMathElementList != nil {
			var valueForMathElementList []json.RawMessage
			err = json.Unmarshal(*valMathElementList, &valueForMathElementList)
			if err != nil {
				return err
			}
			valueForIMathElementList := make([]IMathElement, len(valueForMathElementList))
			for i, v := range valueForMathElementList {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIMathElementList[i] = vObject.(IMathElement)
				}
			}
			this.MathElementList = valueForIMathElementList
		}
	}
	if valMathElementListCap, ok := objMap["MathElementList"]; ok {
		if valMathElementListCap != nil {
			var valueForMathElementList []json.RawMessage
			err = json.Unmarshal(*valMathElementListCap, &valueForMathElementList)
			if err != nil {
				return err
			}
			valueForIMathElementList := make([]IMathElement, len(valueForMathElementList))
			for i, v := range valueForMathElementList {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIMathElementList[i] = vObject.(IMathElement)
				}
			}
			this.MathElementList = valueForIMathElementList
		}
	}

	return nil
}
