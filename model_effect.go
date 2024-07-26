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
	GetType() string
	SetType(newValue string)

	// Effect subtype.
	GetSubtype() string
	SetSubtype(newValue string)

	// Preset class type.
	GetPresetClassType() string
	SetPresetClassType(newValue string)

	// Preset class type.
	GetAnimateTextType() string
	SetAnimateTextType(newValue string)

	// Shape index.
	GetShapeIndex() int32
	SetShapeIndex(newValue int32)

	// Paragraph index.
	GetParagraphIndex() int32
	SetParagraphIndex(newValue int32)

	// Effect trigger type.
	GetTriggerType() string
	SetTriggerType(newValue string)

	// The percentage of duration accelerate behavior effect.
	GetAccelerate() float64
	SetAccelerate(newValue float64)

	// True to automatically play the animation in reverse after playing it in the forward direction.
	GetAutoReverse() *bool
	SetAutoReverse(newValue *bool)

	// The percentage of duration decelerate behavior effect.
	GetDecelerate() float64
	SetDecelerate(newValue float64)

	// The duration of animation effect.
	GetDuration() float64
	SetDuration(newValue float64)

	// The number of times the effect should repeat.
	GetRepeatCount() float64
	SetRepeatCount(newValue float64)

	// The number of times the effect should repeat.
	GetRepeatDuration() float64
	SetRepeatDuration(newValue float64)

	// The way for a effect to restart after complete.
	GetRestart() string
	SetRestart(newValue string)

	// The percentage by which to speed up (or slow down) the timing.
	GetSpeed() float64
	SetSpeed(newValue float64)

	// Delay time after trigger.
	GetTriggerDelayTime() float64
	SetTriggerDelayTime(newValue float64)

	// Specifies if the effect will repeat until the end of slide.
	GetRepeatUntilEndSlide() *bool
	SetRepeatUntilEndSlide(newValue *bool)

	// Specifies if the effect will repeat until the next click.
	GetRepeatUntilNextClick() *bool
	SetRepeatUntilNextClick(newValue *bool)

	// This attribute specifies if the animation effect stops the previous sound.
	GetStopPreviousSound() *bool
	SetStopPreviousSound(newValue *bool)

	// This attribute specifies if the effect will rewind when done playing.
	GetRewind() *bool
	SetRewind(newValue *bool)

	// Defined an after animation color for effect.
	GetAfterAnimationType() string
	SetAfterAnimationType(newValue string)

	// Defined an after animation color for effect. Applied when the AfterAnimationType property is set to Color.
	GetAfterAnimationColor() string
	SetAfterAnimationColor(newValue string)
}

type Effect struct {

	// Effect type.
	Type_ string `json:"Type,omitempty"`

	// Effect subtype.
	Subtype string `json:"Subtype,omitempty"`

	// Preset class type.
	PresetClassType string `json:"PresetClassType,omitempty"`

	// Preset class type.
	AnimateTextType string `json:"AnimateTextType,omitempty"`

	// Shape index.
	ShapeIndex int32 `json:"ShapeIndex"`

	// Paragraph index.
	ParagraphIndex int32 `json:"ParagraphIndex,omitempty"`

	// Effect trigger type.
	TriggerType string `json:"TriggerType,omitempty"`

	// The percentage of duration accelerate behavior effect.
	Accelerate float64 `json:"Accelerate,omitempty"`

	// True to automatically play the animation in reverse after playing it in the forward direction.
	AutoReverse *bool `json:"AutoReverse"`

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

	// Specifies if the effect will repeat until the end of slide.
	RepeatUntilEndSlide *bool `json:"RepeatUntilEndSlide"`

	// Specifies if the effect will repeat until the next click.
	RepeatUntilNextClick *bool `json:"RepeatUntilNextClick"`

	// This attribute specifies if the animation effect stops the previous sound.
	StopPreviousSound *bool `json:"StopPreviousSound"`

	// This attribute specifies if the effect will rewind when done playing.
	Rewind *bool `json:"Rewind"`

	// Defined an after animation color for effect.
	AfterAnimationType string `json:"AfterAnimationType,omitempty"`

	// Defined an after animation color for effect. Applied when the AfterAnimationType property is set to Color.
	AfterAnimationColor string `json:"AfterAnimationColor,omitempty"`
}

