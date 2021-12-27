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

// Represents response for ApiInfo  DTO
type IApiInfo interface {

	// Product name.
	getName() string
	setName(newValue string)

	// API version.
	getVersion() string
	setVersion(newValue string)
}

type ApiInfo struct {

	// Product name.
	Name string `json:"Name,omitempty"`

	// API version.
	Version string `json:"Version,omitempty"`
}

func NewApiInfo() *ApiInfo {
	instance := new(ApiInfo)
	return instance
}

func (this *ApiInfo) getName() string {
	return this.Name
}

func (this *ApiInfo) setName(newValue string) {
	this.Name = newValue
}
func (this *ApiInfo) getVersion() string {
	return this.Version
}

func (this *ApiInfo) setVersion(newValue string) {
	this.Version = newValue
}

func (this *ApiInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valName, ok := objMap["name"]; ok {
		if valName != nil {
			var valueForName string
			err = json.Unmarshal(*valName, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	if valNameCap, ok := objMap["Name"]; ok {
		if valNameCap != nil {
			var valueForName string
			err = json.Unmarshal(*valNameCap, &valueForName)
			if err != nil {
				return err
			}
			this.Name = valueForName
		}
	}
	
	if valVersion, ok := objMap["version"]; ok {
		if valVersion != nil {
			var valueForVersion string
			err = json.Unmarshal(*valVersion, &valueForVersion)
			if err != nil {
				return err
			}
			this.Version = valueForVersion
		}
	}
	if valVersionCap, ok := objMap["Version"]; ok {
		if valVersionCap != nil {
			var valueForVersion string
			err = json.Unmarshal(*valVersionCap, &valueForVersion)
			if err != nil {
				return err
			}
			this.Version = valueForVersion
		}
	}

	return nil
}
