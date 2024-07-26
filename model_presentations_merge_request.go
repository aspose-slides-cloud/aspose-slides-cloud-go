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

// Request for presentations merge
type IPresentationsMergeRequest interface {

	// Gets or sets the presentation paths.
	GetPresentationPaths() []string
	SetPresentationPaths(newValue []string)

	// Gets or sets the presentation passwords.
	GetPresentationPasswords() []string
	SetPresentationPasswords(newValue []string)
}

type PresentationsMergeRequest struct {

	// Gets or sets the presentation paths.
	PresentationPaths []string `json:"PresentationPaths,omitempty"`

	// Gets or sets the presentation passwords.
	PresentationPasswords []string `json:"PresentationPasswords,omitempty"`
}

func NewPresentationsMergeRequest() *PresentationsMergeRequest {
	instance := new(PresentationsMergeRequest)
	return instance
}

func (this *PresentationsMergeRequest) GetPresentationPaths() []string {
	return this.PresentationPaths
}

func (this *PresentationsMergeRequest) SetPresentationPaths(newValue []string) {
	this.PresentationPaths = newValue
}
func (this *PresentationsMergeRequest) GetPresentationPasswords() []string {
	return this.PresentationPasswords
}

func (this *PresentationsMergeRequest) SetPresentationPasswords(newValue []string) {
	this.PresentationPasswords = newValue
}

func (this *PresentationsMergeRequest) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPresentationPaths, ok := GetMapValue(objMap, "presentationPaths"); ok {
		if valPresentationPaths != nil {
			var valueForPresentationPaths []string
			err = json.Unmarshal(*valPresentationPaths, &valueForPresentationPaths)
			if err != nil {
				return err
			}
			this.PresentationPaths = valueForPresentationPaths
		}
	}
	
	if valPresentationPasswords, ok := GetMapValue(objMap, "presentationPasswords"); ok {
		if valPresentationPasswords != nil {
			var valueForPresentationPasswords []string
			err = json.Unmarshal(*valPresentationPasswords, &valueForPresentationPasswords)
			if err != nil {
				return err
			}
			this.PresentationPasswords = valueForPresentationPasswords
		}
	}

	return nil
}
