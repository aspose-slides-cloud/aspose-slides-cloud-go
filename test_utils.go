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
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	testApiClient *APIClient
	testFolderName = "TempSlidesSDK"
	testFileName = "test.ppt"
	unprotectedTestFileName = "test-unprotected.ppt"
	changedTestFileName = "changedtest.ppt"
	testTemplateFileName = "TemplateCV.pptx"
        testFilePassword = "password"
)

func initializeTest(functionName string, invalidParamName string, invalidParamValue interface{}) error {
	err := callStorage(testFolderName + "/" + testFileName, "PUT", testFileName)
	if err != nil {
		return err
	}
	if functionName == "DeleteSlidesCleanSlidesList" || functionName == "PutSlidesSlide" {
		err = callStorage(testFolderName + "/" + unprotectedTestFileName, "PUT", unprotectedTestFileName)
		if err != nil {
			return err
		}
	}
	if functionName == "PostSlidesDocument" {
		err = callStorage(testFolderName + "/" + testTemplateFileName, "PUT", testTemplateFileName)
		if err != nil {
			return err
		}
		err = callStorage(testFolderName + "/" + testFileName, "DELETE", "")
		if err != nil {
			return err
		}
	}
	if invalidParamName == "name" {
		err = callStorage(testFolderName + "/" + invalidParamValue.(string), "DELETE", "")
		if err != nil {
			return err
		}
	}
	if invalidParamName == "folder" {
		err = callStorage(invalidParamValue.(string) + "/" + testFileName, "DELETE", "")
		if err != nil {
			return err
		}
	}
	if functionName == "PostSlidesDocument" || functionName == "PutNewPresentation" {
		err = callStorage(testFolderName + "invalid/" + changedTestFileName, "DELETE", "")
		if err != nil {
			return err
		}
		err = callStorage(testFolderName + "/invalid" + changedTestFileName, "DELETE", "")
		if err != nil {
			return err
		}
		err = callStorage(testFolderName + "/" + changedTestFileName, "DELETE", "")
		if err != nil {
			return err
		}
	}
	return nil
}

