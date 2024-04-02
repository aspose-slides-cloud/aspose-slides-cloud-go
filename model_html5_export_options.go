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

// Provides options that control how a presentation is saved in Html5 format.
type IHtml5ExportOptions interface {

	// Default regular font for rendering the presentation. 
	GetDefaultRegularFont() string
	SetDefaultRegularFont(newValue string)

	// Gets of sets list of font fallback rules.
	GetFontFallbackRules() []IFontFallbackRule
	SetFontFallbackRules(newValue []IFontFallbackRule)

	// Gets of sets list of font substitution rules.
	GetFontSubstRules() []IFontSubstRule
	SetFontSubstRules(newValue []IFontSubstRule)

	// Export format.
	GetFormat() string
	SetFormat(newValue string)

	// Gets or sets transitions animation option.
	GetAnimateTransitions() bool
	SetAnimateTransitions(newValue bool)

	// Gets or sets shapes animation option.
	GetAnimateShapes() bool
	SetAnimateShapes(newValue bool)

	// Gets or sets embed images option.
	GetEmbedImages() bool
	SetEmbedImages(newValue bool)

	// Slides layouting options
	GetNotesCommentsLayouting() INotesCommentsLayoutingOptions
	SetNotesCommentsLayouting(newValue INotesCommentsLayoutingOptions)
}

type Html5ExportOptions struct {

	// Default regular font for rendering the presentation. 
	DefaultRegularFont string `json:"DefaultRegularFont,omitempty"`

	// Gets of sets list of font fallback rules.
	FontFallbackRules []IFontFallbackRule `json:"FontFallbackRules,omitempty"`

	// Gets of sets list of font substitution rules.
	FontSubstRules []IFontSubstRule `json:"FontSubstRules,omitempty"`

	// Export format.
	Format string `json:"Format,omitempty"`

	// Gets or sets transitions animation option.
	AnimateTransitions bool `json:"AnimateTransitions"`

	// Gets or sets shapes animation option.
	AnimateShapes bool `json:"AnimateShapes"`

	// Gets or sets embed images option.
	EmbedImages bool `json:"EmbedImages"`

	// Slides layouting options
	NotesCommentsLayouting INotesCommentsLayoutingOptions `json:"NotesCommentsLayouting,omitempty"`
}

func NewHtml5ExportOptions() *Html5ExportOptions {
	instance := new(Html5ExportOptions)
	return instance
}

func (this *Html5ExportOptions) GetDefaultRegularFont() string {
	return this.DefaultRegularFont
}

func (this *Html5ExportOptions) SetDefaultRegularFont(newValue string) {
	this.DefaultRegularFont = newValue
}
func (this *Html5ExportOptions) GetFontFallbackRules() []IFontFallbackRule {
	return this.FontFallbackRules
}

func (this *Html5ExportOptions) SetFontFallbackRules(newValue []IFontFallbackRule) {
	this.FontFallbackRules = newValue
}
func (this *Html5ExportOptions) GetFontSubstRules() []IFontSubstRule {
	return this.FontSubstRules
}

func (this *Html5ExportOptions) SetFontSubstRules(newValue []IFontSubstRule) {
	this.FontSubstRules = newValue
}
func (this *Html5ExportOptions) GetFormat() string {
	return this.Format
}

func (this *Html5ExportOptions) SetFormat(newValue string) {
	this.Format = newValue
}
func (this *Html5ExportOptions) GetAnimateTransitions() bool {
	return this.AnimateTransitions
}

func (this *Html5ExportOptions) SetAnimateTransitions(newValue bool) {
	this.AnimateTransitions = newValue
}
func (this *Html5ExportOptions) GetAnimateShapes() bool {
	return this.AnimateShapes
}

func (this *Html5ExportOptions) SetAnimateShapes(newValue bool) {
	this.AnimateShapes = newValue
}
func (this *Html5ExportOptions) GetEmbedImages() bool {
	return this.EmbedImages
}

func (this *Html5ExportOptions) SetEmbedImages(newValue bool) {
	this.EmbedImages = newValue
}
func (this *Html5ExportOptions) GetNotesCommentsLayouting() INotesCommentsLayoutingOptions {
	return this.NotesCommentsLayouting
}

func (this *Html5ExportOptions) SetNotesCommentsLayouting(newValue INotesCommentsLayoutingOptions) {
	this.NotesCommentsLayouting = newValue
}

