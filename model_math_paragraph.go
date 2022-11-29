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

// Mathematical paragraph that is a container for mathematical blocks
type IMathParagraph interface {

	// List of math blocks
	GetMathBlockList() []IBlockElement
	SetMathBlockList(newValue []IBlockElement)

	// Justification of the math paragraph
	GetJustification() string
	SetJustification(newValue string)
}

type MathParagraph struct {

	// List of math blocks
	MathBlockList []IBlockElement `json:"MathBlockList,omitempty"`

	// Justification of the math paragraph
	Justification string `json:"Justification,omitempty"`
}

func NewMathParagraph() *MathParagraph {
	instance := new(MathParagraph)
	return instance
}

func (this *MathParagraph) GetMathBlockList() []IBlockElement {
	return this.MathBlockList
}

func (this *MathParagraph) SetMathBlockList(newValue []IBlockElement) {
	this.MathBlockList = newValue
}
func (this *MathParagraph) GetJustification() string {
	return this.Justification
}

func (this *MathParagraph) SetJustification(newValue string) {
	this.Justification = newValue
}

func (this *MathParagraph) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valMathBlockList, ok := objMap["mathBlockList"]; ok {
		if valMathBlockList != nil {
			var valueForMathBlockList []json.RawMessage
			err = json.Unmarshal(*valMathBlockList, &valueForMathBlockList)
			if err != nil {
				return err
			}
			valueForIMathBlockList := make([]IBlockElement, len(valueForMathBlockList))
			for i, v := range valueForMathBlockList {
				vObject, err := createObjectForType("BlockElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIMathBlockList[i] = vObject.(IBlockElement)
				}
			}
			this.MathBlockList = valueForIMathBlockList
		}
	}
	if valMathBlockListCap, ok := objMap["MathBlockList"]; ok {
		if valMathBlockListCap != nil {
			var valueForMathBlockList []json.RawMessage
			err = json.Unmarshal(*valMathBlockListCap, &valueForMathBlockList)
			if err != nil {
				return err
			}
			valueForIMathBlockList := make([]IBlockElement, len(valueForMathBlockList))
			for i, v := range valueForMathBlockList {
				vObject, err := createObjectForType("BlockElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIMathBlockList[i] = vObject.(IBlockElement)
				}
			}
			this.MathBlockList = valueForIMathBlockList
		}
	}
	
	if valJustification, ok := objMap["justification"]; ok {
		if valJustification != nil {
			var valueForJustification string
			err = json.Unmarshal(*valJustification, &valueForJustification)
			if err != nil {
				var valueForJustificationInt int32
				err = json.Unmarshal(*valJustification, &valueForJustificationInt)
				if err != nil {
					return err
				}
				this.Justification = string(valueForJustificationInt)
			} else {
				this.Justification = valueForJustification
			}
		}
	}
	if valJustificationCap, ok := objMap["Justification"]; ok {
		if valJustificationCap != nil {
			var valueForJustification string
			err = json.Unmarshal(*valJustificationCap, &valueForJustification)
			if err != nil {
				var valueForJustificationInt int32
				err = json.Unmarshal(*valJustificationCap, &valueForJustificationInt)
				if err != nil {
					return err
				}
				this.Justification = string(valueForJustificationInt)
			} else {
				this.Justification = valueForJustification
			}
		}
	}

	return nil
}
