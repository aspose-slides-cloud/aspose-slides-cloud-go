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

// Provides options that control how a presentation is saved in Pdf format.
type IPdfExportOptions interface {

	getFormat() string
	setFormat(newValue string)

	// Specifies compression type to be used for all textual content in the document.
	getTextCompression() string
	setTextCompression(newValue string)

	// Determines if all characters of font should be embedded or only used subset.
	getEmbedFullFonts() bool
	setEmbedFullFonts(newValue bool)

	// Desired conformance level for generated PDF document.
	getCompliance() string
	setCompliance(newValue string)

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	getSufficientResolution() float64
	setSufficientResolution(newValue float64)

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	getJpegQuality() int32
	setJpegQuality(newValue int32)

	// True to draw black frame around each slide.
	getDrawSlidesFrame() bool
	setDrawSlidesFrame(newValue bool)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	getShowHiddenSlides() bool
	setShowHiddenSlides(newValue bool)

	// True to convert all metafiles used in a presentation to the PNG images.
	getSaveMetafilesAsPng() bool
	setSaveMetafilesAsPng(newValue bool)

	// Setting user password to protect the PDF document. 
	getPassword() string
	setPassword(newValue string)

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	getEmbedTrueTypeFontsForASCII() bool
	setEmbedTrueTypeFontsForASCII(newValue bool)

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	getAdditionalCommonFontFamilies() []string
	setAdditionalCommonFontFamilies(newValue []string)

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

	// Image transparent color.
	getImageTransparentColor() string
	setImageTransparentColor(newValue string)

	// True to apply specified ImageTransparentColor  to an image.
	getApplyImageTransparent() bool
	setApplyImageTransparent(newValue bool)
}

type PdfExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Specifies compression type to be used for all textual content in the document.
	TextCompression string `json:"TextCompression"`

	// Determines if all characters of font should be embedded or only used subset.
	EmbedFullFonts bool `json:"EmbedFullFonts"`

	// Desired conformance level for generated PDF document.
	Compliance string `json:"Compliance"`

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	SufficientResolution float64 `json:"SufficientResolution"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality"`

	// True to draw black frame around each slide.
	DrawSlidesFrame bool `json:"DrawSlidesFrame"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// True to convert all metafiles used in a presentation to the PNG images.
	SaveMetafilesAsPng bool `json:"SaveMetafilesAsPng"`

	// Setting user password to protect the PDF document. 
	Password string `json:"Password,omitempty"`

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	EmbedTrueTypeFontsForASCII bool `json:"EmbedTrueTypeFontsForASCII"`

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	AdditionalCommonFontFamilies []string `json:"AdditionalCommonFontFamilies,omitempty"`

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

	// Image transparent color.
	ImageTransparentColor string `json:"ImageTransparentColor,omitempty"`

	// True to apply specified ImageTransparentColor  to an image.
	ApplyImageTransparent bool `json:"ApplyImageTransparent"`
}

func (this PdfExportOptions) getFormat() string {
	return this.Format
}

func (this PdfExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this PdfExportOptions) getTextCompression() string {
	return this.TextCompression
}

func (this PdfExportOptions) setTextCompression(newValue string) {
	this.TextCompression = newValue
}
func (this PdfExportOptions) getEmbedFullFonts() bool {
	return this.EmbedFullFonts
}

func (this PdfExportOptions) setEmbedFullFonts(newValue bool) {
	this.EmbedFullFonts = newValue
}
func (this PdfExportOptions) getCompliance() string {
	return this.Compliance
}

func (this PdfExportOptions) setCompliance(newValue string) {
	this.Compliance = newValue
}
func (this PdfExportOptions) getSufficientResolution() float64 {
	return this.SufficientResolution
}

func (this PdfExportOptions) setSufficientResolution(newValue float64) {
	this.SufficientResolution = newValue
}
func (this PdfExportOptions) getJpegQuality() int32 {
	return this.JpegQuality
}

func (this PdfExportOptions) setJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this PdfExportOptions) getDrawSlidesFrame() bool {
	return this.DrawSlidesFrame
}

func (this PdfExportOptions) setDrawSlidesFrame(newValue bool) {
	this.DrawSlidesFrame = newValue
}
func (this PdfExportOptions) getShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this PdfExportOptions) setShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this PdfExportOptions) getSaveMetafilesAsPng() bool {
	return this.SaveMetafilesAsPng
}

func (this PdfExportOptions) setSaveMetafilesAsPng(newValue bool) {
	this.SaveMetafilesAsPng = newValue
}
func (this PdfExportOptions) getPassword() string {
	return this.Password
}

func (this PdfExportOptions) setPassword(newValue string) {
	this.Password = newValue
}
func (this PdfExportOptions) getEmbedTrueTypeFontsForASCII() bool {
	return this.EmbedTrueTypeFontsForASCII
}

func (this PdfExportOptions) setEmbedTrueTypeFontsForASCII(newValue bool) {
	this.EmbedTrueTypeFontsForASCII = newValue
}
func (this PdfExportOptions) getAdditionalCommonFontFamilies() []string {
	return this.AdditionalCommonFontFamilies
}

func (this PdfExportOptions) setAdditionalCommonFontFamilies(newValue []string) {
	this.AdditionalCommonFontFamilies = newValue
}
func (this PdfExportOptions) getNotesPosition() string {
	return this.NotesPosition
}

func (this PdfExportOptions) setNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this PdfExportOptions) getCommentsPosition() string {
	return this.CommentsPosition
}

func (this PdfExportOptions) setCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this PdfExportOptions) getCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this PdfExportOptions) setCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this PdfExportOptions) getCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this PdfExportOptions) setCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this PdfExportOptions) getShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this PdfExportOptions) setShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}
func (this PdfExportOptions) getImageTransparentColor() string {
	return this.ImageTransparentColor
}

func (this PdfExportOptions) setImageTransparentColor(newValue string) {
	this.ImageTransparentColor = newValue
}
func (this PdfExportOptions) getApplyImageTransparent() bool {
	return this.ApplyImageTransparent
}

func (this PdfExportOptions) setApplyImageTransparent(newValue bool) {
	this.ApplyImageTransparent = newValue
}

func (this *PdfExportOptions) UnmarshalJSON(b []byte) error {
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

	if valTextCompression, ok := objMap["TextCompression"]; ok {
		if valTextCompression != nil {
			var valueForTextCompression string
			err = json.Unmarshal(*valTextCompression, &valueForTextCompression)
			if err != nil {
				return err
			}
			this.TextCompression = valueForTextCompression
		}
	}

	if valEmbedFullFonts, ok := objMap["EmbedFullFonts"]; ok {
		if valEmbedFullFonts != nil {
			var valueForEmbedFullFonts bool
			err = json.Unmarshal(*valEmbedFullFonts, &valueForEmbedFullFonts)
			if err != nil {
				return err
			}
			this.EmbedFullFonts = valueForEmbedFullFonts
		}
	}

	if valCompliance, ok := objMap["Compliance"]; ok {
		if valCompliance != nil {
			var valueForCompliance string
			err = json.Unmarshal(*valCompliance, &valueForCompliance)
			if err != nil {
				return err
			}
			this.Compliance = valueForCompliance
		}
	}

	if valSufficientResolution, ok := objMap["SufficientResolution"]; ok {
		if valSufficientResolution != nil {
			var valueForSufficientResolution float64
			err = json.Unmarshal(*valSufficientResolution, &valueForSufficientResolution)
			if err != nil {
				return err
			}
			this.SufficientResolution = valueForSufficientResolution
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

	if valDrawSlidesFrame, ok := objMap["DrawSlidesFrame"]; ok {
		if valDrawSlidesFrame != nil {
			var valueForDrawSlidesFrame bool
			err = json.Unmarshal(*valDrawSlidesFrame, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
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

	if valSaveMetafilesAsPng, ok := objMap["SaveMetafilesAsPng"]; ok {
		if valSaveMetafilesAsPng != nil {
			var valueForSaveMetafilesAsPng bool
			err = json.Unmarshal(*valSaveMetafilesAsPng, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}

	if valPassword, ok := objMap["Password"]; ok {
		if valPassword != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPassword, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}

	if valEmbedTrueTypeFontsForASCII, ok := objMap["EmbedTrueTypeFontsForASCII"]; ok {
		if valEmbedTrueTypeFontsForASCII != nil {
			var valueForEmbedTrueTypeFontsForASCII bool
			err = json.Unmarshal(*valEmbedTrueTypeFontsForASCII, &valueForEmbedTrueTypeFontsForASCII)
			if err != nil {
				return err
			}
			this.EmbedTrueTypeFontsForASCII = valueForEmbedTrueTypeFontsForASCII
		}
	}

	if valAdditionalCommonFontFamilies, ok := objMap["AdditionalCommonFontFamilies"]; ok {
		if valAdditionalCommonFontFamilies != nil {
			var valueForAdditionalCommonFontFamilies []string
			err = json.Unmarshal(*valAdditionalCommonFontFamilies, &valueForAdditionalCommonFontFamilies)
			if err != nil {
				return err
			}
			this.AdditionalCommonFontFamilies = valueForAdditionalCommonFontFamilies
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

	if valImageTransparentColor, ok := objMap["ImageTransparentColor"]; ok {
		if valImageTransparentColor != nil {
			var valueForImageTransparentColor string
			err = json.Unmarshal(*valImageTransparentColor, &valueForImageTransparentColor)
			if err != nil {
				return err
			}
			this.ImageTransparentColor = valueForImageTransparentColor
		}
	}

	if valApplyImageTransparent, ok := objMap["ApplyImageTransparent"]; ok {
		if valApplyImageTransparent != nil {
			var valueForApplyImageTransparent bool
			err = json.Unmarshal(*valApplyImageTransparent, &valueForApplyImageTransparent)
			if err != nil {
				return err
			}
			this.ApplyImageTransparent = valueForApplyImageTransparent
		}
	}

    return nil
}
