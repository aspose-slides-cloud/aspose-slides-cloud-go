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

// Contains the TextFrame's formatting properties.
type ITextFrameFormat interface {

	// Represents 3d effect properties for a text.
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets text wrapping shape.
	GetTransform() string
	SetTransform(newValue string)

	// Left margin. Left margin.
	GetMarginLeft() float64
	SetMarginLeft(newValue float64)

	// Right margin.
	GetMarginRight() float64
	SetMarginRight(newValue float64)

	// Top margin.
	GetMarginTop() float64
	SetMarginTop(newValue float64)

	// Bottom margin.
	GetMarginBottom() float64
	SetMarginBottom(newValue float64)

	// True if text is wrapped at TextFrame's margins.
	GetWrapText() string
	SetWrapText(newValue string)

	// Returns or sets vertical anchor text in a TextFrame.
	GetAnchoringType() string
	SetAnchoringType(newValue string)

	// If True then text should be centered in box horizontally.
	GetCenterText() string
	SetCenterText(newValue string)

	// Determines text orientation. The resulted value of visual text rotation summarized from this property and custom angle in property RotationAngle.
	GetTextVerticalType() string
	SetTextVerticalType(newValue string)

	// Returns or sets text's auto-fit mode.
	GetAutofitType() string
	SetAutofitType(newValue string)

	// Returns or sets number of columns in the text area. This value must be a positive number. Otherwise, the value will be set to zero.  Value 0 means undefined value.
	GetColumnCount() int32
	SetColumnCount(newValue int32)

	// Returns or sets the space between text columns in the text area (in points). This should only apply  when there is more than 1 column present. This value must be a positive number. Otherwise, the value will be set to zero. 
	GetColumnSpacing() float64
	SetColumnSpacing(newValue float64)

	// Returns or set keeping text out of 3D scene entirely.
	GetKeepTextFlat() *bool
	SetKeepTextFlat(newValue *bool)

	// Specifies the custom rotation that is being applied to the text within the bounding box.
	GetRotationAngle() float64
	SetRotationAngle(newValue float64)

	// Default portion format.
	GetDefaultParagraphFormat() IParagraphFormat
	SetDefaultParagraphFormat(newValue IParagraphFormat)
}

type TextFrameFormat struct {

	// Represents 3d effect properties for a text.
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets text wrapping shape.
	Transform string `json:"Transform,omitempty"`

	// Left margin. Left margin.
	MarginLeft float64 `json:"MarginLeft,omitempty"`

	// Right margin.
	MarginRight float64 `json:"MarginRight,omitempty"`

	// Top margin.
	MarginTop float64 `json:"MarginTop,omitempty"`

	// Bottom margin.
	MarginBottom float64 `json:"MarginBottom,omitempty"`

	// True if text is wrapped at TextFrame's margins.
	WrapText string `json:"WrapText,omitempty"`

	// Returns or sets vertical anchor text in a TextFrame.
	AnchoringType string `json:"AnchoringType,omitempty"`

	// If True then text should be centered in box horizontally.
	CenterText string `json:"CenterText,omitempty"`

	// Determines text orientation. The resulted value of visual text rotation summarized from this property and custom angle in property RotationAngle.
	TextVerticalType string `json:"TextVerticalType,omitempty"`

	// Returns or sets text's auto-fit mode.
	AutofitType string `json:"AutofitType,omitempty"`

	// Returns or sets number of columns in the text area. This value must be a positive number. Otherwise, the value will be set to zero.  Value 0 means undefined value.
	ColumnCount int32 `json:"ColumnCount,omitempty"`

	// Returns or sets the space between text columns in the text area (in points). This should only apply  when there is more than 1 column present. This value must be a positive number. Otherwise, the value will be set to zero. 
	ColumnSpacing float64 `json:"ColumnSpacing,omitempty"`

	// Returns or set keeping text out of 3D scene entirely.
	KeepTextFlat *bool `json:"KeepTextFlat"`

	// Specifies the custom rotation that is being applied to the text within the bounding box.
	RotationAngle float64 `json:"RotationAngle,omitempty"`

	// Default portion format.
	DefaultParagraphFormat IParagraphFormat `json:"DefaultParagraphFormat,omitempty"`
}

