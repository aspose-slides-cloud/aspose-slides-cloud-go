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


type ITiffExportOptions interface {

	getFormat() string
	setFormat(newValue string)

	getCompression() TiffCompressionType
	setCompression(newValue TiffCompressionType)

	getWidth() int32
	setWidth(newValue int32)

	getHeight() int32
	setHeight(newValue int32)

	getDpiX() int32
	setDpiX(newValue int32)

	getDpiY() int32
	setDpiY(newValue int32)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	getShowHiddenSlides() bool
	setShowHiddenSlides(newValue bool)

	// Specifies the pixel format for the generated images. Read/write .
	getPixelFormat() ImagePixelFormat
	setPixelFormat(newValue ImagePixelFormat)

	// Gets or sets the position of the notes on the page.
	getNotesPosition() NotesPositions
	setNotesPosition(newValue NotesPositions)

	// Gets or sets the position of the comments on the page.
	getCommentsPosition() CommentsPositions
	setCommentsPosition(newValue CommentsPositions)

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

type TiffExportOptions struct {

	Format string `json:"Format,omitempty"`

	Compression TiffCompressionType `json:"Compression,omitempty"`

	Width int32 `json:"Width,omitempty"`

	Height int32 `json:"Height,omitempty"`

	DpiX int32 `json:"DpiX,omitempty"`

	DpiY int32 `json:"DpiY,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides,omitempty"`

	// Specifies the pixel format for the generated images. Read/write .
	PixelFormat ImagePixelFormat `json:"PixelFormat,omitempty"`

	// Gets or sets the position of the notes on the page.
	NotesPosition NotesPositions `json:"NotesPosition,omitempty"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition CommentsPositions `json:"CommentsPosition,omitempty"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth,omitempty"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor bool `json:"ShowCommentsByNoAuthor,omitempty"`
}

func (this TiffExportOptions) getFormat() string {
	return this.Format
}

func (this TiffExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this TiffExportOptions) getCompression() TiffCompressionType {
	return this.Compression
}

func (this TiffExportOptions) setCompression(newValue TiffCompressionType) {
	this.Compression = newValue
}
func (this TiffExportOptions) getWidth() int32 {
	return this.Width
}

func (this TiffExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this TiffExportOptions) getHeight() int32 {
	return this.Height
}

func (this TiffExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this TiffExportOptions) getDpiX() int32 {
	return this.DpiX
}

func (this TiffExportOptions) setDpiX(newValue int32) {
	this.DpiX = newValue
}
func (this TiffExportOptions) getDpiY() int32 {
	return this.DpiY
}

func (this TiffExportOptions) setDpiY(newValue int32) {
	this.DpiY = newValue
}
func (this TiffExportOptions) getShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this TiffExportOptions) setShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this TiffExportOptions) getPixelFormat() ImagePixelFormat {
	return this.PixelFormat
}

func (this TiffExportOptions) setPixelFormat(newValue ImagePixelFormat) {
	this.PixelFormat = newValue
}
func (this TiffExportOptions) getNotesPosition() NotesPositions {
	return this.NotesPosition
}

func (this TiffExportOptions) setNotesPosition(newValue NotesPositions) {
	this.NotesPosition = newValue
}
func (this TiffExportOptions) getCommentsPosition() CommentsPositions {
	return this.CommentsPosition
}

func (this TiffExportOptions) setCommentsPosition(newValue CommentsPositions) {
	this.CommentsPosition = newValue
}
func (this TiffExportOptions) getCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this TiffExportOptions) setCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this TiffExportOptions) getCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this TiffExportOptions) setCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this TiffExportOptions) getShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this TiffExportOptions) setShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *TiffExportOptions) UnmarshalJSON(b []byte) error {
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

	if valCompression, ok := objMap["Compression"]; ok {
		if valCompression != nil {
			var valueForCompression TiffCompressionType
			err = json.Unmarshal(*valCompression, &valueForCompression)
			if err != nil {
				return err
			}
			this.Compression = valueForCompression
		}
	}

	if valWidth, ok := objMap["Width"]; ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}

	if valHeight, ok := objMap["Height"]; ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}

	if valDpiX, ok := objMap["DpiX"]; ok {
		if valDpiX != nil {
			var valueForDpiX int32
			err = json.Unmarshal(*valDpiX, &valueForDpiX)
			if err != nil {
				return err
			}
			this.DpiX = valueForDpiX
		}
	}

	if valDpiY, ok := objMap["DpiY"]; ok {
		if valDpiY != nil {
			var valueForDpiY int32
			err = json.Unmarshal(*valDpiY, &valueForDpiY)
			if err != nil {
				return err
			}
			this.DpiY = valueForDpiY
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

	if valPixelFormat, ok := objMap["PixelFormat"]; ok {
		if valPixelFormat != nil {
			var valueForPixelFormat ImagePixelFormat
			err = json.Unmarshal(*valPixelFormat, &valueForPixelFormat)
			if err != nil {
				return err
			}
			this.PixelFormat = valueForPixelFormat
		}
	}

	if valNotesPosition, ok := objMap["NotesPosition"]; ok {
		if valNotesPosition != nil {
			var valueForNotesPosition NotesPositions
			err = json.Unmarshal(*valNotesPosition, &valueForNotesPosition)
			if err != nil {
				return err
			}
			this.NotesPosition = valueForNotesPosition
		}
	}

	if valCommentsPosition, ok := objMap["CommentsPosition"]; ok {
		if valCommentsPosition != nil {
			var valueForCommentsPosition CommentsPositions
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
