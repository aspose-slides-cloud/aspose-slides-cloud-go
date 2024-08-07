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

// A set of properties specifying which access permissions should be granted when the document is opened with user access.
type IAccessPermissions interface {

	// The user may print the document (possibly not at the highest quality level, depending on whether bit HighQualityPrint is also set).
	GetPrintDocument() *bool
	SetPrintDocument(newValue *bool)

	// The user may modify the contents of the document by operations other than those controlled by bits AddOrModifyFields, FillExistingFields, AssembleDocument.
	GetModifyContent() *bool
	SetModifyContent(newValue *bool)

	// The user may copy or otherwise extract text and graphics from the document by operations other than that controlled by bit ExtractTextAndGraphics.
	GetCopyTextAndGraphics() *bool
	SetCopyTextAndGraphics(newValue *bool)

	// The user may add or modify text annotations, fill in interactive form fields, and, if bit ModifyContent is also set, create or modify interactive form fields (including signature fields).
	GetAddOrModifyFields() *bool
	SetAddOrModifyFields(newValue *bool)

	// The user may fill in existing interactive form fields (including signature fields), even if bit AddOrModifyFields is clear.
	GetFillExistingFields() *bool
	SetFillExistingFields(newValue *bool)

	// The user may extract text and graphics in support of accessibility to users with disabilities or for other purposes.
	GetExtractTextAndGraphics() *bool
	SetExtractTextAndGraphics(newValue *bool)

	// The user may assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if bit ModifyContent is clear.
	GetAssembleDocument() *bool
	SetAssembleDocument(newValue *bool)

	// The user may print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and bit PrintDocument is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
	GetHighQualityPrint() *bool
	SetHighQualityPrint(newValue *bool)
}

type AccessPermissions struct {

	// The user may print the document (possibly not at the highest quality level, depending on whether bit HighQualityPrint is also set).
	PrintDocument *bool `json:"PrintDocument"`

	// The user may modify the contents of the document by operations other than those controlled by bits AddOrModifyFields, FillExistingFields, AssembleDocument.
	ModifyContent *bool `json:"ModifyContent"`

	// The user may copy or otherwise extract text and graphics from the document by operations other than that controlled by bit ExtractTextAndGraphics.
	CopyTextAndGraphics *bool `json:"CopyTextAndGraphics"`

	// The user may add or modify text annotations, fill in interactive form fields, and, if bit ModifyContent is also set, create or modify interactive form fields (including signature fields).
	AddOrModifyFields *bool `json:"AddOrModifyFields"`

	// The user may fill in existing interactive form fields (including signature fields), even if bit AddOrModifyFields is clear.
	FillExistingFields *bool `json:"FillExistingFields"`

	// The user may extract text and graphics in support of accessibility to users with disabilities or for other purposes.
	ExtractTextAndGraphics *bool `json:"ExtractTextAndGraphics"`

	// The user may assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if bit ModifyContent is clear.
	AssembleDocument *bool `json:"AssembleDocument"`

	// The user may print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and bit PrintDocument is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
	HighQualityPrint *bool `json:"HighQualityPrint"`
}

func NewAccessPermissions() *AccessPermissions {
	instance := new(AccessPermissions)
	return instance
}

func (this *AccessPermissions) GetPrintDocument() *bool {
	return this.PrintDocument
}

func (this *AccessPermissions) SetPrintDocument(newValue *bool) {
	this.PrintDocument = newValue
}
func (this *AccessPermissions) GetModifyContent() *bool {
	return this.ModifyContent
}

func (this *AccessPermissions) SetModifyContent(newValue *bool) {
	this.ModifyContent = newValue
}
func (this *AccessPermissions) GetCopyTextAndGraphics() *bool {
	return this.CopyTextAndGraphics
}

func (this *AccessPermissions) SetCopyTextAndGraphics(newValue *bool) {
	this.CopyTextAndGraphics = newValue
}
func (this *AccessPermissions) GetAddOrModifyFields() *bool {
	return this.AddOrModifyFields
}

func (this *AccessPermissions) SetAddOrModifyFields(newValue *bool) {
	this.AddOrModifyFields = newValue
}
func (this *AccessPermissions) GetFillExistingFields() *bool {
	return this.FillExistingFields
}

