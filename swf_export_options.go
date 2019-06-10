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

	getFormat() string
	setFormat(newValue string)

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	getShowHiddenSlides() bool
	setShowHiddenSlides(newValue bool)

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	getCompressed() bool
	setCompressed(newValue bool)

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	getViewerIncluded() bool
	setViewerIncluded(newValue bool)

	// Specifies whether border around pages should be shown. Default is true. 
	getShowPageBorder() bool
	setShowPageBorder(newValue bool)

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	getShowFullScreen() bool
	setShowFullScreen(newValue bool)

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	getShowPageStepper() bool
	setShowPageStepper(newValue bool)

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	getShowSearch() bool
	setShowSearch(newValue bool)

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	getShowTopPane() bool
	setShowTopPane(newValue bool)

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	getShowBottomPane() bool
	setShowBottomPane(newValue bool)

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	getShowLeftPane() bool
	setShowLeftPane(newValue bool)

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	getStartOpenLeftPane() bool
	setStartOpenLeftPane(newValue bool)

	// Enable/disable context menu. Default is true. 
	getEnableContextMenu() bool
	setEnableContextMenu(newValue bool)

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	getLogoImage() string
	setLogoImage(newValue string)

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	getLogoLink() string
	setLogoLink(newValue string)

	// Specifies the quality of JPEG images. Default is 95.
	getJpegQuality() int32
	setJpegQuality(newValue int32)

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

type SwfExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides"`

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	Compressed bool `json:"Compressed"`

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	ViewerIncluded bool `json:"ViewerIncluded"`

	// Specifies whether border around pages should be shown. Default is true. 
	ShowPageBorder bool `json:"ShowPageBorder"`

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	ShowFullScreen bool `json:"ShowFullScreen"`

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	ShowPageStepper bool `json:"ShowPageStepper"`

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	ShowSearch bool `json:"ShowSearch"`

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	ShowTopPane bool `json:"ShowTopPane"`

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	ShowBottomPane bool `json:"ShowBottomPane"`

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	ShowLeftPane bool `json:"ShowLeftPane"`

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	StartOpenLeftPane bool `json:"StartOpenLeftPane"`

	// Enable/disable context menu. Default is true. 
	EnableContextMenu bool `json:"EnableContextMenu"`

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	LogoImage string `json:"LogoImage,omitempty"`

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a LogoImage is specified. 
	LogoLink string `json:"LogoLink,omitempty"`

	// Specifies the quality of JPEG images. Default is 95.
	JpegQuality int32 `json:"JpegQuality"`

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

func (this SwfExportOptions) getFormat() string {
	return this.Format
}

func (this SwfExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this SwfExportOptions) getShowHiddenSlides() bool {
	return this.ShowHiddenSlides
}

func (this SwfExportOptions) setShowHiddenSlides(newValue bool) {
	this.ShowHiddenSlides = newValue
}
func (this SwfExportOptions) getCompressed() bool {
	return this.Compressed
}

func (this SwfExportOptions) setCompressed(newValue bool) {
	this.Compressed = newValue
}
func (this SwfExportOptions) getViewerIncluded() bool {
	return this.ViewerIncluded
}

func (this SwfExportOptions) setViewerIncluded(newValue bool) {
	this.ViewerIncluded = newValue
}
func (this SwfExportOptions) getShowPageBorder() bool {
	return this.ShowPageBorder
}

func (this SwfExportOptions) setShowPageBorder(newValue bool) {
	this.ShowPageBorder = newValue
}
func (this SwfExportOptions) getShowFullScreen() bool {
	return this.ShowFullScreen
}

func (this SwfExportOptions) setShowFullScreen(newValue bool) {
	this.ShowFullScreen = newValue
}
func (this SwfExportOptions) getShowPageStepper() bool {
	return this.ShowPageStepper
}

func (this SwfExportOptions) setShowPageStepper(newValue bool) {
	this.ShowPageStepper = newValue
}
func (this SwfExportOptions) getShowSearch() bool {
	return this.ShowSearch
}

func (this SwfExportOptions) setShowSearch(newValue bool) {
	this.ShowSearch = newValue
}
func (this SwfExportOptions) getShowTopPane() bool {
	return this.ShowTopPane
}

func (this SwfExportOptions) setShowTopPane(newValue bool) {
	this.ShowTopPane = newValue
}
func (this SwfExportOptions) getShowBottomPane() bool {
	return this.ShowBottomPane
}

func (this SwfExportOptions) setShowBottomPane(newValue bool) {
	this.ShowBottomPane = newValue
}
func (this SwfExportOptions) getShowLeftPane() bool {
	return this.ShowLeftPane
}

func (this SwfExportOptions) setShowLeftPane(newValue bool) {
	this.ShowLeftPane = newValue
}
func (this SwfExportOptions) getStartOpenLeftPane() bool {
	return this.StartOpenLeftPane
}

func (this SwfExportOptions) setStartOpenLeftPane(newValue bool) {
	this.StartOpenLeftPane = newValue
}
func (this SwfExportOptions) getEnableContextMenu() bool {
	return this.EnableContextMenu
}

func (this SwfExportOptions) setEnableContextMenu(newValue bool) {
	this.EnableContextMenu = newValue
}
func (this SwfExportOptions) getLogoImage() string {
	return this.LogoImage
}

func (this SwfExportOptions) setLogoImage(newValue string) {
	this.LogoImage = newValue
}
func (this SwfExportOptions) getLogoLink() string {
	return this.LogoLink
}

func (this SwfExportOptions) setLogoLink(newValue string) {
	this.LogoLink = newValue
}
func (this SwfExportOptions) getJpegQuality() int32 {
	return this.JpegQuality
}

