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

// An array of elements.
type IArrayElement interface {

	// Type
	getType() string
	setType(newValue string)

	// Arguments
	getArguments() []IMathElement
	setArguments(newValue []IMathElement)

	// Specifies alignment of the array relative to surrounding text
	getBaseJustification() string
	setBaseJustification(newValue string)

	// Maximum Distribution
	getMaximumDistribution() bool
	setMaximumDistribution(newValue bool)

	// Object Distribution
	getObjectDistribution() bool
	setObjectDistribution(newValue bool)

	// The type of vertical spacing between array elements
	getRowSpacingRule() string
	setRowSpacingRule(newValue string)

	// Spacing between rows of an array
	getRowSpacing() int32
	setRowSpacing(newValue int32)
}

type ArrayElement struct {

	// Type
	Type_ string `json:"Type"`

	// Arguments
	Arguments []IMathElement `json:"Arguments,omitempty"`

	// Specifies alignment of the array relative to surrounding text
	BaseJustification string `json:"BaseJustification,omitempty"`

	// Maximum Distribution
	MaximumDistribution bool `json:"MaximumDistribution"`

	// Object Distribution
	ObjectDistribution bool `json:"ObjectDistribution"`

	// The type of vertical spacing between array elements
	RowSpacingRule string `json:"RowSpacingRule,omitempty"`

	// Spacing between rows of an array
	RowSpacing int32 `json:"RowSpacing,omitempty"`
}

func NewArrayElement() *ArrayElement {
	instance := new(ArrayElement)
	instance.Type_ = "Array"
	instance.BaseJustification = ""
	instance.RowSpacingRule = ""
	return instance
}

func (this *ArrayElement) getType() string {
	return this.Type_
}

func (this *ArrayElement) setType(newValue string) {
	this.Type_ = newValue
}
func (this *ArrayElement) getArguments() []IMathElement {
	return this.Arguments
}

func (this *ArrayElement) setArguments(newValue []IMathElement) {
	this.Arguments = newValue
}
func (this *ArrayElement) getBaseJustification() string {
	return this.BaseJustification
}

func (this *ArrayElement) setBaseJustification(newValue string) {
	this.BaseJustification = newValue
}
func (this *ArrayElement) getMaximumDistribution() bool {
	return this.MaximumDistribution
}

func (this *ArrayElement) setMaximumDistribution(newValue bool) {
	this.MaximumDistribution = newValue
}
func (this *ArrayElement) getObjectDistribution() bool {
	return this.ObjectDistribution
}

func (this *ArrayElement) setObjectDistribution(newValue bool) {
	this.ObjectDistribution = newValue
}
func (this *ArrayElement) getRowSpacingRule() string {
	return this.RowSpacingRule
}

func (this *ArrayElement) setRowSpacingRule(newValue string) {
	this.RowSpacingRule = newValue
}
func (this *ArrayElement) getRowSpacing() int32 {
	return this.RowSpacing
}

func (this *ArrayElement) setRowSpacing(newValue int32) {
	this.RowSpacing = newValue
}

