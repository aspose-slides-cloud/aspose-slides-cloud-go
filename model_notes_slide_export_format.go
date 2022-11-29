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

// NotesSlideExportFormat : Represents a format for notes slide export.
type NotesSlideExportFormat string

// List of NotesSlideExportFormat NotesSlideExportFormat
const (
	NotesSlideExportFormat_Jpeg NotesSlideExportFormat = "Jpeg"
	NotesSlideExportFormat_Png NotesSlideExportFormat = "Png"
	NotesSlideExportFormat_Gif NotesSlideExportFormat = "Gif"
	NotesSlideExportFormat_Bmp NotesSlideExportFormat = "Bmp"
	NotesSlideExportFormat_Tiff NotesSlideExportFormat = "Tiff"
	NotesSlideExportFormat_Html NotesSlideExportFormat = "Html"
	NotesSlideExportFormat_Pdf NotesSlideExportFormat = "Pdf"
	NotesSlideExportFormat_Xps NotesSlideExportFormat = "Xps"
	NotesSlideExportFormat_Pptx NotesSlideExportFormat = "Pptx"
	NotesSlideExportFormat_Odp NotesSlideExportFormat = "Odp"
	NotesSlideExportFormat_Otp NotesSlideExportFormat = "Otp"
	NotesSlideExportFormat_Ppt NotesSlideExportFormat = "Ppt"
	NotesSlideExportFormat_Pps NotesSlideExportFormat = "Pps"
	NotesSlideExportFormat_Ppsx NotesSlideExportFormat = "Ppsx"
	NotesSlideExportFormat_Pptm NotesSlideExportFormat = "Pptm"
	NotesSlideExportFormat_Ppsm NotesSlideExportFormat = "Ppsm"
	NotesSlideExportFormat_Potx NotesSlideExportFormat = "Potx"
	NotesSlideExportFormat_Pot NotesSlideExportFormat = "Pot"
	NotesSlideExportFormat_Potm NotesSlideExportFormat = "Potm"
	NotesSlideExportFormat_Svg NotesSlideExportFormat = "Svg"
	NotesSlideExportFormat_Fodp NotesSlideExportFormat = "Fodp"
	NotesSlideExportFormat_Xaml NotesSlideExportFormat = "Xaml"
	NotesSlideExportFormat_Html5 NotesSlideExportFormat = "Html5"
)

func NotesSlideExportFormat_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Jpeg)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Png)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Gif)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Bmp)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Tiff)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Html)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Pdf)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Xps)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Pptx)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Odp)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Otp)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Ppt)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Pps)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Ppsx)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Pptm)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Ppsm)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Potx)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Pot)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Potm)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Svg)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Fodp)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Xaml)) || strings.ToLower(value) == strings.ToLower(string(NotesSlideExportFormat_Html5))
}