func (this *Html5ExportOptions) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valDefaultRegularFont, ok := objMap["defaultRegularFont"]; ok {
		if valDefaultRegularFont != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFont, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	if valDefaultRegularFontCap, ok := objMap["DefaultRegularFont"]; ok {
		if valDefaultRegularFontCap != nil {
			var valueForDefaultRegularFont string
			err = json.Unmarshal(*valDefaultRegularFontCap, &valueForDefaultRegularFont)
			if err != nil {
				return err
			}
			this.DefaultRegularFont = valueForDefaultRegularFont
		}
	}
	
	if valFontFallbackRules, ok := objMap["fontFallbackRules"]; ok {
		if valFontFallbackRules != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRules, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	if valFontFallbackRulesCap, ok := objMap["FontFallbackRules"]; ok {
		if valFontFallbackRulesCap != nil {
			var valueForFontFallbackRules []json.RawMessage
			err = json.Unmarshal(*valFontFallbackRulesCap, &valueForFontFallbackRules)
			if err != nil {
				return err
			}
			valueForIFontFallbackRules := make([]IFontFallbackRule, len(valueForFontFallbackRules))
			for i, v := range valueForFontFallbackRules {
				vObject, err := createObjectForType("FontFallbackRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontFallbackRules[i] = vObject.(IFontFallbackRule)
				}
			}
			this.FontFallbackRules = valueForIFontFallbackRules
		}
	}
	
	if valFontSubstRules, ok := objMap["fontSubstRules"]; ok {
		if valFontSubstRules != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRules, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	if valFontSubstRulesCap, ok := objMap["FontSubstRules"]; ok {
		if valFontSubstRulesCap != nil {
			var valueForFontSubstRules []json.RawMessage
			err = json.Unmarshal(*valFontSubstRulesCap, &valueForFontSubstRules)
			if err != nil {
				return err
			}
			valueForIFontSubstRules := make([]IFontSubstRule, len(valueForFontSubstRules))
			for i, v := range valueForFontSubstRules {
				vObject, err := createObjectForType("FontSubstRule", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIFontSubstRules[i] = vObject.(IFontSubstRule)
				}
			}
			this.FontSubstRules = valueForIFontSubstRules
		}
	}
	
	if valFormat, ok := objMap["format"]; ok {
		if valFormat != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormat, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	if valFormatCap, ok := objMap["Format"]; ok {
		if valFormatCap != nil {
			var valueForFormat string
			err = json.Unmarshal(*valFormatCap, &valueForFormat)
			if err != nil {
				return err
			}
			this.Format = valueForFormat
		}
	}
	
	if valAnimateTransitions, ok := objMap["animateTransitions"]; ok {
		if valAnimateTransitions != nil {
			var valueForAnimateTransitions bool
			err = json.Unmarshal(*valAnimateTransitions, &valueForAnimateTransitions)
			if err != nil {
				return err
			}
			this.AnimateTransitions = valueForAnimateTransitions
		}
	}
	if valAnimateTransitionsCap, ok := objMap["AnimateTransitions"]; ok {
		if valAnimateTransitionsCap != nil {
			var valueForAnimateTransitions bool
			err = json.Unmarshal(*valAnimateTransitionsCap, &valueForAnimateTransitions)
			if err != nil {
				return err
			}
			this.AnimateTransitions = valueForAnimateTransitions
		}
	}
	
	if valAnimateShapes, ok := objMap["animateShapes"]; ok {
		if valAnimateShapes != nil {
			var valueForAnimateShapes bool
			err = json.Unmarshal(*valAnimateShapes, &valueForAnimateShapes)
			if err != nil {
				return err
			}
			this.AnimateShapes = valueForAnimateShapes
		}
	}
	if valAnimateShapesCap, ok := objMap["AnimateShapes"]; ok {
		if valAnimateShapesCap != nil {
			var valueForAnimateShapes bool
			err = json.Unmarshal(*valAnimateShapesCap, &valueForAnimateShapes)
			if err != nil {
				return err
			}
			this.AnimateShapes = valueForAnimateShapes
		}
	}
	
	if valEmbedImages, ok := objMap["embedImages"]; ok {
		if valEmbedImages != nil {
			var valueForEmbedImages bool
			err = json.Unmarshal(*valEmbedImages, &valueForEmbedImages)
			if err != nil {
				return err
			}
			this.EmbedImages = valueForEmbedImages
		}
	}
	if valEmbedImagesCap, ok := objMap["EmbedImages"]; ok {
		if valEmbedImagesCap != nil {
			var valueForEmbedImages bool
			err = json.Unmarshal(*valEmbedImagesCap, &valueForEmbedImages)
			if err != nil {
				return err
			}
			this.EmbedImages = valueForEmbedImages
		}
	}
	
	if valNotesCommentsLayouting, ok := objMap["notesCommentsLayouting"]; ok {
		if valNotesCommentsLayouting != nil {
			var valueForNotesCommentsLayouting NotesCommentsLayoutingOptions
			err = json.Unmarshal(*valNotesCommentsLayouting, &valueForNotesCommentsLayouting)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("NotesCommentsLayoutingOptions", *valNotesCommentsLayouting)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNotesCommentsLayouting, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(INotesCommentsLayoutingOptions)
			if ok {
				this.NotesCommentsLayouting = vInterfaceObject
			}
		}
	}
	if valNotesCommentsLayoutingCap, ok := objMap["NotesCommentsLayouting"]; ok {
		if valNotesCommentsLayoutingCap != nil {
			var valueForNotesCommentsLayouting NotesCommentsLayoutingOptions
			err = json.Unmarshal(*valNotesCommentsLayoutingCap, &valueForNotesCommentsLayouting)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("NotesCommentsLayoutingOptions", *valNotesCommentsLayoutingCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valNotesCommentsLayoutingCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(INotesCommentsLayoutingOptions)
			if ok {
				this.NotesCommentsLayouting = vInterfaceObject
			}
		}
	}

	return nil
}