func NewEffect() *Effect {
	instance := new(Effect)
	return instance
}

func (this *Effect) GetType() string {
	return this.Type_
}

func (this *Effect) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *Effect) GetSubtype() string {
	return this.Subtype
}

func (this *Effect) SetSubtype(newValue string) {
	this.Subtype = newValue
}
func (this *Effect) GetPresetClassType() string {
	return this.PresetClassType
}

func (this *Effect) SetPresetClassType(newValue string) {
	this.PresetClassType = newValue
}
func (this *Effect) GetAnimateTextType() string {
	return this.AnimateTextType
}

func (this *Effect) SetAnimateTextType(newValue string) {
	this.AnimateTextType = newValue
}
func (this *Effect) GetShapeIndex() int32 {
	return this.ShapeIndex
}

func (this *Effect) SetShapeIndex(newValue int32) {
	this.ShapeIndex = newValue
}
func (this *Effect) GetParagraphIndex() int32 {
	return this.ParagraphIndex
}

func (this *Effect) SetParagraphIndex(newValue int32) {
	this.ParagraphIndex = newValue
}
func (this *Effect) GetTriggerType() string {
	return this.TriggerType
}

func (this *Effect) SetTriggerType(newValue string) {
	this.TriggerType = newValue
}
func (this *Effect) GetAccelerate() float64 {
	return this.Accelerate
}

func (this *Effect) SetAccelerate(newValue float64) {
	this.Accelerate = newValue
}
func (this *Effect) GetAutoReverse() *bool {
	return this.AutoReverse
}

func (this *Effect) SetAutoReverse(newValue *bool) {
	this.AutoReverse = newValue
}
func (this *Effect) GetDecelerate() float64 {
	return this.Decelerate
}

func (this *Effect) SetDecelerate(newValue float64) {
	this.Decelerate = newValue
}
func (this *Effect) GetDuration() float64 {
	return this.Duration
}

func (this *Effect) SetDuration(newValue float64) {
	this.Duration = newValue
}
func (this *Effect) GetRepeatCount() float64 {
	return this.RepeatCount
}

func (this *Effect) SetRepeatCount(newValue float64) {
	this.RepeatCount = newValue
}
func (this *Effect) GetRepeatDuration() float64 {
	return this.RepeatDuration
}

func (this *Effect) SetRepeatDuration(newValue float64) {
	this.RepeatDuration = newValue
}
func (this *Effect) GetRestart() string {
	return this.Restart
}

func (this *Effect) SetRestart(newValue string) {
	this.Restart = newValue
}
func (this *Effect) GetSpeed() float64 {
	return this.Speed
}

func (this *Effect) SetSpeed(newValue float64) {
	this.Speed = newValue
}
func (this *Effect) GetTriggerDelayTime() float64 {
	return this.TriggerDelayTime
}

func (this *Effect) SetTriggerDelayTime(newValue float64) {
	this.TriggerDelayTime = newValue
}
func (this *Effect) GetRepeatUntilEndSlide() *bool {
	return this.RepeatUntilEndSlide
}

func (this *Effect) SetRepeatUntilEndSlide(newValue *bool) {
	this.RepeatUntilEndSlide = newValue
}
func (this *Effect) GetRepeatUntilNextClick() *bool {
	return this.RepeatUntilNextClick
}

func (this *Effect) SetRepeatUntilNextClick(newValue *bool) {
	this.RepeatUntilNextClick = newValue
}
func (this *Effect) GetStopPreviousSound() *bool {
	return this.StopPreviousSound
}

