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

// Merge presentations task.
type IMerge interface {

	// Task type.
	GetType() string
	SetType(newValue string)

	// Information about documents and slides being merging sources.
	GetPresentations() []IMergingSource
	SetPresentations(newValue []IMergingSource)
}

type Merge struct {

	// Task type.
	Type_ string `json:"Type"`

	// Information about documents and slides being merging sources.
	Presentations []IMergingSource `json:"Presentations,omitempty"`
}

func NewMerge() *Merge {
	instance := new(Merge)
	instance.Type_ = "Merge"
	return instance
}

func (this *Merge) GetType() string {
	return this.Type_
}

func (this *Merge) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Merge) GetPresentations() []IMergingSource {
	return this.Presentations
}

func (this *Merge) SetPresentations(newValue []IMergingSource) {
	this.Presentations = newValue
}

func (this *Merge) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "Merge"
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
	
	if valPresentations, ok := objMap["presentations"]; ok {
		if valPresentations != nil {
			var valueForPresentations []json.RawMessage
			err = json.Unmarshal(*valPresentations, &valueForPresentations)
			if err != nil {
				return err
			}
			valueForIPresentations := make([]IMergingSource, len(valueForPresentations))
			for i, v := range valueForPresentations {
				vObject, err := createObjectForType("MergingSource", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPresentations[i] = vObject.(IMergingSource)
				}
			}
			this.Presentations = valueForIPresentations
		}
	}
	if valPresentationsCap, ok := objMap["Presentations"]; ok {
		if valPresentationsCap != nil {
			var valueForPresentations []json.RawMessage
			err = json.Unmarshal(*valPresentationsCap, &valueForPresentations)
			if err != nil {
				return err
			}
			valueForIPresentations := make([]IMergingSource, len(valueForPresentations))
			for i, v := range valueForPresentations {
				vObject, err := createObjectForType("MergingSource", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIPresentations[i] = vObject.(IMergingSource)
				}
			}
			this.Presentations = valueForIPresentations
		}
	}

	return nil
}