func (this *ArrayElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Array"
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
	
	if valArguments, ok := objMap["arguments"]; ok {
		if valArguments != nil {
			var valueForArguments []json.RawMessage
			err = json.Unmarshal(*valArguments, &valueForArguments)
			if err != nil {
				return err
			}
			valueForIArguments := make([]IMathElement, len(valueForArguments))
			for i, v := range valueForArguments {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIArguments[i] = vObject.(IMathElement)
				}
			}
			this.Arguments = valueForIArguments
		}
	}
	if valArgumentsCap, ok := objMap["Arguments"]; ok {
		if valArgumentsCap != nil {
			var valueForArguments []json.RawMessage
			err = json.Unmarshal(*valArgumentsCap, &valueForArguments)
			if err != nil {
				return err
			}
			valueForIArguments := make([]IMathElement, len(valueForArguments))
			for i, v := range valueForArguments {
				vObject, err := createObjectForType("MathElement", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIArguments[i] = vObject.(IMathElement)
				}
			}
			this.Arguments = valueForIArguments
		}
	}
	this.BaseJustification = ""
	if valBaseJustification, ok := objMap["baseJustification"]; ok {
		if valBaseJustification != nil {
			var valueForBaseJustification string
			err = json.Unmarshal(*valBaseJustification, &valueForBaseJustification)
			if err != nil {
				var valueForBaseJustificationInt int32
				err = json.Unmarshal(*valBaseJustification, &valueForBaseJustificationInt)
				if err != nil {
					return err
				}
				this.BaseJustification = string(valueForBaseJustificationInt)
			} else {
				this.BaseJustification = valueForBaseJustification
			}
		}
	}
	if valBaseJustificationCap, ok := objMap["BaseJustification"]; ok {
		if valBaseJustificationCap != nil {
			var valueForBaseJustification string
			err = json.Unmarshal(*valBaseJustificationCap, &valueForBaseJustification)
			if err != nil {
				var valueForBaseJustificationInt int32
				err = json.Unmarshal(*valBaseJustificationCap, &valueForBaseJustificationInt)
				if err != nil {
					return err
				}
				this.BaseJustification = string(valueForBaseJustificationInt)
			} else {
				this.BaseJustification = valueForBaseJustification
			}
		}
	}
	
	if valMaximumDistribution, ok := objMap["maximumDistribution"]; ok {
		if valMaximumDistribution != nil {
			var valueForMaximumDistribution bool
			err = json.Unmarshal(*valMaximumDistribution, &valueForMaximumDistribution)
			if err != nil {
				return err
			}
			this.MaximumDistribution = valueForMaximumDistribution
		}
	}
	if valMaximumDistributionCap, ok := objMap["MaximumDistribution"]; ok {
		if valMaximumDistributionCap != nil {
			var valueForMaximumDistribution bool
			err = json.Unmarshal(*valMaximumDistributionCap, &valueForMaximumDistribution)
			if err != nil {
				return err
			}
			this.MaximumDistribution = valueForMaximumDistribution
		}
	}
	
	if valObjectDistribution, ok := objMap["objectDistribution"]; ok {
		if valObjectDistribution != nil {
			var valueForObjectDistribution bool
			err = json.Unmarshal(*valObjectDistribution, &valueForObjectDistribution)
			if err != nil {
				return err
			}
			this.ObjectDistribution = valueForObjectDistribution
		}
	}
	if valObjectDistributionCap, ok := objMap["ObjectDistribution"]; ok {
		if valObjectDistributionCap != nil {
			var valueForObjectDistribution bool
			err = json.Unmarshal(*valObjectDistributionCap, &valueForObjectDistribution)
			if err != nil {
				return err
			}
			this.ObjectDistribution = valueForObjectDistribution
		}
	}
	this.RowSpacingRule = ""
	if valRowSpacingRule, ok := objMap["rowSpacingRule"]; ok {
		if valRowSpacingRule != nil {
			var valueForRowSpacingRule string
			err = json.Unmarshal(*valRowSpacingRule, &valueForRowSpacingRule)
			if err != nil {
				var valueForRowSpacingRuleInt int32
				err = json.Unmarshal(*valRowSpacingRule, &valueForRowSpacingRuleInt)
				if err != nil {
					return err
				}
				this.RowSpacingRule = string(valueForRowSpacingRuleInt)
			} else {
				this.RowSpacingRule = valueForRowSpacingRule
			}
		}
	}
	if valRowSpacingRuleCap, ok := objMap["RowSpacingRule"]; ok {
		if valRowSpacingRuleCap != nil {
			var valueForRowSpacingRule string
			err = json.Unmarshal(*valRowSpacingRuleCap, &valueForRowSpacingRule)
			if err != nil {
				var valueForRowSpacingRuleInt int32
				err = json.Unmarshal(*valRowSpacingRuleCap, &valueForRowSpacingRuleInt)
				if err != nil {
					return err
				}
				this.RowSpacingRule = string(valueForRowSpacingRuleInt)
			} else {
				this.RowSpacingRule = valueForRowSpacingRule
			}
		}
	}
	
	if valRowSpacing, ok := objMap["rowSpacing"]; ok {
		if valRowSpacing != nil {
			var valueForRowSpacing int32
			err = json.Unmarshal(*valRowSpacing, &valueForRowSpacing)
			if err != nil {
				return err
			}
			this.RowSpacing = valueForRowSpacing
		}
	}
	if valRowSpacingCap, ok := objMap["RowSpacing"]; ok {
		if valRowSpacingCap != nil {
			var valueForRowSpacing int32
			err = json.Unmarshal(*valRowSpacingCap, &valueForRowSpacing)
			if err != nil {
				return err
			}
			this.RowSpacing = valueForRowSpacing
		}
	}

	return nil
}
