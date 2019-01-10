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
	"bytes"
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"context"
	"os"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type NotesSlideShapesApiService service


/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ParagraphListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlideParagraph(request DeleteNotesSlideParagraphRequest) (ParagraphListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlideParagraph
*/
type DeleteNotesSlideParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ParagraphListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlideParagraphs(request DeleteNotesSlideParagraphsRequest) (ParagraphListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.paragraphs, "[]int32", "paragraphs"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlideParagraphs
*/
type DeleteNotesSlideParagraphsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphs []int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return PortionListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlidePortion(request DeleteNotesSlidePortionRequest) (PortionListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlidePortion
*/
type DeleteNotesSlidePortionRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    portionIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return PortionListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlidePortions(request DeleteNotesSlidePortionsRequest) (PortionListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.portions, "[]int32", "portions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlidePortions
*/
type DeleteNotesSlidePortionsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    portions []int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ShapeListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlideShape(request DeleteNotesSlideShapeRequest) (ShapeListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlideShape
*/
type DeleteNotesSlideShapeRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Removes a shape, specified shapes or all shapes.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ShapeListResponse*/
func (a *NotesSlideShapesApiService) DeleteNotesSlideShapes(request DeleteNotesSlideShapesRequest) (ShapeListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.shapes, "[]int32", "shapes"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.DeleteNotesSlideShapes
*/
type DeleteNotesSlideShapesRequest struct {
    name string
    slideIndex int32
    path string
    shapes []int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ShapeResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShape(request GetNotesSlideShapeRequest) (ShapeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShape
*/
type GetNotesSlideShapeRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ParagraphResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapeParagraph(request GetNotesSlideShapeParagraphRequest) (ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapeParagraph
*/
type GetNotesSlideShapeParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ParagraphListResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapeParagraphs(request GetNotesSlideShapeParagraphsRequest) (ParagraphListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapeParagraphs
*/
type GetNotesSlideShapeParagraphsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return PortionResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapePortion(request GetNotesSlideShapePortionRequest) (PortionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapePortion
*/
type GetNotesSlideShapePortionRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    portionIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return PortionListResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapePortions(request GetNotesSlideShapePortionsRequest) (PortionListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapePortions
*/
type GetNotesSlideShapePortionsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Render shape to specified picture format.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "outPath" (string) Output path.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapeWithFormat(request GetNotesSlideShapeWithFormatRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/saveAs/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.scaleX != nil {
		if err := typeCheckParameter(*request.scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.scaleY != nil {
		if err := typeCheckParameter(*request.scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.outPath, "string", "outPath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("scaleX", parameterToString(request.scaleX, ""))
	localVarQueryParams.Add("scaleY", parameterToString(request.scaleY, ""))
	if localVarTempParam := request.bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.outPath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("outPath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("fontsFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	defer successPayload.Close()
        successPayload, err = processFileResponse(responseBody)
        if err != nil {
		return successPayload, localVarHttpResponse, err
        }

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapeWithFormat
*/
type GetNotesSlideShapeWithFormatRequest struct {
    name string
    slideIndex int32
    shapeIndex int32
    format string
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds string
    outPath string
    fontsFolder string
}

/* NotesSlideShapesApiService Read slide shapes or shape info.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ShapeListResponse*/
func (a *NotesSlideShapesApiService) GetNotesSlideShapes(request GetNotesSlideShapesRequest) (ShapeListResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.GetNotesSlideShapes
*/
type GetNotesSlideShapesRequest struct {
    name string
    slideIndex int32
    path string
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Creates new paragraph.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return ParagraphResponse*/
func (a *NotesSlideShapesApiService) PostNotesSlideAddNewParagraph(request PostNotesSlideAddNewParagraphRequest) (ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.position != nil {
		if err := typeCheckParameter(*request.position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if request.position != nil {
		localVarQueryParams.Add("position", parameterToString(*request.position, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PostNotesSlideAddNewParagraph
*/
type PostNotesSlideAddNewParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    dto IParagraph
    password string
    folder string
    storage string
    position *int32
}

/* NotesSlideShapesApiService Creates new portion.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return PortionResponse*/
func (a *NotesSlideShapesApiService) PostNotesSlideAddNewPortion(request PostNotesSlideAddNewPortionRequest) (PortionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.position != nil {
		if err := typeCheckParameter(*request.position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if request.position != nil {
		localVarQueryParams.Add("position", parameterToString(*request.position, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PostNotesSlideAddNewPortion
*/
type PostNotesSlideAddNewPortionRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    dto IPortion
    password string
    folder string
    storage string
    position *int32
}

/* NotesSlideShapesApiService Creates new shape.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding the new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeResponse*/
func (a *NotesSlideShapesApiService) PostNotesSlideAddNewShape(request PostNotesSlideAddNewShapeRequest) (ShapeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.shapeToClone != nil {
		if err := typeCheckParameter(*request.shapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.position != nil {
		if err := typeCheckParameter(*request.position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if request.shapeToClone != nil {
		localVarQueryParams.Add("shapeToClone", parameterToString(*request.shapeToClone, ""))
	}
	if request.position != nil {
		localVarQueryParams.Add("position", parameterToString(*request.position, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PostNotesSlideAddNewShape
*/
type PostNotesSlideAddNewShapeRequest struct {
    name string
    slideIndex int32
    path string
    dto IShapeBase
    password string
    folder string
    storage string
    shapeToClone *int32
    position *int32
}

/* NotesSlideShapesApiService Render shape to specified picture format.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "outPath" (string) Output path.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *NotesSlideShapesApiService) PostNotesSlideShapeSaveAs(request PostNotesSlideShapeSaveAsRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/saveAs/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.scaleX != nil {
		if err := typeCheckParameter(*request.scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.scaleY != nil {
		if err := typeCheckParameter(*request.scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.outPath, "string", "outPath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("scaleX", parameterToString(request.scaleX, ""))
	localVarQueryParams.Add("scaleY", parameterToString(request.scaleY, ""))
	if localVarTempParam := request.bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.outPath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("outPath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("fontsFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.options
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	defer successPayload.Close()
        successPayload, err = processFileResponse(responseBody)
        if err != nil {
		return successPayload, localVarHttpResponse, err
        }

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PostNotesSlideShapeSaveAs
*/
type PostNotesSlideShapeSaveAsRequest struct {
    name string
    slideIndex int32
    shapeIndex int32
    format string
    options IIShapeExportOptions
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds string
    outPath string
    fontsFolder string
}

/* NotesSlideShapesApiService Updates shape properties.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Object path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ShapeResponse*/
func (a *NotesSlideShapesApiService) PutUpdateNotesSlideShape(request PutUpdateNotesSlideShapeRequest) (ShapeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PutUpdateNotesSlideShape
*/
type PutUpdateNotesSlideShapeRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    dto IShapeBase
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Updates shape properties.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Object path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return ParagraphResponse*/
func (a *NotesSlideShapesApiService) PutUpdateNotesSlideShapeParagraph(request PutUpdateNotesSlideShapeParagraphRequest) (ParagraphResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ParagraphResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PutUpdateNotesSlideShapeParagraph
*/
type PutUpdateNotesSlideShapeParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    dto IParagraph
    password string
    folder string
    storage string
}

/* NotesSlideShapesApiService Updates shape properties.

 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Object path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Shape dto.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
 @return PortionResponse*/
func (a *NotesSlideShapesApiService) PutUpdateNotesSlideShapePortion(request PutUpdateNotesSlideShapePortionRequest) (PortionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PortionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &request.dto
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
                successPayload.Code = int32(localVarHttpResponse.StatusCode)
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(errorMessage.getMessage())
	}

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for NotesSlideShapesApiService.PutUpdateNotesSlideShapePortion
*/
type PutUpdateNotesSlideShapePortionRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    portionIndex int32
    dto IPortion
    password string
    folder string
    storage string
}

