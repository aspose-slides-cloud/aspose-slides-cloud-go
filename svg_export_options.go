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

// Provides options that control how a presentation is saved in SVG format.
type ISvgExportOptions interface {

	// Default regular font for rendering the presentation. 
	getDefaultRegularFont() string
	setDefaultRegularFont(newValue string)

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getHeight() int32
	setHeight(newValue int32)

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	getWidth() int32
	setWidth(newValue int32)

	//          Export format.          
	getFormat() string
	setFormat(newValue string)

	// Determines whether the text on a slide will be saved as graphics.
	getVectorizeText() bool
	setVectorizeText(newValue bool)

	// Returns or sets the lower resolution limit for metafile rasterization.
	getMetafileRasterizationDpi() int32
	setMetafileRasterizationDpi(newValue int32)

	// Determines whether the 3D text is disabled in SVG.
	getDisable3DText() bool
	setDisable3DText(newValue bool)

	// Disables splitting FromCornerX and FromCenter gradients.
	getDisableGradientSplit() bool
	setDisableGradientSplit(newValue bool)

	// SVG 1.1 lacks ability to define insets for markers. Aspose.Slides SVG writing engine has workaround for that problem: it crops end of line with arrow, so, line doesn't overlap markers. This option switches off such behavior.
	getDisableLineEndCropping() bool
	setDisableLineEndCropping(newValue bool)

	// Determines JPEG encoding quality.
	getJpegQuality() int32
	setJpegQuality(newValue int32)

	// Represents the pictures compression level
	getPicturesCompression() string
	setPicturesCompression(newValue string)

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	getDeletePicturesCroppedAreas() bool
	setDeletePicturesCroppedAreas(newValue bool)

	// Determines a way of handling externally loaded fonts.
	getExternalFontsHandling() string
	setExternalFontsHandling(newValue string)
}

type SvgExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Height int32 `json:"Height,omitempty"`

	// Gets or sets the height of slides in the output format, e.g. image size, pdf page size etc.
	Width int32 `json:"Width,omitempty"`

	//          Export format.          
	Format string `json:"Format,omitempty"`

	// Determines whether the text on a slide will be saved as graphics.
	VectorizeText bool `json:"VectorizeText"`

	// Returns or sets the lower resolution limit for metafile rasterization.
	MetafileRasterizationDpi int32 `json:"MetafileRasterizationDpi,omitempty"`

	// Determines whether the 3D text is disabled in SVG.
	Disable3DText bool `json:"Disable3DText"`

	// Disables splitting FromCornerX and FromCenter gradients.
	DisableGradientSplit bool `json:"DisableGradientSplit"`

	// SVG 1.1 lacks ability to define insets for markers. Aspose.Slides SVG writing engine has workaround for that problem: it crops end of line with arrow, so, line doesn't overlap markers. This option switches off such behavior.
	DisableLineEndCropping bool `json:"DisableLineEndCropping"`

	// Determines JPEG encoding quality.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Represents the pictures compression level
	PicturesCompression string `json:"PicturesCompression,omitempty"`

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	DeletePicturesCroppedAreas bool `json:"DeletePicturesCroppedAreas"`

	// Determines a way of handling externally loaded fonts.
	ExternalFontsHandling string `json:"ExternalFontsHandling,omitempty"`
}

func NewSvgExportOptions() *SvgExportOptions {
	instance := new(SvgExportOptions)
	instance.PicturesCompression = ""
	instance.ExternalFontsHandling = ""
	return instance
}

func (this *SvgExportOptions) getDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *SvgExportOptions) setDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *SvgExportOptions) getHeight() int32 {
	return this.Height
}

func (this *SvgExportOptions) setHeight(newValue int32) {
	this.Height = newValue
}
func (this *SvgExportOptions) getWidth() int32 {
	return this.Width
}

func (this *SvgExportOptions) setWidth(newValue int32) {
	this.Width = newValue
}
func (this *SvgExportOptions) getFormat() string {
	return this.Format
}

func (this *SvgExportOptions) setFormat(newValue string) {
	this.Format = newValue
}
func (this *SvgExportOptions) getVectorizeText() bool {
	return this.VectorizeText
}

