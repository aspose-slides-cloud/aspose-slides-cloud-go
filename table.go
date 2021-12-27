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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Gets or sets the name.
	getName() string
	setName(newValue string)

	// Gets or sets the width.
	getWidth() float64
	setWidth(newValue float64)

	// Gets or sets the height.
	getHeight() float64
	setHeight(newValue float64)

	// Gets or sets the alternative text.
	getAlternativeText() string
	setAlternativeText(newValue string)

	// The title of alternative text associated with the shape.
	getAlternativeTextTitle() string
	setAlternativeTextTitle(newValue string)

	// Gets or sets a value indicating whether this ShapeBase is hidden.
	getHidden() bool
	setHidden(newValue bool)

	// Gets or sets the X
	getX() float64
	setX(newValue float64)

	// Gets or sets the Y.
	getY() float64
	setY(newValue float64)

	// Gets z-order position of shape
	getZOrderPosition() int32
	setZOrderPosition(newValue int32)

	// Gets or sets the link to shapes.
	getShapes() IResourceUri
	setShapes(newValue IResourceUri)

	// Gets or sets the fill format.
	getFillFormat() IFillFormat
	setFillFormat(newValue IFillFormat)

	// Gets or sets the effect format.
	getEffectFormat() IEffectFormat
	setEffectFormat(newValue IEffectFormat)

	// Gets or sets the 3D format
	getThreeDFormat() IThreeDFormat
	setThreeDFormat(newValue IThreeDFormat)

	// Gets or sets the line format.
	getLineFormat() ILineFormat
	setLineFormat(newValue ILineFormat)

	// Shape type.
	getType() string
	setType(newValue string)

	// Builtin table style.
	getStyle() string
	setStyle(newValue string)

	// Rows.
	getRows() []ITableRow
	setRows(newValue []ITableRow)

	// Columns.
	getColumns() []ITableColumn
	setColumns(newValue []ITableColumn)

	// Determines whether the first column of a table has to be drawn with a special formatting.
	getFirstCol() bool
	setFirstCol(newValue bool)

	// Determines whether the first row of a table has to be drawn with a special formatting.
	getFirstRow() bool
	setFirstRow(newValue bool)

	// Determines whether the even rows has to be drawn with a different formatting.
	getHorizontalBanding() bool
	setHorizontalBanding(newValue bool)

	// Determines whether the last column of a table has to be drawn with a special formatting.
	getLastCol() bool
	setLastCol(newValue bool)

	// Determines whether the last row of a table has to be drawn with a special formatting.
	getLastRow() bool
	setLastRow(newValue bool)

	// Determines whether the table has right to left reading order.
	getRightToLeft() bool
	setRightToLeft(newValue bool)

	// Determines whether the even columns has to be drawn with a different formatting.
	getVerticalBanding() bool
	setVerticalBanding(newValue bool)
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
	Hidden bool `json:"Hidden"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition"`

	// Gets or sets the link to shapes.
	Shapes IResourceUri `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat IFillFormat `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat IEffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the 3D format
	ThreeDFormat IThreeDFormat `json:"ThreeDFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat ILineFormat `json:"LineFormat,omitempty"`

	// Shape type.
	Type_ string `json:"Type"`

	// Builtin table style.
	Style string `json:"Style,omitempty"`

	// Rows.
	Rows []ITableRow `json:"Rows,omitempty"`

	// Columns.
	Columns []ITableColumn `json:"Columns,omitempty"`

	// Determines whether the first column of a table has to be drawn with a special formatting.
	FirstCol bool `json:"FirstCol"`

	// Determines whether the first row of a table has to be drawn with a special formatting.
	FirstRow bool `json:"FirstRow"`

	// Determines whether the even rows has to be drawn with a different formatting.
	HorizontalBanding bool `json:"HorizontalBanding"`

	// Determines whether the last column of a table has to be drawn with a special formatting.
	LastCol bool `json:"LastCol"`

	// Determines whether the last row of a table has to be drawn with a special formatting.
	LastRow bool `json:"LastRow"`

	// Determines whether the table has right to left reading order.
	RightToLeft bool `json:"RightToLeft"`

	// Determines whether the even columns has to be drawn with a different formatting.
	VerticalBanding bool `json:"VerticalBanding"`
}

func NewTable() *Table {
	instance := new(Table)
	instance.Type_ = "Table"
	instance.Style = ""
	return instance
}

func (this *Table) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *Table) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *Table) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *Table) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *Table) getName() string {
	return this.Name
}

func (this *Table) setName(newValue string) {
	this.Name = newValue
}
func (this *Table) getWidth() float64 {
	return this.Width
}

func (this *Table) setWidth(newValue float64) {
	this.Width = newValue
}
func (this *Table) getHeight() float64 {
	return this.Height
}

func (this *Table) setHeight(newValue float64) {
	this.Height = newValue
}
func (this *Table) getAlternativeText() string {
	return this.AlternativeText
}

func (this *Table) setAlternativeText(newValue string) {
	this.AlternativeText = newValue
}
func (this *Table) getAlternativeTextTitle() string {
	return this.AlternativeTextTitle
}

func (this *Table) setAlternativeTextTitle(newValue string) {
	this.AlternativeTextTitle = newValue
}
func (this *Table) getHidden() bool {
	return this.Hidden
}

func (this *Table) setHidden(newValue bool) {
	this.Hidden = newValue
}
func (this *Table) getX() float64 {
	return this.X
}

func (this *Table) setX(newValue float64) {
	this.X = newValue
}
func (this *Table) getY() float64 {
	return this.Y
}

func (this *Table) setY(newValue float64) {
	this.Y = newValue
}
func (this *Table) getZOrderPosition() int32 {
	return this.ZOrderPosition
}

func (this *Table) setZOrderPosition(newValue int32) {
	this.ZOrderPosition = newValue
}
func (this *Table) getShapes() IResourceUri {
	return this.Shapes
}

func (this *Table) setShapes(newValue IResourceUri) {
	this.Shapes = newValue
}
func (this *Table) getFillFormat() IFillFormat {
	return this.FillFormat
}

func (this *Table) setFillFormat(newValue IFillFormat) {
	this.FillFormat = newValue
}
func (this *Table) getEffectFormat() IEffectFormat {
	return this.EffectFormat
}

func (this *Table) setEffectFormat(newValue IEffectFormat) {
	this.EffectFormat = newValue
}
func (this *Table) getThreeDFormat() IThreeDFormat {
	return this.ThreeDFormat
}

func (this *Table) setThreeDFormat(newValue IThreeDFormat) {
	this.ThreeDFormat = newValue
}
func (this *Table) getLineFormat() ILineFormat {
	return this.LineFormat
}

func (this *Table) setLineFormat(newValue ILineFormat) {
	this.LineFormat = newValue
}
func (this *Table) getType() string {
	return this.Type_
}

func (this *Table) setType(newValue string) {
	this.Type_ = newValue
}
func (this *Table) getStyle() string {
	return this.Style
}

func (this *Table) setStyle(newValue string) {
	this.Style = newValue
}
func (this *Table) getRows() []ITableRow {
	return this.Rows
}

func (this *Table) setRows(newValue []ITableRow) {
	this.Rows = newValue
}
func (this *Table) getColumns() []ITableColumn {
	return this.Columns
}

func (this *Table) setColumns(newValue []ITableColumn) {
	this.Columns = newValue
}
func (this *Table) getFirstCol() bool {
	return this.FirstCol
}

func (this *Table) setFirstCol(newValue bool) {
	this.FirstCol = newValue
}
func (this *Table) getFirstRow() bool {
	return this.FirstRow
}

func (this *Table) setFirstRow(newValue bool) {
	this.FirstRow = newValue
}
func (this *Table) getHorizontalBanding() bool {
	return this.HorizontalBanding
}

func (this *Table) setHorizontalBanding(newValue bool) {
	this.HorizontalBanding = newValue
}
func (this *Table) getLastCol() bool {
	return this.LastCol
}

func (this *Table) setLastCol(newValue bool) {
	this.LastCol = newValue
}
func (this *Table) getLastRow() bool {
	return this.LastRow
}

func (this *Table) setLastRow(newValue bool) {
	this.LastRow = newValue
}
func (this *Table) getRightToLeft() bool {
	return this.RightToLeft
}

func (this *Table) setRightToLeft(newValue bool) {
	this.RightToLeft = newValue
}
func (this *Table) getVerticalBanding() bool {
	return this.VerticalBanding
}

func (this *Table) setVerticalBanding(newValue bool) {
	this.VerticalBanding = newValue
}

func (this *Table) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valSelfUri, ok := objMap["selfUri"]; ok {
		if valSelfUri != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUri, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			this.SelfUri = &valueForSelfUri
		}
	}
	
	if valAlternateLinks, ok := objMap["alternateLinks"]; ok {
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
	if valAlternateLinksCap, ok := objMap["AlternateLinks"]; ok {
		if valAlternateLinksCap != nil {
			var valueForAlternateLinks []json.RawMessage
			err = json.Unmarshal(*valAlternateLinksCap, &valueForAlternateLinks)
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
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valWidth, ok := objMap["width"]; ok {
		if valWidth != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidth, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	if valWidthCap, ok := objMap["Width"]; ok {
		if valWidthCap != nil {
			var valueForWidth float64
			err = json.Unmarshal(*valWidthCap, &valueForWidth)
			if err != nil {
				return err
			}
			this.Width = valueForWidth
		}
	}
	
	if valHeight, ok := objMap["height"]; ok {
		if valHeight != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeight, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	if valHeightCap, ok := objMap["Height"]; ok {
		if valHeightCap != nil {
			var valueForHeight float64
			err = json.Unmarshal(*valHeightCap, &valueForHeight)
			if err != nil {
				return err
			}
			this.Height = valueForHeight
		}
	}
	
	if valAlternativeText, ok := objMap["alternativeText"]; ok {
		if valAlternativeText != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeText, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	if valAlternativeTextCap, ok := objMap["AlternativeText"]; ok {
		if valAlternativeTextCap != nil {
			var valueForAlternativeText string
			err = json.Unmarshal(*valAlternativeTextCap, &valueForAlternativeText)
			if err != nil {
				return err
			}
			this.AlternativeText = valueForAlternativeText
		}
	}
	
	if valAlternativeTextTitle, ok := objMap["alternativeTextTitle"]; ok {
		if valAlternativeTextTitle != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitle, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	if valAlternativeTextTitleCap, ok := objMap["AlternativeTextTitle"]; ok {
		if valAlternativeTextTitleCap != nil {
			var valueForAlternativeTextTitle string
			err = json.Unmarshal(*valAlternativeTextTitleCap, &valueForAlternativeTextTitle)
			if err != nil {
				return err
			}
			this.AlternativeTextTitle = valueForAlternativeTextTitle
		}
	}
	
	if valHidden, ok := objMap["hidden"]; ok {
		if valHidden != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHidden, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	if valHiddenCap, ok := objMap["Hidden"]; ok {
		if valHiddenCap != nil {
			var valueForHidden bool
			err = json.Unmarshal(*valHiddenCap, &valueForHidden)
			if err != nil {
				return err
			}
			this.Hidden = valueForHidden
		}
	}
	
	if valX, ok := objMap["x"]; ok {
		if valX != nil {
			var valueForX float64
			err = json.Unmarshal(*valX, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	if valXCap, ok := objMap["X"]; ok {
		if valXCap != nil {
			var valueForX float64
			err = json.Unmarshal(*valXCap, &valueForX)
			if err != nil {
				return err
			}
			this.X = valueForX
		}
	}
	
	if valY, ok := objMap["y"]; ok {
		if valY != nil {
			var valueForY float64
			err = json.Unmarshal(*valY, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	if valYCap, ok := objMap["Y"]; ok {
		if valYCap != nil {
			var valueForY float64
			err = json.Unmarshal(*valYCap, &valueForY)
			if err != nil {
				return err
			}
			this.Y = valueForY
		}
	}
	
	if valZOrderPosition, ok := objMap["zOrderPosition"]; ok {
		if valZOrderPosition != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPosition, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	if valZOrderPositionCap, ok := objMap["ZOrderPosition"]; ok {
		if valZOrderPositionCap != nil {
			var valueForZOrderPosition int32
			err = json.Unmarshal(*valZOrderPositionCap, &valueForZOrderPosition)
			if err != nil {
				return err
			}
			this.ZOrderPosition = valueForZOrderPosition
		}
	}
	
	if valShapes, ok := objMap["shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapesCap, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = &valueForShapes
		}
	}
	
	if valFillFormat, ok := objMap["fillFormat"]; ok {
		if valFillFormat != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormat, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	if valFillFormatCap, ok := objMap["FillFormat"]; ok {
		if valFillFormatCap != nil {
			var valueForFillFormat FillFormat
			err = json.Unmarshal(*valFillFormatCap, &valueForFillFormat)
			if err != nil {
				return err
			}
			this.FillFormat = &valueForFillFormat
		}
	}
	
	if valEffectFormat, ok := objMap["effectFormat"]; ok {
		if valEffectFormat != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormat, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	if valEffectFormatCap, ok := objMap["EffectFormat"]; ok {
		if valEffectFormatCap != nil {
			var valueForEffectFormat EffectFormat
			err = json.Unmarshal(*valEffectFormatCap, &valueForEffectFormat)
			if err != nil {
				return err
			}
			this.EffectFormat = &valueForEffectFormat
		}
	}
	
	if valThreeDFormat, ok := objMap["threeDFormat"]; ok {
		if valThreeDFormat != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormat, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			this.ThreeDFormat = &valueForThreeDFormat
		}
	}
	if valThreeDFormatCap, ok := objMap["ThreeDFormat"]; ok {
		if valThreeDFormatCap != nil {
			var valueForThreeDFormat ThreeDFormat
			err = json.Unmarshal(*valThreeDFormatCap, &valueForThreeDFormat)
			if err != nil {
				return err
			}
			this.ThreeDFormat = &valueForThreeDFormat
		}
	}
	
	if valLineFormat, ok := objMap["lineFormat"]; ok {
		if valLineFormat != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormat, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	if valLineFormatCap, ok := objMap["LineFormat"]; ok {
		if valLineFormatCap != nil {
			var valueForLineFormat LineFormat
			err = json.Unmarshal(*valLineFormatCap, &valueForLineFormat)
			if err != nil {
				return err
			}
			this.LineFormat = &valueForLineFormat
		}
	}
	this.Type_ = "Table"
	if valType, ok := objMap["type"]; ok {
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
	if valTypeCap, ok := objMap["Type"]; ok {
		if valTypeCap != nil {
			var valueForType string
			err = json.Unmarshal(*valTypeCap, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valTypeCap, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	this.Style = ""
	if valStyle, ok := objMap["style"]; ok {
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
	if valStyleCap, ok := objMap["Style"]; ok {
		if valStyleCap != nil {
			var valueForStyle string
			err = json.Unmarshal(*valStyleCap, &valueForStyle)
			if err != nil {
				var valueForStyleInt int32
				err = json.Unmarshal(*valStyleCap, &valueForStyleInt)
				if err != nil {
					return err
				}
				this.Style = string(valueForStyleInt)
			} else {
				this.Style = valueForStyle
			}
		}
	}
	
	if valRows, ok := objMap["rows"]; ok {
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
	if valRowsCap, ok := objMap["Rows"]; ok {
		if valRowsCap != nil {
			var valueForRows []json.RawMessage
			err = json.Unmarshal(*valRowsCap, &valueForRows)
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
	
	if valColumns, ok := objMap["columns"]; ok {
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
	if valColumnsCap, ok := objMap["Columns"]; ok {
		if valColumnsCap != nil {
			var valueForColumns []json.RawMessage
			err = json.Unmarshal(*valColumnsCap, &valueForColumns)
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
	
	if valFirstCol, ok := objMap["firstCol"]; ok {
		if valFirstCol != nil {
			var valueForFirstCol bool
			err = json.Unmarshal(*valFirstCol, &valueForFirstCol)
			if err != nil {
				return err
			}
			this.FirstCol = valueForFirstCol
		}
	}
	if valFirstColCap, ok := objMap["FirstCol"]; ok {
		if valFirstColCap != nil {
			var valueForFirstCol bool
			err = json.Unmarshal(*valFirstColCap, &valueForFirstCol)
			if err != nil {
				return err
			}
			this.FirstCol = valueForFirstCol
		}
	}
	
	if valFirstRow, ok := objMap["firstRow"]; ok {
		if valFirstRow != nil {
			var valueForFirstRow bool
			err = json.Unmarshal(*valFirstRow, &valueForFirstRow)
			if err != nil {
				return err
			}
			this.FirstRow = valueForFirstRow
		}
	}
	if valFirstRowCap, ok := objMap["FirstRow"]; ok {
		if valFirstRowCap != nil {
			var valueForFirstRow bool
			err = json.Unmarshal(*valFirstRowCap, &valueForFirstRow)
			if err != nil {
				return err
			}
			this.FirstRow = valueForFirstRow
		}
	}
	
	if valHorizontalBanding, ok := objMap["horizontalBanding"]; ok {
		if valHorizontalBanding != nil {
			var valueForHorizontalBanding bool
			err = json.Unmarshal(*valHorizontalBanding, &valueForHorizontalBanding)
			if err != nil {
				return err
			}
			this.HorizontalBanding = valueForHorizontalBanding
		}
	}
	if valHorizontalBandingCap, ok := objMap["HorizontalBanding"]; ok {
		if valHorizontalBandingCap != nil {
			var valueForHorizontalBanding bool
			err = json.Unmarshal(*valHorizontalBandingCap, &valueForHorizontalBanding)
			if err != nil {
				return err
			}
			this.HorizontalBanding = valueForHorizontalBanding
		}
	}
	
	if valLastCol, ok := objMap["lastCol"]; ok {
		if valLastCol != nil {
			var valueForLastCol bool
			err = json.Unmarshal(*valLastCol, &valueForLastCol)
			if err != nil {
				return err
			}
			this.LastCol = valueForLastCol
		}
	}
	if valLastColCap, ok := objMap["LastCol"]; ok {
		if valLastColCap != nil {
			var valueForLastCol bool
			err = json.Unmarshal(*valLastColCap, &valueForLastCol)
			if err != nil {
				return err
			}
			this.LastCol = valueForLastCol
		}
	}
	
	if valLastRow, ok := objMap["lastRow"]; ok {
		if valLastRow != nil {
			var valueForLastRow bool
			err = json.Unmarshal(*valLastRow, &valueForLastRow)
			if err != nil {
				return err
			}
			this.LastRow = valueForLastRow
		}
	}
	if valLastRowCap, ok := objMap["LastRow"]; ok {
		if valLastRowCap != nil {
			var valueForLastRow bool
			err = json.Unmarshal(*valLastRowCap, &valueForLastRow)
			if err != nil {
				return err
			}
			this.LastRow = valueForLastRow
		}
	}
	
	if valRightToLeft, ok := objMap["rightToLeft"]; ok {
		if valRightToLeft != nil {
			var valueForRightToLeft bool
			err = json.Unmarshal(*valRightToLeft, &valueForRightToLeft)
			if err != nil {
				return err
			}
			this.RightToLeft = valueForRightToLeft
		}
	}
	if valRightToLeftCap, ok := objMap["RightToLeft"]; ok {
		if valRightToLeftCap != nil {
			var valueForRightToLeft bool
			err = json.Unmarshal(*valRightToLeftCap, &valueForRightToLeft)
			if err != nil {
				return err
			}
			this.RightToLeft = valueForRightToLeft
		}
	}
	
	if valVerticalBanding, ok := objMap["verticalBanding"]; ok {
		if valVerticalBanding != nil {
			var valueForVerticalBanding bool
			err = json.Unmarshal(*valVerticalBanding, &valueForVerticalBanding)
			if err != nil {
				return err
			}
			this.VerticalBanding = valueForVerticalBanding
		}
	}
	if valVerticalBandingCap, ok := objMap["VerticalBanding"]; ok {
		if valVerticalBandingCap != nil {
			var valueForVerticalBanding bool
			err = json.Unmarshal(*valVerticalBandingCap, &valueForVerticalBanding)
			if err != nil {
				return err
			}
			this.VerticalBanding = valueForVerticalBanding
		}
	}

	return nil
}
