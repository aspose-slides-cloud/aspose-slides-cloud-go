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

// ShapeType : Shape type
type ShapeType string

// List of ShapeType ShapeType
const (
	ShapeType_Shape ShapeType = "Shape"
	ShapeType_Chart ShapeType = "Chart"
	ShapeType_Table ShapeType = "Table"
	ShapeType_PictureFrame ShapeType = "PictureFrame"
	ShapeType_VideoFrame ShapeType = "VideoFrame"
	ShapeType_AudioFrame ShapeType = "AudioFrame"
	ShapeType_SmartArt ShapeType = "SmartArt"
	ShapeType_OleObjectFrame ShapeType = "OleObjectFrame"
	ShapeType_GroupShape ShapeType = "GroupShape"
	ShapeType_GraphicalObject ShapeType = "GraphicalObject"
	ShapeType_Connector ShapeType = "Connector"
	ShapeType_SmartArtShape ShapeType = "SmartArtShape"
	ShapeType_ZoomFrame ShapeType = "ZoomFrame"
	ShapeType_SectionZoomFrame ShapeType = "SectionZoomFrame"
	ShapeType_SummaryZoomFrame ShapeType = "SummaryZoomFrame"
	ShapeType_SummaryZoomSection ShapeType = "SummaryZoomSection"
)

func ShapeType_Validate(value string) (bool) {
	return strings.ToLower(value) == strings.ToLower(string(ShapeType_Shape)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_Chart)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_Table)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_PictureFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_VideoFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_AudioFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_SmartArt)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_OleObjectFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_GroupShape)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_GraphicalObject)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_Connector)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_SmartArtShape)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_ZoomFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_SectionZoomFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_SummaryZoomFrame)) || strings.ToLower(value) == strings.ToLower(string(ShapeType_SummaryZoomSection))
}
