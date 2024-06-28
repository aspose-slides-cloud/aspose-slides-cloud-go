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

// Provides options that control how a presentation is saved in SWF format.
type ISwfExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Gets of sets list of font substitution rules.
	GetFontSubstRules() []IFontSubstRule
	SetFontSubstRules(newValue []IFontSubstRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	GetShowHiddenSlides() *bool
	SetShowHiddenSlides(newValue *bool)

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	GetCompressed() *bool
	SetCompressed(newValue *bool)

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	GetViewerIncluded() *bool
	SetViewerIncluded(newValue *bool)

	// Specifies whether border around pages should be shown. Default is true. 
	GetShowPageBorder() *bool
	SetShowPageBorder(newValue *bool)

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	GetShowFullScreen() *bool
	SetShowFullScreen(newValue *bool)

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	GetShowPageStepper() *bool
	SetShowPageStepper(newValue *bool)

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	GetShowSearch() *bool
	SetShowSearch(newValue *bool)

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	GetShowTopPane() *bool
	SetShowTopPane(newValue *bool)

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	GetShowBottomPane() *bool
	SetShowBottomPane(newValue *bool)

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	GetShowLeftPane() *bool
	SetShowLeftPane(newValue *bool)

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	GetStartOpenLeftPane() *bool
	SetStartOpenLeftPane(newValue *bool)

	// Enable/disable context menu. Default is true. 
	GetEnableContextMenu() *bool
	SetEnableContextMenu(newValue *bool)

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	GetLogoImage() string
	SetLogoImage(newValue string)

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	GetLogoLink() string
	SetLogoLink(newValue string)

	// Specifies the quality of JPEG images. Default is 95.
	GetJpegQuality() int32
	SetJpegQuality(newValue int32)

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
	GetShowCommentsByNoAuthor() *bool
	SetShowCommentsByNoAuthor(newValue *bool)
}

type SwfExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides *bool `json:"ShowHiddenSlides"`

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	Compressed *bool `json:"Compressed"`

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	ViewerIncluded *bool `json:"ViewerIncluded"`

	// Specifies whether border around pages should be shown. Default is true. 
	ShowPageBorder *bool `json:"ShowPageBorder"`

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	ShowFullScreen *bool `json:"ShowFullScreen"`

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	ShowPageStepper *bool `json:"ShowPageStepper"`

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	ShowSearch *bool `json:"ShowSearch"`

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	ShowTopPane *bool `json:"ShowTopPane"`

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	ShowBottomPane *bool `json:"ShowBottomPane"`

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	ShowLeftPane *bool `json:"ShowLeftPane"`

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	StartOpenLeftPane *bool `json:"StartOpenLeftPane"`

	// Enable/disable context menu. Default is true. 
	EnableContextMenu *bool `json:"EnableContextMenu"`

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	LogoImage string `json:"LogoImage,omitempty"`

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	LogoLink string `json:"LogoLink,omitempty"`

	// Specifies the quality of JPEG images. Default is 95.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Gets or sets the position of the notes on the page.
	NotesPosition string `json:"NotesPosition,omitempty"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition string `json:"CommentsPosition,omitempty"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth,omitempty"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor *bool `json:"ShowCommentsByNoAuthor"`
}

func NewSwfExportOptions() *SwfExportOptions {
	instance := new(SwfExportOptions)
	return instance
}

func (this *SwfExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *SwfExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *SwfExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *SwfExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *SwfExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *SwfExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *SwfExportOptions) GetFormat() string {
	return this.Format
}

func (this *SwfExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *SwfExportOptions) GetShowHiddenSlides() *bool {
	return this.ShowHiddenSlides
}

func (this *SwfExportOptions) SetShowHiddenSlides(newValue *bool) {
	this.ShowHiddenSlides = newValue
}
func (this *SwfExportOptions) GetCompressed() *bool {
	return this.Compressed
}

func (this *SwfExportOptions) SetCompressed(newValue *bool) {
	this.Compressed = newValue
}
func (this *SwfExportOptions) GetViewerIncluded() *bool {
	return this.ViewerIncluded
}

func (this *SwfExportOptions) SetViewerIncluded(newValue *bool) {
	this.ViewerIncluded = newValue
}
func (this *SwfExportOptions) GetShowPageBorder() *bool {
	return this.ShowPageBorder
}

func (this *SwfExportOptions) SetShowPageBorder(newValue *bool) {
	this.ShowPageBorder = newValue
}
func (this *SwfExportOptions) GetShowFullScreen() *bool {
	return this.ShowFullScreen
}

