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

// Provides options that control how a presentation is saved in TIFF format.
type ITiffExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Gets or sets the height of slides in the output image format.
	GetHeight() int32
	SetHeight(newValue int32)

	// Gets or sets the height of slides in the output the output image format.
	GetWidth() int32
	SetWidth(newValue int32)

	// Compression type.
	GetCompression() string
	SetCompression(newValue string)

	// Horizontal resolution, in dots per inch.
	GetDpiX() int32
	SetDpiX(newValue int32)

	// Vertical resolution, in dots per inch.
	GetDpiY() int32
	SetDpiY(newValue int32)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() bool
	SetShowHiddenSlides(newValue bool)

	// Specifies the pixel format for the generated images. Read/write ImagePixelFormat.
	GetPixelFormat() string
	SetPixelFormat(newValue string)

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
}

type TiffExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Gets or sets the height of slides in the output image format.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output the output image format.
	Width int32 `json:"Width,omitempty"`

	// Compression type.
	Compression string `json:"Compression,omitempty"`

	// Horizontal resolution, in dots per inch.
	DpiX int32 `json:"DpiX,omitempty"`

	// Vertical resolution, in dots per inch.
	DpiY int32 `json:"DpiY,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// Specifies the pixel format for the generated images. Read/write ImagePixelFormat.
	PixelFormat string `json:"PixelFormat,omitempty"`

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
}

func NewTiffExportOptions() *TiffExportOptions {
	instance := new(TiffExportOptions)
	instance.Compression = ""
	instance.PixelFormat = ""
	instance.NotesPosition = ""
	instance.CommentsPosition = ""
	return instance
}

func (this *TiffExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *TiffExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *TiffExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *TiffExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *TiffExportOptions) GetFormat() string {
	return this.Format
}

func (this *TiffExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *TiffExportOptions) GetHeight() int32 {
	return this.Height
}

func (this *TiffExportOptions) SetHeight(newValue int32) {
	this.Height = newValue
}
func (this *TiffExportOptions) GetWidth() int32 {
	return this.Width
}

func (this *TiffExportOptions) SetWidth(newValue int32) {
	this.Width = newValue
}
func (this *TiffExportOptions) GetCompression() string {
	return this.Compression
}

func (this *TiffExportOptions) SetCompression(newValue string) {
	this.Compression = newValue
}
func (this *TiffExportOptions) GetDpiX() int32 {
	return this.DpiX
}

func (this *TiffExportOptions) SetDpiX(newValue int32) {
	this.DpiX = newValue
}
func (this *TiffExportOptions) GetDpiY() int32 {
	return this.DpiY
}

func (this *TiffExportOptions) SetDpiY(newValue int32) {
	this.DpiY = newValue
}
func (this *TiffExportOptions) GetShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this *TiffExportOptions) SetShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this *TiffExportOptions) GetPixelFormat() string {
	return this.PixelFormat
}

func (this *TiffExportOptions) SetPixelFormat(newValue string) {
	this.PixelFormat = newValue
}
func (this *TiffExportOptions) GetNotesPosition() string {
	return this.NotesPosition
}

func (this *TiffExportOptions) SetNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this *TiffExportOptions) GetCommentsPosition() string {
	return this.CommentsPosition
}

func (this *TiffExportOptions) SetCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this *TiffExportOptions) GetCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this *TiffExportOptions) SetCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this *TiffExportOptions) GetCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this *TiffExportOptions) SetCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this *TiffExportOptions) GetShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this *TiffExportOptions) SetShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *TiffExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight int32
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth int32
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	this.Compression = ""
	if valCompression, ok := objMap["compression"]; ok {
		if valCompression != nil {
			var valueForCompression string
			err = json.Unmarshal(*valCompression, &valueForCompression)
			if err != nil {
				var valueForCompressionInt int32
				err = json.Unmarshal(*valCompression, &valueForCompressionInt)
				if err != nil {
					return err
				}
				this.Compression = string(valueForCompressionInt)
			} else {
				this.Compression = valueForCompression
			}
		}
	}
	if valCompressionCap, ok := objMap["Compression"]; ok {
		if valCompressionCap != nil {
			var valueForCompression string
			err = json.Unmarshal(*valCompressionCap, &valueForCompression)
			if err != nil {
				var valueForCompressionInt int32
				err = json.Unmarshal(*valCompressionCap, &valueForCompressionInt)
				if err != nil {
					return err
				}
				this.Compression = string(valueForCompressionInt)
			} else {
				this.Compression = valueForCompression
			}
		}
	}
	
	if valDpiX, ok := objMap["dpiX"]; ok {
		if valDpiX != nil {
			var valueForDpiX int32
			err = json.Unmarshal(*valDpiX, &valueForDpiX)
			if err != nil {
				return err
			}
			this.DpiX = valueForDpiX
		}
	}
	if valDpiXCap, ok := objMap["DpiX"]; ok {
		if valDpiXCap != nil {
			var valueForDpiX int32
			err = json.Unmarshal(*valDpiXCap, &valueForDpiX)
			if err != nil {
				return err
			}
			this.DpiX = valueForDpiX
		}
	}
	
	if valDpiY, ok := objMap["dpiY"]; ok {
		if valDpiY != nil {
			var valueForDpiY int32
			err = json.Unmarshal(*valDpiY, &valueForDpiY)
			if err != nil {
				return err
			}
			this.DpiY = valueForDpiY
		}
	}
	if valDpiYCap, ok := objMap["DpiY"]; ok {
		if valDpiYCap != nil {
			var valueForDpiY int32
			err = json.Unmarshal(*valDpiYCap, &valueForDpiY)
			if err != nil {
				return err
			}
			this.DpiY = valueForDpiY
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
	this.PixelFormat = ""
	if valPixelFormat, ok := objMap["pixelFormat"]; ok {
		if valPixelFormat != nil {
			var valueForPixelFormat string
			err = json.Unmarshal(*valPixelFormat, &valueForPixelFormat)
			if err != nil {
				var valueForPixelFormatInt int32
				err = json.Unmarshal(*valPixelFormat, &valueForPixelFormatInt)
				if err != nil {
					return err
				}
				this.PixelFormat = string(valueForPixelFormatInt)
			} else {
				this.PixelFormat = valueForPixelFormat
			}
		}
	}
	if valPixelFormatCap, ok := objMap["PixelFormat"]; ok {
		if valPixelFormatCap != nil {
			var valueForPixelFormat string
			err = json.Unmarshal(*valPixelFormatCap, &valueForPixelFormat)
			if err != nil {
				var valueForPixelFormatInt int32
				err = json.Unmarshal(*valPixelFormatCap, &valueForPixelFormatInt)
				if err != nil {
					return err
				}
				this.PixelFormat = string(valueForPixelFormatInt)
			} else {
				this.PixelFormat = valueForPixelFormat
			}
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

	return nil
}
