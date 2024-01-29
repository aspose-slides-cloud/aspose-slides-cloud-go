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

package usecasetests

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v24"
)

/*
   Test for asynchronous convert presentation
*/
func TestAsyncConvert(t *testing.T) {
	var sleepTimeout int32 = 3
	maxRetries := 10

	c := slidescloud.GetTestSlidesAsyncApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	operationId, _, e := c.SlidesAsyncApi.StartConvert(source, "pdf", password, "", "", nil, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	converted, _, e := c.SlidesAsyncApi.GetOperationResult(operationId)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	convertedStat, e := os.Stat(converted.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if convertedStat.Size() <= 0 {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", 0, convertedStat.Size())
		return
	}
}

/*
   Test for asynchronous download presentation
*/
func TestAsyncDownloadPresentation(t *testing.T) {
	sleepTimeout := 3
	maxRetries := 10
	fileName := "test.pptx"

	_, e := slidescloud.GetTestSlidesApiClient().SlidesApi.CopyFile("TempTests/" + fileName, folderName +"/"+ fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := slidescloud.GetTestSlidesAsyncApiClient()
	operationId, _, e := c.SlidesAsyncApi.StartDownloadPresentation(fileName, "pdf", nil, password, folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	converted, _, e := c.SlidesAsyncApi.GetOperationResult(operationId)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	convertedStat, e := os.Stat(converted.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if convertedStat.Size() <= 0 {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", 0, convertedStat.Size())
		return
	}
}

/*
   Test for asynchronous convert and save presentation
*/
func TestAsyncConvertAndSave(t *testing.T) {
	var sleepTimeout int32 = 3
	maxRetries := 10

	c := slidescloud.GetTestSlidesAsyncApiClient()
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	outPath := "TempSlidesSDK/out.pptx"
	operationId, _, e := c.SlidesAsyncApi.StartConvertAndSave(source, "pdf", outPath, "password", "", "", nil, nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for asynchronous save presentation
*/
func TestAsyncSavePresentation(t *testing.T) {
	var sleepTimeout int32 = 3
	maxRetries := 10
	fileName := "test.pptx"
	outPath := "TempSlidesSDK/out.pptx"

	a := slidescloud.GetTestSlidesApiClient()
	_, e := a.SlidesApi.CopyFile("TempTests/" + fileName, folderName +"/"+ fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = a.SlidesApi.DeleteFile(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := slidescloud.GetTestSlidesAsyncApiClient()

	operationId, _, e := c.SlidesAsyncApi.StartSavePresentation(fileName, "pdf", outPath, nil, "password", folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	resultExists, _, e := a.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for asynchronous merge presentations
*/
func TestAsyncMerge(t *testing.T) {
	var sleepTimeout int32 = 3
	maxRetries := 10

	source1, e := ioutil.ReadFile("../TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("../TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := slidescloud.GetTestSlidesAsyncApiClient()

	operationId, _, e := c.SlidesAsyncApi.StartMerge([][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetProgress() == nil {
		t.Errorf("Progress must not be empty.")
	}
	if operation.GetProgress().GetStepCount() != 2 {
		t.Errorf("Wrong progress step count. Expected %v but was %v.", 2, operation.GetProgress().GetStepCount())
	}
	if operation.GetProgress().GetStepCount() != operation.GetProgress().GetStepIndex() {
		t.Errorf("Wrong progress step index. Expected %v but was %v.", operation.GetProgress().GetStepCount(), operation.GetProgress().GetStepIndex())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	merged, _, e := c.SlidesAsyncApi.GetOperationResult(operationId)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	mergedStat, e := os.Stat(merged.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if mergedStat.Size() <= 0 {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", 0, mergedStat.Size())
		return
	}
}

/*
   Test for asynchronous merge and save presentations
*/
func TestAsyncMergeAndSave(t *testing.T) {
	var sleepTimeout int32 = 3
	maxRetries := 10

	source1, e := ioutil.ReadFile("../TestData/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("../TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	outPath := "TempSlidesSDK/merged.pptx"

	a := slidescloud.GetTestSlidesApiClient()
	_, e = a.SlidesApi.DeleteFile(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c := slidescloud.GetTestSlidesAsyncApiClient()

	operationId, _, e := c.SlidesAsyncApi.StartMergeAndSave(outPath, [][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Finished" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Finished", operation.GetStatus())
	}
	if operation.GetProgress() == nil {
		t.Errorf("Progress must not be empty.")
	}
	if operation.GetProgress().GetStepCount() != 2 {
		t.Errorf("Wrong progress step count. Expected %v but was %v.", 2, operation.GetProgress().GetStepCount())
	}
	if operation.GetProgress().GetStepCount() != operation.GetProgress().GetStepIndex() {
		t.Errorf("Wrong progress step index. Expected %v but was %v.", operation.GetProgress().GetStepCount(), operation.GetProgress().GetStepIndex())
	}
	if operation.GetError() != "" {
		t.Errorf("Unexpected error: %v.", operation.GetError())
	}

	resultExists, _, e := a.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for asynchronous bad operation
*/
func TestAsyncBadOperation(t *testing.T) {
	sleepTimeout := 3
	maxRetries := 10
	c := slidescloud.GetTestSlidesAsyncApiClient()

	operationId, _, e := c.SlidesAsyncApi.StartDownloadPresentation("IDoNotExist.pptx", "pdf", nil, "", "", "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var operation slidescloud.IOperation
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(sleepTimeout) * time.Second)
		operation, _, e = c.SlidesAsyncApi.GetOperationStatus(operationId)
		if e != nil {
			t.Errorf("Error: %v.", e)
			return
		}
		if operation.GetStatus() != "Created" && operation.GetStatus() != "Enqueued" && operation.GetStatus() != "Started" {
			break
		}
	}

	if operation.GetStatus() != "Failed" {
		t.Errorf("Wrong operation status. Expected %v but was %v.", "Failed", operation.GetStatus())
	}
	if operation.GetError() == "" {
		t.Errorf("An error expected.")
	}
}
