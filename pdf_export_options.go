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

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Specifies compression type to be used for all textual content in the document.
	GetTextCompression() string
	SetTextCompression(newValue string)

	// Determines if all characters of font should be embedded or only used subset.
	GetEmbedFullFonts() bool
	SetEmbedFullFonts(newValue bool)

	// Desired conformance level for generated PDF document.
	GetCompliance() string
	SetCompliance(newValue string)

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	GetSufficientResolution() float64
	SetSufficientResolution(newValue float64)

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	GetJpegQuality() int32
	SetJpegQuality(newValue int32)

	// True to draw black frame around each slide.
	GetDrawSlidesFrame() bool
	SetDrawSlidesFrame(newValue bool)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() bool
	SetShowHiddenSlides(newValue bool)

	// True to convert all metafiles used in a presentation to the PNG images.
	GetSaveMetafilesAsPng() bool
	SetSaveMetafilesAsPng(newValue bool)

	// Setting user password to protect the PDF document. 
	GetPassword() string
	SetPassword(newValue string)

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	GetEmbedTrueTypeFontsForASCII() bool
	SetEmbedTrueTypeFontsForASCII(newValue bool)

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	GetAdditionalCommonFontFamilies() []string
	SetAdditionalCommonFontFamilies(newValue []string)

	// Gets or sets the position of the notes on the page.
	GetNotesPosition() string
	SetNotesPosition(newValue string)

	// Gets or sets the position of the comments on the page.
	GetCommentsPosition() string
	SetCommentsPosition(newValue string)

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	GetCommentsAreaWidth() int32
	SetCommentsAreaWidth(newValue int32)

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	GetCommentsAreaColor() string
	SetCommentsAreaColor(newValue string)

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	GetShowCommentsByNoAuthor() bool
	SetShowCommentsByNoAuthor(newValue bool)

	// Image transparent color.
	GetImageTransparentColor() string
	SetImageTransparentColor(newValue string)

	// True to apply specified ImageTransparentColor  to an image.
	GetApplyImageTransparent() bool
	SetApplyImageTransparent(newValue bool)

	// Access permissions that should be granted when the document is opened with user access.  Default is AccessPermissions.None.             
	GetAccessPermissions() IAccessPermissions
	SetAccessPermissions(newValue IAccessPermissions)
}

type PdfExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies compression type to be used for all textual content in the document.
	TextCompression string `json:"TextCompression,omitempty"`

	// Determines if all characters of font should be embedded or only used subset.
	EmbedFullFonts bool `json:"EmbedFullFonts"`

	// Desired conformance level for generated PDF document.
	Compliance string `json:"Compliance,omitempty"`

	// Returns or sets a value determining resolution of images inside PDF document.  Property affects on file size, time of export and image quality. The default value is 96.
	SufficientResolution float64 `json:"SufficientResolution,omitempty"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

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
	NotesPosition string `json:"NotesPosition,omitempty"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition string `json:"CommentsPosition,omitempty"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth,omitempty"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor bool `json:"ShowCommentsByNoAuthor"`

	// Image transparent color.
	ImageTransparentColor string `json:"ImageTransparentColor,omitempty"`

	// True to apply specified ImageTransparentColor  to an image.
	ApplyImageTransparent bool `json:"ApplyImageTransparent"`

	// Access permissions that should be granted when the document is opened with user access.  Default is AccessPermissions.None.             
	AccessPermissions IAccessPermissions `json:"AccessPermissions,omitempty"`
}

func NewPdfExportOptions() *PdfExportOptions {
	instance := new(PdfExportOptions)
	instance.TextCompression = ""
	instance.Compliance = ""
	instance.NotesPosition = ""
	instance.CommentsPosition = ""
	return instance
}

func (this *PdfExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *PdfExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *PdfExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *PdfExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *PdfExportOptions) GetFormat() string {
	return this.Format
}

func (this *PdfExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *PdfExportOptions) GetTextCompression() string {
	return this.TextCompression
}

func (this *PdfExportOptions) SetTextCompression(newValue string) {
	this.TextCompression = newValue
}
func (this *PdfExportOptions) GetEmbedFullFonts() bool {
	return this.EmbedFullFonts
}

func (this *PdfExportOptions) SetEmbedFullFonts(newValue bool) {
	this.EmbedFullFonts = newValue
}
func (this *PdfExportOptions) GetCompliance() string {
	return this.Compliance
}

func (this *PdfExportOptions) SetCompliance(newValue string) {
	this.Compliance = newValue
}
func (this *PdfExportOptions) GetSufficientResolution() float64 {
	return this.SufficientResolution
}

func (this *PdfExportOptions) SetSufficientResolution(newValue float64) {
	this.SufficientResolution = newValue
}
func (this *PdfExportOptions) GetJpegQuality() int32 {
	return this.JpegQuality
}

func (this *PdfExportOptions) SetJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *PdfExportOptions) GetDrawSlidesFrame() bool {
	return this.DrawSlidesFrame
}

func (this *PdfExportOptions) SetDrawSlidesFrame(newValue bool) {
	this.DrawSlidesFrame = newValue
}
func (this *PdfExportOptions) GetShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this *PdfExportOptions) SetShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this *PdfExportOptions) GetSaveMetafilesAsPng() bool {
	return this.SaveMetafilesAsPng
}

func (this *PdfExportOptions) SetSaveMetafilesAsPng(newValue bool) {
	this.SaveMetafilesAsPng = newValue
}
func (this *PdfExportOptions) GetPassword() string {
	return this.Password
}

func (this *PdfExportOptions) SetPassword(newValue string) {
	this.Password = newValue
}
func (this *PdfExportOptions) GetEmbedTrueTypeFontsForASCII() bool {
	return this.EmbedTrueTypeFontsForASCII
}

func (this *PdfExportOptions) SetEmbedTrueTypeFontsForASCII(newValue bool) {
	this.EmbedTrueTypeFontsForASCII = newValue
}
func (this *PdfExportOptions) GetAdditionalCommonFontFamilies() []string {
	return this.AdditionalCommonFontFamilies
}

func (this *PdfExportOptions) SetAdditionalCommonFontFamilies(newValue []string) {
	this.AdditionalCommonFontFamilies = newValue
}
func (this *PdfExportOptions) GetNotesPosition() string {
	return this.NotesPosition
}

func (this *PdfExportOptions) SetNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this *PdfExportOptions) GetCommentsPosition() string {
	return this.CommentsPosition
}

func (this *PdfExportOptions) SetCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this *PdfExportOptions) GetCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this *PdfExportOptions) SetCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this *PdfExportOptions) GetCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this *PdfExportOptions) SetCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this *PdfExportOptions) GetShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this *PdfExportOptions) SetShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}
func (this *PdfExportOptions) GetImageTransparentColor() string {
	return this.ImageTransparentColor
}

func (this *PdfExportOptions) SetImageTransparentColor(newValue string) {
	this.ImageTransparentColor = newValue
}
func (this *PdfExportOptions) GetApplyImageTransparent() bool {
	return this.ApplyImageTransparent
}

func (this *PdfExportOptions) SetApplyImageTransparent(newValue bool) {
	this.ApplyImageTransparent = newValue
}
func (this *PdfExportOptions) GetAccessPermissions() IAccessPermissions {
	return this.AccessPermissions
}

func (this *PdfExportOptions) SetAccessPermissions(newValue IAccessPermissions) {
	this.AccessPermissions = newValue
}

func (this *PdfExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := objMap["defaultRegularFont"]; ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	if valDefaultRegularFontCap, ok := objMap["DefaultRegularFont"]; ok {
		if valDefaultRegularFontCap != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFontCap, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valFontFallbackRules, ok := objMap["fontFallbackRules"]; ok {
		if valFontFallbackRules != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRules, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	if valFontFallbackRulesCap, ok := objMap["FontFallbackRules"]; ok {
		if valFontFallbackRulesCap != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRulesCap, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
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
	this.TextCompression = ""
	if valTextCompression, ok := objMap["textCompression"]; ok {
		if valTextCompression != nil {
			var valueForTextCompression string
			err = json.Unmarshal(*valTextCompression, &valueForTextCompression)
			if err != nil {
				var valueForTextCompressionInt int32
				err = json.Unmarshal(*valTextCompression, &valueForTextCompressionInt)
				if err != nil {
					return err
				}
				this.TextCompression = string(valueForTextCompressionInt)
			} else {
				this.TextCompression = valueForTextCompression
			}
		}
	}
	if valTextCompressionCap, ok := objMap["TextCompression"]; ok {
		if valTextCompressionCap != nil {
			var valueForTextCompression string
			err = json.Unmarshal(*valTextCompressionCap, &valueForTextCompression)
			if err != nil {
				var valueForTextCompressionInt int32
				err = json.Unmarshal(*valTextCompressionCap, &valueForTextCompressionInt)
				if err != nil {
					return err
				}
				this.TextCompression = string(valueForTextCompressionInt)
			} else {
				this.TextCompression = valueForTextCompression
			}
		}
	}
	
	if valEmbedFullFonts, ok := objMap["embedFullFonts"]; ok {
		if valEmbedFullFonts != nil {
			var valueForEmbedFullFonts bool
			err = json.Unmarshal(*valEmbedFullFonts, &valueForEmbedFullFonts)
			if err != nil {
				return err
			}
			this.EmbedFullFonts = valueForEmbedFullFonts
		}
	}
	if valEmbedFullFontsCap, ok := objMap["EmbedFullFonts"]; ok {
		if valEmbedFullFontsCap != nil {
			var valueForEmbedFullFonts bool
			err = json.Unmarshal(*valEmbedFullFontsCap, &valueForEmbedFullFonts)
			if err != nil {
				return err
			}
			this.EmbedFullFonts = valueForEmbedFullFonts
		}
	}
	this.Compliance = ""
	if valCompliance, ok := objMap["compliance"]; ok {
		if valCompliance != nil {
			var valueForCompliance string
			err = json.Unmarshal(*valCompliance, &valueForCompliance)
			if err != nil {
				var valueForComplianceInt int32
				err = json.Unmarshal(*valCompliance, &valueForComplianceInt)
				if err != nil {
					return err
				}
				this.Compliance = string(valueForComplianceInt)
			} else {
				this.Compliance = valueForCompliance
			}
		}
	}
	if valComplianceCap, ok := objMap["Compliance"]; ok {
		if valComplianceCap != nil {
			var valueForCompliance string
			err = json.Unmarshal(*valComplianceCap, &valueForCompliance)
			if err != nil {
				var valueForComplianceInt int32
				err = json.Unmarshal(*valComplianceCap, &valueForComplianceInt)
				if err != nil {
					return err
				}
				this.Compliance = string(valueForComplianceInt)
			} else {
				this.Compliance = valueForCompliance
			}
		}
	}
	
	if valSufficientResolution, ok := objMap["sufficientResolution"]; ok {
		if valSufficientResolution != nil {
			var valueForSufficientResolution float64
			err = json.Unmarshal(*valSufficientResolution, &valueForSufficientResolution)
			if err != nil {
				return err
			}
			this.SufficientResolution = valueForSufficientResolution
		}
	}
	if valSufficientResolutionCap, ok := objMap["SufficientResolution"]; ok {
		if valSufficientResolutionCap != nil {
			var valueForSufficientResolution float64
			err = json.Unmarshal(*valSufficientResolutionCap, &valueForSufficientResolution)
			if err != nil {
				return err
			}
			this.SufficientResolution = valueForSufficientResolution
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
	
	if valDrawSlidesFrame, ok := objMap["drawSlidesFrame"]; ok {
		if valDrawSlidesFrame != nil {
			var valueForDrawSlidesFrame bool
			err = json.Unmarshal(*valDrawSlidesFrame, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
		}
	}
	if valDrawSlidesFrameCap, ok := objMap["DrawSlidesFrame"]; ok {
		if valDrawSlidesFrameCap != nil {
			var valueForDrawSlidesFrame bool
			err = json.Unmarshal(*valDrawSlidesFrameCap, &valueForDrawSlidesFrame)
			if err != nil {
				return err
			}
			this.DrawSlidesFrame = valueForDrawSlidesFrame
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
	
	if valSaveMetafilesAsPng, ok := objMap["saveMetafilesAsPng"]; ok {
		if valSaveMetafilesAsPng != nil {
			var valueForSaveMetafilesAsPng bool
			err = json.Unmarshal(*valSaveMetafilesAsPng, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	if valSaveMetafilesAsPngCap, ok := objMap["SaveMetafilesAsPng"]; ok {
		if valSaveMetafilesAsPngCap != nil {
			var valueForSaveMetafilesAsPng bool
			err = json.Unmarshal(*valSaveMetafilesAsPngCap, &valueForSaveMetafilesAsPng)
			if err != nil {
				return err
			}
			this.SaveMetafilesAsPng = valueForSaveMetafilesAsPng
		}
	}
	
	if valPassword, ok := objMap["password"]; ok {
		if valPassword != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPassword, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}
	if valPasswordCap, ok := objMap["Password"]; ok {
		if valPasswordCap != nil {
			var valueForPassword string
			err = json.Unmarshal(*valPasswordCap, &valueForPassword)
			if err != nil {
				return err
			}
			this.Password = valueForPassword
		}
	}
	
	if valEmbedTrueTypeFontsForASCII, ok := objMap["embedTrueTypeFontsForASCII"]; ok {
		if valEmbedTrueTypeFontsForASCII != nil {
			var valueForEmbedTrueTypeFontsForASCII bool
			err = json.Unmarshal(*valEmbedTrueTypeFontsForASCII, &valueForEmbedTrueTypeFontsForASCII)
			if err != nil {
				return err
			}
			this.EmbedTrueTypeFontsForASCII = valueForEmbedTrueTypeFontsForASCII
		}
	}
	if valEmbedTrueTypeFontsForASCIICap, ok := objMap["EmbedTrueTypeFontsForASCII"]; ok {
		if valEmbedTrueTypeFontsForASCIICap != nil {
			var valueForEmbedTrueTypeFontsForASCII bool
			err = json.Unmarshal(*valEmbedTrueTypeFontsForASCIICap, &valueForEmbedTrueTypeFontsForASCII)
			if err != nil {
				return err
			}
			this.EmbedTrueTypeFontsForASCII = valueForEmbedTrueTypeFontsForASCII
		}
	}
	
	if valAdditionalCommonFontFamilies, ok := objMap["additionalCommonFontFamilies"]; ok {
		if valAdditionalCommonFontFamilies != nil {
			var valueForAdditionalCommonFontFamilies []string
			err = json.Unmarshal(*valAdditionalCommonFontFamilies, &valueForAdditionalCommonFontFamilies)
			if err != nil {
				return err
			}
			this.AdditionalCommonFontFamilies = valueForAdditionalCommonFontFamilies
		}
	}
	if valAdditionalCommonFontFamiliesCap, ok := objMap["AdditionalCommonFontFamilies"]; ok {
		if valAdditionalCommonFontFamiliesCap != nil {
			var valueForAdditionalCommonFontFamilies []string
			err = json.Unmarshal(*valAdditionalCommonFontFamiliesCap, &valueForAdditionalCommonFontFamilies)
			if err != nil {
				return err
			}
			this.AdditionalCommonFontFamilies = valueForAdditionalCommonFontFamilies
		}
	}
	this.NotesPosition = ""
	if valNotesPosition, ok := objMap["notesPosition"]; ok {
		if valNotesPosition != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPosition, &valueForNotesPosition)
			if err != nil {
				var valueForNotesPositionInt int32
				err = json.Unmarshal(*valNotesPosition, &valueForNotesPositionInt)
				if err != nil {
					return err
				}
				this.NotesPosition = string(valueForNotesPositionInt)
			} else {
				this.NotesPosition = valueForNotesPosition
			}
		}
	}
	if valNotesPositionCap, ok := objMap["NotesPosition"]; ok {
		if valNotesPositionCap != nil {
			var valueForNotesPosition string
			err = json.Unmarshal(*valNotesPositionCap, &valueForNotesPosition)
			if err != nil {
				var valueForNotesPositionInt int32
				err = json.Unmarshal(*valNotesPositionCap, &valueForNotesPositionInt)
				if err != nil {
					return err
				}
				this.NotesPosition = string(valueForNotesPositionInt)
			} else {
				this.NotesPosition = valueForNotesPosition
			}
		}
	}
	this.CommentsPosition = ""
	if valCommentsPosition, ok := objMap["commentsPosition"]; ok {
		if valCommentsPosition != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPosition)
			if err != nil {
				var valueForCommentsPositionInt int32
				err = json.Unmarshal(*valCommentsPosition, &valueForCommentsPositionInt)
				if err != nil {
					return err
				}
				this.CommentsPosition = string(valueForCommentsPositionInt)
			} else {
				this.CommentsPosition = valueForCommentsPosition
			}
		}
	}
	if valCommentsPositionCap, ok := objMap["CommentsPosition"]; ok {
		if valCommentsPositionCap != nil {
			var valueForCommentsPosition string
			err = json.Unmarshal(*valCommentsPositionCap, &valueForCommentsPosition)
			if err != nil {
				var valueForCommentsPositionInt int32
				err = json.Unmarshal(*valCommentsPositionCap, &valueForCommentsPositionInt)
				if err != nil {
					return err
				}
				this.CommentsPosition = string(valueForCommentsPositionInt)
			} else {
				this.CommentsPosition = valueForCommentsPosition
			}
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
	
	if valImageTransparentColor, ok := objMap["imageTransparentColor"]; ok {
		if valImageTransparentColor != nil {
			var valueForImageTransparentColor string
			err = json.Unmarshal(*valImageTransparentColor, &valueForImageTransparentColor)
			if err != nil {
				return err
			}
			this.ImageTransparentColor = valueForImageTransparentColor
		}
	}
	if valImageTransparentColorCap, ok := objMap["ImageTransparentColor"]; ok {
		if valImageTransparentColorCap != nil {
			var valueForImageTransparentColor string
			err = json.Unmarshal(*valImageTransparentColorCap, &valueForImageTransparentColor)
			if err != nil {
				return err
			}
			this.ImageTransparentColor = valueForImageTransparentColor
		}
	}
	
	if valApplyImageTransparent, ok := objMap["applyImageTransparent"]; ok {
		if valApplyImageTransparent != nil {
			var valueForApplyImageTransparent bool
			err = json.Unmarshal(*valApplyImageTransparent, &valueForApplyImageTransparent)
			if err != nil {
				return err
			}
			this.ApplyImageTransparent = valueForApplyImageTransparent
		}
	}
	if valApplyImageTransparentCap, ok := objMap["ApplyImageTransparent"]; ok {
		if valApplyImageTransparentCap != nil {
			var valueForApplyImageTransparent bool
			err = json.Unmarshal(*valApplyImageTransparentCap, &valueForApplyImageTransparent)
			if err != nil {
				return err
			}
			this.ApplyImageTransparent = valueForApplyImageTransparent
		}
	}
	
	if valAccessPermissions, ok := objMap["accessPermissions"]; ok {
		if valAccessPermissions != nil {
			var valueForAccessPermissions AccessPermissions
			err = json.Unmarshal(*valAccessPermissions, &valueForAccessPermissions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("AccessPermissions", *valAccessPermissions)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valAccessPermissions, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAccessPermissions)
			if ok {
				this.AccessPermissions = vInterfaceObject
			}
		}
	}
	if valAccessPermissionsCap, ok := objMap["AccessPermissions"]; ok {
		if valAccessPermissionsCap != nil {
			var valueForAccessPermissions AccessPermissions
			err = json.Unmarshal(*valAccessPermissionsCap, &valueForAccessPermissions)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("AccessPermissions", *valAccessPermissionsCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valAccessPermissionsCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IAccessPermissions)
			if ok {
				this.AccessPermissions = vInterfaceObject
			}
		}
	}

	return nil
}