func callStorage(path string, method string, testFileName string) error {
	cl := getTestApiClient()
	var body interface{}
	var err error
	if len(testFileName) > 0 {
		body, err = os.Open("testData/" + testFileName)
		if err != nil {
			return err
		}
	}
	r, err := cl.prepareRequest(
		nil, cl.cfg.GetApiUrl() + "/storage/file/" + path, method, body, make(map[string]string), nil, nil, "", nil)
	if err != nil {
		return err
	}
	rs, err := cl.callAPI(r)
	if err != nil {
		return err
	}
	if rs.StatusCode >= 300 {
		return errors.New("Storage error.")
	}
	return nil
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
	if paramType == "IShapeExportOptions" {
		var value IShapeExportOptions
		return value
	}
	if paramType == "SlideBackground" {
		var value SlideBackground
		return value
	}
	if paramType == "ExportOptions" || (functionName == "PostSlideSaveAs" && paramName == "options") {
		var value ExportOptions
		return value
	}
	if paramType == "Pipeline" {
		var value Pipeline
		return value
	}
	if paramType == "PresentationsMergeRequest" {
		var value PresentationsMergeRequest
		return value
	}
	if paramType == "OrderedMergeRequest" {
		var value OrderedMergeRequest
		return value
	}
	if paramType == "Slide" {
		var slide Slide
		var masterSlideUri ResourceUriElement
		var uri ResourceUri
		uri.Href = "TitleOnly"
		masterSlideUri.Uri = &uri
		slide.LayoutSlide = &masterSlideUri
		return slide
	}
	if paramType == "LayoutSlide" {
		var slide LayoutSlide
		var masterSlideUri ResourceUriElement
		var uri ResourceUri
		uri.Href = "masterSlides/2"
		masterSlideUri.Uri = &uri
		slide.MasterSlide = &masterSlideUri
		return slide
	}
	if paramType == "NotesSlide" {
		var slide NotesSlide
		slide.Text = "testNote"
		return slide
	}
	if paramType == "Paragraph" {
		var paragraph Paragraph
		return paragraph
	}
	if paramType == "Portion" {
		var portion Portion
		portion.Text = "testPortion"
		return portion
	}
	if paramType == "DocumentProperty" {
		var property DocumentProperty
		property.Name = "testProperty001"
		property.Value = "testValue002"
		return property
	}
	if paramType == "DocumentProperties" {
		var documentProperties DocumentProperties
		var property DocumentProperty
		property.Name = "testProperty001"
		property.Value = "testValue002"
		documentProperties.List = []DocumentProperty{ property }
		return documentProperties
	}
	if paramName == "dto" {
		var shape Shape
		shape.Text = "testShape"
		shape.Type_ = ShapeType_SHAPE
		shape.ShapeType = CombinedShapeType_BENT_ARROW
		shape.GeometryShapeType = GeometryShapeType_RECTANGLE
                return shape
	}
	if paramType == "[]byte" {
		data, _ := ioutil.ReadFile("testData/" + testFileName)
		return data
	}
	if paramType == "[]int32" {
		value := []int32{}
		return value
	}
	if paramType == "bool" {
		return true
	}
	if paramType == "float64" {
		var value float64
		return value
	}
	if paramType == "int32" {
		var value int32
		if paramName == "shapeIndex" {
			value = 3
		} else if paramName == "shapeToClone" {
			value = 0
		} else {
			value = 1
		}
		return value
	}
	if paramName == "format" {
		if functionName == "GetNotesSlideShapeWithFormat" ||
			functionName == "GetNotesSlideWithFormat" ||
			functionName == "GetSlidesImageWithFormat" ||
			functionName == "GetShapeWithFormat" ||
			functionName == "PostShapeSaveAs" ||
			functionName == "PostNotesSlideShapeSaveAs" {
			return "png"
		}
		return "pdf"
	}
	if paramName == "name" {
		if functionName == "PutNewPresentation" {
			return changedTestFileName
		}
		if (functionName == "DeleteSlidesCleanSlidesList" || functionName == "PutSlidesSlide") {
			return unprotectedTestFileName
		}
		return testFileName
	}
	if paramName == "propertyName" {
		return "testProperty"
	}
	if paramName == "folder" {
		return testFolderName
	}
        if (paramName == "templatePath" || paramName == "cloneFrom") {
            if (functionName == "PostSlidesDocument") {
                return testFolderName + "/" + testTemplateFileName
            }
            return testFolderName + "/" + testFileName
        }
	if strings.HasSuffix(strings.ToLower(paramName), "password") {
		if (functionName == "DeleteSlidesCleanSlidesList" || functionName == "PutSlidesSlide") {
			return ""
		}
		return testFilePassword
	}
        if (paramName == "data") {
            return "<staff><person><name>John Doe</name><address><line1>10 Downing Street</line1><line2>London</line2></address><phone>+457 123456</phone><bio>Hi, I'm John and this is my CV</bio><skills><skill><title>C#</title><level>Excellent</level></skill><skill><title>C++</title><level>Good</level></skill><skill><title>Java</title><level>Average</level></skill></skills></person></staff>"
        }
	if strings.HasSuffix(strings.ToLower(paramName), "storage") || paramName == "outPath" || paramName == "path" {
		return ""
	}
	return "test" + paramName
}

func invalidizeTestParamValue(value interface{}, paramName string, paramType string) interface{} {
	if paramType == "bool" {
		return false
	}
	if paramType == "float64" {
		var value float64
		value = 536
		return value
	}
	if paramType == "int32" {
		var value int32
		value = 536
		return value
	}
	if paramType == "[]int32" {
		value := []int32{ 1, 536 }
		return value
	}
	if paramType == "string" {
		if paramName == "name" {
			return "invalid" + value.(string)
		}
		return value.(string) + "invalid"
	}
	return value
}

