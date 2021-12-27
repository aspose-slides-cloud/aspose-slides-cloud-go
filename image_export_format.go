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
	"strings"
)

// ImageExportFormat : Represents a format for image export.
type ImageExportFormat string

// List of ImageExportFormat ImageExportFormat
const (
	ImageExportFormat_Jpeg ImageExportFormat = "Jpeg"
	ImageExportFormat_Png ImageExportFormat = "Png"
	ImageExportFormat_Gif ImageExportFormat = "Gif"
	ImageExportFormat_Bmp ImageExportFormat = "Bmp"
	ImageExportFormat_Tiff ImageExportFormat = "Tiff"
)

func ImageExportFormat_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(ImageExportFormat_Jpeg)) || strings.ToLower(value) == strings.ToLower(string(ImageExportFormat_Png)) || strings.ToLower(value) == strings.ToLower(string(ImageExportFormat_Gif)) || strings.ToLower(value) == strings.ToLower(string(ImageExportFormat_Bmp)) || strings.ToLower(value) == strings.ToLower(string(ImageExportFormat_Tiff))
}