func (this *Effect) SetStopPreviousSound(newValue *bool) {
	this.StopPreviousSound = newValue
}
func (this *Effect) GetRewind() *bool {
	return this.Rewind
}

func (this *Effect) SetRewind(newValue *bool) {
	this.Rewind = newValue
}
func (this *Effect) GetAfterAnimationType() string {
	return this.AfterAnimationType
}

func (this *Effect) SetAfterAnimationType(newValue string) {
	this.AfterAnimationType = newValue
}
func (this *Effect) GetAfterAnimationColor() string {
	return this.AfterAnimationColor
}

func (this *Effect) SetAfterAnimationColor(newValue string) {
	this.AfterAnimationColor = newValue
}

func (this *Effect) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valType, ok := GetMapValue(objMap, "type"); ok {
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
	
	if valSubtype, ok := GetMapValue(objMap, "subtype"); ok {
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
	
	if valPresetClassType, ok := GetMapValue(objMap, "presetClassType"); ok {
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
	
	if valAnimateTextType, ok := GetMapValue(objMap, "animateTextType"); ok {
		if valAnimateTextType != nil {
			var valueForAnimateTextType string
			err = json.Unmarshal(*valAnimateTextType, &valueForAnimateTextType)
			if err != nil {
				var valueForAnimateTextTypeInt int32
				err = json.Unmarshal(*valAnimateTextType, &valueForAnimateTextTypeInt)
				if err != nil {
					return err
				}
				this.AnimateTextType = string(valueForAnimateTextTypeInt)
			} else {
				this.AnimateTextType = valueForAnimateTextType
			}
		}
	}
	
	if valShapeIndex, ok := GetMapValue(objMap, "shapeIndex"); ok {
		if valShapeIndex != nil {
			var valueForShapeIndex int32
			err = json.Unmarshal(*valShapeIndex, &valueForShapeIndex)
			if err != nil {
				return err
			}
			this.ShapeIndex = valueForShapeIndex
		}
	}
	
	if valParagraphIndex, ok := GetMapValue(objMap, "paragraphIndex"); ok {
		if valParagraphIndex != nil {
			var valueForParagraphIndex int32
			err = json.Unmarshal(*valParagraphIndex, &valueForParagraphIndex)
			if err != nil {
				return err
			}
			this.ParagraphIndex = valueForParagraphIndex
		}
	}
	
	if valTriggerType, ok := GetMapValue(objMap, "triggerType"); ok {
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
	
	if valAccelerate, ok := GetMapValue(objMap, "accelerate"); ok {
		if valAccelerate != nil {
			var valueForAccelerate float64
			err = json.Unmarshal(*valAccelerate, &valueForAccelerate)
			if err != nil {
				return err
			}
			this.Accelerate = valueForAccelerate
		}
	}
	
	if valAutoReverse, ok := GetMapValue(objMap, "autoReverse"); ok {
		if valAutoReverse != nil {
			var valueForAutoReverse *bool
			err = json.Unmarshal(*valAutoReverse, &valueForAutoReverse)
			if err != nil {
				return err
			}
			this.AutoReverse = valueForAutoReverse
		}
	}
	
	if valDecelerate, ok := GetMapValue(objMap, "decelerate"); ok {
		if valDecelerate != nil {
			var valueForDecelerate float64
			err = json.Unmarshal(*valDecelerate, &valueForDecelerate)
			if err != nil {
				return err
			}
			this.Decelerate = valueForDecelerate
		}
	}
	
	if valDuration, ok := GetMapValue(objMap, "duration"); ok {
		if valDuration != nil {
			var valueForDuration float64
			err = json.Unmarshal(*valDuration, &valueForDuration)
			if err != nil {
				return err
			}
			this.Duration = valueForDuration
		}
	}
	
	if valRepeatCount, ok := GetMapValue(objMap, "repeatCount"); ok {
		if valRepeatCount != nil {
			var valueForRepeatCount float64
			err = json.Unmarshal(*valRepeatCount, &valueForRepeatCount)
			if err != nil {
				return err
			}
			this.RepeatCount = valueForRepeatCount
		}
	}
	
	if valRepeatDuration, ok := GetMapValue(objMap, "repeatDuration"); ok {
		if valRepeatDuration != nil {
			var valueForRepeatDuration float64
			err = json.Unmarshal(*valRepeatDuration, &valueForRepeatDuration)
			if err != nil {
				return err
			}
			this.RepeatDuration = valueForRepeatDuration
		}
	}
	
	if valRestart, ok := GetMapValue(objMap, "restart"); ok {
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
	
	if valSpeed, ok := GetMapValue(objMap, "speed"); ok {
		if valSpeed != nil {
			var valueForSpeed float64
			err = json.Unmarshal(*valSpeed, &valueForSpeed)
			if err != nil {
				return err
			}
			this.Speed = valueForSpeed
		}
	}
	
	if valTriggerDelayTime, ok := GetMapValue(objMap, "triggerDelayTime"); ok {
		if valTriggerDelayTime != nil {
			var valueForTriggerDelayTime float64
			err = json.Unmarshal(*valTriggerDelayTime, &valueForTriggerDelayTime)
			if err != nil {
				return err
			}
			this.TriggerDelayTime = valueForTriggerDelayTime
		}
	}
	
	if valRepeatUntilEndSlide, ok := GetMapValue(objMap, "repeatUntilEndSlide"); ok {
		if valRepeatUntilEndSlide != nil {
			var valueForRepeatUntilEndSlide *bool
			err = json.Unmarshal(*valRepeatUntilEndSlide, &valueForRepeatUntilEndSlide)
			if err != nil {
				return err
			}
			this.RepeatUntilEndSlide = valueForRepeatUntilEndSlide
		}
	}
	
	if valRepeatUntilNextClick, ok := GetMapValue(objMap, "repeatUntilNextClick"); ok {
		if valRepeatUntilNextClick != nil {
			var valueForRepeatUntilNextClick *bool
			err = json.Unmarshal(*valRepeatUntilNextClick, &valueForRepeatUntilNextClick)
			if err != nil {
				return err
			}
			this.RepeatUntilNextClick = valueForRepeatUntilNextClick
		}
	}
	
	if valStopPreviousSound, ok := GetMapValue(objMap, "stopPreviousSound"); ok {
		if valStopPreviousSound != nil {
			var valueForStopPreviousSound *bool
			err = json.Unmarshal(*valStopPreviousSound, &valueForStopPreviousSound)
			if err != nil {
				return err
			}
			this.StopPreviousSound = valueForStopPreviousSound
		}
	}
	
	if valRewind, ok := GetMapValue(objMap, "rewind"); ok {
		if valRewind != nil {
			var valueForRewind *bool
			err = json.Unmarshal(*valRewind, &valueForRewind)
			if err != nil {
				return err
			}
			this.Rewind = valueForRewind
		}
	}
	
	if valAfterAnimationType, ok := GetMapValue(objMap, "afterAnimationType"); ok {
		if valAfterAnimationType != nil {
			var valueForAfterAnimationType string
			err = json.Unmarshal(*valAfterAnimationType, &valueForAfterAnimationType)
			if err != nil {
				var valueForAfterAnimationTypeInt int32
				err = json.Unmarshal(*valAfterAnimationType, &valueForAfterAnimationTypeInt)
				if err != nil {
					return err
				}
				this.AfterAnimationType = string(valueForAfterAnimationTypeInt)
			} else {
				this.AfterAnimationType = valueForAfterAnimationType
			}
		}
	}
	
	if valAfterAnimationColor, ok := GetMapValue(objMap, "afterAnimationColor"); ok {
		if valAfterAnimationColor != nil {
			var valueForAfterAnimationColor string
			err = json.Unmarshal(*valAfterAnimationColor, &valueForAfterAnimationColor)
			if err != nil {
				return err
			}
			this.AfterAnimationColor = valueForAfterAnimationColor
		}
	}

	return nil
}