func (this *SwfExportOptions) SetShowFullScreen(newValue *bool) {
	this.ShowFullScreen = newValue
}
func (this *SwfExportOptions) GetShowPageStepper() *bool {
	return this.ShowPageStepper
}

func (this *SwfExportOptions) SetShowPageStepper(newValue *bool) {
	this.ShowPageStepper = newValue
}
func (this *SwfExportOptions) GetShowSearch() *bool {
	return this.ShowSearch
}

func (this *SwfExportOptions) SetShowSearch(newValue *bool) {
	this.ShowSearch = newValue
}
func (this *SwfExportOptions) GetShowTopPane() *bool {
	return this.ShowTopPane
}

func (this *SwfExportOptions) SetShowTopPane(newValue *bool) {
	this.ShowTopPane = newValue
}
func (this *SwfExportOptions) GetShowBottomPane() *bool {
	return this.ShowBottomPane
}

func (this *SwfExportOptions) SetShowBottomPane(newValue *bool) {
	this.ShowBottomPane = newValue
}
func (this *SwfExportOptions) GetShowLeftPane() *bool {
	return this.ShowLeftPane
}

func (this *SwfExportOptions) SetShowLeftPane(newValue *bool) {
	this.ShowLeftPane = newValue
}
func (this *SwfExportOptions) GetStartOpenLeftPane() *bool {
	return this.StartOpenLeftPane
}

func (this *SwfExportOptions) SetStartOpenLeftPane(newValue *bool) {
	this.StartOpenLeftPane = newValue
}
func (this *SwfExportOptions) GetEnableContextMenu() *bool {
	return this.EnableContextMenu
}

func (this *SwfExportOptions) SetEnableContextMenu(newValue *bool) {
	this.EnableContextMenu = newValue
}
func (this *SwfExportOptions) GetLogoImage() string {
	return this.LogoImage
}

func (this *SwfExportOptions) SetLogoImage(newValue string) {
	this.LogoImage = newValue
}
func (this *SwfExportOptions) GetLogoLink() string {
	return this.LogoLink
}

func (this *SwfExportOptions) SetLogoLink(newValue string) {
	this.LogoLink = newValue
}
func (this *SwfExportOptions) GetJpegQuality() int32 {
	return this.JpegQuality
}

func (this *SwfExportOptions) SetJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *SwfExportOptions) GetNotesPosition() string {
	return this.NotesPosition
}

func (this *SwfExportOptions) SetNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this *SwfExportOptions) GetCommentsPosition() string {
	return this.CommentsPosition
}

func (this *SwfExportOptions) SetCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this *SwfExportOptions) GetCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this *SwfExportOptions) SetCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this *SwfExportOptions) GetCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this *SwfExportOptions) SetCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this *SwfExportOptions) GetShowCommentsByNoAuthor() *bool {
	return this.ShowCommentsByNoAuthor
}

