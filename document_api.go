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

type DocumentApiService service


/* DocumentApiService Get API info.

 @return ApiInfoResponse*/
func (a *DocumentApiService) GetSlidesApiInfo() (ApiInfoResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ApiInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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


/* DocumentApiService Read presentation info.

 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) The document password.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) Document&#39;s folder.
 @return DocumentResponse*/
func (a *DocumentApiService) GetSlidesDocument(request GetSlidesDocumentRequest) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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

/* Request for DocumentApiService.GetSlidesDocument
*/
type GetSlidesDocumentRequest struct {
    name string
    password string
    storage string
    folder string
}

/* DocumentApiService Export presentation to some format.

 @param name The document name.
 @param format The conversion format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "jpegQuality" (int32) Quality of the JPEG images inside PDF document.
     @param "password" (string) The document password.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) Document&#39;s folder.
     @param "outPath" (string) Path to save result
     @param "fontsFolder" (string) The optional custom fonts folder.
 @return *os.File*/
func (a *DocumentApiService) GetSlidesDocumentWithFormat(request GetSlidesDocumentWithFormatRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/saveAs/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if request.jpegQuality != nil {
		if err := typeCheckParameter(*request.jpegQuality, "int32", "jpegQuality"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.outPath, "string", "outPath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if request.jpegQuality != nil {
		localVarQueryParams.Add("jpegQuality", parameterToString(*request.jpegQuality, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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

/* Request for DocumentApiService.GetSlidesDocumentWithFormat
*/
type GetSlidesDocumentWithFormatRequest struct {
    name string
    format string
    jpegQuality *int32
    password string
    storage string
    folder string
    outPath string
    fontsFolder string
}

/* DocumentApiService Create presentation 

 @param name The document name.
 @param data 
 @param templatePath Template file path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "templateStorage" (string) Template storage name.
     @param "isImageDataEmbedded" (bool) Is Image Data Embedded
     @param "password" (string) The document password.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) Document&#39;s folder.
 @return DocumentResponse*/
func (a *DocumentApiService) PostSlidesDocument(request PostSlidesDocumentRequest) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.templateStorage, "string", "templateStorage"); err != nil {
		return successPayload, nil, err
	}
	if request.isImageDataEmbedded != nil {
		if err := typeCheckParameter(*request.isImageDataEmbedded, "bool", "isImageDataEmbedded"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("templatePath", parameterToString(request.templatePath, ""))
	if localVarTempParam := request.templateStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("templateStorage", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("isImageDataEmbedded", parameterToString(request.isImageDataEmbedded, ""))
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = request.data
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

/* Request for DocumentApiService.PostSlidesDocument
*/
type PostSlidesDocumentRequest struct {
    name string
    data string
    templatePath string
    templateStorage string
    isImageDataEmbedded *bool
    password string
    storage string
    folder string
}

/* DocumentApiService Performs slides pipeline. Http-request contains pipeline DTO.

 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pipeline" (Pipeline) 
 @return *os.File*/
func (a *DocumentApiService) PostSlidesPipeline(request PostSlidesPipelineRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/pipeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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
	localVarPostBody = &request.pipeline
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

/* Request for DocumentApiService.PostSlidesPipeline
*/
type PostSlidesPipelineRequest struct {
    pipeline IPipeline
}

/* DocumentApiService Saves presentation with options

 @param name The presentation name
 @param format Export format
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Tiff export options
     @param "password" (string) The password to open presentation.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) The presentation folder.
     @param "outPath" (string) The optional output path.
     @param "fontsFolder" (string) The optional custom fonts folder.
 @return *os.File*/
func (a *DocumentApiService) PostSlidesSaveAs(request PostSlidesSaveAsRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/saveAs/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
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
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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

/* Request for DocumentApiService.PostSlidesSaveAs
*/
type PostSlidesSaveAsRequest struct {
    name string
    format string
    options IExportOptions
    password string
    storage string
    folder string
    outPath string
    fontsFolder string
}

/* DocumentApiService Splitting presentations. Create one image per slide.

 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) export options
     @param "format" (string) The format. Default value is jpeg.
     @param "width" (int32) The width of created images.
     @param "height" (int32) The height of created images.
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "destFolder" (string) Folder on storage where images are going to be uploaded. If not specified then images are uploaded to same folder as presentation.
     @param "password" (string) The document password.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SplitDocumentResponse*/
func (a *DocumentApiService) PostSlidesSplit(request PostSlidesSplitRequest) (SplitDocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SplitDocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/split"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.format, "string", "format"); err != nil {
		return successPayload, nil, err
	}
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
	if request.to != nil {
		if err := typeCheckParameter(*request.to, "int32", "to"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.from != nil {
		if err := typeCheckParameter(*request.from, "int32", "from"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.destFolder, "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.format; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, ""))
	}
	if request.width != nil {
		localVarQueryParams.Add("width", parameterToString(*request.width, ""))
	}
	if request.height != nil {
		localVarQueryParams.Add("height", parameterToString(*request.height, ""))
	}
	if request.to != nil {
		localVarQueryParams.Add("to", parameterToString(*request.to, ""))
	}
	if request.from != nil {
		localVarQueryParams.Add("from", parameterToString(*request.from, ""))
	}
	if localVarTempParam := request.destFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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

/* Request for DocumentApiService.PostSlidesSplit
*/
type PostSlidesSplitRequest struct {
    name string
    options IExportOptions
    format string
    width *int32
    height *int32
    to *int32
    from *int32
    destFolder string
    password string
    storage string
    folder string
}

/* DocumentApiService Create presentation 

 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "stream" ([]byte) 
     @param "templatePath" (string) Template file path.
     @param "templatePassword" (string) 
     @param "templateStorage" (string) Template storage name.
     @param "password" (string) The document password.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) Document&#39;s folder.
 @return DocumentResponse*/
func (a *DocumentApiService) PutNewPresentation(request PutNewPresentationRequest) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.templatePath, "string", "templatePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.templatePassword, "string", "templatePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.templateStorage, "string", "templateStorage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.templatePath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("templatePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.templatePassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("templatePassword", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.templateStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("templateStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = request.stream
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

/* Request for DocumentApiService.PutNewPresentation
*/
type PutNewPresentationRequest struct {
    name string
    stream []byte
    templatePath string
    templatePassword string
    templateStorage string
    password string
    storage string
    folder string
}

/* DocumentApiService Convert presentation from request content to format specified.

 @param format The format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "document" ([]byte) 
     @param "password" (string) The document password.
     @param "outPath" (string) Path to save result
     @param "fontsFolder" (string) The optional custom fonts folder.
 @return *os.File*/
func (a *DocumentApiService) PutSlidesConvert(request PutSlidesConvertRequest) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.outPath, "string", "outPath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("format", parameterToString(request.format, ""))
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = request.document
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

/* Request for DocumentApiService.PutSlidesConvert
*/
type PutSlidesConvertRequest struct {
    format string
    document []byte
    password string
    outPath string
    fontsFolder string
}

/* DocumentApiService Create presentation document from html

 @param name The document name.
 @param html 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) The document password.
     @param "storage" (string) Document&#39;s storage.
     @param "folder" (string) Document&#39;s folder.
 @return DocumentResponse*/
func (a *DocumentApiService) PutSlidesDocumentFromHtml(request PutSlidesDocumentFromHtmlRequest) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromHtml"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = request.html
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

/* Request for DocumentApiService.PutSlidesDocumentFromHtml
*/
type PutSlidesDocumentFromHtmlRequest struct {
    name string
    html string
    password string
    storage string
    folder string
}

/* DocumentApiService Set slide size for the presentation.

 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) The document password.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "width" (int32) Slide width.
     @param "height" (int32) Slide height.
     @param "sizeType" (string) Standard slide size type.
     @param "scaleType" (string) Standard slide scale type.
 @return DocumentResponse*/
func (a *DocumentApiService) PutSlidesSlideSize(request PutSlidesSlideSizeRequest) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slidesize"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
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
	if err := typeCheckParameter(request.sizeType, "string", "sizeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.scaleType, "string", "scaleType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if request.width != nil {
		localVarQueryParams.Add("width", parameterToString(*request.width, ""))
	}
	if request.height != nil {
		localVarQueryParams.Add("height", parameterToString(*request.height, ""))
	}
	if localVarTempParam := request.sizeType; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sizeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.scaleType; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("scaleType", parameterToString(localVarTempParam, ""))
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

/* Request for DocumentApiService.PutSlidesSlideSize
*/
type PutSlidesSlideSizeRequest struct {
    name string
    password string
    storage string
    folder string
    width *int32
    height *int32
    sizeType string
    scaleType string
}

