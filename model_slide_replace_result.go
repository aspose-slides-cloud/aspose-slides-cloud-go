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

// Represents slide replace result DTO.
type ISlideReplaceResult interface {

	// Gets or sets the link to this resource.
	GetSelfUri() IResourceUri
	SetSelfUri(newValue IResourceUri)

	// List of alternate links.
	GetAlternateLinks() []IResourceUri
	SetAlternateLinks(newValue []IResourceUri)

	// Gets or sets the width.
	GetWidth() float64
	SetWidth(newValue float64)

	// Gets or sets the height.
	GetHeight() float64
	SetHeight(newValue float64)

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	GetShowMasterShapes() *bool
	SetShowMasterShapes(newValue *bool)

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	GetSlideShowTransition() ISlideShowTransition
	SetSlideShowTransition(newValue ISlideShowTransition)

	// Gets or sets the  link to the layout slide.
	GetLayoutSlide() IResourceUri
	SetLayoutSlide(newValue IResourceUri)

	// Gets or sets the  link to list of top-level shapes.
	GetShapes() IResourceUri
	SetShapes(newValue IResourceUri)

	// Gets or sets the link to theme.
	GetTheme() IResourceUri
	SetTheme(newValue IResourceUri)

	// Gets or sets the  link to placeholders.
	GetPlaceholders() IResourceUri
	SetPlaceholders(newValue IResourceUri)

	// Gets or sets the link to images.
	GetImages() IResourceUri
	SetImages(newValue IResourceUri)

	// Gets or sets the link to comments.
	GetComments() IResourceUri
	SetComments(newValue IResourceUri)

	// Get or sets the link to slide's background
	GetBackground() IResourceUri
	SetBackground(newValue IResourceUri)

	// Get or sets the link to notes slide.
	GetNotesSlide() IResourceUri
	SetNotesSlide(newValue IResourceUri)

	// Gets or sets the number of matches 
	GetMatches() int32
	SetMatches(newValue int32)
}

type SlideReplaceResult struct {

	// Gets or sets the link to this resource.
	SelfUri IResourceUri `json:"SelfUri,omitempty"`

	// List of alternate links.
	AlternateLinks []IResourceUri `json:"AlternateLinks,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width"`

	// Gets or sets the height.
	Height float64 `json:"Height"`

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	ShowMasterShapes *bool `json:"ShowMasterShapes"`

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	SlideShowTransition ISlideShowTransition `json:"SlideShowTransition,omitempty"`

	// Gets or sets the  link to the layout slide.
	LayoutSlide IResourceUri `json:"LayoutSlide,omitempty"`

	// Gets or sets the  link to list of top-level shapes.
	Shapes IResourceUri `json:"Shapes,omitempty"`

	// Gets or sets the link to theme.
	Theme IResourceUri `json:"Theme,omitempty"`

	// Gets or sets the  link to placeholders.
	Placeholders IResourceUri `json:"Placeholders,omitempty"`

	// Gets or sets the link to images.
	Images IResourceUri `json:"Images,omitempty"`

	// Gets or sets the link to comments.
	Comments IResourceUri `json:"Comments,omitempty"`

	// Get or sets the link to slide's background
	Background IResourceUri `json:"Background,omitempty"`

	// Get or sets the link to notes slide.
	NotesSlide IResourceUri `json:"NotesSlide,omitempty"`

	// Gets or sets the number of matches 
	Matches int32 `json:"Matches"`
}

func NewSlideReplaceResult() *SlideReplaceResult {
	instance := new(SlideReplaceResult)
	return instance
}

func (this *SlideReplaceResult) GetSelfUri() IResourceUri {
	return this.SelfUri
}

func (this *SlideReplaceResult) SetSelfUri(newValue IResourceUri) {
	this.SelfUri = newValue
}
func (this *SlideReplaceResult) GetAlternateLinks() []IResourceUri {
	return this.AlternateLinks
}

func (this *SlideReplaceResult) SetAlternateLinks(newValue []IResourceUri) {
	this.AlternateLinks = newValue
}
func (this *SlideReplaceResult) GetWidth() float64 {
	return this.Width
}

func (this *SlideReplaceResult) SetWidth(newValue float64) {
	this.Width = newValue
}
func (this *SlideReplaceResult) GetHeight() float64 {
	return this.Height
}

func (this *SlideReplaceResult) SetHeight(newValue float64) {
	this.Height = newValue
}
func (this *SlideReplaceResult) GetShowMasterShapes() *bool {
	return this.ShowMasterShapes
}

func (this *SlideReplaceResult) SetShowMasterShapes(newValue *bool) {
	this.ShowMasterShapes = newValue
}
func (this *SlideReplaceResult) GetSlideShowTransition() ISlideShowTransition {
	return this.SlideShowTransition
}