func assertError(t *testing.T, functionName string, paramName string, errorCode int32, e error) {
	if errorCode < 300 {
		if paramName != "jpegQuality" &&
			paramName != "width" &&
			paramName != "height" &&
			paramName != "from" &&
			paramName != "sizeType" &&
			paramName != "scaleType" &&
			paramName != "scaleX" &&
			paramName != "scaleY" &&
			paramName != "bounds" &&
			paramName != "withEmpty" &&
			paramName != "ignoreCase" &&
			paramName != "oldValue" &&
			paramName != "newValue" &&
			paramName != "outPath" &&
			paramName != "pipeline" &&
			paramName != "options" &&
			paramName != "stream" &&
			paramName != "document" &&
			paramName != "html" &&
			paramName != "request" &&
			paramName != "dto" &&
			paramName != "slideDto" &&
			paramName != "background" &&
			paramName != "color" &&
			paramName != "properties" &&
			paramName != "propertyName" &&
			paramName != "property" &&
			paramName != "applyToAll" &&
			paramName != "destFolder" &&
			paramName != "fontsFolder" &&
			paramName != "isImageDataEmbedded" &&
			!(functionName == "PostSlidesReorderPosition" &&
				(paramName == "slideToCopy" ||
					paramName == "position" ||
					paramName == "slideToClone" ||
					paramName == "source" ||
					paramName == "layoutAlias")) &&
			!((functionName == "PutSlidesDocumentFromHtml" ||
					functionName == "PutNewPresentation" ||
					functionName == "PostSlidesDocument") &&
				(paramName == "name" || paramName == "folder" || paramName == "password")) {
			t.Errorf("Must have failed.")
		}
		return
	}
	if (paramName == "name" || paramName == "folder" || paramName == "cloneFrom" || paramName == "propertyName") &&
		(functionName != "PostAddNotesSlide") {
		if errorCode != 404 {
			t.Errorf("Unexpected error code: %v.", errorCode)
			return
		}
		if paramName == "propertyName" {
			if !strings.Contains(e.Error(), " not found.") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else {
			if !strings.Contains(e.Error(), "The specified key does not exist") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		}
        } else if paramName == "path" && strings.HasSuffix(functionName, "AddNewShape") {
		if errorCode != 405 {
			t.Errorf("Unexpected error code: %v.", errorCode)
			return
		}
	} else {
		if errorCode != 400 {
			t.Errorf("Unexpected error code: %v.", errorCode)
			return
		}
		if functionName == "PutNewPresentation" ||
			functionName == "PutSlidesDocumentFromHtml" ||
			(functionName == "PostAddNotesSlide" && paramName != "slideIndex") {
			if !strings.Contains(e.Error(), "Object reference not set to an instance of an object") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if strings.HasSuffix(strings.ToLower(paramName), "password") {
			if functionName == "DeleteSlidesCleanSlidesList" || functionName == "PutSlidesSlide" {
				if !strings.Contains(e.Error(), "An attempt was made to move the position before the beginning of the stream") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			} else {
				if !strings.Contains(e.Error(), "Invalid password") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			}
		} else if paramName == "format" {
			if !strings.Contains(e.Error(), " is not supported.") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "to" {
			if !strings.Contains(e.Error(), "Invalid 'to' parameter") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "cloneFromPosition" {
			if !strings.Contains(e.Error(), "Invalid index: index should be in the range") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "index" {
			if !strings.Contains(e.Error(), "Invalid index: index should be in the range") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "placeholderIndex" {
			if !strings.Contains(e.Error(), "Placeholder with specified index doesn't exist") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "path" {
			if strings.HasSuffix(functionName, "Shapes") {
				if !strings.Contains(e.Error(), "The request is invalid") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			} else {
				if !strings.Contains(e.Error(), "The HTTP resource that matches the request URI") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			}
		} else if paramName == "slideIndex" || paramName == "slides" || paramName == "shapeToClone" {
			if paramName == "shapeToClone" ||
				functionName == "GetSlidesSlideComments" ||
				(strings.HasPrefix(functionName, "GetNotesSlide") &&
					!strings.HasPrefix(functionName, "GetNotesSlideShape")) {
				if !strings.Contains(e.Error(), "Invalid index: index should be in the range") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			} else {
				if !strings.Contains(e.Error(), "Wrong slide index") {
					t.Errorf("Unexpected error message: %v.", e)
					return
				}
			}
		} else if paramName == "shapeIndex" || paramName == "shapes" {
			if !strings.Contains(e.Error(), "Wrong shape index") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "paragraphIndex" || paramName == "paragraphs" {
			if !strings.Contains(e.Error(), "Wrong paragraph index") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "portionIndex" || paramName == "portions" {
			if !strings.Contains(e.Error(), "Wrong portion index") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "position" {
			if !strings.Contains(e.Error(), "Index must be within the bounds of the List") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "oldPosition" {
			if !strings.Contains(e.Error(), "Index was out of range") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "newPosition" {
			if !strings.Contains(e.Error(), "Specified argument was out of the range of valid values") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else if paramName == "data" ||
			paramName == "templatePath" ||
			(strings.HasSuffix(strings.ToLower(paramName), "storage") && functionName == "PostSlidesDocument") {
			if !strings.Contains(e.Error(), "Object reference not set to an instance of an object") {
				t.Errorf("Unexpected error message: %v.", e)
				return
			}
		} else {
			if !strings.Contains(e.Error(), "The specified storage was not found or is not associated with the application") {
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
