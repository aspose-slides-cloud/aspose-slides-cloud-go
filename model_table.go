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

// Represents Table shape resource.
type ITable interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	GetName() string
	SetName(newValue string)

	// Gets or sets the width.
	GetWidth() float64
	SetWidth(newValue float64)

	// Gets or sets the height.
	GetHeight() float64
	SetHeight(newValue float64)

	// Gets or sets the alternative text.
	GetAlternativeText() string
	SetAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	GetAlternativeTextTitle() string
	SetAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	GetHidden() *bool
	SetHidden(newValue *bool)

	// Gets or sets 'Mark as decorative' option.
	GetIsDecorative() *bool
	SetIsDecorative(newValue *bool)

	// Gets or sets the X
	GetX() float64
	SetX(newValue float64)

	// Gets or sets the Y.
	GetY() float64
	SetY(newValue float64)

	// Gets z-order position of shape
	GetZOrderPosition() int32
	SetZOrderPosition(newValue int32)

	// Gets or sets the fill format.
	GetFillFormat() IFillFormat
	SetFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	GetEffectFormat() IEffectFormat
	SetEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	GetThreeDFormat() IThreeDFormat
	SetThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	GetLineFormat() ILineFormat
	SetLineFormat(newValue ILineFormat)

	// Hyperlink defined for mouse click.
	GetHyperlinkClick() IHyperlink
	SetHyperlinkClick(newValue IHyperlink)

	// Hyperlink defined for mouse over.
	GetHyperlinkMouseOver() IHyperlink
	SetHyperlinkMouseOver(newValue IHyperlink)

	// Shape type.
	GetType() string
	SetType(newValue string)

	// Builtin table style.
	GetStyle() string
	SetStyle(newValue string)

	// Rows.
	GetRows() []ITableRow
	SetRows(newValue []ITableRow)

	// Columns.
	GetColumns() []ITableColumn
	SetColumns(newValue []ITableColumn)

	// Determines whether the first column of a table has to be drawn with a special formatting.
	GetFirstCol() *bool
	SetFirstCol(newValue *bool)

	// Determines whether the first row of a table has to be drawn with a special formatting.
	GetFirstRow() *bool
	SetFirstRow(newValue *bool)

	// Determines whether the even rows has to be drawn with a different formatting.
	GetHorizontalBanding() *bool
	SetHorizontalBanding(newValue *bool)

	// Determines whether the last column of a table has to be drawn with a special formatting.
	GetLastCol() *bool
	SetLastCol(newValue *bool)

	// Determines whether the last row of a table has to be drawn with a special formatting.
	GetLastRow() *bool
	SetLastRow(newValue *bool)

	// Determines whether the table has right to left reading order.
	GetRightToLeft() *bool
	SetRightToLeft(newValue *bool)

	// Determines whether the even columns has to be drawn with a different formatting.
	GetVerticalBanding() *bool
	SetVerticalBanding(newValue *bool)

	// Transparency.
	GetTransparency() float64
	SetTransparency(newValue float64)
}

type Table struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// The title of alternative text associated with the shape.
	AlternativeTextTitle string `json:"AlternativeTextTitle,omitempty"`

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	Hidden *bool `json:"Hidden"`

	// Gets or sets 'Mark as decorative' option.
	IsDecorative *bool `json:"IsDecorative"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Hyperlink defined for mouse click.
	HyperlinkClick IHyperlink `json:"HyperlinkClick,omitempty"`

	// Hyperlink defined for mouse over.
	HyperlinkMouseOver IHyperlink `json:"HyperlinkMouseOver,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Builtin table style.
	Style string `json:"Style,omitempty"`

	// Rows.
	Rows []ITableRow `json:"Rows,omitempty"`

	// Columns.
	Columns []ITableColumn `json:"Columns,omitempty"`

	// Determines whether the first column of a table has to be drawn with a special formatting.
	FirstCol *bool `json:"FirstCol"`

	// Determines whether the first row of a table has to be drawn with a special formatting.
	FirstRow *bool `json:"FirstRow"`

	// Determines whether the even rows has to be drawn with a different formatting.
	HorizontalBanding *bool `json:"HorizontalBanding"`

	// Determines whether the last column of a table has to be drawn with a special formatting.
	LastCol *bool `json:"LastCol"`

	// Determines whether the last row of a table has to be drawn with a special formatting.
	LastRow *bool `json:"LastRow"`

	// Determines whether the table has right to left reading order.
	RightToLeft *bool `json:"RightToLeft"`

	// Determines whether the even columns has to be drawn with a different formatting.
	VerticalBanding *bool `json:"VerticalBanding"`

	// Transparency.
	Transparency float64 `json:"Transparency,omitempty"`
}

