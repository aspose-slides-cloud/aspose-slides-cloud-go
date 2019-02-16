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

type NotesSlideApiService service


/* NotesSlideApiService Remove Notes Slide.

 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideResponse*/
func (a *NotesSlideApiService) DeleteNotesSlide(request DeleteNotesSlideRequest) (SlideResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
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

/* Request for NotesSlideApiService.DeleteNotesSlide
*/
type DeleteNotesSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideApiService Read Notes slide info.

 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlideResponse*/
func (a *NotesSlideApiService) GetNotesSlide(request GetNotesSlideRequest) (NotesSlideResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlideResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
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

/* Request for NotesSlideApiService.GetNotesSlide
*/
type GetNotesSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* NotesSlideApiService Convert Notes Slide to the specified picture format.

 @param name Document name.
 @param slideIndex Slide index.
 @param format Output file format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Output file width.
     @param "height" (int32) Output file height.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *NotesSlideApiService) GetNotesSlideWithFormat(request GetNotesSlideWithFormatRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/saveAs/{format}"
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
	formatPathStringValue := fmt.Sprintf("%v", request.format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.width != nil {
		if err := typeCheckParameter(*request.width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.height != nil {
		if err := typeCheckParameter(*request.height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
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

	if request.width != nil {
		localVarQueryParams.Add("width", parameterToString(*request.width, ""))
	}
	if request.height != nil {
		localVarQueryParams.Add("height", parameterToString(*request.height, ""))
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

/* Request for NotesSlideApiService.GetNotesSlideWithFormat
*/
type GetNotesSlideWithFormatRequest struct {
    name string
    slideIndex int32
    format string
    width *int32
    height *int32
    password string
    folder string
    storage string
}

/* NotesSlideApiService Add new Notes Slide.

 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (NotesSlide) A  object with notes slide data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlideResponse*/
func (a *NotesSlideApiService) PostAddNotesSlide(request PostAddNotesSlideRequest) (NotesSlideResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlideResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
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

/* Request for NotesSlideApiService.PostAddNotesSlide
*/
type PostAddNotesSlideRequest struct {
    name string
    slideIndex int32
    dto INotesSlide
    password string
    folder string
    storage string
}

/* NotesSlideApiService Update Notes Slide properties.

 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (NotesSlide) A  object with notes slide data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlideResponse*/
func (a *NotesSlideApiService) PutUpdateNotesSlide(request PutUpdateNotesSlideRequest) (NotesSlideResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlideResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
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

/* Request for NotesSlideApiService.PutUpdateNotesSlide
*/
type PutUpdateNotesSlideRequest struct {
    name string
    slideIndex int32
    dto INotesSlide
    password string
    folder string
    storage string
}