func (this *SlideReplaceResult) SetSlideShowTransition(newValue ISlideShowTransition) {
	this.SlideShowTransition = newValue
}
func (this *SlideReplaceResult) GetLayoutSlide() IResourceUri {
	return this.LayoutSlide
}

func (this *SlideReplaceResult) SetLayoutSlide(newValue IResourceUri) {
	this.LayoutSlide = newValue
}
func (this *SlideReplaceResult) GetShapes() IResourceUri {
	return this.Shapes
}

func (this *SlideReplaceResult) SetShapes(newValue IResourceUri) {
	this.Shapes = newValue
}
func (this *SlideReplaceResult) GetTheme() IResourceUri {
	return this.Theme
}

func (this *SlideReplaceResult) SetTheme(newValue IResourceUri) {
	this.Theme = newValue
}
func (this *SlideReplaceResult) GetPlaceholders() IResourceUri {
	return this.Placeholders
}

func (this *SlideReplaceResult) SetPlaceholders(newValue IResourceUri) {
	this.Placeholders = newValue
}
func (this *SlideReplaceResult) GetImages() IResourceUri {
	return this.Images
}

func (this *SlideReplaceResult) SetImages(newValue IResourceUri) {
	this.Images = newValue
}
func (this *SlideReplaceResult) GetComments() IResourceUri {
	return this.Comments
}

func (this *SlideReplaceResult) SetComments(newValue IResourceUri) {
	this.Comments = newValue
}
func (this *SlideReplaceResult) GetBackground() IResourceUri {
	return this.Background
}

func (this *SlideReplaceResult) SetBackground(newValue IResourceUri) {
	this.Background = newValue
}
func (this *SlideReplaceResult) GetNotesSlide() IResourceUri {
	return this.NotesSlide
}

func (this *SlideReplaceResult) SetNotesSlide(newValue IResourceUri) {
	this.NotesSlide = newValue
}
func (this *SlideReplaceResult) GetMatches() int32 {
	return this.Matches
}

func (this *SlideReplaceResult) SetMatches(newValue int32) {
	this.Matches = newValue
}

