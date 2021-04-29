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

/* Request for SlidesApiService.CopyFile
*/
type CopyFileRequest struct {
    SrcPath string
    DestPath string
    SrcStorageName string
    DestStorageName string
    VersionId string
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

/* Request for SlidesApiService.CopyFolder
*/
type CopyFolderRequest struct {
    SrcPath string
    DestPath string
    SrcStorageName string
    DestStorageName string
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

/* Request for SlidesApiService.CreateFolder
*/
type CreateFolderRequest struct {
    Path string
    StorageName string
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

/* Request for SlidesApiService.DeleteChartCategory
*/
type DeleteChartCategoryRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    CategoryIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteChartDataPoint
*/
type DeleteChartDataPointRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    SeriesIndex int32
    PointIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteChartSeries
*/
type DeleteChartSeriesRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    SeriesIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteFile
*/
type DeleteFileRequest struct {
    Path string
    StorageName string
    VersionId string
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

/* Request for SlidesApiService.DeleteFolder
*/
type DeleteFolderRequest struct {
    Path string
    StorageName string
    Recursive *bool
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

/* Request for SlidesApiService.DeleteNotesSlide
*/
type DeleteNotesSlideRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlideParagraph
*/
type DeleteNotesSlideParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlideParagraphs
*/
type DeleteNotesSlideParagraphsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Paragraphs []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlidePortion
*/
type DeleteNotesSlidePortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlidePortions
*/
type DeleteNotesSlidePortionsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Portions []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlideShape
*/
type DeleteNotesSlideShapeRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteNotesSlideShapes
*/
type DeleteNotesSlideShapesRequest struct {
    Name string
    SlideIndex int32
    Shapes []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteParagraph
*/
type DeleteParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteParagraphs
*/
type DeleteParagraphsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Paragraphs []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeletePortion
*/
type DeletePortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeletePortions
*/
type DeletePortionsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Portions []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSection
*/
type DeleteSectionRequest struct {
    Name string
    SectionIndex int32
    WithSlides *bool
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSections
*/
type DeleteSectionsRequest struct {
    Name string
    Sections []int32
    WithSlides *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimation(request DeleteSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimation(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimation
*/
type DeleteSlideAnimationRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimationEffect(request DeleteSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimationEffect(request.Name, request.SlideIndex, request.EffectIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimationEffect
*/
type DeleteSlideAnimationEffectRequest struct {
    Name string
    SlideIndex int32
    EffectIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequence(request DeleteSlideAnimationInteractiveSequenceRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimationInteractiveSequence(request.Name, request.SlideIndex, request.SequenceIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimationInteractiveSequence
*/
type DeleteSlideAnimationInteractiveSequenceRequest struct {
    Name string
    SlideIndex int32
    SequenceIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequenceEffect(request DeleteSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimationInteractiveSequenceEffect(request.Name, request.SlideIndex, request.SequenceIndex, request.EffectIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimationInteractiveSequenceEffect
*/
type DeleteSlideAnimationInteractiveSequenceEffectRequest struct {
    Name string
    SlideIndex int32
    SequenceIndex int32
    EffectIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequences(request DeleteSlideAnimationInteractiveSequencesRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimationInteractiveSequences(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimationInteractiveSequences
*/
type DeleteSlideAnimationInteractiveSequencesRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideAnimationMainSequence(request DeleteSlideAnimationMainSequenceRequest) (ISlideAnimation, *http.Response, error) {
	return a.DeleteAnimationMainSequence(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideAnimationMainSequence
*/
type DeleteSlideAnimationMainSequenceRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideByIndex(request DeleteSlideByIndexRequest) (ISlides, *http.Response, error) {
	return a.DeleteSlide(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideByIndex
*/
type DeleteSlideByIndexRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideShape(request DeleteSlideShapeRequest) (IShapes, *http.Response, error) {
	return a.DeleteShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideShape
*/
type DeleteSlideShapeRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideShapes(request DeleteSlideShapesRequest) (IShapes, *http.Response, error) {
	return a.DeleteShapes(request.Name, request.SlideIndex, request.Shapes, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideShapes
*/
type DeleteSlideShapesRequest struct {
    Name string
    SlideIndex int32
    Shapes []int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideSubshape(request DeleteSlideSubshapeRequest) (IShapes, *http.Response, error) {
	return a.DeleteSubshape(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideSubshape
*/
type DeleteSlideSubshapeRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlideSubshapes(request DeleteSlideSubshapesRequest) (IShapes, *http.Response, error) {
	return a.DeleteSubshapes(request.Name, request.SlideIndex, request.Path, request.Shapes, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlideSubshapes
*/
type DeleteSlideSubshapesRequest struct {
    Name string
    SlideIndex int32
    Path string
    Shapes []int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlidesCleanSlidesList(request DeleteSlidesCleanSlidesListRequest) (ISlides, *http.Response, error) {
	return a.DeleteSlides(request.Name, request.Slides, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlidesCleanSlidesList
*/
type DeleteSlidesCleanSlidesListRequest struct {
    Name string
    Slides []int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlidesDocumentProperties(request DeleteSlidesDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	return a.DeleteDocumentProperties(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlidesDocumentProperties
*/
type DeleteSlidesDocumentPropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlidesDocumentProperty(request DeleteSlidesDocumentPropertyRequest) (IDocumentProperties, *http.Response, error) {
	return a.DeleteDocumentProperty(request.Name, request.PropertyName, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlidesDocumentProperty
*/
type DeleteSlidesDocumentPropertyRequest struct {
    Name string
    PropertyName string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlidesProtectionProperties(request DeleteSlidesProtectionPropertiesRequest) (IProtectionProperties, *http.Response, error) {
	return a.DeleteProtection(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlidesProtectionProperties
*/
type DeleteSlidesProtectionPropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) DeleteSlidesProtectionPropertiesOnline(request DeleteSlidesProtectionPropertiesOnlineRequest) (*os.File, *http.Response, error) {
	return a.DeleteProtectionOnline(request.Document, request.Password)
}

/* Request for SlidesApiService.DeleteSlidesProtectionPropertiesOnline
*/
type DeleteSlidesProtectionPropertiesOnlineRequest struct {
    Document []byte
    Password string
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

func (a *SlidesApiService) DeleteSlidesSlideBackground(request DeleteSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	return a.DeleteBackground(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.DeleteSlidesSlideBackground
*/
type DeleteSlidesSlideBackgroundRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSubshapeParagraph
*/
type DeleteSubshapeParagraphRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSubshapeParagraphs
*/
type DeleteSubshapeParagraphsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Paragraphs []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSubshapePortion
*/
type DeleteSubshapePortionRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteSubshapePortions
*/
type DeleteSubshapePortionsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Portions []int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DeleteWatermark
*/
type DeleteWatermarkRequest struct {
    Name string
    ShapeName string
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.DownloadFile
*/
type DownloadFileRequest struct {
    Path string
    StorageName string
    VersionId string
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

/* Request for SlidesApiService.GetDiscUsage
*/
type GetDiscUsageRequest struct {
    StorageName string
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

/* Request for SlidesApiService.GetFileVersions
*/
type GetFileVersionsRequest struct {
    Path string
    StorageName string
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

/* Request for SlidesApiService.GetFilesList
*/
type GetFilesListRequest struct {
    Path string
    StorageName string
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

/* Request for SlidesApiService.GetLayoutSlide
*/
type GetLayoutSlideRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetLayoutSlidesList(request GetLayoutSlidesListRequest) (ILayoutSlides, *http.Response, error) {
	return a.GetLayoutSlides(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetLayoutSlidesList
*/
type GetLayoutSlidesListRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetMasterSlide
*/
type GetMasterSlideRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetMasterSlidesList(request GetMasterSlidesListRequest) (IMasterSlides, *http.Response, error) {
	return a.GetMasterSlides(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetMasterSlidesList
*/
type GetMasterSlidesListRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetNotesSlide
*/
type GetNotesSlideRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideExists(request GetNotesSlideExistsRequest) (IEntityExists, *http.Response, error) {
	return a.NotesSlideExists(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetNotesSlideExists
*/
type GetNotesSlideExistsRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetNotesSlideHeaderFooter
*/
type GetNotesSlideHeaderFooterRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetNotesSlideShape
*/
type GetNotesSlideShapeRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideShapeParagraph(request GetNotesSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	return a.GetNotesSlideParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetNotesSlideShapeParagraph
*/
type GetNotesSlideShapeParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideShapeParagraphs(request GetNotesSlideShapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	return a.GetNotesSlideParagraphs(request.Name, request.SlideIndex, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetNotesSlideShapeParagraphs
*/
type GetNotesSlideShapeParagraphsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideShapePortion(request GetNotesSlideShapePortionRequest) (IPortion, *http.Response, error) {
	return a.GetNotesSlidePortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetNotesSlideShapePortion
*/
type GetNotesSlideShapePortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideShapePortions(request GetNotesSlideShapePortionsRequest) (IPortions, *http.Response, error) {
	return a.GetNotesSlidePortions(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetNotesSlideShapePortions
*/
type GetNotesSlideShapePortionsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetNotesSlideShapes
*/
type GetNotesSlideShapesRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetNotesSlideWithFormat(request GetNotesSlideWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadNotesSlide(request.Name, request.SlideIndex, request.Format, request.Width, request.Height, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.GetNotesSlideWithFormat
*/
type GetNotesSlideWithFormatRequest struct {
    Name string
    SlideIndex int32
    Format string
    Width *int32
    Height *int32
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) GetParagraphPortion(request GetParagraphPortionRequest) (IPortion, *http.Response, error) {
	return a.GetPortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetParagraphPortion
*/
type GetParagraphPortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetParagraphPortions(request GetParagraphPortionsRequest) (IPortions, *http.Response, error) {
	return a.GetPortions(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetParagraphPortions
*/
type GetParagraphPortionsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetSections
*/
type GetSectionsRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideAnimation(request GetSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	return a.GetAnimation(request.Name, request.SlideIndex, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideAnimation
*/
type GetSlideAnimationRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex *int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.GetSlideHeaderFooter
*/
type GetSlideHeaderFooterRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideShape(request GetSlideShapeRequest) (IShapeBase, *http.Response, error) {
	return a.GetShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideShape
*/
type GetSlideShapeRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideShapeParagraph(request GetSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	return a.GetParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideShapeParagraph
*/
type GetSlideShapeParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideShapeParagraphs(request GetSlideShapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	return a.GetParagraphs(request.Name, request.SlideIndex, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideShapeParagraphs
*/
type GetSlideShapeParagraphsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideShapes(request GetSlideShapesRequest) (IShapes, *http.Response, error) {
	return a.GetShapes(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideShapes
*/
type GetSlideShapesRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideSubshape(request GetSlideSubshapeRequest) (IShapeBase, *http.Response, error) {
	return a.GetSubshape(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideSubshape
*/
type GetSlideSubshapeRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideSubshapeParagraph(request GetSlideSubshapeParagraphRequest) (IParagraph, *http.Response, error) {
	return a.GetSubshapeParagraph(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideSubshapeParagraph
*/
type GetSlideSubshapeParagraphRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideSubshapeParagraphs(request GetSlideSubshapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	return a.GetSubshapeParagraphs(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideSubshapeParagraphs
*/
type GetSlideSubshapeParagraphsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlideSubshapes(request GetSlideSubshapesRequest) (IShapes, *http.Response, error) {
	return a.GetSubshapes(request.Name, request.SlideIndex, request.Path, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlideSubshapes
*/
type GetSlideSubshapesRequest struct {
    Name string
    SlideIndex int32
    Path string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesApiInfo() (IApiInfo, *http.Response, error) {
	return a.GetApiInfo()
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

func (a *SlidesApiService) GetSlidesDocument(request GetSlidesDocumentRequest) (IDocument, *http.Response, error) {
	return a.GetPresentation(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesDocument
*/
type GetSlidesDocumentRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesDocumentProperties(request GetSlidesDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	return a.GetDocumentProperties(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesDocumentProperties
*/
type GetSlidesDocumentPropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesDocumentProperty(request GetSlidesDocumentPropertyRequest) (IDocumentProperty, *http.Response, error) {
	return a.GetDocumentProperty(request.Name, request.PropertyName, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesDocumentProperty
*/
type GetSlidesDocumentPropertyRequest struct {
    Name string
    PropertyName string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesImageWithDefaultFormat(request GetSlidesImageWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImageDefaultFormat(request.Name, request.Index, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesImageWithDefaultFormat
*/
type GetSlidesImageWithDefaultFormatRequest struct {
    Name string
    Index int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesImageWithFormat(request GetSlidesImageWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImage(request.Name, request.Index, request.Format, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesImageWithFormat
*/
type GetSlidesImageWithFormatRequest struct {
    Name string
    Index int32
    Format string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesImages(request GetSlidesImagesRequest) (IImages, *http.Response, error) {
	return a.GetPresentationImages(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesImages
*/
type GetSlidesImagesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesPlaceholder(request GetSlidesPlaceholderRequest) (IPlaceholder, *http.Response, error) {
	return a.GetPlaceholder(request.Name, request.SlideIndex, request.PlaceholderIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesPlaceholder
*/
type GetSlidesPlaceholderRequest struct {
    Name string
    SlideIndex int32
    PlaceholderIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesPlaceholders(request GetSlidesPlaceholdersRequest) (IPlaceholders, *http.Response, error) {
	return a.GetPlaceholders(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesPlaceholders
*/
type GetSlidesPlaceholdersRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesPresentationTextItems(request GetSlidesPresentationTextItemsRequest) (ITextItems, *http.Response, error) {
	return a.GetPresentationTextItems(request.Name, request.WithEmpty, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesPresentationTextItems
*/
type GetSlidesPresentationTextItemsRequest struct {
    Name string
    WithEmpty *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesProtectionProperties(request GetSlidesProtectionPropertiesRequest) (IProtectionProperties, *http.Response, error) {
	return a.GetProtectionProperties(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesProtectionProperties
*/
type GetSlidesProtectionPropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlide(request GetSlidesSlideRequest) (ISlide, *http.Response, error) {
	return a.GetSlide(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlide
*/
type GetSlidesSlideRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlideBackground(request GetSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	return a.GetBackground(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlideBackground
*/
type GetSlidesSlideBackgroundRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
}

/* SlidesApiService Read presentation slide comments.
 @param name Document name.
 @param slideIndex The position of the slide to be reordered.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideComments*/
func (a *SlidesApiService) GetComments(name string, slideIndex int32, password string, folder string, storage string) (ISlideComments, *http.Response, error) {
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

func (a *SlidesApiService) GetSlidesSlideComments(request GetSlidesSlideCommentsRequest) (ISlideComments, *http.Response, error) {
	return a.GetComments(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlideComments
*/
type GetSlidesSlideCommentsRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlideImages(request GetSlidesSlideImagesRequest) (IImages, *http.Response, error) {
	return a.GetSlideImages(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlideImages
*/
type GetSlidesSlideImagesRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlideProperties(request GetSlidesSlidePropertiesRequest) (ISlideProperties, *http.Response, error) {
	return a.GetSlideProperties(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlideProperties
*/
type GetSlidesSlidePropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlideTextItems(request GetSlidesSlideTextItemsRequest) (ITextItems, *http.Response, error) {
	return a.GetSlideTextItems(request.Name, request.SlideIndex, request.WithEmpty, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlideTextItems
*/
type GetSlidesSlideTextItemsRequest struct {
    Name string
    SlideIndex int32
    WithEmpty *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesSlidesList(request GetSlidesSlidesListRequest) (ISlides, *http.Response, error) {
	return a.GetSlides(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesSlidesList
*/
type GetSlidesSlidesListRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesTheme(request GetSlidesThemeRequest) (ITheme, *http.Response, error) {
	return a.GetTheme(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesTheme
*/
type GetSlidesThemeRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesThemeColorScheme(request GetSlidesThemeColorSchemeRequest) (IColorScheme, *http.Response, error) {
	return a.GetColorScheme(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesThemeColorScheme
*/
type GetSlidesThemeColorSchemeRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesThemeFontScheme(request GetSlidesThemeFontSchemeRequest) (IFontScheme, *http.Response, error) {
	return a.GetFontScheme(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesThemeFontScheme
*/
type GetSlidesThemeFontSchemeRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesThemeFormatScheme(request GetSlidesThemeFormatSchemeRequest) (IFormatScheme, *http.Response, error) {
	return a.GetFormatScheme(request.Name, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesThemeFormatScheme
*/
type GetSlidesThemeFormatSchemeRequest struct {
    Name string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSlidesViewProperties(request GetSlidesViewPropertiesRequest) (IViewProperties, *http.Response, error) {
	return a.GetViewProperties(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSlidesViewProperties
*/
type GetSlidesViewPropertiesRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSubshapeParagraphPortion(request GetSubshapeParagraphPortionRequest) (IPortion, *http.Response, error) {
	return a.GetSubshapePortion(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSubshapeParagraphPortion
*/
type GetSubshapeParagraphPortionRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) GetSubshapeParagraphPortions(request GetSubshapeParagraphPortionsRequest) (IPortions, *http.Response, error) {
	return a.GetSubshapePortions(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.GetSubshapeParagraphPortions
*/
type GetSubshapeParagraphPortionsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.MoveFile
*/
type MoveFileRequest struct {
    SrcPath string
    DestPath string
    SrcStorageName string
    DestStorageName string
    VersionId string
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

/* Request for SlidesApiService.MoveFolder
*/
type MoveFolderRequest struct {
    SrcPath string
    DestPath string
    SrcStorageName string
    DestStorageName string
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

/* Request for SlidesApiService.ObjectExists
*/
type ObjectExistsRequest struct {
    Path string
    StorageName string
    VersionId string
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

func (a *SlidesApiService) PostAddNewParagraph(request PostAddNewParagraphRequest) (IParagraph, *http.Response, error) {
	return a.CreateParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewParagraph
*/
type PostAddNewParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IParagraph
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNewPortion(request PostAddNewPortionRequest) (IPortion, *http.Response, error) {
	return a.CreatePortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewPortion
*/
type PostAddNewPortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNewShape(request PostAddNewShapeRequest) (IShapeBase, *http.Response, error) {
	return a.CreateShape(request.Name, request.SlideIndex, request.Dto, request.ShapeToClone, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewShape
*/
type PostAddNewShapeRequest struct {
    Name string
    SlideIndex int32
    Dto IShapeBase
    ShapeToClone *int32
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNewSubshape(request PostAddNewSubshapeRequest) (IShapeBase, *http.Response, error) {
	return a.CreateSubshape(request.Name, request.SlideIndex, request.Path, request.Dto, request.ShapeToClone, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewSubshape
*/
type PostAddNewSubshapeRequest struct {
    Name string
    SlideIndex int32
    Path string
    Dto IShapeBase
    ShapeToClone *int32
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNewSubshapeParagraph(request PostAddNewSubshapeParagraphRequest) (IParagraph, *http.Response, error) {
	return a.CreateSubshapeParagraph(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewSubshapeParagraph
*/
type PostAddNewSubshapeParagraphRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Dto IParagraph
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNewSubshapePortion(request PostAddNewSubshapePortionRequest) (IPortion, *http.Response, error) {
	return a.CreateSubshapePortion(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNewSubshapePortion
*/
type PostAddNewSubshapePortionRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostAddNotesSlide(request PostAddNotesSlideRequest) (INotesSlide, *http.Response, error) {
	return a.CreateNotesSlide(request.Name, request.SlideIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAddNotesSlide
*/
type PostAddNotesSlideRequest struct {
    Name string
    SlideIndex int32
    Dto INotesSlide
    Password string
    Folder string
    Storage string
}

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

func (a *SlidesApiService) PostAlignShapes(request PostAlignShapesRequest) (IShapes, *http.Response, error) {
	return a.AlignShapes(request.Name, request.SlideIndex, request.AlignmentType, request.AlignToSlide, request.Shapes, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostAlignShapes
*/
type PostAlignShapesRequest struct {
    Name string
    SlideIndex int32
    AlignmentType string
    AlignToSlide *bool
    Shapes []int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostChartCategory(request PostChartCategoryRequest) (IChart, *http.Response, error) {
	return a.CreateChartCategory(request.Name, request.SlideIndex, request.ShapeIndex, request.Category, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostChartCategory
*/
type PostChartCategoryRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Category IChartCategory
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostChartDataPoint(request PostChartDataPointRequest) (IChart, *http.Response, error) {
	return a.CreateChartDataPoint(request.Name, request.SlideIndex, request.ShapeIndex, request.SeriesIndex, request.DataPoint, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostChartDataPoint
*/
type PostChartDataPointRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    SeriesIndex int32
    DataPoint IDataPoint
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostChartSeries(request PostChartSeriesRequest) (IChart, *http.Response, error) {
	return a.CreateChartSeries(request.Name, request.SlideIndex, request.ShapeIndex, request.Series, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostChartSeries
*/
type PostChartSeriesRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Series ISeries
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostCopyLayoutSlideFromSourcePresentation(request PostCopyLayoutSlideFromSourcePresentationRequest) (ILayoutSlide, *http.Response, error) {
	return a.CopyLayoutSlide(request.Name, request.CloneFrom, request.CloneFromPosition, request.CloneFromPassword, request.CloneFromStorage, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostCopyLayoutSlideFromSourcePresentation
*/
type PostCopyLayoutSlideFromSourcePresentationRequest struct {
    Name string
    CloneFrom string
    CloneFromPosition int32
    CloneFromPassword string
    CloneFromStorage string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostCopyMasterSlideFromSourcePresentation(request PostCopyMasterSlideFromSourcePresentationRequest) (IMasterSlide, *http.Response, error) {
	return a.CopyMasterSlide(request.Name, request.CloneFrom, request.CloneFromPosition, request.CloneFromPassword, request.CloneFromStorage, request.ApplyToAll, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostCopyMasterSlideFromSourcePresentation
*/
type PostCopyMasterSlideFromSourcePresentationRequest struct {
    Name string
    CloneFrom string
    CloneFromPosition int32
    CloneFromPassword string
    CloneFromStorage string
    ApplyToAll *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostExportImageWithDefaultFormat(request PostExportImageWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImageDefaultFormatOnline(request.Document, request.Index, request.Password)
}

/* Request for SlidesApiService.PostExportImageWithDefaultFormat
*/
type PostExportImageWithDefaultFormatRequest struct {
    Document []byte
    Index int32
    Password string
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

func (a *SlidesApiService) PostExportImageWithFormat(request PostExportImageWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImageOnline(request.Document, request.Index, request.Format, request.Password)
}

/* Request for SlidesApiService.PostExportImageWithFormat
*/
type PostExportImageWithFormatRequest struct {
    Document []byte
    Index int32
    Format string
    Password string
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

func (a *SlidesApiService) PostExportImagesFromRequestWithFormat(request PostExportImagesFromRequestWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImagesOnline(request.Document, request.Format, request.Password)
}

/* Request for SlidesApiService.PostExportImagesFromRequestWithFormat
*/
type PostExportImagesFromRequestWithFormatRequest struct {
    Document []byte
    Format string
    Password string
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

func (a *SlidesApiService) PostExportImagesWithDefaultFormat(request PostExportImagesWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImagesDefaultFormat(request.Name, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostExportImagesWithDefaultFormat
*/
type PostExportImagesWithDefaultFormatRequest struct {
    Name string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostExportImagesWithFormat(request PostExportImagesWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImages(request.Name, request.Format, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostExportImagesWithFormat
*/
type PostExportImagesWithFormatRequest struct {
    Name string
    Format string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostExportShape(request PostExportShapeRequest) (*os.File, *http.Response, error) {
	return a.DownloadShapeOnline(request.Document, request.SlideIndex, request.ShapeIndex, request.Format, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostExportShape
*/
type PostExportShapeRequest struct {
    Document []byte
    SlideIndex int32
    ShapeIndex int32
    Format string
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostExportSlide(request PostExportSlideRequest) (*os.File, *http.Response, error) {
	return a.DownloadSlideOnline(request.Document, request.SlideIndex, request.Format, request.Width, request.Height, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostExportSlide
*/
type PostExportSlideRequest struct {
    Document []byte
    SlideIndex int32
    Format string
    Width *int32
    Height *int32
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostGetNotesSlide(request PostGetNotesSlideRequest) (INotesSlide, *http.Response, error) {
	return a.GetNotesSlideOnline(request.Document, request.SlideIndex, request.Password)
}

/* Request for SlidesApiService.PostGetNotesSlide
*/
type PostGetNotesSlideRequest struct {
    Document []byte
    SlideIndex int32
    Password string
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

func (a *SlidesApiService) PostGetNotesSlideExists(request PostGetNotesSlideExistsRequest) (IEntityExists, *http.Response, error) {
	return a.NotesSlideExistsOnline(request.Document, request.SlideIndex, request.Password)
}

/* Request for SlidesApiService.PostGetNotesSlideExists
*/
type PostGetNotesSlideExistsRequest struct {
    Document []byte
    SlideIndex int32
    Password string
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

func (a *SlidesApiService) PostGetNotesSlideWithFormat(request PostGetNotesSlideWithFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadNotesSlideOnline(request.Document, request.SlideIndex, request.Format, request.Width, request.Height, request.Password, request.FontsFolder)
}

/* Request for SlidesApiService.PostGetNotesSlideWithFormat
*/
type PostGetNotesSlideWithFormatRequest struct {
    Document []byte
    SlideIndex int32
    Format string
    Width *int32
    Height *int32
    Password string
    FontsFolder string
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

func (a *SlidesApiService) PostImageWatermark(request PostImageWatermarkRequest) (*http.Response, error) {
	return a.CreateImageWatermark(request.Name, request.Image, request.PictureFrame, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostImageWatermark
*/
type PostImageWatermarkRequest struct {
    Name string
    Image []byte
    PictureFrame IPictureFrame
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostImageWatermarkOnline(request PostImageWatermarkOnlineRequest) (*os.File, *http.Response, error) {
	return a.CreateImageWatermarkOnline(request.Document, request.Image, request.PictureFrame, request.Password)
}

/* Request for SlidesApiService.PostImageWatermarkOnline
*/
type PostImageWatermarkOnlineRequest struct {
    Document []byte
    Image []byte
    PictureFrame IPictureFrame
    Password string
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

func (a *SlidesApiService) PostImagesFromRequestWithDefaultFormat(request PostImagesFromRequestWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	return a.DownloadImagesDefaultFormatOnline(request.Document, request.Password)
}

/* Request for SlidesApiService.PostImagesFromRequestWithDefaultFormat
*/
type PostImagesFromRequestWithDefaultFormatRequest struct {
    Document []byte
    Password string
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

func (a *SlidesApiService) PostNotesSlideAddNewParagraph(request PostNotesSlideAddNewParagraphRequest) (IParagraph, *http.Response, error) {
	return a.CreateNotesSlideParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostNotesSlideAddNewParagraph
*/
type PostNotesSlideAddNewParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IParagraph
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostNotesSlideAddNewPortion(request PostNotesSlideAddNewPortionRequest) (IPortion, *http.Response, error) {
	return a.CreateNotesSlidePortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostNotesSlideAddNewPortion
*/
type PostNotesSlideAddNewPortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostNotesSlideAddNewShape(request PostNotesSlideAddNewShapeRequest) (IShapeBase, *http.Response, error) {
	return a.CreateNotesSlideShape(request.Name, request.SlideIndex, request.Dto, request.ShapeToClone, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostNotesSlideAddNewShape
*/
type PostNotesSlideAddNewShapeRequest struct {
    Name string
    SlideIndex int32
    Dto IShapeBase
    ShapeToClone *int32
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostNotesSlideShapeSaveAs(request PostNotesSlideShapeSaveAsRequest) (*os.File, *http.Response, error) {
	return a.DownloadNotesSlideShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Format, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostNotesSlideShapeSaveAs
*/
type PostNotesSlideShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostPresentationMerge(request PostPresentationMergeRequest) (IDocument, *http.Response, error) {
	return a.Merge(request.Name, request.Request, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostPresentationMerge
*/
type PostPresentationMergeRequest struct {
    Name string
    Request IPresentationsMergeRequest
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostPresentationReplaceText(request PostPresentationReplaceTextRequest) (*os.File, *http.Response, error) {
	return a.ReplacePresentationTextOnline(request.Document, request.OldValue, request.NewValue, request.IgnoreCase, request.Password)
}

/* Request for SlidesApiService.PostPresentationReplaceText
*/
type PostPresentationReplaceTextRequest struct {
    Document []byte
    OldValue string
    NewValue string
    IgnoreCase *bool
    Password string
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

func (a *SlidesApiService) PostPresentationSplit(request PostPresentationSplitRequest) (*os.File, *http.Response, error) {
	return a.SplitOnline(request.Document, request.Format, request.Width, request.Height, request.From, request.To, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostPresentationSplit
*/
type PostPresentationSplitRequest struct {
    Document []byte
    Format string
    Width *int32
    Height *int32
    From *int32
    To *int32
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSection(request PostSectionRequest) (ISections, *http.Response, error) {
	return a.CreateSection(request.Name, request.SectionName, request.SlideIndex, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSection
*/
type PostSectionRequest struct {
    Name string
    SectionName string
    SlideIndex int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSectionMove(request PostSectionMoveRequest) (ISections, *http.Response, error) {
	return a.MoveSection(request.Name, request.SectionIndex, request.NewPosition, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSectionMove
*/
type PostSectionMoveRequest struct {
    Name string
    SectionIndex int32
    NewPosition int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostShapeSaveAs(request PostShapeSaveAsRequest) (*os.File, *http.Response, error) {
	return a.DownloadShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Format, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostShapeSaveAs
*/
type PostShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSlideAnimationEffect(request PostSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.CreateAnimationEffect(request.Name, request.SlideIndex, request.Effect, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlideAnimationEffect
*/
type PostSlideAnimationEffectRequest struct {
    Name string
    SlideIndex int32
    Effect IEffect
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlideAnimationInteractiveSequence(request PostSlideAnimationInteractiveSequenceRequest) (ISlideAnimation, *http.Response, error) {
	return a.CreateAnimationInteractiveSequence(request.Name, request.SlideIndex, request.Sequence, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlideAnimationInteractiveSequence
*/
type PostSlideAnimationInteractiveSequenceRequest struct {
    Name string
    SlideIndex int32
    Sequence IInteractiveSequence
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlideAnimationInteractiveSequenceEffect(request PostSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.CreateAnimationInteractiveSequenceEffect(request.Name, request.SlideIndex, request.SequenceIndex, request.Effect, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlideAnimationInteractiveSequenceEffect
*/
type PostSlideAnimationInteractiveSequenceEffectRequest struct {
    Name string
    SlideIndex int32
    SequenceIndex int32
    Effect IEffect
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlideReplaceText(request PostSlideReplaceTextRequest) (*os.File, *http.Response, error) {
	return a.ReplaceSlideTextOnline(request.Document, request.SlideIndex, request.OldValue, request.NewValue, request.IgnoreCase, request.Password)
}

/* Request for SlidesApiService.PostSlideReplaceText
*/
type PostSlideReplaceTextRequest struct {
    Document []byte
    SlideIndex int32
    OldValue string
    NewValue string
    IgnoreCase *bool
    Password string
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

func (a *SlidesApiService) PostSlideSaveAs(request PostSlideSaveAsRequest) (*os.File, *http.Response, error) {
	return a.DownloadSlide(request.Name, request.SlideIndex, request.Format, request.Options, request.Width, request.Height, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostSlideSaveAs
*/
type PostSlideSaveAsRequest struct {
    Name string
    SlideIndex int32
    Format string
    Options IExportOptions
    Width *int32
    Height *int32
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSlidesAdd(request PostSlidesAddRequest) (ISlides, *http.Response, error) {
	return a.CreateSlide(request.Name, request.LayoutAlias, request.Position, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesAdd
*/
type PostSlidesAddRequest struct {
    Name string
    LayoutAlias string
    Position *int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesConvert(request PostSlidesConvertRequest) (*os.File, *http.Response, error) {
	return a.Convert(request.Document, request.Format, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostSlidesConvert
*/
type PostSlidesConvertRequest struct {
    Document []byte
    Format string
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSlidesCopy(request PostSlidesCopyRequest) (ISlides, *http.Response, error) {
	return a.CopySlide(request.Name, request.SlideToCopy, request.Position, request.Source, request.SourcePassword, request.SourceStorage, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesCopy
*/
type PostSlidesCopyRequest struct {
    Name string
    SlideToCopy int32
    Position *int32
    Source string
    SourcePassword string
    SourceStorage string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesDocument(request PostSlidesDocumentRequest) (IDocument, *http.Response, error) {
	return a.CreatePresentation(request.Name, request.Data, request.InputPassword, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesDocument
*/
type PostSlidesDocumentRequest struct {
    Name string
    Data []byte
    InputPassword string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesDocumentFromHtml(request PostSlidesDocumentFromHtmlRequest) (IDocument, *http.Response, error) {
	return a.ImportFromHtml(request.Name, request.Html, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesDocumentFromHtml
*/
type PostSlidesDocumentFromHtmlRequest struct {
    Name string
    Html string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesDocumentFromPdf(request PostSlidesDocumentFromPdfRequest) (IDocument, *http.Response, error) {
	return a.ImportFromPdf(request.Name, request.Pdf, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesDocumentFromPdf
*/
type PostSlidesDocumentFromPdfRequest struct {
    Name string
    Pdf []byte
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesDocumentFromSource(request PostSlidesDocumentFromSourceRequest) (IDocument, *http.Response, error) {
	return a.CreatePresentationFromSource(request.Name, request.SourcePath, request.SourcePassword, request.SourceStorage, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesDocumentFromSource
*/
type PostSlidesDocumentFromSourceRequest struct {
    Name string
    SourcePath string
    SourcePassword string
    SourceStorage string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesDocumentFromTemplate(request PostSlidesDocumentFromTemplateRequest) (IDocument, *http.Response, error) {
	return a.CreatePresentationFromTemplate(request.Name, request.TemplatePath, request.Data, request.TemplatePassword, request.TemplateStorage, request.IsImageDataEmbedded, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesDocumentFromTemplate
*/
type PostSlidesDocumentFromTemplateRequest struct {
    Name string
    TemplatePath string
    Data string
    TemplatePassword string
    TemplateStorage string
    IsImageDataEmbedded *bool
    Password string
    Folder string
    Storage string
}

/* SlidesApiService Merges the presentation with other presentations or some of their slides specified in the request parameter. Returns result file in the response. 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) Merge request.
     @param "password" (string) Document password.
 @return *os.File*/
func (a *SlidesApiService) MergeOnline(files [][]byte, request IOrderedMergeRequest, password string) (*os.File, *http.Response, error) {
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

func (a *SlidesApiService) PostSlidesMerge(request PostSlidesMergeRequest) (*os.File, *http.Response, error) {
	return a.MergeOnline(request.Files, request.Request, request.Password)
}

/* Request for SlidesApiService.PostSlidesMerge
*/
type PostSlidesMergeRequest struct {
    Files [][]byte
    Request IOrderedMergeRequest
    Password string
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

func (a *SlidesApiService) PostSlidesPipeline(request PostSlidesPipelineRequest) (*os.File, *http.Response, error) {
	return a.Pipeline(request.Pipeline, request.Files)
}

/* Request for SlidesApiService.PostSlidesPipeline
*/
type PostSlidesPipelineRequest struct {
    Pipeline IPipeline
    Files [][]byte
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

func (a *SlidesApiService) PostSlidesPresentationReplaceText(request PostSlidesPresentationReplaceTextRequest) (IDocumentReplaceResult, *http.Response, error) {
	return a.ReplacePresentationText(request.Name, request.OldValue, request.NewValue, request.IgnoreCase, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesPresentationReplaceText
*/
type PostSlidesPresentationReplaceTextRequest struct {
    Name string
    OldValue string
    NewValue string
    IgnoreCase *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesReorder(request PostSlidesReorderRequest) (ISlides, *http.Response, error) {
	return a.MoveSlide(request.Name, request.SlideIndex, request.NewPosition, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesReorder
*/
type PostSlidesReorderRequest struct {
    Name string
    SlideIndex int32
    NewPosition int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesReorderMany(request PostSlidesReorderManyRequest) (ISlides, *http.Response, error) {
	return a.ReorderSlides(request.Name, request.OldPositions, request.NewPositions, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesReorderMany
*/
type PostSlidesReorderManyRequest struct {
    Name string
    OldPositions []int32
    NewPositions []int32
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesSaveAs(request PostSlidesSaveAsRequest) (*os.File, *http.Response, error) {
	return a.DownloadPresentation(request.Name, request.Format, request.Options, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostSlidesSaveAs
*/
type PostSlidesSaveAsRequest struct {
    Name string
    Format string
    Options IExportOptions
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSlidesSetDocumentProperties(request PostSlidesSetDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	return a.SetDocumentProperties(request.Name, request.Properties, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesSetDocumentProperties
*/
type PostSlidesSetDocumentPropertiesRequest struct {
    Name string
    Properties IDocumentProperties
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesSlideReplaceText(request PostSlidesSlideReplaceTextRequest) (ISlideReplaceResult, *http.Response, error) {
	return a.ReplaceSlideText(request.Name, request.SlideIndex, request.OldValue, request.NewValue, request.IgnoreCase, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostSlidesSlideReplaceText
*/
type PostSlidesSlideReplaceTextRequest struct {
    Name string
    SlideIndex int32
    OldValue string
    NewValue string
    IgnoreCase *bool
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostSlidesSplit(request PostSlidesSplitRequest) (ISplitDocumentResult, *http.Response, error) {
	return a.Split(request.Name, request.Options, request.Format, request.Width, request.Height, request.From, request.To, request.DestFolder, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostSlidesSplit
*/
type PostSlidesSplitRequest struct {
    Name string
    Options IExportOptions
    Format string
    Width *int32
    Height *int32
    From *int32
    To *int32
    DestFolder string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostSubshapeSaveAs(request PostSubshapeSaveAsRequest) (*os.File, *http.Response, error) {
	return a.DownloadSubshape(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Format, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PostSubshapeSaveAs
*/
type PostSubshapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PostWatermark(request PostWatermarkRequest) (*http.Response, error) {
	return a.CreateWatermark(request.Name, request.Shape, request.FontHeight, request.Text, request.FontName, request.FontColor, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PostWatermark
*/
type PostWatermarkRequest struct {
    Name string
    Shape IShape
    FontHeight *float64
    Text string
    FontName string
    FontColor string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PostWatermarkDeleteOnline(request PostWatermarkDeleteOnlineRequest) (*os.File, *http.Response, error) {
	return a.DeleteWatermarkOnline(request.Document, request.ShapeName, request.Password)
}

/* Request for SlidesApiService.PostWatermarkDeleteOnline
*/
type PostWatermarkDeleteOnlineRequest struct {
    Document []byte
    ShapeName string
    Password string
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

func (a *SlidesApiService) PostWatermarkOnline(request PostWatermarkOnlineRequest) (*os.File, *http.Response, error) {
	return a.CreateWatermarkOnline(request.Document, request.Shape, request.FontHeight, request.Text, request.FontName, request.FontColor, request.Password)
}

/* Request for SlidesApiService.PostWatermarkOnline
*/
type PostWatermarkOnlineRequest struct {
    Document []byte
    Shape IShape
    FontHeight *float64
    Text string
    FontName string
    FontColor string
    Password string
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

func (a *SlidesApiService) PutChartCategory(request PutChartCategoryRequest) (IChart, *http.Response, error) {
	return a.UpdateChartCategory(request.Name, request.SlideIndex, request.ShapeIndex, request.CategoryIndex, request.Category, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutChartCategory
*/
type PutChartCategoryRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    CategoryIndex int32
    Category IChartCategory
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutChartDataPoint(request PutChartDataPointRequest) (IChart, *http.Response, error) {
	return a.UpdateChartDataPoint(request.Name, request.SlideIndex, request.ShapeIndex, request.SeriesIndex, request.PointIndex, request.DataPoint, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutChartDataPoint
*/
type PutChartDataPointRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    SeriesIndex int32
    PointIndex int32
    DataPoint IDataPoint
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutChartSeries(request PutChartSeriesRequest) (IChart, *http.Response, error) {
	return a.UpdateChartSeries(request.Name, request.SlideIndex, request.ShapeIndex, request.SeriesIndex, request.Series, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutChartSeries
*/
type PutChartSeriesRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    SeriesIndex int32
    Series ISeries
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutExportShape(request PutExportShapeRequest) (*http.Response, error) {
	return a.SaveShapeOnline(request.Document, request.SlideIndex, request.ShapeIndex, request.Format, request.OutPath, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutExportShape
*/
type PutExportShapeRequest struct {
    Document []byte
    SlideIndex int32
    ShapeIndex int32
    Format string
    OutPath string
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutExportSlide(request PutExportSlideRequest) (*http.Response, error) {
	return a.SaveSlideOnline(request.Document, request.SlideIndex, request.Format, request.OutPath, request.Width, request.Height, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutExportSlide
*/
type PutExportSlideRequest struct {
    Document []byte
    SlideIndex int32
    Format string
    OutPath string
    Width *int32
    Height *int32
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutLayoutSlide(request PutLayoutSlideRequest) (ILayoutSlide, *http.Response, error) {
	return a.UpdateLayoutSlide(request.Name, request.SlideIndex, request.SlideDto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutLayoutSlide
*/
type PutLayoutSlideRequest struct {
    Name string
    SlideIndex int32
    SlideDto ILayoutSlide
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutNotesSlideHeaderFooter(request PutNotesSlideHeaderFooterRequest) (INotesSlideHeaderFooter, *http.Response, error) {
	return a.SetNotesSlideHeaderFooter(request.Name, request.SlideIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutNotesSlideHeaderFooter
*/
type PutNotesSlideHeaderFooterRequest struct {
    Name string
    SlideIndex int32
    Dto INotesSlideHeaderFooter
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutNotesSlideShapeSaveAs(request PutNotesSlideShapeSaveAsRequest) (*http.Response, error) {
	return a.SaveNotesSlideShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Format, request.OutPath, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutNotesSlideShapeSaveAs
*/
type PutNotesSlideShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    OutPath string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutPresentationMerge(request PutPresentationMergeRequest) (IDocument, *http.Response, error) {
	return a.OrderedMerge(request.Name, request.Request, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutPresentationMerge
*/
type PutPresentationMergeRequest struct {
    Name string
    Request IOrderedMergeRequest
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutPresentationSplit(request PutPresentationSplitRequest) (ISplitDocumentResult, *http.Response, error) {
	return a.SplitAndSaveOnline(request.Document, request.Format, request.DestFolder, request.Width, request.Height, request.From, request.To, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutPresentationSplit
*/
type PutPresentationSplitRequest struct {
    Document []byte
    Format string
    DestFolder string
    Width *int32
    Height *int32
    From *int32
    To *int32
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutSection(request PutSectionRequest) (ISections, *http.Response, error) {
	return a.UpdateSection(request.Name, request.SectionIndex, request.SectionName, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSection
*/
type PutSectionRequest struct {
    Name string
    SectionIndex int32
    SectionName string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSections(request PutSectionsRequest) (ISections, *http.Response, error) {
	return a.SetSections(request.Name, request.Sections, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSections
*/
type PutSectionsRequest struct {
    Name string
    Sections ISections
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSetParagraphPortionProperties(request PutSetParagraphPortionPropertiesRequest) (IPortion, *http.Response, error) {
	return a.UpdatePortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSetParagraphPortionProperties
*/
type PutSetParagraphPortionPropertiesRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSetParagraphProperties(request PutSetParagraphPropertiesRequest) (IParagraph, *http.Response, error) {
	return a.UpdateParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSetParagraphProperties
*/
type PutSetParagraphPropertiesRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSetSubshapeParagraphPortionProperties(request PutSetSubshapeParagraphPortionPropertiesRequest) (IPortion, *http.Response, error) {
	return a.UpdateSubshapePortion(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSetSubshapeParagraphPortionProperties
*/
type PutSetSubshapeParagraphPortionPropertiesRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSetSubshapeParagraphProperties(request PutSetSubshapeParagraphPropertiesRequest) (IParagraph, *http.Response, error) {
	return a.UpdateSubshapeParagraph(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSetSubshapeParagraphProperties
*/
type PutSetSubshapeParagraphPropertiesRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutShapeSaveAs(request PutShapeSaveAsRequest) (*http.Response, error) {
	return a.SaveShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Format, request.OutPath, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutShapeSaveAs
*/
type PutShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    OutPath string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutSlideAnimation(request PutSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	return a.SetAnimation(request.Name, request.SlideIndex, request.Animation, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideAnimation
*/
type PutSlideAnimationRequest struct {
    Name string
    SlideIndex int32
    Animation ISlideAnimation
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlideAnimationEffect(request PutSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.UpdateAnimationEffect(request.Name, request.SlideIndex, request.EffectIndex, request.Effect, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideAnimationEffect
*/
type PutSlideAnimationEffectRequest struct {
    Name string
    SlideIndex int32
    EffectIndex int32
    Effect IEffect
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlideAnimationInteractiveSequenceEffect(request PutSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	return a.UpdateAnimationInteractiveSequenceEffect(request.Name, request.SlideIndex, request.SequenceIndex, request.EffectIndex, request.Effect, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideAnimationInteractiveSequenceEffect
*/
type PutSlideAnimationInteractiveSequenceEffectRequest struct {
    Name string
    SlideIndex int32
    SequenceIndex int32
    EffectIndex int32
    Effect IEffect
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlideHeaderFooter(request PutSlideHeaderFooterRequest) (IHeaderFooter, *http.Response, error) {
	return a.SetSlideHeaderFooter(request.Name, request.SlideIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideHeaderFooter
*/
type PutSlideHeaderFooterRequest struct {
    Name string
    SlideIndex int32
    Dto IHeaderFooter
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlideSaveAs(request PutSlideSaveAsRequest) (*http.Response, error) {
	return a.SaveSlide(request.Name, request.SlideIndex, request.Format, request.OutPath, request.Options, request.Width, request.Height, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutSlideSaveAs
*/
type PutSlideSaveAsRequest struct {
    Name string
    SlideIndex int32
    Format string
    OutPath string
    Options IExportOptions
    Width *int32
    Height *int32
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutSlideShapeInfo(request PutSlideShapeInfoRequest) (IShapeBase, *http.Response, error) {
	return a.UpdateShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideShapeInfo
*/
type PutSlideShapeInfoRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IShapeBase
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlideSubshapeInfo(request PutSlideSubshapeInfoRequest) (IShapeBase, *http.Response, error) {
	return a.UpdateSubshape(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlideSubshapeInfo
*/
type PutSlideSubshapeInfoRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Dto IShapeBase
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesConvert(request PutSlidesConvertRequest) (*http.Response, error) {
	return a.ConvertAndSave(request.Document, request.Format, request.OutPath, request.Password, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutSlidesConvert
*/
type PutSlidesConvertRequest struct {
    Document []byte
    Format string
    OutPath string
    Password string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutSlidesHeaderFooter(request PutSlidesHeaderFooterRequest) (IDocument, *http.Response, error) {
	return a.SetPresentationHeaderFooter(request.Name, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesHeaderFooter
*/
type PutSlidesHeaderFooterRequest struct {
    Name string
    Dto IHeaderFooter
    Password string
    Folder string
    Storage string
}

/* SlidesApiService Merges the presentation with other presentations or some of their slides specified in the request parameter. Puts result in the storage. 
 @param outPath Path to save result.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "files" ([][]byte) Files to merge
     @param "request" (OrderedMergeRequest) Merge request.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
 @return */
func (a *SlidesApiService) MergeAndSaveOnline(outPath string, files [][]byte, request IOrderedMergeRequest, password string, storage string) (*http.Response, error) {
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

	if err := typeCheckParameter(password, "string", "password"); err != nil {
		return nil, err
	}
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
	if localVarTempParam := password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
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

func (a *SlidesApiService) PutSlidesMerge(request PutSlidesMergeRequest) (*http.Response, error) {
	return a.MergeAndSaveOnline(request.OutPath, request.Files, request.Request, request.Password, request.Storage)
}

/* Request for SlidesApiService.PutSlidesMerge
*/
type PutSlidesMergeRequest struct {
    OutPath string
    Files [][]byte
    Request IOrderedMergeRequest
    Password string
    Storage string
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

func (a *SlidesApiService) PutSlidesProtectionProperties(request PutSlidesProtectionPropertiesRequest) (IProtectionProperties, *http.Response, error) {
	return a.SetProtection(request.Name, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesProtectionProperties
*/
type PutSlidesProtectionPropertiesRequest struct {
    Name string
    Dto IProtectionProperties
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesProtectionPropertiesOnline(request PutSlidesProtectionPropertiesOnlineRequest) (*os.File, *http.Response, error) {
	return a.SetProtectionOnline(request.Document, request.Dto, request.Password)
}

/* Request for SlidesApiService.PutSlidesProtectionPropertiesOnline
*/
type PutSlidesProtectionPropertiesOnlineRequest struct {
    Document []byte
    Dto IProtectionProperties
    Password string
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

func (a *SlidesApiService) PutSlidesSaveAs(request PutSlidesSaveAsRequest) (*http.Response, error) {
	return a.SavePresentation(request.Name, request.Format, request.OutPath, request.Options, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutSlidesSaveAs
*/
type PutSlidesSaveAsRequest struct {
    Name string
    Format string
    OutPath string
    Options IExportOptions
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutSlidesSetDocumentProperty(request PutSlidesSetDocumentPropertyRequest) (IDocumentProperty, *http.Response, error) {
	return a.SetDocumentProperty(request.Name, request.PropertyName, request.Property, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesSetDocumentProperty
*/
type PutSlidesSetDocumentPropertyRequest struct {
    Name string
    PropertyName string
    Property IDocumentProperty
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesSlide(request PutSlidesSlideRequest) (ISlide, *http.Response, error) {
	return a.UpdateSlide(request.Name, request.SlideIndex, request.SlideDto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesSlide
*/
type PutSlidesSlideRequest struct {
    Name string
    SlideIndex int32
    SlideDto ISlide
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesSlideBackground(request PutSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	return a.SetBackground(request.Name, request.SlideIndex, request.Background, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesSlideBackground
*/
type PutSlidesSlideBackgroundRequest struct {
    Name string
    SlideIndex int32
    Background ISlideBackground
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesSlideBackgroundColor(request PutSlidesSlideBackgroundColorRequest) (ISlideBackground, *http.Response, error) {
	return a.SetBackgroundColor(request.Name, request.SlideIndex, request.Color, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesSlideBackgroundColor
*/
type PutSlidesSlideBackgroundColorRequest struct {
    Name string
    SlideIndex int32
    Color string
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesSlideProperties(request PutSlidesSlidePropertiesRequest) (ISlideProperties, *http.Response, error) {
	return a.SetSlideProperties(request.Name, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesSlideProperties
*/
type PutSlidesSlidePropertiesRequest struct {
    Name string
    Dto ISlideProperties
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSlidesViewProperties(request PutSlidesViewPropertiesRequest) (IViewProperties, *http.Response, error) {
	return a.SetViewProperties(request.Name, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutSlidesViewProperties
*/
type PutSlidesViewPropertiesRequest struct {
    Name string
    Dto IViewProperties
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutSubshapeSaveAs(request PutSubshapeSaveAsRequest) (*http.Response, error) {
	return a.SaveSubshape(request.Name, request.SlideIndex, request.Path, request.ShapeIndex, request.Format, request.OutPath, request.Options, request.ScaleX, request.ScaleY, request.Bounds, request.Password, request.Folder, request.Storage, request.FontsFolder)
}

/* Request for SlidesApiService.PutSubshapeSaveAs
*/
type PutSubshapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Format string
    OutPath string
    Options IIShapeExportOptions
    ScaleX *float64
    ScaleY *float64
    Bounds string
    Password string
    Folder string
    Storage string
    FontsFolder string
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

func (a *SlidesApiService) PutUpdateNotesSlide(request PutUpdateNotesSlideRequest) (INotesSlide, *http.Response, error) {
	return a.UpdateNotesSlide(request.Name, request.SlideIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutUpdateNotesSlide
*/
type PutUpdateNotesSlideRequest struct {
    Name string
    SlideIndex int32
    Dto INotesSlide
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutUpdateNotesSlideShape(request PutUpdateNotesSlideShapeRequest) (IShapeBase, *http.Response, error) {
	return a.UpdateNotesSlideShape(request.Name, request.SlideIndex, request.ShapeIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutUpdateNotesSlideShape
*/
type PutUpdateNotesSlideShapeRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IShapeBase
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutUpdateNotesSlideShapeParagraph(request PutUpdateNotesSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	return a.UpdateNotesSlideParagraph(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutUpdateNotesSlideShapeParagraph
*/
type PutUpdateNotesSlideShapeParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
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

func (a *SlidesApiService) PutUpdateNotesSlideShapePortion(request PutUpdateNotesSlideShapePortionRequest) (IPortion, *http.Response, error) {
	return a.UpdateNotesSlidePortion(request.Name, request.SlideIndex, request.ShapeIndex, request.ParagraphIndex, request.PortionIndex, request.Dto, request.Password, request.Folder, request.Storage)
}

/* Request for SlidesApiService.PutUpdateNotesSlideShapePortion
*/
type PutUpdateNotesSlideShapePortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    PortionIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
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

/* Request for SlidesApiService.StorageExists
*/
type StorageExistsRequest struct {
    StorageName string
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

/* Request for SlidesApiService.UploadFile
*/
type UploadFileRequest struct {
    Path string
    File []byte
    StorageName string
}

