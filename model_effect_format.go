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
	GetBlur() IBlurEffect
	SetBlur(newValue IBlurEffect)

	// glow effect
	GetGlow() IGlowEffect
	SetGlow(newValue IGlowEffect)

	// inner shadow effect
	GetInnerShadow() IInnerShadowEffect
	SetInnerShadow(newValue IInnerShadowEffect)

	// outer shadow effect
	GetOuterShadow() IOuterShadowEffect
	SetOuterShadow(newValue IOuterShadowEffect)

	// preset shadow effect
	GetPresetShadow() IPresetShadowEffect
	SetPresetShadow(newValue IPresetShadowEffect)

	// soft edge effect
	GetSoftEdge() ISoftEdgeEffect
	SetSoftEdge(newValue ISoftEdgeEffect)

	// reflection effect
	GetReflection() IReflectionEffect
	SetReflection(newValue IReflectionEffect)

	// fill overlay effect
	GetFillOverlay() IFillOverlayEffect
	SetFillOverlay(newValue IFillOverlayEffect)
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

func NewEffectFormat() *EffectFormat {
	instance := new(EffectFormat)
	return instance
}

func (this *EffectFormat) GetBlur() IBlurEffect {
	return this.Blur
}

func (this *EffectFormat) SetBlur(newValue IBlurEffect) {
	this.Blur = newValue
}
func (this *EffectFormat) GetGlow() IGlowEffect {
	return this.Glow
}

func (this *EffectFormat) SetGlow(newValue IGlowEffect) {
	this.Glow = newValue
}
func (this *EffectFormat) GetInnerShadow() IInnerShadowEffect {
	return this.InnerShadow
}

func (this *EffectFormat) SetInnerShadow(newValue IInnerShadowEffect) {
	this.InnerShadow = newValue
}
func (this *EffectFormat) GetOuterShadow() IOuterShadowEffect {
	return this.OuterShadow
}

func (this *EffectFormat) SetOuterShadow(newValue IOuterShadowEffect) {
	this.OuterShadow = newValue
}
func (this *EffectFormat) GetPresetShadow() IPresetShadowEffect {
	return this.PresetShadow
}

func (this *EffectFormat) SetPresetShadow(newValue IPresetShadowEffect) {
	this.PresetShadow = newValue
}
func (this *EffectFormat) GetSoftEdge() ISoftEdgeEffect {
	return this.SoftEdge
}

func (this *EffectFormat) SetSoftEdge(newValue ISoftEdgeEffect) {
	this.SoftEdge = newValue
}
func (this *EffectFormat) GetReflection() IReflectionEffect {
	return this.Reflection
}

func (this *EffectFormat) SetReflection(newValue IReflectionEffect) {
	this.Reflection = newValue
}
func (this *EffectFormat) GetFillOverlay() IFillOverlayEffect {
	return this.FillOverlay
}

func (this *EffectFormat) SetFillOverlay(newValue IFillOverlayEffect) {
	this.FillOverlay = newValue
}

func (this *EffectFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valBlur, ok := GetMapValue(objMap, "blur"); ok {
		if valBlur != nil {
			var valueForBlur BlurEffect
			err = json.Unmarshal(*valBlur, &valueForBlur)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("BlurEffect", *valBlur)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBlur, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IBlurEffect)
			if ok {
				this.Blur = vInterfaceObject
			}
		}
	}
	
	if valGlow, ok := GetMapValue(objMap, "glow"); ok {
		if valGlow != nil {
			var valueForGlow GlowEffect
			err = json.Unmarshal(*valGlow, &valueForGlow)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("GlowEffect", *valGlow)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valGlow, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IGlowEffect)
			if ok {
				this.Glow = vInterfaceObject
			}
		}
	}
	
	if valInnerShadow, ok := GetMapValue(objMap, "innerShadow"); ok {
		if valInnerShadow != nil {
			var valueForInnerShadow InnerShadowEffect
			err = json.Unmarshal(*valInnerShadow, &valueForInnerShadow)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("InnerShadowEffect", *valInnerShadow)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valInnerShadow, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IInnerShadowEffect)
			if ok {
				this.InnerShadow = vInterfaceObject
			}
		}
	}
	
	if valOuterShadow, ok := GetMapValue(objMap, "outerShadow"); ok {
		if valOuterShadow != nil {
			var valueForOuterShadow OuterShadowEffect
			err = json.Unmarshal(*valOuterShadow, &valueForOuterShadow)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("OuterShadowEffect", *valOuterShadow)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valOuterShadow, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IOuterShadowEffect)
			if ok {
				this.OuterShadow = vInterfaceObject
			}
		}
	}
	
	if valPresetShadow, ok := GetMapValue(objMap, "presetShadow"); ok {
		if valPresetShadow != nil {
			var valueForPresetShadow PresetShadowEffect
			err = json.Unmarshal(*valPresetShadow, &valueForPresetShadow)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("PresetShadowEffect", *valPresetShadow)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valPresetShadow, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IPresetShadowEffect)
			if ok {
				this.PresetShadow = vInterfaceObject
			}
		}
	}
	
	if valSoftEdge, ok := GetMapValue(objMap, "softEdge"); ok {
		if valSoftEdge != nil {
			var valueForSoftEdge SoftEdgeEffect
			err = json.Unmarshal(*valSoftEdge, &valueForSoftEdge)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("SoftEdgeEffect", *valSoftEdge)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valSoftEdge, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ISoftEdgeEffect)
			if ok {
				this.SoftEdge = vInterfaceObject
			}
		}
	}
	
	if valReflection, ok := GetMapValue(objMap, "reflection"); ok {
		if valReflection != nil {
			var valueForReflection ReflectionEffect
			err = json.Unmarshal(*valReflection, &valueForReflection)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ReflectionEffect", *valReflection)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valReflection, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IReflectionEffect)
			if ok {
				this.Reflection = vInterfaceObject
			}
		}
	}
	
	if valFillOverlay, ok := GetMapValue(objMap, "fillOverlay"); ok {
		if valFillOverlay != nil {
			var valueForFillOverlay FillOverlayEffect
			err = json.Unmarshal(*valFillOverlay, &valueForFillOverlay)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("FillOverlayEffect", *valFillOverlay)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valFillOverlay, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IFillOverlayEffect)
			if ok {
				this.FillOverlay = vInterfaceObject
			}
		}
	}

	return nil
}
