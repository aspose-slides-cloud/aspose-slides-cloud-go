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

// Effect format
type IEffectFormat interface {

	// blur effect
	getBlur() IBlurEffect
	setBlur(newValue IBlurEffect)

	// glow effect
	getGlow() IGlowEffect
	setGlow(newValue IGlowEffect)

	// inner shadow effect
	getInnerShadow() IInnerShadowEffect
	setInnerShadow(newValue IInnerShadowEffect)

	// outer shadow effect
	getOuterShadow() IOuterShadowEffect
	setOuterShadow(newValue IOuterShadowEffect)

	// preset shadow effect
	getPresetShadow() IPresetShadowEffect
	setPresetShadow(newValue IPresetShadowEffect)

	// soft edge effect
	getSoftEdge() ISoftEdgeEffect
	setSoftEdge(newValue ISoftEdgeEffect)

	// reflection effect
	getReflection() IReflectionEffect
	setReflection(newValue IReflectionEffect)

	// fill overlay effect
	getFillOverlay() IFillOverlayEffect
	setFillOverlay(newValue IFillOverlayEffect)
}

type EffectFormat struct {

	// blur effect
	Blur IBlurEffect `json:"Blur,omitempty"`

	// glow effect
	Glow IGlowEffect `json:"Glow,omitempty"`

	// inner shadow effect
	InnerShadow IInnerShadowEffect `json:"InnerShadow,omitempty"`

	// outer shadow effect
	OuterShadow IOuterShadowEffect `json:"OuterShadow,omitempty"`

	// preset shadow effect
	PresetShadow IPresetShadowEffect `json:"PresetShadow,omitempty"`

	// soft edge effect
	SoftEdge ISoftEdgeEffect `json:"SoftEdge,omitempty"`

	// reflection effect
	Reflection IReflectionEffect `json:"Reflection,omitempty"`

	// fill overlay effect
	FillOverlay IFillOverlayEffect `json:"FillOverlay,omitempty"`
}

func (this EffectFormat) getBlur() IBlurEffect {
	return this.Blur
}

func (this EffectFormat) setBlur(newValue IBlurEffect) {
	this.Blur = newValue
}
func (this EffectFormat) getGlow() IGlowEffect {
	return this.Glow
}

func (this EffectFormat) setGlow(newValue IGlowEffect) {
	this.Glow = newValue
}
func (this EffectFormat) getInnerShadow() IInnerShadowEffect {
	return this.InnerShadow
}

func (this EffectFormat) setInnerShadow(newValue IInnerShadowEffect) {
	this.InnerShadow = newValue
}
func (this EffectFormat) getOuterShadow() IOuterShadowEffect {
	return this.OuterShadow
}

func (this EffectFormat) setOuterShadow(newValue IOuterShadowEffect) {
	this.OuterShadow = newValue
}
func (this EffectFormat) getPresetShadow() IPresetShadowEffect {
	return this.PresetShadow
}

func (this EffectFormat) setPresetShadow(newValue IPresetShadowEffect) {
	this.PresetShadow = newValue
}
func (this EffectFormat) getSoftEdge() ISoftEdgeEffect {
	return this.SoftEdge
}

func (this EffectFormat) setSoftEdge(newValue ISoftEdgeEffect) {
	this.SoftEdge = newValue
}
func (this EffectFormat) getReflection() IReflectionEffect {
	return this.Reflection
}

func (this EffectFormat) setReflection(newValue IReflectionEffect) {
	this.Reflection = newValue
}
func (this EffectFormat) getFillOverlay() IFillOverlayEffect {
	return this.FillOverlay
}

func (this EffectFormat) setFillOverlay(newValue IFillOverlayEffect) {
	this.FillOverlay = newValue
}

func (this *EffectFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if valBlur, ok := objMap["Blur"]; ok {
		if valBlur != nil {
			var valueForBlur BlurEffect
			err = json.Unmarshal(*valBlur, &valueForBlur)
			if err != nil {
				return err
			}
			this.Blur = valueForBlur
		}
	}

	if valGlow, ok := objMap["Glow"]; ok {
		if valGlow != nil {
			var valueForGlow GlowEffect
			err = json.Unmarshal(*valGlow, &valueForGlow)
			if err != nil {
				return err
			}
			this.Glow = valueForGlow
		}
	}

	if valInnerShadow, ok := objMap["InnerShadow"]; ok {
		if valInnerShadow != nil {
			var valueForInnerShadow InnerShadowEffect
			err = json.Unmarshal(*valInnerShadow, &valueForInnerShadow)
			if err != nil {
				return err
			}
			this.InnerShadow = valueForInnerShadow
		}
	}

	if valOuterShadow, ok := objMap["OuterShadow"]; ok {
		if valOuterShadow != nil {
			var valueForOuterShadow OuterShadowEffect
			err = json.Unmarshal(*valOuterShadow, &valueForOuterShadow)
			if err != nil {
				return err
			}
			this.OuterShadow = valueForOuterShadow
		}
	}

	if valPresetShadow, ok := objMap["PresetShadow"]; ok {
		if valPresetShadow != nil {
			var valueForPresetShadow PresetShadowEffect
			err = json.Unmarshal(*valPresetShadow, &valueForPresetShadow)
			if err != nil {
				return err
			}
			this.PresetShadow = valueForPresetShadow
		}
	}

	if valSoftEdge, ok := objMap["SoftEdge"]; ok {
		if valSoftEdge != nil {
			var valueForSoftEdge SoftEdgeEffect
			err = json.Unmarshal(*valSoftEdge, &valueForSoftEdge)
			if err != nil {
				return err
			}
			this.SoftEdge = valueForSoftEdge
		}
	}

	if valReflection, ok := objMap["Reflection"]; ok {
		if valReflection != nil {
			var valueForReflection ReflectionEffect
			err = json.Unmarshal(*valReflection, &valueForReflection)
			if err != nil {
				return err
			}
			this.Reflection = valueForReflection
		}
	}

	if valFillOverlay, ok := objMap["FillOverlay"]; ok {
		if valFillOverlay != nil {
			var valueForFillOverlay FillOverlayEffect
			err = json.Unmarshal(*valFillOverlay, &valueForFillOverlay)
			if err != nil {
				return err
			}
			this.FillOverlay = valueForFillOverlay
		}
	}

    return nil
}
