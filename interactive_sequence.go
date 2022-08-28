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

// Represents comments collection of slide
type IInteractiveSequence interface {

	// Effect list.
	GetEffects() []IEffect
	SetEffects(newValue []IEffect)

	// Index of the shape that triggers the sequence.
	GetTriggerShapeIndex() int32
	SetTriggerShapeIndex(newValue int32)
}

type InteractiveSequence struct {

	// Effect list.
	Effects []IEffect `json:"Effects,omitempty"`

	// Index of the shape that triggers the sequence.
	TriggerShapeIndex int32 `json:"TriggerShapeIndex"`
}

func NewInteractiveSequence() *InteractiveSequence {
	instance := new(InteractiveSequence)
	return instance
}

func (this *InteractiveSequence) GetEffects() []IEffect {
	return this.Effects
}

func (this *InteractiveSequence) SetEffects(newValue []IEffect) {
	this.Effects = newValue
}
func (this *InteractiveSequence) GetTriggerShapeIndex() int32 {
	return this.TriggerShapeIndex
}

func (this *InteractiveSequence) SetTriggerShapeIndex(newValue int32) {
	this.TriggerShapeIndex = newValue
}

func (this *InteractiveSequence) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valEffects, ok := objMap["effects"]; ok {
		if valEffects != nil {
			var valueForEffects []json.RawMessage
			err = json.Unmarshal(*valEffects, &valueForEffects)
			if err != nil {
				return err
			}
			valueForIEffects := make([]IEffect, len(valueForEffects))
			for i, v := range valueForEffects {
				vObject, err := createObjectForType("Effect", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIEffects[i] = vObject.(IEffect)
				}
			}
			this.Effects = valueForIEffects
		}
	}
	if valEffectsCap, ok := objMap["Effects"]; ok {
		if valEffectsCap != nil {
			var valueForEffects []json.RawMessage
			err = json.Unmarshal(*valEffectsCap, &valueForEffects)
			if err != nil {
				return err
			}
			valueForIEffects := make([]IEffect, len(valueForEffects))
			for i, v := range valueForEffects {
				vObject, err := createObjectForType("Effect", v)
				if err != nil {
					return err
				}
				err = json.Unmarshal(v, &vObject)
				if err != nil {
					return err
				}
				if vObject != nil {
					valueForIEffects[i] = vObject.(IEffect)
				}
			}
			this.Effects = valueForIEffects
		}
	}
	
	if valTriggerShapeIndex, ok := objMap["triggerShapeIndex"]; ok {
		if valTriggerShapeIndex != nil {
			var valueForTriggerShapeIndex int32
			err = json.Unmarshal(*valTriggerShapeIndex, &valueForTriggerShapeIndex)
			if err != nil {
				return err
			}
			this.TriggerShapeIndex = valueForTriggerShapeIndex
		}
	}
	if valTriggerShapeIndexCap, ok := objMap["TriggerShapeIndex"]; ok {
		if valTriggerShapeIndexCap != nil {
			var valueForTriggerShapeIndex int32
			err = json.Unmarshal(*valTriggerShapeIndexCap, &valueForTriggerShapeIndex)
			if err != nil {
				return err
			}
			this.TriggerShapeIndex = valueForTriggerShapeIndex
		}
	}

	return nil
}
