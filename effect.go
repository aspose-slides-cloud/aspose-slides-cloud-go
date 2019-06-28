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

// Represents comment of slide
type IEffect interface {

	// Effect type.
	getType() string
	setType(newValue string)

	// Effect subtype.
	getSubtype() string
	setSubtype(newValue string)

	// Preset class type.
	getPresetClassType() string
	setPresetClassType(newValue string)

	// Shape index.
	getShapeIndex() int32
	setShapeIndex(newValue int32)

	// Effect trigger type.
	getTriggerType() string
	setTriggerType(newValue string)

	// The percentage of duration accelerate behavior effect.
	getAccelerate() float64
	setAccelerate(newValue float64)

	// True to automatically play the animation in reverse after playing it in the forward direction.
	getAutoReverse() bool
	setAutoReverse(newValue bool)

	// The percentage of duration decelerate behavior effect.
	getDecelerate() float64
	setDecelerate(newValue float64)

	// The duration of animation effect.
	getDuration() float64
	setDuration(newValue float64)

	// The number of times the effect should repeat.
	getRepeatCount() float64
	setRepeatCount(newValue float64)

	// The number of times the effect should repeat.
	getRepeatDuration() float64
	setRepeatDuration(newValue float64)

	// The way for a effect to restart after complete.
	getRestart() string
	setRestart(newValue string)

	// The percentage by which to speed up (or slow down) the timing.
	getSpeed() float64
	setSpeed(newValue float64)

	// Delay time after trigger.
	getTriggerDelayTime() float64
	setTriggerDelayTime(newValue float64)
}

type Effect struct {

	// Effect type.
	Type_ string `json:"Type,omitempty"`

	// Effect subtype.
	Subtype string `json:"Subtype,omitempty"`

	// Preset class type.
	PresetClassType string `json:"PresetClassType,omitempty"`

	// Shape index.
	ShapeIndex int32 `json:"ShapeIndex"`

	// Effect trigger type.
	TriggerType string `json:"TriggerType,omitempty"`

	// The percentage of duration accelerate behavior effect.
	Accelerate float64 `json:"Accelerate,omitempty"`

	// True to automatically play the animation in reverse after playing it in the forward direction.
	AutoReverse bool `json:"AutoReverse,omitempty"`

	// The percentage of duration decelerate behavior effect.
	Decelerate float64 `json:"Decelerate,omitempty"`

	// The duration of animation effect.
	Duration float64 `json:"Duration,omitempty"`

	// The number of times the effect should repeat.
	RepeatCount float64 `json:"RepeatCount,omitempty"`

	// The number of times the effect should repeat.
	RepeatDuration float64 `json:"RepeatDuration,omitempty"`

	// The way for a effect to restart after complete.
	Restart string `json:"Restart,omitempty"`

	// The percentage by which to speed up (or slow down) the timing.
	Speed float64 `json:"Speed,omitempty"`

	// Delay time after trigger.
	TriggerDelayTime float64 `json:"TriggerDelayTime,omitempty"`
}

func (this Effect) getType() string {
	return this.Type_
}

func (this Effect) setType(newValue string) {
	this.Type_ = newValue
}
func (this Effect) getSubtype() string {
	return this.Subtype
}

func (this Effect) setSubtype(newValue string) {
	this.Subtype = newValue
}
func (this Effect) getPresetClassType() string {
	return this.PresetClassType
}

func (this Effect) setPresetClassType(newValue string) {
	this.PresetClassType = newValue
}
func (this Effect) getShapeIndex() int32 {
	return this.ShapeIndex
}

func (this Effect) setShapeIndex(newValue int32) {
	this.ShapeIndex = newValue
}
func (this Effect) getTriggerType() string {
	return this.TriggerType
}

func (this Effect) setTriggerType(newValue string) {
	this.TriggerType = newValue
}
func (this Effect) getAccelerate() float64 {
	return this.Accelerate
}

func (this Effect) setAccelerate(newValue float64) {
	this.Accelerate = newValue
}
func (this Effect) getAutoReverse() bool {
	return this.AutoReverse
}

func (this Effect) setAutoReverse(newValue bool) {
	this.AutoReverse = newValue
}
func (this Effect) getDecelerate() float64 {
	return this.Decelerate
}

func (this Effect) setDecelerate(newValue float64) {
	this.Decelerate = newValue
}
func (this Effect) getDuration() float64 {
	return this.Duration
}

func (this Effect) setDuration(newValue float64) {
	this.Duration = newValue
}
func (this Effect) getRepeatCount() float64 {
	return this.RepeatCount
}

func (this Effect) setRepeatCount(newValue float64) {
	this.RepeatCount = newValue
}
func (this Effect) getRepeatDuration() float64 {
	return this.RepeatDuration
}

func (this Effect) setRepeatDuration(newValue float64) {
	this.RepeatDuration = newValue
}
func (this Effect) getRestart() string {
	return this.Restart
}

