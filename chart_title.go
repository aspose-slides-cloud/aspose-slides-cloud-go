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

// Represents chart title
type IChartTitle interface {

	// Get or sets the text.
	getText() string
	setText(newValue string)

	// Get or sets value determines visibility of title
	getHasTitle() bool
	setHasTitle(newValue bool)
}

type ChartTitle struct {

	// Get or sets the text.
	Text string `json:"Text,omitempty"`

	// Get or sets value determines visibility of title
	HasTitle bool `json:"HasTitle"`
}

func NewChartTitle() *ChartTitle {
	instance := new(ChartTitle)
	return instance
}

func (this *ChartTitle) getText() string {
	return this.Text
}

func (this *ChartTitle) setText(newValue string) {
	this.Text = newValue
}
func (this *ChartTitle) getHasTitle() bool {
	return this.HasTitle
}

func (this *ChartTitle) setHasTitle(newValue bool) {
	this.HasTitle = newValue
}

func (this *ChartTitle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
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
	
	if valHasTitle, ok := objMap["hasTitle"]; ok {
		if valHasTitle != nil {
			var valueForHasTitle bool
			err = json.Unmarshal(*valHasTitle, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}
	if valHasTitleCap, ok := objMap["HasTitle"]; ok {
		if valHasTitleCap != nil {
			var valueForHasTitle bool
			err = json.Unmarshal(*valHasTitleCap, &valueForHasTitle)
			if err != nil {
				return err
			}
			this.HasTitle = valueForHasTitle
		}
	}

    return nil
}