func NewTextFrameFormat() *TextFrameFormat {
	instance := new(TextFrameFormat)
	return instance
}

func (this *TextFrameFormat) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *TextFrameFormat) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *TextFrameFormat) GetTransform() string {
	return this.Transform
}

func (this *TextFrameFormat) SetTransform(newValue string) {
	this.Transform = newValue
}
func (this *TextFrameFormat) GetMarginLeft() float64 {
	return this.MarginLeft
}

func (this *TextFrameFormat) SetMarginLeft(newValue float64) {
	this.MarginLeft = newValue
}
func (this *TextFrameFormat) GetMarginRight() float64 {
	return this.MarginRight
}

func (this *TextFrameFormat) SetMarginRight(newValue float64) {
	this.MarginRight = newValue
}
func (this *TextFrameFormat) GetMarginTop() float64 {
	return this.MarginTop
}

func (this *TextFrameFormat) SetMarginTop(newValue float64) {
	this.MarginTop = newValue
}
func (this *TextFrameFormat) GetMarginBottom() float64 {
	return this.MarginBottom
}

func (this *TextFrameFormat) SetMarginBottom(newValue float64) {
	this.MarginBottom = newValue
}
func (this *TextFrameFormat) GetWrapText() string {
	return this.WrapText
}

func (this *TextFrameFormat) SetWrapText(newValue string) {
	this.WrapText = newValue
}
func (this *TextFrameFormat) GetAnchoringType() string {
	return this.AnchoringType
}

func (this *TextFrameFormat) SetAnchoringType(newValue string) {
	this.AnchoringType = newValue
}
func (this *TextFrameFormat) GetCenterText() string {
	return this.CenterText
}

func (this *TextFrameFormat) SetCenterText(newValue string) {
	this.CenterText = newValue
}
func (this *TextFrameFormat) GetTextVerticalType() string {
	return this.TextVerticalType
}

func (this *TextFrameFormat) SetTextVerticalType(newValue string) {
	this.TextVerticalType = newValue
}
func (this *TextFrameFormat) GetAutofitType() string {
	return this.AutofitType
}

func (this *TextFrameFormat) SetAutofitType(newValue string) {
	this.AutofitType = newValue
}
func (this *TextFrameFormat) GetColumnCount() int32 {
	return this.ColumnCount
}

func (this *TextFrameFormat) SetColumnCount(newValue int32) {
	this.ColumnCount = newValue
}
func (this *TextFrameFormat) GetColumnSpacing() float64 {
	return this.ColumnSpacing
}

func (this *TextFrameFormat) SetColumnSpacing(newValue float64) {
	this.ColumnSpacing = newValue
}
func (this *TextFrameFormat) GetKeepTextFlat() *bool {
	return this.KeepTextFlat
}

func (this *TextFrameFormat) SetKeepTextFlat(newValue *bool) {
	this.KeepTextFlat = newValue
}
func (this *TextFrameFormat) GetRotationAngle() float64 {
	return this.RotationAngle
}

func (this *TextFrameFormat) SetRotationAngle(newValue float64) {
	this.RotationAngle = newValue
}
func (this *TextFrameFormat) GetDefaultParagraphFormat() IParagraphFormat {
	return this.DefaultParagraphFormat
}

func (this *TextFrameFormat) SetDefaultParagraphFormat(newValue IParagraphFormat) {
	this.DefaultParagraphFormat = newValue
}

