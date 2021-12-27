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

// SpecialSlideType : 
type SpecialSlideType string

// List of SpecialSlideType SpecialSlideType
const (
	SpecialSlideType_MasterSlide SpecialSlideType = "MasterSlide"
	SpecialSlideType_LayoutSlide SpecialSlideType = "LayoutSlide"
	SpecialSlideType_NotesSlide SpecialSlideType = "NotesSlide"
)

func SpecialSlideType_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(SpecialSlideType_MasterSlide)) || strings.ToLower(value) == strings.ToLower(string(SpecialSlideType_LayoutSlide)) || strings.ToLower(value) == strings.ToLower(string(SpecialSlideType_NotesSlide))
}
