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
	getPrintDocument() bool
	setPrintDocument(newValue bool)

	// The user may modify the contents of the document by operations other than those controlled by bits AddOrModifyFields, FillExistingFields, AssembleDocument.
	getModifyContent() bool
	setModifyContent(newValue bool)

	// The user may copy or otherwise extract text and graphics from the document by operations other than that controlled by bit ExtractTextAndGraphics.
	getCopyTextAndGraphics() bool
	setCopyTextAndGraphics(newValue bool)

	// The user may add or modify text annotations, fill in interactive form fields, and, if bit ModifyContent is also set, create or modify interactive form fields (including signature fields).
	getAddOrModifyFields() bool
	setAddOrModifyFields(newValue bool)

	// The user may fill in existing interactive form fields (including signature fields), even if bit AddOrModifyFields is clear.
	getFillExistingFields() bool
	setFillExistingFields(newValue bool)

	// The user may extract text and graphics in support of accessibility to users with disabilities or for other purposes.
	getExtractTextAndGraphics() bool
	setExtractTextAndGraphics(newValue bool)

	// The user may assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if bit ModifyContent is clear.
	getAssembleDocument() bool
	setAssembleDocument(newValue bool)

	// The user may print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and bit PrintDocument is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
	getHighQualityPrint() bool
	setHighQualityPrint(newValue bool)
}

type AccessPermissions struct {

	// The user may print the document (possibly not at the highest quality level, depending on whether bit HighQualityPrint is also set).
	PrintDocument bool `json:"PrintDocument"`

	// The user may modify the contents of the document by operations other than those controlled by bits AddOrModifyFields, FillExistingFields, AssembleDocument.
	ModifyContent bool `json:"ModifyContent"`

	// The user may copy or otherwise extract text and graphics from the document by operations other than that controlled by bit ExtractTextAndGraphics.
	CopyTextAndGraphics bool `json:"CopyTextAndGraphics"`

	// The user may add or modify text annotations, fill in interactive form fields, and, if bit ModifyContent is also set, create or modify interactive form fields (including signature fields).
	AddOrModifyFields bool `json:"AddOrModifyFields"`

	// The user may fill in existing interactive form fields (including signature fields), even if bit AddOrModifyFields is clear.
	FillExistingFields bool `json:"FillExistingFields"`

	// The user may extract text and graphics in support of accessibility to users with disabilities or for other purposes.
	ExtractTextAndGraphics bool `json:"ExtractTextAndGraphics"`

	// The user may assemble the document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if bit ModifyContent is clear.
	AssembleDocument bool `json:"AssembleDocument"`

	// The user may print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and bit PrintDocument is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
	HighQualityPrint bool `json:"HighQualityPrint"`
}

func NewAccessPermissions() *AccessPermissions {
	instance := new(AccessPermissions)
	return instance
}

func (this *AccessPermissions) getPrintDocument() bool {
	return this.PrintDocument
}

func (this *AccessPermissions) setPrintDocument(newValue bool) {
	this.PrintDocument = newValue
}
func (this *AccessPermissions) getModifyContent() bool {
	return this.ModifyContent
}

func (this *AccessPermissions) setModifyContent(newValue bool) {
	this.ModifyContent = newValue
}
func (this *AccessPermissions) getCopyTextAndGraphics() bool {
	return this.CopyTextAndGraphics
}

func (this *AccessPermissions) setCopyTextAndGraphics(newValue bool) {
	this.CopyTextAndGraphics = newValue
}
func (this *AccessPermissions) getAddOrModifyFields() bool {
	return this.AddOrModifyFields
}

func (this *AccessPermissions) setAddOrModifyFields(newValue bool) {
	this.AddOrModifyFields = newValue
}
func (this *AccessPermissions) getFillExistingFields() bool {
	return this.FillExistingFields
}

func (this *AccessPermissions) setFillExistingFields(newValue bool) {
	this.FillExistingFields = newValue
}
func (this *AccessPermissions) getExtractTextAndGraphics() bool {
	return this.ExtractTextAndGraphics
}

func (this *AccessPermissions) setExtractTextAndGraphics(newValue bool) {
	this.ExtractTextAndGraphics = newValue
}
func (this *AccessPermissions) getAssembleDocument() bool {
	return this.AssembleDocument
}

func (this *AccessPermissions) setAssembleDocument(newValue bool) {
	this.AssembleDocument = newValue
}
func (this *AccessPermissions) getHighQualityPrint() bool {
	return this.HighQualityPrint
}

func (this *AccessPermissions) setHighQualityPrint(newValue bool) {
	this.HighQualityPrint = newValue
}

