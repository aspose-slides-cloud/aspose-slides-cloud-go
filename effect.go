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
	AutoReverse bool `json:"AutoReverse"`

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

func NewEffect() *Effect {
	instance := new(Effect)
	instance.Type_ = ""
	instance.Subtype = ""
	instance.PresetClassType = ""
	instance.TriggerType = ""
	instance.Restart = ""
	return instance
}

func (this *Effect) getType() string {
	return this.Type_
}

func (this *Effect) setType(newValue string) {
	this.Type_ = newValue
}
func (this *Effect) getSubtype() string {
	return this.Subtype
}

func (this *Effect) setSubtype(newValue string) {
	this.Subtype = newValue
}
func (this *Effect) getPresetClassType() string {
	return this.PresetClassType
}

func (this *Effect) setPresetClassType(newValue string) {
	this.PresetClassType = newValue
}
func (this *Effect) getShapeIndex() int32 {
	return this.ShapeIndex
}

func (this *Effect) setShapeIndex(newValue int32) {
	this.ShapeIndex = newValue
}
func (this *Effect) getTriggerType() string {
	return this.TriggerType
}

func (this *Effect) setTriggerType(newValue string) {
	this.TriggerType = newValue
}
func (this *Effect) getAccelerate() float64 {
	return this.Accelerate
}

func (this *Effect) setAccelerate(newValue float64) {
	this.Accelerate = newValue
}
func (this *Effect) getAutoReverse() bool {
	return this.AutoReverse
}

func (this *Effect) setAutoReverse(newValue bool) {
	this.AutoReverse = newValue
}
func (this *Effect) getDecelerate() float64 {
	return this.Decelerate
}

func (this *Effect) setDecelerate(newValue float64) {
	this.Decelerate = newValue
}
func (this *Effect) getDuration() float64 {
	return this.Duration
}

func (this *Effect) setDuration(newValue float64) {
	this.Duration = newValue
}
func (this *Effect) getRepeatCount() float64 {
	return this.RepeatCount
}

func (this *Effect) setRepeatCount(newValue float64) {
	this.RepeatCount = newValue
}
func (this *Effect) getRepeatDuration() float64 {
	return this.RepeatDuration
}

func (this *Effect) setRepeatDuration(newValue float64) {
	this.RepeatDuration = newValue
}
func (this *Effect) getRestart() string {
	return this.Restart
}

func (this *Effect) setRestart(newValue string) {
	this.Restart = newValue
}
func (this *Effect) getSpeed() float64 {
	return this.Speed
}

func (this *Effect) setSpeed(newValue float64) {
	this.Speed = newValue
}
func (this *Effect) getTriggerDelayTime() float64 {
	return this.TriggerDelayTime
}

func (this *Effect) setTriggerDelayTime(newValue float64) {
	this.TriggerDelayTime = newValue
}