func (this *SvgExportOptions) setVectorizeText(newValue bool) {
	this.VectorizeText = newValue
}
func (this *SvgExportOptions) getMetafileRasterizationDpi() int32 {
	return this.MetafileRasterizationDpi
}

func (this *SvgExportOptions) setMetafileRasterizationDpi(newValue int32) {
	this.MetafileRasterizationDpi = newValue
}
func (this *SvgExportOptions) getDisable3DText() bool {
	return this.Disable3DText
}

func (this *SvgExportOptions) setDisable3DText(newValue bool) {
	this.Disable3DText = newValue
}
func (this *SvgExportOptions) getDisableGradientSplit() bool {
	return this.DisableGradientSplit
}

func (this *SvgExportOptions) setDisableGradientSplit(newValue bool) {
	this.DisableGradientSplit = newValue
}
func (this *SvgExportOptions) getDisableLineEndCropping() bool {
	return this.DisableLineEndCropping
}

func (this *SvgExportOptions) setDisableLineEndCropping(newValue bool) {
	this.DisableLineEndCropping = newValue
}
func (this *SvgExportOptions) getJpegQuality() int32 {
	return this.JpegQuality
}

func (this *SvgExportOptions) setJpegQuality(newValue int32) {
	this.JpegQuality = newValue
}
func (this *SvgExportOptions) getPicturesCompression() string {
	return this.PicturesCompression
}

func (this *SvgExportOptions) setPicturesCompression(newValue string) {
	this.PicturesCompression = newValue
}
func (this *SvgExportOptions) getDeletePicturesCroppedAreas() bool {
	return this.DeletePicturesCroppedAreas
}

func (this *SvgExportOptions) setDeletePicturesCroppedAreas(newValue bool) {
	this.DeletePicturesCroppedAreas = newValue
}
func (this *SvgExportOptions) getExternalFontsHandling() string {
	return this.ExternalFontsHandling
}

func (this *SvgExportOptions) setExternalFontsHandling(newValue string) {
	this.ExternalFontsHandling = newValue
}

