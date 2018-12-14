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
// SmartArtLayoutType : Represents layout type of a SmartArt diagram.
type SmartArtLayoutType string

// List of SmartArtLayoutType SmartArtLayoutType
const (
	SmartArtLayoutType_ACCENT_PROCESS SmartArtLayoutType = "AccentProcess"
	SmartArtLayoutType_ACCENTED_PICTURE SmartArtLayoutType = "AccentedPicture"
	SmartArtLayoutType_ALTERNATING_FLOW SmartArtLayoutType = "AlternatingFlow"
	SmartArtLayoutType_ALTERNATING_HEXAGONS SmartArtLayoutType = "AlternatingHexagons"
	SmartArtLayoutType_ALTERNATING_PICTURE_BLOCKS SmartArtLayoutType = "AlternatingPictureBlocks"
	SmartArtLayoutType_ALTERNATING_PICTURE_CIRCLES SmartArtLayoutType = "AlternatingPictureCircles"
	SmartArtLayoutType_ARROW_RIBBON SmartArtLayoutType = "ArrowRibbon"
	SmartArtLayoutType_ASCENDING_PICTURE_ACCENT_PROCESS SmartArtLayoutType = "AscendingPictureAccentProcess"
	SmartArtLayoutType_BALANCE SmartArtLayoutType = "Balance"
	SmartArtLayoutType_BASIC_BENDING_PROCESS SmartArtLayoutType = "BasicBendingProcess"
	SmartArtLayoutType_BASIC_BLOCK_LIST SmartArtLayoutType = "BasicBlockList"
	SmartArtLayoutType_BASIC_CHEVRON_PROCESS SmartArtLayoutType = "BasicChevronProcess"
	SmartArtLayoutType_BASIC_CYCLE SmartArtLayoutType = "BasicCycle"
	SmartArtLayoutType_BASIC_MATRIX SmartArtLayoutType = "BasicMatrix"
	SmartArtLayoutType_BASIC_PIE SmartArtLayoutType = "BasicPie"
	SmartArtLayoutType_BASIC_PROCESS SmartArtLayoutType = "BasicProcess"
	SmartArtLayoutType_BASIC_PYRAMID SmartArtLayoutType = "BasicPyramid"
	SmartArtLayoutType_BASIC_RADIAL SmartArtLayoutType = "BasicRadial"
	SmartArtLayoutType_BASIC_TARGET SmartArtLayoutType = "BasicTarget"
	SmartArtLayoutType_BASIC_TIMELINE SmartArtLayoutType = "BasicTimeline"
	SmartArtLayoutType_BASIC_VENN SmartArtLayoutType = "BasicVenn"
	SmartArtLayoutType_BENDING_PICTURE_ACCENT_LIST SmartArtLayoutType = "BendingPictureAccentList"
	SmartArtLayoutType_BENDING_PICTURE_BLOCKS SmartArtLayoutType = "BendingPictureBlocks"
	SmartArtLayoutType_BENDING_PICTURE_CAPTION SmartArtLayoutType = "BendingPictureCaption"
	SmartArtLayoutType_BENDING_PICTURE_CAPTION_LIST SmartArtLayoutType = "BendingPictureCaptionList"
	SmartArtLayoutType_BENDING_PICTURE_SEMI_TRANSPARENT_TEXT SmartArtLayoutType = "BendingPictureSemiTransparentText"
	SmartArtLayoutType_BLOCK_CYCLE SmartArtLayoutType = "BlockCycle"
	SmartArtLayoutType_BUBBLE_PICTURE_LIST SmartArtLayoutType = "BubblePictureList"
	SmartArtLayoutType_CAPTIONED_PICTURES SmartArtLayoutType = "CaptionedPictures"
	SmartArtLayoutType_CHEVRON_LIST SmartArtLayoutType = "ChevronList"
	SmartArtLayoutType_CIRCLE_ACCENT_TIMELINE SmartArtLayoutType = "CircleAccentTimeline"
	SmartArtLayoutType_CIRCLE_ARROW_PROCESS SmartArtLayoutType = "CircleArrowProcess"
	SmartArtLayoutType_CIRCLE_PICTURE_HIERARCHY SmartArtLayoutType = "CirclePictureHierarchy"
	SmartArtLayoutType_CIRCLE_RELATIONSHIP SmartArtLayoutType = "CircleRelationship"
	SmartArtLayoutType_CIRCULAR_BENDING_PROCESS SmartArtLayoutType = "CircularBendingProcess"
	SmartArtLayoutType_CIRCULAR_PICTURE_CALLOUT SmartArtLayoutType = "CircularPictureCallout"
	SmartArtLayoutType_CLOSED_CHEVRON_PROCESS SmartArtLayoutType = "ClosedChevronProcess"
	SmartArtLayoutType_CONTINUOUS_ARROW_PROCESS SmartArtLayoutType = "ContinuousArrowProcess"
	SmartArtLayoutType_CONTINUOUS_BLOCK_PROCESS SmartArtLayoutType = "ContinuousBlockProcess"
	SmartArtLayoutType_CONTINUOUS_CYCLE SmartArtLayoutType = "ContinuousCycle"
	SmartArtLayoutType_CONTINUOUS_PICTURE_LIST SmartArtLayoutType = "ContinuousPictureList"
	SmartArtLayoutType_CONVERGING_ARROWS SmartArtLayoutType = "ConvergingArrows"
	SmartArtLayoutType_CONVERGING_RADIAL SmartArtLayoutType = "ConvergingRadial"
	SmartArtLayoutType_COUNTERBALANCE_ARROWS SmartArtLayoutType = "CounterbalanceArrows"
	SmartArtLayoutType_CYCLE_MATRIX SmartArtLayoutType = "CycleMatrix"
	SmartArtLayoutType_DESCENDING_BLOCK_LIST SmartArtLayoutType = "DescendingBlockList"
	SmartArtLayoutType_DESCENDING_PROCESS SmartArtLayoutType = "DescendingProcess"
	SmartArtLayoutType_DETAILED_PROCESS SmartArtLayoutType = "DetailedProcess"
	SmartArtLayoutType_DIVERGING_ARROWS SmartArtLayoutType = "DivergingArrows"
	SmartArtLayoutType_DIVERGING_RADIAL SmartArtLayoutType = "DivergingRadial"
	SmartArtLayoutType_EQUATION SmartArtLayoutType = "Equation"
	SmartArtLayoutType_FRAMED_TEXT_PICTURE SmartArtLayoutType = "FramedTextPicture"
	SmartArtLayoutType_FUNNEL SmartArtLayoutType = "Funnel"
	SmartArtLayoutType_GEAR SmartArtLayoutType = "Gear"
	SmartArtLayoutType_GRID_MATRIX SmartArtLayoutType = "GridMatrix"
	SmartArtLayoutType_GROUPED_LIST SmartArtLayoutType = "GroupedList"
	SmartArtLayoutType_HALF_CIRCLE_ORGANIZATION_CHART SmartArtLayoutType = "HalfCircleOrganizationChart"
	SmartArtLayoutType_HEXAGON_CLUSTER SmartArtLayoutType = "HexagonCluster"
	SmartArtLayoutType_HIERARCHY SmartArtLayoutType = "Hierarchy"
	SmartArtLayoutType_HIERARCHY_LIST SmartArtLayoutType = "HierarchyList"
	SmartArtLayoutType_HORIZONTAL_BULLET_LIST SmartArtLayoutType = "HorizontalBulletList"
	SmartArtLayoutType_HORIZONTAL_HIERARCHY SmartArtLayoutType = "HorizontalHierarchy"
	SmartArtLayoutType_HORIZONTAL_LABELED_HIERARCHY SmartArtLayoutType = "HorizontalLabeledHierarchy"
	SmartArtLayoutType_HORIZONTAL_MULTI_LEVEL_HIERARCHY SmartArtLayoutType = "HorizontalMultiLevelHierarchy"
	SmartArtLayoutType_HORIZONTAL_ORGANIZATION_CHART SmartArtLayoutType = "HorizontalOrganizationChart"
	SmartArtLayoutType_HORIZONTAL_PICTURE_LIST SmartArtLayoutType = "HorizontalPictureList"
	SmartArtLayoutType_INCREASING_ARROWS_PROCESS SmartArtLayoutType = "IncreasingArrowsProcess"
	SmartArtLayoutType_INCREASING_CIRCLE_PROCESS SmartArtLayoutType = "IncreasingCircleProcess"
	SmartArtLayoutType_INVERTED_PYRAMID SmartArtLayoutType = "InvertedPyramid"
	SmartArtLayoutType_LABELED_HIERARCHY SmartArtLayoutType = "LabeledHierarchy"
	SmartArtLayoutType_LINEAR_VENN SmartArtLayoutType = "LinearVenn"
	SmartArtLayoutType_LINED_LIST SmartArtLayoutType = "LinedList"
	SmartArtLayoutType_MULTIDIRECTIONAL_CYCLE SmartArtLayoutType = "MultidirectionalCycle"
	SmartArtLayoutType_NAMEAND_TITLE_ORGANIZATION_CHART SmartArtLayoutType = "NameandTitleOrganizationChart"
	SmartArtLayoutType_NESTED_TARGET SmartArtLayoutType = "NestedTarget"
	SmartArtLayoutType_NONDIRECTIONAL_CYCLE SmartArtLayoutType = "NondirectionalCycle"
	SmartArtLayoutType_OPPOSING_ARROWS SmartArtLayoutType = "OpposingArrows"
	SmartArtLayoutType_OPPOSING_IDEAS SmartArtLayoutType = "OpposingIdeas"
	SmartArtLayoutType_ORGANIZATION_CHART SmartArtLayoutType = "OrganizationChart"
	SmartArtLayoutType_PHASED_PROCESS SmartArtLayoutType = "PhasedProcess"
	SmartArtLayoutType_PICTURE_ACCENT_BLOCKS SmartArtLayoutType = "PictureAccentBlocks"
	SmartArtLayoutType_PICTURE_ACCENT_LIST SmartArtLayoutType = "PictureAccentList"
	SmartArtLayoutType_PICTURE_ACCENT_PROCESS SmartArtLayoutType = "PictureAccentProcess"
	SmartArtLayoutType_PICTURE_CAPTION_LIST SmartArtLayoutType = "PictureCaptionList"
	SmartArtLayoutType_PICTURE_GRID SmartArtLayoutType = "PictureGrid"
	SmartArtLayoutType_PICTURE_LINEUP SmartArtLayoutType = "PictureLineup"
	SmartArtLayoutType_PICTURE_STRIPS SmartArtLayoutType = "PictureStrips"
	SmartArtLayoutType_PIE_PROCESS SmartArtLayoutType = "PieProcess"
	SmartArtLayoutType_PLUSAND_MINUS SmartArtLayoutType = "PlusandMinus"
	SmartArtLayoutType_PROCESS_ARROWS SmartArtLayoutType = "ProcessArrows"
	SmartArtLayoutType_PROCESS_LIST SmartArtLayoutType = "ProcessList"
	SmartArtLayoutType_PYRAMID_LIST SmartArtLayoutType = "PyramidList"
	SmartArtLayoutType_RADIAL_CLUSTER SmartArtLayoutType = "RadialCluster"
	SmartArtLayoutType_RADIAL_CYCLE SmartArtLayoutType = "RadialCycle"
	SmartArtLayoutType_RADIAL_LIST SmartArtLayoutType = "RadialList"
	SmartArtLayoutType_RADIAL_VENN SmartArtLayoutType = "RadialVenn"
	SmartArtLayoutType_RANDOM_TO_RESULT_PROCESS SmartArtLayoutType = "RandomToResultProcess"
	SmartArtLayoutType_REPEATING_BENDING_PROCESS SmartArtLayoutType = "RepeatingBendingProcess"
	SmartArtLayoutType_REVERSE_LIST SmartArtLayoutType = "ReverseList"
	SmartArtLayoutType_SEGMENTED_CYCLE SmartArtLayoutType = "SegmentedCycle"
	SmartArtLayoutType_SEGMENTED_PROCESS SmartArtLayoutType = "SegmentedProcess"
	SmartArtLayoutType_SEGMENTED_PYRAMID SmartArtLayoutType = "SegmentedPyramid"
	SmartArtLayoutType_SNAPSHOT_PICTURE_LIST SmartArtLayoutType = "SnapshotPictureList"
	SmartArtLayoutType_SPIRAL_PICTURE SmartArtLayoutType = "SpiralPicture"
	SmartArtLayoutType_SQUARE_ACCENT_LIST SmartArtLayoutType = "SquareAccentList"
	SmartArtLayoutType_STACKED_LIST SmartArtLayoutType = "StackedList"
	SmartArtLayoutType_STACKED_VENN SmartArtLayoutType = "StackedVenn"
	SmartArtLayoutType_STAGGERED_PROCESS SmartArtLayoutType = "StaggeredProcess"
	SmartArtLayoutType_STEP_DOWN_PROCESS SmartArtLayoutType = "StepDownProcess"
	SmartArtLayoutType_STEP_UP_PROCESS SmartArtLayoutType = "StepUpProcess"
	SmartArtLayoutType_SUB_STEP_PROCESS SmartArtLayoutType = "SubStepProcess"
	SmartArtLayoutType_TABLE_HIERARCHY SmartArtLayoutType = "TableHierarchy"
	SmartArtLayoutType_TABLE_LIST SmartArtLayoutType = "TableList"
	SmartArtLayoutType_TARGET_LIST SmartArtLayoutType = "TargetList"
	SmartArtLayoutType_TEXT_CYCLE SmartArtLayoutType = "TextCycle"
	SmartArtLayoutType_TITLE_PICTURE_LINEUP SmartArtLayoutType = "TitlePictureLineup"
	SmartArtLayoutType_TITLED_MATRIX SmartArtLayoutType = "TitledMatrix"
	SmartArtLayoutType_TITLED_PICTURE_ACCENT_LIST SmartArtLayoutType = "TitledPictureAccentList"
	SmartArtLayoutType_TITLED_PICTURE_BLOCKS SmartArtLayoutType = "TitledPictureBlocks"
	SmartArtLayoutType_TRAPEZOID_LIST SmartArtLayoutType = "TrapezoidList"
	SmartArtLayoutType_UPWARD_ARROW SmartArtLayoutType = "UpwardArrow"
	SmartArtLayoutType_VERTICAL_ACCENT_LIST SmartArtLayoutType = "VerticalAccentList"
	SmartArtLayoutType_VERTICAL_ARROW_LIST SmartArtLayoutType = "VerticalArrowList"
	SmartArtLayoutType_VERTICAL_BENDING_PROCESS SmartArtLayoutType = "VerticalBendingProcess"
	SmartArtLayoutType_VERTICAL_BLOCK_LIST SmartArtLayoutType = "VerticalBlockList"
	SmartArtLayoutType_VERTICAL_BOX_LIST SmartArtLayoutType = "VerticalBoxList"
	SmartArtLayoutType_VERTICAL_BULLET_LIST SmartArtLayoutType = "VerticalBulletList"
	SmartArtLayoutType_VERTICAL_CHEVRON_LIST SmartArtLayoutType = "VerticalChevronList"
	SmartArtLayoutType_VERTICAL_CIRCLE_LIST SmartArtLayoutType = "VerticalCircleList"
	SmartArtLayoutType_VERTICAL_CURVED_LIST SmartArtLayoutType = "VerticalCurvedList"
	SmartArtLayoutType_VERTICAL_EQUATION SmartArtLayoutType = "VerticalEquation"
	SmartArtLayoutType_VERTICAL_PICTURE_ACCENT_LIST SmartArtLayoutType = "VerticalPictureAccentList"
	SmartArtLayoutType_VERTICAL_PICTURE_LIST SmartArtLayoutType = "VerticalPictureList"
	SmartArtLayoutType_VERTICAL_PROCESS SmartArtLayoutType = "VerticalProcess"
	SmartArtLayoutType_CUSTOM SmartArtLayoutType = "Custom"
	SmartArtLayoutType_PICTURE_ORGANIZATION_CHART SmartArtLayoutType = "PictureOrganizationChart"
)