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
	"net/url"
	"net/http"
	"strings"
	"context"
	"fmt"
	"os"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type SlidesAsyncApiService service


/* SlidesAsyncApiService 
 @param id 
 @return *os.File*/
func (a *SlidesAsyncApiService) GetOperationResult(id string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(id) == 0 {
		return successPayload, nil, reportError("Missing required parameter id")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/{id}/result"
	idPathStringValue := fmt.Sprintf("%v", id)
	if len(idPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", idPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"id"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	defer successPayload.Close()
        successPayload, err = processFileResponse(responseBody)
        if err != nil {
		return successPayload, localVarHttpResponse, err
        }

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param id 
 @return Operation*/
func (a *SlidesAsyncApiService) GetOperationStatus(id string) (IOperation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IOperation
	)

	if len(id) == 0 {
		return successPayload, nil, reportError("Missing required parameter id")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/{id}"
	idPathStringValue := fmt.Sprintf("%v", id)
	if len(idPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", idPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"id"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Operation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IOperation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IOperation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param document Document data.
 @param format 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
     @param "slides" ([]int32) 
     @param "options" (ExportOptions) 
 @return string*/
func (a *SlidesAsyncApiService) StartConvert(document []byte, format string, password string, storage string, fontsFolder string, slides []int32, options IExportOptions) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	if !ExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/convert/{format}"
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
	}
	if slides != nil {
		localVarQueryParams.Add("Slides", parameterToString(slides, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarPostBody = &options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param document Document data.
 @param format 
 @param outPath 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
     @param "slides" ([]int32) 
     @param "options" (ExportOptions) 
 @return string*/
func (a *SlidesAsyncApiService) StartConvertAndSave(document []byte, format string, outPath string, password string, storage string, fontsFolder string, slides []int32, options IExportOptions) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	if !ExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	if len(outPath) == 0 {
		return successPayload, nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/convert/{format}"
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
	}
	if slides != nil {
		localVarQueryParams.Add("Slides", parameterToString(slides, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarPostBody = &options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param name 
 @param format 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) 
     @param "password" (string) 
     @param "folder" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
     @param "slides" ([]int32) 
 @return string*/
func (a *SlidesAsyncApiService) StartDownloadPresentation(name string, format string, options IExportOptions, password string, folder string, storage string, fontsFolder string, slides []int32) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	if !ExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/{name}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
	}
	if slides != nil {
		localVarQueryParams.Add("Slides", parameterToString(slides, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) 
     @param "storage" (string) 
 @return string*/
func (a *SlidesAsyncApiService) StartMerge(files [][]byte, request IOrderedMergeRequest, storage string) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/merge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	localVarFiles = files
	localVarPostBody = &request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param outPath 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) 
     @param "storage" (string) 
 @return string*/
func (a *SlidesAsyncApiService) StartMergeAndSave(outPath string, files [][]byte, request IOrderedMergeRequest, storage string) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	if len(outPath) == 0 {
		return successPayload, nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/merge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	localVarFiles = files
	localVarPostBody = &request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesAsyncApiService 
 @param name 
 @param format 
 @param outPath 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) 
     @param "password" (string) 
     @param "folder" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
     @param "slides" ([]int32) 
 @return string*/
func (a *SlidesAsyncApiService) StartSavePresentation(name string, format string, outPath string, options IExportOptions, password string, folder string, storage string, fontsFolder string, slides []int32) (string, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload string
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	if !ExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	if len(outPath) == 0 {
		return successPayload, nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/{name}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
	}
	if slides != nil {
		localVarQueryParams.Add("Slides", parameterToString(slides, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json" }

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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse != nil && localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayload = string(responseBytes)
	if len(successPayload) > 1 && successPayload[0] == '"' && successPayload[len(successPayload)-1] == '"' {
		successPayload = successPayload[1:len(successPayload)-1]
	}

	return successPayload, localVarHttpResponse, err
}