func (this *SvgExportOptions) UnmarshalJSON(b []byte) error {
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
	
	if valVectorizeText, ok := objMap["vectorizeText"]; ok {
		if valVectorizeText != nil {
			var valueForVectorizeText bool
			err = json.Unmarshal(*valVectorizeText, &valueForVectorizeText)
			if err != nil {
				return err
			}
			this.VectorizeText = valueForVectorizeText
		}
	}
	if valVectorizeTextCap, ok := objMap["VectorizeText"]; ok {
		if valVectorizeTextCap != nil {
			var valueForVectorizeText bool
			err = json.Unmarshal(*valVectorizeTextCap, &valueForVectorizeText)
			if err != nil {
				return err
			}
			this.VectorizeText = valueForVectorizeText
		}
	}
	
	if valMetafileRasterizationDpi, ok := objMap["metafileRasterizationDpi"]; ok {
		if valMetafileRasterizationDpi != nil {
			var valueForMetafileRasterizationDpi int32
			err = json.Unmarshal(*valMetafileRasterizationDpi, &valueForMetafileRasterizationDpi)
			if err != nil {
				return err
			}
			this.MetafileRasterizationDpi = valueForMetafileRasterizationDpi
		}
	}
	if valMetafileRasterizationDpiCap, ok := objMap["MetafileRasterizationDpi"]; ok {
		if valMetafileRasterizationDpiCap != nil {
			var valueForMetafileRasterizationDpi int32
			err = json.Unmarshal(*valMetafileRasterizationDpiCap, &valueForMetafileRasterizationDpi)
			if err != nil {
				return err
			}
			this.MetafileRasterizationDpi = valueForMetafileRasterizationDpi
		}
	}
	
	if valDisable3DText, ok := objMap["disable3DText"]; ok {
		if valDisable3DText != nil {
			var valueForDisable3DText bool
			err = json.Unmarshal(*valDisable3DText, &valueForDisable3DText)
			if err != nil {
				return err
			}
			this.Disable3DText = valueForDisable3DText
		}
	}
	if valDisable3DTextCap, ok := objMap["Disable3DText"]; ok {
		if valDisable3DTextCap != nil {
			var valueForDisable3DText bool
			err = json.Unmarshal(*valDisable3DTextCap, &valueForDisable3DText)
			if err != nil {
				return err
			}
			this.Disable3DText = valueForDisable3DText
		}
	}
	
	if valDisableGradientSplit, ok := objMap["disableGradientSplit"]; ok {
		if valDisableGradientSplit != nil {
			var valueForDisableGradientSplit bool
			err = json.Unmarshal(*valDisableGradientSplit, &valueForDisableGradientSplit)
			if err != nil {
				return err
			}
			this.DisableGradientSplit = valueForDisableGradientSplit
		}
	}
	if valDisableGradientSplitCap, ok := objMap["DisableGradientSplit"]; ok {
		if valDisableGradientSplitCap != nil {
			var valueForDisableGradientSplit bool
			err = json.Unmarshal(*valDisableGradientSplitCap, &valueForDisableGradientSplit)
			if err != nil {
				return err
			}
			this.DisableGradientSplit = valueForDisableGradientSplit
		}
	}
	
	if valDisableLineEndCropping, ok := objMap["disableLineEndCropping"]; ok {
		if valDisableLineEndCropping != nil {
			var valueForDisableLineEndCropping bool
			err = json.Unmarshal(*valDisableLineEndCropping, &valueForDisableLineEndCropping)
			if err != nil {
				return err
			}
			this.DisableLineEndCropping = valueForDisableLineEndCropping
		}
	}
	if valDisableLineEndCroppingCap, ok := objMap["DisableLineEndCropping"]; ok {
		if valDisableLineEndCroppingCap != nil {
			var valueForDisableLineEndCropping bool
			err = json.Unmarshal(*valDisableLineEndCroppingCap, &valueForDisableLineEndCropping)
			if err != nil {
				return err
			}
			this.DisableLineEndCropping = valueForDisableLineEndCropping
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
				var valueForPicturesCompressionInt int32
				err = json.Unmarshal(*valPicturesCompression, &valueForPicturesCompressionInt)
				if err != nil {
					return err
				}
				this.PicturesCompression = string(valueForPicturesCompressionInt)
			} else {
				this.PicturesCompression = valueForPicturesCompression
			}
		}
	}
	if valPicturesCompressionCap, ok := objMap["PicturesCompression"]; ok {
		if valPicturesCompressionCap != nil {
			var valueForPicturesCompression string
			err = json.Unmarshal(*valPicturesCompressionCap, &valueForPicturesCompression)
			if err != nil {
				var valueForPicturesCompressionInt int32
				err = json.Unmarshal(*valPicturesCompressionCap, &valueForPicturesCompressionInt)
				if err != nil {
					return err
				}
				this.PicturesCompression = string(valueForPicturesCompressionInt)
			} else {
				this.PicturesCompression = valueForPicturesCompression
			}
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
	this.ExternalFontsHandling = ""
	if valExternalFontsHandling, ok := objMap["externalFontsHandling"]; ok {
		if valExternalFontsHandling != nil {
			var valueForExternalFontsHandling string
			err = json.Unmarshal(*valExternalFontsHandling, &valueForExternalFontsHandling)
			if err != nil {
				var valueForExternalFontsHandlingInt int32
				err = json.Unmarshal(*valExternalFontsHandling, &valueForExternalFontsHandlingInt)
				if err != nil {
					return err
				}
				this.ExternalFontsHandling = string(valueForExternalFontsHandlingInt)
			} else {
				this.ExternalFontsHandling = valueForExternalFontsHandling
			}
		}
	}
	if valExternalFontsHandlingCap, ok := objMap["ExternalFontsHandling"]; ok {
		if valExternalFontsHandlingCap != nil {
			var valueForExternalFontsHandling string
			err = json.Unmarshal(*valExternalFontsHandlingCap, &valueForExternalFontsHandling)
			if err != nil {
				var valueForExternalFontsHandlingInt int32
				err = json.Unmarshal(*valExternalFontsHandlingCap, &valueForExternalFontsHandlingInt)
				if err != nil {
					return err
				}
				this.ExternalFontsHandling = string(valueForExternalFontsHandlingInt)
			} else {
				this.ExternalFontsHandling = valueForExternalFontsHandling
			}
		}
	}

	return nil
}