func (this *AccessPermissions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valPrintDocument, ok := objMap["printDocument"]; ok {
		if valPrintDocument != nil {
			var valueForPrintDocument bool
			err = json.Unmarshal(*valPrintDocument, &valueForPrintDocument)
			if err != nil {
				return err
			}
			this.PrintDocument = valueForPrintDocument
		}
	}
	if valPrintDocumentCap, ok := objMap["PrintDocument"]; ok {
		if valPrintDocumentCap != nil {
			var valueForPrintDocument bool
			err = json.Unmarshal(*valPrintDocumentCap, &valueForPrintDocument)
			if err != nil {
				return err
			}
			this.PrintDocument = valueForPrintDocument
		}
	}
	
	if valModifyContent, ok := objMap["modifyContent"]; ok {
		if valModifyContent != nil {
			var valueForModifyContent bool
			err = json.Unmarshal(*valModifyContent, &valueForModifyContent)
			if err != nil {
				return err
			}
			this.ModifyContent = valueForModifyContent
		}
	}
	if valModifyContentCap, ok := objMap["ModifyContent"]; ok {
		if valModifyContentCap != nil {
			var valueForModifyContent bool
			err = json.Unmarshal(*valModifyContentCap, &valueForModifyContent)
			if err != nil {
				return err
			}
			this.ModifyContent = valueForModifyContent
		}
	}
	
	if valCopyTextAndGraphics, ok := objMap["copyTextAndGraphics"]; ok {
		if valCopyTextAndGraphics != nil {
			var valueForCopyTextAndGraphics bool
			err = json.Unmarshal(*valCopyTextAndGraphics, &valueForCopyTextAndGraphics)
			if err != nil {
				return err
			}
			this.CopyTextAndGraphics = valueForCopyTextAndGraphics
		}
	}
	if valCopyTextAndGraphicsCap, ok := objMap["CopyTextAndGraphics"]; ok {
		if valCopyTextAndGraphicsCap != nil {
			var valueForCopyTextAndGraphics bool
			err = json.Unmarshal(*valCopyTextAndGraphicsCap, &valueForCopyTextAndGraphics)
			if err != nil {
				return err
			}
			this.CopyTextAndGraphics = valueForCopyTextAndGraphics
		}
	}
	
	if valAddOrModifyFields, ok := objMap["addOrModifyFields"]; ok {
		if valAddOrModifyFields != nil {
			var valueForAddOrModifyFields bool
			err = json.Unmarshal(*valAddOrModifyFields, &valueForAddOrModifyFields)
			if err != nil {
				return err
			}
			this.AddOrModifyFields = valueForAddOrModifyFields
		}
	}
	if valAddOrModifyFieldsCap, ok := objMap["AddOrModifyFields"]; ok {
		if valAddOrModifyFieldsCap != nil {
			var valueForAddOrModifyFields bool
			err = json.Unmarshal(*valAddOrModifyFieldsCap, &valueForAddOrModifyFields)
			if err != nil {
				return err
			}
			this.AddOrModifyFields = valueForAddOrModifyFields
		}
	}
	
	if valFillExistingFields, ok := objMap["fillExistingFields"]; ok {
		if valFillExistingFields != nil {
			var valueForFillExistingFields bool
			err = json.Unmarshal(*valFillExistingFields, &valueForFillExistingFields)
			if err != nil {
				return err
			}
			this.FillExistingFields = valueForFillExistingFields
		}
	}
	if valFillExistingFieldsCap, ok := objMap["FillExistingFields"]; ok {
		if valFillExistingFieldsCap != nil {
			var valueForFillExistingFields bool
			err = json.Unmarshal(*valFillExistingFieldsCap, &valueForFillExistingFields)
			if err != nil {
				return err
			}
			this.FillExistingFields = valueForFillExistingFields
		}
	}
	
	if valExtractTextAndGraphics, ok := objMap["extractTextAndGraphics"]; ok {
		if valExtractTextAndGraphics != nil {
			var valueForExtractTextAndGraphics bool
			err = json.Unmarshal(*valExtractTextAndGraphics, &valueForExtractTextAndGraphics)
			if err != nil {
				return err
			}
			this.ExtractTextAndGraphics = valueForExtractTextAndGraphics
		}
	}
	if valExtractTextAndGraphicsCap, ok := objMap["ExtractTextAndGraphics"]; ok {
		if valExtractTextAndGraphicsCap != nil {
			var valueForExtractTextAndGraphics bool
			err = json.Unmarshal(*valExtractTextAndGraphicsCap, &valueForExtractTextAndGraphics)
			if err != nil {
				return err
			}
			this.ExtractTextAndGraphics = valueForExtractTextAndGraphics
		}
	}
	
	if valAssembleDocument, ok := objMap["assembleDocument"]; ok {
		if valAssembleDocument != nil {
			var valueForAssembleDocument bool
			err = json.Unmarshal(*valAssembleDocument, &valueForAssembleDocument)
			if err != nil {
				return err
			}
			this.AssembleDocument = valueForAssembleDocument
		}
	}
	if valAssembleDocumentCap, ok := objMap["AssembleDocument"]; ok {
		if valAssembleDocumentCap != nil {
			var valueForAssembleDocument bool
			err = json.Unmarshal(*valAssembleDocumentCap, &valueForAssembleDocument)
			if err != nil {
				return err
			}
			this.AssembleDocument = valueForAssembleDocument
		}
	}
	
	if valHighQualityPrint, ok := objMap["highQualityPrint"]; ok {
		if valHighQualityPrint != nil {
			var valueForHighQualityPrint bool
			err = json.Unmarshal(*valHighQualityPrint, &valueForHighQualityPrint)
			if err != nil {
				return err
			}
			this.HighQualityPrint = valueForHighQualityPrint
		}
	}
	if valHighQualityPrintCap, ok := objMap["HighQualityPrint"]; ok {
		if valHighQualityPrintCap != nil {
			var valueForHighQualityPrint bool
			err = json.Unmarshal(*valHighQualityPrintCap, &valueForHighQualityPrint)
			if err != nil {
				return err
			}
			this.HighQualityPrint = valueForHighQualityPrint
		}
	}

	return nil
}