func (this *SlideReplaceResult) UnmarshalJSON(b []byte) error {
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
	if valSelfUriCap, ok := objMap["SelfUri"]; ok {
		if valSelfUriCap != nil {
			var valueForSelfUri ResourceUri
			err = json.Unmarshal(*valSelfUriCap, &valueForSelfUri)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valSelfUriCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSelfUriCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.SelfUri = vInterfaceObject
			}
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
	
	if valShowMasterShapes, ok := objMap["showMasterShapes"]; ok {
		if valShowMasterShapes != nil {
			var valueForShowMasterShapes *bool
			err = json.Unmarshal(*valShowMasterShapes, &valueForShowMasterShapes)
			if err != nil {
				return err
			}
			this.ShowMasterShapes = valueForShowMasterShapes
		}
	}
	if valShowMasterShapesCap, ok := objMap["ShowMasterShapes"]; ok {
		if valShowMasterShapesCap != nil {
			var valueForShowMasterShapes *bool
			err = json.Unmarshal(*valShowMasterShapesCap, &valueForShowMasterShapes)
			if err != nil {
				return err
			}
			this.ShowMasterShapes = valueForShowMasterShapes
		}
	}
	
	if valSlideShowTransition, ok := objMap["slideShowTransition"]; ok {
		if valSlideShowTransition != nil {
			var valueForSlideShowTransition SlideShowTransition
			err = json.Unmarshal(*valSlideShowTransition, &valueForSlideShowTransition)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SlideShowTransition", *valSlideShowTransition)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlideShowTransition, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISlideShowTransition)
			if ok {
				this.SlideShowTransition = vInterfaceObject
			}
		}
	}
	if valSlideShowTransitionCap, ok := objMap["SlideShowTransition"]; ok {
		if valSlideShowTransitionCap != nil {
			var valueForSlideShowTransition SlideShowTransition
			err = json.Unmarshal(*valSlideShowTransitionCap, &valueForSlideShowTransition)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SlideShowTransition", *valSlideShowTransitionCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSlideShowTransitionCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISlideShowTransition)
			if ok {
				this.SlideShowTransition = vInterfaceObject
			}
		}
	}
	
	if valLayoutSlide, ok := objMap["layoutSlide"]; ok {
		if valLayoutSlide != nil {
			var valueForLayoutSlide ResourceUri
			err = json.Unmarshal(*valLayoutSlide, &valueForLayoutSlide)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valLayoutSlide)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLayoutSlide, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.LayoutSlide = vInterfaceObject
			}
		}
	}
	if valLayoutSlideCap, ok := objMap["LayoutSlide"]; ok {
		if valLayoutSlideCap != nil {
			var valueForLayoutSlide ResourceUri
			err = json.Unmarshal(*valLayoutSlideCap, &valueForLayoutSlide)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valLayoutSlideCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLayoutSlideCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.LayoutSlide = vInterfaceObject
			}
		}
	}
	
	if valShapes, ok := objMap["shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valShapes)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShapes, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Shapes = vInterfaceObject
			}
		}
	}
	if valShapesCap, ok := objMap["Shapes"]; ok {
		if valShapesCap != nil {
			var valueForShapes ResourceUri
			err = json.Unmarshal(*valShapesCap, &valueForShapes)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valShapesCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valShapesCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Shapes = vInterfaceObject
			}
		}
	}
	
	if valTheme, ok := objMap["theme"]; ok {
		if valTheme != nil {
			var valueForTheme ResourceUri
			err = json.Unmarshal(*valTheme, &valueForTheme)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valTheme)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valTheme, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Theme = vInterfaceObject
			}
		}
	}
	if valThemeCap, ok := objMap["Theme"]; ok {
		if valThemeCap != nil {
			var valueForTheme ResourceUri
			err = json.Unmarshal(*valThemeCap, &valueForTheme)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valThemeCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valThemeCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Theme = vInterfaceObject
			}
		}
	}
	
	if valPlaceholders, ok := objMap["placeholders"]; ok {
		if valPlaceholders != nil {
			var valueForPlaceholders ResourceUri
			err = json.Unmarshal(*valPlaceholders, &valueForPlaceholders)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valPlaceholders)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPlaceholders, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Placeholders = vInterfaceObject
			}
		}
	}
	if valPlaceholdersCap, ok := objMap["Placeholders"]; ok {
		if valPlaceholdersCap != nil {
			var valueForPlaceholders ResourceUri
			err = json.Unmarshal(*valPlaceholdersCap, &valueForPlaceholders)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valPlaceholdersCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPlaceholdersCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Placeholders = vInterfaceObject
			}
		}
	}
	
	if valImages, ok := objMap["images"]; ok {
		if valImages != nil {
			var valueForImages ResourceUri
			err = json.Unmarshal(*valImages, &valueForImages)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valImages)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valImages, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Images = vInterfaceObject
			}
		}
	}
	if valImagesCap, ok := objMap["Images"]; ok {
		if valImagesCap != nil {
			var valueForImages ResourceUri
			err = json.Unmarshal(*valImagesCap, &valueForImages)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valImagesCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valImagesCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Images = vInterfaceObject
			}
		}
	}
	
	if valComments, ok := objMap["comments"]; ok {
		if valComments != nil {
			var valueForComments ResourceUri
			err = json.Unmarshal(*valComments, &valueForComments)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valComments)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valComments, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Comments = vInterfaceObject
			}
		}
	}
	if valCommentsCap, ok := objMap["Comments"]; ok {
		if valCommentsCap != nil {
			var valueForComments ResourceUri
			err = json.Unmarshal(*valCommentsCap, &valueForComments)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valCommentsCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCommentsCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Comments = vInterfaceObject
			}
		}
	}
	
	if valBackground, ok := objMap["background"]; ok {
		if valBackground != nil {
			var valueForBackground ResourceUri
			err = json.Unmarshal(*valBackground, &valueForBackground)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valBackground)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBackground, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Background = vInterfaceObject
			}
		}
	}
	if valBackgroundCap, ok := objMap["Background"]; ok {
		if valBackgroundCap != nil {
			var valueForBackground ResourceUri
			err = json.Unmarshal(*valBackgroundCap, &valueForBackground)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valBackgroundCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBackgroundCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.Background = vInterfaceObject
			}
		}
	}
	
	if valNotesSlide, ok := objMap["notesSlide"]; ok {
		if valNotesSlide != nil {
			var valueForNotesSlide ResourceUri
			err = json.Unmarshal(*valNotesSlide, &valueForNotesSlide)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valNotesSlide)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNotesSlide, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.NotesSlide = vInterfaceObject
			}
		}
	}
	if valNotesSlideCap, ok := objMap["NotesSlide"]; ok {
		if valNotesSlideCap != nil {
			var valueForNotesSlide ResourceUri
			err = json.Unmarshal(*valNotesSlideCap, &valueForNotesSlide)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ResourceUri", *valNotesSlideCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNotesSlideCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IResourceUri)
			if ok {
				this.NotesSlide = vInterfaceObject
			}
		}
	}
	
	if valMatches, ok := objMap["matches"]; ok {
		if valMatches != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatches, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}
	if valMatchesCap, ok := objMap["Matches"]; ok {
		if valMatchesCap != nil {
			var valueForMatches int32
			err = json.Unmarshal(*valMatchesCap, &valueForMatches)
			if err != nil {
				return err
			}
			this.Matches = valueForMatches
		}
	}

	return nil
}