func NewTable() *Table {
	instance := new(Table)
	instance.Type_ = "Table"
	return instance
}

func (this *Table) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Table) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Table) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Table) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Table) GetName() string {
	return this.Name
}

func (this *Table) SetName(newValue string) {
	this.Name = newValue
}
func (this *Table) GetWidth() float64 {
	return this.Width
}

func (this *Table) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *Table) GetHeight() float64 {
	return this.Height
}

func (this *Table) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *Table) GetAlternativeText() string {
	return this.AlternativeText
}

func (this *Table) SetAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *Table) GetAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *Table) SetAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *Table) GetHidden() *bool {
	return this.Hidden
}

func (this *Table) SetHidden(newValue *bool) {
	this.Hidden = newValue
}
func (this *Table) GetIsDecorative() *bool {
	return this.IsDecorative
}

func (this *Table) SetIsDecorative(newValue *bool) {
	this.IsDecorative = newValue
}
func (this *Table) GetX() float64 {
	return this.X
}

func (this *Table) SetX(newValue float64) {
	this.X = newValue
}
func (this *Table) GetY() float64 {
	return this.Y
}

func (this *Table) SetY(newValue float64) {
	this.Y = newValue
}
func (this *Table) GetZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *Table) SetZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *Table) GetFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Table) SetFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Table) GetEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Table) SetEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Table) GetThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *Table) SetThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *Table) GetLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Table) SetLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Table) GetHyperlinkClick() IHyperlink {
	return this.HyperlinkClick
}

func (this *Table) SetHyperlinkClick(newValue IHyperlink) {
	this.HyperlinkClick = newValue
}
func (this *Table) GetHyperlinkMouseOver() IHyperlink {
	return this.HyperlinkMouseOver
}

func (this *Table) SetHyperlinkMouseOver(newValue IHyperlink) {
	this.HyperlinkMouseOver = newValue
}
func (this *Table) GetType() string {
	return this.Type_
}

func (this *Table) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Table) GetStyle() string {
	return this.Style
}

func (this *Table) SetStyle(newValue string) {
	this.Style = newValue
}
func (this *Table) GetRows() []ITableRow {
	return this.Rows
}

func (this *Table) SetRows(newValue []ITableRow) {
	this.Rows = newValue
}
func (this *Table) GetColumns() []ITableColumn {
	return this.Columns
}

func (this *Table) SetColumns(newValue []ITableColumn) {
	this.Columns = newValue
}
func (this *Table) GetFirstCol() *bool {
	return this.FirstCol
}

func (this *Table) SetFirstCol(newValue *bool) {
	this.FirstCol = newValue
}
func (this *Table) GetFirstRow() *bool {
	return this.FirstRow
}

func (this *Table) SetFirstRow(newValue *bool) {
	this.FirstRow = newValue
}
func (this *Table) GetHorizontalBanding() *bool {
	return this.HorizontalBanding
}

func (this *Table) SetHorizontalBanding(newValue *bool) {
	this.HorizontalBanding = newValue
}
func (this *Table) GetLastCol() *bool {
	return this.LastCol
}

func (this *Table) SetLastCol(newValue *bool) {
	this.LastCol = newValue
}
func (this *Table) GetLastRow() *bool {
	return this.LastRow
}

func (this *Table) SetLastRow(newValue *bool) {
	this.LastRow = newValue
}
func (this *Table) GetRightToLeft() *bool {
	return this.RightToLeft
}

