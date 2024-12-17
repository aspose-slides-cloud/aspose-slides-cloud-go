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

// Slides document properties.
type IViewProperties interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Last used view mode.
	GetLastView() string
	SetLastView(newValue string)

	// Horizontal bar state.
	GetHorizontalBarState() string
	SetHorizontalBarState(newValue string)

	// Vertical bar state.
	GetVerticalBarState() string
	SetVerticalBarState(newValue string)

	// True to prefer single view.
	GetPreferSingleView() *bool
	SetPreferSingleView(newValue *bool)

	// The sizing of the side content region of the normal view, when the region is of a variable restored size.
	GetRestoredLeft() INormalViewRestoredProperties
	SetRestoredLeft(newValue INormalViewRestoredProperties)

	// The sizing of the top slide region of the normal view, when the region is of a variable restored size.
	GetRestoredTop() INormalViewRestoredProperties
	SetRestoredTop(newValue INormalViewRestoredProperties)

	// Slide view mode properties.
	GetSlideViewProperties() ICommonSlideViewProperties
	SetSlideViewProperties(newValue ICommonSlideViewProperties)

	// Notes view mode properties.
	GetNotesViewProperties() ICommonSlideViewProperties
	SetNotesViewProperties(newValue ICommonSlideViewProperties)

	// True if the comments should be shown.
	GetShowComments() string
	SetShowComments(newValue string)

	// The grid spacing that should be used for the grid underlying the presentation document, in points.
	GetGridSpacing() float64
	SetGridSpacing(newValue float64)
}

type ViewProperties struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Last used view mode.
	LastView string `json:"LastView,omitempty"`

	// Horizontal bar state.
	HorizontalBarState string `json:"HorizontalBarState,omitempty"`

	// Vertical bar state.
	VerticalBarState string `json:"VerticalBarState,omitempty"`

	// True to prefer single view.
	PreferSingleView *bool `json:"PreferSingleView"`

	// The sizing of the side content region of the normal view, when the region is of a variable restored size.
	RestoredLeft INormalViewRestoredProperties `json:"RestoredLeft,omitempty"`

	// The sizing of the top slide region of the normal view, when the region is of a variable restored size.
	RestoredTop INormalViewRestoredProperties `json:"RestoredTop,omitempty"`

	// Slide view mode properties.
	SlideViewProperties ICommonSlideViewProperties `json:"SlideViewProperties,omitempty"`

	// Notes view mode properties.
	NotesViewProperties ICommonSlideViewProperties `json:"NotesViewProperties,omitempty"`

	// True if the comments should be shown.
	ShowComments string `json:"ShowComments,omitempty"`

	// The grid spacing that should be used for the grid underlying the presentation document, in points.
	GridSpacing float64 `json:"GridSpacing,omitempty"`
}

func NewViewProperties() *ViewProperties {
	instance := new(ViewProperties)
	return instance
}

func (this *ViewProperties) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ViewProperties) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ViewProperties) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ViewProperties) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ViewProperties) GetLastView() string {
	return this.LastView
}

func (this *ViewProperties) SetLastView(newValue string) {
	this.LastView = newValue
}
func (this *ViewProperties) GetHorizontalBarState() string {
	return this.HorizontalBarState
}

func (this *ViewProperties) SetHorizontalBarState(newValue string) {
	this.HorizontalBarState = newValue
}
func (this *ViewProperties) GetVerticalBarState() string {
	return this.VerticalBarState
}

func (this *ViewProperties) SetVerticalBarState(newValue string) {
	this.VerticalBarState = newValue
}
func (this *ViewProperties) GetPreferSingleView() *bool {
	return this.PreferSingleView
}

func (this *ViewProperties) SetPreferSingleView(newValue *bool) {
	this.PreferSingleView = newValue
}
func (this *ViewProperties) GetRestoredLeft() INormalViewRestoredProperties {
	return this.RestoredLeft
}

func (this *ViewProperties) SetRestoredLeft(newValue INormalViewRestoredProperties) {
	this.RestoredLeft = newValue
}
func (this *ViewProperties) GetRestoredTop() INormalViewRestoredProperties {
	return this.RestoredTop
}

func (this *ViewProperties) SetRestoredTop(newValue INormalViewRestoredProperties) {
	this.RestoredTop = newValue
}
func (this *ViewProperties) GetSlideViewProperties() ICommonSlideViewProperties {
	return this.SlideViewProperties
}

func (this *ViewProperties) SetSlideViewProperties(newValue ICommonSlideViewProperties) {
	this.SlideViewProperties = newValue
}
func (this *ViewProperties) GetNotesViewProperties() ICommonSlideViewProperties {
	return this.NotesViewProperties
}

func (this *ViewProperties) SetNotesViewProperties(newValue ICommonSlideViewProperties) {
	this.NotesViewProperties = newValue
}
func (this *ViewProperties) GetShowComments() string {
	return this.ShowComments
}

func (this *ViewProperties) SetShowComments(newValue string) {
	this.ShowComments = newValue
}
func (this *ViewProperties) GetGridSpacing() float64 {
	return this.GridSpacing
}

