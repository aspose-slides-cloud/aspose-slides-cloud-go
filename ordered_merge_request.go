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

// Request for presentations merge with optional order of slides
type IOrderedMergeRequest interface {

	// Gets or sets the presentation paths.
	getPresentations() []PresentationToMerge
	setPresentations(newValue []PresentationToMerge)
}

type OrderedMergeRequest struct {

	// Gets or sets the presentation paths.
	Presentations []PresentationToMerge `json:"Presentations,omitempty"`
}

func (this OrderedMergeRequest) getPresentations() []PresentationToMerge {
	return this.Presentations
}

func (this OrderedMergeRequest) setPresentations(newValue []PresentationToMerge) {
	this.Presentations = newValue
}

func (this *OrderedMergeRequest) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valPresentations, ok := objMap["Presentations"]; ok {
		if valPresentations != nil {
			var valueForPresentations []PresentationToMerge
			err = json.Unmarshal(*valPresentations, &valueForPresentations)
			if err != nil {
				return err
			}
			this.Presentations = valueForPresentations
		}
	}

    return nil
}
