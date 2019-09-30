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

	// Export format.
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

	// True to make layout responsive by excluding width and height attributes from svg container.
	getSvgResponsiveLayout() bool
	setSvgResponsiveLayout(newValue bool)

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

	// Export format.
	Format string `json:"Format,omitempty"`

	// Get or sets flag for save presentation as zip file
	SaveAsZip bool `json:"SaveAsZip"`

	// Get or set name of subdirectory in zip-file for store external files
	SubDirectoryName string `json:"SubDirectoryName,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// True to make layout responsive by excluding width and height attributes from svg container.
	SvgResponsiveLayout bool `json:"SvgResponsiveLayout"`

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
func (this HtmlExportOptions) getSvgResponsiveLayout() bool {
	return this.SvgResponsiveLayout
}

func (this HtmlExportOptions) setSvgResponsiveLayout(newValue bool) {
	this.SvgResponsiveLayout = newValue
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
	
	if valFormat, ok := objMap["format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	
	if valSaveAsZip, ok := objMap["saveAsZip"]; ok {
		if valSaveAsZip != nil {
			var valueForSaveAsZip bool
			err = json.Unmarshal(*valSaveAsZip, &valueForSaveAsZip)
			if err != nil {
				return err
			}
			this.SaveAsZip = valueForSaveAsZip
		}
	}
	if valSaveAsZipCap, ok := objMap["SaveAsZip"]; ok {
		if valSaveAsZipCap != nil {
			var valueForSaveAsZip bool
			err = json.Unmarshal(*valSaveAsZipCap, &valueForSaveAsZip)
			if err != nil {
				return err
			}
			this.SaveAsZip = valueForSaveAsZip
		}
	}
	
	if valSubDirectoryName, ok := objMap["subDirectoryName"]; ok {
		if valSubDirectoryName != nil {
			var valueForSubDirectoryName string
			err = json.Unmarshal(*valSubDirectoryName, &valueForSubDirectoryName)
			if err != nil {
				return err
			}
			this.SubDirectoryName = valueForSubDirectoryName
		}
	}
	if valSubDirectoryNameCap, ok := objMap["SubDirectoryName"]; ok {
		if valSubDirectoryNameCap != nil {
			var valueForSubDirectoryName string
			err = json.Unmarshal(*valSubDirectoryNameCap, &valueForSubDirectoryName)
			if err != nil {
				return err
			}
			this.SubDirectoryName = valueForSubDirectoryName
		}
	}
	
	if valShowHiddenSlides, ok := objMap["showHiddenSlides"]; ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	if valShowHiddenSlidesCap, ok := objMap["ShowHiddenSlides"]; ok {
		if valShowHiddenSlidesCap != nil {
			var valueForShowHiddenSlides bool
			err = json.Unmarshal(*valShowHiddenSlidesCap, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valSvgResponsiveLayout, ok := objMap["svgResponsiveLayout"]; ok {
		if valSvgResponsiveLayout != nil {
			var valueForSvgResponsiveLayout bool
			err = json.Unmarshal(*valSvgResponsiveLayout, &valueForSvgResponsiveLayout)
			if err != nil {
				return err
			}
			this.SvgResponsiveLayout = valueForSvgResponsiveLayout
		}
	}
	if valSvgResponsiveLayoutCap, ok := objMap["SvgResponsiveLayout"]; ok {
		if valSvgResponsiveLayoutCap != nil {
			var valueForSvgResponsiveLayout bool
			err = json.Unmarshal(*valSvgResponsiveLayoutCap, &valueForSvgResponsiveLayout)
			if err != nil {
				return err
			}
			this.SvgResponsiveLayout = valueForSvgResponsiveLayout
		}
	}
	
	if valJpegQuality, ok := objMap["jpegQuality"]; ok {
		if valJpegQuality != nil {
			var valueForJpegQuality int32
			err = json.Unmarshal(*valJpegQuality, &valueForJpegQuality)
			if err != nil {
				return err
			}
			this.JpegQuality = valueForJpegQuality
		}
	}
	if valJpegQualityCap, ok := objMap["JpegQuality"]; ok {
		if valJpegQualityCap != nil {
			var valueForJpegQuality int32
			err = json.Unmarshal(*valJpegQualityCap, &valueForJpegQuality)
			if err != nil {
				return err
			}
			this.JpegQuality = valueForJpegQuality
		}
	}
	this.PicturesCompression = ""
	if valPicturesCompression, ok := objMap["picturesCompression"]; ok {
		if valPicturesCompression != nil {
			var valueForPicturesCompression string
			err = json.Unmarshal(*valPicturesCompression, &valueForPicturesCompression)
			if err != nil {
				return err
			}
			this.PicturesCompression = valueForPicturesCompression
		}
	}
	if valPicturesCompressionCap, ok := objMap["PicturesCompression"]; ok {
		if valPicturesCompressionCap != nil {
			var valueForPicturesCompression string
			err = json.Unmarshal(*valPicturesCompressionCap, &valueForPicturesCompression)
			if err != nil {
				return err
			}
			this.PicturesCompression = valueForPicturesCompression
		}
	}
	
	if valDeletePicturesCroppedAreas, ok := objMap["deletePicturesCroppedAreas"]; ok {
		if valDeletePicturesCroppedAreas != nil {
			var valueForDeletePicturesCroppedAreas bool
			err = json.Unmarshal(*valDeletePicturesCroppedAreas, &valueForDeletePicturesCroppedAreas)
			if err != nil {
				return err
			}
			this.DeletePicturesCroppedAreas = valueForDeletePicturesCroppedAreas
		}
	}
	if valDeletePicturesCroppedAreasCap, ok := objMap["DeletePicturesCroppedAreas"]; ok {
		if valDeletePicturesCroppedAreasCap != nil {
			var valueForDeletePicturesCroppedAreas bool
			err = json.Unmarshal(*valDeletePicturesCroppedAreasCap, &valueForDeletePicturesCroppedAreas)
			if err != nil {
				return err
			}
			this.DeletePicturesCroppedAreas = valueForDeletePicturesCroppedAreas
		}
	}
	this.NotesPosition = "None"
	if valNotesPosition, ok := objMap["notesPosition"]; ok {
		if valNotesPosition != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPosition, &valueForNotesPosition)
			if err != nil {
				return err
			}
			this.NotesPosition = valueForNotesPosition
		}
	}
	if valNotesPositionCap, ok := objMap["NotesPosition"]; ok {
		if valNotesPositionCap != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPositionCap, &valueForNotesPosition)
			if err != nil {
				return err
			}
			this.NotesPosition = valueForNotesPosition
		}
	}
	this.CommentsPosition = "None"
	if valCommentsPosition, ok := objMap["commentsPosition"]; ok {
		if valCommentsPosition != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPosition)
			if err != nil {
				return err
			}
			this.CommentsPosition = valueForCommentsPosition
		}
	}
	if valCommentsPositionCap, ok := objMap["CommentsPosition"]; ok {
		if valCommentsPositionCap != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPositionCap, &valueForCommentsPosition)
			if err != nil {
				return err
			}
			this.CommentsPosition = valueForCommentsPosition
		}
	}
	
	if valCommentsAreaWidth, ok := objMap["commentsAreaWidth"]; ok {
		if valCommentsAreaWidth != nil {
			var valueForCommentsAreaWidth int32
			err = json.Unmarshal(*valCommentsAreaWidth, &valueForCommentsAreaWidth)
			if err != nil {
				return err
			}
			this.CommentsAreaWidth = valueForCommentsAreaWidth
		}
	}
	if valCommentsAreaWidthCap, ok := objMap["CommentsAreaWidth"]; ok {
		if valCommentsAreaWidthCap != nil {
			var valueForCommentsAreaWidth int32
			err = json.Unmarshal(*valCommentsAreaWidthCap, &valueForCommentsAreaWidth)
			if err != nil {
				return err
			}
			this.CommentsAreaWidth = valueForCommentsAreaWidth
		}
	}
	
	if valCommentsAreaColor, ok := objMap["commentsAreaColor"]; ok {
		if valCommentsAreaColor != nil {
			var valueForCommentsAreaColor string
			err = json.Unmarshal(*valCommentsAreaColor, &valueForCommentsAreaColor)
			if err != nil {
				return err
			}
			this.CommentsAreaColor = valueForCommentsAreaColor
		}
	}
	if valCommentsAreaColorCap, ok := objMap["CommentsAreaColor"]; ok {
		if valCommentsAreaColorCap != nil {
			var valueForCommentsAreaColor string
			err = json.Unmarshal(*valCommentsAreaColorCap, &valueForCommentsAreaColor)
			if err != nil {
				return err
			}
			this.CommentsAreaColor = valueForCommentsAreaColor
		}
	}
	
	if valShowCommentsByNoAuthor, ok := objMap["showCommentsByNoAuthor"]; ok {
		if valShowCommentsByNoAuthor != nil {
			var valueForShowCommentsByNoAuthor bool
			err = json.Unmarshal(*valShowCommentsByNoAuthor, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}
	if valShowCommentsByNoAuthorCap, ok := objMap["ShowCommentsByNoAuthor"]; ok {
		if valShowCommentsByNoAuthorCap != nil {
			var valueForShowCommentsByNoAuthor bool
			err = json.Unmarshal(*valShowCommentsByNoAuthorCap, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}

    return nil
}
