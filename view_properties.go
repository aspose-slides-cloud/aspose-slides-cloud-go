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
	getSelfUri() IResourceUri
	setSelfUri(newValue IResourceUri)

	// List of alternate links.
	getAlternateLinks() []IResourceUri
	setAlternateLinks(newValue []IResourceUri)

	// Last used view mode.
	getLastView() string
	setLastView(newValue string)

	// Horizontal bar state.
	getHorizontalBarState() string
	setHorizontalBarState(newValue string)

	// Vertical bar state.
	getVerticalBarState() string
	setVerticalBarState(newValue string)

	// True to prefer single view.
	getPreferSingleView() bool
	setPreferSingleView(newValue bool)

	// The sizing of the side content region of the normal view, when the region is of a variable restored size.
	getRestoredLeft() INormalViewRestoredProperties
	setRestoredLeft(newValue INormalViewRestoredProperties)

	// The sizing of the top slide region of the normal view, when the region is of a variable restored size.
	getRestoredTop() INormalViewRestoredProperties
	setRestoredTop(newValue INormalViewRestoredProperties)

	// Slide view mode properties.
	getSlideViewProperties() ICommonSlideViewProperties
	setSlideViewProperties(newValue ICommonSlideViewProperties)

	// Notes view mode properties.
	getNotesViewProperties() ICommonSlideViewProperties
	setNotesViewProperties(newValue ICommonSlideViewProperties)

	// True if the comments should be shown.
	getShowComments() string
	setShowComments(newValue string)
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
	PreferSingleView bool `json:"PreferSingleView"`

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
}

func NewViewProperties() *ViewProperties {
	instance := new(ViewProperties)
	instance.LastView = ""
	instance.HorizontalBarState = ""
	instance.VerticalBarState = ""
	instance.ShowComments = ""
	return instance
}

func (this *ViewProperties) getSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *ViewProperties) setSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *ViewProperties) getAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *ViewProperties) setAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *ViewProperties) getLastView() string {
	return this.LastView
}

func (this *ViewProperties) setLastView(newValue string) {
	this.LastView = newValue
}
func (this *ViewProperties) getHorizontalBarState() string {
	return this.HorizontalBarState
}

func (this *ViewProperties) setHorizontalBarState(newValue string) {
	this.HorizontalBarState = newValue
}
func (this *ViewProperties) getVerticalBarState() string {
	return this.VerticalBarState
}

func (this *ViewProperties) setVerticalBarState(newValue string) {
	this.VerticalBarState = newValue
}
func (this *ViewProperties) getPreferSingleView() bool {
	return this.PreferSingleView
}

func (this *ViewProperties) setPreferSingleView(newValue bool) {
	this.PreferSingleView = newValue
}
func (this *ViewProperties) getRestoredLeft() INormalViewRestoredProperties {
	return this.RestoredLeft
}

func (this *ViewProperties) setRestoredLeft(newValue INormalViewRestoredProperties) {
	this.RestoredLeft = newValue
}
func (this *ViewProperties) getRestoredTop() INormalViewRestoredProperties {
	return this.RestoredTop
}

func (this *ViewProperties) setRestoredTop(newValue INormalViewRestoredProperties) {
	this.RestoredTop = newValue
}
func (this *ViewProperties) getSlideViewProperties() ICommonSlideViewProperties {
	return this.SlideViewProperties
}

func (this *ViewProperties) setSlideViewProperties(newValue ICommonSlideViewProperties) {
	this.SlideViewProperties = newValue
}
func (this *ViewProperties) getNotesViewProperties() ICommonSlideViewProperties {
	return this.NotesViewProperties
}

func (this *ViewProperties) setNotesViewProperties(newValue ICommonSlideViewProperties) {
	this.NotesViewProperties = newValue
}
func (this *ViewProperties) getShowComments() string {
	return this.ShowComments
}

func (this *ViewProperties) setShowComments(newValue string) {
	this.ShowComments = newValue
}

