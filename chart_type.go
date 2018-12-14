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
// ChartType : Represents a type of chart.
type ChartType string

// List of ChartType ChartType
const (
	ChartType_CLUSTERED_COLUMN ChartType = "ClusteredColumn"
	ChartType_STACKED_COLUMN ChartType = "StackedColumn"
	ChartType_PERCENTS_STACKED_COLUMN ChartType = "PercentsStackedColumn"
	ChartType_CLUSTERED_COLUMN3_D ChartType = "ClusteredColumn3D"
	ChartType_STACKED_COLUMN3_D ChartType = "StackedColumn3D"
	ChartType_PERCENTS_STACKED_COLUMN3_D ChartType = "PercentsStackedColumn3D"
	ChartType_COLUMN3_D ChartType = "Column3D"
	ChartType_CLUSTERED_CYLINDER ChartType = "ClusteredCylinder"
	ChartType_STACKED_CYLINDER ChartType = "StackedCylinder"
	ChartType_PERCENTS_STACKED_CYLINDER ChartType = "PercentsStackedCylinder"
	ChartType_CYLINDER3_D ChartType = "Cylinder3D"
	ChartType_CLUSTERED_CONE ChartType = "ClusteredCone"
	ChartType_STACKED_CONE ChartType = "StackedCone"
	ChartType_PERCENTS_STACKED_CONE ChartType = "PercentsStackedCone"
	ChartType_CONE3_D ChartType = "Cone3D"
	ChartType_CLUSTERED_PYRAMID ChartType = "ClusteredPyramid"
	ChartType_STACKED_PYRAMID ChartType = "StackedPyramid"
	ChartType_PERCENTS_STACKED_PYRAMID ChartType = "PercentsStackedPyramid"
	ChartType_PYRAMID3_D ChartType = "Pyramid3D"
	ChartType_LINE ChartType = "Line"
	ChartType_STACKED_LINE ChartType = "StackedLine"
	ChartType_PERCENTS_STACKED_LINE ChartType = "PercentsStackedLine"
	ChartType_LINE_WITH_MARKERS ChartType = "LineWithMarkers"
	ChartType_STACKED_LINE_WITH_MARKERS ChartType = "StackedLineWithMarkers"
	ChartType_PERCENTS_STACKED_LINE_WITH_MARKERS ChartType = "PercentsStackedLineWithMarkers"
	ChartType_LINE3_D ChartType = "Line3D"
	ChartType_PIE ChartType = "Pie"
	ChartType_PIE3_D ChartType = "Pie3D"
	ChartType_PIE_OF_PIE ChartType = "PieOfPie"
	ChartType_EXPLODED_PIE ChartType = "ExplodedPie"
	ChartType_EXPLODED_PIE3_D ChartType = "ExplodedPie3D"
	ChartType_BAR_OF_PIE ChartType = "BarOfPie"
	ChartType_PERCENTS_STACKED_BAR ChartType = "PercentsStackedBar"
	ChartType_CLUSTERED_BAR3_D ChartType = "ClusteredBar3D"
	ChartType_CLUSTERED_BAR ChartType = "ClusteredBar"
	ChartType_STACKED_BAR ChartType = "StackedBar"
	ChartType_STACKED_BAR3_D ChartType = "StackedBar3D"
	ChartType_PERCENTS_STACKED_BAR3_D ChartType = "PercentsStackedBar3D"
	ChartType_CLUSTERED_HORIZONTAL_CYLINDER ChartType = "ClusteredHorizontalCylinder"
	ChartType_STACKED_HORIZONTAL_CYLINDER ChartType = "StackedHorizontalCylinder"
	ChartType_PERCENTS_STACKED_HORIZONTAL_CYLINDER ChartType = "PercentsStackedHorizontalCylinder"
	ChartType_CLUSTERED_HORIZONTAL_CONE ChartType = "ClusteredHorizontalCone"
	ChartType_STACKED_HORIZONTAL_CONE ChartType = "StackedHorizontalCone"
	ChartType_PERCENTS_STACKED_HORIZONTAL_CONE ChartType = "PercentsStackedHorizontalCone"
	ChartType_CLUSTERED_HORIZONTAL_PYRAMID ChartType = "ClusteredHorizontalPyramid"
	ChartType_STACKED_HORIZONTAL_PYRAMID ChartType = "StackedHorizontalPyramid"
	ChartType_PERCENTS_STACKED_HORIZONTAL_PYRAMID ChartType = "PercentsStackedHorizontalPyramid"
	ChartType_AREA ChartType = "Area"
	ChartType_STACKED_AREA ChartType = "StackedArea"
	ChartType_PERCENTS_STACKED_AREA ChartType = "PercentsStackedArea"
	ChartType_AREA3_D ChartType = "Area3D"
	ChartType_STACKED_AREA3_D ChartType = "StackedArea3D"
	ChartType_PERCENTS_STACKED_AREA3_D ChartType = "PercentsStackedArea3D"
	ChartType_SCATTER_WITH_MARKERS ChartType = "ScatterWithMarkers"
	ChartType_SCATTER_WITH_SMOOTH_LINES_AND_MARKERS ChartType = "ScatterWithSmoothLinesAndMarkers"
	ChartType_SCATTER_WITH_SMOOTH_LINES ChartType = "ScatterWithSmoothLines"
	ChartType_SCATTER_WITH_STRAIGHT_LINES_AND_MARKERS ChartType = "ScatterWithStraightLinesAndMarkers"
	ChartType_SCATTER_WITH_STRAIGHT_LINES ChartType = "ScatterWithStraightLines"
	ChartType_HIGH_LOW_CLOSE ChartType = "HighLowClose"
	ChartType_OPEN_HIGH_LOW_CLOSE ChartType = "OpenHighLowClose"
	ChartType_VOLUME_HIGH_LOW_CLOSE ChartType = "VolumeHighLowClose"
	ChartType_VOLUME_OPEN_HIGH_LOW_CLOSE ChartType = "VolumeOpenHighLowClose"
	ChartType_SURFACE3_D ChartType = "Surface3D"
	ChartType_WIREFRAME_SURFACE3_D ChartType = "WireframeSurface3D"
	ChartType_CONTOUR ChartType = "Contour"
	ChartType_WIREFRAME_CONTOUR ChartType = "WireframeContour"
	ChartType_DOUGHNUT ChartType = "Doughnut"
	ChartType_EXPLODED_DOUGHNUT ChartType = "ExplodedDoughnut"
	ChartType_BUBBLE ChartType = "Bubble"
	ChartType_BUBBLE_WITH3_D ChartType = "BubbleWith3D"
	ChartType_RADAR ChartType = "Radar"
	ChartType_RADAR_WITH_MARKERS ChartType = "RadarWithMarkers"
	ChartType_FILLED_RADAR ChartType = "FilledRadar"
	ChartType_SERIES_OF_MIXED_TYPES ChartType = "SeriesOfMixedTypes"
)