func (this SwfExportOptions) setJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this SwfExportOptions) getNotesPosition() string {
	return this.NotesPosition
}

func (this SwfExportOptions) setNotesPosition(newValue string) {
	this.NotesPosition = newValue
}
func (this SwfExportOptions) getCommentsPosition() string {
	return this.CommentsPosition
}

func (this SwfExportOptions) setCommentsPosition(newValue string) {
	this.CommentsPosition = newValue
}
func (this SwfExportOptions) getCommentsAreaWidth() int32 {
	return this.CommentsAreaWidth
}

func (this SwfExportOptions) setCommentsAreaWidth(newValue int32) {
	this.CommentsAreaWidth = newValue
}
func (this SwfExportOptions) getCommentsAreaColor() string {
	return this.CommentsAreaColor
}

func (this SwfExportOptions) setCommentsAreaColor(newValue string) {
	this.CommentsAreaColor = newValue
}
func (this SwfExportOptions) getShowCommentsByNoAuthor() bool {
	return this.ShowCommentsByNoAuthor
}

func (this SwfExportOptions) setShowCommentsByNoAuthor(newValue bool) {
	this.ShowCommentsByNoAuthor = newValue
}

func (this *SwfExportOptions) UnmarshalJSON(b []byte) error {
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

	if valCompressed, ok := objMap["Compressed"]; ok {
		if valCompressed != nil {
			var valueForCompressed bool
			err = json.Unmarshal(*valCompressed, &valueForCompressed)
			if err != nil {
				return err
			}
			this.Compressed = valueForCompressed
		}
	}

	if valViewerIncluded, ok := objMap["ViewerIncluded"]; ok {
		if valViewerIncluded != nil {
			var valueForViewerIncluded bool
			err = json.Unmarshal(*valViewerIncluded, &valueForViewerIncluded)
			if err != nil {
				return err
			}
			this.ViewerIncluded = valueForViewerIncluded
		}
	}

	if valShowPageBorder, ok := objMap["ShowPageBorder"]; ok {
		if valShowPageBorder != nil {
			var valueForShowPageBorder bool
			err = json.Unmarshal(*valShowPageBorder, &valueForShowPageBorder)
			if err != nil {
				return err
			}
			this.ShowPageBorder = valueForShowPageBorder
		}
	}

	if valShowFullScreen, ok := objMap["ShowFullScreen"]; ok {
		if valShowFullScreen != nil {
			var valueForShowFullScreen bool
			err = json.Unmarshal(*valShowFullScreen, &valueForShowFullScreen)
			if err != nil {
				return err
			}
			this.ShowFullScreen = valueForShowFullScreen
		}
	}

	if valShowPageStepper, ok := objMap["ShowPageStepper"]; ok {
		if valShowPageStepper != nil {
			var valueForShowPageStepper bool
			err = json.Unmarshal(*valShowPageStepper, &valueForShowPageStepper)
			if err != nil {
				return err
			}
			this.ShowPageStepper = valueForShowPageStepper
		}
	}

	if valShowSearch, ok := objMap["ShowSearch"]; ok {
		if valShowSearch != nil {
			var valueForShowSearch bool
			err = json.Unmarshal(*valShowSearch, &valueForShowSearch)
			if err != nil {
				return err
			}
			this.ShowSearch = valueForShowSearch
		}
	}

	if valShowTopPane, ok := objMap["ShowTopPane"]; ok {
		if valShowTopPane != nil {
			var valueForShowTopPane bool
			err = json.Unmarshal(*valShowTopPane, &valueForShowTopPane)
			if err != nil {
				return err
			}
			this.ShowTopPane = valueForShowTopPane
		}
	}

	if valShowBottomPane, ok := objMap["ShowBottomPane"]; ok {
		if valShowBottomPane != nil {
			var valueForShowBottomPane bool
			err = json.Unmarshal(*valShowBottomPane, &valueForShowBottomPane)
			if err != nil {
				return err
			}
			this.ShowBottomPane = valueForShowBottomPane
		}
	}

	if valShowLeftPane, ok := objMap["ShowLeftPane"]; ok {
		if valShowLeftPane != nil {
			var valueForShowLeftPane bool
			err = json.Unmarshal(*valShowLeftPane, &valueForShowLeftPane)
			if err != nil {
				return err
			}
			this.ShowLeftPane = valueForShowLeftPane
		}
	}

	if valStartOpenLeftPane, ok := objMap["StartOpenLeftPane"]; ok {
		if valStartOpenLeftPane != nil {
			var valueForStartOpenLeftPane bool
			err = json.Unmarshal(*valStartOpenLeftPane, &valueForStartOpenLeftPane)
			if err != nil {
				return err
			}
			this.StartOpenLeftPane = valueForStartOpenLeftPane
		}
	}

	if valEnableContextMenu, ok := objMap["EnableContextMenu"]; ok {
		if valEnableContextMenu != nil {
			var valueForEnableContextMenu bool
			err = json.Unmarshal(*valEnableContextMenu, &valueForEnableContextMenu)
			if err != nil {
				return err
			}
			this.EnableContextMenu = valueForEnableContextMenu
		}
	}

	if valLogoImage, ok := objMap["LogoImage"]; ok {
		if valLogoImage != nil {
			var valueForLogoImage string
			err = json.Unmarshal(*valLogoImage, &valueForLogoImage)
			if err != nil {
				return err
			}
			this.LogoImage = valueForLogoImage
		}
	}

	if valLogoLink, ok := objMap["LogoLink"]; ok {
		if valLogoLink != nil {
			var valueForLogoLink string
			err = json.Unmarshal(*valLogoLink, &valueForLogoLink)
			if err != nil {
				return err
			}
			this.LogoLink = valueForLogoLink
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
