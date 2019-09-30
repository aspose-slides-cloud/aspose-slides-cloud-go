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

// Class for disc space information.
type IDiscUsage interface {

	// Application used disc space.
	getUsedSize() int64
	setUsedSize(newValue int64)

	// Total disc space.
	getTotalSize() int64
	setTotalSize(newValue int64)
}

type DiscUsage struct {

	// Application used disc space.
	UsedSize int64 `json:"UsedSize"`

	// Total disc space.
	TotalSize int64 `json:"TotalSize"`
}

func (this DiscUsage) getUsedSize() int64 {
	return this.UsedSize
}

func (this DiscUsage) setUsedSize(newValue int64) {
	this.UsedSize = newValue
}
func (this DiscUsage) getTotalSize() int64 {
	return this.TotalSize
}

func (this DiscUsage) setTotalSize(newValue int64) {
	this.TotalSize = newValue
}

func (this *DiscUsage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valUsedSize, ok := objMap["usedSize"]; ok {
		if valUsedSize != nil {
			var valueForUsedSize int64
			err = json.Unmarshal(*valUsedSize, &valueForUsedSize)
			if err != nil {
				return err
			}
			this.UsedSize = valueForUsedSize
		}
	}
	if valUsedSizeCap, ok := objMap["UsedSize"]; ok {
		if valUsedSizeCap != nil {
			var valueForUsedSize int64
			err = json.Unmarshal(*valUsedSizeCap, &valueForUsedSize)
			if err != nil {
				return err
			}
			this.UsedSize = valueForUsedSize
		}
	}
	
	if valTotalSize, ok := objMap["totalSize"]; ok {
		if valTotalSize != nil {
			var valueForTotalSize int64
			err = json.Unmarshal(*valTotalSize, &valueForTotalSize)
			if err != nil {
				return err
			}
			this.TotalSize = valueForTotalSize
		}
	}
	if valTotalSizeCap, ok := objMap["TotalSize"]; ok {
		if valTotalSizeCap != nil {
			var valueForTotalSize int64
			err = json.Unmarshal(*valTotalSizeCap, &valueForTotalSize)
			if err != nil {
				return err
			}
			this.TotalSize = valueForTotalSize
		}
	}

    return nil
}