func (this *Table) SetRightToLeft(newValue *bool) {
	this.RightToLeft = newValue
}
func (this *Table) GetVerticalBanding() *bool {
	return this.VerticalBanding
}

func (this *Table) SetVerticalBanding(newValue *bool) {
	this.VerticalBanding = newValue
}
func (this *Table) GetTransparency() float64 {
	return this.Transparency
}

func (this *Table) SetTransparency(newValue float64) {
	this.Transparency = newValue
}

func (this *Table) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := GetMapValue(objMap, "selfUri"); ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUri)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUri, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
		}
	}
	
	if valAlternateLinks, ok := GetMapValue(objMap, "alternateLinks"); ok {
		if valAlternateLinks != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinks, &valueForAlternateLinks)
			if err != nil {
				return err
			}
			valueForIAlternateLinks := make([]IResourceUri, len(valueForAlternateLinks))
			for i, v := range valueForAlternateLinks {
				vObject, err := createObjectForType("ResourceUri", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIAlternateLinks[i] = vObject.(IResourceUri)
				}
			}
			this.AlternateLinks = valueForIAlternateLinks
		}
	}
	
	if valName, ok := GetMapValue(objMap, "name"); ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := GetMapValue(objMap, "width"); ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := GetMapValue(objMap, "height"); ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := GetMapValue(objMap, "alternativeText"); ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := GetMapValue(objMap, "alternativeTextTitle"); ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := GetMapValue(objMap, "hidden"); ok {
		if valHidden != nil {
			var valueForHidden *bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valIsDecorative, ok := GetMapValue(objMap, "isDecorative"); ok {
		if valIsDecorative != nil {
			var valueForIsDecorative *bool
			err = json.Unmarshal(*valIsDecorative, &valueForIsDecorative)
			if err != nil {
				return err
			}
			this.IsDecorative = valueForIsDecorative
		}
	}
	
	if valX, ok := GetMapValue(objMap, "x"); ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := GetMapValue(objMap, "y"); ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := GetMapValue(objMap, "zOrderPosition"); ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valFillFormat, ok := GetMapValue(objMap, "fillFormat"); ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillFormat", *valFillFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillFormat)
			if ok {
				this.FillFormat = vInterfaceObject
			}
		}
	}
	
	if valEffectFormat, ok := GetMapValue(objMap, "effectFormat"); ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("EffectFormat", *valEffectFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valEffectFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IEffectFormat)
			if ok {
				this.EffectFormat = vInterfaceObject
			}
		}
	}
	
	if valThreeDFormat, ok := GetMapValue(objMap, "threeDFormat"); ok {
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
	
	if valLineFormat, ok := GetMapValue(objMap, "lineFormat"); ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LineFormat", *valLineFormat)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLineFormat, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILineFormat)
			if ok {
				this.LineFormat = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkClick, ok := GetMapValue(objMap, "hyperlinkClick"); ok {
		if valHyperlinkClick != nil {
			var valueForHyperlinkClick Hyperlink
			err = json.Unmarshal(*valHyperlinkClick, &valueForHyperlinkClick)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkClick)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkClick, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkClick = vInterfaceObject
			}
		}
	}
	
	if valHyperlinkMouseOver, ok := GetMapValue(objMap, "hyperlinkMouseOver"); ok {
		if valHyperlinkMouseOver != nil {
			var valueForHyperlinkMouseOver Hyperlink
			err = json.Unmarshal(*valHyperlinkMouseOver, &valueForHyperlinkMouseOver)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Hyperlink", *valHyperlinkMouseOver)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valHyperlinkMouseOver, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IHyperlink)
			if ok {
				this.HyperlinkMouseOver = vInterfaceObject
			}
		}
	}
	this.Type_ = "Table"
	if valType, ok := GetMapValue(objMap, "type"); ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valStyle, ok := GetMapValue(objMap, "style"); ok {
		if valStyle != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyle, &valueForStyle)
			if err != nil {
				var valueForStyleInt int32
				err = json.Unmarshal(*valStyle, &valueForStyleInt)
				if err != nil {
					return err
				}
				this.Style = string(valueForStyleInt)
			} else {
				this.Style = valueForStyle
			}
		}
	}
	
	if valRows, ok := GetMapValue(objMap, "rows"); ok {
		if valRows != nil {
			var valueForRows []json.RawMessage
			err = json.Unmarshal(*valRows, &valueForRows)
			if err != nil {
				return err
			}
			valueForIRows := make([]ITableRow, len(valueForRows))
			for i, v := range valueForRows {
				vObject, err := createObjectForType("TableRow", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIRows[i] = vObject.(ITableRow)
				}
			}
			this.Rows = valueForIRows
		}
	}
	
	if valColumns, ok := GetMapValue(objMap, "columns"); ok {
		if valColumns != nil {
			var valueForColumns []json.RawMessage
			err = json.Unmarshal(*valColumns, &valueForColumns)
			if err != nil {
				return err
			}
			valueForIColumns := make([]ITableColumn, len(valueForColumns))
			for i, v := range valueForColumns {
				vObject, err := createObjectForType("TableColumn", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIColumns[i] = vObject.(ITableColumn)
				}
			}
			this.Columns = valueForIColumns
		}
	}
	
	if valFirstCol, ok := GetMapValue(objMap, "firstCol"); ok {
		if valFirstCol != nil {
			var valueForFirstCol *bool
			err = json.Unmarshal(*valFirstCol, &valueForFirstCol)
			if err != nil {
				return err
			}
			this.FirstCol = valueForFirstCol
		}
	}
	
	if valFirstRow, ok := GetMapValue(objMap, "firstRow"); ok {
		if valFirstRow != nil {
			var valueForFirstRow *bool
			err = json.Unmarshal(*valFirstRow, &valueForFirstRow)
			if err != nil {
				return err
			}
			this.FirstRow = valueForFirstRow
		}
	}
	
	if valHorizontalBanding, ok := GetMapValue(objMap, "horizontalBanding"); ok {
		if valHorizontalBanding != nil {
			var valueForHorizontalBanding *bool
			err = json.Unmarshal(*valHorizontalBanding, &valueForHorizontalBanding)
			if err != nil {
				return err
			}
			this.HorizontalBanding = valueForHorizontalBanding
		}
	}
	
	if valLastCol, ok := GetMapValue(objMap, "lastCol"); ok {
		if valLastCol != nil {
			var valueForLastCol *bool
			err = json.Unmarshal(*valLastCol, &valueForLastCol)
			if err != nil {
				return err
			}
			this.LastCol = valueForLastCol
		}
	}
	
	if valLastRow, ok := GetMapValue(objMap, "lastRow"); ok {
		if valLastRow != nil {
			var valueForLastRow *bool
			err = json.Unmarshal(*valLastRow, &valueForLastRow)
			if err != nil {
				return err
			}
			this.LastRow = valueForLastRow
		}
	}
	
	if valRightToLeft, ok := GetMapValue(objMap, "rightToLeft"); ok {
		if valRightToLeft != nil {
			var valueForRightToLeft *bool
			err = json.Unmarshal(*valRightToLeft, &valueForRightToLeft)
			if err != nil {
				return err
			}
			this.RightToLeft = valueForRightToLeft
		}
	}
	
	if valVerticalBanding, ok := GetMapValue(objMap, "verticalBanding"); ok {
		if valVerticalBanding != nil {
			var valueForVerticalBanding *bool
			err = json.Unmarshal(*valVerticalBanding, &valueForVerticalBanding)
			if err != nil {
				return err
			}
			this.VerticalBanding = valueForVerticalBanding
		}
	}
	
	if valTransparency, ok := GetMapValue(objMap, "transparency"); ok {
		if valTransparency != nil {
			var valueForTransparency float64
			err = json.Unmarshal(*valTransparency, &valueForTransparency)
			if err != nil {
				return err
			}
			this.Transparency = valueForTransparency
		}
	}

	return nil
}
