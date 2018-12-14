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
// GeometryShapeType : 
type GeometryShapeType string

// List of GeometryShapeType GeometryShapeType
const (
	GeometryShapeType_CUSTOM GeometryShapeType = "Custom"
	GeometryShapeType_LINE GeometryShapeType = "Line"
	GeometryShapeType_LINE_INVERSE GeometryShapeType = "LineInverse"
	GeometryShapeType_TRIANGLE GeometryShapeType = "Triangle"
	GeometryShapeType_RIGHT_TRIANGLE GeometryShapeType = "RightTriangle"
	GeometryShapeType_RECTANGLE GeometryShapeType = "Rectangle"
	GeometryShapeType_DIAMOND GeometryShapeType = "Diamond"
	GeometryShapeType_PARALLELOGRAM GeometryShapeType = "Parallelogram"
	GeometryShapeType_TRAPEZOID GeometryShapeType = "Trapezoid"
	GeometryShapeType_NON_ISOSCELES_TRAPEZOID GeometryShapeType = "NonIsoscelesTrapezoid"
	GeometryShapeType_PENTAGON GeometryShapeType = "Pentagon"
	GeometryShapeType_HEXAGON GeometryShapeType = "Hexagon"
	GeometryShapeType_HEPTAGON GeometryShapeType = "Heptagon"
	GeometryShapeType_OCTAGON GeometryShapeType = "Octagon"
	GeometryShapeType_DECAGON GeometryShapeType = "Decagon"
	GeometryShapeType_DODECAGON GeometryShapeType = "Dodecagon"
	GeometryShapeType_FOUR_POINTED_STAR GeometryShapeType = "FourPointedStar"
	GeometryShapeType_FIVE_POINTED_STAR GeometryShapeType = "FivePointedStar"
	GeometryShapeType_SIX_POINTED_STAR GeometryShapeType = "SixPointedStar"
	GeometryShapeType_SEVEN_POINTED_STAR GeometryShapeType = "SevenPointedStar"
	GeometryShapeType_EIGHT_POINTED_STAR GeometryShapeType = "EightPointedStar"
	GeometryShapeType_TEN_POINTED_STAR GeometryShapeType = "TenPointedStar"
	GeometryShapeType_TWELVE_POINTED_STAR GeometryShapeType = "TwelvePointedStar"
	GeometryShapeType_SIXTEEN_POINTED_STAR GeometryShapeType = "SixteenPointedStar"
	GeometryShapeType_TWENTY_FOUR_POINTED_STAR GeometryShapeType = "TwentyFourPointedStar"
	GeometryShapeType_THIRTY_TWO_POINTED_STAR GeometryShapeType = "ThirtyTwoPointedStar"
	GeometryShapeType_ROUND_CORNER_RECTANGLE GeometryShapeType = "RoundCornerRectangle"
	GeometryShapeType_ONE_ROUND_CORNER_RECTANGLE GeometryShapeType = "OneRoundCornerRectangle"
	GeometryShapeType_TWO_SAMESIDE_ROUND_CORNER_RECTANGLE GeometryShapeType = "TwoSamesideRoundCornerRectangle"
	GeometryShapeType_TWO_DIAGONAL_ROUND_CORNER_RECTANGLE GeometryShapeType = "TwoDiagonalRoundCornerRectangle"
	GeometryShapeType_ONE_SNIP_ONE_ROUND_CORNER_RECTANGLE GeometryShapeType = "OneSnipOneRoundCornerRectangle"
	GeometryShapeType_ONE_SNIP_CORNER_RECTANGLE GeometryShapeType = "OneSnipCornerRectangle"
	GeometryShapeType_TWO_SAMESIDE_SNIP_CORNER_RECTANGLE GeometryShapeType = "TwoSamesideSnipCornerRectangle"
	GeometryShapeType_TWO_DIAGONAL_SNIP_CORNER_RECTANGLE GeometryShapeType = "TwoDiagonalSnipCornerRectangle"
	GeometryShapeType_PLAQUE GeometryShapeType = "Plaque"
	GeometryShapeType_ELLIPSE GeometryShapeType = "Ellipse"
	GeometryShapeType_TEARDROP GeometryShapeType = "Teardrop"
	GeometryShapeType_HOME_PLATE GeometryShapeType = "HomePlate"
	GeometryShapeType_CHEVRON GeometryShapeType = "Chevron"
	GeometryShapeType_PIE_WEDGE GeometryShapeType = "PieWedge"
	GeometryShapeType_PIE GeometryShapeType = "Pie"
	GeometryShapeType_BLOCK_ARC GeometryShapeType = "BlockArc"
	GeometryShapeType_DONUT GeometryShapeType = "Donut"
	GeometryShapeType_NO_SMOKING GeometryShapeType = "NoSmoking"
	GeometryShapeType_RIGHT_ARROW GeometryShapeType = "RightArrow"
	GeometryShapeType_LEFT_ARROW GeometryShapeType = "LeftArrow"
	GeometryShapeType_UP_ARROW GeometryShapeType = "UpArrow"
	GeometryShapeType_DOWN_ARROW GeometryShapeType = "DownArrow"
	GeometryShapeType_STRIPED_RIGHT_ARROW GeometryShapeType = "StripedRightArrow"
	GeometryShapeType_NOTCHED_RIGHT_ARROW GeometryShapeType = "NotchedRightArrow"
	GeometryShapeType_BENT_UP_ARROW GeometryShapeType = "BentUpArrow"
	GeometryShapeType_LEFT_RIGHT_ARROW GeometryShapeType = "LeftRightArrow"
	GeometryShapeType_UP_DOWN_ARROW GeometryShapeType = "UpDownArrow"
	GeometryShapeType_LEFT_UP_ARROW GeometryShapeType = "LeftUpArrow"
	GeometryShapeType_LEFT_RIGHT_UP_ARROW GeometryShapeType = "LeftRightUpArrow"
	GeometryShapeType_QUAD_ARROW GeometryShapeType = "QuadArrow"
	GeometryShapeType_CALLOUT_LEFT_ARROW GeometryShapeType = "CalloutLeftArrow"
	GeometryShapeType_CALLOUT_RIGHT_ARROW GeometryShapeType = "CalloutRightArrow"
	GeometryShapeType_CALLOUT_UP_ARROW GeometryShapeType = "CalloutUpArrow"
	GeometryShapeType_CALLOUT_DOWN_ARROW GeometryShapeType = "CalloutDownArrow"
	GeometryShapeType_CALLOUT_LEFT_RIGHT_ARROW GeometryShapeType = "CalloutLeftRightArrow"
	GeometryShapeType_CALLOUT_UP_DOWN_ARROW GeometryShapeType = "CalloutUpDownArrow"
	GeometryShapeType_CALLOUT_QUAD_ARROW GeometryShapeType = "CalloutQuadArrow"
	GeometryShapeType_BENT_ARROW GeometryShapeType = "BentArrow"
	GeometryShapeType_U_TURN_ARROW GeometryShapeType = "UTurnArrow"
	GeometryShapeType_CIRCULAR_ARROW GeometryShapeType = "CircularArrow"
	GeometryShapeType_LEFT_CIRCULAR_ARROW GeometryShapeType = "LeftCircularArrow"
	GeometryShapeType_LEFT_RIGHT_CIRCULAR_ARROW GeometryShapeType = "LeftRightCircularArrow"
	GeometryShapeType_CURVED_RIGHT_ARROW GeometryShapeType = "CurvedRightArrow"
	GeometryShapeType_CURVED_LEFT_ARROW GeometryShapeType = "CurvedLeftArrow"
	GeometryShapeType_CURVED_UP_ARROW GeometryShapeType = "CurvedUpArrow"
	GeometryShapeType_CURVED_DOWN_ARROW GeometryShapeType = "CurvedDownArrow"
	GeometryShapeType_SWOOSH_ARROW GeometryShapeType = "SwooshArrow"
	GeometryShapeType_CUBE GeometryShapeType = "Cube"
	GeometryShapeType_CAN GeometryShapeType = "Can"
	GeometryShapeType_LIGHTNING_BOLT GeometryShapeType = "LightningBolt"
	GeometryShapeType_HEART GeometryShapeType = "Heart"
	GeometryShapeType_SUN GeometryShapeType = "Sun"
	GeometryShapeType_MOON GeometryShapeType = "Moon"
	GeometryShapeType_SMILEY_FACE GeometryShapeType = "SmileyFace"
	GeometryShapeType_IRREGULAR_SEAL1 GeometryShapeType = "IrregularSeal1"
	GeometryShapeType_IRREGULAR_SEAL2 GeometryShapeType = "IrregularSeal2"
	GeometryShapeType_FOLDED_CORNER GeometryShapeType = "FoldedCorner"
	GeometryShapeType_BEVEL GeometryShapeType = "Bevel"
	GeometryShapeType_FRAME GeometryShapeType = "Frame"
	GeometryShapeType_HALF_FRAME GeometryShapeType = "HalfFrame"
	GeometryShapeType_CORNER GeometryShapeType = "Corner"
	GeometryShapeType_DIAGONAL_STRIPE GeometryShapeType = "DiagonalStripe"
	GeometryShapeType_CHORD GeometryShapeType = "Chord"
	GeometryShapeType_CURVED_ARC GeometryShapeType = "CurvedArc"
	GeometryShapeType_LEFT_BRACKET GeometryShapeType = "LeftBracket"
	GeometryShapeType_RIGHT_BRACKET GeometryShapeType = "RightBracket"
	GeometryShapeType_LEFT_BRACE GeometryShapeType = "LeftBrace"
	GeometryShapeType_RIGHT_BRACE GeometryShapeType = "RightBrace"
	GeometryShapeType_BRACKET_PAIR GeometryShapeType = "BracketPair"
	GeometryShapeType_BRACE_PAIR GeometryShapeType = "BracePair"
	GeometryShapeType_STRAIGHT_CONNECTOR1 GeometryShapeType = "StraightConnector1"
	GeometryShapeType_BENT_CONNECTOR2 GeometryShapeType = "BentConnector2"
	GeometryShapeType_BENT_CONNECTOR3 GeometryShapeType = "BentConnector3"
	GeometryShapeType_BENT_CONNECTOR4 GeometryShapeType = "BentConnector4"
	GeometryShapeType_BENT_CONNECTOR5 GeometryShapeType = "BentConnector5"
	GeometryShapeType_CURVED_CONNECTOR2 GeometryShapeType = "CurvedConnector2"
	GeometryShapeType_CURVED_CONNECTOR3 GeometryShapeType = "CurvedConnector3"
	GeometryShapeType_CURVED_CONNECTOR4 GeometryShapeType = "CurvedConnector4"
	GeometryShapeType_CURVED_CONNECTOR5 GeometryShapeType = "CurvedConnector5"
	GeometryShapeType_CALLOUT1 GeometryShapeType = "Callout1"
	GeometryShapeType_CALLOUT2 GeometryShapeType = "Callout2"
	GeometryShapeType_CALLOUT3 GeometryShapeType = "Callout3"
	GeometryShapeType_CALLOUT1_WITH_ACCENT GeometryShapeType = "Callout1WithAccent"
	GeometryShapeType_CALLOUT2_WITH_ACCENT GeometryShapeType = "Callout2WithAccent"
	GeometryShapeType_CALLOUT3_WITH_ACCENT GeometryShapeType = "Callout3WithAccent"
	GeometryShapeType_CALLOUT1_WITH_BORDER GeometryShapeType = "Callout1WithBorder"
	GeometryShapeType_CALLOUT2_WITH_BORDER GeometryShapeType = "Callout2WithBorder"
	GeometryShapeType_CALLOUT3_WITH_BORDER GeometryShapeType = "Callout3WithBorder"
	GeometryShapeType_CALLOUT1_WITH_BORDER_AND_ACCENT GeometryShapeType = "Callout1WithBorderAndAccent"
	GeometryShapeType_CALLOUT2_WITH_BORDER_AND_ACCENT GeometryShapeType = "Callout2WithBorderAndAccent"
	GeometryShapeType_CALLOUT3_WITH_BORDER_AND_ACCENT GeometryShapeType = "Callout3WithBorderAndAccent"
	GeometryShapeType_CALLOUT_WEDGE_RECTANGLE GeometryShapeType = "CalloutWedgeRectangle"
	GeometryShapeType_CALLOUT_WEDGE_ROUND_RECTANGLE GeometryShapeType = "CalloutWedgeRoundRectangle"
	GeometryShapeType_CALLOUT_WEDGE_ELLIPSE GeometryShapeType = "CalloutWedgeEllipse"
	GeometryShapeType_CALLOUT_CLOUD GeometryShapeType = "CalloutCloud"
	GeometryShapeType_CLOUD GeometryShapeType = "Cloud"
	GeometryShapeType_RIBBON GeometryShapeType = "Ribbon"
	GeometryShapeType_RIBBON2 GeometryShapeType = "Ribbon2"
	GeometryShapeType_ELLIPSE_RIBBON GeometryShapeType = "EllipseRibbon"
	GeometryShapeType_ELLIPSE_RIBBON2 GeometryShapeType = "EllipseRibbon2"
	GeometryShapeType_LEFT_RIGHT_RIBBON GeometryShapeType = "LeftRightRibbon"
	GeometryShapeType_VERTICAL_SCROLL GeometryShapeType = "VerticalScroll"
	GeometryShapeType_HORIZONTAL_SCROLL GeometryShapeType = "HorizontalScroll"
	GeometryShapeType_WAVE GeometryShapeType = "Wave"
	GeometryShapeType_DOUBLE_WAVE GeometryShapeType = "DoubleWave"
	GeometryShapeType_PLUS GeometryShapeType = "Plus"
	GeometryShapeType_PROCESS_FLOW GeometryShapeType = "ProcessFlow"
	GeometryShapeType_DECISION_FLOW GeometryShapeType = "DecisionFlow"
	GeometryShapeType_INPUT_OUTPUT_FLOW GeometryShapeType = "InputOutputFlow"
	GeometryShapeType_PREDEFINED_PROCESS_FLOW GeometryShapeType = "PredefinedProcessFlow"
	GeometryShapeType_INTERNAL_STORAGE_FLOW GeometryShapeType = "InternalStorageFlow"
	GeometryShapeType_DOCUMENT_FLOW GeometryShapeType = "DocumentFlow"
	GeometryShapeType_MULTI_DOCUMENT_FLOW GeometryShapeType = "MultiDocumentFlow"
	GeometryShapeType_TERMINATOR_FLOW GeometryShapeType = "TerminatorFlow"
	GeometryShapeType_PREPARATION_FLOW GeometryShapeType = "PreparationFlow"
	GeometryShapeType_MANUAL_INPUT_FLOW GeometryShapeType = "ManualInputFlow"
	GeometryShapeType_MANUAL_OPERATION_FLOW GeometryShapeType = "ManualOperationFlow"
	GeometryShapeType_CONNECTOR_FLOW GeometryShapeType = "ConnectorFlow"
	GeometryShapeType_PUNCHED_CARD_FLOW GeometryShapeType = "PunchedCardFlow"
	GeometryShapeType_PUNCHED_TAPE_FLOW GeometryShapeType = "PunchedTapeFlow"
	GeometryShapeType_SUMMING_JUNCTION_FLOW GeometryShapeType = "SummingJunctionFlow"
	GeometryShapeType_OR_FLOW GeometryShapeType = "OrFlow"
	GeometryShapeType_COLLATE_FLOW GeometryShapeType = "CollateFlow"
	GeometryShapeType_SORT_FLOW GeometryShapeType = "SortFlow"
	GeometryShapeType_EXTRACT_FLOW GeometryShapeType = "ExtractFlow"
	GeometryShapeType_MERGE_FLOW GeometryShapeType = "MergeFlow"
	GeometryShapeType_OFFLINE_STORAGE_FLOW GeometryShapeType = "OfflineStorageFlow"
	GeometryShapeType_ONLINE_STORAGE_FLOW GeometryShapeType = "OnlineStorageFlow"
	GeometryShapeType_MAGNETIC_TAPE_FLOW GeometryShapeType = "MagneticTapeFlow"
	GeometryShapeType_MAGNETIC_DISK_FLOW GeometryShapeType = "MagneticDiskFlow"
	GeometryShapeType_MAGNETIC_DRUM_FLOW GeometryShapeType = "MagneticDrumFlow"
	GeometryShapeType_DISPLAY_FLOW GeometryShapeType = "DisplayFlow"
	GeometryShapeType_DELAY_FLOW GeometryShapeType = "DelayFlow"
	GeometryShapeType_ALTERNATE_PROCESS_FLOW GeometryShapeType = "AlternateProcessFlow"
	GeometryShapeType_OFF_PAGE_CONNECTOR_FLOW GeometryShapeType = "OffPageConnectorFlow"
	GeometryShapeType_BLANK_BUTTON GeometryShapeType = "BlankButton"
	GeometryShapeType_HOME_BUTTON GeometryShapeType = "HomeButton"
	GeometryShapeType_HELP_BUTTON GeometryShapeType = "HelpButton"
	GeometryShapeType_INFORMATION_BUTTON GeometryShapeType = "InformationButton"
	GeometryShapeType_FORWARD_OR_NEXT_BUTTON GeometryShapeType = "ForwardOrNextButton"
	GeometryShapeType_BACK_OR_PREVIOUS_BUTTON GeometryShapeType = "BackOrPreviousButton"
	GeometryShapeType_END_BUTTON GeometryShapeType = "EndButton"
	GeometryShapeType_BEGINNING_BUTTON GeometryShapeType = "BeginningButton"
	GeometryShapeType_RETURN_BUTTON GeometryShapeType = "ReturnButton"
	GeometryShapeType_DOCUMENT_BUTTON GeometryShapeType = "DocumentButton"
	GeometryShapeType_SOUND_BUTTON GeometryShapeType = "SoundButton"
	GeometryShapeType_MOVIE_BUTTON GeometryShapeType = "MovieButton"
	GeometryShapeType_GEAR6 GeometryShapeType = "Gear6"
	GeometryShapeType_GEAR9 GeometryShapeType = "Gear9"
	GeometryShapeType_FUNNEL GeometryShapeType = "Funnel"
	GeometryShapeType_PLUS_MATH GeometryShapeType = "PlusMath"
	GeometryShapeType_MINUS_MATH GeometryShapeType = "MinusMath"
	GeometryShapeType_MULTIPLY_MATH GeometryShapeType = "MultiplyMath"
	GeometryShapeType_DIVIDE_MATH GeometryShapeType = "DivideMath"
	GeometryShapeType_EQUAL_MATH GeometryShapeType = "EqualMath"
	GeometryShapeType_NOT_EQUAL_MATH GeometryShapeType = "NotEqualMath"
	GeometryShapeType_CORNER_TABS GeometryShapeType = "CornerTabs"
	GeometryShapeType_SQUARE_TABS GeometryShapeType = "SquareTabs"
	GeometryShapeType_PLAQUE_TABS GeometryShapeType = "PlaqueTabs"
	GeometryShapeType_CHART_X GeometryShapeType = "ChartX"
	GeometryShapeType_CHART_STAR GeometryShapeType = "ChartStar"
	GeometryShapeType_CHART_PLUS GeometryShapeType = "ChartPlus"
	GeometryShapeType_NOT_DEFINED GeometryShapeType = "NotDefined"
)