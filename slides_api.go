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
	"os"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type SlidesApiService service


/* SlidesApiService Changes the placement of selected shapes on the slide. Aligns shapes to the margins or the edge of the slide or aligns them relative to each other.
 @param name Document name.
 @param slideIndex Slide index.
 @param alignmentType Alignment type that will be applied to the shapes.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "alignToSlide" (bool) If true, shapes will be aligned relative to the slide edges.
     @param "shapes" ([]int32) Shapes indexes.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) AlignShapes(name string, slideIndex int32, alignmentType string, alignToSlide *bool, shapes []int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(alignmentType) == 0 {
		return successPayload, nil, reportError("Missing required parameter alignmentType")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/align/{alignmentType}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	alignmentTypePathStringValue := fmt.Sprintf("%v", alignmentType)
	if len(alignmentTypePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"alignmentType"+"}", alignmentTypePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"alignmentType"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if alignToSlide != nil {
		if err := typeCheckParameter(*alignToSlide, "bool", "alignToSlide"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(shapes, "[]int32", "shapes"); err != nil {
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

	if alignToSlide != nil {
		localVarQueryParams.Add("AlignToSlide", parameterToString(alignToSlide, ""))
	}
	if localVarTempParam := shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Convert presentation from request content to format specified.
 @param document Document data.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) Convert(document []byte, format string, password string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Convert presentation from request content to format specified.
 @param document Document data.
 @param format Export format.
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return */
func (a *SlidesApiService) ConvertAndSave(document []byte, format string, outPath string, password string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(document) == 0 {
		return nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Copy file
 @param srcPath Source file path e.g. &#39;/folder/file.ext&#39;
 @param destPath Destination file path
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
     @param "versionId" (string) File version ID to copy
 @return */
func (a *SlidesApiService) CopyFile(srcPath string, destPath string, srcStorageName string, destStorageName string, versionId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(srcPath) == 0 {
		return nil, reportError("Missing required parameter srcPath")
	}
	if len(destPath) == 0 {
		return nil, reportError("Missing required parameter destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(versionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam := srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Copy folder
 @param srcPath Source folder path e.g. &#39;/src&#39;
 @param destPath Destination folder path e.g. &#39;/dst&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *SlidesApiService) CopyFolder(srcPath string, destPath string, srcStorageName string, destStorageName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(srcPath) == 0 {
		return nil, reportError("Missing required parameter srcPath")
	}
	if len(destPath) == 0 {
		return nil, reportError("Missing required parameter destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam := srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Copy layoutSlide from source presentation.
 @param name Document name.
 @param cloneFrom Name of the document to clone layoutSlide from.
 @param cloneFromPosition Position of cloned layout slide.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "cloneFromPassword" (string) Password for the document to clone layoutSlide from.
     @param "cloneFromStorage" (string) Storage of the document to clone layoutSlide from.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) CopyLayoutSlide(name string, cloneFrom string, cloneFromPosition int32, cloneFromPassword string, cloneFromStorage string, password string, folder string, storage string) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(cloneFrom) == 0 {
		return successPayload, nil, reportError("Missing required parameter cloneFrom")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(cloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(cloneFromStorage, "string", "cloneFromStorage"); err != nil {
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

	localVarQueryParams.Add("CloneFrom", parameterToString(cloneFrom, ""))
	localVarQueryParams.Add("CloneFromPosition", parameterToString(cloneFromPosition, ""))
	if localVarTempParam := cloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("CloneFromStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := cloneFromPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["CloneFromPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("LayoutSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ILayoutSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ILayoutSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Copy masterSlide from source presentation.
 @param name Document name.
 @param cloneFrom Name of the document to clone masterSlide from.
 @param cloneFromPosition Position of cloned master slide.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "cloneFromPassword" (string) Password for the document to clone masterSlide from.
     @param "cloneFromStorage" (string) Storage of the document to clone masterSlide from.
     @param "applyToAll" (bool) True to apply cloned master slide to every existing slide.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return MasterSlide*/
func (a *SlidesApiService) CopyMasterSlide(name string, cloneFrom string, cloneFromPosition int32, cloneFromPassword string, cloneFromStorage string, applyToAll *bool, password string, folder string, storage string) (IMasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(cloneFrom) == 0 {
		return successPayload, nil, reportError("Missing required parameter cloneFrom")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(cloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(cloneFromStorage, "string", "cloneFromStorage"); err != nil {
		return successPayload, nil, err
	}
	if applyToAll != nil {
		if err := typeCheckParameter(*applyToAll, "bool", "applyToAll"); err != nil {
			return successPayload, nil, err
		}
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

	localVarQueryParams.Add("CloneFrom", parameterToString(cloneFrom, ""))
	localVarQueryParams.Add("CloneFromPosition", parameterToString(cloneFromPosition, ""))
	if localVarTempParam := cloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("CloneFromStorage", parameterToString(localVarTempParam, ""))
	}
	if applyToAll != nil {
		localVarQueryParams.Add("ApplyToAll", parameterToString(applyToAll, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := cloneFromPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["CloneFromPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("MasterSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IMasterSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IMasterSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Copy a slide from the current or another presentation.
 @param name Document name.
 @param slideToCopy The index of the slide to be copied from the source presentation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) The target position at which to copy the slide. Copy to the end by default.
     @param "source" (string) Name of the document to copy a slide from.
     @param "sourcePassword" (string) Password for the document to copy a slide from.
     @param "sourceStorage" (string) Template storage name.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) CopySlide(name string, slideToCopy int32, position *int32, source string, sourcePassword string, sourceStorage string, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/copy"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(source, "string", "source"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(sourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(sourceStorage, "string", "sourceStorage"); err != nil {
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

	localVarQueryParams.Add("SlideToCopy", parameterToString(slideToCopy, ""))
	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := source; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Source", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := sourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourceStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := sourcePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["SourcePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Add an effect to slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param effect Animation effect DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) CreateAnimationEffect(name string, slideIndex int32, effect IEffect, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if effect == nil {
		return successPayload, nil, reportError("Missing required parameter effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param sequence Animation sequence DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) CreateAnimationInteractiveSequence(name string, slideIndex int32, sequence IInteractiveSequence, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if sequence == nil {
		return successPayload, nil, reportError("Missing required parameter sequence")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &sequence
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Add an animation effect to a slide interactive sequence.
 @param name Document name.
 @param slideIndex Slide index.
 @param sequenceIndex The position of the interactive sequence.
 @param effect Animation effect DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) CreateAnimationInteractiveSequenceEffect(name string, slideIndex int32, sequenceIndex int32, effect IEffect, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if effect == nil {
		return successPayload, nil, reportError("Missing required parameter effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", sequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Add a new category to a chart.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param category Category DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) CreateChartCategory(name string, slideIndex int32, shapeIndex int32, category IChartCategory, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if category == nil {
		return successPayload, nil, reportError("Missing required parameter category")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &category
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Add a new data point to a chart series.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param seriesIndex Series index.
 @param dataPoint Data point DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) CreateChartDataPoint(name string, slideIndex int32, shapeIndex int32, seriesIndex int32, dataPoint IDataPoint, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dataPoint == nil {
		return successPayload, nil, reportError("Missing required parameter dataPoint")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", seriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dataPoint
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Add a new series to a chart.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index (must be a chart).
 @param series Series DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) CreateChartSeries(name string, slideIndex int32, shapeIndex int32, series ISeries, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if series == nil {
		return successPayload, nil, reportError("Missing required parameter series")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &series
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Adds the comment on the slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto Comment DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideComments*/
func (a *SlidesApiService) CreateComment(name string, slideIndex int32, dto ISlideComment, password string, folder string, storage string) (ISlideComments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideComments
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/comments"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideComments", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideComments); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideComments); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Adds the comment on the slide.
 @param document Document data.
 @param slideIndex Slide index.
 @param dto Comment DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) CreateCommentOnline(document []byte, slideIndex int32, dto ISlideComment, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/comments"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Create the folder
 @param path Folder path to create e.g. &#39;folder_1/folder_2/&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return */
func (a *SlidesApiService) CreateFolder(path string, storageName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(path) == 0 {
		return nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storageName, "string", "storageName"); err != nil {
		return nil, err
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Adds an image watermark to each slide of the presentation.  Image can be provided as a part of the form or withing PictureFrame DTO for detailed customization. Both options are applicable simultaneously. 
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "image" ([]byte) Image data.
     @param "pictureFrame" (PictureFrame) PictureFrame DTO
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) CreateImageWatermark(name string, image []byte, pictureFrame IPictureFrame, password string, folder string, storage string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/watermark/image"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(image) > 0 {
		localVarFiles = append(localVarFiles, image)
	}
	localVarPostBody = &pictureFrame
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Adds an image watermark to each slide of the presentation.  Image can be provided as a part of the form or withing PictureFrame DTO for detailed customization. Both options are applicable simultaneously. 
 @param document Document data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "image" ([]byte) Image data.
     @param "pictureFrame" (PictureFrame) PictureFrame DTO.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) CreateImageWatermarkOnline(document []byte, image []byte, pictureFrame IPictureFrame, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/watermark/image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	if len(image) > 0 {
		localVarFiles = append(localVarFiles, image)
	}
	localVarPostBody = &pictureFrame
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Add new notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto A NotesSlide object with notes slide data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) CreateNotesSlide(name string, slideIndex int32, dto INotesSlide, password string, folder string, storage string) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) CreateNotesSlideParagraph(name string, slideIndex int32, shapeIndex int32, dto IParagraph, position *int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) CreateNotesSlidePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, dto IPortion, position *int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto Shape DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) CreateNotesSlideShape(name string, slideIndex int32, dto IShapeBase, shapeToClone *int32, position *int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if shapeToClone != nil {
		if err := typeCheckParameter(*shapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if shapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*shapeToClone, ""))
	}
	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) CreateParagraph(name string, slideIndex int32, shapeIndex int32, dto IParagraph, position *int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) CreatePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, dto IPortion, position *int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create a presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" ([]byte) Document input data.
     @param "inputPassword" (string) The password for input document.
     @param "password" (string) The document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) CreatePresentation(name string, data []byte, inputPassword string, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(inputPassword, "string", "inputPassword"); err != nil {
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := inputPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["InputPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(data) > 0 {
		localVarFiles = append(localVarFiles, data)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create a presentation from an existing source.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "sourcePath" (string) Template file path.
     @param "sourcePassword" (string) Template file password.
     @param "sourceStorage" (string) Template storage name.
     @param "password" (string) The document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) CreatePresentationFromSource(name string, sourcePath string, sourcePassword string, sourceStorage string, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromSource"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(sourcePath, "string", "sourcePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(sourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(sourceStorage, "string", "sourceStorage"); err != nil {
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

	if localVarTempParam := sourcePath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourcePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := sourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourceStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := sourcePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["SourcePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create a presentation.
 @param name Document name.
 @param templatePath Template file path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" (string) Document input data.
     @param "templatePassword" (string) Template file password.
     @param "templateStorage" (string) Template storage name.
     @param "isImageDataEmbedded" (bool) True if image data is embedded.
     @param "password" (string) The document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) CreatePresentationFromTemplate(name string, templatePath string, data string, templatePassword string, templateStorage string, isImageDataEmbedded *bool, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(templatePath) == 0 {
		return successPayload, nil, reportError("Missing required parameter templatePath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromTemplate"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(data, "string", "data"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(templatePassword, "string", "templatePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(templateStorage, "string", "templateStorage"); err != nil {
		return successPayload, nil, err
	}
	if isImageDataEmbedded != nil {
		if err := typeCheckParameter(*isImageDataEmbedded, "bool", "isImageDataEmbedded"); err != nil {
			return successPayload, nil, err
		}
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

	localVarQueryParams.Add("TemplatePath", parameterToString(templatePath, ""))
	if localVarTempParam := templateStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("TemplateStorage", parameterToString(localVarTempParam, ""))
	}
	if isImageDataEmbedded != nil {
		localVarQueryParams.Add("IsImageDataEmbedded", parameterToString(isImageDataEmbedded, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := templatePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["TemplatePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = data
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create a section starting at a specified slide index.
 @param name Document name.
 @param sectionName Section name.
 @param slideIndex Slide index (one-based).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) CreateSection(name string, sectionName string, slideIndex int32, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(sectionName) == 0 {
		return successPayload, nil, reportError("Missing required parameter sectionName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	localVarQueryParams.Add("SectionName", parameterToString(sectionName, ""))
	localVarQueryParams.Add("SlideIndex", parameterToString(slideIndex, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) CreateShape(name string, slideIndex int32, dto IShapeBase, shapeToClone *int32, position *int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if shapeToClone != nil {
		if err := typeCheckParameter(*shapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if shapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*shapeToClone, ""))
	}
	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create a slide.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "layoutAlias" (string) Alias of layout slide for new slide. Alias may be the type of layout, name of layout slide or index
     @param "position" (int32) The target position at which to create the slide. Add to the end by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) CreateSlide(name string, layoutAlias string, position *int32, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(layoutAlias, "string", "layoutAlias"); err != nil {
		return successPayload, nil, err
	}
	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if localVarTempParam := layoutAlias; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("LayoutAlias", parameterToString(localVarTempParam, ""))
	}
	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create new shape (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) CreateSubshape(name string, slideIndex int32, path string, dto IShapeBase, shapeToClone *int32, position *int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if shapeToClone != nil {
		if err := typeCheckParameter(*shapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if shapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*shapeToClone, ""))
	}
	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new paragraph (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) CreateSubshapeParagraph(name string, slideIndex int32, path string, shapeIndex int32, dto IParagraph, position *int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Creates new portion (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) CreateSubshapePortion(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, dto IPortion, position *int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if position != nil {
		if err := typeCheckParameter(*position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
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

	if position != nil {
		localVarQueryParams.Add("Position", parameterToString(*position, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Adds a text watermark to each slide of the presentation. Text watermark can be setup via method arguments or withing Shape DTO for detailed customization. Both options are applicable simultaneously. 
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shape" (Shape) Shape DTO
     @param "fontHeight" (float64) Watermark font height.
     @param "text" (string) Watermark text.
     @param "fontName" (string) Watermark font name.
     @param "fontColor" (string) Watermark font color.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) CreateWatermark(name string, shape IShape, fontHeight *float64, text string, fontName string, fontColor string, password string, folder string, storage string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/watermark"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if fontHeight != nil {
		if err := typeCheckParameter(*fontHeight, "float64", "fontHeight"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(text, "string", "text"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontName, "string", "fontName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontColor, "string", "fontColor"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}

	if fontHeight != nil {
		localVarQueryParams.Add("FontHeight", parameterToString(*fontHeight, ""))
	}
	if localVarTempParam := text; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Text", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontColor; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontColor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &shape
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Adds a text watermark to each slide of the presentation. Text watermark can be setup via method arguments or withing Shape DTO for detailed customization. Both options are applicable simultaneously. 
 @param document Document data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shape" (Shape) Shape DTO
     @param "fontHeight" (float64) Watermark font height.
     @param "text" (string) Watermark text.
     @param "fontName" (string) Watermark font name.
     @param "fontColor" (string) Watermark font color.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) CreateWatermarkOnline(document []byte, shape IShape, fontHeight *float64, text string, fontName string, fontColor string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/watermark"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if fontHeight != nil {
		if err := typeCheckParameter(*fontHeight, "float64", "fontHeight"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(text, "string", "text"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontName, "string", "fontName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(fontColor, "string", "fontColor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	if fontHeight != nil {
		localVarQueryParams.Add("FontHeight", parameterToString(*fontHeight, ""))
	}
	if localVarTempParam := text; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Text", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := fontColor; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontColor", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarPostBody = &shape
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Remove animation from a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimation(name string, slideIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove an effect from slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param effectIndex Index of the effect to be removed.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimationEffect(name string, slideIndex int32, effectIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", effectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove an interactive sequence from slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param sequenceIndex The index of an interactive sequence to be deleted.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimationInteractiveSequence(name string, slideIndex int32, sequenceIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", sequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove an effect from slide animation interactive sequence.
 @param name Document name.
 @param slideIndex Slide index.
 @param sequenceIndex Interactive sequence index.
 @param effectIndex Index of the effect to be removed.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimationInteractiveSequenceEffect(name string, slideIndex int32, sequenceIndex int32, effectIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", sequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", effectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Clear all interactive sequences from slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimationInteractiveSequences(name string, slideIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Clear main sequence in slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) DeleteAnimationMainSequence(name string, slideIndex int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove background from a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) DeleteBackground(name string, slideIndex int32, password string, folder string, storage string) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideBackground", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideBackground); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideBackground); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete a category from a chart.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param categoryIndex Category index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) DeleteChartCategory(name string, slideIndex int32, shapeIndex int32, categoryIndex int32, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories/{categoryIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	categoryIndexPathStringValue := fmt.Sprintf("%v", categoryIndex)
	if len(categoryIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"categoryIndex"+"}", categoryIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"categoryIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete a data point from a chart series.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param seriesIndex Series index.
 @param pointIndex Data point index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) DeleteChartDataPoint(name string, slideIndex int32, shapeIndex int32, seriesIndex int32, pointIndex int32, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints/{pointIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", seriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}
	pointIndexPathStringValue := fmt.Sprintf("%v", pointIndex)
	if len(pointIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"pointIndex"+"}", pointIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"pointIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete a series from a chart.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index (must be a chart).
 @param seriesIndex Series index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) DeleteChartSeries(name string, slideIndex int32, shapeIndex int32, seriesIndex int32, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", seriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Removes comments of the specified author from the presentation. If author value is not provided all comments will be removed. 
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "author" (string) Author of comments.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) DeleteComments(name string, author string, password string, folder string, storage string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/comments"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(author, "string", "author"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}

	if localVarTempParam := author; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Author", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Removes comments of the specified author from the presentation. If author value is not provided all comments will be removed. 
 @param document Document data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "author" (string) Author of comments.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DeleteCommentsOnline(document []byte, author string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/comments/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(author, "string", "author"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := author; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Author", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Clean document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) DeleteDocumentProperties(name string, password string, folder string, storage string) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete document property.
 @param name Document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) DeleteDocumentProperty(name string, propertyName string, password string, folder string, storage string) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(propertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter propertyName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete file
 @param path File path e.g. &#39;/folder/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID to delete
 @return */
func (a *SlidesApiService) DeleteFile(path string, storageName string, versionId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(path) == 0 {
		return nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(versionId, "string", "versionId"); err != nil {
		return nil, err
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
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Delete folder
 @param path Folder path e.g. &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "recursive" (bool) Enable to delete folders, subfolders and files
 @return */
func (a *SlidesApiService) DeleteFolder(path string, storageName string, recursive *bool) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(path) == 0 {
		return nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if recursive != nil {
		if err := typeCheckParameter(*recursive, "bool", "recursive"); err != nil {
			return nil, err
		}
	}

	if localVarTempParam := storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if recursive != nil {
		localVarQueryParams.Add("Recursive", parameterToString(recursive, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Remove notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) DeleteNotesSlide(name string, slideIndex int32, password string, folder string, storage string) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteNotesSlideParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of paragraphs.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteNotesSlideParagraphs(name string, slideIndex int32, shapeIndex int32, paragraphs []int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(paragraphs, "[]int32", "paragraphs"); err != nil {
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

	if localVarTempParam := paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteNotesSlidePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of portions.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteNotesSlidePortions(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portions []int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(portions, "[]int32", "portions"); err != nil {
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

	if localVarTempParam := portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteNotesSlideShape(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of shapes.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteNotesSlideShapes(name string, slideIndex int32, shapes []int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(shapes, "[]int32", "shapes"); err != nil {
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

	if localVarTempParam := shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of paragraphs.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteParagraphs(name string, slideIndex int32, shapeIndex int32, paragraphs []int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(paragraphs, "[]int32", "paragraphs"); err != nil {
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

	if localVarTempParam := paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeletePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of portions.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeletePortions(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portions []int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(portions, "[]int32", "portions"); err != nil {
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

	if localVarTempParam := portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Resets all presentation protection settings. 
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Presentation password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ProtectionProperties*/
func (a *SlidesApiService) DeleteProtection(name string, password string, folder string, storage string) (IProtectionProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IProtectionProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/protection"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ProtectionProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IProtectionProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IProtectionProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Resets all presentation protection settings. 
 @param document Document data.
 @param password Presentation password.
 @return *os.File*/
func (a *SlidesApiService) DeleteProtectionOnline(document []byte, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(password) == 0 {
		return successPayload, nil, reportError("Missing required parameter password")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/protection/delete"

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
	localVarHeaderParams["Password"] = parameterToString(password, "")
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Delete a presentation section.
 @param name Document name.
 @param sectionIndex Section index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withSlides" (bool) True to delete the slides related to the deleted section; move them to the remaining sections otherwise.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) DeleteSection(name string, sectionIndex int32, withSlides *bool, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", sectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if withSlides != nil {
		if err := typeCheckParameter(*withSlides, "bool", "withSlides"); err != nil {
			return successPayload, nil, err
		}
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

	if withSlides != nil {
		localVarQueryParams.Add("WithSlides", parameterToString(withSlides, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete presentation sections.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "sections" ([]int32) The indices of the sections to be deleted; delete all by default.
     @param "withSlides" (bool) True to delete the slides related to the deleted sections; move them to the remaining sections otherwise.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) DeleteSections(name string, sections []int32, withSlides *bool, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(sections, "[]int32", "sections"); err != nil {
		return successPayload, nil, err
	}
	if withSlides != nil {
		if err := typeCheckParameter(*withSlides, "bool", "withSlides"); err != nil {
			return successPayload, nil, err
		}
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

	if localVarTempParam := sections; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Sections", parameterToString(localVarTempParam, ""))
	}
	if withSlides != nil {
		localVarQueryParams.Add("WithSlides", parameterToString(withSlides, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteShape(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of shapes.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteShapes(name string, slideIndex int32, shapes []int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(shapes, "[]int32", "shapes"); err != nil {
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

	if localVarTempParam := shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Delete a presentation slide by index.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) DeleteSlide(name string, slideIndex int32, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Removes comments of the specified author from the slide. If author value is not provided all comments will be removed. 
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "author" (string) Author of comments.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideComments*/
func (a *SlidesApiService) DeleteSlideComments(name string, slideIndex int32, author string, password string, folder string, storage string) (ISlideComments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideComments
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/comments"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(author, "string", "author"); err != nil {
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

	if localVarTempParam := author; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Author", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideComments", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideComments); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideComments); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Removes comments of the specified author from the slide. If author value is not provided all comments will be removed.              
 @param document Document data.
 @param slideIndex 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "author" (string) Author of comments.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DeleteSlideCommentsOnline(document []byte, slideIndex int32, author string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/comments/delete"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(author, "string", "author"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := author; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Author", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Delete presentation slides.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "slides" ([]int32) The indices of the slides to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) DeleteSlides(name string, slides []int32, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(slides, "[]int32", "slides"); err != nil {
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

	if localVarTempParam := slides; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Slides", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a shape (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteSubshape(name string, slideIndex int32, path string, shapeIndex int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a paragraph (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteSubshapeParagraph(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of paragraphs (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteSubshapeParagraphs(name string, slideIndex int32, path string, shapeIndex int32, paragraphs []int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(paragraphs, "[]int32", "paragraphs"); err != nil {
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

	if localVarTempParam := paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a portion (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteSubshapePortion(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of portions (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteSubshapePortions(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, portions []int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(portions, "[]int32", "portions"); err != nil {
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

	if localVarTempParam := portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Remove a range of shapes (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteSubshapes(name string, slideIndex int32, path string, shapes []int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(shapes, "[]int32", "shapes"); err != nil {
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

	if localVarTempParam := shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Removes shapes with name \&quot;watermark\&quot; from the presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapeName" (string) Name of the watermark shape. If null, default value \&quot;watermark\&quot;is used.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) DeleteWatermark(name string, shapeName string, password string, folder string, storage string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/watermark/delete"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(shapeName, "string", "shapeName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}

	if localVarTempParam := shapeName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("ShapeName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Removes shapes with name \&quot;watermark\&quot; from the presentation.
 @param document Document data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapeName" (string) Name of the watermark shape. If null, default value \&quot;watermark\&quot;is used.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DeleteWatermarkOnline(document []byte, shapeName string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/watermark/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(shapeName, "string", "shapeName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := shapeName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("ShapeName", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Download file
 @param path File path e.g. &#39;/folder/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID to download
 @return *os.File*/
func (a *SlidesApiService) DownloadFile(path string, storageName string, versionId string) (*os.File, *http.Response, error) {
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
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get image in specified format.
 @param name Document name.
 @param index Image index.
 @param format Export format (png, jpg, gif).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) DownloadImage(name string, index int32, format string, password string, folder string, storage string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get image binary data.
 @param name Document name.
 @param index Image index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) DownloadImageDefaultFormat(name string, index int32, password string, folder string, storage string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get image binary data.
 @param document Document data.
 @param index Image index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DownloadImageDefaultFormatOnline(document []byte, index int32, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/images/{index}"
	indexPathStringValue := fmt.Sprintf("%v", index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get image in specified format.
 @param document Document data.
 @param index Image index.
 @param format Export format (png, jpg, gif).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DownloadImageOnline(document []byte, index int32, format string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/images/{index}/{format}"
	indexPathStringValue := fmt.Sprintf("%v", index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get all presentation images in specified format.
 @param name 
 @param format Export format (png, jpg, gif).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) DownloadImages(name string, format string, password string, folder string, storage string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/download/{format}"
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get all presentation images.
 @param name 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) DownloadImagesDefaultFormat(name string, password string, folder string, storage string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/download"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get all presentation images.
 @param document Document data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DownloadImagesDefaultFormatOnline(document []byte, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/images/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Get all presentation images in specified format. 
 @param document Document data.
 @param format Export format (png, jpg, gif).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) DownloadImagesOnline(document []byte, format string, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/images/download/{format}"
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Convert notes slide to the specified image format.
 @param name Document name.
 @param slideIndex Slide index.
 @param format Output file format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Output file width.
     @param "height" (int32) Output file height.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return *os.File*/
func (a *SlidesApiService) DownloadNotesSlide(name string, slideIndex int32, format string, width *int32, height *int32, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Convert notes slide to the specified image format.
 @param document Document data.
 @param slideIndex Slide index.
 @param format Output file format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Output file width.
     @param "height" (int32) Output file height.
     @param "password" (string) Document password.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return *os.File*/
func (a *SlidesApiService) DownloadNotesSlideOnline(document []byte, slideIndex int32, format string, width *int32, height *int32, password string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
	if err := typeCheckParameter(password, "string", "password"); err != nil {
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
		"multipart/form-data",
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) DownloadNotesSlideShape(name string, slideIndex int32, shapeIndex int32, format string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
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

	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Save a presentation to a specified format.
 @param name Document name.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) DownloadPresentation(name string, format string, options IExportOptions, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) DownloadShape(name string, slideIndex int32, shapeIndex int32, format string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
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

	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Render shape to specified picture format.
 @param document Document data.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) DownloadShapeOnline(document []byte, slideIndex int32, shapeIndex int32, format string, scaleX *float64, scaleY *float64, bounds string, password string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
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

	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Save a slide to a specified format.
 @param name Document name.
 @param slideIndex Slide index.
 @param format Output file format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "width" (int32) Output file width; 0 to not adjust the size. Default is 0.
     @param "height" (int32) Output file height; 0 to not adjust the size. Default is 0.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return *os.File*/
func (a *SlidesApiService) DownloadSlide(name string, slideIndex int32, format string, options IExportOptions, width *int32, height *int32, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
		"multipart/form-data",
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Save a slide to a specified format.
 @param document Document data.
 @param slideIndex Slide index.
 @param format Output file format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Output file width; 0 to not adjust the size. Default is 0.
     @param "height" (int32) Output file height; 0 to not adjust the size. Default is 0.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return *os.File*/
func (a *SlidesApiService) DownloadSlideOnline(document []byte, slideIndex int32, format string, width *int32, height *int32, password string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
	if err := typeCheckParameter(password, "string", "password"); err != nil {
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
		"multipart/form-data",
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Render shape to specified picture format (for smart art and group shapes).
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) DownloadSubshape(name string, slideIndex int32, path string, shapeIndex int32, format string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
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

	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
		"multipart/form-data",
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Read slide animation effects.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapeIndex" (int32) Shape index. If specified, only effects related to that shape are returned.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) GetAnimation(name string, slideIndex int32, shapeIndex *int32, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if shapeIndex != nil {
		if err := typeCheckParameter(*shapeIndex, "int32", "shapeIndex"); err != nil {
			return successPayload, nil, err
		}
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

	if shapeIndex != nil {
		localVarQueryParams.Add("ShapeIndex", parameterToString(*shapeIndex, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get API info.
 @return ApiInfo*/
func (a *SlidesApiService) GetApiInfo() (IApiInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IApiInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/info"

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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ApiInfo", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IApiInfo); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IApiInfo); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide background info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) GetBackground(name string, slideIndex int32, password string, folder string, storage string) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideBackground", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideBackground); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideBackground); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide theme color scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ColorScheme*/
func (a *SlidesApiService) GetColorScheme(name string, slideIndex int32, password string, folder string, storage string) (IColorScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IColorScheme
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/colorScheme"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ColorScheme", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IColorScheme); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IColorScheme); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get disc usage
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return DiscUsage*/
func (a *SlidesApiService) GetDiscUsage(storageName string) (IDiscUsage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDiscUsage
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/disc"

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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DiscUsage", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDiscUsage); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDiscUsage); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) GetDocumentProperties(name string, password string, folder string, storage string) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation document property.
 @param name Document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperty*/
func (a *SlidesApiService) GetDocumentProperty(name string, propertyName string, password string, folder string, storage string) (IDocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperty
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(propertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter propertyName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperty", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperty); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperty); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get file versions
 @param path File path e.g. &#39;/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FileVersions*/
func (a *SlidesApiService) GetFileVersions(path string, storageName string) (IFileVersions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFileVersions
	)

	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/version/{path}"
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("FileVersions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IFileVersions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IFileVersions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get all files and folders within a folder
 @param path Folder path e.g. &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FilesList*/
func (a *SlidesApiService) GetFilesList(path string, storageName string) (IFilesList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFilesList
	)

	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("FilesList", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IFilesList); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IFilesList); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide theme font scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return FontScheme*/
func (a *SlidesApiService) GetFontScheme(name string, slideIndex int32, password string, folder string, storage string) (IFontScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFontScheme
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/fontScheme"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("FontScheme", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IFontScheme); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IFontScheme); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide theme format scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return FormatScheme*/
func (a *SlidesApiService) GetFormatScheme(name string, slideIndex int32, password string, folder string, storage string) (IFormatScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFormatScheme
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/formatScheme"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("FormatScheme", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IFormatScheme); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IFormatScheme); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation layoutSlide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) GetLayoutSlide(name string, slideIndex int32, password string, folder string, storage string) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("LayoutSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ILayoutSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ILayoutSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation layoutSlides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlides*/
func (a *SlidesApiService) GetLayoutSlides(name string, password string, folder string, storage string) (ILayoutSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("LayoutSlides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ILayoutSlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ILayoutSlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation masterSlide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return MasterSlide*/
func (a *SlidesApiService) GetMasterSlide(name string, slideIndex int32, password string, folder string, storage string) (IMasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("MasterSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IMasterSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IMasterSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation masterSlides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return MasterSlides*/
func (a *SlidesApiService) GetMasterSlides(name string, password string, folder string, storage string) (IMasterSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("MasterSlides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IMasterSlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IMasterSlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read notes slide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) GetNotesSlide(name string, slideIndex int32, password string, folder string, storage string) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get header/footer info for the notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlideHeaderFooter*/
func (a *SlidesApiService) GetNotesSlideHeaderFooter(name string, slideIndex int32, password string, folder string, storage string) (INotesSlideHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlideHeaderFooter
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlideHeaderFooter", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlideHeaderFooter); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlideHeaderFooter); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read notes slide info.
 @param document Document data.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return NotesSlide*/
func (a *SlidesApiService) GetNotesSlideOnline(document []byte, slideIndex int32, password string) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraph info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) GetNotesSlideParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraphs info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) GetNotesSlideParagraphs(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portion info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) GetNotesSlidePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portions info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) GetNotesSlidePortions(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shape info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) GetNotesSlideShape(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shapes info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) GetNotesSlideShapes(name string, slideIndex int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraph info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) GetParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraphs info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) GetParagraphs(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide placeholder info.
 @param name Document name.
 @param slideIndex Slide index.
 @param placeholderIndex Placeholder index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Placeholder*/
func (a *SlidesApiService) GetPlaceholder(name string, slideIndex int32, placeholderIndex int32, password string, folder string, storage string) (IPlaceholder, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPlaceholder
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders/{placeholderIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	placeholderIndexPathStringValue := fmt.Sprintf("%v", placeholderIndex)
	if len(placeholderIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"placeholderIndex"+"}", placeholderIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"placeholderIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Placeholder", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPlaceholder); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPlaceholder); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide placeholders info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Placeholders*/
func (a *SlidesApiService) GetPlaceholders(name string, slideIndex int32, password string, folder string, storage string) (IPlaceholders, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPlaceholders
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Placeholders", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPlaceholders); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPlaceholders); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portion info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) GetPortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portions info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) GetPortions(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) GetPresentation(name string, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation images info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Images*/
func (a *SlidesApiService) GetPresentationImages(name string, password string, folder string, storage string) (IImages, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IImages
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Images", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IImages); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IImages); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Extract presentation text items.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withEmpty" (bool) True to incude empty items.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return TextItems*/
func (a *SlidesApiService) GetPresentationTextItems(name string, withEmpty *bool, password string, folder string, storage string) (ITextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITextItems
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/textItems"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if withEmpty != nil {
		if err := typeCheckParameter(*withEmpty, "bool", "withEmpty"); err != nil {
			return successPayload, nil, err
		}
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

	if withEmpty != nil {
		localVarQueryParams.Add("WithEmpty", parameterToString(withEmpty, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("TextItems", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ITextItems); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ITextItems); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation protection properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ProtectionProperties*/
func (a *SlidesApiService) GetProtectionProperties(name string, password string, folder string, storage string) (IProtectionProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IProtectionProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/protection"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ProtectionProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IProtectionProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IProtectionProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation sections info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) GetSections(name string, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shape info.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) GetShape(name string, slideIndex int32, shapeIndex int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shapes info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) GetShapes(name string, slideIndex int32, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation slide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) GetSlide(name string, slideIndex int32, password string, folder string, storage string) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation slide comments.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideComments*/
func (a *SlidesApiService) GetSlideComments(name string, slideIndex int32, password string, folder string, storage string) (ISlideComments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideComments
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/comments"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideComments", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideComments); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideComments); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get footer info for the slide.
 @param name Document name.
 @param slideIndex The position of the slide to be reordered.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return HeaderFooter*/
func (a *SlidesApiService) GetSlideHeaderFooter(name string, slideIndex int32, password string, folder string, storage string) (IHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IHeaderFooter
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("HeaderFooter", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IHeaderFooter); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IHeaderFooter); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide images info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Images*/
func (a *SlidesApiService) GetSlideImages(name string, slideIndex int32, password string, folder string, storage string) (IImages, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IImages
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/images"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Images", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IImages); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IImages); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation slide properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideProperties*/
func (a *SlidesApiService) GetSlideProperties(name string, password string, folder string, storage string) (ISlideProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideProperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Extract slide text items.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withEmpty" (bool) True to include empty items.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return TextItems*/
func (a *SlidesApiService) GetSlideTextItems(name string, slideIndex int32, withEmpty *bool, password string, folder string, storage string) (ITextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITextItems
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/textItems"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if withEmpty != nil {
		if err := typeCheckParameter(*withEmpty, "bool", "withEmpty"); err != nil {
			return successPayload, nil, err
		}
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

	if withEmpty != nil {
		localVarQueryParams.Add("WithEmpty", parameterToString(withEmpty, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("TextItems", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ITextItems); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ITextItems); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation slides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) GetSlides(name string, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shape info (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) GetSubshape(name string, slideIndex int32, path string, shapeIndex int32, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraph info (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) GetSubshapeParagraph(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read shape paragraphs info (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) GetSubshapeParagraphs(name string, slideIndex int32, path string, shapeIndex int32, password string, folder string, storage string) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraphs", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraphs); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraphs); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portion info (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) GetSubshapePortion(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, portionIndex int32, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read paragraph portions info (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) GetSubshapePortions(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, password string, folder string, storage string) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portions", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortions); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortions); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide shapes info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) GetSubshapes(name string, slideIndex int32, path string, password string, folder string, storage string) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Shapes", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapes); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapes); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read slide theme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Theme*/
func (a *SlidesApiService) GetTheme(name string, slideIndex int32, password string, folder string, storage string) (ITheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITheme
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Theme", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ITheme); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ITheme); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Read presentation document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ViewProperties*/
func (a *SlidesApiService) GetViewProperties(name string, password string, folder string, storage string) (IViewProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IViewProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/viewProperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ViewProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IViewProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IViewProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create presentation document from html.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "html" (string) HTML data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) ImportFromHtml(name string, html string, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromHtml"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(html, "string", "html"); err != nil {
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = html
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Create presentation document from pdf or append pdf to an existing presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pdf" ([]byte) PDF data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) ImportFromPdf(name string, pdf []byte, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromPdf"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(pdf) > 0 {
		localVarFiles = append(localVarFiles, pdf)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Merge the presentation with other presentations specified in the request parameter.
 @param name Document name.
 @param request PresentationsMergeRequest with a list of presentations to merge.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) Merge(name string, request IPresentationsMergeRequest, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if request == nil {
		return successPayload, nil, reportError("Missing required parameter request")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Merges presentations or some of their slides specified in the request parameter. Result will be save in the storage.
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) Merge request.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) MergeAndSaveOnline(outPath string, files [][]byte, request IOrderedMergeRequest, storage string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/merge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Merges presentations or some of their slides specified in the request parameter. Returns result file in the response. 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) Merge request.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) MergeOnline(files [][]byte, request IOrderedMergeRequest, storage string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/merge"

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
		"multipart/form-data",
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
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Move file
 @param srcPath Source file path e.g. &#39;/src.ext&#39;
 @param destPath Destination file path e.g. &#39;/dest.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
     @param "versionId" (string) File version ID to move
 @return */
func (a *SlidesApiService) MoveFile(srcPath string, destPath string, srcStorageName string, destStorageName string, versionId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(srcPath) == 0 {
		return nil, reportError("Missing required parameter srcPath")
	}
	if len(destPath) == 0 {
		return nil, reportError("Missing required parameter destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(versionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam := srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Move folder
 @param srcPath Folder path to move e.g. &#39;/folder&#39;
 @param destPath Destination folder path to move to e.g &#39;/dst&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *SlidesApiService) MoveFolder(srcPath string, destPath string, srcStorageName string, destStorageName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(srcPath) == 0 {
		return nil, reportError("Missing required parameter srcPath")
	}
	if len(destPath) == 0 {
		return nil, reportError("Missing required parameter destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam := srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Move presentation section to a specified position.
 @param name Document name.
 @param sectionIndex The position of the section to be reordered.
 @param newPosition The new position of the reordered section.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) MoveSection(name string, sectionIndex int32, newPosition int32, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}/move"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", sectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
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

	localVarQueryParams.Add("NewPosition", parameterToString(newPosition, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Reorder presentation slide position.
 @param name Document name.
 @param slideIndex The position of the slide to be reordered.
 @param newPosition The new position of the reordered slide.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) MoveSlide(name string, slideIndex int32, newPosition int32, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/move"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	localVarQueryParams.Add("NewPosition", parameterToString(newPosition, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get info whether a notes slide exists.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return EntityExists*/
func (a *SlidesApiService) NotesSlideExists(name string, slideIndex int32, password string, folder string, storage string) (IEntityExists, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IEntityExists
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/exist"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("EntityExists", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IEntityExists); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IEntityExists); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Get info whether a notes slide exists.
 @param document Document data.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return EntityExists*/
func (a *SlidesApiService) NotesSlideExistsOnline(document []byte, slideIndex int32, password string) (IEntityExists, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IEntityExists
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide/exist"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("EntityExists", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IEntityExists); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IEntityExists); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Check if file or folder exists
 @param path File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID
 @return ObjectExist*/
func (a *SlidesApiService) ObjectExists(path string, storageName string, versionId string) (IObjectExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IObjectExist
	)

	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/exist/{path}"
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
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ObjectExist", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IObjectExist); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IObjectExist); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Merge the presentation with other presentations or some of their slides specified in the request parameter.
 @param name Document name.
 @param request OrderedMergeRequest with a list of presentations and slide indices to merge.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) OrderedMerge(name string, request IOrderedMergeRequest, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if request == nil {
		return successPayload, nil, reportError("Missing required parameter request")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Performs slides pipeline.
 @param pipeline A Pipeline object.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to upload with the pipeline
 @return *os.File*/
func (a *SlidesApiService) Pipeline(pipeline IPipeline, files [][]byte) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if pipeline == nil {
		return successPayload, nil, reportError("Missing required parameter pipeline")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/pipeline"

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
	localVarFiles = files
	localVarPostBody = &pipeline
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Reorder presentation slides positions.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "oldPositions" ([]int32) A comma separated array of positions of slides to be reordered.
     @param "newPositions" ([]int32) A comma separated array of new slide positions.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) ReorderSlides(name string, oldPositions []int32, newPositions []int32, password string, folder string, storage string) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/reorder"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(oldPositions, "[]int32", "oldPositions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(newPositions, "[]int32", "newPositions"); err != nil {
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

	if localVarTempParam := oldPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("OldPositions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := newPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("NewPositions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slides", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlides); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlides); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Replace text with a new value.
 @param name Document name.
 @param oldValue Text value to be replaced.
 @param newValue Text value to replace with.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "ignoreCase" (bool) True if character case must be ignored.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentReplaceResult*/
func (a *SlidesApiService) ReplacePresentationText(name string, oldValue string, newValue string, ignoreCase *bool, password string, folder string, storage string) (IDocumentReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentReplaceResult
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(oldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter oldValue")
	}
	if len(newValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/replaceText"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if ignoreCase != nil {
		if err := typeCheckParameter(*ignoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
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

	localVarQueryParams.Add("OldValue", parameterToString(oldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(newValue, ""))
	if ignoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(ignoreCase, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentReplaceResult", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentReplaceResult); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentReplaceResult); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Replace text with a new value.
 @param document Document data.
 @param oldValue Text value to be replaced.
 @param newValue Text value to replace with.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "ignoreCase" (bool) True if character case must be ignored.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) ReplacePresentationTextOnline(document []byte, oldValue string, newValue string, ignoreCase *bool, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(oldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter oldValue")
	}
	if len(newValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/replaceText"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if ignoreCase != nil {
		if err := typeCheckParameter(*ignoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OldValue", parameterToString(oldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(newValue, ""))
	if ignoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(ignoreCase, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Replace text with a new value.
 @param name Document name.
 @param slideIndex Slide index.
 @param oldValue Text value to be replaced.
 @param newValue Text value to replace with.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "ignoreCase" (bool) True if character case must be ignored.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideReplaceResult*/
func (a *SlidesApiService) ReplaceSlideText(name string, slideIndex int32, oldValue string, newValue string, ignoreCase *bool, password string, folder string, storage string) (ISlideReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideReplaceResult
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(oldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter oldValue")
	}
	if len(newValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/replaceText"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if ignoreCase != nil {
		if err := typeCheckParameter(*ignoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
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

	localVarQueryParams.Add("OldValue", parameterToString(oldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(newValue, ""))
	if ignoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(ignoreCase, ""))
	}
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideReplaceResult", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideReplaceResult); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideReplaceResult); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Replace slide text with a new value.
 @param document Document data.
 @param slideIndex Index of target slide.
 @param oldValue Text value to be replaced.
 @param newValue Text value to replace with.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "ignoreCase" (bool) True if character case must be ignored.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) ReplaceSlideTextOnline(document []byte, slideIndex int32, oldValue string, newValue string, ignoreCase *bool, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(oldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter oldValue")
	}
	if len(newValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/replaceText"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if ignoreCase != nil {
		if err := typeCheckParameter(*ignoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OldValue", parameterToString(oldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(newValue, ""))
	if ignoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(ignoreCase, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) SaveNotesSlideShape(name string, slideIndex int32, shapeIndex int32, format string, outPath string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Save a presentation to a specified format.
 @param name Document name.
 @param format Export format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return */
func (a *SlidesApiService) SavePresentation(name string, format string, outPath string, options IExportOptions, password string, folder string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
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
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) SaveShape(name string, slideIndex int32, shapeIndex int32, format string, outPath string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Render shape to specified picture format.
 @param document Document data.
 @param slideIndex Slide index.
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) SaveShapeOnline(document []byte, slideIndex int32, shapeIndex int32, format string, outPath string, scaleX *float64, scaleY *float64, bounds string, password string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(document) == 0 {
		return nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Save a slide to a specified format.
 @param name Document name.
 @param slideIndex Slide index.
 @param format Output file format.
 @param outPath Path to upload the output file to.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "width" (int32) Output file width; 0 to not adjust the size. Default is 0.
     @param "height" (int32) Output file height; 0 to not adjust the size. Default is 0.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return */
func (a *SlidesApiService) SaveSlide(name string, slideIndex int32, format string, outPath string, options IExportOptions, width *int32, height *int32, password string, folder string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
			return nil, err
		}
	}
	if height != nil {
		if err := typeCheckParameter(*height, "int32", "height"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if width != nil {
		localVarQueryParams.Add("Width", parameterToString(*width, ""))
	}
	if height != nil {
		localVarQueryParams.Add("Height", parameterToString(*height, ""))
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Save a slide to a specified format.
 @param document Document data.
 @param slideIndex Slide index.
 @param format Output file format.
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Output file width; 0 to not adjust the size. Default is 0.
     @param "height" (int32) Output file height; 0 to not adjust the size. Default is 0.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Storage folder containing custom fonts to be used with the document.
 @return */
func (a *SlidesApiService) SaveSlideOnline(document []byte, slideIndex int32, format string, outPath string, width *int32, height *int32, password string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(document) == 0 {
		return nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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
			return nil, err
		}
	}
	if height != nil {
		if err := typeCheckParameter(*height, "int32", "height"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if width != nil {
		localVarQueryParams.Add("Width", parameterToString(*width, ""))
	}
	if height != nil {
		localVarQueryParams.Add("Height", parameterToString(*height, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) SaveSubshape(name string, slideIndex int32, path string, shapeIndex int32, format string, outPath string, options IIShapeExportOptions, scaleX *float64, scaleY *float64, bounds string, password string, folder string, storage string, fontsFolder string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(name) == 0 {
		return nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return nil, reportError("Missing required parameter path")
	}
	if len(format) == 0 {
		return nil, reportError("Missing required parameter format")
	}
	if len(outPath) == 0 {
		return nil, reportError("Missing required parameter outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if scaleX != nil {
		if err := typeCheckParameter(*scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if scaleY != nil {
		if err := typeCheckParameter(*scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(outPath, ""))
	if scaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*scaleX, ""))
	}
	if scaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*scaleY, ""))
	}
	if localVarTempParam := bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(string(responseBytes))
	}


	return localVarHttpResponse, err
}

/* SlidesApiService Set slide animation.
 @param name Document name.
 @param slideIndex Slide index.
 @param animation Animation DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) SetAnimation(name string, slideIndex int32, animation ISlideAnimation, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if animation == nil {
		return successPayload, nil, reportError("Missing required parameter animation")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &animation
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set background for a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param background Slide background update data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) SetBackground(name string, slideIndex int32, background ISlideBackground, password string, folder string, storage string) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if background == nil {
		return successPayload, nil, reportError("Missing required parameter background")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &background
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideBackground", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideBackground); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideBackground); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set background color for a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param color Slide background target color in RRGGBB format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) SetBackgroundColor(name string, slideIndex int32, color string, password string, folder string, storage string) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(color) == 0 {
		return successPayload, nil, reportError("Missing required parameter color")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/backgroundColor"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	localVarQueryParams.Add("Color", parameterToString(color, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideBackground", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideBackground); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideBackground); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set document properties.
 @param name Document name.
 @param properties New properties.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) SetDocumentProperties(name string, properties IDocumentProperties, password string, folder string, storage string) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if properties == nil {
		return successPayload, nil, reportError("Missing required parameter properties")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &properties
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set document property.
 @param name Document name.
 @param propertyName The property name.
 @param property Property with the value.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperty*/
func (a *SlidesApiService) SetDocumentProperty(name string, propertyName string, property IDocumentProperty, password string, folder string, storage string) (IDocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperty
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(propertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter propertyName")
	}
	if property == nil {
		return successPayload, nil, reportError("Missing required parameter property")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &property
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("DocumentProperty", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocumentProperty); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocumentProperty); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set header/footer the notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto Header/footer to set.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlideHeaderFooter*/
func (a *SlidesApiService) SetNotesSlideHeaderFooter(name string, slideIndex int32, dto INotesSlideHeaderFooter, password string, folder string, storage string) (INotesSlideHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlideHeaderFooter
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlideHeaderFooter", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlideHeaderFooter); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlideHeaderFooter); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set footers for all slides in a presentation.
 @param name Document name.
 @param dto HeaderFooter instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Document*/
func (a *SlidesApiService) SetPresentationHeaderFooter(name string, dto IHeaderFooter, password string, folder string, storage string) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Document", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IDocument); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IDocument); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Updates presentation protection properties.
 @param name Document name.
 @param dto Protection properties.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ProtectionProperties*/
func (a *SlidesApiService) SetProtection(name string, dto IProtectionProperties, password string, folder string, storage string) (IProtectionProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IProtectionProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/protection"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ProtectionProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IProtectionProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IProtectionProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Sets presentation protection options. 
 @param document Document data.
 @param dto Protection properties.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) SetProtectionOnline(document []byte, dto IProtectionProperties, password string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/protection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	if len(document) > 0 {
		localVarFiles = append(localVarFiles, document)
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Replace existing presentation sections with the ones provided in the sections DTO.
 @param name Document name.
 @param sections Sections DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) SetSections(name string, sections ISections, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if sections == nil {
		return successPayload, nil, reportError("Missing required parameter sections")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &sections
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Set footer the slide.
 @param name Document name.
 @param slideIndex The position of the slide to be reordered.
 @param dto Footer to set.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return HeaderFooter*/
func (a *SlidesApiService) SetSlideHeaderFooter(name string, slideIndex int32, dto IHeaderFooter, password string, folder string, storage string) (IHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IHeaderFooter
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("HeaderFooter", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IHeaderFooter); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IHeaderFooter); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update presentation slide properties.
 @param name Document name.
 @param dto The view properties data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideProperties*/
func (a *SlidesApiService) SetSlideProperties(name string, dto ISlideProperties, password string, folder string, storage string) (ISlideProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideProperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update presentation document properties.
 @param name Document name.
 @param dto The view properties data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ViewProperties*/
func (a *SlidesApiService) SetViewProperties(name string, dto IViewProperties, password string, folder string, storage string) (IViewProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IViewProperties
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/viewProperties"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ViewProperties", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IViewProperties); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IViewProperties); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Splitting presentations. Create one image per slide.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "format" (string) Export format. Default value is jpeg.
     @param "width" (int32) The width of created images.
     @param "height" (int32) The height of created images.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "destFolder" (string) Folder on storage where images are going to be uploaded. If not specified then images are uploaded to same folder as presentation.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return SplitDocumentResult*/
func (a *SlidesApiService) Split(name string, options IExportOptions, format string, width *int32, height *int32, from *int32, to *int32, destFolder string, password string, folder string, storage string, fontsFolder string) (ISplitDocumentResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISplitDocumentResult
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/split"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	if err := typeCheckParameter(format, "string", "format"); err != nil {
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

	if localVarTempParam := format; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Format", parameterToString(localVarTempParam, ""))
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SplitDocumentResult", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISplitDocumentResult); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISplitDocumentResult); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Splits PowerPoint presentation slides from the specified range into separate files and exports them in the specified file format. If the range is not provided all slides will be processed. 
 @param document Document data.
 @param format ExportFormat
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "destFolder" (string) Folder on storage where images are going to be uploaded. If not specified then images are uploaded to the root folder.
     @param "width" (int32) Slide width.
     @param "height" (int32) Slide height.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return SplitDocumentResult*/
func (a *SlidesApiService) SplitAndSaveOnline(document []byte, format string, destFolder string, width *int32, height *int32, from *int32, to *int32, password string, storage string, fontsFolder string) (ISplitDocumentResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISplitDocumentResult
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/split/{format}"
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SplitDocumentResult", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISplitDocumentResult); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISplitDocumentResult); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Splits PowerPoint presentation slides from the specified range into separate files and exports them in the specified file format. If the range is not provided all slides will be processed. 
 @param document Document data.
 @param format ExportFormat
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) Slide width.
     @param "height" (int32) Slide height.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) SplitOnline(document []byte, format string, width *int32, height *int32, from *int32, to *int32, password string, storage string, fontsFolder string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(document) == 0 {
		return successPayload, nil, reportError("Missing required parameter document")
	}
	if len(format) == 0 {
		return successPayload, nil, reportError("Missing required parameter format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/split/{format}"
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
	if err := typeCheckParameter(password, "string", "password"); err != nil {
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
		"multipart/form-data",
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
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

/* SlidesApiService Check if storage exists
 @param storageName Storage name
 @return StorageExist*/
func (a *SlidesApiService) StorageExists(storageName string) (IStorageExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IStorageExist
	)

	if len(storageName) == 0 {
		return successPayload, nil, reportError("Missing required parameter storageName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/{storageName}/exist"
	storageNamePathStringValue := fmt.Sprintf("%v", storageName)
	if len(storageNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", storageNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"storageName"+"}", "", -1)
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
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("StorageExist", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IStorageExist); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IStorageExist); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Modify an animation effect for a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param effectIndex The position of the effect to be modified.
 @param effect Animation effect DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) UpdateAnimationEffect(name string, slideIndex int32, effectIndex int32, effect IEffect, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if effect == nil {
		return successPayload, nil, reportError("Missing required parameter effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", effectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Modify an animation effect for a slide interactive sequence.
 @param name Document name.
 @param slideIndex Slide index.
 @param sequenceIndex The position of the interactive sequence.
 @param effectIndex The position of the effect to be modified.
 @param effect Animation effect DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideAnimation*/
func (a *SlidesApiService) UpdateAnimationInteractiveSequenceEffect(name string, slideIndex int32, sequenceIndex int32, effectIndex int32, effect IEffect, password string, folder string, storage string) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if effect == nil {
		return successPayload, nil, reportError("Missing required parameter effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", sequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", effectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("SlideAnimation", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlideAnimation); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlideAnimation); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update a chart category.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param categoryIndex Category index.
 @param category Category DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) UpdateChartCategory(name string, slideIndex int32, shapeIndex int32, categoryIndex int32, category IChartCategory, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if category == nil {
		return successPayload, nil, reportError("Missing required parameter category")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories/{categoryIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	categoryIndexPathStringValue := fmt.Sprintf("%v", categoryIndex)
	if len(categoryIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"categoryIndex"+"}", categoryIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"categoryIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &category
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update a data point in a chart series.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param seriesIndex Series index.
 @param pointIndex Data point index.
 @param dataPoint Data point DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) UpdateChartDataPoint(name string, slideIndex int32, shapeIndex int32, seriesIndex int32, pointIndex int32, dataPoint IDataPoint, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dataPoint == nil {
		return successPayload, nil, reportError("Missing required parameter dataPoint")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints/{pointIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", seriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}
	pointIndexPathStringValue := fmt.Sprintf("%v", pointIndex)
	if len(pointIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"pointIndex"+"}", pointIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"pointIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dataPoint
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update a series in a chart.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index (must be a chart).
 @param seriesIndex Series index.
 @param series Series DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Chart*/
func (a *SlidesApiService) UpdateChartSeries(name string, slideIndex int32, shapeIndex int32, seriesIndex int32, series ISeries, password string, folder string, storage string) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if series == nil {
		return successPayload, nil, reportError("Missing required parameter series")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", seriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &series
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Chart", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IChart); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IChart); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update a layoutSlide.
 @param name Document name.
 @param slideIndex Slide index.
 @param slideDto Slide update data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) UpdateLayoutSlide(name string, slideIndex int32, slideDto ILayoutSlide, password string, folder string, storage string) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if slideDto == nil {
		return successPayload, nil, reportError("Missing required parameter slideDto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &slideDto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("LayoutSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ILayoutSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ILayoutSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update notes slide properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto A NotesSlide object with notes slide data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) UpdateNotesSlide(name string, slideIndex int32, dto INotesSlide, password string, folder string, storage string) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("NotesSlide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(INotesSlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(INotesSlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update paragraph properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) UpdateNotesSlideParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, dto IParagraph, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update portion properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) UpdateNotesSlidePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, dto IPortion, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update shape properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param dto Shape DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) UpdateNotesSlideShape(name string, slideIndex int32, shapeIndex int32, dto IShapeBase, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update paragraph properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) UpdateParagraph(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, dto IParagraph, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update portion properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) UpdatePortion(name string, slideIndex int32, shapeIndex int32, paragraphIndex int32, portionIndex int32, dto IPortion, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update section name.
 @param name Document name.
 @param sectionIndex The position of the section to be updated.
 @param sectionName Section name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Sections*/
func (a *SlidesApiService) UpdateSection(name string, sectionIndex int32, sectionName string, password string, folder string, storage string) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(sectionName) == 0 {
		return successPayload, nil, reportError("Missing required parameter sectionName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", sectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
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

	localVarQueryParams.Add("SectionName", parameterToString(sectionName, ""))
	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Sections", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISections); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISections); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update shape properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param dto Shape DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) UpdateShape(name string, slideIndex int32, shapeIndex int32, dto IShapeBase, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param slideDto Slide update data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) UpdateSlide(name string, slideIndex int32, slideDto ISlide, password string, folder string, storage string) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if slideDto == nil {
		return successPayload, nil, reportError("Missing required parameter slideDto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &slideDto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Slide", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(ISlide); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(ISlide); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update shape properties (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param dto Shape DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) UpdateSubshape(name string, slideIndex int32, path string, shapeIndex int32, dto IShapeBase, password string, folder string, storage string) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("ShapeBase", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IShapeBase); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IShapeBase); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update paragraph properties (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) UpdateSubshapeParagraph(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, dto IParagraph, password string, folder string, storage string) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Paragraph", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IParagraph); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IParagraph); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Update portion properties (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) UpdateSubshapePortion(name string, slideIndex int32, path string, shapeIndex int32, paragraphIndex int32, portionIndex int32, dto IPortion, password string, folder string, storage string) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(name) == 0 {
		return successPayload, nil, reportError("Missing required parameter name")
	}
	if len(path) == 0 {
		return successPayload, nil, reportError("Missing required parameter path")
	}
	if dto == nil {
		return successPayload, nil, reportError("Missing required parameter dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", slideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", shapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", paragraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", portionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
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

	if localVarTempParam := folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarPostBody = &dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFiles)
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return successPayload, localVarHttpResponse, reportError(string(responseBytes))
	}

	successPayloadObject, err := createObjectForType("Portion", responseBytes)
	if err != nil {
		return successPayload, localVarHttpResponse, err
	}
	if err = json.NewDecoder(responseBody).Decode(successPayloadObject); err != nil {
		if sp, ok := successPayloadObject.(IPortion); ok {
			return sp, localVarHttpResponse, err
		}
		return successPayload, localVarHttpResponse, err
	}
	if successPayload, _ = successPayloadObject.(IPortion); true {
	}

	return successPayload, localVarHttpResponse, err
}

/* SlidesApiService Upload file
 @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.             
 @param file File to upload
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FilesUploadResult*/
func (a *SlidesApiService) UploadFile(path string, file []byte, storageName string) (IFilesUploadResult, *http.Response, error) {
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
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
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
	if localVarHttpResponse.StatusCode >= 300 {
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

