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

type SlidesApiService service


/* SlidesApiService Copy file
 @param srcPath Source file path e.g. &#39;/folder/file.ext&#39;
 @param destPath Destination file path
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
     @param "versionId" (string) File version ID to copy
 @return */
func (a *SlidesApiService) CopyFile(request CopyFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.versionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("destPath", parameterToString(request.destPath, ""))
	if localVarTempParam := request.srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.CopyFile
*/
type CopyFileRequest struct {
    srcPath string
    destPath string
    srcStorageName string
    destStorageName string
    versionId string
}

/* SlidesApiService Copy folder
 @param srcPath Source folder path e.g. &#39;/src&#39;
 @param destPath Destination folder path e.g. &#39;/dst&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *SlidesApiService) CopyFolder(request CopyFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("destPath", parameterToString(request.destPath, ""))
	if localVarTempParam := request.srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.CopyFolder
*/
type CopyFolderRequest struct {
    srcPath string
    destPath string
    srcStorageName string
    destStorageName string
}

/* SlidesApiService Create the folder
 @param path Folder path to create e.g. &#39;folder_1/folder_2/&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return */
func (a *SlidesApiService) CreateFolder(request CreateFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.CreateFolder
*/
type CreateFolderRequest struct {
    path string
    storageName string
}

/* SlidesApiService Delete file
 @param path File path e.g. &#39;/folder/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID to delete
 @return */
func (a *SlidesApiService) DeleteFile(request DeleteFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.versionId, "string", "versionId"); err != nil {
		return nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.DeleteFile
*/
type DeleteFileRequest struct {
    path string
    storageName string
    versionId string
}

/* SlidesApiService Delete folder
 @param path Folder path e.g. &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "recursive" (bool) Enable to delete folders, subfolders and files
 @return */
func (a *SlidesApiService) DeleteFolder(request DeleteFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if request.recursive != nil {
		if err := typeCheckParameter(*request.recursive, "bool", "recursive"); err != nil {
			return nil, err
		}
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("recursive", parameterToString(request.recursive, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.DeleteFolder
*/
type DeleteFolderRequest struct {
    path string
    storageName string
    recursive *bool
}

/* SlidesApiService Remove notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) DeleteNotesSlide(request DeleteNotesSlideRequest) (Slide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slide
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

/* Request for SlidesApiService.DeleteNotesSlide
*/
type DeleteNotesSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Remove a paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteNotesSlideParagraph(request DeleteNotesSlideParagraphRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
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

/* Request for SlidesApiService.DeleteNotesSlideParagraph
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

/* SlidesApiService Remove a range of paragraphs.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteNotesSlideParagraphs(request DeleteNotesSlideParagraphsRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
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

/* Request for SlidesApiService.DeleteNotesSlideParagraphs
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

/* SlidesApiService Remove a portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteNotesSlidePortion(request DeleteNotesSlidePortionRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
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

/* Request for SlidesApiService.DeleteNotesSlidePortion
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

/* SlidesApiService Remove a range of portions.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeleteNotesSlidePortions(request DeleteNotesSlidePortionsRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
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

/* Request for SlidesApiService.DeleteNotesSlidePortions
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

/* SlidesApiService Remove a shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteNotesSlideShape(request DeleteNotesSlideShapeRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
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

/* Request for SlidesApiService.DeleteNotesSlideShape
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

/* SlidesApiService Remove a range of shapes.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteNotesSlideShapes(request DeleteNotesSlideShapesRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
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

/* Request for SlidesApiService.DeleteNotesSlideShapes
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

/* SlidesApiService Remove a paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteParagraph(request DeleteParagraphRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
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

/* Request for SlidesApiService.DeleteParagraph
*/
type DeleteParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Remove a range of paragraphs.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraphs" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) DeleteParagraphs(request DeleteParagraphsRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
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

/* Request for SlidesApiService.DeleteParagraphs
*/
type DeleteParagraphsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphs []int32
    password string
    folder string
    storage string
}

/* SlidesApiService Remove a portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeletePortion(request DeletePortionRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
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

/* Request for SlidesApiService.DeletePortion
*/
type DeletePortionRequest struct {
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

/* SlidesApiService Remove a range of portions.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "portions" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) DeletePortions(request DeletePortionsRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
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

/* Request for SlidesApiService.DeletePortions
*/
type DeletePortionsRequest struct {
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

/* SlidesApiService Delete a presentation slide by index.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) DeleteSlideByIndex(request DeleteSlideByIndexRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
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

/* Request for SlidesApiService.DeleteSlideByIndex
*/
type DeleteSlideByIndexRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Remove a shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteSlideShape(request DeleteSlideShapeRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
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

/* Request for SlidesApiService.DeleteSlideShape
*/
type DeleteSlideShapeRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Remove a range of shapes.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "shapes" ([]int32) The indices of the shapes to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) DeleteSlideShapes(request DeleteSlideShapesRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
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

/* Request for SlidesApiService.DeleteSlideShapes
*/
type DeleteSlideShapesRequest struct {
    name string
    slideIndex int32
    path string
    shapes []int32
    password string
    folder string
    storage string
}

/* SlidesApiService Delete presentation slides.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "slides" ([]int32) The indices of the slides to be deleted; delete all by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) DeleteSlidesCleanSlidesList(request DeleteSlidesCleanSlidesListRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.slides, "[]int32", "slides"); err != nil {
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

	if localVarTempParam := request.slides; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("slides", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.DeleteSlidesCleanSlidesList
*/
type DeleteSlidesCleanSlidesListRequest struct {
    name string
    slides []int32
    password string
    folder string
    storage string
}

/* SlidesApiService Clean document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) DeleteSlidesDocumentProperties(request DeleteSlidesDocumentPropertiesRequest) (DocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperties
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
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

/* Request for SlidesApiService.DeleteSlidesDocumentProperties
*/
type DeleteSlidesDocumentPropertiesRequest struct {
    name string
    password string
    folder string
    storage string
}

/* SlidesApiService Delete document property.
 @param name Document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) DeleteSlidesDocumentProperty(request DeleteSlidesDocumentPropertyRequest) (DocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperties
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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

/* Request for SlidesApiService.DeleteSlidesDocumentProperty
*/
type DeleteSlidesDocumentPropertyRequest struct {
    name string
    propertyName string
    password string
    folder string
    storage string
}

/* SlidesApiService Remove background from a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) DeleteSlidesSlideBackground(request DeleteSlidesSlideBackgroundRequest) (SlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideBackground
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
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

/* Request for SlidesApiService.DeleteSlidesSlideBackground
*/
type DeleteSlidesSlideBackgroundRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Download file
 @param path File path e.g. &#39;/folder/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID to download
 @return *os.File*/
func (a *SlidesApiService) DownloadFile(request DownloadFileRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.versionId, "string", "versionId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.DownloadFile
*/
type DownloadFileRequest struct {
    path string
    storageName string
    versionId string
}

/* SlidesApiService Get disc usage
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return DiscUsage*/
func (a *SlidesApiService) GetDiscUsage(request GetDiscUsageRequest) (DiscUsage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DiscUsage
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/disc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.GetDiscUsage
*/
type GetDiscUsageRequest struct {
    storageName string
}

/* SlidesApiService Get file versions
 @param path File path e.g. &#39;/file.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FileVersions*/
func (a *SlidesApiService) GetFileVersions(request GetFileVersionsRequest) (FileVersions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FileVersions
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/version/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.GetFileVersions
*/
type GetFileVersionsRequest struct {
    path string
    storageName string
}

/* SlidesApiService Get all files and folders within a folder
 @param path Folder path e.g. &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FilesList*/
func (a *SlidesApiService) GetFilesList(request GetFilesListRequest) (FilesList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.GetFilesList
*/
type GetFilesListRequest struct {
    path string
    storageName string
}

/* SlidesApiService Read presentation layoutSlide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) GetLayoutSlide(request GetLayoutSlideRequest) (LayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LayoutSlide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
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

/* Request for SlidesApiService.GetLayoutSlide
*/
type GetLayoutSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation layoutSlides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlides*/
func (a *SlidesApiService) GetLayoutSlidesList(request GetLayoutSlidesListRequest) (LayoutSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LayoutSlides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
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

/* Request for SlidesApiService.GetLayoutSlidesList
*/
type GetLayoutSlidesListRequest struct {
    name string
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation masterSlide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return MasterSlide*/
func (a *SlidesApiService) GetMasterSlide(request GetMasterSlideRequest) (MasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  MasterSlide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides/{slideIndex}"
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

/* Request for SlidesApiService.GetMasterSlide
*/
type GetMasterSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation masterSlides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return MasterSlides*/
func (a *SlidesApiService) GetMasterSlidesList(request GetMasterSlidesListRequest) (MasterSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  MasterSlides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
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

/* Request for SlidesApiService.GetMasterSlidesList
*/
type GetMasterSlidesListRequest struct {
    name string
    password string
    folder string
    storage string
}

/* SlidesApiService Read notes slide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) GetNotesSlide(request GetNotesSlideRequest) (NotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlide
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

/* Request for SlidesApiService.GetNotesSlide
*/
type GetNotesSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide shape info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) GetNotesSlideShape(request GetNotesSlideShapeRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
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

/* Request for SlidesApiService.GetNotesSlideShape
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

/* SlidesApiService Read shape paragraph info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) GetNotesSlideShapeParagraph(request GetNotesSlideShapeParagraphRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
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

/* Request for SlidesApiService.GetNotesSlideShapeParagraph
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

/* SlidesApiService Read shape paragraphs info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) GetNotesSlideShapeParagraphs(request GetNotesSlideShapeParagraphsRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
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

/* Request for SlidesApiService.GetNotesSlideShapeParagraphs
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

/* SlidesApiService Read paragraph portion info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) GetNotesSlideShapePortion(request GetNotesSlideShapePortionRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
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

/* Request for SlidesApiService.GetNotesSlideShapePortion
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

/* SlidesApiService Read paragraph portions info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) GetNotesSlideShapePortions(request GetNotesSlideShapePortionsRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
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

/* Request for SlidesApiService.GetNotesSlideShapePortions
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

/* SlidesApiService Read slide shapes info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Shapes*/
func (a *SlidesApiService) GetNotesSlideShapes(request GetNotesSlideShapesRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
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

/* Request for SlidesApiService.GetNotesSlideShapes
*/
type GetNotesSlideShapesRequest struct {
    name string
    slideIndex int32
    path string
    password string
    folder string
    storage string
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
func (a *SlidesApiService) GetNotesSlideWithFormat(request GetNotesSlideWithFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/{format}"
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
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
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

/* Request for SlidesApiService.GetNotesSlideWithFormat
*/
type GetNotesSlideWithFormatRequest struct {
    name string
    slideIndex int32
    format int32
    width *int32
    height *int32
    password string
    folder string
    storage string
    fontsFolder string
}

/* SlidesApiService Read paragraph portion info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) GetParagraphPortion(request GetParagraphPortionRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
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

/* Request for SlidesApiService.GetParagraphPortion
*/
type GetParagraphPortionRequest struct {
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

/* SlidesApiService Read paragraph portions info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portions*/
func (a *SlidesApiService) GetParagraphPortions(request GetParagraphPortionsRequest) (Portions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portions
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
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

/* Request for SlidesApiService.GetParagraphPortions
*/
type GetParagraphPortionsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide shape info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) GetSlideShape(request GetSlideShapeRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
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

/* Request for SlidesApiService.GetSlideShape
*/
type GetSlideShapeRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read shape paragraph info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) GetSlideShapeParagraph(request GetSlideShapeParagraphRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
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

/* Request for SlidesApiService.GetSlideShapeParagraph
*/
type GetSlideShapeParagraphRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    paragraphIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read shape paragraphs info.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraphs*/
func (a *SlidesApiService) GetSlideShapeParagraphs(request GetSlideShapeParagraphsRequest) (Paragraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraphs
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
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

/* Request for SlidesApiService.GetSlideShapeParagraphs
*/
type GetSlideShapeParagraphsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    password string
    folder string
    storage string
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
func (a *SlidesApiService) GetSlideShapes(request GetSlideShapesRequest) (Shapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Shapes
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
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

/* Request for SlidesApiService.GetSlideShapes
*/
type GetSlideShapesRequest struct {
    name string
    slideIndex int32
    path string
    password string
    folder string
    storage string
}

/* SlidesApiService Get API info.
 @return ApiInfo*/
func (a *SlidesApiService) GetSlidesApiInfo() (ApiInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ApiInfo
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


/* SlidesApiService Read presentation info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Documentstorage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) GetSlidesDocument(request GetSlidesDocumentRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
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

/* Request for SlidesApiService.GetSlidesDocument
*/
type GetSlidesDocumentRequest struct {
    name string
    password string
    storage string
    folder string
}

/* SlidesApiService Read presentation document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) GetSlidesDocumentProperties(request GetSlidesDocumentPropertiesRequest) (DocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperties
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
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

/* Request for SlidesApiService.GetSlidesDocumentProperties
*/
type GetSlidesDocumentPropertiesRequest struct {
    name string
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation document property.
 @param name Document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperty*/
func (a *SlidesApiService) GetSlidesDocumentProperty(request GetSlidesDocumentPropertyRequest) (DocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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

/* Request for SlidesApiService.GetSlidesDocumentProperty
*/
type GetSlidesDocumentPropertyRequest struct {
    name string
    propertyName string
    password string
    folder string
    storage string
}

/* SlidesApiService Get image binary data.
 @param name Document name.
 @param index Image index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return *os.File*/
func (a *SlidesApiService) GetSlidesImageWithDefaultFormat(request GetSlidesImageWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", request.index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
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

/* Request for SlidesApiService.GetSlidesImageWithDefaultFormat
*/
type GetSlidesImageWithDefaultFormatRequest struct {
    name string
    index int32
    password string
    folder string
    storage string
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
func (a *SlidesApiService) GetSlidesImageWithFormat(request GetSlidesImageWithFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", request.index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
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

/* Request for SlidesApiService.GetSlidesImageWithFormat
*/
type GetSlidesImageWithFormatRequest struct {
    name string
    index int32
    format int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation images info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Images*/
func (a *SlidesApiService) GetSlidesImages(request GetSlidesImagesRequest) (Images, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Images
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images"
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

/* Request for SlidesApiService.GetSlidesImages
*/
type GetSlidesImagesRequest struct {
    name string
    password string
    folder string
    storage string
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
func (a *SlidesApiService) GetSlidesPlaceholder(request GetSlidesPlaceholderRequest) (Placeholder, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Placeholder
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders/{placeholderIndex}"
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
	placeholderIndexPathStringValue := fmt.Sprintf("%v", request.placeholderIndex)
	if len(placeholderIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"placeholderIndex"+"}", placeholderIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"placeholderIndex"+"}", "", -1)
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

/* Request for SlidesApiService.GetSlidesPlaceholder
*/
type GetSlidesPlaceholderRequest struct {
    name string
    slideIndex int32
    placeholderIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide placeholders info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Placeholders*/
func (a *SlidesApiService) GetSlidesPlaceholders(request GetSlidesPlaceholdersRequest) (Placeholders, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Placeholders
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders"
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

/* Request for SlidesApiService.GetSlidesPlaceholders
*/
type GetSlidesPlaceholdersRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Extract presentation text items.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withEmpty" (bool) True to incude empty items.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return TextItems*/
func (a *SlidesApiService) GetSlidesPresentationTextItems(request GetSlidesPresentationTextItemsRequest) (TextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextItems
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/textItems"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.withEmpty != nil {
		if err := typeCheckParameter(*request.withEmpty, "bool", "withEmpty"); err != nil {
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

	localVarQueryParams.Add("withEmpty", parameterToString(request.withEmpty, ""))
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

/* Request for SlidesApiService.GetSlidesPresentationTextItems
*/
type GetSlidesPresentationTextItemsRequest struct {
    name string
    withEmpty *bool
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation slide info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) GetSlidesSlide(request GetSlidesSlideRequest) (Slide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
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

/* Request for SlidesApiService.GetSlidesSlide
*/
type GetSlidesSlideRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide background info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) GetSlidesSlideBackground(request GetSlidesSlideBackgroundRequest) (SlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideBackground
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
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

/* Request for SlidesApiService.GetSlidesSlideBackground
*/
type GetSlidesSlideBackgroundRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation slide comments.
 @param name Document name.
 @param slideIndex The position of the slide to be reordered.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideComments*/
func (a *SlidesApiService) GetSlidesSlideComments(request GetSlidesSlideCommentsRequest) (SlideComments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideComments
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/comments"
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

/* Request for SlidesApiService.GetSlidesSlideComments
*/
type GetSlidesSlideCommentsRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide images info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Images*/
func (a *SlidesApiService) GetSlidesSlideImages(request GetSlidesSlideImagesRequest) (Images, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Images
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/images"
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

/* Request for SlidesApiService.GetSlidesSlideImages
*/
type GetSlidesSlideImagesRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Extract slide text items.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "withEmpty" (bool) True to incude empty items.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return TextItems*/
func (a *SlidesApiService) GetSlidesSlideTextItems(request GetSlidesSlideTextItemsRequest) (TextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextItems
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/textItems"
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

	if request.withEmpty != nil {
		if err := typeCheckParameter(*request.withEmpty, "bool", "withEmpty"); err != nil {
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

	localVarQueryParams.Add("withEmpty", parameterToString(request.withEmpty, ""))
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

/* Request for SlidesApiService.GetSlidesSlideTextItems
*/
type GetSlidesSlideTextItemsRequest struct {
    name string
    slideIndex int32
    withEmpty *bool
    password string
    folder string
    storage string
}

/* SlidesApiService Read presentation slides info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slides*/
func (a *SlidesApiService) GetSlidesSlidesList(request GetSlidesSlidesListRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
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

/* Request for SlidesApiService.GetSlidesSlidesList
*/
type GetSlidesSlidesListRequest struct {
    name string
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide theme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Theme*/
func (a *SlidesApiService) GetSlidesTheme(request GetSlidesThemeRequest) (Theme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Theme
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme"
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

/* Request for SlidesApiService.GetSlidesTheme
*/
type GetSlidesThemeRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide theme color scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ColorScheme*/
func (a *SlidesApiService) GetSlidesThemeColorScheme(request GetSlidesThemeColorSchemeRequest) (ColorScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ColorScheme
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/colorScheme"
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

/* Request for SlidesApiService.GetSlidesThemeColorScheme
*/
type GetSlidesThemeColorSchemeRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide theme font scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return FontScheme*/
func (a *SlidesApiService) GetSlidesThemeFontScheme(request GetSlidesThemeFontSchemeRequest) (FontScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FontScheme
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/fontScheme"
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

/* Request for SlidesApiService.GetSlidesThemeFontScheme
*/
type GetSlidesThemeFontSchemeRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Read slide theme format scheme info.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return FormatScheme*/
func (a *SlidesApiService) GetSlidesThemeFormatScheme(request GetSlidesThemeFormatSchemeRequest) (FormatScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FormatScheme
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/formatScheme"
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

/* Request for SlidesApiService.GetSlidesThemeFormatScheme
*/
type GetSlidesThemeFormatSchemeRequest struct {
    name string
    slideIndex int32
    password string
    folder string
    storage string
}

/* SlidesApiService Move file
 @param srcPath Source file path e.g. &#39;/src.ext&#39;
 @param destPath Destination file path e.g. &#39;/dest.ext&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
     @param "versionId" (string) File version ID to move
 @return */
func (a *SlidesApiService) MoveFile(request MoveFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.versionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("destPath", parameterToString(request.destPath, ""))
	if localVarTempParam := request.srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.MoveFile
*/
type MoveFileRequest struct {
    srcPath string
    destPath string
    srcStorageName string
    destStorageName string
    versionId string
}

/* SlidesApiService Move folder
 @param srcPath Folder path to move e.g. &#39;/folder&#39;
 @param destPath Destination folder path to move to e.g &#39;/dst&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *SlidesApiService) MoveFolder(request MoveFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.srcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.srcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.destStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("destPath", parameterToString(request.destPath, ""))
	if localVarTempParam := request.srcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.destStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.MoveFolder
*/
type MoveFolderRequest struct {
    srcPath string
    destPath string
    srcStorageName string
    destStorageName string
}

/* SlidesApiService Check if file or folder exists
 @param path File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39;
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "versionId" (string) File version ID
 @return ObjectExist*/
func (a *SlidesApiService) ObjectExists(request ObjectExistsRequest) (ObjectExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ObjectExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/exist/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.versionId, "string", "versionId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.versionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.ObjectExists
*/
type ObjectExistsRequest struct {
    path string
    storageName string
    versionId string
}

/* SlidesApiService Creates new paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Paragraph DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return Paragraph*/
func (a *SlidesApiService) PostAddNewParagraph(request PostAddNewParagraphRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
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

/* Request for SlidesApiService.PostAddNewParagraph
*/
type PostAddNewParagraphRequest struct {
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

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Portion DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
 @return Portion*/
func (a *SlidesApiService) PostAddNewPortion(request PostAddNewPortionRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
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

/* Request for SlidesApiService.PostAddNewPortion
*/
type PostAddNewPortionRequest struct {
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

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeBase*/
func (a *SlidesApiService) PostAddNewShape(request PostAddNewShapeRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
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

/* Request for SlidesApiService.PostAddNewShape
*/
type PostAddNewShapeRequest struct {
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

/* SlidesApiService Add new notes slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (NotesSlide) A NotesSlide object with notes slide data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) PostAddNotesSlide(request PostAddNotesSlideRequest) (NotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlide
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

/* Request for SlidesApiService.PostAddNotesSlide
*/
type PostAddNotesSlideRequest struct {
    name string
    slideIndex int32
    dto INotesSlide
    password string
    folder string
    storage string
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
func (a *SlidesApiService) PostCopyLayoutSlideFromSourcePresentation(request PostCopyLayoutSlideFromSourcePresentationRequest) (LayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LayoutSlide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.cloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.cloneFromStorage, "string", "cloneFromStorage"); err != nil {
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

	localVarQueryParams.Add("cloneFrom", parameterToString(request.cloneFrom, ""))
	localVarQueryParams.Add("cloneFromPosition", parameterToString(request.cloneFromPosition, ""))
	if localVarTempParam := request.cloneFromPassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("cloneFromPassword", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.cloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("cloneFromStorage", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.PostCopyLayoutSlideFromSourcePresentation
*/
type PostCopyLayoutSlideFromSourcePresentationRequest struct {
    name string
    cloneFrom string
    cloneFromPosition int32
    cloneFromPassword string
    cloneFromStorage string
    password string
    folder string
    storage string
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
func (a *SlidesApiService) PostCopyMasterSlideFromSourcePresentation(request PostCopyMasterSlideFromSourcePresentationRequest) (MasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  MasterSlide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.cloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.cloneFromStorage, "string", "cloneFromStorage"); err != nil {
		return successPayload, nil, err
	}
	if request.applyToAll != nil {
		if err := typeCheckParameter(*request.applyToAll, "bool", "applyToAll"); err != nil {
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

	localVarQueryParams.Add("cloneFrom", parameterToString(request.cloneFrom, ""))
	localVarQueryParams.Add("cloneFromPosition", parameterToString(request.cloneFromPosition, ""))
	if localVarTempParam := request.cloneFromPassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("cloneFromPassword", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.cloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("cloneFromStorage", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("applyToAll", parameterToString(request.applyToAll, ""))
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

/* Request for SlidesApiService.PostCopyMasterSlideFromSourcePresentation
*/
type PostCopyMasterSlideFromSourcePresentationRequest struct {
    name string
    cloneFrom string
    cloneFromPosition int32
    cloneFromPassword string
    cloneFromStorage string
    applyToAll *bool
    password string
    folder string
    storage string
}

/* SlidesApiService Creates new paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Paragraph DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return Paragraph*/
func (a *SlidesApiService) PostNotesSlideAddNewParagraph(request PostNotesSlideAddNewParagraphRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
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

/* Request for SlidesApiService.PostNotesSlideAddNewParagraph
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

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Portion DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
 @return Portion*/
func (a *SlidesApiService) PostNotesSlideAddNewPortion(request PostNotesSlideAddNewPortionRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
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

/* Request for SlidesApiService.PostNotesSlideAddNewPortion
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

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeBase*/
func (a *SlidesApiService) PostNotesSlideAddNewShape(request PostNotesSlideAddNewShapeRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
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

/* Request for SlidesApiService.PostNotesSlideAddNewShape
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

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path 
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (int32) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostNotesSlideShapeSaveAs(request PostNotesSlideShapeSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/{format}"
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
	if request.bounds != nil {
		if err := typeCheckParameter(*request.bounds, "int32", "bounds"); err != nil {
			return successPayload, nil, err
		}
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
	if request.bounds != nil {
		localVarQueryParams.Add("bounds", parameterToString(*request.bounds, ""))
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

/* Request for SlidesApiService.PostNotesSlideShapeSaveAs
*/
type PostNotesSlideShapeSaveAsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    format int32
    options IIShapeExportOptions
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds *int32
    fontsFolder string
}

/* SlidesApiService Merge the presentation with other presentations specified in the request parameter.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "request" (PresentationsMergeRequest) PresentationsMergeRequest with a list of presentations to merge.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostPresentationMerge(request PostPresentationMergeRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
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
	localVarPostBody = &request.request
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PostPresentationMerge
*/
type PostPresentationMergeRequest struct {
    name string
    request IPresentationsMergeRequest
    password string
    storage string
    folder string
}

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path 
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (int32) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostShapeSaveAs(request PostShapeSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
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
	if request.bounds != nil {
		if err := typeCheckParameter(*request.bounds, "int32", "bounds"); err != nil {
			return successPayload, nil, err
		}
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
	if request.bounds != nil {
		localVarQueryParams.Add("bounds", parameterToString(*request.bounds, ""))
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

/* Request for SlidesApiService.PostShapeSaveAs
*/
type PostShapeSaveAsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    format int32
    options IIShapeExportOptions
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds *int32
    fontsFolder string
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
func (a *SlidesApiService) PostSlideSaveAs(request PostSlideSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
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
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
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

/* Request for SlidesApiService.PostSlideSaveAs
*/
type PostSlideSaveAsRequest struct {
    name string
    slideIndex int32
    format int32
    options IExportOptions
    width *int32
    height *int32
    password string
    folder string
    storage string
    fontsFolder string
}

/* SlidesApiService Create a slide.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "position" (int32) The target position at which to create the slide. Add to the end by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "layoutAlias" (string) Alias of layout slide for new slide. Alias may be the type of layout, name of layout slide or index
 @return Slides*/
func (a *SlidesApiService) PostSlidesAdd(request PostSlidesAddRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.position != nil {
		if err := typeCheckParameter(*request.position, "int32", "position"); err != nil {
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
	if err := typeCheckParameter(request.layoutAlias, "string", "layoutAlias"); err != nil {
		return successPayload, nil, err
	}

	if request.position != nil {
		localVarQueryParams.Add("position", parameterToString(*request.position, ""))
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
	if localVarTempParam := request.layoutAlias; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("layoutAlias", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.PostSlidesAdd
*/
type PostSlidesAddRequest struct {
    name string
    position *int32
    password string
    folder string
    storage string
    layoutAlias string
}

/* SlidesApiService Convert presentation from request content to format specified.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "document" ([]byte) Document data.
     @param "password" (string) Document password.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesConvert(request PostSlidesConvertRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
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
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("fontsFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/octet-stream", "multipart/form-data",  }

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

/* Request for SlidesApiService.PostSlidesConvert
*/
type PostSlidesConvertRequest struct {
    format int32
    document []byte
    password string
    fontsFolder string
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
func (a *SlidesApiService) PostSlidesCopy(request PostSlidesCopyRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/copy"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.position != nil {
		if err := typeCheckParameter(*request.position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.source, "string", "source"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.sourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.sourceStorage, "string", "sourceStorage"); err != nil {
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

	localVarQueryParams.Add("slideToCopy", parameterToString(request.slideToCopy, ""))
	if request.position != nil {
		localVarQueryParams.Add("position", parameterToString(*request.position, ""))
	}
	if localVarTempParam := request.source; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("source", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.sourcePassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sourcePassword", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.sourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sourceStorage", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.PostSlidesCopy
*/
type PostSlidesCopyRequest struct {
    name string
    slideToCopy int32
    position *int32
    source string
    sourcePassword string
    sourceStorage string
    password string
    folder string
    storage string
}

/* SlidesApiService Create a presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "data" ([]byte) Document input data.
     @param "inputPassword" (string) The password for input document.
     @param "password" (string) The document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocument(request PostSlidesDocumentRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
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

	if err := typeCheckParameter(request.inputPassword, "string", "inputPassword"); err != nil {
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

	if localVarTempParam := request.inputPassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("inputPassword", parameterToString(localVarTempParam, ""))
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
	localVarHttpContentTypes := []string{ "application/octet-stream", "multipart/form-data",  }

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

/* Request for SlidesApiService.PostSlidesDocument
*/
type PostSlidesDocumentRequest struct {
    name string
    data []byte
    inputPassword string
    password string
    storage string
    folder string
}

/* SlidesApiService Create presentation document from html.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "html" (string) HTML data.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocumentFromHtml(request PostSlidesDocumentFromHtmlRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
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

	if err := typeCheckParameter(request.html, "string", "html"); err != nil {
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

/* Request for SlidesApiService.PostSlidesDocumentFromHtml
*/
type PostSlidesDocumentFromHtmlRequest struct {
    name string
    html string
    password string
    storage string
    folder string
}

/* SlidesApiService Create a presentation from an existing source.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "sourcePath" (string) Template file path.
     @param "sourcePassword" (string) Template file password.
     @param "sourceStorage" (string) Template storage name.
     @param "password" (string) The document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocumentFromSource(request PostSlidesDocumentFromSourceRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromSource"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.sourcePath, "string", "sourcePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.sourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.sourceStorage, "string", "sourceStorage"); err != nil {
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

	if localVarTempParam := request.sourcePath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sourcePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.sourcePassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sourcePassword", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.sourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("sourceStorage", parameterToString(localVarTempParam, ""))
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PostSlidesDocumentFromSource
*/
type PostSlidesDocumentFromSourceRequest struct {
    name string
    sourcePath string
    sourcePassword string
    sourceStorage string
    password string
    storage string
    folder string
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
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocumentFromTemplate(request PostSlidesDocumentFromTemplateRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromTemplate"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.data, "string", "data"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.templatePassword, "string", "templatePassword"); err != nil {
		return successPayload, nil, err
	}
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
	if localVarTempParam := request.templatePassword; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("templatePassword", parameterToString(localVarTempParam, ""))
	}
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

/* Request for SlidesApiService.PostSlidesDocumentFromTemplate
*/
type PostSlidesDocumentFromTemplateRequest struct {
    name string
    templatePath string
    data string
    templatePassword string
    templateStorage string
    isImageDataEmbedded *bool
    password string
    storage string
    folder string
}

/* SlidesApiService Performs slides pipeline.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pipeline" (Pipeline) A Pipeline object.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesPipeline(request PostSlidesPipelineRequest) (*os.File, *http.Response, error) {
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
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

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

/* Request for SlidesApiService.PostSlidesPipeline
*/
type PostSlidesPipelineRequest struct {
    pipeline IPipeline
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
func (a *SlidesApiService) PostSlidesPresentationReplaceText(request PostSlidesPresentationReplaceTextRequest) (DocumentReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentReplaceResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/replaceText"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ignoreCase != nil {
		if err := typeCheckParameter(*request.ignoreCase, "bool", "ignoreCase"); err != nil {
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

	localVarQueryParams.Add("oldValue", parameterToString(request.oldValue, ""))
	localVarQueryParams.Add("newValue", parameterToString(request.newValue, ""))
	localVarQueryParams.Add("ignoreCase", parameterToString(request.ignoreCase, ""))
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

/* Request for SlidesApiService.PostSlidesPresentationReplaceText
*/
type PostSlidesPresentationReplaceTextRequest struct {
    name string
    oldValue string
    newValue string
    ignoreCase *bool
    password string
    folder string
    storage string
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
func (a *SlidesApiService) PostSlidesReorder(request PostSlidesReorderRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/move"
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

	localVarQueryParams.Add("newPosition", parameterToString(request.newPosition, ""))
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

/* Request for SlidesApiService.PostSlidesReorder
*/
type PostSlidesReorderRequest struct {
    name string
    slideIndex int32
    newPosition int32
    password string
    folder string
    storage string
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
func (a *SlidesApiService) PostSlidesReorderMany(request PostSlidesReorderManyRequest) (Slides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slides
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/reorder"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.oldPositions, "[]int32", "oldPositions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.newPositions, "[]int32", "newPositions"); err != nil {
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

	if localVarTempParam := request.oldPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("oldPositions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.newPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("newPositions", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.PostSlidesReorderMany
*/
type PostSlidesReorderManyRequest struct {
    name string
    oldPositions []int32
    newPositions []int32
    password string
    folder string
    storage string
}

/* SlidesApiService Save a presentation to a specified format.
 @param name Document name.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesSaveAs(request PostSlidesSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
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

/* Request for SlidesApiService.PostSlidesSaveAs
*/
type PostSlidesSaveAsRequest struct {
    name string
    format int32
    options IExportOptions
    password string
    storage string
    folder string
    fontsFolder string
}

/* SlidesApiService Set document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "properties" (DocumentProperties) New properties.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) PostSlidesSetDocumentProperties(request PostSlidesSetDocumentPropertiesRequest) (DocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperties
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
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
	localVarPostBody = &request.properties
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PostSlidesSetDocumentProperties
*/
type PostSlidesSetDocumentPropertiesRequest struct {
    name string
    properties IDocumentProperties
    password string
    folder string
    storage string
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
func (a *SlidesApiService) PostSlidesSlideReplaceText(request PostSlidesSlideReplaceTextRequest) (SlideReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideReplaceResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/replaceText"
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

	if request.ignoreCase != nil {
		if err := typeCheckParameter(*request.ignoreCase, "bool", "ignoreCase"); err != nil {
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

	localVarQueryParams.Add("oldValue", parameterToString(request.oldValue, ""))
	localVarQueryParams.Add("newValue", parameterToString(request.newValue, ""))
	localVarQueryParams.Add("ignoreCase", parameterToString(request.ignoreCase, ""))
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

/* Request for SlidesApiService.PostSlidesSlideReplaceText
*/
type PostSlidesSlideReplaceTextRequest struct {
    name string
    slideIndex int32
    oldValue string
    newValue string
    ignoreCase *bool
    password string
    folder string
    storage string
}

/* SlidesApiService Splitting presentations. Create one image per slide.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "format" (int32) Export format. Default value is jpeg.
     @param "width" (int32) The width of created images.
     @param "height" (int32) The height of created images.
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "destFolder" (string) Folder on storage where images are going to be uploaded. If not specified then images are uploaded to same folder as presentation.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "fontsFolder" (string) Custom fonts folder.
 @return SplitDocumentResult*/
func (a *SlidesApiService) PostSlidesSplit(request PostSlidesSplitRequest) (SplitDocumentResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SplitDocumentResult
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

	if request.format != nil {
		if err := typeCheckParameter(*request.format, "int32", "format"); err != nil {
			return successPayload, nil, err
		}
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
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if request.format != nil {
		localVarQueryParams.Add("format", parameterToString(*request.format, ""))
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

/* Request for SlidesApiService.PostSlidesSplit
*/
type PostSlidesSplitRequest struct {
    name string
    options IExportOptions
    format *int32
    width *int32
    height *int32
    to *int32
    from *int32
    destFolder string
    password string
    storage string
    folder string
    fontsFolder string
}

/* SlidesApiService Update a layoutSlide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "slideDto" (LayoutSlide) Slide update data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) PutLayoutSlide(request PutLayoutSlideRequest) (LayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LayoutSlide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
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
	localVarPostBody = &request.slideDto
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PutLayoutSlide
*/
type PutLayoutSlideRequest struct {
    name string
    slideIndex int32
    slideDto ILayoutSlide
    password string
    folder string
    storage string
}

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path 
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (int32) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) PutNotesSlideShapeSaveAs(request PutNotesSlideShapeSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{path}/{shapeIndex}/{format}"
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
		return nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return nil, err
	}
	if request.scaleX != nil {
		if err := typeCheckParameter(*request.scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if request.scaleY != nil {
		if err := typeCheckParameter(*request.scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if request.bounds != nil {
		if err := typeCheckParameter(*request.bounds, "int32", "bounds"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(request.outPath, ""))
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
	if request.bounds != nil {
		localVarQueryParams.Add("bounds", parameterToString(*request.bounds, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.PutNotesSlideShapeSaveAs
*/
type PutNotesSlideShapeSaveAsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    format int32
    outPath string
    options IIShapeExportOptions
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds *int32
    fontsFolder string
}

/* SlidesApiService Merge the presentation with other presentations or some of their slides specified in the request parameter.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "request" (OrderedMergeRequest) OrderedMergeRequest with a list of presentations and slide indices to merge.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PutPresentationMerge(request PutPresentationMergeRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
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
	localVarPostBody = &request.request
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PutPresentationMerge
*/
type PutPresentationMergeRequest struct {
    name string
    request IOrderedMergeRequest
    password string
    storage string
    folder string
}

/* SlidesApiService Update portion properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Portion DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) PutSetParagraphPortionProperties(request PutSetParagraphPortionPropertiesRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
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

/* Request for SlidesApiService.PutSetParagraphPortionProperties
*/
type PutSetParagraphPortionPropertiesRequest struct {
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

/* SlidesApiService Update paragraph properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Paragraph DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) PutSetParagraphProperties(request PutSetParagraphPropertiesRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
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

/* Request for SlidesApiService.PutSetParagraphProperties
*/
type PutSetParagraphPropertiesRequest struct {
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

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
 @param path 
 @param shapeIndex Index of shape starting from 1
 @param format Export picture format.
 @param outPath Output path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (IShapeExportOptions) export options
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (int32) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) PutShapeSaveAs(request PutShapeSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
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
		return nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return nil, err
	}
	if request.scaleX != nil {
		if err := typeCheckParameter(*request.scaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if request.scaleY != nil {
		if err := typeCheckParameter(*request.scaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if request.bounds != nil {
		if err := typeCheckParameter(*request.bounds, "int32", "bounds"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(request.outPath, ""))
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
	if request.bounds != nil {
		localVarQueryParams.Add("bounds", parameterToString(*request.bounds, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.PutShapeSaveAs
*/
type PutShapeSaveAsRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    format int32
    outPath string
    options IIShapeExportOptions
    password string
    folder string
    storage string
    scaleX *float64
    scaleY *float64
    bounds *int32
    fontsFolder string
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
func (a *SlidesApiService) PutSlideSaveAs(request PutSlideSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
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
			return nil, err
		}
	}
	if request.height != nil {
		if err := typeCheckParameter(*request.height, "int32", "height"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(request.outPath, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlideSaveAs
*/
type PutSlideSaveAsRequest struct {
    name string
    slideIndex int32
    format int32
    outPath string
    options IExportOptions
    width *int32
    height *int32
    password string
    folder string
    storage string
    fontsFolder string
}

/* SlidesApiService Update shape properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) PutSlideShapeInfo(request PutSlideShapeInfoRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
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

/* Request for SlidesApiService.PutSlideShapeInfo
*/
type PutSlideShapeInfoRequest struct {
    name string
    slideIndex int32
    path string
    shapeIndex int32
    dto IShapeBase
    password string
    folder string
    storage string
}

/* SlidesApiService Convert presentation from request content to format specified.
 @param format Export format.
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "document" ([]byte) Document data.
     @param "password" (string) Document password.
     @param "fontsFolder" (string) Custom fonts folder.
 @return */
func (a *SlidesApiService) PutSlidesConvert(request PutSlidesConvertRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
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
		return nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(request.outPath, ""))
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.fontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("fontsFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/octet-stream", "multipart/form-data",  }

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
	localVarPostBody = request.document
	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlidesConvert
*/
type PutSlidesConvertRequest struct {
    format int32
    outPath string
    document []byte
    password string
    fontsFolder string
}

/* SlidesApiService Update presentation document from html.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "html" (string) HTML data.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PutSlidesDocumentFromHtml(request PutSlidesDocumentFromHtmlRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
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

	if err := typeCheckParameter(request.html, "string", "html"); err != nil {
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

/* Request for SlidesApiService.PutSlidesDocumentFromHtml
*/
type PutSlidesDocumentFromHtmlRequest struct {
    name string
    html string
    password string
    storage string
    folder string
}

/* SlidesApiService Save a presentation to a specified format.
 @param name Document name.
 @param outPath Output path.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (ExportOptions) Export options.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "fontsFolder" (string) Custom fonts folder.
 @return */
func (a *SlidesApiService) PutSlidesSaveAs(request PutSlidesSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
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
		return nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.fontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(request.outPath, ""))
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
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
		return nil, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("-->: %v\n", r)
	}
	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--: %v\n", localVarHttpResponse)
	}
	defer localVarHttpResponse.Body.Close()
	responseBytes, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if a.client.cfg.Debug {
		fmt.Printf("<--BODY: %v\n", string(responseBytes))
	}
	responseBytes = bytes.Replace(responseBytes, []byte(":\"NaN\""), []byte(":null"), -1) 
	responseBody := bytes.NewReader(responseBytes)
	if localVarHttpResponse.StatusCode >= 300 {
		var errorMessage ErrorMessage
		if err = json.NewDecoder(responseBody).Decode(&errorMessage); err != nil {
			return localVarHttpResponse, reportError(localVarHttpResponse.Status)
		}
		return localVarHttpResponse, reportError(errorMessage.getMessage())
	}


	return localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlidesSaveAs
*/
type PutSlidesSaveAsRequest struct {
    name string
    outPath string
    format int32
    options IExportOptions
    password string
    storage string
    folder string
    fontsFolder string
}

/* SlidesApiService Set document property.
 @param name Document name.
 @param propertyName The property name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "property" (DocumentProperty) Property with the value.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperty*/
func (a *SlidesApiService) PutSlidesSetDocumentProperty(request PutSlidesSetDocumentPropertyRequest) (DocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.propertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
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
	localVarPostBody = &request.property
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlidesSetDocumentProperty
*/
type PutSlidesSetDocumentPropertyRequest struct {
    name string
    propertyName string
    property IDocumentProperty
    password string
    folder string
    storage string
}

/* SlidesApiService Update a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "slideDto" (Slide) Slide update data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Slide*/
func (a *SlidesApiService) PutSlidesSlide(request PutSlidesSlideRequest) (Slide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Slide
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
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
	localVarPostBody = &request.slideDto
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlidesSlide
*/
type PutSlidesSlideRequest struct {
    name string
    slideIndex int32
    slideDto ISlide
    password string
    folder string
    storage string
}

/* SlidesApiService Set background for a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "background" (SlideBackground) Slide background update data.
     @param "folder" (string) Document folder.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) PutSlidesSlideBackground(request PutSlidesSlideBackgroundRequest) (SlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideBackground
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
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

	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = &request.background
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.PutSlidesSlideBackground
*/
type PutSlidesSlideBackgroundRequest struct {
    name string
    slideIndex int32
    background ISlideBackground
    folder string
    password string
    storage string
}

/* SlidesApiService Set background color for a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param color Slide background target color in RRGGBB format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) Document folder.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) PutSlidesSlideBackgroundColor(request PutSlidesSlideBackgroundColorRequest) (SlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SlideBackground
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/backgroundColor"
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

	if err := typeCheckParameter(request.folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("color", parameterToString(request.color, ""))
	if localVarTempParam := request.folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.password; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("password", parameterToString(localVarTempParam, ""))
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

/* Request for SlidesApiService.PutSlidesSlideBackgroundColor
*/
type PutSlidesSlideBackgroundColorRequest struct {
    name string
    slideIndex int32
    color string
    folder string
    password string
    storage string
}

/* SlidesApiService Set slide size for a presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "width" (int32) Slide width.
     @param "height" (int32) Slide height.
     @param "sizeType" (int32) Standard slide size type.
     @param "scaleType" (int32) Standard slide scale type.
 @return Document*/
func (a *SlidesApiService) PutSlidesSlideSize(request PutSlidesSlideSizeRequest) (Document, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Document
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideSize"
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
	if request.sizeType != nil {
		if err := typeCheckParameter(*request.sizeType, "int32", "sizeType"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.scaleType != nil {
		if err := typeCheckParameter(*request.scaleType, "int32", "scaleType"); err != nil {
			return successPayload, nil, err
		}
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
	if request.sizeType != nil {
		localVarQueryParams.Add("sizeType", parameterToString(*request.sizeType, ""))
	}
	if request.scaleType != nil {
		localVarQueryParams.Add("scaleType", parameterToString(*request.scaleType, ""))
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

/* Request for SlidesApiService.PutSlidesSlideSize
*/
type PutSlidesSlideSizeRequest struct {
    name string
    password string
    storage string
    folder string
    width *int32
    height *int32
    sizeType *int32
    scaleType *int32
}

/* SlidesApiService Update notes slide properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (NotesSlide) A NotesSlide object with notes slide data.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return NotesSlide*/
func (a *SlidesApiService) PutUpdateNotesSlide(request PutUpdateNotesSlideRequest) (NotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  NotesSlide
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

/* Request for SlidesApiService.PutUpdateNotesSlide
*/
type PutUpdateNotesSlideRequest struct {
    name string
    slideIndex int32
    dto INotesSlide
    password string
    folder string
    storage string
}

/* SlidesApiService Update shape properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ShapeBase*/
func (a *SlidesApiService) PutUpdateNotesSlideShape(request PutUpdateNotesSlideShapeRequest) (ShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ShapeBase
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

/* Request for SlidesApiService.PutUpdateNotesSlideShape
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

/* SlidesApiService Update paragraph properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Paragraph) Paragraph DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Paragraph*/
func (a *SlidesApiService) PutUpdateNotesSlideShapeParagraph(request PutUpdateNotesSlideShapeParagraphRequest) (Paragraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Paragraph
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

/* Request for SlidesApiService.PutUpdateNotesSlideShapeParagraph
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

/* SlidesApiService Update portion properties.
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path (for smart art and group shapes).
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param portionIndex Portion index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (Portion) Portion DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return Portion*/
func (a *SlidesApiService) PutUpdateNotesSlideShapePortion(request PutUpdateNotesSlideShapePortionRequest) (Portion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Portion
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

/* Request for SlidesApiService.PutUpdateNotesSlideShapePortion
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

/* SlidesApiService Check if storage exists
 @param storageName Storage name
 @return StorageExist*/
func (a *SlidesApiService) StorageExists(request StorageExistsRequest) (StorageExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  StorageExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/{storageName}/exist"
	storageNamePathStringValue := fmt.Sprintf("%v", request.storageName)
	if len(storageNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", storageNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"storageName"+"}", "", -1)
	}

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

/* Request for SlidesApiService.StorageExists
*/
type StorageExistsRequest struct {
    storageName string
}

/* SlidesApiService Upload file
 @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.             
 @param file File to upload
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FilesUploadResult*/
func (a *SlidesApiService) UploadFile(request UploadFileRequest) (FilesUploadResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FilesUploadResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.storageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.storageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("storageName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

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
	localVarPostBody = &request.file
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

	if err = json.NewDecoder(responseBody).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* Request for SlidesApiService.UploadFile
*/
type UploadFileRequest struct {
    path string
    file []byte
    storageName string
}

