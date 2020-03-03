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
	"io/ioutil"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

var (
	testApiClient           *APIClient
	testFolderName          = "TempSlidesSDK"
	testFileName            = "test.ppt"
	unprotectedTestFileName = "test-unprotected.ppt"
	changedTestFileName     = "changedtest.ppt"
	testTemplateFileName    = "TemplateCV.pptx"
	testFilePassword        = "password"
	isInitialized		= false
	expectedTestDataVersion	= "1"
)

func initializeTest(functionName string, invalidParamName string, invalidParamValue interface{}) error {
	if !isInitialized {
		version := "0"
		c := getTestApiClient()
		var request DownloadFileRequest
		request.Path = "TempTests/version.txt"
		f, _, e := c.SlidesApi.DownloadFile(request)
		if e == nil {
			data, e := ioutil.ReadFile(f.Name())
			if e == nil {
				version = string(data)
			}
		}
		if version != expectedTestDataVersion {
			files, e := ioutil.ReadDir("TestData")
			if e != nil {
				return e
			}
			for _, file := range files {
				var request UploadFileRequest
				request.Path = "TempTests/" + file.Name()
				bytes, e := ioutil.ReadFile("TestData/" + file.Name())
				if e != nil {
					return e
				}
				request.File = bytes
				_, _, e = c.SlidesApi.UploadFile(request)
				if e != nil {
					return e
				}
			}
			var request UploadFileRequest
			request.Path = "TempTests/version.txt"
			request.File = []byte(expectedTestDataVersion)
			_, _, e = c.SlidesApi.UploadFile(request)
			if e != nil {
				return e
			}
		}
		isInitialized = true
	}
	files := make(map[string]FileRule)
	for _, rule := range getRules(getTestRules().Files, functionName, invalidParamName) {
		fileRule := rule.(FileRule)
		actualName := (untemplatize(fileRule.File, invalidParamValue)).(string)
		path := "TempSlidesSDK";
		if fileRule.Folder != "" {
			path = untemplatize(fileRule.Folder, invalidParamValue).(string)
		}
		path = path + "/" + actualName
		fileRule.ActualName = actualName
		files[path] = fileRule
	}

	for path, rule := range files {
		if rule.Action == "Put" {
			c := getTestApiClient()
			var request CopyFileRequest
			request.SrcPath = "TempTests/" + rule.ActualName
			request.DestPath = path
			_, e := c.SlidesApi.CopyFile(request)
			if e != nil {
				return e
			}
		} else if rule.Action == "Delete" {
			c := getTestApiClient()
			var request DeleteFileRequest
			request.Path = path
			_, e := c.SlidesApi.DeleteFile(request)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

var testRules *TestRules

func getTestRules() *TestRules {
	if testRules == nil {
		testRules = &TestRules{}
		file, err := os.Open("testRules.json")
		if err == nil {
			json.NewDecoder(file).Decode(testRules)
		}
	}
	return testRules
}

func getTestApiClient() *APIClient {
	cfg := NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient = NewAPIClient(cfg)
	return testApiClient
}

func createTestParamValue(functionName string, paramName string, paramType string) interface{} {
    if paramType == "[]byte" {
        data, _ := ioutil.ReadFile("testData/" + testFileName)
        return data
    }
    var value interface{}
    value = "test" + paramName
    for _, rule := range getRules(getTestRules().Values, functionName, paramName) {
        valueRule := rule.(ValueRule)
        if valueRule.ValueSet {
            value = valueRule.Value
        }
    }
    return undefaultize(value, paramType)
}

func undefaultize(value interface{}, paramType string) interface{} {
    if value == nil {
        if paramType == "[]int32" {
            return []int32{}
        }
        if paramType == "string" {
            return ""
        }
        if paramType == "IShapeExportOptions" {
            var options IShapeExportOptions
            return &options
        }
        if paramType == "ExportOptions" {
            var options ExportOptions
            return &options
        }
        return nil
    }
    if paramType == "[]byte" {
        return []byte{}
    }
    if paramType == "[]int32" {
        var arr = []int32{}
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &arr)
        return arr
    }
    if paramType == "int32" {
        return int32(value.(float64))
    }
    if paramType == "Paragraph" {
        var para Paragraph
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &para)
        return &para
    }
    if paramType == "Portion" {
        var portion Portion
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &portion)
        return &portion
    }
    if paramType == "ShapeBase" {
        var shape Shape
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &shape)
        return &shape
    }
    if paramType == "NotesSlide" {
        var slide NotesSlide
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "Slide" {
        var slide Slide
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "LayoutSlide" {
        var slide LayoutSlide
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "MasterSlide" {
        var slide MasterSlide
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "Pipeline" {
        var slide Pipeline
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "PresentationsMergeRequest" {
        var slide PresentationsMergeRequest
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "OrderedMergeRequest" {
        var slide OrderedMergeRequest
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "SlideBackground" {
        var slide SlideBackground
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "SlideAnimation" {
        var slide SlideAnimation
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "InteractiveSequence" {
        var slide InteractiveSequence
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        return &slide
    }
    if paramType == "Effect" {
        var slide Effect
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &slide)
        slide.Type_ = value.(map[string]interface{})["Type"].(string)
        return &slide
    }
    if paramType == "DocumentProperties" {
        var para DocumentProperties
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &para)
        return &para
    }
    if paramType == "DocumentProperty" {
        var para DocumentProperty
        b, _ := json.Marshal(value)
        json.Unmarshal(b, &para)
        return &para
    }
    return value
}

func getRules(rules interface{}, functionName string, paramName string) []ITestRule {
    filteredRules := []ITestRule{}
    ruleArray := reflect.ValueOf(rules)
    for i := 0; i < ruleArray.Len(); i++ {
        rule := ruleArray.Index(i).Interface().(ITestRule)
        if ruleApplies(rule, functionName, paramName) {
            filteredRules = append(filteredRules, rule)
        }
    }
    return filteredRules
}

func ruleApplies(rule ITestRule, functionName string, paramName string) bool {
    return (rule.getMethod() == "" || (functionName != "" && strings.ToLower(rule.getMethod()) == strings.ToLower(functionName))) &&
        (rule.getInvalid() == nil || (*(rule.getInvalid()) && (paramName != ""))) &&
        (rule.getParameter() == "" || (paramName != "" && strings.ToLower(rule.getParameter()) == strings.ToLower(paramName))) &&
        (rule.getLanguage() == "" || strings.ToLower(rule.getLanguage()) == "go")
}

func invalidizeTestParamValue(value interface{}, functionName string, paramName string, paramType string) interface{} {
    var invalidValue interface{}
    for _, rule := range getRules(getTestRules().Values, functionName, paramName) {
        valueRule := rule.(ValueRule)
        if valueRule.InvalidValueSet {
            invalidValue = valueRule.InvalidValue
        }
    }
    invalidValue = undefaultize(invalidValue, paramType)
    return untemplatize(invalidValue, value)
}

func assertError(t *testing.T, functionName string, paramName string, paramValue interface{}, errorCode int32, e error) {
        if (e == nil) {
		failed := true
		for _, _ = range getRules(getTestRules().OkToNotFail, functionName, paramName) {
		    failed = false
		}
		if (failed) {
		    t.Errorf("Must have failed.")
		    return
		}
	} else {
		var code int32
		message := "Unexpeceted message"
		for _, rule := range getRules(getTestRules().Results, functionName, paramName) {
			resultRule := rule.(ResultRule)
			if resultRule.Code != 0 {
				code = resultRule.Code
			}
			if resultRule.Message != "" {
				message = resultRule.Message
			}
		}
		if errorCode != code {
			t.Errorf("Unexpected error code: %v.", errorCode)
			return
		}
        	if (e != nil) {
			if !strings.Contains(e.Error(), untemplatize(message, paramValue).(string)) {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		}
	}
}

func assertBinaryResponse(file *os.File, t *testing.T) {
	if len(file.Name()) < 0 {
		t.Errorf("Empty or corrupt output file: %v.", file)
	}
}

func untemplatize(template interface{}, value interface{}) interface{} {
	if (template != nil && value != nil && reflect.TypeOf(template).Name() == "string") {
		if (reflect.TypeOf(value).Name() == "string") {
			return strings.Replace(template.(string), "%v", value.(string), -1)
		}
		if (reflect.TypeOf(value).Name() == "int32") {
			return strings.Replace(template.(string), "%v", fmt.Sprint(value), -1)
		}
	}
	return template
}