func (this *Effect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	this.Type_ = ""
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
	this.Subtype = ""
	if valSubtype, ok := objMap["subtype"]; ok {
		if valSubtype != nil {
			var valueForSubtype string
			err = json.Unmarshal(*valSubtype, &valueForSubtype)
			if err != nil {
				var valueForSubtypeInt int32
				err = json.Unmarshal(*valSubtype, &valueForSubtypeInt)
				if err != nil {
					return err
				}
				this.Subtype = string(valueForSubtypeInt)
			} else {
				this.Subtype = valueForSubtype
			}
		}
	}
	if valSubtypeCap, ok := objMap["Subtype"]; ok {
		if valSubtypeCap != nil {
			var valueForSubtype string
			err = json.Unmarshal(*valSubtypeCap, &valueForSubtype)
			if err != nil {
				var valueForSubtypeInt int32
				err = json.Unmarshal(*valSubtypeCap, &valueForSubtypeInt)
				if err != nil {
					return err
				}
				this.Subtype = string(valueForSubtypeInt)
			} else {
				this.Subtype = valueForSubtype
			}
		}
	}
	this.PresetClassType = ""
	if valPresetClassType, ok := objMap["presetClassType"]; ok {
		if valPresetClassType != nil {
			var valueForPresetClassType string
			err = json.Unmarshal(*valPresetClassType, &valueForPresetClassType)
			if err != nil {
				var valueForPresetClassTypeInt int32
				err = json.Unmarshal(*valPresetClassType, &valueForPresetClassTypeInt)
				if err != nil {
					return err
				}
				this.PresetClassType = string(valueForPresetClassTypeInt)
			} else {
				this.PresetClassType = valueForPresetClassType
			}
		}
	}
	if valPresetClassTypeCap, ok := objMap["PresetClassType"]; ok {
		if valPresetClassTypeCap != nil {
			var valueForPresetClassType string
			err = json.Unmarshal(*valPresetClassTypeCap, &valueForPresetClassType)
			if err != nil {
				var valueForPresetClassTypeInt int32
				err = json.Unmarshal(*valPresetClassTypeCap, &valueForPresetClassTypeInt)
				if err != nil {
					return err
				}
				this.PresetClassType = string(valueForPresetClassTypeInt)
			} else {
				this.PresetClassType = valueForPresetClassType
			}
		}
	}
	
	if valShapeIndex, ok := objMap["shapeIndex"]; ok {
		if valShapeIndex != nil {
			var valueForShapeIndex int32
			err = json.Unmarshal(*valShapeIndex, &valueForShapeIndex)
			if err != nil {
				return err
			}
			this.ShapeIndex = valueForShapeIndex
		}
	}
	if valShapeIndexCap, ok := objMap["ShapeIndex"]; ok {
		if valShapeIndexCap != nil {
			var valueForShapeIndex int32
			err = json.Unmarshal(*valShapeIndexCap, &valueForShapeIndex)
			if err != nil {
				return err
			}
			this.ShapeIndex = valueForShapeIndex
		}
	}
	this.TriggerType = ""
	if valTriggerType, ok := objMap["triggerType"]; ok {
		if valTriggerType != nil {
			var valueForTriggerType string
			err = json.Unmarshal(*valTriggerType, &valueForTriggerType)
			if err != nil {
				var valueForTriggerTypeInt int32
				err = json.Unmarshal(*valTriggerType, &valueForTriggerTypeInt)
				if err != nil {
					return err
				}
				this.TriggerType = string(valueForTriggerTypeInt)
			} else {
				this.TriggerType = valueForTriggerType
			}
		}
	}
	if valTriggerTypeCap, ok := objMap["TriggerType"]; ok {
		if valTriggerTypeCap != nil {
			var valueForTriggerType string
			err = json.Unmarshal(*valTriggerTypeCap, &valueForTriggerType)
			if err != nil {
				var valueForTriggerTypeInt int32
				err = json.Unmarshal(*valTriggerTypeCap, &valueForTriggerTypeInt)
				if err != nil {
					return err
				}
				this.TriggerType = string(valueForTriggerTypeInt)
			} else {
				this.TriggerType = valueForTriggerType
			}
		}
	}
	
	if valAccelerate, ok := objMap["accelerate"]; ok {
		if valAccelerate != nil {
			var valueForAccelerate float64
			err = json.Unmarshal(*valAccelerate, &valueForAccelerate)
			if err != nil {
				return err
			}
			this.Accelerate = valueForAccelerate
		}
	}
	if valAccelerateCap, ok := objMap["Accelerate"]; ok {
		if valAccelerateCap != nil {
			var valueForAccelerate float64
			err = json.Unmarshal(*valAccelerateCap, &valueForAccelerate)
			if err != nil {
				return err
			}
			this.Accelerate = valueForAccelerate
		}
	}
	
	if valAutoReverse, ok := objMap["autoReverse"]; ok {
		if valAutoReverse != nil {
			var valueForAutoReverse bool
			err = json.Unmarshal(*valAutoReverse, &valueForAutoReverse)
			if err != nil {
				return err
			}
			this.AutoReverse = valueForAutoReverse
		}
	}
	if valAutoReverseCap, ok := objMap["AutoReverse"]; ok {
		if valAutoReverseCap != nil {
			var valueForAutoReverse bool
			err = json.Unmarshal(*valAutoReverseCap, &valueForAutoReverse)
			if err != nil {
				return err
			}
			this.AutoReverse = valueForAutoReverse
		}
	}
	
	if valDecelerate, ok := objMap["decelerate"]; ok {
		if valDecelerate != nil {
			var valueForDecelerate float64
			err = json.Unmarshal(*valDecelerate, &valueForDecelerate)
			if err != nil {
				return err
			}
			this.Decelerate = valueForDecelerate
		}
	}
	if valDecelerateCap, ok := objMap["Decelerate"]; ok {
		if valDecelerateCap != nil {
			var valueForDecelerate float64
			err = json.Unmarshal(*valDecelerateCap, &valueForDecelerate)
			if err != nil {
				return err
			}
			this.Decelerate = valueForDecelerate
		}
	}
	
	if valDuration, ok := objMap["duration"]; ok {
		if valDuration != nil {
			var valueForDuration float64
			err = json.Unmarshal(*valDuration, &valueForDuration)
			if err != nil {
				return err
			}
			this.Duration = valueForDuration
		}
	}
	if valDurationCap, ok := objMap["Duration"]; ok {
		if valDurationCap != nil {
			var valueForDuration float64
			err = json.Unmarshal(*valDurationCap, &valueForDuration)
			if err != nil {
				return err
			}
			this.Duration = valueForDuration
		}
	}
	
	if valRepeatCount, ok := objMap["repeatCount"]; ok {
		if valRepeatCount != nil {
			var valueForRepeatCount float64
			err = json.Unmarshal(*valRepeatCount, &valueForRepeatCount)
			if err != nil {
				return err
			}
			this.RepeatCount = valueForRepeatCount
		}
	}
	if valRepeatCountCap, ok := objMap["RepeatCount"]; ok {
		if valRepeatCountCap != nil {
			var valueForRepeatCount float64
			err = json.Unmarshal(*valRepeatCountCap, &valueForRepeatCount)
			if err != nil {
				return err
			}
			this.RepeatCount = valueForRepeatCount
		}
	}
	
	if valRepeatDuration, ok := objMap["repeatDuration"]; ok {
		if valRepeatDuration != nil {
			var valueForRepeatDuration float64
			err = json.Unmarshal(*valRepeatDuration, &valueForRepeatDuration)
			if err != nil {
				return err
			}
			this.RepeatDuration = valueForRepeatDuration
		}
	}
	if valRepeatDurationCap, ok := objMap["RepeatDuration"]; ok {
		if valRepeatDurationCap != nil {
			var valueForRepeatDuration float64
			err = json.Unmarshal(*valRepeatDurationCap, &valueForRepeatDuration)
			if err != nil {
				return err
			}
			this.RepeatDuration = valueForRepeatDuration
		}
	}
	this.Restart = ""
	if valRestart, ok := objMap["restart"]; ok {
		if valRestart != nil {
			var valueForRestart string
			err = json.Unmarshal(*valRestart, &valueForRestart)
			if err != nil {
				var valueForRestartInt int32
				err = json.Unmarshal(*valRestart, &valueForRestartInt)
				if err != nil {
					return err
				}
				this.Restart = string(valueForRestartInt)
			} else {
				this.Restart = valueForRestart
			}
		}
	}
	if valRestartCap, ok := objMap["Restart"]; ok {
		if valRestartCap != nil {
			var valueForRestart string
			err = json.Unmarshal(*valRestartCap, &valueForRestart)
			if err != nil {
				var valueForRestartInt int32
				err = json.Unmarshal(*valRestartCap, &valueForRestartInt)
				if err != nil {
					return err
				}
				this.Restart = string(valueForRestartInt)
			} else {
				this.Restart = valueForRestart
			}
		}
	}
	
	if valSpeed, ok := objMap["speed"]; ok {
		if valSpeed != nil {
			var valueForSpeed float64
			err = json.Unmarshal(*valSpeed, &valueForSpeed)
			if err != nil {
				return err
			}
			this.Speed = valueForSpeed
		}
	}
	if valSpeedCap, ok := objMap["Speed"]; ok {
		if valSpeedCap != nil {
			var valueForSpeed float64
			err = json.Unmarshal(*valSpeedCap, &valueForSpeed)
			if err != nil {
				return err
			}
			this.Speed = valueForSpeed
		}
	}
	
	if valTriggerDelayTime, ok := objMap["triggerDelayTime"]; ok {
		if valTriggerDelayTime != nil {
			var valueForTriggerDelayTime float64
			err = json.Unmarshal(*valTriggerDelayTime, &valueForTriggerDelayTime)
			if err != nil {
				return err
			}
			this.TriggerDelayTime = valueForTriggerDelayTime
		}
	}
	if valTriggerDelayTimeCap, ok := objMap["TriggerDelayTime"]; ok {
		if valTriggerDelayTimeCap != nil {
			var valueForTriggerDelayTime float64
			err = json.Unmarshal(*valTriggerDelayTimeCap, &valueForTriggerDelayTime)
			if err != nil {
				return err
			}
			this.TriggerDelayTime = valueForTriggerDelayTime
		}
	}

    return nil
}