func (this *ViewProperties) UnmarshalJSON(b []byte) error {
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
	this.LastView = ""
	if valLastView, ok := objMap["lastView"]; ok {
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
	if valLastViewCap, ok := objMap["LastView"]; ok {
		if valLastViewCap != nil {
			var valueForLastView string
			err = json.Unmarshal(*valLastViewCap, &valueForLastView)
			if err != nil {
				var valueForLastViewInt int32
				err = json.Unmarshal(*valLastViewCap, &valueForLastViewInt)
				if err != nil {
					return err
				}
				this.LastView = string(valueForLastViewInt)
			} else {
				this.LastView = valueForLastView
			}
		}
	}
	this.HorizontalBarState = ""
	if valHorizontalBarState, ok := objMap["horizontalBarState"]; ok {
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
	if valHorizontalBarStateCap, ok := objMap["HorizontalBarState"]; ok {
		if valHorizontalBarStateCap != nil {
			var valueForHorizontalBarState string
			err = json.Unmarshal(*valHorizontalBarStateCap, &valueForHorizontalBarState)
			if err != nil {
				var valueForHorizontalBarStateInt int32
				err = json.Unmarshal(*valHorizontalBarStateCap, &valueForHorizontalBarStateInt)
				if err != nil {
					return err
				}
				this.HorizontalBarState = string(valueForHorizontalBarStateInt)
			} else {
				this.HorizontalBarState = valueForHorizontalBarState
			}
		}
	}
	this.VerticalBarState = ""
	if valVerticalBarState, ok := objMap["verticalBarState"]; ok {
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
	if valVerticalBarStateCap, ok := objMap["VerticalBarState"]; ok {
		if valVerticalBarStateCap != nil {
			var valueForVerticalBarState string
			err = json.Unmarshal(*valVerticalBarStateCap, &valueForVerticalBarState)
			if err != nil {
				var valueForVerticalBarStateInt int32
				err = json.Unmarshal(*valVerticalBarStateCap, &valueForVerticalBarStateInt)
				if err != nil {
					return err
				}
				this.VerticalBarState = string(valueForVerticalBarStateInt)
			} else {
				this.VerticalBarState = valueForVerticalBarState
			}
		}
	}
	
	if valPreferSingleView, ok := objMap["preferSingleView"]; ok {
		if valPreferSingleView != nil {
			var valueForPreferSingleView bool
			err = json.Unmarshal(*valPreferSingleView, &valueForPreferSingleView)
			if err != nil {
				return err
			}
			this.PreferSingleView = valueForPreferSingleView
		}
	}
	if valPreferSingleViewCap, ok := objMap["PreferSingleView"]; ok {
		if valPreferSingleViewCap != nil {
			var valueForPreferSingleView bool
			err = json.Unmarshal(*valPreferSingleViewCap, &valueForPreferSingleView)
			if err != nil {
				return err
			}
			this.PreferSingleView = valueForPreferSingleView
		}
	}
	
	if valRestoredLeft, ok := objMap["restoredLeft"]; ok {
		if valRestoredLeft != nil {
			var valueForRestoredLeft NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredLeft, &valueForRestoredLeft)
			if err != nil {
				return err
			}
			this.RestoredLeft = &valueForRestoredLeft
		}
	}
	if valRestoredLeftCap, ok := objMap["RestoredLeft"]; ok {
		if valRestoredLeftCap != nil {
			var valueForRestoredLeft NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredLeftCap, &valueForRestoredLeft)
			if err != nil {
				return err
			}
			this.RestoredLeft = &valueForRestoredLeft
		}
	}
	
	if valRestoredTop, ok := objMap["restoredTop"]; ok {
		if valRestoredTop != nil {
			var valueForRestoredTop NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredTop, &valueForRestoredTop)
			if err != nil {
				return err
			}
			this.RestoredTop = &valueForRestoredTop
		}
	}
	if valRestoredTopCap, ok := objMap["RestoredTop"]; ok {
		if valRestoredTopCap != nil {
			var valueForRestoredTop NormalViewRestoredProperties
			err = json.Unmarshal(*valRestoredTopCap, &valueForRestoredTop)
			if err != nil {
				return err
			}
			this.RestoredTop = &valueForRestoredTop
		}
	}
	
	if valSlideViewProperties, ok := objMap["slideViewProperties"]; ok {
		if valSlideViewProperties != nil {
			var valueForSlideViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valSlideViewProperties, &valueForSlideViewProperties)
			if err != nil {
				return err
			}
			this.SlideViewProperties = &valueForSlideViewProperties
		}
	}
	if valSlideViewPropertiesCap, ok := objMap["SlideViewProperties"]; ok {
		if valSlideViewPropertiesCap != nil {
			var valueForSlideViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valSlideViewPropertiesCap, &valueForSlideViewProperties)
			if err != nil {
				return err
			}
			this.SlideViewProperties = &valueForSlideViewProperties
		}
	}
	
	if valNotesViewProperties, ok := objMap["notesViewProperties"]; ok {
		if valNotesViewProperties != nil {
			var valueForNotesViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valNotesViewProperties, &valueForNotesViewProperties)
			if err != nil {
				return err
			}
			this.NotesViewProperties = &valueForNotesViewProperties
		}
	}
	if valNotesViewPropertiesCap, ok := objMap["NotesViewProperties"]; ok {
		if valNotesViewPropertiesCap != nil {
			var valueForNotesViewProperties CommonSlideViewProperties
			err = json.Unmarshal(*valNotesViewPropertiesCap, &valueForNotesViewProperties)
			if err != nil {
				return err
			}
			this.NotesViewProperties = &valueForNotesViewProperties
		}
	}
	this.ShowComments = ""
	if valShowComments, ok := objMap["showComments"]; ok {
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
	if valShowCommentsCap, ok := objMap["ShowComments"]; ok {
		if valShowCommentsCap != nil {
			var valueForShowComments string
			err = json.Unmarshal(*valShowCommentsCap, &valueForShowComments)
			if err != nil {
				var valueForShowCommentsInt int32
				err = json.Unmarshal(*valShowCommentsCap, &valueForShowCommentsInt)
				if err != nil {
					return err
				}
				this.ShowComments = string(valueForShowCommentsInt)
			} else {
				this.ShowComments = valueForShowComments
			}
		}
	}

	return nil
}
