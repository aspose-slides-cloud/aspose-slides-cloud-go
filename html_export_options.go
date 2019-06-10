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

// Provides options that control how a presentation is saved in Html format.
type IHtmlExportOptions interface {

	getFormat() string
	setFormat(newValue string)

	// Get or sets flag for save presentation as zip file
	getSaveAsZip() bool
	setSaveAsZip(newValue bool)

	// Get or set name of subdirectory in zip-file for store external files
	getSubDirectoryName() string
	setSubDirectoryName(newValue string)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	getShowHiddenSlides() bool
	setShowHiddenSlides(newValue bool)

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	getJpegQuality() int32
	setJpegQuality(newValue int32)

	// Represents the pictures compression level
	getPicturesCompression() string
	setPicturesCompression(newValue string)

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	getDeletePicturesCroppedAreas() bool
	setDeletePicturesCroppedAreas(newValue bool)

	// Gets or sets the position of the notes on the page.
	getNotesPosition() string
	setNotesPosition(newValue string)

	// Gets or sets the position of the comments on the page.
	getCommentsPosition() string
	setCommentsPosition(newValue string)

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	getCommentsAreaWidth() int32
	setCommentsAreaWidth(newValue int32)

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	getCommentsAreaColor() string
	setCommentsAreaColor(newValue string)

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	getShowCommentsByNoAuthor() bool
	setShowCommentsByNoAuthor(newValue bool)
}

type HtmlExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Get or sets flag for save presentation as zip file
	SaveAsZip bool `json:"SaveAsZip"`

	// Get or set name of subdirectory in zip-file for store external files
	SubDirectoryName string `json:"SubDirectoryName,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality"`

	// Represents the pictures compression level
	PicturesCompression string `json:"PicturesCompression,omitempty"`

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	DeletePicturesCroppedAreas bool `json:"DeletePicturesCroppedAreas"`

	// Gets or sets the position of the notes on the page.
	NotesPosition string `json:"NotesPosition"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition string `json:"CommentsPosition"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor bool `json:"ShowCommentsByNoAuthor"`
}

func (this HtmlExportOptions) getFormat() string {
	return this.Format
}

func (this HtmlExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this HtmlExportOptions) getSaveAsZip() bool {
	return this.SaveAsZip
}

func (this HtmlExportOptions) setSaveAsZip(newValue bool) {
	this.SaveAsZip = newValue
}
func (this HtmlExportOptions) getSubDirectoryName() string {
	return this.SubDirectoryName
}

func (this HtmlExportOptions) setSubDirectoryName(newValue string) {
	this.SubDirectoryName = newValue
}
func (this HtmlExportOptions) getShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this HtmlExportOptions) setShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this HtmlExportOptions) getJpegQuality() int32 {
	return this.JpegQuality
}

func (this HtmlExportOptions) setJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this HtmlExportOptions) getPicturesCompression() string {
	return this.PicturesCompression
}

func (this HtmlExportOptions) setPicturesCompression(newValue string) {
	this.PicturesCompression = newValue
}
func (this HtmlExportOptions) getDeletePicturesCroppedAreas() bool {
	return this.DeletePicturesCroppedAreas
}

func (this HtmlExportOptions) setDeletePicturesCroppedAreas(newValue bool) {
	this.DeletePicturesCroppedAreas = newValue
}
func (this HtmlExportOptions) getNotesPosition() string {
	return this.NotesPosition
}

func (this HtmlExportOptions) setNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this HtmlExportOptions) getCommentsPosition() string {
	return this.CommentsPosition
}

func (this HtmlExportOptions) setCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this HtmlExportOptions) getCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this HtmlExportOptions) setCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this HtmlExportOptions) getCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this HtmlExportOptions) setCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this HtmlExportOptions) getShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this HtmlExportOptions) setShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *HtmlExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valFormat, ok := objMap["Format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}

	if valSaveAsZip, ok := objMap["SaveAsZip"]; ok {
		if valSaveAsZip != nil {
			var valueForSaveAsZip bool
			err = json.Unmarshal(*valSaveAsZip, &valueForSaveAsZip)
			if err != nil {
				return err
			}
			this.SaveAsZip = valueForSaveAsZip
		}
	}

	if valSubDirectoryName, ok := objMap["SubDirectoryName"]; ok {
		if valSubDirectoryName != nil {
			var valueForSubDirectoryName string
			err = json.Unmarshal(*valSubDirectoryName, &valueForSubDirectoryName)
			if err != nil {
				return err
			}
			this.SubDirectoryName = valueForSubDirectoryName
		}
	}

	if valShowHiddenSlides, ok := objMap["ShowHiddenSlides"]; ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}

	if valJpegQuality, ok := objMap["JpegQuality"]; ok {
		if valJpegQuality != nil {
			var valueForJpegQuality int32
			err = json.Unmarshal(*valJpegQuality, &valueForJpegQuality)
			if err != nil {
				return err
			}
			this.JpegQuality = valueForJpegQuality
		}
	}

	if valPicturesCompression, ok := objMap["PicturesCompression"]; ok {
		if valPicturesCompression != nil {
			var valueForPicturesCompression string
			err = json.Unmarshal(*valPicturesCompression, &valueForPicturesCompression)
			if err != nil {
				return err
			}
			this.PicturesCompression = valueForPicturesCompression
		}
	}

	if valDeletePicturesCroppedAreas, ok := objMap["DeletePicturesCroppedAreas"]; ok {
		if valDeletePicturesCroppedAreas != nil {
			var valueForDeletePicturesCroppedAreas bool
			err = json.Unmarshal(*valDeletePicturesCroppedAreas, &valueForDeletePicturesCroppedAreas)
			if err != nil {
				return err
			}
			this.DeletePicturesCroppedAreas = valueForDeletePicturesCroppedAreas
		}
	}

	if valNotesPosition, ok := objMap["NotesPosition"]; ok {
		if valNotesPosition != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPosition, &valueForNotesPosition)
			if err != nil {
				return err
			}
			this.NotesPosition = valueForNotesPosition
		}
	}

	if valCommentsPosition, ok := objMap["CommentsPosition"]; ok {
		if valCommentsPosition != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPosition)
			if err != nil {
				return err
			}
			this.CommentsPosition = valueForCommentsPosition
		}
	}

	if valCommentsAreaWidth, ok := objMap["CommentsAreaWidth"]; ok {
		if valCommentsAreaWidth != nil {
			var valueForCommentsAreaWidth int32
			err = json.Unmarshal(*valCommentsAreaWidth, &valueForCommentsAreaWidth)
			if err != nil {
				return err
			}
			this.CommentsAreaWidth = valueForCommentsAreaWidth
		}
	}

	if valCommentsAreaColor, ok := objMap["CommentsAreaColor"]; ok {
		if valCommentsAreaColor != nil {
			var valueForCommentsAreaColor string
			err = json.Unmarshal(*valCommentsAreaColor, &valueForCommentsAreaColor)
			if err != nil {
				return err
			}
			this.CommentsAreaColor = valueForCommentsAreaColor
		}
	}

	if valShowCommentsByNoAuthor, ok := objMap["ShowCommentsByNoAuthor"]; ok {
		if valShowCommentsByNoAuthor != nil {
			var valueForShowCommentsByNoAuthor bool
			err = json.Unmarshal(*valShowCommentsByNoAuthor, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}

    return nil
}
