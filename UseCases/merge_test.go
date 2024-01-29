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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v24"
)

/*
   Test for merge from storage
*/
func TestMergeStorage(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	fileNamePdf := "test.pdf"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileNamePdf, folderName+"/"+fileNamePdf, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewPresentationsMergeRequest()
	request.PresentationPaths = []string{folderName + "/" + fileName2, folderName + "/" + fileNamePdf}
	_, _, e = c.SlidesApi.Merge(fileName, request, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for ordered merge from storage
*/
func TestMergeOrderedStorage(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation := slidescloud.NewPresentationToMerge()
	presentation.Path = folderName + "/" + fileName2
	presentation.Slides = []int32{2, 1}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation}
	_, _, e = c.SlidesApi.OrderedMerge(fileName, request, "password", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge from request
*/
func TestMergeRequest(t *testing.T) {
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

	_, _, e = slidescloud.GetTestSlidesApiClient().SlidesApi.MergeOnline([][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge and save from request
*/
func TestMergeAndSaveRequest(t *testing.T) {
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
	outPath := "TestData/out.pptx"
	c := slidescloud.GetTestSlidesApiClient()

	_, e = c.SlidesApi.MergeAndSaveOnline(outPath, [][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
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
   Test for ordered merge from request
*/
func TestMergeOrderedRequest(t *testing.T) {
	source1, e := ioutil.ReadFile("../TestData/test.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile("../TestData/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = "password"
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Path = "file2"
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}
	_, _, e = slidescloud.GetTestSlidesApiClient().SlidesApi.MergeOnline([][]byte{source1, source2}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for combined merge from request
*/
func TestMergeCombinedRequest(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName2, folderName+"/"+fileName2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = "password"
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Source = "Storage"
	presentation2.Path = folderName + "/" + fileName2
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}
	_, _, e = slidescloud.GetTestSlidesApiClient().SlidesApi.MergeOnline([][]byte{source}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for
*/
func TestMergeOrderedUrl(t *testing.T) {

	c := slidescloud.GetTestSlidesApiClient()
	_, e := c.SlidesApi.CopyFile("TempTests/"+fileName, folderName+"/"+fileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Source = "Storage"
	presentation1.Path = folderName + "/" + fileName
	presentation1.Slides = []int32{1, 2}
	presentation1.Password = password
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Source = "Url"
	presentation2.Path = "https://drive.google.com/uc?export=download&id=1ycMzd7e--Ro9H8eH2GL5fPP7-2HjX4My"
	presentation2.Slides = []int32{1}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}

	_, _, e = c.SlidesApi.MergeOnline(nil, request, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
