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
// TableStylePreset : 
type TableStylePreset string

// List of TableStylePreset TableStylePreset
const (
	TableStylePreset_NONE TableStylePreset = "None"
	TableStylePreset_MEDIUM_STYLE2_ACCENT1 TableStylePreset = "MediumStyle2Accent1"
	TableStylePreset_MEDIUM_STYLE2 TableStylePreset = "MediumStyle2"
	TableStylePreset_NO_STYLE_NO_GRID TableStylePreset = "NoStyleNoGrid"
	TableStylePreset_THEMED_STYLE1_ACCENT1 TableStylePreset = "ThemedStyle1Accent1"
	TableStylePreset_THEMED_STYLE1_ACCENT2 TableStylePreset = "ThemedStyle1Accent2"
	TableStylePreset_THEMED_STYLE1_ACCENT3 TableStylePreset = "ThemedStyle1Accent3"
	TableStylePreset_THEMED_STYLE1_ACCENT4 TableStylePreset = "ThemedStyle1Accent4"
	TableStylePreset_THEMED_STYLE1_ACCENT5 TableStylePreset = "ThemedStyle1Accent5"
	TableStylePreset_THEMED_STYLE1_ACCENT6 TableStylePreset = "ThemedStyle1Accent6"
	TableStylePreset_NO_STYLE_TABLE_GRID TableStylePreset = "NoStyleTableGrid"
	TableStylePreset_THEMED_STYLE2_ACCENT1 TableStylePreset = "ThemedStyle2Accent1"
	TableStylePreset_THEMED_STYLE2_ACCENT2 TableStylePreset = "ThemedStyle2Accent2"
	TableStylePreset_THEMED_STYLE2_ACCENT3 TableStylePreset = "ThemedStyle2Accent3"
	TableStylePreset_THEMED_STYLE2_ACCENT4 TableStylePreset = "ThemedStyle2Accent4"
	TableStylePreset_THEMED_STYLE2_ACCENT5 TableStylePreset = "ThemedStyle2Accent5"
	TableStylePreset_THEMED_STYLE2_ACCENT6 TableStylePreset = "ThemedStyle2Accent6"
	TableStylePreset_LIGHT_STYLE1 TableStylePreset = "LightStyle1"
	TableStylePreset_LIGHT_STYLE1_ACCENT1 TableStylePreset = "LightStyle1Accent1"
	TableStylePreset_LIGHT_STYLE1_ACCENT2 TableStylePreset = "LightStyle1Accent2"
	TableStylePreset_LIGHT_STYLE1_ACCENT3 TableStylePreset = "LightStyle1Accent3"
	TableStylePreset_LIGHT_STYLE1_ACCENT4 TableStylePreset = "LightStyle1Accent4"
	TableStylePreset_LIGHT_STYLE2_ACCENT5 TableStylePreset = "LightStyle2Accent5"
	TableStylePreset_LIGHT_STYLE1_ACCENT6 TableStylePreset = "LightStyle1Accent6"
	TableStylePreset_LIGHT_STYLE2 TableStylePreset = "LightStyle2"
	TableStylePreset_LIGHT_STYLE2_ACCENT1 TableStylePreset = "LightStyle2Accent1"
	TableStylePreset_LIGHT_STYLE2_ACCENT2 TableStylePreset = "LightStyle2Accent2"
	TableStylePreset_LIGHT_STYLE2_ACCENT3 TableStylePreset = "LightStyle2Accent3"
	TableStylePreset_MEDIUM_STYLE2_ACCENT3 TableStylePreset = "MediumStyle2Accent3"
	TableStylePreset_MEDIUM_STYLE2_ACCENT4 TableStylePreset = "MediumStyle2Accent4"
	TableStylePreset_MEDIUM_STYLE2_ACCENT5 TableStylePreset = "MediumStyle2Accent5"
	TableStylePreset_LIGHT_STYLE2_ACCENT6 TableStylePreset = "LightStyle2Accent6"
	TableStylePreset_LIGHT_STYLE2_ACCENT4 TableStylePreset = "LightStyle2Accent4"
	TableStylePreset_LIGHT_STYLE3 TableStylePreset = "LightStyle3"
	TableStylePreset_LIGHT_STYLE3_ACCENT1 TableStylePreset = "LightStyle3Accent1"
	TableStylePreset_MEDIUM_STYLE2_ACCENT2 TableStylePreset = "MediumStyle2Accent2"
	TableStylePreset_LIGHT_STYLE3_ACCENT2 TableStylePreset = "LightStyle3Accent2"
	TableStylePreset_LIGHT_STYLE3_ACCENT3 TableStylePreset = "LightStyle3Accent3"
	TableStylePreset_LIGHT_STYLE3_ACCENT4 TableStylePreset = "LightStyle3Accent4"
	TableStylePreset_LIGHT_STYLE3_ACCENT5 TableStylePreset = "LightStyle3Accent5"
	TableStylePreset_LIGHT_STYLE3_ACCENT6 TableStylePreset = "LightStyle3Accent6"
	TableStylePreset_MEDIUM_STYLE1 TableStylePreset = "MediumStyle1"
	TableStylePreset_MEDIUM_STYLE1_ACCENT1 TableStylePreset = "MediumStyle1Accent1"
	TableStylePreset_MEDIUM_STYLE1_ACCENT2 TableStylePreset = "MediumStyle1Accent2"
	TableStylePreset_MEDIUM_STYLE1_ACCENT3 TableStylePreset = "MediumStyle1Accent3"
	TableStylePreset_MEDIUM_STYLE1_ACCENT4 TableStylePreset = "MediumStyle1Accent4"
	TableStylePreset_MEDIUM_STYLE1_ACCENT5 TableStylePreset = "MediumStyle1Accent5"
	TableStylePreset_MEDIUM_STYLE1_ACCENT6 TableStylePreset = "MediumStyle1Accent6"
	TableStylePreset_MEDIUM_STYLE2_ACCENT6 TableStylePreset = "MediumStyle2Accent6"
	TableStylePreset_MEDIUM_STYLE3 TableStylePreset = "MediumStyle3"
	TableStylePreset_MEDIUM_STYLE3_ACCENT1 TableStylePreset = "MediumStyle3Accent1"
	TableStylePreset_MEDIUM_STYLE3_ACCENT2 TableStylePreset = "MediumStyle3Accent2"
	TableStylePreset_MEDIUM_STYLE3_ACCENT3 TableStylePreset = "MediumStyle3Accent3"
	TableStylePreset_MEDIUM_STYLE3_ACCENT4 TableStylePreset = "MediumStyle3Accent4"
	TableStylePreset_MEDIUM_STYLE3_ACCENT5 TableStylePreset = "MediumStyle3Accent5"
	TableStylePreset_MEDIUM_STYLE3_ACCENT6 TableStylePreset = "MediumStyle3Accent6"
	TableStylePreset_MEDIUM_STYLE4 TableStylePreset = "MediumStyle4"
	TableStylePreset_MEDIUM_STYLE4_ACCENT1 TableStylePreset = "MediumStyle4Accent1"
	TableStylePreset_MEDIUM_STYLE4_ACCENT2 TableStylePreset = "MediumStyle4Accent2"
	TableStylePreset_MEDIUM_STYLE4_ACCENT3 TableStylePreset = "MediumStyle4Accent3"
	TableStylePreset_MEDIUM_STYLE4_ACCENT4 TableStylePreset = "MediumStyle4Accent4"
	TableStylePreset_MEDIUM_STYLE4_ACCENT5 TableStylePreset = "MediumStyle4Accent5"
	TableStylePreset_MEDIUM_STYLE4_ACCENT6 TableStylePreset = "MediumStyle4Accent6"
	TableStylePreset_DARK_STYLE1 TableStylePreset = "DarkStyle1"
	TableStylePreset_DARK_STYLE1_ACCENT1 TableStylePreset = "DarkStyle1Accent1"
	TableStylePreset_DARK_STYLE1_ACCENT2 TableStylePreset = "DarkStyle1Accent2"
	TableStylePreset_DARK_STYLE1_ACCENT3 TableStylePreset = "DarkStyle1Accent3"
	TableStylePreset_DARK_STYLE1_ACCENT4 TableStylePreset = "DarkStyle1Accent4"
	TableStylePreset_DARK_STYLE1_ACCENT5 TableStylePreset = "DarkStyle1Accent5"
	TableStylePreset_DARK_STYLE1_ACCENT6 TableStylePreset = "DarkStyle1Accent6"
	TableStylePreset_DARK_STYLE2 TableStylePreset = "DarkStyle2"
	TableStylePreset_DARK_STYLE2_ACCENT1_ACCENT2 TableStylePreset = "DarkStyle2Accent1Accent2"
	TableStylePreset_DARK_STYLE2_ACCENT3_ACCENT4 TableStylePreset = "DarkStyle2Accent3Accent4"
	TableStylePreset_DARK_STYLE2_ACCENT5_ACCENT6 TableStylePreset = "DarkStyle2Accent5Accent6"
	TableStylePreset_LIGHT_STYLE1_ACCENT5 TableStylePreset = "LightStyle1Accent5"
	TableStylePreset_CUSTOM TableStylePreset = "Custom"
)