func (this *AccessPermissions) SetFillExistingFields(newValue *bool) {
	this.FillExistingFields = newValue
}
func (this *AccessPermissions) GetExtractTextAndGraphics() *bool {
	return this.ExtractTextAndGraphics
}

func (this *AccessPermissions) SetExtractTextAndGraphics(newValue *bool) {
	this.ExtractTextAndGraphics = newValue
}
func (this *AccessPermissions) GetAssembleDocument() *bool {
	return this.AssembleDocument
}

func (this *AccessPermissions) SetAssembleDocument(newValue *bool) {
	this.AssembleDocument = newValue
}
func (this *AccessPermissions) GetHighQualityPrint() *bool {
	return this.HighQualityPrint
}

func (this *AccessPermissions) SetHighQualityPrint(newValue *bool) {
	this.HighQualityPrint = newValue
}

func (this *AccessPermissions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPrintDocument, ok := GetMapValue(objMap, "printDocument"); ok {
		if valPrintDocument != nil {
			var valueForPrintDocument *bool
			err = json.Unmarshal(*valPrintDocument, &valueForPrintDocument)
			if err != nil {
				return err
			}
			this.PrintDocument = valueForPrintDocument
		}
	}
	
	if valModifyContent, ok := GetMapValue(objMap, "modifyContent"); ok {
		if valModifyContent != nil {
			var valueForModifyContent *bool
			err = json.Unmarshal(*valModifyContent, &valueForModifyContent)
			if err != nil {
				return err
			}
			this.ModifyContent = valueForModifyContent
		}
	}
	
	if valCopyTextAndGraphics, ok := GetMapValue(objMap, "copyTextAndGraphics"); ok {
		if valCopyTextAndGraphics != nil {
			var valueForCopyTextAndGraphics *bool
			err = json.Unmarshal(*valCopyTextAndGraphics, &valueForCopyTextAndGraphics)
			if err != nil {
				return err
			}
			this.CopyTextAndGraphics = valueForCopyTextAndGraphics
		}
	}
	
	if valAddOrModifyFields, ok := GetMapValue(objMap, "addOrModifyFields"); ok {
		if valAddOrModifyFields != nil {
			var valueForAddOrModifyFields *bool
			err = json.Unmarshal(*valAddOrModifyFields, &valueForAddOrModifyFields)
			if err != nil {
				return err
			}
			this.AddOrModifyFields = valueForAddOrModifyFields
		}
	}
	
	if valFillExistingFields, ok := GetMapValue(objMap, "fillExistingFields"); ok {
		if valFillExistingFields != nil {
			var valueForFillExistingFields *bool
			err = json.Unmarshal(*valFillExistingFields, &valueForFillExistingFields)
			if err != nil {
				return err
			}
			this.FillExistingFields = valueForFillExistingFields
		}
	}
	
	if valExtractTextAndGraphics, ok := GetMapValue(objMap, "extractTextAndGraphics"); ok {
		if valExtractTextAndGraphics != nil {
			var valueForExtractTextAndGraphics *bool
			err = json.Unmarshal(*valExtractTextAndGraphics, &valueForExtractTextAndGraphics)
			if err != nil {
				return err
			}
			this.ExtractTextAndGraphics = valueForExtractTextAndGraphics
		}
	}
	
	if valAssembleDocument, ok := GetMapValue(objMap, "assembleDocument"); ok {
		if valAssembleDocument != nil {
			var valueForAssembleDocument *bool
			err = json.Unmarshal(*valAssembleDocument, &valueForAssembleDocument)
			if err != nil {
				return err
			}
			this.AssembleDocument = valueForAssembleDocument
		}
	}
	
	if valHighQualityPrint, ok := GetMapValue(objMap, "highQualityPrint"); ok {
		if valHighQualityPrint != nil {
			var valueForHighQualityPrint *bool
			err = json.Unmarshal(*valHighQualityPrint, &valueForHighQualityPrint)
			if err != nil {
				return err
			}
			this.HighQualityPrint = valueForHighQualityPrint
		}
	}

	return nil
}