func (this Effect) setRestart(newValue string) {
	this.Restart = newValue
}
func (this Effect) getSpeed() float64 {
	return this.Speed
}

func (this Effect) setSpeed(newValue float64) {
	this.Speed = newValue
}
func (this Effect) getTriggerDelayTime() float64 {
	return this.TriggerDelayTime
}

func (this Effect) setTriggerDelayTime(newValue float64) {
	this.TriggerDelayTime = newValue
}

func (this *Effect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = "null"
	if valType, ok := objMap["Type"]; ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				return err
			}
			this.Type_ = valueForType
		}
	}
	this.Subtype = "null"
	if valSubtype, ok := objMap["Subtype"]; ok {
		if valSubtype != nil {
			var valueForSubtype string
			err = json.Unmarshal(*valSubtype, &valueForSubtype)
			if err != nil {
				return err
			}
			this.Subtype = valueForSubtype
		}
	}
	this.PresetClassType = "null"
	if valPresetClassType, ok := objMap["PresetClassType"]; ok {
		if valPresetClassType != nil {
			var valueForPresetClassType string
			err = json.Unmarshal(*valPresetClassType, &valueForPresetClassType)
			if err != nil {
				return err
			}
			this.PresetClassType = valueForPresetClassType
		}
	}
	
	if valShapeIndex, ok := objMap["ShapeIndex"]; ok {
		if valShapeIndex != nil {
			var valueForShapeIndex int32
			err = json.Unmarshal(*valShapeIndex, &valueForShapeIndex)
			if err != nil {
				return err
			}
			this.ShapeIndex = valueForShapeIndex
		}
	}
	this.TriggerType = "null"
	if valTriggerType, ok := objMap["TriggerType"]; ok {
		if valTriggerType != nil {
			var valueForTriggerType string
			err = json.Unmarshal(*valTriggerType, &valueForTriggerType)
			if err != nil {
				return err
			}
			this.TriggerType = valueForTriggerType
		}
	}
	
	if valAccelerate, ok := objMap["Accelerate"]; ok {
		if valAccelerate != nil {
			var valueForAccelerate float64
			err = json.Unmarshal(*valAccelerate, &valueForAccelerate)
			if err != nil {
				return err
			}
			this.Accelerate = valueForAccelerate
		}
	}
	
	if valAutoReverse, ok := objMap["AutoReverse"]; ok {
		if valAutoReverse != nil {
			var valueForAutoReverse bool
			err = json.Unmarshal(*valAutoReverse, &valueForAutoReverse)
			if err != nil {
				return err
			}
			this.AutoReverse = valueForAutoReverse
		}
	}
	
	if valDecelerate, ok := objMap["Decelerate"]; ok {
		if valDecelerate != nil {
			var valueForDecelerate float64
			err = json.Unmarshal(*valDecelerate, &valueForDecelerate)
			if err != nil {
				return err
			}
			this.Decelerate = valueForDecelerate
		}
	}
	
	if valDuration, ok := objMap["Duration"]; ok {
		if valDuration != nil {
			var valueForDuration float64
			err = json.Unmarshal(*valDuration, &valueForDuration)
			if err != nil {
				return err
			}
			this.Duration = valueForDuration
		}
	}
	
	if valRepeatCount, ok := objMap["RepeatCount"]; ok {
		if valRepeatCount != nil {
			var valueForRepeatCount float64
			err = json.Unmarshal(*valRepeatCount, &valueForRepeatCount)
			if err != nil {
				return err
			}
			this.RepeatCount = valueForRepeatCount
		}
	}
	
	if valRepeatDuration, ok := objMap["RepeatDuration"]; ok {
		if valRepeatDuration != nil {
			var valueForRepeatDuration float64
			err = json.Unmarshal(*valRepeatDuration, &valueForRepeatDuration)
			if err != nil {
				return err
			}
			this.RepeatDuration = valueForRepeatDuration
		}
	}
	this.Restart = "null"
	if valRestart, ok := objMap["Restart"]; ok {
		if valRestart != nil {
			var valueForRestart string
			err = json.Unmarshal(*valRestart, &valueForRestart)
			if err != nil {
				return err
			}
			this.Restart = valueForRestart
		}
	}
	
	if valSpeed, ok := objMap["Speed"]; ok {
		if valSpeed != nil {
			var valueForSpeed float64
			err = json.Unmarshal(*valSpeed, &valueForSpeed)
			if err != nil {
				return err
			}
			this.Speed = valueForSpeed
		}
	}
	
	if valTriggerDelayTime, ok := objMap["TriggerDelayTime"]; ok {
		if valTriggerDelayTime != nil {
			var valueForTriggerDelayTime float64
			err = json.Unmarshal(*valTriggerDelayTime, &valueForTriggerDelayTime)
			if err != nil {
				return err
			}
			this.TriggerDelayTime = valueForTriggerDelayTime
		}
	}

    return nil
}