func (this *SwfExportOptions) SetShowCommentsByNoAuthor(newValue *bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *SwfExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valFontSubstRules, ok := objMap["fontSubstRules"]; ok {
		if valFontSubstRules != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRules, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	if valFontSubstRulesCap, ok := objMap["FontSubstRules"]; ok {
		if valFontSubstRulesCap != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRulesCap, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
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
	
	if valShowHiddenSlides, ok := objMap["showHiddenSlides"]; ok {
		if valShowHiddenSlides != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlides, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	if valShowHiddenSlidesCap, ok := objMap["ShowHiddenSlides"]; ok {
		if valShowHiddenSlidesCap != nil {
			var valueForShowHiddenSlides *bool
			err = json.Unmarshal(*valShowHiddenSlidesCap, &valueForShowHiddenSlides)
			if err != nil {
				return err
			}
			this.ShowHiddenSlides = valueForShowHiddenSlides
		}
	}
	
	if valCompressed, ok := objMap["compressed"]; ok {
		if valCompressed != nil {
			var valueForCompressed *bool
			err = json.Unmarshal(*valCompressed, &valueForCompressed)
			if err != nil {
				return err
			}
			this.Compressed = valueForCompressed
		}
	}
	if valCompressedCap, ok := objMap["Compressed"]; ok {
		if valCompressedCap != nil {
			var valueForCompressed *bool
			err = json.Unmarshal(*valCompressedCap, &valueForCompressed)
			if err != nil {
				return err
			}
			this.Compressed = valueForCompressed
		}
	}
	
	if valViewerIncluded, ok := objMap["viewerIncluded"]; ok {
		if valViewerIncluded != nil {
			var valueForViewerIncluded *bool
			err = json.Unmarshal(*valViewerIncluded, &valueForViewerIncluded)
			if err != nil {
				return err
			}
			this.ViewerIncluded = valueForViewerIncluded
		}
	}
	if valViewerIncludedCap, ok := objMap["ViewerIncluded"]; ok {
		if valViewerIncludedCap != nil {
			var valueForViewerIncluded *bool
			err = json.Unmarshal(*valViewerIncludedCap, &valueForViewerIncluded)
			if err != nil {
				return err
			}
			this.ViewerIncluded = valueForViewerIncluded
		}
	}
	
	if valShowPageBorder, ok := objMap["showPageBorder"]; ok {
		if valShowPageBorder != nil {
			var valueForShowPageBorder *bool
			err = json.Unmarshal(*valShowPageBorder, &valueForShowPageBorder)
			if err != nil {
				return err
			}
			this.ShowPageBorder = valueForShowPageBorder
		}
	}
	if valShowPageBorderCap, ok := objMap["ShowPageBorder"]; ok {
		if valShowPageBorderCap != nil {
			var valueForShowPageBorder *bool
			err = json.Unmarshal(*valShowPageBorderCap, &valueForShowPageBorder)
			if err != nil {
				return err
			}
			this.ShowPageBorder = valueForShowPageBorder
		}
	}
	
	if valShowFullScreen, ok := objMap["showFullScreen"]; ok {
		if valShowFullScreen != nil {
			var valueForShowFullScreen *bool
			err = json.Unmarshal(*valShowFullScreen, &valueForShowFullScreen)
			if err != nil {
				return err
			}
			this.ShowFullScreen = valueForShowFullScreen
		}
	}
	if valShowFullScreenCap, ok := objMap["ShowFullScreen"]; ok {
		if valShowFullScreenCap != nil {
			var valueForShowFullScreen *bool
			err = json.Unmarshal(*valShowFullScreenCap, &valueForShowFullScreen)
			if err != nil {
				return err
			}
			this.ShowFullScreen = valueForShowFullScreen
		}
	}
	
	if valShowPageStepper, ok := objMap["showPageStepper"]; ok {
		if valShowPageStepper != nil {
			var valueForShowPageStepper *bool
			err = json.Unmarshal(*valShowPageStepper, &valueForShowPageStepper)
			if err != nil {
				return err
			}
			this.ShowPageStepper = valueForShowPageStepper
		}
	}
	if valShowPageStepperCap, ok := objMap["ShowPageStepper"]; ok {
		if valShowPageStepperCap != nil {
			var valueForShowPageStepper *bool
			err = json.Unmarshal(*valShowPageStepperCap, &valueForShowPageStepper)
			if err != nil {
				return err
			}
			this.ShowPageStepper = valueForShowPageStepper
		}
	}
	
	if valShowSearch, ok := objMap["showSearch"]; ok {
		if valShowSearch != nil {
			var valueForShowSearch *bool
			err = json.Unmarshal(*valShowSearch, &valueForShowSearch)
			if err != nil {
				return err
			}
			this.ShowSearch = valueForShowSearch
		}
	}
	if valShowSearchCap, ok := objMap["ShowSearch"]; ok {
		if valShowSearchCap != nil {
			var valueForShowSearch *bool
			err = json.Unmarshal(*valShowSearchCap, &valueForShowSearch)
			if err != nil {
				return err
			}
			this.ShowSearch = valueForShowSearch
		}
	}
	
	if valShowTopPane, ok := objMap["showTopPane"]; ok {
		if valShowTopPane != nil {
			var valueForShowTopPane *bool
			err = json.Unmarshal(*valShowTopPane, &valueForShowTopPane)
			if err != nil {
				return err
			}
			this.ShowTopPane = valueForShowTopPane
		}
	}
	if valShowTopPaneCap, ok := objMap["ShowTopPane"]; ok {
		if valShowTopPaneCap != nil {
			var valueForShowTopPane *bool
			err = json.Unmarshal(*valShowTopPaneCap, &valueForShowTopPane)
			if err != nil {
				return err
			}
			this.ShowTopPane = valueForShowTopPane
		}
	}
	
	if valShowBottomPane, ok := objMap["showBottomPane"]; ok {
		if valShowBottomPane != nil {
			var valueForShowBottomPane *bool
			err = json.Unmarshal(*valShowBottomPane, &valueForShowBottomPane)
			if err != nil {
				return err
			}
			this.ShowBottomPane = valueForShowBottomPane
		}
	}
	if valShowBottomPaneCap, ok := objMap["ShowBottomPane"]; ok {
		if valShowBottomPaneCap != nil {
			var valueForShowBottomPane *bool
			err = json.Unmarshal(*valShowBottomPaneCap, &valueForShowBottomPane)
			if err != nil {
				return err
			}
			this.ShowBottomPane = valueForShowBottomPane
		}
	}
	
	if valShowLeftPane, ok := objMap["showLeftPane"]; ok {
		if valShowLeftPane != nil {
			var valueForShowLeftPane *bool
			err = json.Unmarshal(*valShowLeftPane, &valueForShowLeftPane)
			if err != nil {
				return err
			}
			this.ShowLeftPane = valueForShowLeftPane
		}
	}
	if valShowLeftPaneCap, ok := objMap["ShowLeftPane"]; ok {
		if valShowLeftPaneCap != nil {
			var valueForShowLeftPane *bool
			err = json.Unmarshal(*valShowLeftPaneCap, &valueForShowLeftPane)
			if err != nil {
				return err
			}
			this.ShowLeftPane = valueForShowLeftPane
		}
	}
	
	if valStartOpenLeftPane, ok := objMap["startOpenLeftPane"]; ok {
		if valStartOpenLeftPane != nil {
			var valueForStartOpenLeftPane *bool
			err = json.Unmarshal(*valStartOpenLeftPane, &valueForStartOpenLeftPane)
			if err != nil {
				return err
			}
			this.StartOpenLeftPane = valueForStartOpenLeftPane
		}
	}
	if valStartOpenLeftPaneCap, ok := objMap["StartOpenLeftPane"]; ok {
		if valStartOpenLeftPaneCap != nil {
			var valueForStartOpenLeftPane *bool
			err = json.Unmarshal(*valStartOpenLeftPaneCap, &valueForStartOpenLeftPane)
			if err != nil {
				return err
			}
			this.StartOpenLeftPane = valueForStartOpenLeftPane
		}
	}
	
	if valEnableContextMenu, ok := objMap["enableContextMenu"]; ok {
		if valEnableContextMenu != nil {
			var valueForEnableContextMenu *bool
			err = json.Unmarshal(*valEnableContextMenu, &valueForEnableContextMenu)
			if err != nil {
				return err
			}
			this.EnableContextMenu = valueForEnableContextMenu
		}
	}
	if valEnableContextMenuCap, ok := objMap["EnableContextMenu"]; ok {
		if valEnableContextMenuCap != nil {
			var valueForEnableContextMenu *bool
			err = json.Unmarshal(*valEnableContextMenuCap, &valueForEnableContextMenu)
			if err != nil {
				return err
			}
			this.EnableContextMenu = valueForEnableContextMenu
		}
	}
	
	if valLogoImage, ok := objMap["logoImage"]; ok {
		if valLogoImage != nil {
			var valueForLogoImage string
			err = json.Unmarshal(*valLogoImage, &valueForLogoImage)
			if err != nil {
				return err
			}
			this.LogoImage = valueForLogoImage
		}
	}
	if valLogoImageCap, ok := objMap["LogoImage"]; ok {
		if valLogoImageCap != nil {
			var valueForLogoImage string
			err = json.Unmarshal(*valLogoImageCap, &valueForLogoImage)
			if err != nil {
				return err
			}
			this.LogoImage = valueForLogoImage
		}
	}
	
	if valLogoLink, ok := objMap["logoLink"]; ok {
		if valLogoLink != nil {
			var valueForLogoLink string
			err = json.Unmarshal(*valLogoLink, &valueForLogoLink)
			if err != nil {
				return err
			}
			this.LogoLink = valueForLogoLink
		}
	}
	if valLogoLinkCap, ok := objMap["LogoLink"]; ok {
		if valLogoLinkCap != nil {
			var valueForLogoLink string
			err = json.Unmarshal(*valLogoLinkCap, &valueForLogoLink)
			if err != nil {
				return err
			}
			this.LogoLink = valueForLogoLink
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
			var valueForShowCommentsByNoAuthor *bool
			err = json.Unmarshal(*valShowCommentsByNoAuthor, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}
	if valShowCommentsByNoAuthorCap, ok := objMap["ShowCommentsByNoAuthor"]; ok {
		if valShowCommentsByNoAuthorCap != nil {
			var valueForShowCommentsByNoAuthor *bool
			err = json.Unmarshal(*valShowCommentsByNoAuthorCap, &valueForShowCommentsByNoAuthor)
			if err != nil {
				return err
			}
			this.ShowCommentsByNoAuthor = valueForShowCommentsByNoAuthor
		}
	}

	return nil
}
