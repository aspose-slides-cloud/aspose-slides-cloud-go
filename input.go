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

// Represents input document for pipeline.
type IInput interface {

	// Get or sets template document. If property is null new empty presentation will be created.
	getTemplate() IInputFile
	setTemplate(newValue IInputFile)

	// Get or sets html data for generate new presentation.
	getHtmlData() IInputFile
	setHtmlData(newValue IInputFile)

	// Get or sets data for template engine.
	getTemplateData() IInputFile
	setTemplateData(newValue IInputFile)
}

type Input struct {

	// Get or sets template document. If property is null new empty presentation will be created.
	Template IInputFile `json:"Template,omitempty"`

	// Get or sets html data for generate new presentation.
	HtmlData IInputFile `json:"HtmlData,omitempty"`

	// Get or sets data for template engine.
	TemplateData IInputFile `json:"TemplateData,omitempty"`
}

func NewInput() *Input {
	instance := new(Input)
	return instance
}

func (this *Input) getTemplate() IInputFile {
	return this.Template
}

func (this *Input) setTemplate(newValue IInputFile) {
	this.Template = newValue
}
func (this *Input) getHtmlData() IInputFile {
	return this.HtmlData
}

func (this *Input) setHtmlData(newValue IInputFile) {
	this.HtmlData = newValue
}
func (this *Input) getTemplateData() IInputFile {
	return this.TemplateData
}

func (this *Input) setTemplateData(newValue IInputFile) {
	this.TemplateData = newValue
}

func (this *Input) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valTemplate, ok := objMap["template"]; ok {
		if valTemplate != nil {
			var valueForTemplate InputFile
			err = json.Unmarshal(*valTemplate, &valueForTemplate)
			if err != nil {
				return err
			}
			this.Template = &valueForTemplate
		}
	}
	if valTemplateCap, ok := objMap["Template"]; ok {
		if valTemplateCap != nil {
			var valueForTemplate InputFile
			err = json.Unmarshal(*valTemplateCap, &valueForTemplate)
			if err != nil {
				return err
			}
			this.Template = &valueForTemplate
		}
	}
	
	if valHtmlData, ok := objMap["htmlData"]; ok {
		if valHtmlData != nil {
			var valueForHtmlData InputFile
			err = json.Unmarshal(*valHtmlData, &valueForHtmlData)
			if err != nil {
				return err
			}
			this.HtmlData = &valueForHtmlData
		}
	}
	if valHtmlDataCap, ok := objMap["HtmlData"]; ok {
		if valHtmlDataCap != nil {
			var valueForHtmlData InputFile
			err = json.Unmarshal(*valHtmlDataCap, &valueForHtmlData)
			if err != nil {
				return err
			}
			this.HtmlData = &valueForHtmlData
		}
	}
	
	if valTemplateData, ok := objMap["templateData"]; ok {
		if valTemplateData != nil {
			var valueForTemplateData InputFile
			err = json.Unmarshal(*valTemplateData, &valueForTemplateData)
			if err != nil {
				return err
			}
			this.TemplateData = &valueForTemplateData
		}
	}
	if valTemplateDataCap, ok := objMap["TemplateData"]; ok {
		if valTemplateDataCap != nil {
			var valueForTemplateData InputFile
			err = json.Unmarshal(*valTemplateDataCap, &valueForTemplateData)
			if err != nil {
				return err
			}
			this.TemplateData = &valueForTemplateData
		}
	}

    return nil
}
