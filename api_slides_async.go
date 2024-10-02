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
 @param path 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) 
     @param "versionId" (string) 
 @return *os.File*/
func (a *SlidesAsyncApiService) Download(path string, storageName string, versionId string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(versionId, "string", "versionId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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

/* SlidesAsyncApiService 
 @param name 
 @param format 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) 
     @param "width" (int32) 
     @param "height" (int32) 
     @param "from" (int32) 
     @param "to" (int32) 
     @param "destFolder" (string) 
     @param "password" (string) 
     @param "folder" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
 @return string*/
func (a *SlidesAsyncApiService) StartSplit(name string, format string, options IExportOptions, width *int32, height *int32, from *int32, to *int32, destFolder string, password string, folder string, storage string, fontsFolder string) (string, *http.Response, error) {
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
	if !SlideExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/{name}/split/{format}"
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

	if width != nil {
		if err := typeCheckParameter(*width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if height != nil {
		if err := typeCheckParameter(*height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if from != nil {
		if err := typeCheckParameter(*from, "int32", "from"); err != nil {
			return successPayload, nil, err
		}
	}
	if to != nil {
		if err := typeCheckParameter(*to, "int32", "to"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(destFolder, "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}
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

	if width != nil {
		localVarQueryParams.Add("Width", parameterToString(*width, ""))
	}
	if height != nil {
		localVarQueryParams.Add("Height", parameterToString(*height, ""))
	}
	if from != nil {
		localVarQueryParams.Add("From", parameterToString(*from, ""))
	}
	if to != nil {
		localVarQueryParams.Add("To", parameterToString(*to, ""))
	}
	if localVarTempParam := destFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestFolder", parameterToString(localVarTempParam, ""))
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
 @param document Document data.
 @param format 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "destFolder" (string) 
     @param "width" (int32) 
     @param "height" (int32) 
     @param "from" (int32) 
     @param "to" (int32) 
     @param "password" (string) 
     @param "storage" (string) 
     @param "fontsFolder" (string) 
     @param "options" (ExportOptions) 
 @return string*/
func (a *SlidesAsyncApiService) StartUploadAndSplit(document []byte, format string, destFolder string, width *int32, height *int32, from *int32, to *int32, password string, storage string, fontsFolder string, options IExportOptions) (string, *http.Response, error) {
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
	if !SlideExportFormat_Validate(format) {
		return successPayload, nil, reportError("Invalid value for parameter format: " + format)
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/split/{format}"
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(destFolder, "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}
	if width != nil {
		if err := typeCheckParameter(*width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if height != nil {
		if err := typeCheckParameter(*height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if from != nil {
		if err := typeCheckParameter(*from, "int32", "from"); err != nil {
			return successPayload, nil, err
		}
	}
	if to != nil {
		if err := typeCheckParameter(*to, "int32", "to"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := destFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestFolder", parameterToString(localVarTempParam, ""))
	}
	if width != nil {
		localVarQueryParams.Add("Width", parameterToString(*width, ""))
	}
	if height != nil {
		localVarQueryParams.Add("Height", parameterToString(*height, ""))
	}
	if from != nil {
		localVarQueryParams.Add("From", parameterToString(*from, ""))
	}
	if to != nil {
		localVarQueryParams.Add("To", parameterToString(*to, ""))
	}
	if localVarTempParam := storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
 @param path 
 @param file File to upload
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) 
 @return FilesUploadResult*/
func (a *SlidesAsyncApiService) Upload(path string, file []byte, storageName string) (IFilesUploadResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFilesUploadResult
	)

	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if len(file) == 0 {
		return successPayload, nil, reportError("Missing required parameter file")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/async/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	if len(file) > 0 {
		localVarFiles = append(localVarFiles, file)
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

	successPayloadObject, err := createObjectForType("FilesUploadResult", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IFilesUploadResult); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IFilesUploadResult); true {
	}

	return successPayload, localVarHttpResponse, err
}