func (this *ViewProperties) SetGridSpacing(newValue float64) {
	this.GridSpacing = newValue
}

func (this *ViewProperties) UnmarshalJSON(b []byte) error {
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
	
	if valLastView, ok := GetMapValue(objMap, "lastView"); ok {
		if valLastView != nil {
			var valueForLastView string
			err = json.Unmarshal(*valLastView, &valueForLastView)
			if err != nil {
				var valueForLastViewInt int32
				err = json.Unmarshal(*valLastView, &valueForLastViewInt)
				if err != nil {
					return err
				}
				this.LastView = string(valueForLastViewInt)
			} else {
				this.LastView = valueForLastView
			}
		}
	}
	
	if valHorizontalBarState, ok := GetMapValue(objMap, "horizontalBarState"); ok {
		if valHorizontalBarState != nil {
			var valueForHorizontalBarState string
			err = json.Unmarshal(*valHorizontalBarState, &valueForHorizontalBarState)
			if err != nil {
				var valueForHorizontalBarStateInt int32
				err = json.Unmarshal(*valHorizontalBarState, &valueForHorizontalBarStateInt)
				if err != nil {
					return err
				}
				this.HorizontalBarState = string(valueForHorizontalBarStateInt)
			} else {
				this.HorizontalBarState = valueForHorizontalBarState
			}
		}
	}
	
	if valVerticalBarState, ok := GetMapValue(objMap, "verticalBarState"); ok {
		if valVerticalBarState != nil {
			var valueForVerticalBarState string
			err = json.Unmarshal(*valVerticalBarState, &valueForVerticalBarState)
			if err != nil {
				var valueForVerticalBarStateInt int32
				err = json.Unmarshal(*valVerticalBarState, &valueForVerticalBarStateInt)
				if err != nil {
					return err
				}
				this.VerticalBarState = string(valueForVerticalBarStateInt)
			} else {
				this.VerticalBarState = valueForVerticalBarState
			}
		}
	}
	
	if valPreferSingleView, ok := GetMapValue(objMap, "preferSingleView"); ok {
		if valPreferSingleView != nil {
			var valueForPreferSingleView *bool
			err = json.Unmarshal(*valPreferSingleView, &valueForPreferSingleView)
			if err != nil {
				return err
			}
			this.PreferSingleView = valueForPreferSingleView
		}
	}
	
	if valRestoredLeft, ok := GetMapValue(objMap, "restoredLeft"); ok {
		if valRestoredLeft != nil {
			var valueForRestoredLeft NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredLeft, &valueForRestoredLeft)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("NormalViewRestoredProperties", *valRestoredLeft)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valRestoredLeft, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(INormalViewRestoredProperties)
			if ok {
				this.RestoredLeft = vInterfaceObject
			}
		}
	}
	
	if valRestoredTop, ok := GetMapValue(objMap, "restoredTop"); ok {
		if valRestoredTop != nil {
			var valueForRestoredTop NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredTop, &valueForRestoredTop)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("NormalViewRestoredProperties", *valRestoredTop)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valRestoredTop, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(INormalViewRestoredProperties)
			if ok {
				this.RestoredTop = vInterfaceObject
			}
		}
	}
	
	if valSlideViewProperties, ok := GetMapValue(objMap, "slideViewProperties"); ok {
		if valSlideViewProperties != nil {
			var valueForSlideViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valSlideViewProperties, &valueForSlideViewProperties)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("CommonSlideViewProperties", *valSlideViewProperties)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlideViewProperties, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ICommonSlideViewProperties)
			if ok {
				this.SlideViewProperties = vInterfaceObject
			}
		}
	}
	
	if valNotesViewProperties, ok := GetMapValue(objMap, "notesViewProperties"); ok {
		if valNotesViewProperties != nil {
			var valueForNotesViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valNotesViewProperties, &valueForNotesViewProperties)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("CommonSlideViewProperties", *valNotesViewProperties)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNotesViewProperties, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ICommonSlideViewProperties)
			if ok {
				this.NotesViewProperties = vInterfaceObject
			}
		}
	}
	
	if valShowComments, ok := GetMapValue(objMap, "showComments"); ok {
		if valShowComments != nil {
			var valueForShowComments string
			err = json.Unmarshal(*valShowComments, &valueForShowComments)
			if err != nil {
				var valueForShowCommentsInt int32
				err = json.Unmarshal(*valShowComments, &valueForShowCommentsInt)
				if err != nil {
					return err
				}
				this.ShowComments = string(valueForShowCommentsInt)
			} else {
				this.ShowComments = valueForShowComments
			}
		}
	}
	
	if valGridSpacing, ok := GetMapValue(objMap, "gridSpacing"); ok {
		if valGridSpacing != nil {
			var valueForGridSpacing float64
			err = json.Unmarshal(*valGridSpacing, &valueForGridSpacing)
			if err != nil {
				return err
			}
			this.GridSpacing = valueForGridSpacing
		}
	}

	return nil
}