func (this *TextFrameFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ThreeDFormat", *valThreeDFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThreeDFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IThreeDFormat)
			if ok {
				this.ThreeDFormat = vInterfaceObject
			}
		}
	}
	
	if valTransform, ok := objMap["transform"]; ok {
		if valTransform != nil {
			var valueForTransform string
			err = json.Unmarshal(*valTransform, &valueForTransform)
			if err != nil {
				var valueForTransformInt int32
				err = json.Unmarshal(*valTransform, &valueForTransformInt)
				if err != nil {
					return err
				}
				this.Transform = string(valueForTransformInt)
			} else {
				this.Transform = valueForTransform
			}
		}
	}
	if valTransformCap, ok := objMap["Transform"]; ok {
		if valTransformCap != nil {
			var valueForTransform string
			err = json.Unmarshal(*valTransformCap, &valueForTransform)
			if err != nil {
				var valueForTransformInt int32
				err = json.Unmarshal(*valTransformCap, &valueForTransformInt)
				if err != nil {
					return err
				}
				this.Transform = string(valueForTransformInt)
			} else {
				this.Transform = valueForTransform
			}
		}
	}
	
	if valMarginLeft, ok := objMap["marginLeft"]; ok {
		if valMarginLeft != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeft, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	if valMarginLeftCap, ok := objMap["MarginLeft"]; ok {
		if valMarginLeftCap != nil {
			var valueForMarginLeft float64
			err = json.Unmarshal(*valMarginLeftCap, &valueForMarginLeft)
			if err != nil {
				return err
			}
			this.MarginLeft = valueForMarginLeft
		}
	}
	
	if valMarginRight, ok := objMap["marginRight"]; ok {
		if valMarginRight != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRight, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	if valMarginRightCap, ok := objMap["MarginRight"]; ok {
		if valMarginRightCap != nil {
			var valueForMarginRight float64
			err = json.Unmarshal(*valMarginRightCap, &valueForMarginRight)
			if err != nil {
				return err
			}
			this.MarginRight = valueForMarginRight
		}
	}
	
	if valMarginTop, ok := objMap["marginTop"]; ok {
		if valMarginTop != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTop, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}
	if valMarginTopCap, ok := objMap["MarginTop"]; ok {
		if valMarginTopCap != nil {
			var valueForMarginTop float64
			err = json.Unmarshal(*valMarginTopCap, &valueForMarginTop)
			if err != nil {
				return err
			}
			this.MarginTop = valueForMarginTop
		}
	}
	
	if valMarginBottom, ok := objMap["marginBottom"]; ok {
		if valMarginBottom != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottom, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}
	if valMarginBottomCap, ok := objMap["MarginBottom"]; ok {
		if valMarginBottomCap != nil {
			var valueForMarginBottom float64
			err = json.Unmarshal(*valMarginBottomCap, &valueForMarginBottom)
			if err != nil {
				return err
			}
			this.MarginBottom = valueForMarginBottom
		}
	}
	
	if valWrapText, ok := objMap["wrapText"]; ok {
		if valWrapText != nil {
			var valueForWrapText string
			err = json.Unmarshal(*valWrapText, &valueForWrapText)
			if err != nil {
				var valueForWrapTextInt int32
				err = json.Unmarshal(*valWrapText, &valueForWrapTextInt)
				if err != nil {
					return err
				}
				this.WrapText = string(valueForWrapTextInt)
			} else {
				this.WrapText = valueForWrapText
			}
		}
	}
	if valWrapTextCap, ok := objMap["WrapText"]; ok {
		if valWrapTextCap != nil {
			var valueForWrapText string
			err = json.Unmarshal(*valWrapTextCap, &valueForWrapText)
			if err != nil {
				var valueForWrapTextInt int32
				err = json.Unmarshal(*valWrapTextCap, &valueForWrapTextInt)
				if err != nil {
					return err
				}
				this.WrapText = string(valueForWrapTextInt)
			} else {
				this.WrapText = valueForWrapText
			}
		}
	}
	
	if valAnchoringType, ok := objMap["anchoringType"]; ok {
		if valAnchoringType != nil {
			var valueForAnchoringType string
			err = json.Unmarshal(*valAnchoringType, &valueForAnchoringType)
			if err != nil {
				var valueForAnchoringTypeInt int32
				err = json.Unmarshal(*valAnchoringType, &valueForAnchoringTypeInt)
				if err != nil {
					return err
				}
				this.AnchoringType = string(valueForAnchoringTypeInt)
			} else {
				this.AnchoringType = valueForAnchoringType
			}
		}
	}
	if valAnchoringTypeCap, ok := objMap["AnchoringType"]; ok {
		if valAnchoringTypeCap != nil {
			var valueForAnchoringType string
			err = json.Unmarshal(*valAnchoringTypeCap, &valueForAnchoringType)
			if err != nil {
				var valueForAnchoringTypeInt int32
				err = json.Unmarshal(*valAnchoringTypeCap, &valueForAnchoringTypeInt)
				if err != nil {
					return err
				}
				this.AnchoringType = string(valueForAnchoringTypeInt)
			} else {
				this.AnchoringType = valueForAnchoringType
			}
		}
	}
	
	if valCenterText, ok := objMap["centerText"]; ok {
		if valCenterText != nil {
			var valueForCenterText string
			err = json.Unmarshal(*valCenterText, &valueForCenterText)
			if err != nil {
				var valueForCenterTextInt int32
				err = json.Unmarshal(*valCenterText, &valueForCenterTextInt)
				if err != nil {
					return err
				}
				this.CenterText = string(valueForCenterTextInt)
			} else {
				this.CenterText = valueForCenterText
			}
		}
	}
	if valCenterTextCap, ok := objMap["CenterText"]; ok {
		if valCenterTextCap != nil {
			var valueForCenterText string
			err = json.Unmarshal(*valCenterTextCap, &valueForCenterText)
			if err != nil {
				var valueForCenterTextInt int32
				err = json.Unmarshal(*valCenterTextCap, &valueForCenterTextInt)
				if err != nil {
					return err
				}
				this.CenterText = string(valueForCenterTextInt)
			} else {
				this.CenterText = valueForCenterText
			}
		}
	}
	
	if valTextVerticalType, ok := objMap["textVerticalType"]; ok {
		if valTextVerticalType != nil {
			var valueForTextVerticalType string
			err = json.Unmarshal(*valTextVerticalType, &valueForTextVerticalType)
			if err != nil {
				var valueForTextVerticalTypeInt int32
				err = json.Unmarshal(*valTextVerticalType, &valueForTextVerticalTypeInt)
				if err != nil {
					return err
				}
				this.TextVerticalType = string(valueForTextVerticalTypeInt)
			} else {
				this.TextVerticalType = valueForTextVerticalType
			}
		}
	}
	if valTextVerticalTypeCap, ok := objMap["TextVerticalType"]; ok {
		if valTextVerticalTypeCap != nil {
			var valueForTextVerticalType string
			err = json.Unmarshal(*valTextVerticalTypeCap, &valueForTextVerticalType)
			if err != nil {
				var valueForTextVerticalTypeInt int32
				err = json.Unmarshal(*valTextVerticalTypeCap, &valueForTextVerticalTypeInt)
				if err != nil {
					return err
				}
				this.TextVerticalType = string(valueForTextVerticalTypeInt)
			} else {
				this.TextVerticalType = valueForTextVerticalType
			}
		}
	}
	
	if valAutofitType, ok := objMap["autofitType"]; ok {
		if valAutofitType != nil {
			var valueForAutofitType string
			err = json.Unmarshal(*valAutofitType, &valueForAutofitType)
			if err != nil {
				var valueForAutofitTypeInt int32
				err = json.Unmarshal(*valAutofitType, &valueForAutofitTypeInt)
				if err != nil {
					return err
				}
				this.AutofitType = string(valueForAutofitTypeInt)
			} else {
				this.AutofitType = valueForAutofitType
			}
		}
	}
	if valAutofitTypeCap, ok := objMap["AutofitType"]; ok {
		if valAutofitTypeCap != nil {
			var valueForAutofitType string
			err = json.Unmarshal(*valAutofitTypeCap, &valueForAutofitType)
			if err != nil {
				var valueForAutofitTypeInt int32
				err = json.Unmarshal(*valAutofitTypeCap, &valueForAutofitTypeInt)
				if err != nil {
					return err
				}
				this.AutofitType = string(valueForAutofitTypeInt)
			} else {
				this.AutofitType = valueForAutofitType
			}
		}
	}
	
	if valColumnCount, ok := objMap["columnCount"]; ok {
		if valColumnCount != nil {
			var valueForColumnCount int32
			err = json.Unmarshal(*valColumnCount, &valueForColumnCount)
			if err != nil {
				return err
			}
			this.ColumnCount = valueForColumnCount
		}
	}
	if valColumnCountCap, ok := objMap["ColumnCount"]; ok {
		if valColumnCountCap != nil {
			var valueForColumnCount int32
			err = json.Unmarshal(*valColumnCountCap, &valueForColumnCount)
			if err != nil {
				return err
			}
			this.ColumnCount = valueForColumnCount
		}
	}
	
	if valColumnSpacing, ok := objMap["columnSpacing"]; ok {
		if valColumnSpacing != nil {
			var valueForColumnSpacing float64
			err = json.Unmarshal(*valColumnSpacing, &valueForColumnSpacing)
			if err != nil {
				return err
			}
			this.ColumnSpacing = valueForColumnSpacing
		}
	}
	if valColumnSpacingCap, ok := objMap["ColumnSpacing"]; ok {
		if valColumnSpacingCap != nil {
			var valueForColumnSpacing float64
			err = json.Unmarshal(*valColumnSpacingCap, &valueForColumnSpacing)
			if err != nil {
				return err
			}
			this.ColumnSpacing = valueForColumnSpacing
		}
	}
	
	if valKeepTextFlat, ok := objMap["keepTextFlat"]; ok {
		if valKeepTextFlat != nil {
			var valueForKeepTextFlat *bool
			err = json.Unmarshal(*valKeepTextFlat, &valueForKeepTextFlat)
			if err != nil {
				return err
			}
			this.KeepTextFlat = valueForKeepTextFlat
		}
	}
	if valKeepTextFlatCap, ok := objMap["KeepTextFlat"]; ok {
		if valKeepTextFlatCap != nil {
			var valueForKeepTextFlat *bool
			err = json.Unmarshal(*valKeepTextFlatCap, &valueForKeepTextFlat)
			if err != nil {
				return err
			}
			this.KeepTextFlat = valueForKeepTextFlat
		}
	}
	
	if valRotationAngle, ok := objMap["rotationAngle"]; ok {
		if valRotationAngle != nil {
			var valueForRotationAngle float64
			err = json.Unmarshal(*valRotationAngle, &valueForRotationAngle)
			if err != nil {
				return err
			}
			this.RotationAngle = valueForRotationAngle
		}
	}
	if valRotationAngleCap, ok := objMap["RotationAngle"]; ok {
		if valRotationAngleCap != nil {
			var valueForRotationAngle float64
			err = json.Unmarshal(*valRotationAngleCap, &valueForRotationAngle)
			if err != nil {
				return err
			}
			this.RotationAngle = valueForRotationAngle
		}
	}
	
	if valDefaultParagraphFormat, ok := objMap["defaultParagraphFormat"]; ok {
		if valDefaultParagraphFormat != nil {
			var valueForDefaultParagraphFormat ParagraphFormat
			err = json.Unmarshal(*valDefaultParagraphFormat, &valueForDefaultParagraphFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ParagraphFormat", *valDefaultParagraphFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDefaultParagraphFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IParagraphFormat)
			if ok {
				this.DefaultParagraphFormat = vInterfaceObject
			}
		}
	}
	if valDefaultParagraphFormatCap, ok := objMap["DefaultParagraphFormat"]; ok {
		if valDefaultParagraphFormatCap != nil {
			var valueForDefaultParagraphFormat ParagraphFormat
			err = json.Unmarshal(*valDefaultParagraphFormatCap, &valueForDefaultParagraphFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ParagraphFormat", *valDefaultParagraphFormatCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valDefaultParagraphFormatCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IParagraphFormat)
			if ok {
				this.DefaultParagraphFormat = vInterfaceObject
			}
		}
	}

	return nil
}
